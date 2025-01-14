package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func GetCurrentFolderName() string {
	// Get the current working directory
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return ""
	}

	// Get the folder name from the path
	return filepath.Base(currentPath)
}

func CreateFile(path, content string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Failed to create file %s: %v\n", path, err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Failed to write to file %s: %v\n", path, err)
	}
}

func CreateMultipleDirectories(base string, directories []string) error {
	for _, dir := range directories {
		path := filepath.Join(base, dir)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", path, err)
		}
	}

	return nil
}

// Copy a file from source to destination
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer destinationFile.Close()

	// Copy file content
	if _, err := io.Copy(destinationFile, sourceFile); err != nil {
		return fmt.Errorf("failed to copy content: %v", err)
	}

	return nil
}
