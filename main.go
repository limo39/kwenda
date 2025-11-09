package main

import (
    "fmt"
    "kwenda/lexer"
    "kwenda/parser"
    "kwenda/interpreter"
    "os"
    "strings"
)

// ModuleCache stores loaded modules
var moduleCache = make(map[string]*interpreter.Environment)

// LoadModule loads and parses a module file
func LoadModule(modulePath string) (*interpreter.Environment, error) {
    // Check if module is already loaded
    if env, exists := moduleCache[modulePath]; exists {
        return env, nil
    }
    
    // Read the module file
    input, err := os.ReadFile(modulePath)
    if err != nil {
        return nil, fmt.Errorf("cannot read module %s: %v", modulePath, err)
    }
    
    // Parse the module
    tokens := lexer.Lex(string(input))
    program := parser.ParseProgram(tokens)
    
    // Create module environment
    moduleEnv := interpreter.NewEnvironment()
    
    // Execute all top-level statements in the module (functions and variables)
    for _, node := range program.Functions {
        interpreter.Interpret(node, moduleEnv)
    }
    
    // Cache the module
    moduleCache[modulePath] = moduleEnv
    return moduleEnv, nil
}

// ProcessImports processes import statements in the source code
func ProcessImports(source string) (string, error) {
    lines := strings.Split(source, "\n")
    var processedLines []string
    
    for _, line := range lines {
        trimmed := strings.TrimSpace(line)
        if strings.HasPrefix(trimmed, "leta ") {
            // Extract module path
            parts := strings.Fields(trimmed)
            if len(parts) >= 2 {
                modulePath := strings.Trim(parts[1], "\"")
                
                // Load the module
                _, err := LoadModule(modulePath)
                if err != nil {
                    return "", err
                }
                
                // Skip the import line (don't include in processed code)
                continue
            }
        }
        processedLines = append(processedLines, line)
    }
    
    return strings.Join(processedLines, "\n"), nil
}

func printHelp() {
    help := `
╔═══════════════════════════════════════════════════════════════════════════╗
║                    KWENDA - Swahili Programming Language                  ║
║                    "Move Forward with Technology"                         ║
╚═══════════════════════════════════════════════════════════════════════════╝

USAGE:
    kwenda <filename.swh>              Run a Kwenda program
    kwenda --help                      Show this help message
    kwenda --version                   Show version information

DESCRIPTION:
    Kwenda is a fully-featured programming language with native Swahili syntax.
    It's designed to make programming accessible to Swahili speakers while
    providing modern programming capabilities.

EXAMPLES:
    kwenda hello.swh                   Run hello.swh program
    kwenda examples/demo.swh           Run demo from examples folder

BASIC SYNTAX:
    kazi kuu() {                       # Main function (entry point)
        andika("Habari!")              # Print to console
    }

KEYWORDS:
    kazi        - Function declaration
    kuu         - Main function name
    andika      - Print/output
    rudisha     - Return
    namba       - Number type
    maneno      - String type
    boolean     - Boolean type
    kamusi      - Dictionary type
    orodha      - Array/list type
    kama        - If statement
    sivyo       - Else statement
    wakati      - While loop
    kwa         - For loop
    darasa      - Class declaration
    unda        - Create/instantiate
    hii         - This/self reference
    lambda      - Anonymous function
    leta        - Import module

FEATURES:
    ✓ Variables and data types (numbers, strings, booleans)
    ✓ Functions with parameters and return values
    ✓ Control flow (if/else, while, for loops)
    ✓ Arrays and dictionaries
    ✓ Object-oriented programming (classes, inheritance)
    ✓ Lambda/anonymous functions
    ✓ Closures
    ✓ Module system
    ✓ Error handling (try/catch)
    ✓ Standard library functions

DOCUMENTATION:
    README.md                          Full documentation
    examples/                          Example programs
    
WEBSITE:
    https://github.com/limo39/kwenda

VERSION:
    Kwenda v1.0.0 - Built with Go

For more information, visit the documentation or run example programs.
`
    fmt.Println(help)
}

func printVersion() {
    version := `
Kwenda Programming Language
Version: 1.0.0
Built with: Go 1.23.3
License: MIT

Kwenda - "Move Forward" in Swahili
A bridge to technology for everyone.
`
    fmt.Println(version)
}

func main() {
    // Check for command line arguments
    if len(os.Args) < 2 {
        fmt.Println("Usage: kwenda <filename.swh>")
        fmt.Println("Try 'kwenda --help' for more information.")
        return
    }
    
    filename := os.Args[1]
    
    // Handle help flag
    if filename == "--help" || filename == "-h" {
        printHelp()
        return
    }
    
    // Handle version flag
    if filename == "--version" || filename == "-v" {
        printVersion()
        return
    }
    
    // Read the source code from a file
    input, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    
    // Process imports
    processedSource, err := ProcessImports(string(input))
    if err != nil {
        fmt.Println("Error processing imports:", err)
        return
    }

    // Lexical analysis
    tokens := lexer.Lex(processedSource)
    fmt.Println("Tokens:", tokens)

    // Parsing
    program := parser.ParseProgram(tokens)
    fmt.Println("Program AST:", program)

    // Interpretation
    env := interpreter.NewEnvironment()
    
    // Add loaded modules to main environment
    for modulePath, moduleEnv := range moduleCache {
        // Extract module name from path
        parts := strings.Split(modulePath, "/")
        moduleName := strings.TrimSuffix(parts[len(parts)-1], ".swh")
        
        // Store module environment in Modules map
        env.Modules[moduleName] = moduleEnv
    }
    
    var result interface{}
    
    // First pass: register all functions
    for _, function := range program.Functions {
        interpreter.Interpret(function, env)
    }
    
    // Second pass: execute main function if it exists
    if mainFunc, exists := env.GetFunction("kuu"); exists {
        result = interpreter.Interpret(mainFunc, env)
    }
    
    fmt.Println("Result:", result)
}