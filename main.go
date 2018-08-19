package main

import (
	"flag"
	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"os"
)

func main() {
	cloneDirPtr := flag.String("clone-dir", os.Args[0], "Directory to clone")
	cloneUrlPtr := flag.String("clone-url", "https://github.com/flutter/flutter", "URL to clone")
	shaPtr := flag.String("sha", "", "sha to clone")
	flag.Parse()

	cloneOptions := git.CloneOptions{
		URL:           *cloneUrlPtr,
		ReferenceName: plumbing.ReferenceName("refs/heads/master"),
		SingleBranch:  true,
		Progress:      os.Stdout,
		Tags:          git.NoTags,
	}
	repo, err := git.PlainClone(*cloneDirPtr, false, &cloneOptions)
	CheckIfError(err)
	reference, err := repo.Head()
	CheckIfError(err)
	Info("Cloned! Head at %s", reference)

	workTree, err := repo.Worktree()
	CheckIfError(err)

	err = workTree.Reset(&git.ResetOptions{
		Commit: plumbing.NewHash(*shaPtr),
		Mode:   git.HardReset,
	})
	CheckIfError(err)
	Info("Hard reseted to %s", *shaPtr)

	status, err := workTree.Status()
	CheckIfError(err)
	Info("Status after reset: %s", status)

	repo.Storer.Index()
}
