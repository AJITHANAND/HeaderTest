package config

type Config struct {
	IncludeDirs        []string `json:"includeDirs"`
	ImportedIncludeDirs []string `json:"importedIncludeDirs"`
	Compiler          string   `json:"compiler"`
	CompilerArgs      []string `json:"compilerArgs"`
	BatchSize         int      `json:"batchSize"`
}

type CompileResult struct {
	HeaderFile string
	Error      error
}

