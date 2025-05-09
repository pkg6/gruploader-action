package clouds

import (
	"fmt"
	"github.com/pkg6/gruploader-action/gruploader/clouds/github"
	"github.com/pkg6/gruploader-action/gruploader/flags"
	"github.com/pkg6/gruploader-action/gruploader/utils"
	"os"
)

const (
	ENV_GITHUB_TOKEN = "GITHUB_TOKEN"
)

const (
	GITHUB = "github"
)

type ICloud interface {
	Upload(assetPath string) error
}

func New(flags *flags.Flags) (ICloud, error) {
	ownerName, repoName, err := utils.ParseRepo(flags.Repo)
	if err != nil {
		return nil, err
	}
	switch flags.Cloud {
	case GITHUB:
		token := os.Getenv(ENV_GITHUB_TOKEN)
		if token == "" {
			return nil, fmt.Errorf("environment variable %s is not set", ENV_GITHUB_TOKEN)
		}
		return &github.Cloud{
			Token:     token,
			Tag:       flags.Tag,
			Owner:     ownerName,
			Repo:      repoName,
			Overwrite: flags.Overwrite,
		}, nil
	default:
		return nil, fmt.Errorf("cloud not supported")
	}
}
