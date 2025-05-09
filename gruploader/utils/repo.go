package utils

import (
	"fmt"
	"strings"
)

func ParseRepo(repo string) (repoOwner string, repoName string, err error) {
	s := strings.Split(repo, "/")
	if len(s) != 2 {
		err = fmt.Errorf("repo has to be 'owner_name/repo_name' format, but got %s", repo)
		return
	}
	repoOwner, repoName = s[0], s[1]
	return
}
