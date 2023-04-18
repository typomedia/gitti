package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/typomedia/gitti/app/msg"
)

func Branches(path string) storer.ReferenceIter {
	repo, err := git.PlainOpen(path)
	msg.Check(err)

	// Get a list of all the local branches
	branches, err := repo.Branches()
	msg.Check(err)

	return branches
}
