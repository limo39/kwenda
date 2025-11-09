package main

import (
	"fmt"
	"kwenda/ast"
	"kwenda/lexer"
	"kwenda/parser"
)

func main() {
	// Test parsing empty dictionary
	code := `kamusi d = {}`
	tokens := lexer.Lex(code)
	
	fmt.Println("Tokens:")
	for i, tok := range tokens {
		fmt.Printf("%d: %s = %q\n", i, tok.Type, tok.Value)
	}
	
	result := parser.Parse(tokens)
	fmt.Printf("\nParsed result type: %T\n", result)
	fmt.Printf("Result: %+v\n", result)
	
	if dictDecl, ok := result.(ast.DictionaryDeclarationNode); ok {
		fmt.Printf("Dictionary name: %s\n", dictDecl.Name)
		fmt.Printf("Dictionary value type: %T\n", dictDecl.Value)
		fmt.Printf("Dictionary value: %+v\n", dictDecl.Value)
	}
}
