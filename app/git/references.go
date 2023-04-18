package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/typomedia/gitti/app/msg"
)

func References(path string) storer.ReferenceIter {
	repo, err := git.PlainOpen(path)
	msg.Check(err)

	refs, err := repo.References()
	msg.Check(err)

	return refs
}
