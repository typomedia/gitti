package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/typomedia/gitti/app/msg"
)

func Checkout(path, branch string) {

	// Fetch the latest changes from remote
	Fetch(path)

	repo, err := git.PlainOpen(path)
	msg.Check(err)

	worktree, err := repo.Worktree()
	msg.Check(err)

	err = worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(branch),
		Create: true,
		Force:  true,
	})
	msg.Check(err)
}
