package setup

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/tahaontech/cmng/utils"
)

// CreateSimpleProject creates a simple C++ project structure.
func CreateSimpleProject(projectName string, withTests bool) {
	fmt.Println("Setting up C++ project...")
	fmt.Println("Name: ", projectName)

	// Create project directories
	directories := []string{".vscode", "build", "cmake", "docs", "include", "src"}
	if withTests {
		directories = append(directories, "tests")
	}

	if err := utils.CreateMultipleDirectories(projectName, directories); err != nil {
		log.Fatalf("Error in create Project: %s,\nError: %s\n", projectName, err)
	}

	// Create files
	utils.CreateFile(filepath.Join(projectName, ".vscode/c_cpp_properties.json"), generateVsCodeConfig())
	utils.CreateFile(filepath.Join(projectName, "CMakeLists.txt"), generateCMakeLists(projectName, withTests))
	utils.CreateFile(filepath.Join(projectName, "README.md"), generateReadme(projectName))
	utils.CreateFile(filepath.Join(projectName, ".gitignore"), generateGitignore())
	utils.CreateFile(filepath.Join(projectName, "src/main.cpp"), generateMainCpp())

	// create Temp Class
	utils.CreateFile(filepath.Join(projectName, "include/my_class.h"), generateClassH())
	utils.CreateFile(filepath.Join(projectName, "src/my_class.cpp"), generateClassCpp())

	if withTests {
		utils.CreateFile(filepath.Join(projectName, "tests/CMakeLists.txt"), generateTestsCmake())
		utils.CreateFile(filepath.Join(projectName, "tests/test_my_class.cpp"), generateClassTest())
	}

	// Initialize Git repository
	if err := utils.InitGitRepo(projectName); err != nil {
		fmt.Printf("Failed to initialize Git repository: %v\n", err)
		return
	}

	// Init vcpkg package manager
	utils.CreateVCPKGManifest(projectName)
	utils.AddTestPkg(projectName)

	fmt.Println("C++ project setup complete!")
	fmt.Printf("for running: \ncd %s & cmng run\n", projectName)
}

func CreateAdvancedProject(projectName string) {
	fmt.Println("Not Implemented :)")
}
