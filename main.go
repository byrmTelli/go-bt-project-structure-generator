package main

import (
	"flag"
	"fmt"
	"go-bt-project-structure-generator/utils"
	"os"
	"path/filepath"
)

func main() {

	// Take the project name from user input
	projectName := flag.String("name", "btelliproject", "Project Name")

	// Check the users choice for installation of orm
	ormInstallationResult, err := utils.ConfirmInstallationORM()
	if err != nil {
		fmt.Println("Error selecting ORM:", err)
		return
	}

	flag.Parse()

	// Initialize the folder structure with users orm choice
	err = createProjectStructure(*projectName, ormInstallationResult)
	if err != nil {
		fmt.Println("Error occurred while creating project folder structure.\nError:", err)
	} else {
		fmt.Println("\nProject folder structure created succesfully!")
	}
}

func createProjectStructure(projectName string, ormInstall bool) error {

	var ormImport string

	if ormInstall {
		err := utils.InstallORM()
		if err != nil {
			fmt.Println("GORM installation error:", err)
			return err
		}
		ormImport = "gorm.io/gorm"
		fmt.Println("Installing GORM...")
	} else {
		fmt.Println("Continue without installing GORM packages...")
	}

	// Define directories for folder structure
	directories := []string{
		filepath.Join(projectName, "cmd"),
		filepath.Join(projectName, "database"),
		filepath.Join(projectName, "routes"),
		filepath.Join(projectName, "handlers"),
		filepath.Join(projectName, "models"),
		filepath.Join(projectName, "utils"),
		filepath.Join(projectName, "services"),
		filepath.Join(projectName, "middlewares"),
	}

	// Define files to be created
	files := map[string]string{
		filepath.Join(projectName, "cmd", "main.go"): `package main

import (
	"fmt"
	"` + ormImport + `" 
)

func main() {
	// This is main function of your application.
	// Write some code here and type "go run main.go" to the command line.

	fmt.Println("Hello, world!")
}`,

		filepath.Join(projectName, "database", "config.go"): `package database

import (
	"` + ormImport + `" 
)

// Here you can apply your database connection actions.
// Also you can create another files named related action to make more modular your project.
`,

		filepath.Join(projectName, "routes", "server.go"): `package routes

func Run() {
	// You can define port, serve and listen here.
}`,

		filepath.Join(projectName, "handlers", "handlers.go"): `package handlers

// You can define handlers for your routes to handle each request.
`,

		filepath.Join(projectName, "models", "dbmodels.go"): `package models

import (
	"` + ormImport + `" 
	)

// You can define your database models here.
`,

		filepath.Join(projectName, "models", "request.go"): `package models

import (
	"` + ormImport + `" 
	)

// You can define your request models here.
`,

		filepath.Join(projectName, "models", "view.go"): `package models

import (
	"` + ormImport + `" 
	)
		
// You can define your view models here.
`,

		filepath.Join(projectName, "utils", "utils.go"): `package utils

// You can define your tools here.
`,

		filepath.Join(projectName, "services", "services.go"): `package services

// You can define your services here. 
// For Instance: jwt_service, auth_service etc.
`,

		filepath.Join(projectName, "middlewares", "middleware.go"): `package middlewares

// You can define your middleware methods here.
// For Instance: jwt_middleware, auth_middleware etc.
`,

		filepath.Join(projectName, ".gitignore"): `*.exe
*.log
`,

		filepath.Join(projectName, "README.md"): "# " + projectName + "\n\nWelcome to another cool project!",
		filepath.Join(projectName, ".env"):      ``,
	}

	totalOperations := len(directories) + len(files)
	currentOperation := 0

	done := make(chan bool)
	go utils.StartSpinner(done)

	// Create directories with progress
	for _, dir := range directories {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Println("\nError occurred while creating file directory. Error:", err)
			done <- true
			return err
		}
		currentOperation++
		utils.PrintProgressBar(currentOperation, totalOperations)
	}

	// Initialize installation according to user's ORM choice
	if ormInstall {
		err := utils.InstallORM()
		if err != nil {
			fmt.Println("GORM installation error:", err)
			return err
		}
	}

	// Create files with progress
	for path, content := range files {
		err := os.WriteFile(path, []byte(content), 0644)
		if err != nil {
			done <- true
			return err
		}
		currentOperation++
		utils.PrintProgressBar(currentOperation, totalOperations)
	}

	done <- true

	return nil
}
