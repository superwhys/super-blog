package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/superwhys/superBlog/models"
	"github.com/superwhys/venkit/lg"
)

func BlogListHandler(hctx *HandlerContext) gin.HandlerFunc {
	ctx := lg.With(hctx.ctx, "[BlogListHandler]")
	return func(c *gin.Context) {
		var pagination models.Pagination
		if err := c.ShouldBind(&pagination); err != nil {
			c.JSON(http.StatusBadRequest, models.PackResponseData(http.StatusBadRequest, "get blog list data failed", nil))
			return
		}
		lg.Infoc(ctx, "get blog list data, pagination: %+v", pagination)

		postList, err := hctx.postManager.GetPostList(ctx, pagination)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.PackResponseData(http.StatusBadRequest, "get blog list data failed", nil))
			return
		}
		postList.BackFillRowTags()
		postList.DecodeBlogItemContent()
		c.JSON(http.StatusOK, models.PackResponseData(http.StatusOK, "get blog list data success", postList))
	}
}
