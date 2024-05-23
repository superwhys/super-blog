package postmanager

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/pkg/errors"
	"github.com/superwhys/super-blog/models"
	"github.com/superwhys/venkit/lg"
)

type LocalGetter struct {
	sync.RWMutex
	postBaseDir string
	postsMap    map[string]*models.BlogItem
	posts       []*models.BlogItem
	postsIdxMap map[string]int
}

func NewLocalPostGetter(postBaseDir string) *LocalGetter {
	entries, err := os.ReadDir(postBaseDir)
	lg.PanicError(err)

	postsMap := make(map[string]*models.BlogItem)
	postsIdxMap := make(map[string]int)
	posts := make([]*models.BlogItem, 0, len(entries))

	l := &LocalGetter{
		postBaseDir: postBaseDir,
		postsMap:    postsMap,
		posts:       posts,
		postsIdxMap: postsIdxMap,
	}

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

		l.insertPost(blogItem)
	}

	return l
}

func (l *LocalGetter) insertPost(item *models.BlogItem) {
	insertIdx := sort.Search(len(l.posts), func(i int) bool {
		if l.posts[i].PostedTime == item.PostedTime {
			return l.posts[i].MetaData.Title > item.MetaData.Title
		}
		return l.posts[i].ToEndPoint < item.ToEndPoint
	})
	l.posts = append(l.posts, nil)
	copy(l.posts[insertIdx+1:], l.posts[insertIdx:])
	l.posts[insertIdx] = item
	l.postsIdxMap[item.FileName] = insertIdx
	l.postsMap[item.FileName] = item
}

func (l *LocalGetter) GetTotalPostCount(ctx context.Context) (int64, error) {
	return int64(len(l.posts)), nil
}

func (l *LocalGetter) GetPostList(ctx context.Context, pagination models.Pagination) (*models.BlogListItems, error) {
	l.RLock()
	defer l.RUnlock()

	startIdx, endIdx := func() (int, int) {
		if pagination.Page == -1 && pagination.Size == -1 {
			return 0, len(l.posts)
		}
		startIdx := (pagination.Page - 1) * pagination.Size
		endIdx := startIdx + pagination.Size
		if endIdx > len(l.posts) {
			endIdx = len(l.posts)
		}
		return startIdx, endIdx
	}()

	posts := l.posts[startIdx:endIdx]

	return &models.BlogListItems{
		Items: posts,
		Total: len(l.posts),
	}, nil
}

func (l *LocalGetter) GetPost(ctx context.Context, postName string) (*models.BlogItem, error) {
	l.RLock()
	defer l.RUnlock()

	post, exists := l.postsMap[postName]
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

func decodeContent(ctx context.Context, content string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		lg.Errorc(ctx, "decode content error: %v", err)
		return "", err
	}
	return string(b), nil
}

func writePostToFile(ctx context.Context, basePostDir string, blogItem *models.BlogItem) error {
	file, err := os.Create(filepath.Join(basePostDir, blogItem.FileName))
	if err != nil {
		return errors.Wrap(err, "create file")
	}

	var (
		metaContent string
		postContent string
	)
	metaContent, err = decodeContent(ctx, blogItem.MetaRow)
	postContent, err = decodeContent(ctx, blogItem.PostContent)
	if err != nil {
		return errors.Wrap(err, "decode content")
	}

	content := fmt.Sprintf("%v\n%v", metaContent, postContent)

	_, err = file.Write([]byte(content))
	if err != nil {
		return errors.Wrap(err, "write content to string")
	}
	return nil
}

func (l *LocalGetter) UpdatePost(ctx context.Context, item *models.BlogItem) error {
	l.Lock()
	defer l.Unlock()

	if err := writePostToFile(ctx, l.postBaseDir, item); err != nil {
		lg.Errorc(ctx, "write new post to file error: %v", err)
		return errors.Wrap(err, "writeNewPostToFile")
	}

	if _, exists := l.postsMap[item.FileName]; exists {
		l.posts[l.postsIdxMap[item.FileName]] = item
		l.postsMap[item.FileName] = item
	}

	return nil
}

func (l *LocalGetter) DeletePost(ctx context.Context, fileName string) error {
	l.Lock()
	defer l.Unlock()

	postIdx, exists := l.postsIdxMap[fileName]
	if !exists {
		return fmt.Errorf("post: %v not exists!", fileName)
	}

	// delete from l.posts
	l.posts = append(l.posts[:postIdx], l.posts[postIdx+1:]...)
	// delete from l.postsIdxMap
	delete(l.postsIdxMap, fileName)
	// delete from l.postsMap
	delete(l.postsMap, fileName)

	return nil
}
