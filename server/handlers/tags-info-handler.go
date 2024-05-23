package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/superwhys/super-blog/models"
	"github.com/superwhys/super-blog/pkg/postmanager"
	"github.com/superwhys/venkit/lg"
)

func getTagsInfo(ctx context.Context, postManager postmanager.PostManager) (map[string][]*models.TagItem, error) {
	postList, err := postManager.GetPostList(ctx, models.Pagination{Page: -1, Size: -1})
	if err != nil {
		return nil, errors.Wrap(err, "get post list")
	}

	tagsInfo := make(map[string][]*models.TagItem)
	for _, post := range postList.Items {
		for _, tag := range post.MetaData.Tags {
			if _, exists := tagsInfo[tag.Tag]; !exists {
				tagsInfo[tag.Tag] = make([]*models.TagItem, 0)
			}

			tagsInfo[tag.Tag] = append(tagsInfo[tag.Tag], &models.TagItem{
				Tag:  tag.Tag,
				Info: post,
			})
		}
	}
	return tagsInfo, nil
}

func GetTagsInfoHandler(hctx *HandlerContext) gin.HandlerFunc {
	ctx := lg.With(hctx.ctx, "[GetTagsInfoHandler]")
	return func(c *gin.Context) {
		tagsInfo, err := getTagsInfo(ctx, hctx.postManager)
		if err != nil {
			lg.Errorc(ctx, "get tags info error: %v", err)
			c.JSON(http.StatusInternalServerError, models.PackResponseData(http.StatusInternalServerError, "get tags info failed", nil))
			return
		}

		c.JSON(http.StatusOK, models.PackResponseData(http.StatusOK, "get tags info success", &models.TagsGroupList{Tags: tagsInfo}))
	}
}

func GetSpecifyTagInfoHandler(hctx *HandlerContext) gin.HandlerFunc {
	ctx := lg.With(hctx.ctx, "[GetSpecifyTagInfoHandler]")
	return func(c *gin.Context) {
		tag := c.Param("tag")
		if tag == "" {
			lg.Errorc(ctx, "no tag param provided")
			c.JSON(http.StatusBadRequest, models.PackResponseData(http.StatusBadRequest, "no tag provided", nil))
			return
		}

		tagsInfo, err := getTagsInfo(ctx, hctx.postManager)
		if err != nil {
			lg.Errorc(ctx, "get tags info error: %v", err)
			c.JSON(http.StatusInternalServerError, models.PackResponseData(http.StatusInternalServerError, "get tags info failed", nil))
			return
		}

		info, exists := tagsInfo[tag]
		if !exists {
			lg.Errorc(ctx, "%v not exists", tag)
			c.JSON(http.StatusBadRequest, models.PackResponseData(http.StatusBadRequest, "tag not exists", nil))
			return
		}
		c.JSON(
			http.StatusOK,
			models.PackResponseData(
				http.StatusOK,
				"get tag info success",
				&models.TagsGroupList{Tags: map[string][]*models.TagItem{tag: info}},
			),
		)
	}
}
