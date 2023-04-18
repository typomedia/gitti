package app

import "fmt"

type Application struct {
	Name        string
	Version     string
	Author      string
	Description string
	Explanation string
	Banner      string
}

var App = Application{
	Name:        "gitti",
	Version:     "0.1.0",
	Author:      "Philipp Speck <philipp@typo.media>",
	Description: "Git HTTP Daemon",
	Explanation: "Git HTTP Daemon for managing multiple repositories via web hooks.",
	Banner: `╔═╗╦╔╦╗╔╦╗╦
║ ╦║ ║  ║ ║
╚═╝╩ ╩  ╩ ╩`,
}

func Logo() string {
	banner := fmt.Sprintf("%s\n", App.Banner)
	banner += fmt.Sprintf("%s %s\n", App.Name, App.Version)
	banner += App.Author

	return banner
}
