package client

import (
	"context"
	"github.com/google/go-github/v31/github"
)

type RepoService interface {
	// GitHub API docs: https://developer.github.com/v3/repos/commits/#list-commits-on-a-repository
	ListCommits(ctx context.Context,
		owner, repo string,
		opts *github.CommitsListOptions) ([]*github.RepositoryCommit, *github.Response, error)

	// GitHub API docs: https://developer.github.com/v3/repos/releases/#get-the-latest-release
	GetLatestRelease(ctx context.Context,
		owner, repo string) (*github.RepositoryRelease, *github.Response, error)

	// GitHub API docs: https://developer.github.com/v3/repos/#list-tags
	ListTags(ctx context.Context,
		owner string, repo string,
		opts *github.ListOptions) ([]*github.RepositoryTag, *github.Response, error)

	// GitHub API docs: https://developer.github.com/v3/repos/releases/#create-a-release
	CreateRelease(ctx context.Context,
		owner, repo string,
		release *github.RepositoryRelease) (*github.RepositoryRelease, *github.Response, error)
}

type Client struct {
	Repositories RepoService
}
