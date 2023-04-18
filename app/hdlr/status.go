package hdlr

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/typomedia/gitti/app/git/ext"
	"net/http"
)

func Status(w http.ResponseWriter, r *http.Request) {
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

	res := ext.Status(repo.Path)

	fmt.Fprintf(w, "%v\n", res)

}
