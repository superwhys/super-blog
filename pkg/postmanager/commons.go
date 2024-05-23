package postmanager

import (
	"context"
	"encoding/base64"
	"fmt"
	"path/filepath"
	"regexp"
	"time"

	"github.com/pkg/errors"
	"github.com/superwhys/super-blog/models"
	"gopkg.in/yaml.v3"
)

var (
	postFilePattern = regexp.MustCompile(`(\d{4}-\d{2}-\d{2})-(.*).md`)
	contentPattern  = regexp.MustCompile(`(?s)^---(.*)---\n`)
)

func FindPostFileNameSubMatch(fileName string) []string {
	fileSubMatch := postFilePattern.FindStringSubmatch(filepath.Base(fileName))
	if len(fileSubMatch) != 3 {
		return nil
	}
	return fileSubMatch
}

func processTags(rowTags []string) (tags []models.Tag) {
	for _, tag := range rowTags {
		tags = append(tags, models.Tag{Tag: tag})
	}
	return
}

func ParsePostContent(ctx context.Context, fileName, content string) (*models.BlogItem, error) {
	fileSubMatch := FindPostFileNameSubMatch(fileName)
	if len(fileSubMatch) != 3 {
		return nil, fmt.Errorf("file: %v not a standard post file", fileName)
	}

	findResp := contentPattern.FindStringSubmatch(content)
	if len(findResp) != 2 {
		return nil, fmt.Errorf("Incorrect format of post: %v", fileName)
	}

	metaData := findResp[1]
	meta := &models.MetaData{}
	if err := yaml.Unmarshal([]byte(metaData), meta); err != nil {
		return nil, errors.Wrapf(err, "decode post: %v meta data", fileName)
	}
	meta.Tags = processTags(meta.RowTags)

	content = contentPattern.ReplaceAllString(content, "")
	encodeContent := base64.StdEncoding.EncodeToString([]byte(content))
	encodeMeta := base64.StdEncoding.EncodeToString([]byte(metaData))

	t, _ := time.Parse(time.DateOnly, fileSubMatch[1])

	item := &models.BlogItem{
		MetaData:    meta,
		FileName:    fileName,
		MetaRow:     encodeMeta,
		PostContent: encodeContent,
		PostedTime:  t.Format(fmt.Sprintf("Posted By %v on Jan 02, 2006", meta.Author)),
		ToEndPoint:  t.Format(fmt.Sprintf("/2006/01/02/%v", fileSubMatch[2])),
	}

	return item, nil
}
