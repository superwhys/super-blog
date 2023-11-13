package postmanager

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/google/go-github/v53/github"
	"github.com/pkg/errors"
	"github.com/superwhys/superBlog/models"
	"github.com/superwhys/superBlog/pkg/share"
)

type GithubGetter struct {
	client *github.Client
}

func NewGithubGetter(client *github.Client) *GithubGetter {
	return &GithubGetter{
		client: client,
	}
}

func (g *GithubGetter) GetSpecifyPost(ctx context.Context, postName string) (*models.BlogItem, error) {
	fileSubMatch := FindPostFileNameSubMatch(filepath.Base(postName))
	if len(fileSubMatch) != 3 {
		return nil, fmt.Errorf("file: %v not a standard post file", filepath.Base(postName))
	}

	fileContent, _, _, err := g.client.Repositories.GetContents(ctx, share.Owner, share.Repo, postName, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "get file: %v content", postName)
	}

	content, err := fileContent.GetContent()
	if err != nil {
		return nil, errors.Wrapf(err, "get file: %v content", postName)
	}

	item, err := ParsePostContent(ctx, filepath.Base(postName), content)
	if err != nil {
		return nil, errors.Wrapf(err, "parse file: %v content", postName)
	}

	return item, nil
}
