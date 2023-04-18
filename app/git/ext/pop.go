package ext

import (
	"github.com/typomedia/gitti/app/msg"
	"github.com/typomedia/gitti/app/str"
	"os/exec"
)

func Pop(path string) string {
	cmd := exec.Command("git", "-C", path, "stash", "pop")

	output, err := cmd.Output()
	msg.Check(err)

	return str.Trim(output)
}
