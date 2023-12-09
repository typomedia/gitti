package hdlr

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/gorilla/mux"
	"github.com/typomedia/gitti/app"
	"github.com/typomedia/gitti/app/git"
	"github.com/typomedia/gitti/app/git/ext"
	"github.com/typomedia/gitti/app/helper"
	"github.com/typomedia/gitti/app/msg"
	"github.com/typomedia/gitti/app/str"
	"net/http"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]

	// Get the repo from config
	repo := helper.GetRepo(project)

	branch := "master"
	if r.URL.Query().Get("branch") != "" {
		branch = r.URL.Query().Get("branch")
	}

	stash := false
	// If url has a query parameter "stash" set to true
	if r.URL.Query().Get("stash") == "true" {
		stash = true
	}

	// Always stash the changes, but pop only if stash is true
	ext.Stash(repo.Path)

	// Always reset the changes to prevent conflicts
	git.Reset(repo.Path)

	// Fetch the latest changes from remote
	git.Fetch(repo.Path)

	// Get a list of all the local branches
	branches := git.Branches(repo.Path)

	// Get a list of all the references
	refs := git.References(repo.Path)

	// Check if the remote branch exists
	exists := false
	err := refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().IsRemote() {

			if str.After(ref.Name().Short(), "/") == branch {
				exists = true

				// Create a temporary branch
				temp := app.App.Name + str.Hex()[:7]
				ext.Checkout(repo.Path, temp, "-b")
				// Check if the local branch exists
				err := branches.ForEach(func(ref *plumbing.Reference) error {
					if ref.Name().Short() == branch {
						res := ext.Branch(repo.Path, branch, "-D")
						msg.Info(res)
					}
					return nil
				})
				msg.Check(err)

				// Checkout the remote branch
				res := ext.Checkout(repo.Path, "origin/"+branch, "-t")
				msg.Info(res)
				fmt.Fprint(w, res)

				if stash {
					// Pop stashed changes
					ext.Pop(repo.Path)
				}

				// Delete the temporary branch
				res = ext.Branch(repo.Path, temp, "-D")
				msg.Info(res)
			}

		}
		return nil
	})
	msg.Check(err)

	if !exists {
		msg.Info("Branch %s does not exist!", branch)
		fmt.Fprintf(w, "Branch %s does not exist!", branch)
	}
}
