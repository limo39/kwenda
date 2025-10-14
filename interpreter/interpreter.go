package interpreter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"lugha-yangu/ast"
)

// Special control flow values
type ControlFlow int

const (
	ControlNormal ControlFlow = iota
	ControlBreak
	ControlContinue
	ControlReturn
	ControlThrow
)

type ControlFlowResult struct {
	Type  ControlFlow
	Value interface{}
}

// ErrorValue represents a runtime error
type ErrorValue struct {
	Message  string
	Context  string // Additional context about where the error occurred
}

// Environment stores variables and their values
type Environment struct {
	Variables map[string]interface{}
	Functions map[string]ast.FunctionNode
	Modules   map[string]*Environment // Module namespaces
	Parent    *Environment // For function scope
}

func NewEnvironment() *Environment {
	return &Environment{
		Variables: make(map[string]interface{}),
		Functions: make(map[string]ast.FunctionNode),
		Modules:   make(map[string]*Environment),
		Parent:    nil,
	}
}

func NewChildEnvironment(parent *Environment) *Environment {
	return &Environment{
		Variables: make(map[string]interface{}),
		Functions: parent.Functions, // Share functions with parent
		Modules:   parent.Modules,   // Share modules with parent
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
	case float64:
		return v != 0.0
	case string:
		return v != ""
	default:
		return false
	}
}

// toNumber converts a value to float64, returns (value, isFloat)
func toNumber(value interface{}) (float64, bool) {
	switch v := value.(type) {
	case int:
		return float64(v), false
	case float64:
		return v, true
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f, strings.Contains(v, ".")
		}
		return 0, false
	case bool:
		if v {
			return 1, false
		}
		return 0, false
	default:
		return 0, false
	}
}

func Interpret(node ast.ASTNode, env *Environment) interface{} {
	switch n := node.(type) {
	case ast.NumberNode:
		// Try to parse as float first
		if strings.Contains(n.Value, ".") {
			value, err := strconv.ParseFloat(n.Value, 64)
			if err == nil {
				return value
			}
		}
		// Fall back to integer
		value, _ := strconv.Atoi(n.Value)
		return value

	case ast.BooleanNode:
		return n.Value

	case ast.StringNode:
		return n.Value

	case ast.ArrayNode:
		// Handle array literals (e.g., [1, 2, 3])
		var elements []interface{}
		for _, element := range n.Elements {
			value := Interpret(element, env)
			elements = append(elements, value)
		}
		return elements

	case ast.ArrayDeclarationNode:
		// Handle array declarations (e.g., orodha namba x = [1, 2, 3])
		var elements []interface{}
		for _, element := range n.Elements {
			value := Interpret(element, env)
			elements = append(elements, value)
		}
		env.Set(n.Name, elements)
		return elements

	case ast.ArrayAccessNode:
		// Handle array access (e.g., arr[0])
		arrayValue := Interpret(n.Array, env)
		indexValue := Interpret(n.Index, env)
		
		if arr, ok := arrayValue.([]interface{}); ok {
			if idx, ok := indexValue.(int); ok {
				if idx >= 0 && idx < len(arr) {
					return arr[idx]
				}
			}
		}
		return nil

	case ast.ArrayAssignmentNode:
		// Handle array assignment (e.g., arr[0] = 5)
		arrayValue := Interpret(n.Array, env)
		indexValue := Interpret(n.Index, env)
		newValue := Interpret(n.Value, env)
		
		if arr, ok := arrayValue.([]interface{}); ok {
			if idx, ok := indexValue.(int); ok {
				if idx >= 0 && idx < len(arr) {
					arr[idx] = newValue
					return newValue
				}
			}
		}
		return nil

	case ast.StringVariableDeclarationNode:
		// Handle string variable declarations (e.g., maneno x = "habari")
		value := Interpret(n.Value, env)
		env.Set(n.Name, value)
		return value

	case ast.IdentifierNode:
		// Check if it's a string literal (starts and ends with quotes)
		if len(n.Value) >= 2 && n.Value[0] == '"' && n.Value[len(n.Value)-1] == '"' {
			return n.Value[1 : len(n.Value)-1] // Remove quotes
		}
		
		// Check if it's a module access (e.g., math.PI or math.ongeza_kubwa)
		if strings.Contains(n.Value, ".") {
			parts := strings.SplitN(n.Value, ".", 2)
			moduleName := parts[0]
			memberName := parts[1]
			
			if moduleEnv, exists := env.Modules[moduleName]; exists {
				// Try to get variable first
				if value := moduleEnv.Get(memberName); value != nil {
					return value
				}
				// Try to get function
				if function, exists := moduleEnv.GetFunction(memberName); exists {
					// Return a callable reference (we'll handle this in FunctionCallNode)
					return function
				}
			}
			// Module or member not found, return as-is for debugging
			return n.Value
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
		
		// Convert to numeric values (int or float64)
		leftFloat, leftIsFloat := toNumber(left)
		rightFloat, rightIsFloat := toNumber(right)
		
		// If either is float, use float arithmetic
		useFloat := leftIsFloat || rightIsFloat
		
		switch n.Op {
		case "+":
			// Handle string concatenation
			if leftStr, leftIsStr := left.(string); leftIsStr {
				if rightStr, rightIsStr := right.(string); rightIsStr {
					return leftStr + rightStr
				}
				// Convert right to string and concatenate
				return leftStr + fmt.Sprintf("%v", right)
			}
			if rightStr, rightIsStr := right.(string); rightIsStr {
				// Convert left to string and concatenate
				return fmt.Sprintf("%v", left) + rightStr
			}
			// Numeric addition
			if useFloat {
				return leftFloat + rightFloat
			}
			return int(leftFloat) + int(rightFloat)
		case "-":
			if useFloat {
				return leftFloat - rightFloat
			}
			return int(leftFloat) - int(rightFloat)
		case "*":
			if useFloat {
				return leftFloat * rightFloat
			}
			return int(leftFloat) * int(rightFloat)
		case "/":
			if rightFloat != 0 {
				// Division always returns float if either operand is float
				if useFloat {
					return leftFloat / rightFloat
				}
				// Integer division
				return int(leftFloat) / int(rightFloat)
			}
			return 0
		case "==":
			// Handle string comparison
			if leftStr, leftIsStr := left.(string); leftIsStr {
				if rightStr, rightIsStr := right.(string); rightIsStr {
					return leftStr == rightStr
				}
			}
			return leftFloat == rightFloat
		case "!=":
			// Handle string comparison
			if leftStr, leftIsStr := left.(string); leftIsStr {
				if rightStr, rightIsStr := right.(string); rightIsStr {
					return leftStr != rightStr
				}
			}
			return leftFloat != rightFloat
		case "<":
			return leftFloat < rightFloat
		case "<=":
			return leftFloat <= rightFloat
		case ">":
			return leftFloat > rightFloat
		case ">=":
			return leftFloat >= rightFloat
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
				// Special formatting for arrays
				if arr, ok := result.([]interface{}); ok {
					fmt.Print("[")
					for j, elem := range arr {
						if j > 0 {
							fmt.Print(", ")
						}
						fmt.Print(elem)
					}
					fmt.Print("]")
				} else {
					fmt.Print(result)
				}
			}
			fmt.Println()
			return nil
		}

		// Array manipulation functions
		if n.Name == "ongeza" && len(n.Args) == 2 {
			// Add element to array: ongeza(array, element)
			element := Interpret(n.Args[1], env)
			
			// Update the original array variable if it's an identifier
			if arrayNode, ok := n.Args[0].(ast.IdentifierNode); ok {
				arrayArg := env.Get(arrayNode.Value)
				if arr, ok := arrayArg.([]interface{}); ok {
					newArr := append(arr, element)
					env.Set(arrayNode.Value, newArr)
					return len(newArr) // Return new length
				}
			}
			return 0
		}

		if n.Name == "ondoa" && len(n.Args) == 2 {
			// Remove element at index: ondoa(array, index)
			indexArg := Interpret(n.Args[1], env)
			
			// Update the original array variable if it's an identifier
			if arrayNode, ok := n.Args[0].(ast.IdentifierNode); ok {
				arrayArg := env.Get(arrayNode.Value)
				if arr, ok := arrayArg.([]interface{}); ok {
					if idx, ok := indexArg.(int); ok {
						if idx >= 0 && idx < len(arr) {
							// Remove element at index
							newArr := append(arr[:idx], arr[idx+1:]...)
							env.Set(arrayNode.Value, newArr)
							return len(newArr) // Return new length
						}
					}
				}
			}
			return 0
		}

		if n.Name == "urefu_orodha" && len(n.Args) == 1 {
			// Get array length: urefu_orodha(array)
			arrayArg := Interpret(n.Args[0], env)
			if arr, ok := arrayArg.([]interface{}); ok {
				return len(arr)
			}
			return 0
		}

		if n.Name == "pata" && len(n.Args) == 2 {
			// Get element at index: pata(array, index)
			arrayArg := Interpret(n.Args[0], env)
			indexArg := Interpret(n.Args[1], env)
			
			if arr, ok := arrayArg.([]interface{}); ok {
				if idx, ok := indexArg.(int); ok {
					if idx >= 0 && idx < len(arr) {
						return arr[idx]
					} else {
						// Throw error for invalid index
						errorMsg := fmt.Sprintf("Index %d ni nje ya mipaka ya orodha (urefu: %d)", idx, len(arr))
						context := fmt.Sprintf("Katika kazi 'pata': Jaribu kutumia index kati ya 0 na %d", len(arr)-1)
						return ControlFlowResult{Type: ControlThrow, Value: ErrorValue{Message: errorMsg, Context: context}}
					}
				} else {
					return ControlFlowResult{Type: ControlThrow, Value: ErrorValue{Message: "Index lazima iwe namba", Context: "Katika kazi 'pata'"}}
				}
			}
			return ControlFlowResult{Type: ControlThrow, Value: ErrorValue{Message: "Hii si orodha", Context: "Katika kazi 'pata': Argument ya kwanza lazima iwe orodha"}}
		}

		// File I/O operations
		if n.Name == "soma" && len(n.Args) == 1 {
			// Read file: soma("filename.txt")
			filenameArg := Interpret(n.Args[0], env)
			if filename, ok := filenameArg.(string); ok {
				content, err := os.ReadFile(filename)
				if err != nil {
					// Throw an error instead of just printing
					errorMsg := fmt.Sprintf("Hitilafu ya kusoma faili '%s': %v", filename, err)
					context := "Katika kazi 'soma': Hakikisha faili ipo na una ruhusa ya kusoma"
					return ControlFlowResult{Type: ControlThrow, Value: ErrorValue{Message: errorMsg, Context: context}}
				}
				return string(content)
			}
			context := "Katika kazi 'soma': Argument lazima iwe jina la faili (maneno)"
			return ControlFlowResult{Type: ControlThrow, Value: ErrorValue{Message: "Jina la faili si sahihi", Context: context}}
		}

		if n.Name == "andika_faili" && len(n.Args) >= 2 {
			// Write to file: andika_faili("filename.txt", "content") or andika_faili("filename.txt", "content", kweli) for append
			filenameArg := Interpret(n.Args[0], env)
			contentArg := Interpret(n.Args[1], env)
			
			if filename, ok := filenameArg.(string); ok {
				// Convert content to string if it's not already
				var content string
				if str, ok := contentArg.(string); ok {
					content = str
				} else {
					content = fmt.Sprintf("%v", contentArg)
				}
				
				// Check if append mode is specified
				append := false
				if len(n.Args) >= 3 {
					appendArg := Interpret(n.Args[2], env)
					if appendVal, ok := appendArg.(bool); ok {
						append = appendVal
					}
				}
				
				var err error
				if append {
					// Append to file
					file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						fmt.Printf("Hitilafu ya kufungua faili '%s': %v\n", filename, err)
						return false
					}
					defer file.Close()
					
					_, err = file.WriteString(content)
				} else {
					// Overwrite file
					err = os.WriteFile(filename, []byte(content), 0644)
				}
				
				if err != nil {
					fmt.Printf("Hitilafu ya kuandika faili '%s': %v\n", filename, err)
					return false
				}
				return true
			}
			return false
		}

		if n.Name == "unda_faili" && len(n.Args) == 1 {
			// Create empty file: unda_faili("filename.txt")
			filenameArg := Interpret(n.Args[0], env)
			if filename, ok := filenameArg.(string); ok {
				file, err := os.Create(filename)
				if err != nil {
					fmt.Printf("Hitilafu ya kuunda faili '%s': %v\n", filename, err)
					return false
				}
				file.Close()
				return true
			}
			return false
		}

		if n.Name == "faili_ipo" && len(n.Args) == 1 {
			// Check if file exists: faili_ipo("filename.txt")
			filenameArg := Interpret(n.Args[0], env)
			if filename, ok := filenameArg.(string); ok {
				_, err := os.Stat(filename)
				return err == nil
			}
			return false
		}

		if n.Name == "ondoa_faili" && len(n.Args) == 1 {
			// Delete file: ondoa_faili("filename.txt")
			filenameArg := Interpret(n.Args[0], env)
			if filename, ok := filenameArg.(string); ok {
				err := os.Remove(filename)
				if err != nil {
					fmt.Printf("Hitilafu ya kuondoa faili '%s': %v\n", filename, err)
					return false
				}
				return true
			}
			return false
		}

		// String manipulation functions
		if n.Name == "urefu" && len(n.Args) == 1 {
			// Get string length
			arg := Interpret(n.Args[0], env)
			if str, ok := arg.(string); ok {
				return len(str)
			}
			return 0
		}

		if n.Name == "unganisha" && len(n.Args) >= 2 {
			// Concatenate strings
			var result strings.Builder
			for _, arg := range n.Args {
				value := Interpret(arg, env)
				if str, ok := value.(string); ok {
					result.WriteString(str)
				} else {
					result.WriteString(fmt.Sprintf("%v", value))
				}
			}
			return result.String()
		}

		if n.Name == "kata" && len(n.Args) >= 2 {
			// Substring function: kata(string, start) or kata(string, start, length)
			str := Interpret(n.Args[0], env)
			start := Interpret(n.Args[1], env)
			
			if strVal, ok := str.(string); ok {
				if startVal, ok := start.(int); ok {
					if startVal < 0 || startVal >= len(strVal) {
						return ""
					}
					
					if len(n.Args) == 3 {
						// kata(string, start, length)
						length := Interpret(n.Args[2], env)
						if lengthVal, ok := length.(int); ok {
							end := startVal + lengthVal
							if end > len(strVal) {
								end = len(strVal)
							}
							return strVal[startVal:end]
						}
					} else {
						// kata(string, start) - from start to end
						return strVal[startVal:]
					}
				}
			}
			return ""
		}

		if n.Name == "badilisha" && len(n.Args) == 3 {
			// Replace function: badilisha(string, old, new)
			str := Interpret(n.Args[0], env)
			old := Interpret(n.Args[1], env)
			new := Interpret(n.Args[2], env)
			
			if strVal, ok := str.(string); ok {
				if oldVal, ok := old.(string); ok {
					if newVal, ok := new.(string); ok {
						return strings.ReplaceAll(strVal, oldVal, newVal)
					}
				}
			}
			return str
		}

		if n.Name == "tafuta" && len(n.Args) == 2 {
			// Find function: tafuta(string, substring) - returns index or -1
			str := Interpret(n.Args[0], env)
			substr := Interpret(n.Args[1], env)
			
			if strVal, ok := str.(string); ok {
				if substrVal, ok := substr.(string); ok {
					return strings.Index(strVal, substrVal)
				}
			}
			return -1
		}

		if n.Name == "awali" && len(n.Args) == 2 {
			// Starts with function: awali(string, prefix) - returns boolean
			str := Interpret(n.Args[0], env)
			prefix := Interpret(n.Args[1], env)
			
			if strVal, ok := str.(string); ok {
				if prefixVal, ok := prefix.(string); ok {
					return strings.HasPrefix(strVal, prefixVal)
				}
			}
			return false
		}

		if n.Name == "mwisho" && len(n.Args) == 2 {
			// Ends with function: mwisho(string, suffix) - returns boolean
			str := Interpret(n.Args[0], env)
			suffix := Interpret(n.Args[1], env)
			
			if strVal, ok := str.(string); ok {
				if suffixVal, ok := suffix.(string); ok {
					return strings.HasSuffix(strVal, suffixVal)
				}
			}
			return false
		}

		if n.Name == "herufi_kubwa" && len(n.Args) == 1 {
			// Convert to uppercase: herufi_kubwa(string)
			str := Interpret(n.Args[0], env)
			if strVal, ok := str.(string); ok {
				return strings.ToUpper(strVal)
			}
			return str
		}

		if n.Name == "herufi_ndogo" && len(n.Args) == 1 {
			// Convert to lowercase: herufi_ndogo(string)
			str := Interpret(n.Args[0], env)
			if strVal, ok := str.(string); ok {
				return strings.ToLower(strVal)
			}
			return str
		}

		if n.Name == "ondoa_nafasi" && len(n.Args) == 1 {
			// Trim whitespace: ondoa_nafasi(string)
			str := Interpret(n.Args[0], env)
			if strVal, ok := str.(string); ok {
				return strings.TrimSpace(strVal)
			}
			return str
		}

		if n.Name == "gawanya_maneno" && len(n.Args) >= 1 {
			// Split string: gawanya_maneno(string) or gawanya_maneno(string, separator)
			str := Interpret(n.Args[0], env)
			if strVal, ok := str.(string); ok {
				if len(n.Args) == 2 {
					separator := Interpret(n.Args[1], env)
					if sepVal, ok := separator.(string); ok {
						parts := strings.Split(strVal, sepVal)
						// Return the number of parts for now (could be enhanced to return array)
						return len(parts)
					}
				} else {
					// Split by whitespace
					parts := strings.Fields(strVal)
					return len(parts)
				}
			}
			return 0
		}

		// Check if it's a module function call (e.g., math.ongeza_kubwa)
		if strings.Contains(n.Name, ".") {
			parts := strings.SplitN(n.Name, ".", 2)
			moduleName := parts[0]
			functionName := parts[1]
			
			if moduleEnv, exists := env.Modules[moduleName]; exists {
				if function, exists := moduleEnv.GetFunction(functionName); exists {
					// Create new environment for function execution
					funcEnv := NewChildEnvironment(moduleEnv)

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

						// Check for return statement or throw
						if cf, ok := result.(ControlFlowResult); ok {
							if cf.Type == ControlReturn {
								return cf.Value
							} else if cf.Type == ControlThrow {
								// Propagate throw from module function
								return cf
							}
							// Other control flow (break/continue) should not escape function
						}
					}

					return result
				}
			}
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

				// Check for return statement or throw
				if cf, ok := result.(ControlFlowResult); ok {
					if cf.Type == ControlReturn {
						return cf.Value
					} else if cf.Type == ControlThrow {
						// Propagate throw from user-defined function
						return cf
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

	case ast.TryNode:
		// Handle try-catch blocks (jaribu ... shika ...)
		var result interface{}
		var caughtError interface{}
		
		// Execute try block
		for _, statement := range n.TryBody {
			result = Interpret(statement, env)
			
			// Check for thrown errors or other control flow
			if cf, ok := result.(ControlFlowResult); ok {
				if cf.Type == ControlThrow {
					caughtError = cf.Value
					break
				} else if cf.Type == ControlReturn {
					// Return statements should propagate up
					if len(n.FinallyBody) > 0 {
						// Execute finally block before returning
						for _, statement := range n.FinallyBody {
							Interpret(statement, env)
						}
					}
					return result
				}
			}
		}
		
		// If an error was caught, execute catch block
		if caughtError != nil && len(n.CatchBody) > 0 {
			// Create new environment for catch block with error variable
			catchEnv := NewChildEnvironment(env)
			if n.CatchVar != "" {
				catchEnv.Set(n.CatchVar, caughtError)
			}
			
			// Execute catch block
			for _, statement := range n.CatchBody {
				result = Interpret(statement, catchEnv)
				
				// Handle control flow in catch block
				if cf, ok := result.(ControlFlowResult); ok {
					if cf.Type == ControlReturn {
						if len(n.FinallyBody) > 0 {
							// Execute finally block before returning
							for _, statement := range n.FinallyBody {
								Interpret(statement, env)
							}
						}
						return result
					}
				}
			}
		}
		
		// Execute finally block if present
		if len(n.FinallyBody) > 0 {
			for _, statement := range n.FinallyBody {
				finallyResult := Interpret(statement, env)
				
				// Finally block can override return values
				if cf, ok := finallyResult.(ControlFlowResult); ok {
					if cf.Type == ControlReturn {
						return finallyResult
					}
				}
			}
		}
		
		// If error wasn't caught, re-throw it
		if caughtError != nil && len(n.CatchBody) == 0 {
			return ControlFlowResult{Type: ControlThrow, Value: caughtError}
		}
		
		return result

	case ast.ThrowNode:
		// Handle throw statements (tupa)
		message := Interpret(n.Message, env)
		errorMsg := fmt.Sprintf("%v", message)
		return ControlFlowResult{Type: ControlThrow, Value: ErrorValue{Message: errorMsg}}

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
					} else if cf.Type == ControlThrow {
						// Unhandled error in main function
						if err, ok := cf.Value.(ErrorValue); ok {
							fmt.Printf("\n╔═══════════════════════════════════════════════════════════╗\n")
							fmt.Printf("║ HITILAFU (ERROR)                                          ║\n")
							fmt.Printf("╚═══════════════════════════════════════════════════════════╝\n")
							fmt.Printf("Ujumbe: %s\n", err.Message)
							if err.Context != "" {
								fmt.Printf("Muktadha: %s\n", err.Context)
							}
							fmt.Printf("\n")
						} else {
							fmt.Printf("Hitilafu isiyoshughulikiwa: %v\n", cf.Value)
						}
						return nil
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