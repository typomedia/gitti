package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/typomedia/gitti/app/msg"
)

func Status(path string) git.Status {
	repo, err := git.PlainOpen(path)
	msg.Check(err)

	worktree, err := repo.Worktree()
	msg.Check(err)

	status, err := worktree.Status()
	msg.Check(err)

	// Check if the status is empty
	// if status.IsClean() {}

	return status
}
