package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/typomedia/gitti/app"
	"github.com/typomedia/gitti/app/msg"
	"strings"
)

// Prune removes all temporary branches created by Gitti
func Prune(path string) []string {
	repo, err := git.PlainOpen(path)
	msg.Check(err)

	// err = repo.Prune(git.PruneOptions{})
	// msg.Check(err)

	branches := Branches(path)

	// list of array with branch names
	var res []string
	err = branches.ForEach(func(ref *plumbing.Reference) error {
		// TODO: Make it more reliable with regex
		if strings.Contains(ref.Name().String(), app.App.Name) {
			err := repo.Storer.RemoveReference(ref.Name())
			msg.Check(err)
			res = append(res, ref.Name().Short())
		}
		return nil
	})
	msg.Check(err)

	return res
}
