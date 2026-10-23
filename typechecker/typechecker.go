package typechecker

import (
	"fmt"
	"kwenda/ast"
)

// Type represents a Kwenda type
type Type string

const (
	TypeNamba   Type = "namba"    // Number type
	TypeManeno  Type = "maneno"   // String type
	TypeBoolean Type = "boolean"  // Boolean type
	TypeOrodha  Type = "orodha"   // Array type
	TypeKamusi  Type = "kamusi"   // Dictionary type
	TypeSeti    Type = "seti"     // Set type
	TypeTuple   Type = "tuple"    // Tuple type
	TypeKazi    Type = "kazi"     // Function type
	TypeAny     Type = "any"      // Any type (no checking)
	TypeVoid    Type = "void"     // Void type (no return)
	TypeUnknown Type = "unknown"  // Unknown type
)

// TypeEnvironment stores type information for variables and functions
type TypeEnvironment struct {
	variables map[string]Type
	functions map[string]FunctionSignature
	parent    *TypeEnvironment
	strict    bool // Enable strict type checking
}

// FunctionSignature represents a function's type signature
type FunctionSignature struct {
	Name       string
	Parameters []ParameterType
	ReturnType Type
}

// ParameterType represents a parameter's type
type ParameterType struct {
	Name string
	Type Type
}

// TypeError represents a type checking error
type TypeError struct {
	Message  string
	Line     int
	Expected Type
	Actual   Type
}

func (e TypeError) Error() string {
	if e.Expected != "" && e.Actual != "" {
		return fmt.Sprintf("Hitilafu ya aina (Type error): %s. Inatarajiwa '%s', lakini imepata '%s' (Expected '%s', got '%s')", 
			e.Message, e.Expected, e.Actual, e.Expected, e.Actual)
	}
	return fmt.Sprintf("Hitilafu ya aina (Type error): %s", e.Message)
}

// NewTypeEnvironment creates a new type environment
func NewTypeEnvironment(strict bool) *TypeEnvironment {
	env := &TypeEnvironment{
		variables: make(map[string]Type),
		functions: make(map[string]FunctionSignature),
		parent:    nil,
		strict:    strict,
	}

	// Register built-in functions
	env.registerBuiltins()
	
	return env
}

// registerBuiltins registers built-in Kwenda functions
func (env *TypeEnvironment) registerBuiltins() {
	// andika - print function (accepts any types)
	env.SetFunction("andika", FunctionSignature{
		Name:       "andika",
		Parameters: []ParameterType{{Name: "values", Type: TypeAny}},
		ReturnType: TypeVoid,
	})

	// ingiza - input function
	env.SetFunction("ingiza", FunctionSignature{
		Name:       "ingiza",
		Parameters: []ParameterType{{Name: "prompt", Type: TypeManeno}},
		ReturnType: TypeManeno,
	})

	// urefu - length function
	env.SetFunction("urefu", FunctionSignature{
		Name:       "urefu",
		Parameters: []ParameterType{{Name: "str", Type: TypeManeno}},
		ReturnType: TypeNamba,
	})

	// Array functions
	env.SetFunction("ongeza", FunctionSignature{
		Name:       "ongeza",
		Parameters: []ParameterType{{Name: "arr", Type: TypeOrodha}, {Name: "item", Type: TypeAny}},
		ReturnType: TypeVoid,
	})

	env.SetFunction("ondoa", FunctionSignature{
		Name:       "ondoa",
		Parameters: []ParameterType{{Name: "arr", Type: TypeOrodha}, {Name: "index", Type: TypeNamba}},
		ReturnType: TypeVoid,
	})

	env.SetFunction("urefu_orodha", FunctionSignature{
		Name:       "urefu_orodha",
		Parameters: []ParameterType{{Name: "arr", Type: TypeOrodha}},
		ReturnType: TypeNamba,
	})

	env.SetFunction("pata", FunctionSignature{
		Name:       "pata",
		Parameters: []ParameterType{{Name: "arr", Type: TypeOrodha}, {Name: "index", Type: TypeNamba}},
		ReturnType: TypeAny,
	})
}

// NewChildEnvironment creates a child environment
func (env *TypeEnvironment) NewChildEnvironment() *TypeEnvironment {
	return &TypeEnvironment{
		variables: make(map[string]Type),
		functions: make(map[string]FunctionSignature),
		parent:    env,
		strict:    env.strict,
	}
}

// SetVariable sets a variable's type
func (env *TypeEnvironment) SetVariable(name string, typ Type) {
	env.variables[name] = typ
}

// GetVariable gets a variable's type
func (env *TypeEnvironment) GetVariable(name string) (Type, bool) {
	if typ, ok := env.variables[name]; ok {
		return typ, true
	}
	if env.parent != nil {
		return env.parent.GetVariable(name)
	}
	return TypeUnknown, false
}

// SetFunction sets a function's signature
func (env *TypeEnvironment) SetFunction(name string, sig FunctionSignature) {
	env.functions[name] = sig
}

// GetFunction gets a function's signature
func (env *TypeEnvironment) GetFunction(name string) (FunctionSignature, bool) {
	if sig, ok := env.functions[name]; ok {
		return sig, true
	}
	if env.parent != nil {
		return env.parent.GetFunction(name)
	}
	return FunctionSignature{}, false
}

// IsStrict returns whether strict type checking is enabled
func (env *TypeEnvironment) IsStrict() bool {
	return env.strict
}

// StringToType converts a string to a Type
func StringToType(s string) Type {
	switch s {
	case "namba":
		return TypeNamba
	case "maneno":
		return TypeManeno
	case "boolean":
		return TypeBoolean
	case "orodha":
		return TypeOrodha
	case "kamusi":
		return TypeKamusi
	case "seti":
		return TypeSeti
	case "tuple":
		return TypeTuple
	case "kazi":
		return TypeKazi
	case "any":
		return TypeAny
	case "void":
		return TypeVoid
	case "":
		return TypeAny // No type annotation means any
	default:
		return TypeUnknown
	}
}

// CheckType performs type checking on an AST node
func CheckType(node ast.ASTNode, env *TypeEnvironment) (Type, error) {
	if node == nil {
		return TypeVoid, nil
	}

	switch n := node.(type) {
	case ast.NumberNode:
		return TypeNamba, nil

	case ast.StringNode:
		return TypeManeno, nil

	case ast.BooleanNode:
		return TypeBoolean, nil

	case ast.IdentifierNode:
		if typ, ok := env.GetVariable(n.Value); ok {
			return typ, nil
		}
		// If not in strict mode, assume any type for undeclared variables
		if !env.IsStrict() {
			return TypeAny, nil
		}
		return TypeUnknown, TypeError{
			Message: fmt.Sprintf("Kigezo '%s' hakijatangazwa (Variable '%s' not declared)", n.Value, n.Value),
		}

	case ast.VariableDeclarationNode:
		// Check the value type
		valueType, err := CheckType(n.Value, env)
		if err != nil {
			return TypeUnknown, err
		}

		// Store variable type
		env.SetVariable(n.Name, valueType)
		return valueType, nil

	case ast.StringVariableDeclarationNode:
		// Check the value type
		valueType, err := CheckType(n.Value, env)
		if err != nil {
			return TypeUnknown, err
		}

		// Verify it's a string type
		if env.IsStrict() && valueType != TypeManeno && valueType != TypeAny {
			return TypeUnknown, TypeError{
				Message:  fmt.Sprintf("Kigezo cha maneno '%s' lazima kiwe cha aina 'maneno' (String variable '%s' must be of type 'maneno')", n.Name, n.Name),
				Expected: TypeManeno,
				Actual:   valueType,
			}
		}

		env.SetVariable(n.Name, TypeManeno)
		return TypeManeno, nil

	case ast.ArrayDeclarationNode:
		env.SetVariable(n.Name, TypeOrodha)
		return TypeOrodha, nil

	case ast.SetDeclarationNode:
		env.SetVariable(n.Name, TypeSeti)
		return TypeSeti, nil

	case ast.TupleDeclarationNode:
		env.SetVariable(n.Name, TypeTuple)
		return TypeTuple, nil

	case ast.DictionaryDeclarationNode:
		env.SetVariable(n.Name, TypeKamusi)
		return TypeKamusi, nil

	case ast.BinaryOpNode:
		leftType, err := CheckType(n.Left, env)
		if err != nil {
			return TypeUnknown, err
		}

		rightType, err := CheckType(n.Right, env)
		if err != nil {
			return TypeUnknown, err
		}

		// Check operator compatibility
		switch n.Op {
		case "+":
			// Addition can be numbers or string concatenation
			if leftType == TypeManeno || rightType == TypeManeno {
				// String concatenation
				return TypeManeno, nil
			}
			// Numeric addition
			if env.IsStrict() {
				if leftType != TypeNamba && leftType != TypeAny {
					return TypeUnknown, TypeError{
						Message:  fmt.Sprintf("Opereta '%s' inahitaji aina 'namba' au 'maneno' (Operator '%s' requires type 'namba' or 'maneno')", n.Op, n.Op),
						Expected: TypeNamba,
						Actual:   leftType,
					}
				}
				if rightType != TypeNamba && rightType != TypeAny {
					return TypeUnknown, TypeError{
						Message:  fmt.Sprintf("Opereta '%s' inahitaji aina 'namba' au 'maneno' (Operator '%s' requires type 'namba' or 'maneno')", n.Op, n.Op),
						Expected: TypeNamba,
						Actual:   rightType,
					}
				}
			}
			return TypeNamba, nil

		case "-", "*", "/", "%":
			// Arithmetic operators require numbers
			if env.IsStrict() {
				if leftType != TypeNamba && leftType != TypeAny {
					return TypeUnknown, TypeError{
						Message:  fmt.Sprintf("Opereta '%s' inahitaji aina 'namba' (Operator '%s' requires type 'namba')", n.Op, n.Op),
						Expected: TypeNamba,
						Actual:   leftType,
					}
				}
				if rightType != TypeNamba && rightType != TypeAny {
					return TypeUnknown, TypeError{
						Message:  fmt.Sprintf("Opereta '%s' inahitaji aina 'namba' (Operator '%s' requires type 'namba')", n.Op, n.Op),
						Expected: TypeNamba,
						Actual:   rightType,
					}
				}
			}
			return TypeNamba, nil

		case "==", "!=", "<", "<=", ">", ">=":
			// Comparison operators return boolean
			return TypeBoolean, nil

		case "na", "au":
			// Logical operators require booleans and return boolean
			if env.IsStrict() {
				if leftType != TypeBoolean && leftType != TypeAny {
					return TypeUnknown, TypeError{
						Message:  fmt.Sprintf("Opereta kimantiki '%s' inahitaji aina 'boolean' (Logical operator '%s' requires type 'boolean')", n.Op, n.Op),
						Expected: TypeBoolean,
						Actual:   leftType,
					}
				}
				if rightType != TypeBoolean && rightType != TypeAny {
					return TypeUnknown, TypeError{
						Message:  fmt.Sprintf("Opereta kimantiki '%s' inahitaji aina 'boolean' (Logical operator '%s' requires type 'boolean')", n.Op, n.Op),
						Expected: TypeBoolean,
						Actual:   rightType,
					}
				}
			}
			return TypeBoolean, nil

		default:
			return TypeAny, nil
		}

	case ast.FunctionNode:
		// Build function signature
		params := make([]ParameterType, len(n.Parameters))
		for i, param := range n.Parameters {
			params[i] = ParameterType{
				Name: param.Name,
				Type: StringToType(param.Type),
			}
		}

		returnType := StringToType(n.ReturnType)
		sig := FunctionSignature{
			Name:       n.Name,
			Parameters: params,
			ReturnType: returnType,
		}

		env.SetFunction(n.Name, sig)

		// Create child environment for function body
		funcEnv := env.NewChildEnvironment()

		// Add parameters to function environment
		for _, param := range params {
			funcEnv.SetVariable(param.Name, param.Type)
		}

		// Check function body
		for _, stmt := range n.Body {
			if retNode, ok := stmt.(ast.ReturnNode); ok {
				if retNode.Value != nil {
					retType, err := CheckType(retNode.Value, funcEnv)
					if err != nil {
						return TypeUnknown, err
					}

					// Check return type matches declared return type
					if env.IsStrict() && returnType != TypeAny && returnType != retType && retType != TypeAny {
						return TypeUnknown, TypeError{
							Message:  fmt.Sprintf("Aina ya kurudi ya kazi '%s' haifanani (Function '%s' return type mismatch)", n.Name, n.Name),
							Expected: returnType,
							Actual:   retType,
						}
					}
				}
			} else {
				_, err := CheckType(stmt, funcEnv)
				if err != nil {
					return TypeUnknown, err
				}
			}
		}

		return TypeKazi, nil

	case ast.FunctionCallNode:
		// Check if function exists
		if sig, ok := env.GetFunction(n.Name); ok {
			// For built-in functions (andika, etc.), skip argument count check
			isBuiltin := n.Name == "andika" || n.Name == "ingiza" || n.Name == "urefu"
			
			// Check argument count (except for builtins which are variadic)
			if env.IsStrict() && !isBuiltin && len(n.Args) != len(sig.Parameters) {
				return TypeUnknown, TypeError{
					Message: fmt.Sprintf("Kazi '%s' inahitaji hoja %d, lakini imepata %d (Function '%s' expects %d arguments, got %d)",
						n.Name, len(sig.Parameters), len(n.Args), n.Name, len(sig.Parameters), len(n.Args)),
				}
			}

			// Check argument types (skip for builtins which accept any)
			if env.IsStrict() && !isBuiltin {
				for i, arg := range n.Args {
					if i < len(sig.Parameters) {
						argType, err := CheckType(arg, env)
						if err != nil {
							return TypeUnknown, err
						}

						expectedType := sig.Parameters[i].Type
						if expectedType != TypeAny && argType != TypeAny && argType != expectedType {
							return TypeUnknown, TypeError{
								Message: fmt.Sprintf("Hoja %d ya kazi '%s' lazima iwe aina '%s' (Argument %d of function '%s' must be of type '%s')",
									i+1, n.Name, expectedType, i+1, n.Name, expectedType),
								Expected: expectedType,
								Actual:   argType,
							}
						}
					}
				}
			}

			return sig.ReturnType, nil
		}

		// If function not found in strict mode, error
		if env.IsStrict() {
			return TypeUnknown, TypeError{
				Message: fmt.Sprintf("Kazi '%s' haijatangazwa (Function '%s' not declared)", n.Name, n.Name),
			}
		}

		return TypeAny, nil

	case ast.IfNode:
		// Check condition is boolean (in strict mode)
		condType, err := CheckType(n.Condition, env)
		if err != nil {
			return TypeUnknown, err
		}

		if env.IsStrict() && condType != TypeBoolean && condType != TypeAny {
			return TypeUnknown, TypeError{
				Message:  "Sharti la 'kama' lazima liwe aina 'boolean' (Condition of 'kama' must be type 'boolean')",
				Expected: TypeBoolean,
				Actual:   condType,
			}
		}

		// Check then body
		childEnv := env.NewChildEnvironment()
		for _, stmt := range n.ThenBody {
			_, err := CheckType(stmt, childEnv)
			if err != nil {
				return TypeUnknown, err
			}
		}

		// Check else body
		if len(n.ElseBody) > 0 {
			elseEnv := env.NewChildEnvironment()
			for _, stmt := range n.ElseBody {
				_, err := CheckType(stmt, elseEnv)
				if err != nil {
					return TypeUnknown, err
				}
			}
		}

		return TypeVoid, nil

	case ast.WhileNode:
		// Check condition is boolean (in strict mode)
		condType, err := CheckType(n.Condition, env)
		if err != nil {
			return TypeUnknown, err
		}

		if env.IsStrict() && condType != TypeBoolean && condType != TypeAny {
			return TypeUnknown, TypeError{
				Message:  "Sharti la 'wakati' lazima liwe aina 'boolean' (Condition of 'wakati' must be type 'boolean')",
				Expected: TypeBoolean,
				Actual:   condType,
			}
		}

		// Check body
		childEnv := env.NewChildEnvironment()
		for _, stmt := range n.Body {
			_, err := CheckType(stmt, childEnv)
			if err != nil {
				return TypeUnknown, err
			}
		}

		return TypeVoid, nil

	case ast.ForNode:
		// Check condition is boolean (in strict mode)
		if n.Condition != nil {
			condType, err := CheckType(n.Condition, env)
			if err != nil {
				return TypeUnknown, err
			}

			if env.IsStrict() && condType != TypeBoolean && condType != TypeAny {
				return TypeUnknown, TypeError{
					Message:  "Sharti la 'kwa' lazima liwe aina 'boolean' (Condition of 'kwa' must be type 'boolean')",
					Expected: TypeBoolean,
					Actual:   condType,
				}
			}
		}

		// Check init, update, and body
		childEnv := env.NewChildEnvironment()
		if n.Init != nil {
			_, err := CheckType(n.Init, childEnv)
			if err != nil {
				return TypeUnknown, err
			}
		}

		for _, stmt := range n.Body {
			_, err := CheckType(stmt, childEnv)
			if err != nil {
				return TypeUnknown, err
			}
		}

		if n.Update != nil {
			_, err := CheckType(n.Update, childEnv)
			if err != nil {
				return TypeUnknown, err
			}
		}

		return TypeVoid, nil

	case ast.ReturnNode:
		if n.Value != nil {
			return CheckType(n.Value, env)
		}
		return TypeVoid, nil

	case ast.ArrayNode:
		return TypeOrodha, nil

	case ast.SetNode:
		return TypeSeti, nil

	case ast.TupleNode:
		return TypeTuple, nil

	case ast.DictionaryNode:
		return TypeKamusi, nil

	default:
		// For unknown node types, return Any in non-strict mode
		if !env.IsStrict() {
			return TypeAny, nil
		}
		return TypeUnknown, nil
	}
}

// CheckProgram performs type checking on a program
func CheckProgram(program ast.ASTNode, strict bool) []error {
	env := NewTypeEnvironment(strict)
	var errors []error

	// Handle different program node types
	switch p := program.(type) {
	case []ast.ASTNode:
		for _, node := range p {
			_, err := CheckType(node, env)
			if err != nil {
				errors = append(errors, err)
			}
		}
	default:
		_, err := CheckType(program, env)
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}
