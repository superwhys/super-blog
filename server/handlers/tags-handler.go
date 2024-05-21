package handlers

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/superwhys/superBlog/models"
	"github.com/superwhys/venkit/lg"
	"github.com/superwhys/venkit/slices"
)

func GetTagsHandler(hctx *HandlerContext) gin.HandlerFunc {
	ctx := lg.With(hctx.ctx, "[GetTagsHandler]")
	return func(c *gin.Context) {
		postList, err := hctx.postManager.GetPostList(ctx, models.Pagination{Page: -1, Size: -1})
		if err != nil {
			lg.Errorc(ctx, "get local post list error: %v", err)
			c.JSON(
				http.StatusInternalServerError,
				models.PackResponseData(http.StatusInternalServerError, "get tags failed", nil),
			)
			return
		}
		tagsSet := slices.NewStringSet()
		for _, post := range postList.Items {

			tagsSet.Add(post.MetaData.GetTagsList()...)
		}
		tags := tagsSet.Slice()

		var tagItems []*models.TagItem
		for _, tag := range tags {
			tagItems = append(tagItems, &models.TagItem{
				Tag:        tag,
				ToEndPoint: fmt.Sprintf("/tag/%v", tag),
			})
		}

		sort.Slice(tagItems, func(i, j int) bool {
			return tagItems[i].Tag < tagItems[j].Tag
		})
		c.JSON(
			http.StatusOK,
			models.PackResponseData(
				http.StatusOK,
				"get tags success",
				&models.TagsList{Tags: tagItems},
			),
		)
	}
}
