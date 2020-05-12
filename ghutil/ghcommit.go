package ghutil

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/go-github/v31/github"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type ReleaseCreator struct {
	username string
	repo     string
	client   *github.Client
	ctx      context.Context
}

func NewReleaseCreator(username string, pat string, repo string) *ReleaseCreator {

	log.Info("Creating new client")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: pat},
	)
	tc := oauth2.NewClient(ctx, ts)

	Client := github.NewClient(tc)

	return &ReleaseCreator{username: username, repo: repo, client: Client, ctx: ctx}
}

func (s *ReleaseCreator) listCommits(sha string) ([]*github.RepositoryCommit, error) {

	opts := &github.CommitsListOptions{}
	commits, _, err := s.client.Repositories.ListCommits(s.ctx, s.username, s.repo, opts)
	if err != nil {
		return nil, err
	}

	var end int
	for i, commit := range commits {

		if commit.GetSHA() == sha {
			end = i
			break
		}
	}

	return commits[0:end], nil
}

func (s *ReleaseCreator) getLatestReleaseTag() (string, error) {

	log.Info("Get latest release tag.")

	var tag string

	release, response, err := s.client.Repositories.GetLatestRelease(s.ctx, s.username, s.repo)
	if err != nil {
		fmt.Printf("%+v", response)
		return tag, err
	}

	if release != nil {
		tag = release.GetTagName()
	}

	return tag, nil
}

func (s *ReleaseCreator) getSHAForTag(tagname string) (string, error) {

	var sha string

	tags, _, err := s.client.Repositories.ListTags(s.ctx, s.username, s.repo, nil)
	if err != nil {
		return sha, err
	}

	for _, tag := range tags {
		if tagname == tag.GetName() {
			sha = tag.GetCommit().GetSHA()
			break
		}
	}

	return sha, nil
}

func (s *ReleaseCreator) createReleaseNotesFromCommits(commits []*github.RepositoryCommit) string {

	var b bytes.Buffer

	for _, commit := range commits {
		messgae := commit.GetCommit().GetMessage()
		b.WriteString(messgae)
		b.WriteString("\n")
	}

	return b.String()
}

func (s *ReleaseCreator) createRelease(name string, tag string, notes string) error {

	release := &github.RepositoryRelease{Name: &name, TagName: &tag, Body: &notes}
	release, _, err := s.client.Repositories.CreateRelease(s.ctx, s.username, s.repo, release)
	if err != nil {
		return err
	}

	fmt.Println(release.HTMLURL)

	return nil
}

func (s *ReleaseCreator) CreateRelease() error {

	releaseTag, err := s.getLatestReleaseTag()
	if err != nil {
		return err
	}
	releaseSHA, err := s.getSHAForTag(releaseTag)
	if err != nil {
		return err
	}

	commits, err := s.listCommits(releaseSHA)
	if err != nil {
		return err
	}

	notes := s.createReleaseNotesFromCommits(commits)

	fmt.Println("---")
	fmt.Println(notes)
	fmt.Println("---")

	err = s.createRelease("v2", "v2", notes)
	if err != nil {
		return err
	}

	return nil
}
