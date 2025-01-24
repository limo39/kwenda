package interpreter

import (
	"bufio"
	"fmt"
	"os"
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
		// Look up the identifier in the environment
		return env.Get(n.Value)

	case ast.BinaryOpNode:
		left := Interpret(n.Left, env)
		right := Interpret(n.Right, env)
		switch n.Op {
		case "+":
			return left.(int) + right.(int)
		case "-":
			return left.(int) - right.(int)
		case "*":
			return left.(int) * right.(int)
		case "/":
			return left.(int) / right.(int)
		default:
			fmt.Println("Operesheni isiyojulikana:", n.Op)
			return nil
		}

	case ast.ReturnNode:
		return Interpret(n.Value, env)

	case ast.InputNode:
		reader := bufio.NewReader(os.Stdin)
		if n.Prompt != "" {
			fmt.Print(n.Prompt + " ")
		} else {
			fmt.Print("Ingiza thamani: ")
		}
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1] // Remove the newline character

		// Try to convert the input to a number
		if num, err := strconv.Atoi(input); err == nil {
			return num
		}
		return input // Return as a string if not a number

	case ast.FunctionCallNode:
		// Handle function calls (e.g., andika(x, y))
		if n.Name == "andika" {
			for _, arg := range n.Args {
				fmt.Print(Interpret(arg, env), " ")
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
		for _, statement := range n.Body {
			Interpret(statement, env)
		}
		return nil

	default:
		fmt.Println("Aina ya nodi haijulikani:", n)
		return nil
	}
}