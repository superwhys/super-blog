package postmanager

import (
	"context"

	"github.com/superwhys/super-blog/models"
)

type PostManager interface {
	GetTotalPostCount(ctx context.Context) (int64, error)
	GetPostList(ctx context.Context, pagination models.Pagination) (*models.BlogListItems, error)
	GetPost(ctx context.Context, postName string) (*models.BlogItem, error)
	DeletePost(ctx context.Context, postName string) error
	UpdatePost(ctx context.Context, item *models.BlogItem) error
}
