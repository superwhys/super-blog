package main

import (
	"context"

	"github.com/google/go-github/v53/github"
	"github.com/superwhys/superBlog/pkg/postgetter"
	"github.com/superwhys/superBlog/server"
	"github.com/yazl-tech/yazl/service"
	"github.com/yazl-tech/yazl/utils/flags"
	"golang.org/x/oauth2"
)

var (
	port        = flags.Int("port", 29915, "the server run port")
	githubToken = flags.String("githubToken", "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", "set the github api token")
	basePostDir = flags.String("basePostDir", "./blog-posts", "set the github api token")
)

func main() {
	flags.Parse()
	localGetter := postgetter.NewLocalPostGetter(basePostDir())

	tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken()}))
	client := github.NewClient(tc)

	blogSrv := server.NewBlogServer(localGetter)

	srv := service.NewYazlService(
		service.WithHTTPCORS(),
		service.WithPprof(),
		service.WithHttpHandler("/", blogSrv.Handler()),
	)
	srv.ListenAndServer(port())
}
