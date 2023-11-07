package postgetter

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/pkg/errors"
	"github.com/superwhys/superBlog/models"
	"github.com/yazl-tech/yazl/utils/lg"
)

var (
	postFilePattern = regexp.MustCompile(`(\d{4}-\d{2}-\d{2})-(.*).md`)
)

type LocalGetter struct {
	postBaseDir string
	posts       map[string]*models.BlogItem
}

func NewLocalPostGetter(postBaseDir string) *LocalGetter {
	entries, err := os.ReadDir(postBaseDir)
	lg.PanicError(err)

	posts := make(map[string]*models.BlogItem)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fileSubMatch := postFilePattern.FindStringSubmatch(entry.Name())
		if len(fileSubMatch) != 3 {
			lg.Warnf("file: %v not a standard post file", entry.Name())
			continue
		}

		content, err := readFile(filepath.Join(postBaseDir, entry.Name()))
		lg.PanicError(err)

		blogItem, err := ParsePostContent(context.Background(), entry.Name(), content)
		if err != nil {
			lg.Errorf("parse %v content error: %v", entry.Name(), err)
			continue
		}

		t, _ := time.Parse(time.DateOnly, fileSubMatch[1])
		blogItem.PostedTime = t.Format(fmt.Sprintf("Posted By %v on Jan 02, 2006", blogItem.MetaData.Author))
		blogItem.Title = fileSubMatch[2]
		blogItem.ToEndPoint = t.Format(fmt.Sprintf("/2006/01/02/%v", fileSubMatch[2]))

		posts[entry.Name()] = blogItem
	}

	return &LocalGetter{
		postBaseDir: postBaseDir,
		posts:       posts,
	}
}

func (l *LocalGetter) GetPostList(ctx context.Context) (*models.BlogListItems, error) {
	var posts []*models.BlogItem
	for _, value := range l.posts {
		posts = append(posts, value)
	}
	return &models.BlogListItems{
		Items: posts,
	}, nil
}

func (l *LocalGetter) GetSpecifyPost(ctx context.Context, postName string) (*models.BlogItem, error) {
	post, exists := l.posts[postName]
	if !exists {
		return nil, fmt.Errorf("post: %v not exists!", postName)
	}
	return post, nil
}

func readFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", errors.Wrap(err, "read file")
	}
	return string(b), nil
}
