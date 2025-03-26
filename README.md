# HeaderTest

HeaderTest is a Go-based utility designed to validate C/C++ header files by testing their compileability. It helps identify issues with header files such as missing includes, syntax errors, or dependency problems.

## Features

- **Header File Validation**: Test if your C/C++ header files can compile successfully
- **Parallel Processing**: Uses Go concurrency for efficient testing of multiple headers
- **Progress Visualization**: Displays real-time progress bars during compilation
- **Configurable**: Supports custom include paths, compiler options, and parallelism settings
- **Cross-platform**: Works on Windows, macOS, and Linux

## Installation

### From Binary Releases

Prebuilt binaries are available on the [GitHub Releases](https://github.com/ajithanand/HeaderTest/releases) page.

### From Source

```bash
# Clone the repository
git clone https://github.com/ajithanand/HeaderTest.git
cd HeaderTest

# Build the project
go build -o HeaderTest ./cmd/HeaderTest
```

## Usage

1. Create a `config.json` file (or use the provided sample):

```json
{
  "includeDirs":[
    "test/project/include",
    "test/project/include/common"
  ],
  "importedIncludeDirs":[
    "test/project/third_party/include"
  ],
  "compiler":"clang",
  "compilerArgs":[
    "-std=c++17"
  ],
  "batchSize": 0
}
```

2. Run the tool:

```bash
./HeaderTest -config config.json
```

## Configuration Options

The `config.json` file supports the following options:

| Option | Description |
|--------|-------------|
| `includeDirs` | Directories to search for header files |
| `importedIncludeDirs` | Additional include directories (e.g., third-party libraries) |
| `compiler` | C/C++ compiler to use (e.g., "clang", "gcc", "g++") |
| `compilerArgs` | Arguments passed to the compiler |
| `batchSize` | Controls parallelism (0 = use wait groups, >0 = use worker pool with that many workers) |

## How It Works

1. The tool scans the specified include directories for header files (.h, .hpp, .hxx, .hh)
2. For each header file, it creates a temporary .cpp file that includes the header
3. It attempts to compile the temporary file using the specified compiler
4. It reports success or failure for each header file

## Example Output

```
Config: {[test/project/include test/project/include/common] [test/project/third_party/include] clang [-std=c++17] 0}
Found 3 header files to compile
✅ test.hpp: compiled successfully
❌ test2.hpp: compilation failed: exit status 1
Output: In file included from /tmp/temp_294872919.cpp:1:
test2.hpp:5:10: fatal error: 'sdfdfddfd.hpp' file not found
#include "sdfdfddfd.hpp"
         ^~~~~~~~~~~~~~~
1 error generated.
✅ common.hpp: compiled successfully

Summary: 2 successful, 1 failed
```

## License

HeaderTest is licensed under the BSD 3-Clause License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

## Requirements

- Go 1.24 or higher
- A C/C++ compiler (clang, gcc, etc.)
