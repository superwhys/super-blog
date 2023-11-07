package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/superwhys/superBlog/pkg/postgetter"
	"github.com/yazl-tech/yazl/utils/lg"
)

// func getGitBlob(ctx context.Context, client *github.Client, sha string) (string, error) {
// 	blob, _, err := client.Git.GetBlob(ctx, owner, repo, sha)
// 	if err != nil {
// 		return "", errors.Wrap(err, "get blob")
// 	}
//
// 	rb, err := base64.StdEncoding.DecodeString(blob.GetContent())
// 	if err != nil {
// 		return "", errors.Wrap(err, "decode blob content")
// 	}
// 	return string(rb), nil
// }

func BlogListHandler(ctx context.Context, getter postgetter.Getter) gin.HandlerFunc {
	ctx = lg.With(ctx, "[BlogListHandler]")
	return func(c *gin.Context) {
		// tree, _, err := client.Git.GetTree(ctx, owner, repo, branch, false)
		// if err != nil {
		// 	lg.Errorc(ctx, "get repositories tree error: %v", err)
		// 	c.JSON(
		// 		http.StatusInternalServerError,
		// 		models.PackResponseData(http.StatusInternalServerError, errors.Wrap(err, "get tree").Error(), nil),
		// 	)
		// 	return
		// }

		// var contents []string
		// lock := sync.Mutex{}
		// grp, ctx := errgroup.WithContext(ctx)
		// grp.SetLimit(len(tree.Entries))
		// for _, entry := range tree.Entries {
		// 	entry := entry
		// 	grp.Go(func() error {
		// 		content, err := getGitBlob(ctx, client, entry.GetSHA())
		// 		if err != nil {
		// 			return errors.Wrapf(err, "get %v blob", entry.GetPath())
		// 		}
		// 		lock.Lock()
		// 		defer lock.Unlock()

		// 		contents = append(contents, content)
		// 		return nil
		// 	})
		// }
	}
}
