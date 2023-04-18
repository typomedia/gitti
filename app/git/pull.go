package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/typomedia/gitti/app/msg"
)

func Pull(path string) error {
	repo, err := git.PlainOpen(path)
	msg.Check(err)

	worktree, err := repo.Worktree()
	msg.Check(err)

	// err = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
	err = worktree.Pull(&git.PullOptions{})
	msg.Check(err)

	return err
}
