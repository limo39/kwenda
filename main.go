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
    ast := parser.Parse(tokens)
    fmt.Println("AST:", ast)

    // Interpretation
    env := interpreter.NewEnvironment()
    result := interpreter.Interpret(ast, env)
    fmt.Println("Result:", result)
}