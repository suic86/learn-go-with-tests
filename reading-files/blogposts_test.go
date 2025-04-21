package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/suic86/learn-go-with-tests/reading-files"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("i always fail")
}

func TestNewBlogPostsErrorHandling(t *testing.T) {
	_, err := blogposts.NewPostsFromFS(StubFailingFS{})
	if err == nil {
		t.Error("expted an error but got nil")
	}
}

func TestNewBlogPosts(t *testing.T) {
	cases := []struct {
		filename string
		content  string
		want     blogposts.Post
	}{
		{
			filename: "hello world.md",
			content: `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`,
			want: blogposts.Post{
				Title:       "Post 1",
				Description: "Description 1",
				Tags:        []string{"tdd", "go"},
				Body: `Hello
World`,
			},
		},
		{
			filename: "hello-world2.md",
			content: `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
F
Y
I`,
			want: blogposts.Post{
				Title:       "Post 2",
				Description: "Description 2",
				Tags:        []string{"rust", "borrow-checker"},
				Body: `F
Y
I`,
			},
		},
	}
	for _, test := range cases {
		t.Run(test.filename, func(t *testing.T) {
			fs := fstest.MapFS{
				test.filename: {Data: []byte(test.content)},
			}

			posts, err := blogposts.NewPostsFromFS(fs)
			if err != nil {
				t.Fatal(err)
			}
			got := posts[0]
			want := test.want
			assertPost(t, got, want)
		})

	}
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
