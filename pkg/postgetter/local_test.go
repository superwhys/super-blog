package postgetter

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/superwhys/superBlog/models"
	"github.com/yazl-tech/yazl/utils/lg"
)

var (
	localGetter *LocalGetter
)

func TestMain(m *testing.M) {
	baseDir := "../../blog-posts"
	localGetter = NewLocalPostGetter(baseDir)
	m.Run()
}

func TestLocalGetter_GetPostList(t *testing.T) {
	wantErr := false
	t.Run("getPostList", func(t *testing.T) {
		got, err := localGetter.GetPostList(context.Background())
		if (err != nil) != wantErr {
			t.Errorf("LocalGetter.GetPostList() error = %v, wantErr %v", err, wantErr)
			return
		}
		fmt.Println(lg.Jsonify(got))
	})
}

func TestLocalGetter_GetSpecifyPost(t *testing.T) {
	type fields struct {
		postBaseDir string
		posts       map[string]*models.BlogItem
	}
	type args struct {
		ctx      context.Context
		postName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.BlogItem
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LocalGetter{
				postBaseDir: tt.fields.postBaseDir,
				posts:       tt.fields.posts,
			}
			got, err := l.GetSpecifyPost(tt.args.ctx, tt.args.postName)
			if (err != nil) != tt.wantErr {
				t.Errorf("LocalGetter.GetSpecifyPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocalGetter.GetSpecifyPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
