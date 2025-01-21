package parser

import (
	"lugha-yangu/ast"
	"lugha-yangu/lexer"
)

type ASTNode interface{}

type FunctionNode struct {
	Name      string
	Params    []string
	Body      []ASTNode
}

type ReturnNode struct {
	Value ASTNode
}

type BinaryOpNode struct {
	Left  ASTNode
	Op    string
	Right ASTNode
}

type NumberNode struct {
	Value string
}

type IdentifierNode struct {
	Value string
}

// Parse converts tokens into an AST
func Parse(tokens []lexer.Token) ast.ASTNode {
	// Placeholder: Add parsing logic here
	// For now, just return a simple AST node
	return ast.NumberNode{Value: "42"} // Example AST node
}