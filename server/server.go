package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
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

func (bs *BlogServer) registerTagHandlers(base *gin.RouterGroup, handlerFuncs ...gin.HandlerFunc) {
	tagGroup := base.Group("/tag")
	{
		tagGroup.GET("/", handlers.GetTagsHandler(bs.ctx, bs.localPostGetter))
		tagGroup.GET("/info", handlers.GetTagsInfoHandler(bs.ctx, bs.localPostGetter))
		tagGroup.GET("/info/:tag", handlers.GetSpecifyTagInfoHandler(bs.ctx, bs.localPostGetter))
	}
}

func (bs *BlogServer) registerGithubHandlers(base *gin.RouterGroup, handlerFuncs ...gin.HandlerFunc) {
	githubGroup := base.Group("/github")
	{
		githubGroup.POST("/hook", handlers.GithubHookHandler(bs.ctx, bs.localPostGetter, bs.githubGetter))
	}
}

func (bs *BlogServer) Handler(box *packr.Box) *gin.Engine {
	bs.engine.StaticFS("/blog", box)
	bs.engine.NoRoute(func(c *gin.Context) {
		file, err := box.Open("index.html")
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		defer file.Close()

		http.ServeContent(c.Writer, c.Request, "index.html", time.Now(), file)
	})

	apiGroup := bs.engine.Group("/api")
	bs.registerPostHandlers(apiGroup)
	bs.registerTagHandlers(apiGroup)
	if bs.autoHookFileChange {
		bs.registerGithubHandlers(apiGroup, middlewares.GithubVerifyMiddleware(bs.ctx, bs.secretToken))
	}

	return bs.engine
}
