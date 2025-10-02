package main

import (
    "fmt"
    "lugha-yangu/lexer"
    "lugha-yangu/parser"
    "lugha-yangu/interpreter"
    "os"
)

func main() {
    // Read the source code from a file
    input, err := os.ReadFile("examples/example1.swh")
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