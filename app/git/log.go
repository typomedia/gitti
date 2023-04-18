package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/typomedia/gitti/app/msg"
	_ "log"
)

func Log(path string) (*plumbing.Reference, *object.Commit) {
	repo, err := git.PlainOpen(path)
	msg.Check(err)

	head, err := repo.Head()
	msg.Check(err)

	// log.Println(head.Hash())
	// log.Println(head.Name().Short())

	commit, err := repo.CommitObject(head.Hash())
	msg.Check(err)

	// log.Println(commit.Message)
	// log.Println(commit.Author)

	return head, commit
}
