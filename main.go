package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/google/go-github/v53/github"
	"github.com/superwhys/superBlog/models"
	"github.com/superwhys/superBlog/pkg/postmanager"
	"github.com/superwhys/superBlog/server"
	"github.com/superwhys/venkit/dialer"
	"github.com/superwhys/venkit/lg"
	"github.com/superwhys/venkit/service"
	"github.com/superwhys/venkit/vflags"
	"golang.org/x/oauth2"
)

var (
	port               = vflags.Int("port", 29915, "the server run port")
	githubToken        = vflags.String("githubToken", "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", "set the github api token")
	githubSecretToken  = vflags.String("githubSecretToken", "SjpvYt4nQVTZ", "set github webhook secret token")
	basePostDir        = vflags.String("basePostDir", "./blog-posts", "set the github api token")
	autoHookFileChange = vflags.Bool("autoHookFileChange", false, "Whether to automatically detect changes of file in github repository")
	mysqlUser          = vflags.String("mysqlUser", "yong", "mysql user name")
	mysqlPwd           = vflags.String("mysqlPwd", "kBcCfmwHFQtL", "mysql password")
	mysqlDbName        = vflags.String("mysqlDbName", "blog", "mysql db name")
)

func main() {
	vflags.Parse()
	if !lg.IsDebug() {
		gin.SetMode(gin.ReleaseMode)
	}
	db, err := dialer.DialMysqlGorm("mysql", dialer.WithAuth(mysqlUser(), mysqlPwd()), dialer.WithDBName(mysqlDbName()))
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
	if autoHookFileChange() {
		lg.PanicError(blogSrv.CheckPosts())
	}

	srv := service.NewVkService(
		service.WithServiceName(vflags.GetServiceName()),
		service.WithHTTPCORS(),
		service.WithPprof(),
		service.WithHttpHandler("/", blogSrv.Handler(box)),
	)
	lg.PanicError(srv.Run(port()))
}
