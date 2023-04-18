package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/typomedia/gitti/app/msg"
)

// Remotes Get a list of all the remote upstreams
func Remotes(path string) []*git.Remote {
	repo, err := git.PlainOpen(path)
	msg.Check(err)

	remotes, err := repo.Remotes()
	msg.Check(err)

	return remotes
}
