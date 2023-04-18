package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/typomedia/gitti/app/msg"
)

func Clean(path string) {
	repo, err := git.PlainOpen(path)
	msg.Check(err)

	worktree, err := repo.Worktree()
	msg.Check(err)

	err = worktree.Clean(&git.CleanOptions{
		Dir: true,
	})
	msg.Check(err)
}
