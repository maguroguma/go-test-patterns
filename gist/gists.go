package main

import (
	"io"
	"strings"
	"testing"
)

type Gist struct {
	Rawurl string `json:"html_url"`
}

// DoerはGistsのAPIにリクエストするインタフェース
type Doer interface {
	doGistsRequest(user string) (io.Reader, error)
}

// ClientはGistのList APIを扱うためのクライアント実装
type Client struct {
	Gister Doer
}

type Gister struct{}

func (g *Gister) doGistsRequest(user string) (io.Reader, error) {
	// 実装は省略
}

func (c *Client) ListGists(user string) ([]string, error) {
	r, err := c.Gister.doGistsRequest(user)
	if err != nil {
		return nil, err
	}
	// 省略
}

// Doerのインタフェースを満たすダミーのstruct
type dummyDoer struct{}

// doGistsRequestのダミー実装
// HTTP Requestを送らないでダミーのデータを返す
func (d *dummyDoer) doGistsRequest(user string) (io.Reader, error) {
	return strings.NewReader(`
[
	{"html_url": "https://gist.github.com/example1"},
	{"html_url": "https://gist.github.com/example2"}
]
	`), nil
}

func TestGetGists2(t *testing.T) {
	// dummyDoerはDoerの実装なので、Clientにわたす事ができる
	c := &Client{&dummyDoer{}}
	urls, err := c.ListGists("test")
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}
	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d", expected, len(urls))
	}
}
