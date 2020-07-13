package ghutil

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/go-github/v31/github"
	"github.com/ianusmagnus/gh-release/ghutil/client"
	log "github.com/sirupsen/logrus"
)

type ReleaseCreator struct {
	username string
	repo     string
	client   *client.Client
	ctx      context.Context
}

func NewReleaseCreator(client *client.Client, ctx context.Context, username string, repo string) *ReleaseCreator {

	return &ReleaseCreator{username: username, repo: repo, client: client, ctx: ctx}
}

func (s *ReleaseCreator) listCommits(sha string) ([]*github.RepositoryCommit, error) {

	opts := &github.CommitsListOptions{}
	commits, _, err := s.client.Repositories.ListCommits(s.ctx, s.username, s.repo, opts)
	if err != nil {
		return nil, err
	}

	var end int
	for i, commit := range commits {

		log.Debugf("commit: %v", *commit.SHA)

		if commit.GetSHA() == sha {
			end = i
			break
		}

		log.Debugf("add commit: %v", *commit.SHA)
	}

	return commits[0:end], nil
}

func (s *ReleaseCreator) getLatestReleaseTag() (string, error) {

	log.Info("Get latest release tag.")

	var tag string

	release, _, err := s.client.Repositories.GetLatestRelease(s.ctx, s.username, s.repo)
	if err != nil {
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

	log.Infof("Created releanse: %v", *release.Name)

	return nil
}

func (s *ReleaseCreator) CreateNewRelease(name string) error {

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

	if len(commits) == 0 {
		return fmt.Errorf("no commits for new release found")
	}

	notes := s.createReleaseNotesFromCommits(commits)

	err = s.createRelease(name, name, notes)
	if err != nil {
		return err
	}

	return nil
}
