package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/superwhys/superBlog/pkg/postgetter"
	"github.com/superwhys/superBlog/server/handlers"
	"github.com/yazl-tech/yazl/utils/lg"
)

type BlogServer struct {
	ctx    context.Context
	engine *gin.Engine
}

func NewBlogServer(postGetter postgetter.Getter) *BlogServer {
	ctx := lg.With(context.Background(), "[BlogServer]")

	router := gin.Default()
	blogGroup := router.Group("/blog")
	{
		blogGroup.GET("/posts", handlers.BlogListHandler(ctx, postGetter))
		blogGroup.GET("/post/:year/:month/:day/:name")
	}

	return &BlogServer{ctx: ctx, engine: router}
}

func (bs *BlogServer) Handler() *gin.Engine {
	return bs.engine
}
