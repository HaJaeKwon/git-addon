package main

import (
	"bytes"
	"flag"
	"fmt"
	"git-addon/pkg/git"
	"os/exec"
)

//var HELP_COMMAND = []string{"help", "--help", "-h"}

var DEVELOP = "develop"
var MASTER = "master"

func main() {
	develop := flag.Bool("d", false, "checkout develop")
	master := flag.Bool("m", false, "checkout master")
	branch := flag.String("b", "", "checkout branch")
	flag.Parse()

	gitExecutable := git.GetGitExecutable()
	git.CheckGitRepository()

	stash(gitExecutable)

	if (*develop && *branch != "") ||
		(*develop && *master) ||
		(*master && *branch != "") {
		panic("use only one branch. develop or master or other branch")
	}
	if *develop {
		branch = &DEVELOP
	}
	if *master {
		branch = &MASTER
	}
	checkout(gitExecutable, branch)

	pull(gitExecutable)
}

func stash(gitExecutable string) {
	var cmd *exec.Cmd
	var stdout, stderr bytes.Buffer
	cmd = exec.Command(gitExecutable, "stash")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	fmt.Println("git stash start")
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Println(string(stderr.Bytes()))
			fmt.Println("ExitCode: " + string(rune(exitError.ExitCode())))
			panic("git stash error")
		}
	}
	fmt.Println(string(stdout.Bytes()), string(stderr.Bytes()))
}

func checkout(gitExecutable string, branch *string) {
	var cmd *exec.Cmd
	var stdout, stderr bytes.Buffer
	cmd = exec.Command(gitExecutable, "checkout", *branch)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	fmt.Printf("git checkout %s start\n", *branch)
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Println(string(stderr.Bytes()))
			fmt.Println("ExitCode: " + string(rune(exitError.ExitCode())))
			panic("git checkout " + *branch + " error")
		}
	}
	fmt.Println(string(stdout.Bytes()), string(stderr.Bytes()))
}

func pull(gitExecutable string) {
	var cmd *exec.Cmd
	var stdout, stderr bytes.Buffer
	cmd = exec.Command(gitExecutable, "pull")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	fmt.Println("git pull start")
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Println(string(stderr.Bytes()))
			fmt.Println("ExitCode: " + string(rune(exitError.ExitCode())))
			panic("git pull error")
		}
	}
	fmt.Println(string(stdout.Bytes()), string(stderr.Bytes()))
}
