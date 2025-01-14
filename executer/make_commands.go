package executer

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/tahaontech/cmng/utils"
)

func BuildProject() {
	projName := getProjectName()
	fmt.Printf("Build %s...\n", projName)
	runCMakeBuild(projName)
}

func RunProject() {
	projName := getProjectName()
	fmt.Printf("Run %s...\n", projName)
	if err := runCMakeBuild(projName); err != nil {
		log.Fatalf("Error in Build: \n%+v\n", err)
	}
	execPath := "./build/" + projName
	runCommand(execPath)
}

func TestProject() {
	projName := getProjectName()
	fmt.Printf("Test %s...\n", projName)
	if err := runCMakeBuild(projName); err != nil {
		log.Fatalf("Error in Build: \n%+v\n", err)
	}

	testFiles, err := processTestFiles()
	if err != nil {
		log.Fatalf("Error in Build: \n%+v\n", err)
	}

	var TestsLen = len(testFiles)

	for idx, tf := range testFiles {
		fmt.Printf("=========== Run Test %d/%d ============\n", idx+1, TestsLen)
		runCommand("./" + tf)
		fmt.Println("")
	}
}

func getProjectName() string {
	return utils.FormatProjectName(utils.GetCurrentFolderName())
}

func runCMakeBuild(projName string) error {
	// Define directory paths
	cmakeDir := "cmake"
	buildDir := "build"

	// Step 1: Remove and recreate cmake and build directories
	dirs := []string{cmakeDir, buildDir}
	for _, dir := range dirs {
		if err := os.RemoveAll(dir); err != nil {
			return fmt.Errorf("failed to remove directory %s: %v", dir, err)
		}
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}

	// Step 2: Change to the cmake directory
	if err := os.Chdir(cmakeDir); err != nil {
		return fmt.Errorf("failed to change directory to %s: %v", cmakeDir, err)
	}

	// Step 3: Run CMake
	if err := runCommand("cmake", ".."); err != nil {
		return fmt.Errorf("cmake command failed: %v", err)
	}

	// Step 4: Run Make
	if err := runCommand("make"); err != nil {
		return fmt.Errorf("make command failed: %v", err)
	}

	// Step 5: Copy the binary to the build directory
	if err := runCommand("cp", projName, fmt.Sprintf("../%s", buildDir)); err != nil {
		return fmt.Errorf("failed to copy binary to build directory: %v", err)
	}

	// Step 6: Return to the original directory
	if err := os.Chdir(".."); err != nil {
		return fmt.Errorf("failed to change back to the original directory: %v", err)
	}

	fmt.Println("Build process completed successfully!")
	return nil
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// List and process .cpp files
func processTestFiles() ([]string, error) {
	// Define source and destination directories
	sourceDir := "tests"
	cmakeDir := "cmake/tests"
	destDir := "build/tests"

	// Ensure the destination directory exists
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create destination directory: %v", err)
	}

	testFiles := []string{}

	// Walk through the source directory to find .cpp files
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing %s: %v", path, err)
		}

		// Process only .cpp files
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".cpp") {
			// Remove .cpp suffix
			baseName := strings.TrimSuffix(info.Name(), ".cpp")

			// Source and destination paths
			sourcePath := filepath.Join(cmakeDir, baseName)
			destPath := filepath.Join(destDir, baseName)

			// Copy file
			if err := utils.CopyFile(sourcePath, destPath); err != nil {
				return fmt.Errorf("failed to copy %s to %s: %v", sourcePath, destPath, err)
			}
			testFiles = append(testFiles, destPath)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error processing test files: %v", err)
	}

	fmt.Println("All test files processed successfully!")
	return testFiles, nil
}
