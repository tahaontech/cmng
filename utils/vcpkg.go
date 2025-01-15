package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func CreateVCPKGManifest(projName string) error {
	if err := os.Chdir(projName); err != nil {
		return fmt.Errorf("failed to change directory to %s: %v", projName, err)
	}

	if err := runCommand("vcpkg", "new", "--name="+projName, "--application"); err != nil {
		return fmt.Errorf("vcpkg command failed: %v", err)
	}

	if err := os.Chdir(".."); err != nil {
		return fmt.Errorf("failed to change back to the original directory: %v", err)
	}

	return nil
}

func AddTestPkg(projName string) error {
	if err := os.Chdir(projName); err != nil {
		return fmt.Errorf("failed to change directory to %s: %v", projName, err)
	}

	if err := runCommand("vcpkg", "add", "port", "fmt"); err != nil {
		return fmt.Errorf("vcpkg command failed: %v", err)
	}

	if err := os.Chdir(".."); err != nil {
		return fmt.Errorf("failed to change back to the original directory: %v", err)
	}

	return nil
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
