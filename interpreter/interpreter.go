package interpreter

import (
	"fmt"
	"strconv"
	"lugha-yangu/ast"
)

// Special control flow values
type ControlFlow int

const (
	ControlNormal ControlFlow = iota
	ControlBreak
	ControlContinue
	ControlReturn
)

type ControlFlowResult struct {
	Type  ControlFlow
	Value interface{}
}

// Environment stores variables and their values
type Environment struct {
	Variables map[string]interface{}
	Functions map[string]ast.FunctionNode
	Parent    *Environment // For function scope
}

func NewEnvironment() *Environment {
	return &Environment{
		Variables: make(map[string]interface{}),
		Functions: make(map[string]ast.FunctionNode),
		Parent:    nil,
	}
}

func NewChildEnvironment(parent *Environment) *Environment {
	return &Environment{
		Variables: make(map[string]interface{}),
		Functions: parent.Functions, // Share functions with parent
		Parent:    parent,
	}
}

func (env *Environment) Set(name string, value interface{}) {
	env.Variables[name] = value
}

func (env *Environment) Get(name string) interface{} {
	if value, exists := env.Variables[name]; exists {
		return value
	}
	// Look in parent environment
	if env.Parent != nil {
		return env.Parent.Get(name)
	}
	return nil
}

func (env *Environment) SetFunction(name string, function ast.FunctionNode) {
	env.Functions[name] = function
}

func (env *Environment) GetFunction(name string) (ast.FunctionNode, bool) {
	function, exists := env.Functions[name]
	return function, exists
}

// toBool converts a value to boolean following Kwenda's rules
func toBool(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v
	case int:
		return v != 0
	case string:
		return v != ""
	default:
		return false
	}
}

func Interpret(node ast.ASTNode, env *Environment) interface{} {
	switch n := node.(type) {
	case ast.NumberNode:
		value, _ := strconv.Atoi(n.Value)
		return value

	case ast.BooleanNode:
		return n.Value

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
		
		// Handle logical operators first
		if n.Op == "na" || n.Op == "au" {
			// Convert to boolean
			leftBool := toBool(left)
			rightBool := toBool(right)
			
			switch n.Op {
			case "na": // AND
				return leftBool && rightBool
			case "au": // OR
				return leftBool || rightBool
			}
		}
		
		// Handle comparison operators that can work with booleans
		if n.Op == "==" || n.Op == "!=" {
			// If both are booleans, compare as booleans
			if leftBool, leftIsBool := left.(bool); leftIsBool {
				if rightBool, rightIsBool := right.(bool); rightIsBool {
					if n.Op == "==" {
						return leftBool == rightBool
					} else {
						return leftBool != rightBool
					}
				}
			}
		}
		
		// Convert to integers for arithmetic and numeric comparisons
		leftInt, leftOk := left.(int)
		if !leftOk {
			if str, ok := left.(string); ok {
				if num, err := strconv.Atoi(str); err == nil {
					leftInt = num
				} else {
					leftInt = 0
				}
			} else if b, ok := left.(bool); ok {
				if b {
					leftInt = 1
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
			} else if b, ok := right.(bool); ok {
				if b {
					rightInt = 1
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
		case "==":
			return leftInt == rightInt
		case "!=":
			return leftInt != rightInt
		case "<":
			return leftInt < rightInt
		case "<=":
			return leftInt <= rightInt
		case ">":
			return leftInt > rightInt
		case ">=":
			return leftInt >= rightInt
		case "=":
			// This should not happen in binary operations - assignment is handled in VariableDeclarationNode
			fmt.Println("Operesheni ya assignment haiwezi kuwa katika binary operation")
			return nil
		default:
			fmt.Println("Operesheni isiyojulikana:", n.Op)
			return nil
		}

	case ast.ReturnNode:
		if n.Value != nil {
			value := Interpret(n.Value, env)
			return ControlFlowResult{Type: ControlReturn, Value: value}
		}
		return ControlFlowResult{Type: ControlReturn, Value: nil}

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
		// Handle built-in function calls
		if n.Name == "andika" {
			for i, arg := range n.Args {
				result := Interpret(arg, env)
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Print(result)
			}
			fmt.Println()
			return nil
		}

		// Handle user-defined function calls
		if function, exists := env.GetFunction(n.Name); exists {
			// Create new environment for function execution
			funcEnv := NewChildEnvironment(env)

			// Evaluate arguments and bind to parameters
			for i, param := range function.Parameters {
				if i < len(n.Args) {
					argValue := Interpret(n.Args[i], env)
					funcEnv.Set(param.Name, argValue)
				}
			}

			// Execute function body
			var result interface{}
			for _, statement := range function.Body {
				result = Interpret(statement, funcEnv)

				// Check for return statement
				if cf, ok := result.(ControlFlowResult); ok {
					if cf.Type == ControlReturn {
						return cf.Value
					}
					// Other control flow (break/continue) should not escape function
				}
			}

			return result
		}

		fmt.Printf("Kazi '%s' haijulikani\n", n.Name)
		return nil

	case ast.VariableDeclarationNode:
		// Handle variable declarations (e.g., namba x = 10)
		value := Interpret(n.Value, env)
		env.Set(n.Name, value)
		return value

	case ast.IfNode:
		// Handle conditional statements (kama ... { ... } sivyo { ... })
		condition := Interpret(n.Condition, env)
		
		// Convert condition to boolean
		conditionBool := toBool(condition)
		
		if conditionBool {
			// Execute then body
			var result interface{}
			for _, statement := range n.ThenBody {
				result = Interpret(statement, env)
				// Propagate return statements
				if cf, ok := result.(ControlFlowResult); ok && cf.Type == ControlReturn {
					return result
				}
			}
			return result
		} else if len(n.ElseBody) > 0 {
			// Execute else body
			var result interface{}
			for _, statement := range n.ElseBody {
				result = Interpret(statement, env)
				// Propagate return statements
				if cf, ok := result.(ControlFlowResult); ok && cf.Type == ControlReturn {
					return result
				}
			}
			return result
		}
		return nil

	case ast.WhileNode:
		// Handle while loops (wakati condition { ... })
		var result interface{}
		for {
			condition := Interpret(n.Condition, env)
			
			// Convert condition to boolean
			conditionBool := toBool(condition)
			
			if !conditionBool {
				break
			}
			
			// Execute loop body
			shouldBreak := false
			for _, statement := range n.Body {
				result = Interpret(statement, env)
				
				// Check for control flow
				if cf, ok := result.(ControlFlowResult); ok {
					if cf.Type == ControlBreak {
						shouldBreak = true
						result = cf.Value
						break
					} else if cf.Type == ControlContinue {
						result = cf.Value
						break // Break inner loop to continue outer loop
					} else if cf.Type == ControlReturn {
						return result // Propagate return up
					}
				}
			}
			
			if shouldBreak {
				break
			}
		}
		return result

	case ast.ForNode:
		// Handle for loops (kwa init; condition; update { ... })
		var result interface{}
		
		// Execute initialization if present
		if n.Init != nil {
			Interpret(n.Init, env)
		}
		
		// Loop while condition is true
		for {
			// Check condition if present
			if n.Condition != nil {
				condition := Interpret(n.Condition, env)
				
				// Convert condition to boolean
				conditionBool := toBool(condition)
				
				if !conditionBool {
					break
				}
			}
			
			// Execute loop body
			shouldBreak := false
			shouldContinue := false
			for _, statement := range n.Body {
				result = Interpret(statement, env)
				
				// Check for control flow
				if cf, ok := result.(ControlFlowResult); ok {
					if cf.Type == ControlBreak {
						shouldBreak = true
						result = cf.Value
						break
					} else if cf.Type == ControlContinue {
						shouldContinue = true
						result = cf.Value
						break
					} else if cf.Type == ControlReturn {
						return result // Propagate return up
					}
				}
			}
			
			if shouldBreak {
				break
			}
			
			// Execute update if present and not continuing
			if !shouldContinue && n.Update != nil {
				Interpret(n.Update, env)
			} else if shouldContinue && n.Update != nil {
				// Still execute update on continue
				Interpret(n.Update, env)
			}
			
			// If no condition, break after first iteration to prevent infinite loop
			if n.Condition == nil {
				break
			}
		}
		return result

	case ast.BreakNode:
		// Handle break statements (vunja)
		return ControlFlowResult{Type: ControlBreak, Value: nil}

	case ast.ContinueNode:
		// Handle continue statements (endelea)
		return ControlFlowResult{Type: ControlContinue, Value: nil}

	case ast.FunctionNode:
		// Handle function definitions
		if n.Name == "kuu" {
			// Execute main function immediately
			var result interface{}
			for _, statement := range n.Body {
				result = Interpret(statement, env)
				
				// Handle return from main function
				if cf, ok := result.(ControlFlowResult); ok {
					if cf.Type == ControlReturn {
						return cf.Value
					}
					// Other control flow statements outside loops are ignored
					result = nil
				}
			}
			return result
		} else {
			// Store user-defined function
			env.SetFunction(n.Name, n)
			return nil
		}

	default:
		fmt.Println("Aina ya nodi haijulikani:", n)
		return nil
	}
}