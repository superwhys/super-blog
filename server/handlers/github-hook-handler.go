package handlers

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/superwhys/goutils/lg"
	"github.com/superwhys/superBlog/models"
	"github.com/superwhys/superBlog/pkg/postmanager"
	"golang.org/x/sync/errgroup"
)

func GithubHookHandler(ctx context.Context, localGetter *postmanager.LocalGetter, githubGetter *postmanager.GithubGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqEvent := c.GetHeader("X-GitHub-Event")
		if reqEvent != "push" {
			c.JSON(http.StatusForbidden, models.PackResponseData(http.StatusForbidden, "event not push", nil))
			return
		}

		var event models.GithubPushEvent
		if err := c.ShouldBind(&event); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				models.PackResponseData(http.StatusInternalServerError, err.Error(), nil),
			)
			return
		}

		handleNewOrModifyFile(lg.With(ctx, "[AddedFile]"), event.HeadCommit.Added, localGetter, githubGetter)
		handleNewOrModifyFile(lg.With(ctx, "[ModifiedFile]"), event.HeadCommit.Modified, localGetter, githubGetter)
		handleDelededFile(lg.With(ctx, "[DeleteFile]"), event.HeadCommit.Removed, localGetter)
	}
}

func handleNewOrModifyFile(ctx context.Context, files []string, localGetter *postmanager.LocalGetter, githubGetter *postmanager.GithubGetter) error {
	if len(files) == 0 {
		return nil
	}

	grp, ctx := errgroup.WithContext(ctx)
	grp.SetLimit(4)
	for _, file := range files {
		file := file
		lg.Debugc(ctx, "get file: %v", file)
		fileSubMatch := postmanager.FindPostFileNameSubMatch(filepath.Base(file))
		if len(fileSubMatch) != 3 {
			lg.Warnc(ctx, "file: %v not a standard post file", filepath.Base(file))
			continue
		}

		grp.Go(func() error {
			blogItem, err := githubGetter.GetSpecifyPost(ctx, file)
			if err != nil {
				lg.Errorc(ctx, "github get post content error: %v", err)
				return errors.Wrap(err, "github get post content")
			}

			if err := localGetter.AddOrUpdatePost(ctx, blogItem); err != nil {
				lg.Errorc(ctx, "process file: %v error: %v", file, err)
				return errors.Wrap(err, "AddOrUpdatePost")
			}

			lg.Debugc(ctx, "process file: %v", file)
			return nil
		})
	}

	return grp.Wait()
}

func handleDelededFile(ctx context.Context, files []string, localGetter *postmanager.LocalGetter) error {
	for _, file := range files {
		fileName := filepath.Base(file)
		localGetter.DeletePost(ctx, fileName)
	}
	return nil
}
