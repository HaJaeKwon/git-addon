package main

import (
	"bytes"
	"fmt"
	"git-addon/pkg/git"
	"os"
	"os/exec"
)

/*
gituser ha

git config --replace-all user.name ${GIT_NAME_ha}
git config --replace-all user.email ${GIT_EMAIL_ha}
*/

var CONFIG = "config"
var SET = "--replace-all"
var USER_NAME = "user.name"
var USER_EMAIL = "user.email"
var ENV_PREFIX_NAME = "GIT_NAME_"
var ENV_PREFIX_EMAIL = "GIT_EMAIL_"

func main() {
	if len(os.Args) < 2 {
		panic("[GIT_ADDON] ENV for git config is must set. ex) gituser company")
	}

	gitExecutable := git.GetGitExecutable()
	git.CheckGitRepository()

	env := os.Args[1:2][0]
	if os.Getenv(ENV_PREFIX_NAME+env) == "" {
		panic("[GIT_ADDON] ENV is not set. please check 'echo $" + ENV_PREFIX_NAME + env + "'")
	}
	if os.Getenv(ENV_PREFIX_EMAIL+env) == "" {
		panic("[GIT_ADDON] ENV is not set. please check 'echo $" + ENV_PREFIX_EMAIL + env + "'")
	}

	cmd := exec.Command(gitExecutable, CONFIG, SET, USER_NAME, os.Getenv(ENV_PREFIX_NAME+env))
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(string(stdout.Bytes()), string(stderr.Bytes()))
		panic("[GIT_ADDON] 'git config --replace-all user.name' fail")
	}

	cmd = exec.Command(gitExecutable, CONFIG, SET, USER_EMAIL, os.Getenv(ENV_PREFIX_EMAIL+env))
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(string(stdout.Bytes()), string(stderr.Bytes()))
		panic("[GIT_ADDON] 'git config --replace-all user.email' fail")
	}
}
