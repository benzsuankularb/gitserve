package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	var repoName string
	var workDir string
	flag.StringVar(&repoName, "repo", "", "repositary name")
	flag.StringVar(&workDir, "dir", "", "work directory")
	flag.Parse()
	if repoName == "" || workDir == "" {
		fmt.Println("Error: repo & dir arguments are required")
		os.Exit(1)
	}
	gitRoot := os.Getenv("GITROOT")
	if gitRoot == "" {
		fmt.Println("Error: $GITROOT required")
		os.Exit(1)
	}
	gitDir := gitRoot + "/" + repoName + ".git"
	os.MkdirAll(gitDir, 0777)
	os.MkdirAll(workDir, 0777)
	args := []string{"init", "--bare", gitDir}
	if err := exec.Command("git", args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	data := []byte("#!/bin/sh\ngit --work-tree=" + workDir + " --git-dir=" + gitDir + " checkout -f")
	ioutil.WriteFile(gitDir+"/hooks/post-receive", data, 777)
	fmt.Println("Setup Done, .git directory is locate at :")
	fmt.Println(gitDir)
}
