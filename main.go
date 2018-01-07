package main

import (
	"flag"
	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"log"
	"os"
	"time"
)

func main() {
	cloneDirPtr := flag.String("clone-dir", os.Args[0], "Directory to clone")
	cloneUrlPtr := flag.String("clone-url", "https://github.com/fkorotkov/go-git-clone-example", "URL to clone")
	refPtr := flag.String("ref", "refs/heads/master", "reference to clone")
	depthPtr := flag.Int("depth", 50, "depth to clone")
	singleBranchPtr := flag.Bool("single-branch", true, "clone only single branch")
	flag.Parse()

	cloneOptions := git.CloneOptions{
		URL:           *cloneUrlPtr,
		ReferenceName: plumbing.ReferenceName(*refPtr),
		Depth:         *depthPtr,
		SingleBranch:  *singleBranchPtr,
		Progress:      os.Stdout,
		Tags: 		   git.NoTags,	 
	}

	start := time.Now()
	Info("Options %s", cloneOptions)
	_, err := git.PlainClone(*cloneDirPtr, false, &cloneOptions)
	elapsed := time.Now().Sub(start)
	log.Printf("Cloned in %s", elapsed)

	CheckIfError(err)
}
