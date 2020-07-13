package testclient

import (
	"context"
	"encoding/json"
	"github.com/google/go-github/v31/github"
	"github.com/ianusmagnus/gh-release/ghutil/client"
	"io/ioutil"
)

type RepoService struct {
}

func (g *RepoService) ListCommits(_ context.Context,
	_, _ string, _ *github.CommitsListOptions) ([]*github.RepositoryCommit, *github.Response, error) {

	bodyBytes, ioerr := ioutil.ReadFile("testdata/list_commits_response.json")
	if ioerr != nil {
		return nil, nil, ioerr
	}

	var commits []*github.RepositoryCommit
	err := json.Unmarshal(bodyBytes, &commits)

	return commits, nil, err
}

func (g *RepoService) GetLatestRelease(_ context.Context,
	_, _ string) (*github.RepositoryRelease, *github.Response, error) {

	bodyBytes, ioerr := ioutil.ReadFile("testdata/get_latest_release_response.json")
	if ioerr != nil {
		return nil, nil, ioerr
	}
	release := github.RepositoryRelease{}
	err := json.Unmarshal(bodyBytes, &release)

	return &release, nil, err
}

func (g *RepoService) ListTags(_ context.Context, _ string, _ string,
	_ *github.ListOptions) ([]*github.RepositoryTag, *github.Response, error) {

	bodyBytes, ioerr := ioutil.ReadFile("testdata/list_tags_response.json")
	if ioerr != nil {
		return nil, nil, ioerr
	}
	var tags []*github.RepositoryTag
	err := json.Unmarshal(bodyBytes, &tags)

	return tags, nil, err
}

func (g *RepoService) CreateRelease(_ context.Context, _, _ string,
	_ *github.RepositoryRelease) (*github.RepositoryRelease, *github.Response, error) {

	bodyBytes, ioerr := ioutil.ReadFile("testdata/create_release_response.json")
	if ioerr != nil {
		return nil, nil, ioerr
	}

	var result *github.RepositoryRelease
	err := json.Unmarshal(bodyBytes, &result)

	return result, nil, err
}

func newGitHubRepoService() *RepoService {
	return &RepoService{}
}

func New(_ context.Context, _ string) *client.Client {
	return &client.Client{
		Repositories: newGitHubRepoService(),
	}
}
