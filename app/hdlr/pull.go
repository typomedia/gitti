package hdlr

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/typomedia/gitti/app/git/ext"
	"github.com/typomedia/gitti/app/helper"
	"net/http"
)

func Pull(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]

	// Get the repo from config
	repo := helper.GetRepo(project)

	pull := ext.Pull(repo.Path)

	fmt.Fprintf(w, "%v\n", pull)

}
