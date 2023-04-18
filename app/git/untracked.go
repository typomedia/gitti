package git

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/typomedia/gitti/app/msg"
	"github.com/typomedia/gitti/app/str"
)

func Untracked(path string) []string {
	refs := References(path)

	locale := make(map[string]bool)
	remote := make(map[string]bool)

	err := refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().IsBranch() {
			locale[ref.Name().Short()] = true
		}

		if ref.Name().IsRemote() {
			short := str.After(ref.Name().Short(), "/")
			remote[short] = true
		}
		return nil
	})
	msg.Check(err)

	untracked := minus(locale, remote)

	var untrackedBranches []string
	for path := range untracked {
		untrackedBranches = append(untrackedBranches, path)
	}

	return untrackedBranches

}

// map a minus b
func minus(a map[string]bool, b map[string]bool) map[string]bool {
	var res = make(map[string]bool)
	for k, v := range a {
		if b[k] != v {
			res[k] = v
		}
	}
	return res
}
