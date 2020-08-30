package main

import (
	"fmt"
	"git-addon/pkg/git"
	"os/exec"
)

//var HELP_COMMAND = []string{"help", "--help", "-h"}

func main() {
	gitExecutable := git.GetGitExecutable()

	git.CheckGitRepository()

	var cmd *exec.Cmd
	cmd = exec.Command(gitExecutable, "stash")
	stdoutStderr, _ := cmd.CombinedOutput()
	fmt.Println(string(stdoutStderr))

	//develop := flag.Bool("d", false, "checkout develop")
	//branch := flag.String("b", "", "checkout branch")
}
