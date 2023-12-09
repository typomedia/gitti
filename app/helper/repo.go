package helper

import (
	"github.com/spf13/viper"
	"sort"
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

	// Sort the repos by name
	sort.Slice(repositories.Repos[:], func(i, j int) bool {
		return repositories.Repos[i].Name < repositories.Repos[j].Name
	})

	return repositories
}

func GetRepo(project string) Repo {
	repository := Repo{}
	repos := GetRepos()

	for _, repo := range repos.Repos {
		if repo.Name == project {
			repository = Repo{
				Name: repo.Name,
				Path: repo.Path,
			}
		}
	}

	return repository
}
