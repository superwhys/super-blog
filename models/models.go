package models

import (
	"encoding/base64"
	"time"

	"github.com/superwhys/venkit/lg"
	"gorm.io/gorm"
)

type Pagination struct {
	Page int `json:"page" form:"page,default=1"`
	Size int `json:"size" form:"size,default=7"`
}

func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.Size
}

type ApiResponseData struct {
	Message    string `json:"message"`
	Data       any    `json:"data"`
	StatusCode int    `json:"statusCode"`
	RequestTs  int64  `json:"requestTs"`
}

func PackResponseData(statusCode int, msg string, data any) *ApiResponseData {
	return &ApiResponseData{
		Message:    msg,
		StatusCode: statusCode,
		Data:       data,
		RequestTs:  time.Now().Unix(),
	}
}

type Tag struct {
	Tag string `gorm:"primarykey"`
}

type MetaData struct {
	Layout   string   `yaml:"layout" json:"layout"`
	Title    string   `yaml:"title" json:"title"`
	SubTitle string   `yaml:"subTitle" json:"subTitle"`
	Date     string   `yarm:"date" json:"date"`
	Author   string   `yaml:"author" json:"author"`
	RowTags  []string `yaml:"tags" json:"tags" gorm:"-:all"`
	Tags     []Tag    `yaml:"-" json:"-" gorm:"many2many:blog_tags"`
}

func (m *MetaData) GetTagsList() (tags []string) {
	for _, tag := range m.Tags {
		tags = append(tags, tag.Tag)
	}
	return
}

func (m *MetaData) BackFillRowTags() {
	for _, tag := range m.Tags {
		m.RowTags = append(m.RowTags, tag.Tag)
	}
}

type BlogItem struct {
	gorm.Model
	MetaData    *MetaData `json:"metaData" gorm:"embedded"`
	MetaRow     string    `json:"-" gorm:"-:all"`
	FileName    string    `json:"fileName" gorm:"unique"`
	PostContent string    `json:"postContent"`
	PostedTime  string    `json:"postedTime"`
	ToEndPoint  string    `json:"toEndPoint"`
}

func (b *BlogItem) DecodeContent() {
	if b.PostContent == "" {
		return
	}
	bc, err := base64.StdEncoding.DecodeString(b.PostContent)
	if err != nil {
		lg.Errorf("decode post: %v content error: %v", b.FileName, err)
		return
	}
	b.PostContent = string(bc)
}

type BlogListItems struct {
	Items []*BlogItem `json:"items"`
	Total int         `json:"total"`
}

func (m *BlogListItems) DecodeBlogItemContent() {
	for _, item := range m.Items {
		item.DecodeContent()
	}
}

func (m *BlogListItems) BackFillRowTags() {
	for _, item := range m.Items {
		item.MetaData.BackFillRowTags()
	}
}

type TagItem struct {
	Info       *BlogItem `json:"info,omitempty"`
	Tag        string    `json:"tag,omitempty"`
	ToEndPoint string    `json:"toEndPoint,omitempty"`
}

type TagsGroupList struct {
	Tags map[string][]*TagItem `json:"tags"`
}

type TagsList struct {
	Tags []*TagItem `json:"tags"`
}

type GithubCommit struct {
	Added    []string `json:"added" form:"added"`
	Modified []string `json:"modified" form:"modified"`
	Removed  []string `json:"removed" form:"removed"`
}

type GithubPushEvent struct {
	Ref        string          `json:"ref" form:"ref"`
	Commits    []*GithubCommit `json:"commits" form:"commits"`
	HeadCommit *GithubCommit   `json:"head_commit" form:"head_commit"`
}
