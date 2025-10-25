package main

import (
	"fmt"
	"lugha-yangu/lexer"
	"lugha-yangu/parser"
)

func main() {
	// Test parsing dictionary literal
	code := `{}`
	tokens := lexer.Lex(code)
	
	fmt.Println("Tokens for '{}':")
	for i, tok := range tokens {
		fmt.Printf("%d: %s = %q\n", i, tok.Type, tok.Value)
	}
	
	result := parser.ParseDictionaryLiteral(tokens)
	fmt.Printf("\nResult type: %T\n", result)
	fmt.Printf("Result: %+v\n", result)
	
	// Test with values
	code2 := `{"name": "Amina"}`
	tokens2 := lexer.Lex(code2)
	
	fmt.Println("\nTokens for '{\"name\": \"Amina\"}':")
	for i, tok := range tokens2 {
		fmt.Printf("%d: %s = %q\n", i, tok.Type, tok.Value)
	}
	
	result2 := parser.ParseDictionaryLiteral(tokens2)
	fmt.Printf("\nResult2 type: %T\n", result2)
	fmt.Printf("Result2: %+v\n", result2)
}
