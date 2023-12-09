package hdlr

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/typomedia/gitti/app/git"
	"github.com/typomedia/gitti/app/helper"
	"net/http"
)

func Log(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]

	// Get the repo from config
	repo := helper.GetRepo(project)

	_, commit := git.Log(repo.Path)

	fmt.Fprintf(w, "%v\n", commit)

}
