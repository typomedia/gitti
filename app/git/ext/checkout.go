package ext

import (
	"github.com/typomedia/gitti/app/msg"
	"github.com/typomedia/gitti/app/str"
	"os/exec"
)

func Checkout(path, branch, option string) string {
	if option == "" {
		option = "-b"
	}
	cmd := exec.Command("git", "-C", path, "checkout", option, branch)

	output, err := cmd.Output()
	msg.Check(err)

	return str.Trim(output)
}
