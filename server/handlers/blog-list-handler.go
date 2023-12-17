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
		var pagination models.Pagination
		if err := c.ShouldBind(&pagination); err != nil {
			c.JSON(http.StatusBadRequest, models.PackResponseData(http.StatusBadRequest, "get blog list data failed", nil))
			return
		}
		lg.Infoc(ctx, "get blog list data, pagination: %+v", pagination)

		postList, _ := localPostGetter.GetPostList(ctx, pagination)
		c.JSON(http.StatusOK, models.PackResponseData(http.StatusOK, "get blog list data success", postList))
	}
}
