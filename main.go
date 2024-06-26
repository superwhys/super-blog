package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/google/go-github/v53/github"
	"github.com/superwhys/super-blog/models"
	"github.com/superwhys/super-blog/pkg/postmanager"
	"github.com/superwhys/super-blog/server"
	"github.com/superwhys/venkit/lg"
	"github.com/superwhys/venkit/service"
	"github.com/superwhys/venkit/vflags"
	"github.com/superwhys/venkit/vgorm"
	"golang.org/x/oauth2"
)

var (
	port               = vflags.Int("port", 29915, "the server run port")
	githubToken        = vflags.String("githubToken", "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", "set the github api token")
	githubSecretToken  = vflags.String("githubSecretToken", "SjpvYt4nQVTZ", "set github webhook secret token")
	basePostDir        = vflags.String("basePostDir", "./blog-posts", "set the github api token")
	autoHookFileChange = vflags.Bool("autoHookFileChange", false, "Whether to automatically detect changes of file in github repository")
	mysqlConfig        = vflags.Struct("mysqlConfig", (*vgorm.MysqlConfig)(nil), "mysql config")
)

func main() {
	vflags.Parse()
	if !lg.IsDebug() {
		gin.SetMode(gin.ReleaseMode)
	}

	mysqlConf := new(vgorm.MysqlConfig)
	lg.PanicError(mysqlConfig(mysqlConf))
	db, err := mysqlConf.DialGorm()
	lg.PanicError(err)
	db.AutoMigrate(&models.BlogItem{}, &models.Tag{})

	var githubGetter *postmanager.GithubGetter
	postStore := postmanager.NewPostStore(db)

	if autoHookFileChange() {
		tc := oauth2.NewClient(
			context.Background(),
			oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken()}),
		)
		client := github.NewClient(tc)
		githubGetter = postmanager.NewGithubGetter(client)
	}

	box := packr.New("blogBox", "./blog-web/dist")
	_, err = box.FindString("index.html")
	lg.PanicError(err)

	blogSrv := server.NewBlogServer(
		postStore,
		githubGetter,
		githubSecretToken(),
		autoHookFileChange(),
	)

	srv := service.NewVkService(
		service.WithServiceName(vflags.GetServiceName()),
		service.WithHTTPCORS(),
		service.WithPprof(),
		service.WithHttpHandler("/", blogSrv.Handler(box)),
	)
	lg.PanicError(srv.Run(port()))
}
