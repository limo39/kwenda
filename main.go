package main

import (
    "fmt"
    "lugha-yangu/lexer"
    "lugha-yangu/parser"
    "lugha-yangu/interpreter"
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

func main() {
    // Check for command line arguments
    if len(os.Args) < 2 {
        fmt.Println("Usage: kwenda <filename.swh>")
        return
    }
    
    filename := os.Args[1]
    
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