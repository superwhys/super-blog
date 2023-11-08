package postmanager

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/pkg/errors"
	"github.com/superwhys/goutils/lg"
	"github.com/superwhys/superBlog/models"
)

type LocalGetter struct {
	sync.RWMutex
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

		content, err := readFile(filepath.Join(postBaseDir, entry.Name()))
		lg.PanicError(err)

		blogItem, err := ParsePostContent(context.Background(), entry.Name(), content)
		if err != nil {
			lg.Errorf("parse %v content error: %v", entry.Name(), err)
			continue
		}

		posts[entry.Name()] = blogItem
	}

	return &LocalGetter{
		postBaseDir: postBaseDir,
		posts:       posts,
	}
}

func (l *LocalGetter) GetPostList(ctx context.Context) (*models.BlogListItems, error) {
	l.RLock()
	defer l.RUnlock()

	var posts []*models.BlogItem
	for _, value := range l.posts {
		posts = append(posts, value)
	}
	return &models.BlogListItems{
		Items: posts,
	}, nil
}

func (l *LocalGetter) GetSpecifyPost(ctx context.Context, postName string) (*models.BlogItem, error) {
	l.RLock()
	defer l.RUnlock()

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

func writePostToFile(ctx context.Context, basePostDir string, blogItem *models.BlogItem) error {
	file, err := os.Create(filepath.Join(basePostDir, blogItem.FileName))
	if err != nil {
		return errors.Wrap(err, "create file")
	}

	content, err := base64.StdEncoding.DecodeString(blogItem.RawContent)
	if err != nil {
		return errors.Wrap(err, "decode raw content")
	}

	_, err = file.Write(content)
	if err != nil {
		return errors.Wrap(err, "write content to string")
	}
	return nil
}

func (l *LocalGetter) AddOrUpdatePost(ctx context.Context, item *models.BlogItem) error {
	l.Lock()
	defer l.Unlock()

	if err := writePostToFile(ctx, l.postBaseDir, item); err != nil {
		lg.Errorc(ctx, "write new post to file error: %v", err)
		return errors.Wrap(err, "writeNewPostToFile")
	}

	l.posts[item.FileName] = item
	return nil
}

func (l *LocalGetter) DeletePost(ctx context.Context, fileName string) error {
	l.Lock()
	defer l.Unlock()

	delete(l.posts, fileName)
	return nil
}
