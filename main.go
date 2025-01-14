package main

import (
	"fmt"
	"os"

	"github.com/tahaontech/cmng/executer"
	"github.com/tahaontech/cmng/setup"
	"github.com/tahaontech/cmng/utils"

	"github.com/manifoldco/promptui"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "new":
		handleNew()
	case "build":
		handleBuild()
	case "run":
		handleRun()
	case "test":
		handleTest()
	case "help":
		showHelp()
	case "version":
		fmt.Println("cmng version:", version)
	default:
		fmt.Println("Unknown command:", command)
		showHelp()
	}
}

func handleNew() {
	// Prompt for project name
	projectName := utils.ScanLn("Enter the project name: ")

	formatedProjectName := utils.FormatProjectName(projectName)

	// Show template selection dialog
	templates := []string{"Simple", "Simple (with tests)", "Advanced"}
	prompt := promptui.Select{
		Label: "Select a project template",
		Items: templates,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Template selection failed: %v\n", err)
		return
	}

	// Create the project based on the selected template
	switch result {
	case "Simple":
		setup.CreateSimpleProject(formatedProjectName, false)
	case "Simple (with tests)":
		setup.CreateSimpleProject(formatedProjectName, true)
	case "Advanced":
		setup.CreateAdvancedProject(formatedProjectName)
	default:
		fmt.Println("Invalid template selected.")
	}
}

func handleBuild() {
	executer.BuildProject()
}

func handleRun() {
	executer.RunProject()
}

func handleTest() {
	executer.TestProject()
}

func showHelp() {
	fmt.Println("C/C++ Manage tool Help:")
	fmt.Println("Usage:")
	fmt.Println("  ./cmng [command]")
	fmt.Println()
	fmt.Println("Available Commands:")
	fmt.Println("  new       Create a new C++ project")
	fmt.Println("  build     Build current project")
	fmt.Println("  run       build & Run current project")
	fmt.Println("  test      build & Run current project unit tests")
	fmt.Println("  help      Show help for this CLI tool")
	fmt.Println("  version   Show the version of this CLI tool")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  ./cmng new")
	fmt.Println("  ./cmng version")
}
