package ext

import (
	"github.com/typomedia/gitti/app/msg"
	"github.com/typomedia/gitti/app/str"
	"os/exec"
)

func Status(path string) string {
	cmd := exec.Command("git", "-C", path, "status")

	output, err := cmd.Output()
	msg.Check(err)

	return str.Trim(output)
}
