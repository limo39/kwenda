package interpreter

import (
	"fmt"
	"strconv"
	"lugha-yangu/ast"
)

// Environment stores variables and their values
type Environment struct {
	Variables map[string]interface{}
}

func NewEnvironment() *Environment {
	return &Environment{
		Variables: make(map[string]interface{}),
	}
}

func (env *Environment) Set(name string, value interface{}) {
	env.Variables[name] = value
}

func (env *Environment) Get(name string) interface{} {
	return env.Variables[name]
}

func Interpret(node ast.ASTNode, env *Environment) interface{} {
	switch n := node.(type) {
	case ast.NumberNode:
		value, _ := strconv.Atoi(n.Value)
		return value

	case ast.IdentifierNode:
		// Check if it's a string literal (starts and ends with quotes)
		if len(n.Value) >= 2 && n.Value[0] == '"' && n.Value[len(n.Value)-1] == '"' {
			return n.Value[1 : len(n.Value)-1] // Remove quotes
		}
		// Look up the identifier in the environment
		value := env.Get(n.Value)
		if value == nil {
			// If not found in environment, return the identifier name itself (for debugging)
			return n.Value
		}
		return value

	case ast.BinaryOpNode:
		left := Interpret(n.Left, env)
		right := Interpret(n.Right, env)
		
		// Convert to integers if they're not already
		leftInt, leftOk := left.(int)
		if !leftOk {
			if str, ok := left.(string); ok {
				if num, err := strconv.Atoi(str); err == nil {
					leftInt = num
				} else {
					leftInt = 0
				}
			} else {
				leftInt = 0
			}
		}
		
		rightInt, rightOk := right.(int)
		if !rightOk {
			if str, ok := right.(string); ok {
				if num, err := strconv.Atoi(str); err == nil {
					rightInt = num
				} else {
					rightInt = 0
				}
			} else {
				rightInt = 0
			}
		}
		
		switch n.Op {
		case "+":
			result := leftInt + rightInt
			return result
		case "-":
			return leftInt - rightInt
		case "*":
			return leftInt * rightInt
		case "/":
			if rightInt != 0 {
				return leftInt / rightInt
			}
			return 0
		default:
			fmt.Println("Operesheni isiyojulikana:", n.Op)
			return nil
		}

	case ast.ReturnNode:
		return Interpret(n.Value, env)

	case ast.InputNode:
		if n.Prompt != "" {
			fmt.Print(n.Prompt + " ")
		} else {
			fmt.Print("Ingiza thamani: ")
		}
		
		var input string
		fmt.Scanln(&input)

		// Always try to convert the input to a number for namba variables
		if num, err := strconv.Atoi(input); err == nil {
			return num
		}
		// If conversion fails, return 0 for numeric operations
		return 0

	case ast.FunctionCallNode:
		// Handle function calls (e.g., andika(x, y))
		if n.Name == "andika" {
			for i, arg := range n.Args {
				result := Interpret(arg, env)
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Print(result)
			}
			fmt.Println()
		}
		return nil

	case ast.VariableDeclarationNode:
		// Handle variable declarations (e.g., namba x = 10)
		value := Interpret(n.Value, env)
		env.Set(n.Name, value)
		return value

	case ast.FunctionNode:
		// Handle function definitions (e.g., kazi kuu() { ... })
		// If this is the main function (kuu), execute it immediately
		if n.Name == "kuu" {
			var result interface{}
			for _, statement := range n.Body {
				result = Interpret(statement, env)
			}
			return result
		}
		// For other functions, just store them (not implemented yet)
		return nil

	default:
		fmt.Println("Aina ya nodi haijulikani:", n)
		return nil
	}
}