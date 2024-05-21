package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/superwhys/superBlog/pkg/postmanager"
	"github.com/superwhys/superBlog/server/handlers"
	"github.com/superwhys/superBlog/server/handlers/middlewares"
	"github.com/superwhys/venkit/lg"
	"github.com/superwhys/venkit/vgin"
)

type BlogServer struct {
	ctx                *handlers.HandlerContext
	secretToken        string
	autoHookFileChange bool
	engine             *gin.Engine
	postManager        postmanager.PostManager
	githubGetter       *postmanager.GithubGetter
}

func NewBlogServer(postManager postmanager.PostManager, githubGetter *postmanager.GithubGetter, secretToken string, autoHookFileChange bool) *BlogServer {
	ctx := handlers.NewContext(postManager)
	router := vgin.NewGinEngine()
	return &BlogServer{
		ctx:                ctx,
		engine:             router,
		secretToken:        secretToken,
		postManager:        postManager,
		githubGetter:       githubGetter,
		autoHookFileChange: autoHookFileChange,
	}
}

func (bs *BlogServer) CheckPosts() error {
	ctx := context.Background()
	count, err := bs.postManager.GetTotalPostCount(ctx)
	if err != nil {
		return err
	}

	if count != 0 {
		return nil
	}

	posts, err := bs.githubGetter.GetPostList(ctx)
	if err != nil {
		return err
	}

	for _, post := range posts {
		postItem, err := bs.githubGetter.GetSpecifyPost(ctx, post)
		if err != nil {
			lg.Errorf("get post: %v error: %v", post, err)
			continue
		}
		bs.postManager.UpdatePost(ctx, postItem)
	}

	return nil
}

func (bs *BlogServer) registerPostHandlers(base *gin.RouterGroup, handlerFuncs ...gin.HandlerFunc) {
	postGroup := base.Group("/post", handlerFuncs...)
	{
		postGroup.GET("/", handlers.BlogListHandler(bs.ctx))
		postGroup.GET("/:year/:month/:day/:name", handlers.PostHandler(bs.ctx))
	}
}

func (bs *BlogServer) registerTagHandlers(base *gin.RouterGroup, handlerFuncs ...gin.HandlerFunc) {
	tagGroup := base.Group("/tag", handlerFuncs...)
	{
		tagGroup.GET("/", handlers.GetTagsHandler(bs.ctx))
		tagGroup.GET("/info", handlers.GetTagsInfoHandler(bs.ctx))
		tagGroup.GET("/info/:tag", handlers.GetSpecifyTagInfoHandler(bs.ctx))
	}
}

func (bs *BlogServer) registerGithubHandlers(base *gin.RouterGroup, handlerFuncs ...gin.HandlerFunc) {
	githubGroup := base.Group("/github", handlerFuncs...)
	{
		githubGroup.POST("/hook", handlers.GithubHookHandler(bs.ctx, bs.githubGetter))
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
		bs.registerGithubHandlers(apiGroup, middlewares.GithubVerifyMiddleware(bs.secretToken))
	}

	return bs.engine
}
