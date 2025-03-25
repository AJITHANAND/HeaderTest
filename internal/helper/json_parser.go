package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"github.com/ajithanand/HeaderTest/config"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// ParseJSONFile parses a JSON file and returns a Config struct data

func ParseJSONFile(filename string, cfg *config.Config) (error) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(cfg)
	check(err)

	return nil
}

func GetHeaderFilesInDir(directoryName string) ([]string) {
	var headerFiles []string
	files, err := os.ReadDir(directoryName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for _, file := range files {
		if strings.Contains(file.Name(), ".") == false {
			continue
		}
		extention := strings.ToLower(file.Name()[strings.LastIndex(file.Name(), "."):])
		switch extention {
		case ".h":
			headerFiles = append(headerFiles, file.Name())
		case ".hpp":
			headerFiles = append(headerFiles, file.Name())
		case ".hxx":
			headerFiles = append(headerFiles, file.Name())
		case ".hh":
			headerFiles = append(headerFiles, file.Name())
		default:
			continue
		}
	}
	return headerFiles
}

func checkCompilerIsAvailable(compilerName string) (bool) {
	_,err :=exec.LookPath(compilerName)
	return err == nil
}
func CompileHeaderFile(headerFile string, compiler string, includeDirs []string, importedIncludeDirs []string, compilerArgs []string) error {
    if checkCompilerIsAvailable(compiler) == false {
        fmt.Printf("Compiler %s is not available\n", compiler)
        return fmt.Errorf("compiler not available")
    }
    fmt.Printf("Compiling %s with %s\n", headerFile, compiler)

    tempFile, err := os.CreateTemp("", "temp_*.cpp")
    check(err)
    tempFilePath := tempFile.Name()
    defer os.Remove(tempFilePath) // Keep this defer

    fmt.Fprintf(tempFile, "#include \"%s\"\n", headerFile)
    tempFile.Close()
    
    args := []string{"-c", tempFilePath}
    
    includeFlags := compilerIncludePathFormat(includeDirs)
    importedIncludeFlags := compilerIncludePathFormat(importedIncludeDirs)

    // fmt.Println("Include Paths:", includeFlags)
    // fmt.Println("Imported Include Paths:", importedIncludeFlags)
    
    args = append(args, includeFlags...)
    args = append(args, importedIncludeFlags...)
    args = append(args, compilerArgs...)

    // fmt.Println("Compiler Args:", args)
    

    cmd := exec.Command(compiler, args...)
    output, err := cmd.CombinedOutput()
	if err != nil {
        return fmt.Errorf("compilation failed: %v\nOutput: %s", err, output)
    }
    
    objFile := strings.Split(tempFilePath,"/")[len(strings.Split(tempFilePath,"/"))-1]
	objFile = strings.Split(objFile,".")[0] + ".o"
	
	err = os.Remove(objFile)
	if err != nil {
		return fmt.Errorf("failed to delete object file: %v", err)
	}
    
    return nil
}

func CompileHeadersWithWorkerPool(headers []string, compiler string, includeDirs []string, importedIncludeDirs []string, compilerArgs []string, numWorkers int) []config.CompileResult {
    jobs := make(chan string, len(headers))
    results := make(chan config.CompileResult, len(headers))
    
    // Start workers
    for w := 1; w <= numWorkers; w++ {
        go func() {
            for headerFile := range jobs {
                err := CompileHeaderFile(headerFile, compiler, includeDirs, importedIncludeDirs, compilerArgs)
                results <- config.CompileResult{
                    HeaderFile: headerFile, 
                    Error: err,
                }
            }
        }()
    }
    
    // Send jobs
    for _, header := range headers {
        jobs <- header
    }
    close(jobs)
    
    // Collect all results, not just errors
    var compileResults []config.CompileResult
    for i := 0; i < len(headers); i++ {
        result := <-results
        compileResults = append(compileResults, result)
    }
    
    return compileResults
}

func compilerIncludePathFormat(includePath []string) []string {
    var includeFlags []string
    currentWorkingDirectory, err := os.Getwd()
    check(err)
    for _, path := range includePath {
        includeFlags = append(includeFlags, "-I"+currentWorkingDirectory+"/"+path)
    }
    return includeFlags
}


func CompileHeadersWithWaitGroup(headers []string, compiler string, includeDirs []string, importedIncludeDirs []string, compilerArgs []string) []config.CompileResult {
	var compileResults []config.CompileResult
	var WaitGroup sync.WaitGroup

	for _, headerFile := range headers {
		WaitGroup.Add(1)
		go func(headerFile string) {
			defer WaitGroup.Done()
			err := CompileHeaderFile(headerFile, compiler, includeDirs, importedIncludeDirs, compilerArgs)
			compileResults = append(compileResults, config.CompileResult{
				HeaderFile: headerFile,
				Error:      err,
			})
		}(headerFile)
	}
	WaitGroup.Wait()
	return compileResults
}