package interpreter

import (
	"fmt"
	"kwenda/ast"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
	Message string
	Context string // Additional context about where the error occurred
}

// Environment stores variables and their values
type Environment struct {
	Variables map[string]interface{}
	Functions map[string]ast.FunctionNode
	Classes   map[string]ast.ClassNode // Class definitions
	Modules   map[string]*Environment  // Module namespaces
	Parent    *Environment             // For function scope
}

func NewEnvironment() *Environment {
	return &Environment{
		Variables: make(map[string]interface{}),
		Functions: make(map[string]ast.FunctionNode),
		Classes:   make(map[string]ast.ClassNode),
		Modules:   make(map[string]*Environment),
		Parent:    nil,
	}
}

func NewChildEnvironment(parent *Environment) *Environment {
	return &Environment{
		Variables: make(map[string]interface{}),
		Functions: parent.Functions, // Share functions with parent
		Classes:   parent.Classes,   // Share classes with parent
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

func (env *Environment) SetClass(name string, class ast.ClassNode) {
	env.Classes[name] = class
}

func (env *Environment) GetClass(name string) (ast.ClassNode, bool) {
	class, exists := env.Classes[name]
	return class, exists
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

	case ast.DictionaryNode:
		// Handle dictionary literals (e.g., {"key": "value", "age": 25})
		dict := make(map[string]interface{})
		for _, pair := range n.Pairs {
			key := Interpret(pair.Key, env)
			value := Interpret(pair.Value, env)
			// Convert key to string
			keyStr := fmt.Sprintf("%v", key)
			dict[keyStr] = value
		}
		return dict

	case ast.DictionaryDeclarationNode:
		// Handle dictionary declarations (e.g., kamusi data = {})
		dictValue := Interpret(n.Value, env)
		env.Set(n.Name, dictValue)
		return dictValue

	case ast.ArrayNode:
		// Handle array literals (e.g., [1, 2, 3])
		var elements []interface{}
		for _, element := range n.Elements {
			value := Interpret(element, env)
			elements = append(elements, value)
		}
		return elements

	case ast.ListComprehensionNode:
		// Handle list comprehensions (e.g., [x * 2 kwa x katika namba kama x > 5])
		// Evaluate the iterable
		iterableValue := Interpret(n.Iterable, env)
		
		// Create result array
		var result []interface{}
		
		// Check if iterable is an array
		if arr, ok := iterableValue.([]interface{}); ok {
			// Create a new environment for the comprehension scope
			compEnv := NewChildEnvironment(env)
			
			// Iterate over the array
			for _, item := range arr {
				// Set the loop variable
				compEnv.Set(n.Variable, item)
				
				// Check the condition if present
				if n.Condition != nil {
					conditionValue := Interpret(n.Condition, compEnv)
					// Skip if condition is false
					if !toBool(conditionValue) {
						continue
					}
				}
				
				// Evaluate the expression for this item
				value := Interpret(n.Expression, compEnv)
				result = append(result, value)
			}
			
			return result
		}
		
		// If iterable is not an array, return empty array
		return []interface{}{}

	case ast.ArrayDeclarationNode:
		// Handle array declarations (e.g., orodha namba x = [1, 2, 3])
		// Special case: if the first element is a ListComprehensionNode, evaluate it directly
		if len(n.Elements) == 1 {
			if _, ok := n.Elements[0].(ast.ListComprehensionNode); ok {
				// Evaluate the list comprehension and use its result
				result := Interpret(n.Elements[0], env)
				if arr, ok := result.([]interface{}); ok {
					env.Set(n.Name, arr)
					return arr
				}
			}
		}
		
		// Regular array declaration
		var elements []interface{}
		for _, element := range n.Elements {
			value := Interpret(element, env)
			elements = append(elements, value)
		}
		env.Set(n.Name, elements)
		return elements

	case ast.ArrayAccessNode:
		// Handle array access (e.g., arr[0]) or dictionary access (e.g., dict["key"])
		arrayValue := Interpret(n.Array, env)
		indexValue := Interpret(n.Index, env)

		// Check if it's a dictionary
		if dict, ok := arrayValue.(map[string]interface{}); ok {
			keyStr := fmt.Sprintf("%v", indexValue)
			if value, exists := dict[keyStr]; exists {
				return value
			}
			return nil
		}

		// Otherwise treat as array
		if arr, ok := arrayValue.([]interface{}); ok {
			if idx, ok := indexValue.(int); ok {
				if idx >= 0 && idx < len(arr) {
					return arr[idx]
				}
			}
		}
		return nil

	case ast.ArrayAssignmentNode:
		// Handle array assignment (e.g., arr[0] = 5) or dictionary assignment (e.g., dict["key"] = value)
		arrayValue := Interpret(n.Array, env)
		indexValue := Interpret(n.Index, env)
		newValue := Interpret(n.Value, env)

		// Check if it's a dictionary
		if dict, ok := arrayValue.(map[string]interface{}); ok {
			keyStr := fmt.Sprintf("%v", indexValue)
			dict[keyStr] = newValue
			return newValue
		}

		// Otherwise treat as array
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

	case ast.ThisNode:
		// Handle 'hii' keyword (this/self)
		value := env.Get("hii")
		if value == nil {
			return ControlFlowResult{Type: ControlThrow, Value: ErrorValue{Message: "'hii' inaweza kutumika tu ndani ya darasa (this can only be used inside a class)"}}
		}
		return value

	case ast.MemberAccessNode:
		// Handle member access (e.g., hii.jina or object.property)
		objectValue := Interpret(n.Object, env)
		if dict, ok := objectValue.(map[string]interface{}); ok {
			if value, exists := dict[n.Member]; exists {
				return value
			}
			return nil
		}
		return nil

	case ast.MemberAssignmentNode:
		// Handle member assignment (e.g., hii.jina = "Amina")
		objectValue := Interpret(n.Object, env)
		newValue := Interpret(n.Value, env)
		if dict, ok := objectValue.(map[string]interface{}); ok {
			dict[n.Member] = newValue
			return newValue
		}
		return nil

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

	case ast.MethodCallNode:
		// Handle method calls with dot notation (e.g., object.method(args))
		objectValue := Interpret(n.Object, env)

		// Get the object's class type
		if dict, ok := objectValue.(map[string]interface{}); ok {
			// Check if this is a class instance with a __class__ field
			if className, hasClass := dict["__class__"].(string); hasClass {
				// Find the method in the class or its parent chain
				method := findMethodInClass(className, n.Method, env)
				if method != nil {
					// Create new environment for method execution
					methodEnv := &Environment{
						Variables: make(map[string]interface{}),
						Functions: env.Functions,
						Classes:   env.Classes,
						Modules:   env.Modules,
						Parent:    env,
					}

					// Set 'hii' to refer to the current instance
					methodEnv.Set("hii", objectValue)

					// Bind parameters
					for i, param := range method.Parameters {
						if i < len(n.Args) {
							argValue := Interpret(n.Args[i], env)
							methodEnv.Set(param.Name, argValue)
						}
					}

					// Execute method body
					var result interface{}
					for _, stmt := range method.Body {
						result = Interpret(stmt, methodEnv)

						// Handle control flow
						if cf, ok := result.(ControlFlowResult); ok {
							if cf.Type == ControlReturn {
								return cf.Value
							}
							if cf.Type == ControlThrow {
								return cf
							}
						}
					}
					return result
				}

				// Method not found
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Mbinu '%s' haipatikani katika darasa '%s'", n.Method, className),
						Context: fmt.Sprintf("Method '%s' not found in class '%s'", n.Method, className),
					},
				}
			}
		}

		// If not a class instance, return error
		return ControlFlowResult{
			Type: ControlThrow,
			Value: ErrorValue{
				Message: "Haiwezi kuita mbinu kwenye kitu ambacho si instance ya darasa",
				Context: "Cannot call method on non-class instance",
			},
		}

	case ast.FunctionCallNode:
		// Handle built-in function calls
		if n.Name == "andika" {
			for i, arg := range n.Args {
				result := Interpret(arg, env)
				if i > 0 {
					fmt.Print(" ")
				}
				// Special formatting for dictionaries
				if dict, ok := result.(map[string]interface{}); ok {
					fmt.Print("{")
					first := true
					for key, value := range dict {
						if !first {
							fmt.Print(", ")
						}
						fmt.Printf("%q: %v", key, value)
						first = false
					}
					fmt.Print("}")
				} else if arr, ok := result.([]interface{}); ok {
					// Special formatting for arrays
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

		// Mathematical rounding functions
		if n.Name == "chini" && len(n.Args) == 1 {
			// Floor function: chini(number)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)
			result := math.Floor(num)

			// Return integer if input was integer, float if input was float
			if isFloat {
				return result
			}
			return int(result)
		}

		if n.Name == "chini" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "chini inahitaji hoja moja (chini requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'chini': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "juu" && len(n.Args) == 1 {
			// Ceiling function: juu(number)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)
			result := math.Ceil(num)

			// Return integer if input was integer, float if input was float
			if isFloat {
				return result
			}
			return int(result)
		}

		if n.Name == "juu" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "juu inahitaji hoja moja (juu requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'juu': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "zunguka" && len(n.Args) == 1 {
			// Round function: zunguka(number)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)
			result := math.Round(num)

			// Return integer if input was integer, float if input was float
			if isFloat {
				return result
			}
			return int(result)
		}

		if n.Name == "zunguka" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "zunguka inahitaji hoja moja (zunguka requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'zunguka': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "kata_desimali" && len(n.Args) == 1 {
			// Truncate function: kata_desimali(number)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)
			result := math.Trunc(num)

			// Return integer if input was integer, float if input was float
			if isFloat {
				return result
			}
			return int(result)
		}

		if n.Name == "kata_desimali" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "kata_desimali inahitaji hoja moja (kata_desimali requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'kata_desimali': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		// Mathematical root functions
		if n.Name == "mzizi_mraba" && len(n.Args) == 1 {
			// Square root function: mzizi_mraba(number)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			// Check for negative numbers
			if num < 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Haiwezekani kupata mzizi mraba wa namba hasi: %v", num),
						Context: "Katika kazi 'mzizi_mraba': Cannot calculate square root of negative number",
					},
				}
			}

			result := math.Sqrt(num)

			// Always return float for square roots to maintain precision
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "mzizi_mraba" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "mzizi_mraba inahitaji hoja moja (mzizi_mraba requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'mzizi_mraba': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "mzizi_mchemraba" && len(n.Args) == 1 {
			// Cube root function: mzizi_mchemraba(number)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.Cbrt(num)

			// Always return float for cube roots to maintain precision
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "mzizi_mchemraba" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "mzizi_mchemraba inahitaji hoja moja (mzizi_mchemraba requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'mzizi_mchemraba': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "mzizi" && len(n.Args) == 2 {
			// Nth root function: mzizi(number, n) calculates the nth root of number
			// Equivalent to: number^(1/n)
			argNum := Interpret(n.Args[0], env)
			argRoot := Interpret(n.Args[1], env)

			num, isFloatNum := toNumber(argNum)
			root, _ := toNumber(argRoot)

			// Check for invalid root
			if root == 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: "Haiwezekani kuhesabu mzizi wa nambari sifuri (Cannot calculate 0th root)",
						Context: "Katika kazi 'mzizi': Root index cannot be zero",
					},
				}
			}

			// Check for even root of negative number
			if num < 0 && int(root)%2 == 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Haiwezekani kupata mzizi wa %v wa namba hasi %v", int(root), num),
						Context: fmt.Sprintf("Katika kazi 'mzizi': Cannot calculate even root (%.0f) of negative number", root),
					},
				}
			}

			// Calculate nth root: x^(1/n)
			// For negative numbers with odd roots, handle specially
			var result float64
			if num < 0 && int(root)%2 != 0 {
				// Negative number with odd root: result is negative
				result = -math.Pow(-num, 1.0/root)
			} else {
				result = math.Pow(num, 1.0/root)
			}

			// Always return float for nth roots to maintain precision
			if isFloatNum || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "mzizi" && len(n.Args) != 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "mzizi inahitaji hoja mbili (mzizi requires two arguments: number, root)",
					Context: fmt.Sprintf("Katika kazi 'mzizi': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		// Exponential and logarithm functions
		if n.Name == "exp" && len(n.Args) == 1 {
			// Exponential function: exp(x) returns e^x
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.Exp(num)

			// Check for infinity (overflow)
			if math.IsInf(result, 1) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Jibu la exp(%v) ni kubwa sana (Result too large)", num),
						Context: "Katika kazi 'exp': Result would overflow (infinity)",
					},
				}
			}

			// Check for NaN
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Jibu la exp(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'exp': Result is NaN",
					},
				}
			}

			// Always return float for exponential unless it's a perfect integer
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "exp" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "exp inahitaji hoja moja (exp requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'exp': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "log_asili" && len(n.Args) == 1 {
			// Natural logarithm function: log_asili(x) returns ln(x)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			// Check for invalid input
			if num <= 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Haiwezekani kuhesabu logarithm ya %v (Must be positive)", num),
						Context: "Katika kazi 'log_asili': Logarithm only works with positive numbers",
					},
				}
			}

			result := math.Log(num)

			// Always return float for logarithm
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "log_asili" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "log_asili inahitaji hoja moja (log_asili requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'log_asili': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "log10" && len(n.Args) == 1 {
			// Base-10 logarithm function: log10(x)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			// Check for invalid input
			if num <= 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Haiwezekani kuhesabu logarithm ya %v (Must be positive)", num),
						Context: "Katika kazi 'log10': Logarithm only works with positive numbers",
					},
				}
			}

			result := math.Log10(num)

			// Always return float for logarithm
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "log10" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "log10 inahitaji hoja moja (log10 requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'log10': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "log2" && len(n.Args) == 1 {
			// Base-2 logarithm function: log2(x)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			// Check for invalid input
			if num <= 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Haiwezekani kuhesabu logarithm ya %v (Must be positive)", num),
						Context: "Katika kazi 'log2': Logarithm only works with positive numbers",
					},
				}
			}

			result := math.Log2(num)

			// Always return float for logarithm
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "log2" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "log2 inahitaji hoja moja (log2 requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'log2': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		// Trigonometric functions (angle in radians)
		if n.Name == "sin" && len(n.Args) == 1 {
			// Sine function: sin(x) where x is in radians
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.Sin(num)

			// Check for NaN
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Jibu la sin(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'sin': Result is NaN",
					},
				}
			}

			// Always return float for trig functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "sin" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "sin inahitaji hoja moja (sin requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'sin': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "cos" && len(n.Args) == 1 {
			// Cosine function: cos(x) where x is in radians
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.Cos(num)

			// Check for NaN
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Jibu la cos(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'cos': Result is NaN",
					},
				}
			}

			// Always return float for trig functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "cos" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "cos inahitaji hoja moja (cos requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'cos': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "tan" && len(n.Args) == 1 {
			// Tangent function: tan(x) where x is in radians
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.Tan(num)

			// Check for NaN
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Jibu la tan(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'tan': Result is NaN",
					},
				}
			}

			// Check for infinity (happens at π/2 + nπ)
			if math.IsInf(result, 0) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("tan(%v) haiwezekani (undefined at π/2 + nπ)", num),
						Context: "Katika kazi 'tan': Result is undefined (approaches infinity)",
					},
				}
			}

			// Always return float for trig functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "tan" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "tan inahitaji hoja moja (tan requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'tan': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		// Inverse trigonometric functions (return angle in radians)
		if n.Name == "asin" && len(n.Args) == 1 {
			// Arc sine function: asin(x) returns angle in radians, domain [-1, 1]
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			// Check domain: asin only works for -1 <= x <= 1
			if num < -1 || num > 1 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("asin(%v) haiwezekani (Domain must be between -1 and 1)", num),
						Context: "Katika kazi 'asin': Input must be in range [-1, 1]",
					},
				}
			}

			result := math.Asin(num)

			// Always return float for inverse trig functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "asin" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "asin inahitaji hoja moja (asin requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'asin': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "acos" && len(n.Args) == 1 {
			// Arc cosine function: acos(x) returns angle in radians, domain [-1, 1]
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			// Check domain: acos only works for -1 <= x <= 1
			if num < -1 || num > 1 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("acos(%v) haiwezekani (Domain must be between -1 and 1)", num),
						Context: "Katika kazi 'acos': Input must be in range [-1, 1]",
					},
				}
			}

			result := math.Acos(num)

			// Always return float for inverse trig functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "acos" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "acos inahitaji hoja moja (acos requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'acos': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "atan" && len(n.Args) == 1 {
			// Arc tangent function: atan(x) returns angle in radians, domain all real numbers
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.Atan(num)

			// Always return float for inverse trig functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "atan" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "atan inahitaji hoja moja (atan requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'atan': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "atan2" && len(n.Args) == 2 {
			// Two-argument arc tangent: atan2(y, x) returns angle in radians
			// Properly handles all quadrants and signs
			yArg := Interpret(n.Args[0], env)
			xArg := Interpret(n.Args[1], env)
			y, yIsFloat := toNumber(yArg)
			x, xIsFloat := toNumber(xArg)

			result := math.Atan2(y, x)

			// Always return float for inverse trig functions
			if yIsFloat || xIsFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "atan2" && len(n.Args) != 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "atan2 inahitaji hoja mbili (atan2 requires two arguments: y, x)",
					Context: fmt.Sprintf("Katika kazi 'atan2': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		// Degree/Radian conversion functions
		if n.Name == "radians" && len(n.Args) == 1 {
			// Convert degrees to radians: radians(degrees)
			arg := Interpret(n.Args[0], env)
			degrees, isFloat := toNumber(arg)

			// Formula: radians = degrees × π / 180
			result := degrees * math.Pi / 180.0

			// Always return float for conversion
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "radians" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "radians inahitaji hoja moja (radians requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'radians': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "degrees" && len(n.Args) == 1 {
			// Convert radians to degrees: degrees(radians)
			arg := Interpret(n.Args[0], env)
			radians, isFloat := toNumber(arg)

			// Formula: degrees = radians × 180 / π
			result := radians * 180.0 / math.Pi

			// Always return float for conversion
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "degrees" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "degrees inahitaji hoja moja (degrees requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'degrees': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		// Hyperbolic functions
		if n.Name == "sinh" && len(n.Args) == 1 {
			// Hyperbolic sine function: sinh(x) = (e^x - e^-x) / 2
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.Sinh(num)

			// Check for infinity (overflow)
			if math.IsInf(result, 0) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Jibu la sinh(%v) ni kubwa sana (Result too large)", num),
						Context: "Katika kazi 'sinh': Result would overflow (infinity)",
					},
				}
			}

			// Check for NaN
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Jibu la sinh(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'sinh': Result is NaN",
					},
				}
			}

			// Always return float for hyperbolic functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "sinh" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "sinh inahitaji hoja moja (sinh requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'sinh': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "cosh" && len(n.Args) == 1 {
			// Hyperbolic cosine function: cosh(x) = (e^x + e^-x) / 2
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.Cosh(num)

			// Check for infinity (overflow)
			if math.IsInf(result, 0) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Jibu la cosh(%v) ni kubwa sana (Result too large)", num),
						Context: "Katika kazi 'cosh': Result would overflow (infinity)",
					},
				}
			}

			// Check for NaN
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Jibu la cosh(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'cosh': Result is NaN",
					},
				}
			}

			// Always return float for hyperbolic functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "cosh" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "cosh inahitaji hoja moja (cosh requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'cosh': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "tanh" && len(n.Args) == 1 {
			// Hyperbolic tangent function: tanh(x) = sinh(x) / cosh(x)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.Tanh(num)

			// Check for NaN
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Jibu la tanh(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'tanh': Result is NaN",
					},
				}
			}

			// Always return float for hyperbolic functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "tanh" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "tanh inahitaji hoja moja (tanh requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'tanh': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		// Inverse hyperbolic functions
		if n.Name == "asinh" && len(n.Args) == 1 {
			// Inverse hyperbolic sine: asinh(x) = ln(x + sqrt(x^2 + 1))
			// Domain: all real numbers
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.Asinh(num)

			// Check for NaN
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Jibu la asinh(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'asinh': Result is NaN",
					},
				}
			}

			// Always return float for inverse hyperbolic functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "asinh" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "asinh inahitaji hoja moja (asinh requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'asinh': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "acosh" && len(n.Args) == 1 {
			// Inverse hyperbolic cosine: acosh(x) = ln(x + sqrt(x^2 - 1))
			// Domain: x >= 1
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			// Check domain: acosh only works for x >= 1
			if num < 1 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("acosh(%v) haiwezekani (Domain must be >= 1)", num),
						Context: "Katika kazi 'acosh': Input must be greater than or equal to 1",
					},
				}
			}

			result := math.Acosh(num)

			// Check for NaN
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Jibu la acosh(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'acosh': Result is NaN",
					},
				}
			}

			// Always return float for inverse hyperbolic functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "acosh" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "acosh inahitaji hoja moja (acosh requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'acosh': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "atanh" && len(n.Args) == 1 {
			// Inverse hyperbolic tangent: atanh(x) = 0.5 * ln((1+x)/(1-x))
			// Domain: -1 < x < 1 (strictly between -1 and 1)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			// Check domain: atanh only works for -1 < x < 1
			if num <= -1 || num >= 1 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("atanh(%v) haiwezekani (Domain must be strictly between -1 and 1)", num),
						Context: "Katika kazi 'atanh': Input must be in range (-1, 1)",
					},
				}
			}

			result := math.Atanh(num)

			// Check for infinity
			if math.IsInf(result, 0) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Jibu la atanh(%v) ni kubwa sana (Result approaches infinity)", num),
						Context: "Katika kazi 'atanh': Result would overflow (approaches infinity at boundaries)",
					},
				}
			}

			// Check for NaN
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Jibu la atanh(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'atanh': Result is NaN",
					},
				}
			}

			// Always return float for inverse hyperbolic functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "atanh" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "atanh inahitaji hoja moja (atanh requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'atanh': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		// Number Property Functions
		if n.Name == "ni_shufwa" && len(n.Args) == 1 {
			// Check if number is even (ni_shufwa = is_even)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)

			// Convert to integer for modulo operation
			intNum := int64(num)
			return intNum%2 == 0
		}

		if n.Name == "ni_shufwa" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "ni_shufwa inahitaji hoja moja (ni_shufwa requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'ni_shufwa': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "ni_witiri" && len(n.Args) == 1 {
			// Check if number is odd (ni_witiri = is_odd)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)

			// Convert to integer for modulo operation
			intNum := int64(num)
			return intNum%2 != 0
		}

		if n.Name == "ni_witiri" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "ni_witiri inahitaji hoja moja (ni_witiri requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'ni_witiri': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "ni_chanya" && len(n.Args) == 1 {
			// Check if number is positive (ni_chanya = is_positive)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)
			return num > 0
		}

		if n.Name == "ni_chanya" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "ni_chanya inahitaji hoja moja (ni_chanya requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'ni_chanya': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "ni_hasi" && len(n.Args) == 1 {
			// Check if number is negative (ni_hasi = is_negative)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)
			return num < 0
		}

		if n.Name == "ni_hasi" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "ni_hasi inahitaji hoja moja (ni_hasi requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'ni_hasi': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "ni_sifuri" && len(n.Args) == 1 {
			// Check if number is zero (ni_sifuri = is_zero)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)
			return num == 0
		}

		if n.Name == "ni_sifuri" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "ni_sifuri inahitaji hoja moja (ni_sifuri requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'ni_sifuri': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "ni_namba_kuu" && len(n.Args) == 1 {
			// Check if number is prime (ni_namba_kuu = is_prime)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)

			// Convert to integer
			intNum := int64(num)

			// Numbers less than 2 are not prime
			if intNum < 2 {
				return false
			}

			// 2 is prime
			if intNum == 2 {
				return true
			}

			// Even numbers are not prime
			if intNum%2 == 0 {
				return false
			}

			// Check for odd divisors up to sqrt(intNum)
			for i := int64(3); i*i <= intNum; i += 2 {
				if intNum%i == 0 {
					return false
				}
			}

			return true
		}

		if n.Name == "ni_namba_kuu" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "ni_namba_kuu inahitaji hoja moja (ni_namba_kuu requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'ni_namba_kuu': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "ni_mraba_kamili" && len(n.Args) == 1 {
			// Check if number is a perfect square (ni_mraba_kamili = is_perfect_square)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)

			// Only non-negative numbers can be perfect squares
			if num < 0 {
				return false
			}

			// Get the square root
			sqrt := math.Sqrt(num)

			// Check if square root is an integer
			return sqrt == math.Floor(sqrt)
		}

		if n.Name == "ni_mraba_kamili" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "ni_mraba_kamili inahitaji hoja moja (ni_mraba_kamili requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'ni_mraba_kamili': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		// Utility Functions
		if n.Name == "kiwango" && len(n.Args) == 1 {
			// Absolute value (kiwango = absolute value)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.Abs(num)

			// Return int if input was int and result is perfect integer
			if !isFloat && result == math.Floor(result) {
				return int(result)
			}
			return result
		}

		if n.Name == "kiwango" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "kiwango inahitaji hoja moja (kiwango requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'kiwango': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "ishara" && len(n.Args) == 1 {
			// Sign function (ishara = sign) - returns -1, 0, or 1
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)

			if num > 0 {
				return 1
			} else if num < 0 {
				return -1
			}
			return 0
		}

		if n.Name == "ishara" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "ishara inahitaji hoja moja (ishara requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'ishara': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "max" && len(n.Args) >= 1 {
			// Maximum value - can take multiple arguments
			if len(n.Args) == 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: "max inahitaji angalau hoja moja (max requires at least one argument)",
						Context: "Katika kazi 'max': Angalau hoja moja inahitajika",
					},
				}
			}

			maxVal, isFloat := toNumber(Interpret(n.Args[0], env))

			for i := 1; i < len(n.Args); i++ {
				val, valIsFloat := toNumber(Interpret(n.Args[i], env))
				if val > maxVal {
					maxVal = val
					isFloat = isFloat || valIsFloat
				}
			}

			// Return int if all inputs were ints
			if !isFloat && maxVal == math.Floor(maxVal) {
				return int(maxVal)
			}
			return maxVal
		}

		if n.Name == "max" && len(n.Args) == 0 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "max inahitaji angalau hoja moja (max requires at least one argument)",
					Context: "Katika kazi 'max': Angalau hoja moja inahitajika",
				},
			}
		}

		if n.Name == "min" && len(n.Args) >= 1 {
			// Minimum value - can take multiple arguments
			if len(n.Args) == 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: "min inahitaji angalau hoja moja (min requires at least one argument)",
						Context: "Katika kazi 'min': Angalau hoja moja inahitajika",
					},
				}
			}

			minVal, isFloat := toNumber(Interpret(n.Args[0], env))

			for i := 1; i < len(n.Args); i++ {
				val, valIsFloat := toNumber(Interpret(n.Args[i], env))
				if val < minVal {
					minVal = val
					isFloat = isFloat || valIsFloat
				}
			}

			// Return int if all inputs were ints
			if !isFloat && minVal == math.Floor(minVal) {
				return int(minVal)
			}
			return minVal
		}

		if n.Name == "min" && len(n.Args) == 0 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "min inahitaji angalau hoja moja (min requires at least one argument)",
					Context: "Katika kazi 'min': Angalau hoja moja inahitajika",
				},
			}
		}

		// Mathematical Utility Functions
		if n.Name == "kigawanyaji_kikuu" && len(n.Args) == 2 {
			// Greatest Common Divisor (kigawanyaji_kikuu = gcd)
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)

			a, _ := toNumber(arg1)
			b, _ := toNumber(arg2)

			// Convert to positive integers
			a = math.Abs(a)
			b = math.Abs(b)
			aInt := int64(a)
			bInt := int64(b)

			// Euclidean algorithm
			for bInt != 0 {
				temp := bInt
				bInt = aInt % bInt
				aInt = temp
			}

			return aInt
		}

		if n.Name == "kigawanyaji_kikuu" && len(n.Args) != 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "kigawanyaji_kikuu inahitaji hoja mbili (kigawanyaji_kikuu requires two arguments)",
					Context: fmt.Sprintf("Katika kazi 'kigawanyaji_kikuu': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "kigawanyaji_ndogo" && len(n.Args) == 2 {
			// Least Common Multiple (kigawanyaji_ndogo = lcm)
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)

			a, _ := toNumber(arg1)
			b, _ := toNumber(arg2)

			// Convert to positive integers
			a = math.Abs(a)
			b = math.Abs(b)
			aInt := int64(a)
			bInt := int64(b)

			// LCM = (a * b) / GCD(a, b)
			// First calculate GCD
			origA := aInt
			origB := bInt
			for bInt != 0 {
				temp := bInt
				bInt = aInt % bInt
				aInt = temp
			}
			gcd := aInt

			// Avoid division by zero
			if gcd == 0 {
				return 0
			}

			lcm := (origA * origB) / gcd
			return lcm
		}

		if n.Name == "kigawanyaji_ndogo" && len(n.Args) != 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "kigawanyaji_ndogo inahitaji hoja mbili (kigawanyaji_ndogo requires two arguments)",
					Context: fmt.Sprintf("Katika kazi 'kigawanyaji_ndogo': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "ukweli" && len(n.Args) == 1 {
			// Factorial (ukweli = factorial)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)

			// Convert to integer
			n := int64(num)

			// Factorial is only defined for non-negative integers
			if n < 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("ukweli(%v) haiwezekani (Factorial not defined for negative numbers)", n),
						Context: "Katika kazi 'ukweli': Factorial requires non-negative integer",
					},
				}
			}

			// Calculate factorial
			result := int64(1)
			for i := int64(2); i <= n; i++ {
				result *= i

				// Check for overflow
				if result < 0 {
					return ControlFlowResult{
						Type: ControlThrow,
						Value: ErrorValue{
							Message: fmt.Sprintf("ukweli(%v) ni kubwa sana (Factorial too large)", num),
							Context: "Katika kazi 'ukweli': Result overflows",
						},
					}
				}
			}

			return result
		}

		if n.Name == "ukweli" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "ukweli inahitaji hoja moja (ukweli requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'ukweli': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "tweza" && len(n.Args) == 2 {
			// Power function (tweza = power)
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)

			base, baseIsFloat := toNumber(arg1)
			exponent, expIsFloat := toNumber(arg2)

			result := math.Pow(base, exponent)

			// Check for special cases
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("tweza(%v, %v) si sahihi (Result is not a valid number)", base, exponent),
						Context: "Katika kazi 'tweza': Result is NaN",
					},
				}
			}

			if math.IsInf(result, 0) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("tweza(%v, %v) ni kubwa sana (Result is too large)", base, exponent),
						Context: "Katika kazi 'tweza': Result overflows",
					},
				}
			}

			// Return int if both inputs were ints and result is perfect integer
			if !baseIsFloat && !expIsFloat && result == math.Floor(result) && result >= -9223372036854775808 && result <= 9223372036854775807 {
				return int(result)
			}
			return result
		}

		if n.Name == "tweza" && len(n.Args) != 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "tweza inahitaji hoja mbili (tweza requires two arguments: base and exponent)",
					Context: fmt.Sprintf("Katika kazi 'tweza': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		// ==================== BITWISE OPERATIONS ====================

		if n.Name == "na_kidogo" && len(n.Args) == 2 {
			// Bitwise AND (na_kidogo = bitwise AND)
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)

			a, _ := toNumber(arg1)
			b, _ := toNumber(arg2)

			aInt := int64(a)
			bInt := int64(b)

			return aInt & bInt
		}

		if n.Name == "na_kidogo" && len(n.Args) != 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "na_kidogo inahitaji hoja mbili (na_kidogo requires two arguments)",
					Context: fmt.Sprintf("Katika kazi 'na_kidogo': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "au_kidogo" && len(n.Args) == 2 {
			// Bitwise OR (au_kidogo = bitwise OR)
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)

			a, _ := toNumber(arg1)
			b, _ := toNumber(arg2)

			aInt := int64(a)
			bInt := int64(b)

			return aInt | bInt
		}

		if n.Name == "au_kidogo" && len(n.Args) != 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "au_kidogo inahitaji hoja mbili (au_kidogo requires two arguments)",
					Context: fmt.Sprintf("Katika kazi 'au_kidogo': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "ama_kidogo" && len(n.Args) == 2 {
			// Bitwise XOR (ama_kidogo = bitwise XOR)
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)

			a, _ := toNumber(arg1)
			b, _ := toNumber(arg2)

			aInt := int64(a)
			bInt := int64(b)

			return aInt ^ bInt
		}

		if n.Name == "ama_kidogo" && len(n.Args) != 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "ama_kidogo inahitaji hoja mbili (ama_kidogo requires two arguments)",
					Context: fmt.Sprintf("Katika kazi 'ama_kidogo': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "si_kidogo" && len(n.Args) == 1 {
			// Bitwise NOT (si_kidogo = bitwise NOT)
			arg := Interpret(n.Args[0], env)

			a, _ := toNumber(arg)
			aInt := int64(a)

			return ^aInt
		}

		if n.Name == "si_kidogo" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "si_kidogo inahitaji hoja moja (si_kidogo requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'si_kidogo': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "gezo_kushoto" && len(n.Args) == 2 {
			// Left shift (gezo_kushoto = left shift)
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)

			a, _ := toNumber(arg1)
			shift, _ := toNumber(arg2)

			aInt := int64(a)
			shiftInt := uint(shift)

			return aInt << shiftInt
		}

		if n.Name == "gezo_kushoto" && len(n.Args) != 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "gezo_kushoto inahitaji hoja mbili (gezo_kushoto requires two arguments)",
					Context: fmt.Sprintf("Katika kazi 'gezo_kushoto': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "gezo_kulia" && len(n.Args) == 2 {
			// Right shift (gezo_kulia = right shift)
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)

			a, _ := toNumber(arg1)
			shift, _ := toNumber(arg2)

			aInt := int64(a)
			shiftInt := uint(shift)

			return aInt >> shiftInt
		}

		if n.Name == "gezo_kulia" && len(n.Args) != 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "gezo_kulia inahitaji hoja mbili (gezo_kulia requires two arguments)",
					Context: fmt.Sprintf("Katika kazi 'gezo_kulia': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		// ==================== MODULAR ARITHMETIC ====================

		if n.Name == "modulo" && len(n.Args) == 2 {
			// Modulo operation (modulo = modulo remainder)
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)

			a, aIsFloat := toNumber(arg1)
			b, bIsFloat := toNumber(arg2)

			if b == 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Imkishaji na sifuri (Division by zero) katika modulo(%v, %v)", a, b),
						Context: "Katika kazi 'modulo': Denominator cannot be zero",
					},
				}
			}

			result := math.Mod(a, b)

			// Return int if both inputs were ints
			if !aIsFloat && !bIsFloat && result == math.Floor(result) {
				return int(result)
			}
			return result
		}

		if n.Name == "modulo" && len(n.Args) != 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "modulo inahitaji hoja mbili (modulo requires two arguments)",
					Context: fmt.Sprintf("Katika kazi 'modulo': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "baki" && len(n.Args) == 2 {
			// Remainder operation (baki = remainder)
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)

			a, _ := toNumber(arg1)
			b, _ := toNumber(arg2)

			if b == 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Imkishaji na sifuri (Division by zero) katika baki(%v, %v)", a, b),
						Context: "Katika kazi 'baki': Denominator cannot be zero",
					},
				}
			}

			// For integers, use modulo directly
			aInt := int64(a)
			bInt := int64(b)

			result := aInt % bInt
			return result
		}

		if n.Name == "baki" && len(n.Args) != 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "baki inahitaji hoja mbili (baki requires two arguments)",
					Context: fmt.Sprintf("Katika kazi 'baki': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "tweza_modulo" && len(n.Args) == 3 {
			// Modular exponentiation (tweza_modulo = power mod)
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)
			arg3 := Interpret(n.Args[2], env)

			baseNum, _ := toNumber(arg1)
			expNum, _ := toNumber(arg2)
			modNum, _ := toNumber(arg3)

			base := int64(baseNum)
			exponent := int64(expNum)
			modulus := int64(modNum)

			if modulus == 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("Imkishaji na sifuri (Division by zero) katika tweza_modulo(%v, %v, %v)", base, exponent, modulus),
						Context: "Katika kazi 'tweza_modulo': Modulus cannot be zero",
					},
				}
			}

			// Handle negative exponent
			if exponent < 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("tweza_modulo(%v, %v, %v) haiwezekani (Negative exponent not supported)", base, exponent, modulus),
						Context: "Katika kazi 'tweza_modulo': Exponent must be non-negative",
					},
				}
			}

			// Simple modular exponentiation
			result := int64(1)
			base = ((base % modulus) + modulus) % modulus

			for exponent > 0 {
				if exponent%2 == 1 {
					result = (result * base) % modulus
				}
				exponent = exponent >> 1
				base = (base * base) % modulus
			}

			return result
		}

		if n.Name == "tweza_modulo" && len(n.Args) != 3 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "tweza_modulo inahitaji hoja tatu (tweza_modulo requires three arguments: base, exponent, modulus)",
					Context: fmt.Sprintf("Katika kazi 'tweza_modulo': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		// ==================== ROUNDING FUNCTIONS ====================

		if n.Name == "pindika" && len(n.Args) == 1 {
			// Round (pindika = round to nearest integer)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)

			result := math.Round(num)
			return int(result)
		}

		if n.Name == "pindika" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "pindika inahitaji hoja moja (pindika requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'pindika': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "sakafu" && len(n.Args) == 1 {
			// Floor (sakafu = floor - round down)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)

			result := math.Floor(num)
			return int(result)
		}

		if n.Name == "sakafu" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "sakafu inahitaji hoja moja (sakafu requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'sakafu': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "dari" && len(n.Args) == 1 {
			// Ceiling (dari = ceil - round up)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)

			result := math.Ceil(num)
			return int(result)
		}

		if n.Name == "dari" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "dari inahitaji hoja moja (dari requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'dari': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "kata" && len(n.Args) == 1 {
			// Truncate (kata = truncate - remove decimal part)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)

			result := math.Trunc(num)
			return int(result)
		}

		if n.Name == "kata" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "kata inahitaji hoja moja (kata requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'kata': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		// ==================== FLOAT UTILITIES ====================

		if n.Name == "bakiza" && len(n.Args) == 2 {
			// Modulo/Remainder (bakiza = remainder/modulo for floating point)
			// Returns the IEEE 754 floating-point remainder of x/y
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)
			
			x, xIsFloat := toNumber(arg1)
			y, yIsFloat := toNumber(arg2)
			
			if y == 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("bakiza(%v, %v): Imegawanya kwa sifuri (Division by zero)", x, y),
						Context: "Katika kazi 'bakiza': Cannot compute remainder with divisor of zero",
					},
				}
			}
			
			result := math.Mod(x, y)
			
			// Return int if both inputs were ints and result is a whole number
			if !xIsFloat && !yIsFloat && result == math.Floor(result) {
				return int(result)
			}
			return result
		}

		if n.Name == "bakiza" && len(n.Args) != 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "bakiza inahitaji hoja mbili (bakiza requires two arguments: dividend, divisor)",
					Context: fmt.Sprintf("Katika kazi 'bakiza': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "salio" && len(n.Args) == 2 {
			// Remainder (salio = remainder - integer remainder)
			// Returns the remainder of integer division
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)
			
			x, _ := toNumber(arg1)
			y, _ := toNumber(arg2)
			
			if y == 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("salio(%v, %v): Imegawanya kwa sifuri (Division by zero)", x, y),
						Context: "Katika kazi 'salio': Cannot compute remainder with divisor of zero",
					},
				}
			}
			
			// Use math.Remainder for IEEE 754 remainder
			result := math.Remainder(x, y)
			
			return result
		}

		if n.Name == "salio" && len(n.Args) != 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "salio inahitaji hoja mbili (salio requires two arguments: dividend, divisor)",
					Context: fmt.Sprintf("Katika kazi 'salio': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "kopanja" && len(n.Args) == 2 {
			// Copysign (kopanja = copysign)
			// Returns a value with the magnitude of x and the sign of y
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)
			
			x, xIsFloat := toNumber(arg1)
			y, _ := toNumber(arg2)
			
			result := math.Copysign(x, y)
			
			// Return int if x was int and result is a whole number
			if !xIsFloat && result == math.Floor(result) {
				return int(result)
			}
			return result
		}

		if n.Name == "kopanja" && len(n.Args) != 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "kopanja inahitaji hoja mbili (kopanja requires two arguments: magnitude, sign)",
					Context: fmt.Sprintf("Katika kazi 'kopanja': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "karibia" && len(n.Args) == 2 {
			// Check if two floats are approximately equal (karibia = approximately equal)
			// Uses a small epsilon for comparison
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)
			
			x, _ := toNumber(arg1)
			y, _ := toNumber(arg2)
			
			epsilon := 1e-9
			return math.Abs(x-y) < epsilon
		}

		if n.Name == "karibia" && len(n.Args) == 3 {
			// Check if two floats are approximately equal with custom epsilon
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)
			arg3 := Interpret(n.Args[2], env)
			
			x, _ := toNumber(arg1)
			y, _ := toNumber(arg2)
			epsilon, _ := toNumber(arg3)
			
			if epsilon < 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("karibia: epsilon lazima iwe chanya (epsilon must be positive), imepatikana %v", epsilon),
						Context: "Katika kazi 'karibia': Epsilon value must be positive",
					},
				}
			}
			
			return math.Abs(x-y) < epsilon
		}

		if n.Name == "karibia" && len(n.Args) != 2 && len(n.Args) != 3 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "karibia inahitaji hoja 2 au 3 (karibia requires 2 or 3 arguments: x, y, [epsilon])",
					Context: fmt.Sprintf("Katika kazi 'karibia': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "ni_kamili" && len(n.Args) == 1 {
			// Check if number is a whole number (ni_kamili = is whole/integer)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)
			
			return num == math.Floor(num)
		}

		if n.Name == "ni_kamili" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "ni_kamili inahitaji hoja moja (ni_kamili requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'ni_kamili': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "ni_usawa" && len(n.Args) == 1 {
			// Check if number is finite (ni_usawa = is finite)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)
			
			return !math.IsInf(num, 0) && !math.IsNaN(num)
		}

		if n.Name == "ni_usawa" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "ni_usawa inahitaji hoja moja (ni_usawa requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'ni_usawa': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "ni_bila_kikomo" && len(n.Args) == 1 {
			// Check if number is infinite (ni_bila_kikomo = is infinite)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)
			
			return math.IsInf(num, 0)
		}

		if n.Name == "ni_bila_kikomo" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "ni_bila_kikomo inahitaji hoja moja (ni_bila_kikomo requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'ni_bila_kikomo': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "ni_sio_namba" && len(n.Args) == 1 {
			// Check if value is NaN (ni_sio_namba = is not a number)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)
			
			return math.IsNaN(num)
		}

		if n.Name == "ni_sio_namba" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "ni_sio_namba inahitaji hoja moja (ni_sio_namba requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'ni_sio_namba': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "upeo" && len(n.Args) == 3 {
			// Clamp a value between min and max (upeo = limit/clamp)
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)
			arg3 := Interpret(n.Args[2], env)
			
			value, valIsFloat := toNumber(arg1)
			minVal, minIsFloat := toNumber(arg2)
			maxVal, maxIsFloat := toNumber(arg3)
			
			if minVal > maxVal {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("upeo(%v, %v, %v): min lazima iwe ndogo au sawa na max", value, minVal, maxVal),
						Context: "Katika kazi 'upeo': min must be less than or equal to max",
					},
				}
			}
			
			result := value
			if result < minVal {
				result = minVal
			}
			if result > maxVal {
				result = maxVal
			}
			
			// Return int if all inputs were ints
			if !valIsFloat && !minIsFloat && !maxIsFloat && result == math.Floor(result) {
				return int(result)
			}
			return result
		}

		if n.Name == "upeo" && len(n.Args) != 3 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "upeo inahitaji hoja tatu (upeo requires three arguments: value, min, max)",
					Context: fmt.Sprintf("Katika kazi 'upeo': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "sehemu_desimali" && len(n.Args) == 1 {
			// Get the fractional part of a number (sehemu_desimali = decimal part)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)
			
			_, frac := math.Modf(num)
			return frac
		}

		if n.Name == "sehemu_desimali" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "sehemu_desimali inahitaji hoja moja (sehemu_desimali requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'sehemu_desimali': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "sehemu_kamili" && len(n.Args) == 1 {
			// Get the integer part of a number (sehemu_kamili = integer part)
			arg := Interpret(n.Args[0], env)
			num, _ := toNumber(arg)
			
			intPart, _ := math.Modf(num)
			return int(intPart)
		}

		if n.Name == "sehemu_kamili" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "sehemu_kamili inahitaji hoja moja (sehemu_kamili requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'sehemu_kamili': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "mzunguko" && len(n.Args) == 2 {
			// Linear interpolation (mzunguko = interpolate/blend)
			// Returns a value between start and end based on t (0 to 1)
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)
			
			start, _ := toNumber(arg1)
			end, _ := toNumber(arg2)
			
			return (start + end) / 2.0
		}

		if n.Name == "mzunguko" && len(n.Args) == 3 {
			// Linear interpolation with parameter t
			arg1 := Interpret(n.Args[0], env)
			arg2 := Interpret(n.Args[1], env)
			arg3 := Interpret(n.Args[2], env)
			
			start, _ := toNumber(arg1)
			end, _ := toNumber(arg2)
			t, _ := toNumber(arg3)
			
			result := start + t*(end-start)
			return result
		}

		if n.Name == "mzunguko" && len(n.Args) != 2 && len(n.Args) != 3 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "mzunguko inahitaji hoja 2 au 3 (mzunguko requires 2 or 3 arguments: start, end, [t])",
					Context: fmt.Sprintf("Katika kazi 'mzunguko': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		// ==================== RANDOM NUMBER GENERATION ====================

		if n.Name == "nasibu" && len(n.Args) == 0 {
			// Random float between 0.0 and 1.0 (nasibu = random)
			return rand.Float64()
		}

		if n.Name == "nasibu" && len(n.Args) == 1 {
			// Random integer between 0 and n-1
			maxArg := Interpret(n.Args[0], env)
			maxNum, _ := toNumber(maxArg)
			maxInt := int64(maxNum)

			if maxInt <= 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("nasibu(%v) inahitaji namba chanya (nasibu requires positive number)", maxInt),
						Context: "Katika kazi 'nasibu': Maximum must be greater than zero",
					},
				}
			}

			return rand.Int63n(maxInt)
		}

		if n.Name == "nasibu" && len(n.Args) == 2 {
			// Random integer between min and max (inclusive on both ends)
			minArg := Interpret(n.Args[0], env)
			maxArg := Interpret(n.Args[1], env)

			minNum, _ := toNumber(minArg)
			maxNum, _ := toNumber(maxArg)

			minInt := int64(minNum)
			maxInt := int64(maxNum)

			if minInt >= maxInt {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("nasibu(%v, %v): min lazima iwe ndogo kuliko max (min must be less than max)", minInt, maxInt),
						Context: "Katika kazi 'nasibu': Invalid range",
					},
				}
			}

			return minInt + rand.Int63n(maxInt-minInt+1)
		}

		if n.Name == "nasibu" && len(n.Args) > 2 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "nasibu inahitaji 0, 1, au 2 hoja (nasibu requires 0, 1, or 2 arguments)",
					Context: fmt.Sprintf("Katika kazi 'nasibu': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "weka_mbegu" && len(n.Args) == 1 {
			// Set random seed (weka_mbegu = set seed)
			seedArg := Interpret(n.Args[0], env)
			seedNum, _ := toNumber(seedArg)
			seedInt := int64(seedNum)

			rand.Seed(seedInt)
			return nil
		}

		if n.Name == "weka_mbegu" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "weka_mbegu inahitaji hoja moja (weka_mbegu requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'weka_mbegu': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		// ==================== ADVANCED SPECIAL FUNCTIONS ====================

		if n.Name == "gamma" && len(n.Args) == 1 {
			// Gamma function (gamma = gamma function)
			// Gamma(n) = (n-1)! for positive integers
			// Extends factorial to real and complex numbers
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.Gamma(num)

			// Check for special cases
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("gamma(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'gamma': Result is NaN",
					},
				}
			}

			if math.IsInf(result, 0) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("gamma(%v) ni kubwa sana (Result is infinite)", num),
						Context: "Katika kazi 'gamma': Result is infinite at poles",
					},
				}
			}

			// Always return float for gamma function
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "gamma" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "gamma inahitaji hoja moja (gamma requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'gamma': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "erf" && len(n.Args) == 1 {
			// Error function (erf = error function)
			// Used in statistics and probability theory
			// erf(x) is the probability integral of a normal distribution
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.Erf(num)

			// Check for special cases
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("erf(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'erf': Result is NaN",
					},
				}
			}

			// Always return float for erf function
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "erf" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "erf inahitaji hoja moja (erf requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'erf': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "erfc" && len(n.Args) == 1 {
			// Complementary error function (erfc = complementary error function)
			// erfc(x) = 1 - erf(x)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.Erfc(num)

			// Check for special cases
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("erfc(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'erfc': Result is NaN",
					},
				}
			}

			// Always return float for erfc function
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "erfc" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "erfc inahitaji hoja moja (erfc requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'erfc': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "j0" && len(n.Args) == 1 {
			// Bessel function of the first kind order 0 (j0 = Bessel J0)
			// Used in physics and engineering (waves, oscillations, etc.)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.J0(num)

			// Check for special cases
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("j0(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'j0': Result is NaN",
					},
				}
			}

			// Always return float for Bessel functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "j0" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "j0 inahitaji hoja moja (j0 requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'j0': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "j1" && len(n.Args) == 1 {
			// Bessel function of the first kind order 1 (j1 = Bessel J1)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			result := math.J1(num)

			// Check for special cases
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("j1(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'j1': Result is NaN",
					},
				}
			}

			// Always return float for Bessel functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "j1" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "j1 inahitaji hoja moja (j1 requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'j1': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "y0" && len(n.Args) == 1 {
			// Bessel function of the second kind order 0 (y0 = Bessel Y0)
			// Y0 is also called Neumann function N0
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			// Y0 is undefined for non-positive numbers
			if num <= 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("y0(%v) haiwezekani (Y0 only defined for positive numbers)", num),
						Context: "Katika kazi 'y0': Bessel Y functions require positive arguments",
					},
				}
			}

			result := math.Y0(num)

			// Check for special cases
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("y0(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'y0': Result is NaN",
					},
				}
			}

			// Always return float for Bessel functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "y0" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "y0 inahitaji hoja moja (y0 requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'y0': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
		}

		if n.Name == "y1" && len(n.Args) == 1 {
			// Bessel function of the second kind order 1 (y1 = Bessel Y1)
			arg := Interpret(n.Args[0], env)
			num, isFloat := toNumber(arg)

			// Y1 is undefined for non-positive numbers
			if num <= 0 {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("y1(%v) haiwezekani (Y1 only defined for positive numbers)", num),
						Context: "Katika kazi 'y1': Bessel Y functions require positive arguments",
					},
				}
			}

			result := math.Y1(num)

			// Check for special cases
			if math.IsNaN(result) {
				return ControlFlowResult{
					Type: ControlThrow,
					Value: ErrorValue{
						Message: fmt.Sprintf("y1(%v) si sahihi (Result is not a valid number)", num),
						Context: "Katika kazi 'y1': Result is NaN",
					},
				}
			}

			// Always return float for Bessel functions
			if isFloat || result != math.Floor(result) {
				return result
			}
			return int(result)
		}

		if n.Name == "y1" && len(n.Args) != 1 {
			return ControlFlowResult{
				Type: ControlThrow,
				Value: ErrorValue{
					Message: "y1 inahitaji hoja moja (y1 requires one argument)",
					Context: fmt.Sprintf("Katika kazi 'y1': Hoja %d zilizotolewa", len(n.Args)),
				},
			}
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

		// Check if it's a lambda stored in a variable
		if lambdaValue := env.Get(n.Name); lambdaValue != nil {
			if lambda, ok := lambdaValue.(map[string]interface{}); ok {
				if lambdaType, hasType := lambda["__type__"].(string); hasType && lambdaType == "lambda" {
					// Call the lambda
					parameters := lambda["__parameters__"].([]ast.Parameter)
					body := lambda["__body__"].([]ast.ASTNode)
					closureEnv := lambda["__env__"].(*Environment)

					// Create new environment for lambda execution (with closure)
					lambdaEnv := NewChildEnvironment(closureEnv)

					// Evaluate arguments and bind to parameters
					for i, param := range parameters {
						if i < len(n.Args) {
							argValue := Interpret(n.Args[i], env)
							lambdaEnv.Set(param.Name, argValue)
						}
					}

					// Execute lambda body
					var result interface{}
					for _, statement := range body {
						result = Interpret(statement, lambdaEnv)

						// Check for return statement or throw
						if cf, ok := result.(ControlFlowResult); ok {
							if cf.Type == ControlReturn {
								return cf.Value
							} else if cf.Type == ControlThrow {
								return cf
							}
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

	case ast.ClassNode:
		// Handle class definitions
		// Store class definition in environment
		env.SetClass(n.Name, n)
		return nil

	case ast.NewInstanceNode:
		// Handle class instantiation (unda ClassName(args))
		classDef, exists := env.GetClass(n.ClassName)
		if !exists {
			return ControlFlowResult{Type: ControlThrow, Value: ErrorValue{Message: fmt.Sprintf("Darasa '%s' halijulikani (Class '%s' not found)", n.ClassName, n.ClassName)}}
		}

		// Create new instance as a dictionary
		instance := make(map[string]interface{})

		// Collect properties from inheritance chain (parent first, then child)
		allProperties := collectInheritedProperties(classDef, env)

		// Initialize properties with default values
		for _, prop := range allProperties {
			instance[prop.Name] = nil
		}

		// Call constructor if it exists
		if classDef.Constructor != nil {
			// Create environment for constructor
			constructorEnv := NewChildEnvironment(env)

			// Set 'hii' to refer to the instance
			constructorEnv.Set("hii", instance)

			// Bind constructor parameters
			for i, param := range classDef.Constructor.Parameters {
				if i < len(n.Args) {
					argValue := Interpret(n.Args[i], env)
					constructorEnv.Set(param.Name, argValue)
				}
			}

			// Execute constructor body
			for _, statement := range classDef.Constructor.Body {
				result := Interpret(statement, constructorEnv)
				if cf, ok := result.(ControlFlowResult); ok {
					if cf.Type == ControlThrow {
						return cf
					}
				}
			}
		}

		// Store class name in instance for method calls
		instance["__class__"] = n.ClassName

		return instance

	case ast.LambdaNode:
		// Handle lambda functions - return the lambda as a callable value
		// Store the lambda with its closure environment
		lambdaValue := map[string]interface{}{
			"__type__":        "lambda",
			"__parameters__":  n.Parameters,
			"__return_type__": n.ReturnType,
			"__body__":        n.Body,
			"__env__":         env, // Capture closure
		}
		return lambdaValue

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

// collectInheritedProperties collects all properties from the class and its parent chain
func collectInheritedProperties(class ast.ClassNode, env *Environment) []ast.PropertyNode {
	var properties []ast.PropertyNode

	// First, collect parent properties if there's a parent
	if class.Parent != "" {
		parentClass, exists := env.GetClass(class.Parent)
		if exists {
			properties = append(properties, collectInheritedProperties(parentClass, env)...)
		}
	}

	// Then add this class's properties
	properties = append(properties, class.Properties...)

	return properties
}

// collectInheritedMethods collects all methods from the class and its parent chain
func collectInheritedMethods(class ast.ClassNode, env *Environment) []ast.FunctionNode {
	methodMap := make(map[string]ast.FunctionNode)

	// First, collect parent methods if there's a parent
	if class.Parent != "" {
		parentClass, exists := env.GetClass(class.Parent)
		if exists {
			parentMethods := collectInheritedMethods(parentClass, env)
			for _, method := range parentMethods {
				methodMap[method.Name] = method
			}
		}
	}

	// Then add/override with this class's methods
	for _, method := range class.Methods {
		methodMap[method.Name] = method
	}

	// Convert map back to slice
	var methods []ast.FunctionNode
	for _, method := range methodMap {
		methods = append(methods, method)
	}

	return methods
}

// findMethodInClass finds a method in the class or its parent chain
func findMethodInClass(className string, methodName string, env *Environment) *ast.FunctionNode {
	classDef, exists := env.GetClass(className)
	if !exists {
		return nil
	}

	// Check this class's methods
	for _, method := range classDef.Methods {
		if method.Name == methodName {
			return &method
		}
	}

	// Check parent class if exists
	if classDef.Parent != "" {
		return findMethodInClass(classDef.Parent, methodName, env)
	}

	return nil
}
