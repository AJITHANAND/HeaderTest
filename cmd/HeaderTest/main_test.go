package main

import (
    "flag"
    "os"
    "path/filepath"
    "testing"
)

func TestMain(t *testing.T) {
    oldArgs := os.Args
    oldFlagCommandLine := flag.CommandLine
    oldDir, _ := os.Getwd()
    
    defer func() {
        os.Args = oldArgs
        flag.CommandLine = oldFlagCommandLine
        os.Chdir(oldDir)
    }()
    
    // Change to project root directory
    projectRoot, _ := filepath.Abs("../..")
    os.Chdir(projectRoot)
    
    os.Args = []string{"HeaderTest", "-config", "config.json"}
    flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
    main()
}