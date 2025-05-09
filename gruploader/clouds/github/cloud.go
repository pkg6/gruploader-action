package github

import (
	"context"
	"github.com/google/go-github/v68/github"
	"github.com/pkg6/gruploader-action/gruploader/logger"
	"golang.org/x/oauth2"
	"os"
	"path/filepath"
)

type Cloud struct {
	Token     string
	Tag       string
	Owner     string
	Repo      string
	Overwrite bool
}

func (c *Cloud) Upload(assetPath string) error {
	rwContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: c.Token})
	tc := oauth2.NewClient(rwContext, ts)
	client := github.NewClient(tc)
	var (
		release *github.RepositoryRelease
		err     error
	)
	release, _, err = client.Repositories.GetReleaseByTag(rwContext, c.Owner, c.Repo, c.Tag)
	if err != nil {
		return err
	}
	assetName := filepath.Base(assetPath)
	if c.Overwrite {
		// remove old one if it's exist already
		if err := deleteReleaseAsset(rwContext, client, c.Owner, c.Repo, release, assetName); err != nil {
			return err
		}
	}
	// open file for uploading
	f, err := os.Open(assetPath) // For read access.
	if err != nil {
		return err
	}
	defer f.Close()
	// upload
	releaseAsset, _, err := client.Repositories.UploadReleaseAsset(
		rwContext,
		c.Owner,
		c.Repo,
		release.GetID(),
		&github.UploadOptions{Name: assetName}, f)
	if err != nil {
		return err
	}
	logger.InfoF("Upload asset succeed, id %d, name '%s', url: '%s'",
		releaseAsset.GetID(),
		releaseAsset.GetName(),
		releaseAsset.GetBrowserDownloadURL(),
	)
	return nil
}
