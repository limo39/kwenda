package main

import (
    "fmt"
    "github.com/limo39/kwenda/lexer"
    "github.com/limo39/kwenda/parser"
    "github.com/limo39/kwenda/interpreter"
    "os"
)

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

    // Lexical analysis
    tokens := lexer.Lex(string(input))
    fmt.Println("Tokens:", tokens)

    // Parsing
    program := parser.ParseProgram(tokens)
    fmt.Println("Program AST:", program)

    // Interpretation
    env := interpreter.NewEnvironment()
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