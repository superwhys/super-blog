package handlers

import (
	"context"

	"github.com/superwhys/superBlog/pkg/postmanager"
	"github.com/superwhys/venkit/lg"
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
