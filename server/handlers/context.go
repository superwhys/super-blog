package handlers

import (
	"context"

	"github.com/superwhys/goutils/lg"
	"github.com/superwhys/superBlog/pkg/postmanager"
)

type HandlerContext struct {
	ctx         context.Context
	postManager postmanager.PostManager
}

func NewContext(pm postmanager.PostManager) *HandlerContext {
	return &HandlerContext{
		ctx:         lg.With(context.Background(), "[BlogServer]"),
		postManager: pm,
	}
}
