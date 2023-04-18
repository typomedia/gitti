package main

import (
	"github.com/typomedia/gitti/app/msg"
	"github.com/typomedia/gitti/cmd"
	"os"
	"os/exec"
)

func main() {
	// check if git is installed
	if err := exec.Command("git", "--version").Run(); err != nil {
		msg.Check(err)
		msg.Info("Git is not installed! Please install it first.")
		os.Exit(1)
	}
	cmd.Execute()
}
