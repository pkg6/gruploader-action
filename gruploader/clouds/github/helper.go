package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v68/github"
	"github.com/pkg6/gruploader-action/gruploader/logger"
	"sort"
)

func createRelease(ctx context.Context, client *github.Client, owner, repo, createTagName string) (createRelease *github.RepositoryRelease, err error) {
	releases, _, err := client.Repositories.ListReleases(ctx, owner, repo, &github.ListOptions{PerPage: 100})
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		if release.GetTagName() != createTagName {
			createRelease, _, err = client.Repositories.CreateRelease(ctx, owner, repo, &github.RepositoryRelease{
				TagName: &createTagName,
			})
			logger.WarningF("Create new Release, id %d, tag '%s', url '%s' commit `%s`",
				createRelease.GetID(),
				createTagName,
				createRelease.GetHTMLURL(),
				createRelease.GetTargetCommitish(),
			)
			return
		}
	}
	return nil, fmt.Errorf("no release for %s/%s found", owner, repo)
}

func deleteReleaseAssets(ctx context.Context, client *github.Client, owner, repo string, release *github.RepositoryRelease) error {
	var (
		assets []*github.ReleaseAsset
		err    error
	)
	assets, _, err = client.Repositories.ListReleaseAssets(ctx, owner, repo, release.GetID(), &github.ListOptions{PerPage: 100})
	if err != nil {
		return err
	}
	for _, asset := range assets {
		if _, err = client.Repositories.DeleteReleaseAsset(ctx, owner, repo, asset.GetID()); err != nil {
			return err
		}
		logger.WarningF("Deleted asset, id %d, name '%s', url '%s'\n", asset.GetID(), asset.GetName(), asset.GetBrowserDownloadURL())
	}
	if len(assets) > 0 {
		logger.WarningF("Total deleted %d assets\n", len(assets))
	}
	return nil
}

func deleteReleaseAsset(ctx context.Context, client *github.Client, owner, repo string, release *github.RepositoryRelease, assetName string) error {
	var (
		assets []*github.ReleaseAsset
		err    error
	)
	assets, _, err = client.Repositories.ListReleaseAssets(ctx, owner, repo, release.GetID(), &github.ListOptions{PerPage: 100})
	if err != nil {
		return err
	}
	for _, asset := range assets {
		if asset.GetName() == assetName {
			// found exist one, delete it
			if _, err = client.Repositories.DeleteReleaseAsset(ctx, owner, repo, asset.GetID()); err != nil {
				return err
			}
			logger.WarningF("Deleted old asset, id %d, name '%s', url '%s'\n",
				asset.GetID(),
				asset.GetName(),
				asset.GetBrowserDownloadURL(),
			)
			break
		}
	}
	return nil
}

func getReleaseByName(ctx context.Context, client *github.Client, repoOwner, repoName, releaseName string) (*github.RepositoryRelease, error) {
	releases, _, err := client.Repositories.ListReleases(ctx, repoOwner, repoName, nil)
	if err != nil {
		return nil, err
	}
	sort.SliceStable(releases, func(i, j int) bool {
		return releases[i].CreatedAt.After(releases[j].CreatedAt.Time)
	})
	for _, release := range releases { // assume they are some kind of sorted, first pick (newest)
		if release.GetName() == releaseName {
			return release, nil
		}
	}
	return nil, fmt.Errorf("no release found with name: %q", releaseName)
}
