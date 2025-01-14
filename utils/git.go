package utils

import (
	"fmt"
	"os/exec"
)

func InitGitRepo(projectName string) error {
	cmd := exec.Command("git", "init", projectName)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run 'git init': %w", err)
	}
	fmt.Println("Git initialized.")
	return nil
}
