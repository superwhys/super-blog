package postgetter

import (
	"context"
	"fmt"
	"regexp"

	"github.com/pkg/errors"
	"github.com/superwhys/superBlog/models"
	"gopkg.in/yaml.v3"
)

type Getter interface {
	GetPostList(ctx context.Context) (*models.BlogListItems, error)
	GetSpecifyPost(ctx context.Context, postName string) (*models.BlogItem, error)
}

var (
	contentPattern = regexp.MustCompile(`(?s)^---(.*)---\n`)
)

func ParsePostContent(ctx context.Context, fileName, content string) (*models.BlogItem, error) {
	findResp := contentPattern.FindStringSubmatch(content)
	if len(findResp) != 2 {
		return nil, fmt.Errorf("Incorrect format of post: %v", fileName)
	}

	metaData := findResp[1]
	meta := &models.MetaData{}
	if err := yaml.Unmarshal([]byte(metaData), meta); err != nil {
		return nil, errors.Wrapf(err, "decode post: %v meta data", fileName)
	}
	content = contentPattern.ReplaceAllString(content, "")

	item := &models.BlogItem{
		MetaData:    meta,
		FileName:    fileName,
		Title:       meta.Title,
		SubTitle:    meta.SubTitle,
		PostContent: content,
	}
	return item, nil
}
