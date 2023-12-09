package hdlr

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/typomedia/gitti/app/git/ext"
	"github.com/typomedia/gitti/app/helper"
	"net/http"
)

func Status(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]

	// Get the repo from config
	repo := helper.GetRepo(project)

	res := ext.Status(repo.Path)

	fmt.Fprintf(w, "%v\n", res)

}
