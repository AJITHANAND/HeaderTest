package main

import (
	"fmt"
	"os"
	"github.com/ajithanand/HeaderTest/internal/helper"
	"github.com/ajithanand/HeaderTest/config"
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
	fmt.Printf("All Header Files: %v\n", headerFilesArray)

	var compileResults []config.CompileResult
	if cfg.BatchSize == 0 {
		compileResults = helper.CompileHeadersWithWaitGroup(headerFilesArray, cfg.Compiler, cfg.IncludeDirs, cfg.ImportedIncludeDirs, cfg.CompilerArgs)
	}else{
		compileResults = helper.CompileHeadersWithWorkerPool(headerFilesArray, cfg.Compiler, cfg.IncludeDirs, cfg.ImportedIncludeDirs, cfg.CompilerArgs, cfg.BatchSize)
	}
	fmt.Println("Compile Results: ", compileResults)

}