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
	return &PostStore{db: db}
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
		err = ps.db.Preload("Tags").Find(&blogs).Error
	} else {
		offset := pagination.GetOffset()
		size := pagination.Size
		err = ps.db.Preload("Tags").Find(&blogs).Offset(offset).Limit(size).Error
	}

	if err != nil {
		return nil, errors.Wrap(err, "getPostList")
	}

	return &models.BlogListItems{
		Items: blogs,
		Total: len(blogs),
	}, nil
}

func (ps *PostStore) GetPost(ctx context.Context, postName string) (*models.BlogItem, error) {
	blog := &models.BlogItem{}
	if err := ps.db.Preload("Tags").Where(&models.BlogItem{FileName: postName}).First(blog).Error; err != nil {
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
	for _, tag := range item.MetaData.Tags {
		if tag.Tag == "" {
			continue
		}
		tag, err := ps.getTag(tag.Tag)
		if err != nil {
			continue
		}
		tags = append(tags, tag)
	}

	item.MetaData.Tags = tags
}

func (ps *PostStore) UpdatePost(ctx context.Context, item *models.BlogItem) error {
	var count int64
	ps.db.Where(&models.BlogItem{FileName: item.FileName}).Count(&count)
	ps.checkItemTag(item)

	if count != 0 {
		return ps.db.Where(&models.BlogItem{FileName: item.FileName}).Updates(item).Error
	}

	return ps.db.Create(item).Error

	// doUpdates := []string{"title", "sub_title", "date", "author", "tags", "meta_row", "post_content", "posted_time", "to_end_point"}
	// return ps.db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "file_name"}},
	// 	DoUpdates: clause.AssignmentColumns(doUpdates),
	// }).Create(item).Error
}
