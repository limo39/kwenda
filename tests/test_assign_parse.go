package main

import (
	"fmt"
	"kwenda/lexer"
	"kwenda/parser"
)

func main() {
	// Test parsing dictionary assignment
	code := `d["x"] = 20`
	tokens := lexer.Lex(code)
	
	fmt.Println("Tokens:")
	for i, tok := range tokens {
		fmt.Printf("%d: %s = %q\n", i, tok.Type, tok.Value)
	}
	
	result := parser.Parse(tokens)
	fmt.Printf("\nResult type: %T\n", result)
	fmt.Printf("Result: %+v\n", result)
}
