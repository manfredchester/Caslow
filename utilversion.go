package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func version() interface{} {
	self := filepath.Base(os.Args[0])
	branch, hash, reversions := getGitInfo()
	return map[string]string{
		"app":     fmt.Sprintf("%s - RESTful MySQL query proxy", self),
		"version": fmt.Sprintf("V%d.%s %s", reversions, hash, branch),
	}
}

func getGitInfo() (branch, hash string, revisions int) {
	branch = "unkown"
	hash = "unkown"
	o, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return
	}
	branch = strings.TrimSpace(string(o))
	o, err = exec.Command("git", "log", "-n1", "--pretty=format:%h").Output()
	if err != nil {
		return
	}
	hash = string(o)
	o, err = exec.Command("git", "log", "--oneline").Output()
	if err != nil {
		return
	}
	revisions = len(strings.Split(string(o), "\n")) - 1
	return
}
