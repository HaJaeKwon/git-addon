package main

import (
	"bytes"
	"fmt"
	"git-addon/pkg/git"
	"os/exec"
)

//var HELP_COMMAND = []string{"help", "--help", "-h"}

func main() {
	gitExecutable := git.GetGitExecutable()

	git.CheckGitRepository()

	var cmd *exec.Cmd
	var stdout, stderr bytes.Buffer
	cmd = exec.Command(gitExecutable, "stash")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	fmt.Println("git stash start")
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Println(exitError.ExitCode())
		}
	}
	fmt.Println(string(stdout.Bytes()), string(stderr.Bytes()))

	//develop := flag.Bool("d", false, "checkout develop")
	//branch := flag.String("b", "", "checkout branch")
}
