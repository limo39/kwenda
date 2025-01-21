package interpreter

import (
	"lugha-yangu/ast"
	"strconv"
)

func Interpret(node ast.ASTNode) interface{} {
	switch n := node.(type) {
	case ast.NumberNode:
		value, _ := strconv.Atoi(n.Value)
		return value
	case ast.IdentifierNode:
		// Look up the identifier in the environment
		return 0 // Placeholder
	case ast.BinaryOpNode:
		left := Interpret(n.Left)
		right := Interpret(n.Right)
		switch n.Op {
		case "+":
			return left.(int) + right.(int)
		case "-":
			return left.(int) - right.(int)
		}
	case ast.ReturnNode:
		return Interpret(n.Value)
	}
	return nil
}