package ext

import (
	"github.com/typomedia/gitti/app/msg"
	"github.com/typomedia/gitti/app/str"
	"os/exec"
)

func Branch(path, branch, option string) string {
	if option == "" {
		option = "-l"
	}
	cmd := exec.Command("git", "-C", path, "branch", option, branch)

	output, err := cmd.Output()
	msg.Check(err)

	return str.Trim(output)
}
