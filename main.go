package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v53/github"
	"github.com/superwhys/superBlog/pkg/postmanager"
	"github.com/superwhys/superBlog/server"
	"github.com/yazl-tech/yazl/service"
	"github.com/yazl-tech/yazl/utils/flags"
	"github.com/yazl-tech/yazl/utils/lg"
	"golang.org/x/oauth2"
)

var (
	port               = flags.Int("port", 29915, "the server run port")
	githubToken        = flags.String("githubToken", "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", "set the github api token")
	githubSecretToken  = flags.String("githubSecretToken", "SjpvYt4nQVTZ", "set github webhook secret token")
	basePostDir        = flags.String("basePostDir", "./blog-posts", "set the github api token")
	autoHookFileChange = flags.Bool("autoHookFileChange", false, "Whether to automatically detect changes of file in github repository")
)

func main() {
	flags.Parse()

	if !lg.IsDebug() {
		gin.SetMode(gin.ReleaseMode)
	}

	var githubGetter *postmanager.GithubGetter
	localGetter := postmanager.NewLocalPostGetter(basePostDir())

	if autoHookFileChange() {
		tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken()}))
		client := github.NewClient(tc)
		githubGetter = postmanager.NewGithubGetter(client)
	}

	blogSrv := server.NewBlogServer(localGetter, githubGetter, githubSecretToken(), autoHookFileChange())

	srv := service.NewYazlService(
		service.WithHTTPCORS(),
		service.WithPprof(),
		service.WithHttpHandler("/", blogSrv.Handler()),
	)
	lg.PanicError(srv.ListenAndServer(port()))
}
