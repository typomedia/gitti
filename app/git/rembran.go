package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/typomedia/gitti/app/msg"
)

// Deprecated: Use git.References instead.
func RemoteBranches(path string) []*plumbing.Reference {
	repo, err := git.PlainOpen(path)
	msg.Check(err)

	remote, _ := repo.Remote("origin")

	refs, _ := remote.List(&git.ListOptions{
		Auth:            nil,
		InsecureSkipTLS: true,
	})

	// Print the remote branches
	// for _, ref := range refs {
	//	if ref.Name().IsRemote() {
	//		fmt.Println(ref.Name().Short())
	//	}
	// }

	return refs
}
