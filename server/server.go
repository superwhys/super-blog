package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/superwhys/goutils/lg"
	"github.com/superwhys/superBlog/pkg/postmanager"
	"github.com/superwhys/superBlog/server/handlers"
	"github.com/superwhys/superBlog/server/handlers/middlewares"
)

type BlogServer struct {
	ctx                context.Context
	secretToken        string
	autoHookFileChange bool
	engine             *gin.Engine
	localPostGetter    *postmanager.LocalGetter
	githubGetter       *postmanager.GithubGetter
}

func NewBlogServer(localGetter *postmanager.LocalGetter, githubGetter *postmanager.GithubGetter, secretToken string, autoHookFileChange bool) *BlogServer {
	ctx := lg.With(context.Background(), "[BlogServer]")

	router := gin.Default()
	return &BlogServer{
		ctx:                ctx,
		engine:             router,
		secretToken:        secretToken,
		localPostGetter:    localGetter,
		githubGetter:       githubGetter,
		autoHookFileChange: autoHookFileChange,
	}
}

func (bs *BlogServer) registerPostHandlers(base *gin.RouterGroup, handlerFuncs ...gin.HandlerFunc) {
	postGroup := base.Group("/post")
	{
		postGroup.GET("/", handlers.BlogListHandler(bs.ctx, bs.localPostGetter))
		postGroup.GET("/:year/:month/:day/:name", handlers.PostHandler(bs.ctx, bs.localPostGetter))
	}
}

func (bs *BlogServer) registerGithubHandlers(base *gin.RouterGroup, handlerFuncs ...gin.HandlerFunc) {
	githubGroup := base.Group("/github")
	{
		githubGroup.POST("/hook", handlers.GithubHookHandler(bs.ctx, bs.localPostGetter, bs.githubGetter))
	}
}

func (bs *BlogServer) Handler() *gin.Engine {
	blogGroup := bs.engine.Group("/blog")
	bs.registerPostHandlers(blogGroup)
	if bs.autoHookFileChange {
		bs.registerGithubHandlers(blogGroup, middlewares.GithubVerifyMiddleware(bs.ctx, bs.secretToken))
	}

	return bs.engine
}
