package helper

import (
	"github.com/spf13/viper"
)

type Repos struct {
	Repos []Repo
}

type Repo struct {
	Name string
	Path string
}

func GetRepos() Repos {
	config := viper.AllSettings()
	repos := config["repos"].(map[string]interface{})

	repositories := Repos{}
	for name, path := range repos {
		repositories.Repos = append(repositories.Repos, Repo{
			Name: name,
			Path: path.(string),
		})

	}

	return repositories
}

func GetRepo(project string) Repo {
	repo := Repo{}
	repos := GetRepos()

	for _, repo := range repos.Repos {
		if repo.Name == project {
			repo = Repo{
				Name: repo.Name,
				Path: repo.Path,
			}
		}
	}

	return repo
}
