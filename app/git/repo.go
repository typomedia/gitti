package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
)

func Repo(path string) (bool, string) {
	_, err := git.PlainOpen(path)

	var msg string

	if err != nil {
		if err == git.ErrRepositoryNotExists {
			msg = fmt.Sprintf("Error: %v\n", err)
			return false, msg
		}
	}

	msg = fmt.Sprintf("%s is a Git repository\n", path)

	return true, msg
}
