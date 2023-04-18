package hdlr

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/typomedia/gitti/app"
	"github.com/typomedia/gitti/app/git"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	repo := Repo{}
	config := viper.AllSettings()
	repos := config["repos"].(map[string]interface{})

	fmt.Fprintln(w, app.App.Banner)
	fmt.Fprintln(w, app.App.Name, app.App.Version)
	fmt.Fprintln(w, app.App.Author)
	fmt.Fprintf(w, "%v \n\n", app.App.Description)

	for name, path := range repos {
		repo = Repo{
			Name: name,
			Path: path.(string),
		}

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
