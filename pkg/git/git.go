package git

import (
	"fmt"
	"os"
	"os/exec"
)

func GetGitExecutable() string {
	gitExucutable, err := exec.LookPath("git")

	if err != nil {
		fmt.Println("'git' is not installed")
		panic(err)
	}

	return gitExucutable
}

func CheckGitRepository() {
	if _, err := os.Stat("./.git"); os.IsNotExist(err) {
		fmt.Println("this directory is not a git repository")
		panic(err)
	}
}
