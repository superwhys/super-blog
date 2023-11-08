package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/superwhys/goutils/lg"
	"github.com/superwhys/superBlog/models"
	"github.com/superwhys/superBlog/pkg/postmanager"
)

func BlogListHandler(ctx context.Context, localPostGetter *postmanager.LocalGetter) gin.HandlerFunc {
	ctx = lg.With(ctx, "[BlogListHandler]")
	return func(c *gin.Context) {
		postList, _ := localPostGetter.GetPostList(ctx)
		c.JSON(http.StatusOK, models.PackResponseData(http.StatusOK, "get blog list data success", postList))
	}
}
