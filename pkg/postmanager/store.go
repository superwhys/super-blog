package postmanager

import (
	"context"

	"github.com/pkg/errors"
	"github.com/superwhys/superBlog/models"
	"github.com/superwhys/venkit/lg"
	"gorm.io/gorm"
)

type PostStore struct {
	db *gorm.DB
}

func NewPostStore(db *gorm.DB) *PostStore {
	ps := &PostStore{db: db}
	return ps
}

func (ps *PostStore) GetTotalPostCount(ctx context.Context) (int64, error) {
	var total int64
	if err := ps.db.Model(&models.BlogItem{}).Count(&total).Error; err != nil {
		return 0, errors.Wrap(err, "count post count")
	}

	return total, nil
}

func (ps *PostStore) GetPostList(ctx context.Context, pagination models.Pagination) (*models.BlogListItems, error) {
	var (
		blogs []*models.BlogItem
		err   error
	)

	if pagination.Page == -1 && pagination.Size == -1 {
		err = ps.db.Preload("Tags").Order("date desc").Find(&blogs).Error
	} else {
		offset := pagination.GetOffset()
		size := pagination.Size
		err = ps.db.Preload("Tags").Offset(offset).Limit(size).Order("date desc").Find(&blogs).Error
	}

	if err != nil {
		return nil, errors.Wrap(err, "getPostList")
	}

	total, err := ps.GetTotalPostCount(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getTotalPostCount")
	}

	return &models.BlogListItems{
		Items: blogs,
		Total: int(total),
	}, nil
}

func (ps *PostStore) GetPost(ctx context.Context, postName string) (*models.BlogItem, error) {
	blog := &models.BlogItem{}
	if err := ps.db.Preload("Tags").Where(&models.BlogItem{FileName: postName}).First(blog).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	blog.MetaData.BackFillRowTags()

	return blog, nil
}

func (ps *PostStore) DeletePost(ctx context.Context, postName string) error {
	return ps.db.Where(&models.BlogItem{FileName: postName}).Delete(&models.BlogItem{}).Error
}

func (ps *PostStore) getTag(tag string) (models.Tag, error) {
	tagItem := models.Tag{Tag: tag}
	if err := ps.db.Where("tag = ?", tag).FirstOrCreate(&tagItem).Error; err != nil {
		lg.Errorf("get tag: %v error: %v", tag, err)
		return tagItem, err
	}

	return tagItem, nil
}

func (ps *PostStore) checkItemTag(item *models.BlogItem) {
	var tags []models.Tag
	for _, t := range item.MetaData.Tags {
		if t.Tag == "" {
			continue
		}
		tag, err := ps.getTag(t.Tag)
		if err != nil {
			continue
		}
		tags = append(tags, tag)
	}

	item.MetaData.Tags = tags
}

func (ps *PostStore) UpdatePost(ctx context.Context, item *models.BlogItem) error {
	bi, err := ps.GetPost(ctx, item.FileName)
	if err != nil {
		return errors.Wrap(err, "GetPost")
	}

	if bi == nil {
		if err = ps.db.Create(item).Error; err != nil {
			return errors.Wrap(err, "Create")
		}
		return nil
	}

	lg.Debugc(ctx, "filename: %v", item.FileName)
	ps.checkItemTag(item)

	bi.MetaData = item.MetaData
	bi.PostContent = item.PostContent
	bi.PostedTime = item.PostedTime
	bi.ToEndPoint = item.ToEndPoint

	if err := ps.db.Model(bi).Association("Tags").Replace(item.MetaData.Tags); err != nil {
		return errors.Wrap(err, "Replace")
	}
	return ps.db.Save(bi).Error
}
