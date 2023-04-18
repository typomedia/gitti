package hdlr

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/typomedia/gitti/app/git"
	"github.com/typomedia/gitti/app/msg"
	"net/http"
)

func Prune(w http.ResponseWriter, r *http.Request) {
	repo := Repo{}
	vars := mux.Vars(r)
	project := vars["project"]
	config := viper.AllSettings()
	repos := config["repos"].(map[string]interface{})

	for name, path := range repos {
		if name == project {
			repo = Repo{
				Name: name,
				Path: path.(string),
			}
		}
	}

	res := git.Prune(repo.Path)

	for _, r := range res {
		fmt.Fprintf(w, "Deleted: %v\n", r)
		msg.Info("Deleted: " + r)
	}

	// TODO: Not stable enough. In some cases it removes tracked branches.
	//res = git.Untracked(repo.Path)
	//
	//for _, v := range res {
	//	// soft delete untracked branch
	//	result := git.Branch(repo.Path, v, "-d")
	//
	//	if result != "" {
	//		fmt.Fprintf(w, "%v\n", result)
	//		msg.Info(fmt.Sprintf("%v", result))
	//	}
	//}
}
