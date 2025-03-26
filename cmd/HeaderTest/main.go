package main

import (
	"fmt"
	"os"
	"github.com/ajithanand/HeaderTest/internal/helper"
	"github.com/ajithanand/HeaderTest/config"
	"flag"
)


func main() {
	fmt.Printf("Hello, world!\n")
	var cfg config.Config

	// parse the JSON file
	err := helper.ParseJSONFile("config.json", &cfg)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Config: %v\n", cfg)

	// get all header files in the include directories
	var headerFilesArray []string
	for _, Dir := range cfg.IncludeDirs {
		headerFiles := helper.GetHeaderFilesInDir(Dir)
		headerFilesArray = append(headerFilesArray, headerFiles...)
	}
	fmt.Printf("Found %d header files to compile\n", len(headerFilesArray))

	// Use the progress bar versions of the compilation functions
	var compileResults []config.CompileResult
	if cfg.BatchSize == 0 {
		compileResults = helper.CompileHeadersWithWaitGroupAndProgressBars(headerFilesArray, cfg.Compiler, cfg.IncludeDirs, cfg.ImportedIncludeDirs, cfg.CompilerArgs)
	} else {
		compileResults = helper.CompileHeadersWithWorkerPoolAndProgressBars(headerFilesArray, cfg.Compiler, cfg.IncludeDirs, cfg.ImportedIncludeDirs, cfg.CompilerArgs, cfg.BatchSize)
	}
	
	// Display a summary of compilation results
	fmt.Println("\nCompilation Results Summary:")
	errorCount := 0
	for _, result := range compileResults {
		if result.Error != nil {
			fmt.Printf("❌ %s: %v\n", result.HeaderFile, result.Error)
			errorCount++
		} else {
			fmt.Printf("✅ %s: compiled successfully\n", result.HeaderFile)
		}
	}
	fmt.Printf("\nSummary: %d successful, %d failed\n", len(compileResults)-errorCount, errorCount)
	if errorCount > 0 {
		os.Exit(1)
	}
}