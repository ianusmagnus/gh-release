package client

import (
	"context"
	"github.com/google/go-github/v31/github"
	"golang.org/x/oauth2"
	"net/http"
)

func NewClient(ctx context.Context, pat string) *Client {
	var tc *http.Client

	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: pat},
	)
	tc = oauth2.NewClient(ctx, tokenSource)

	client := github.NewClient(tc)

	return &Client{
		Repositories: client.Repositories,
	}
}
