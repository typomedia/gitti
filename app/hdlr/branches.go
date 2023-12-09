package hdlr

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/typomedia/gitti/app/git"
	"github.com/typomedia/gitti/app/helper"
	"github.com/typomedia/gitti/app/msg"
	"net/http"
)

func Branches(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]

	// Get the repo from config
	repo := helper.GetRepo(project)

	// Fetch the latest changes from remote
	git.Fetch(repo.Path)

	// Get a list of all the remote branches
	refs := git.RemoteBranches(repo.Path)

	branches := Refs{}
	for _, ref := range refs {
		if ref.Name().IsBranch() {
			branches.Branches = append(branches.Branches, Branch{
				Name:     ref.Name().Short(),
				Revision: ref.Hash().String(),
			})
		}
	}

	//json.NewEncoder(w).Encode(branches)
	jsonData, err := json.Marshal(branches)
	msg.Check(err)

	// set application/json header
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
