package hdlr

import (
	"fmt"
	"github.com/typomedia/gitti/app"
	"github.com/typomedia/gitti/app/git"
	"github.com/typomedia/gitti/app/helper"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, app.App.Banner)
	fmt.Fprintln(w, app.App.Name, app.App.Version)
	fmt.Fprintln(w, app.App.Author)
	fmt.Fprintf(w, "%v \n\n", app.App.Description)

	// Get the repo from config
	repos := helper.GetRepos()

	for _, repo := range repos.Repos {

		fmt.Fprintf(w, "%v: %v \n", repo.Name, repo.Path)

		res, msg := git.Repo(repo.Path)
		if !res {
			fmt.Fprintf(w, "%v \n\n", msg)
			continue
		}

		head, _ := git.Log(repo.Path)
		fmt.Fprintf(w, "%v %v\n\n", head.Name().Short(), head.Hash())
	}

}
