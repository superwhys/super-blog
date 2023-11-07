package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/superwhys/superBlog/models"
	"github.com/superwhys/superBlog/pkg/postmanager"
	"github.com/yazl-tech/yazl/utils/lg"
)

type postHandlerParams struct {
	Year  string `form:"year" binding:"required"`
	Month string `form:"month" binding:"required"`
	Day   string `form:"day" binding:"required"`
	Name  string `form:"name" binding:"required"`
}

func PostHandler(ctx context.Context, localGetter *postmanager.LocalGetter) gin.HandlerFunc {
	ctx = lg.With(ctx, "[PostHandler]")
	return func(c *gin.Context) {
		params := postHandlerParams{
			Year:  c.Param("year"),
			Month: c.Param("month"),
			Day:   c.Param("day"),
			Name:  c.Param("name"),
		}

		fileName := fmt.Sprintf("%v-%v-%v-%v.md", params.Year, params.Month, params.Day, params.Name)
		blogItem, err := localGetter.GetSpecifyPost(ctx, fileName)
		if err != nil {
			lg.Errorc(ctx, "get post: %v error: %v", fileName, err)
			c.JSON(http.StatusBadRequest, models.PackResponseData(http.StatusBadRequest, fmt.Sprintf("get post: %v failed", fileName), nil))
			return
		}

		c.JSON(http.StatusOK, models.PackResponseData(http.StatusOK, fmt.Sprintf("get post: %v success", fileName), blogItem))
	}
}
