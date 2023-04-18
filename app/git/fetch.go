package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/typomedia/gitti/app/msg"
)

func Fetch(path string) {
	repo, err := git.PlainOpen(path)
	msg.Check(err)

	err = repo.Fetch(&git.FetchOptions{})
	msg.Check(err)

}
