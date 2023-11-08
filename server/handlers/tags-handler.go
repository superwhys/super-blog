package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/superwhys/goutils/lg"
	"github.com/superwhys/goutils/slices"
	"github.com/superwhys/superBlog/models"
	"github.com/superwhys/superBlog/pkg/postmanager"
)

func GetTagsHandler(ctx context.Context, localPostGetter *postmanager.LocalGetter) gin.HandlerFunc {
	ctx = lg.With(ctx, "[GetTagsHandler]")
	return func(c *gin.Context) {
		postList, err := localPostGetter.GetPostList(ctx)
		if err != nil {
			lg.Errorc(ctx, "get local post list error: %v", err)
			c.JSON(http.StatusInternalServerError, models.PackResponseData(http.StatusInternalServerError, "get tags failed", nil))
			return
		}

		tagsSet := slices.NewStringSet()
		for _, post := range postList.Items {
			tagsSet.Add(post.MetaData.Tags...)
		}
		tags := tagsSet.Slice()
		c.JSON(http.StatusOK, models.PackResponseData(http.StatusOK, "get tags success", tags))
	}
}
