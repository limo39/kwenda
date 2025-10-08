package parser

import (
	"lugha-yangu/ast"
	"lugha-yangu/lexer"
)

// ProgramNode represents the entire program
type ProgramNode struct {
	Functions []ast.ASTNode
	Imports   []ast.ImportNode
}

// ParseProgram parses the entire program with multiple functions and imports
func ParseProgram(tokens []lexer.Token) ProgramNode {
	var functions []ast.ASTNode
	var imports []ast.ImportNode
	i := 0

	for i < len(tokens) {
		// Skip empty tokens
		if tokens[i].Value == "" {
			i++
			continue
		}

		// Handle import statements (leta "module.swh")
		if tokens[i].Value == "leta" && i+1 < len(tokens) && tokens[i+1].Type == lexer.TokenString {
			imports = append(imports, ast.ImportNode{
				ModulePath: tokens[i+1].Value,
			})
			i += 2
			continue
		}

		// Handle top-level variable declarations
		if (tokens[i].Value == "namba" || tokens[i].Value == "maneno") && i+3 < len(tokens) && tokens[i+2].Value == "=" {
			end := i + 3
			for end < len(tokens) && tokens[end].Value != "namba" && tokens[end].Value != "maneno" && tokens[end].Value != "kazi" {
				end++
			}
			stmt := Parse(tokens[i:end])
			if stmt != nil {
				functions = append(functions, stmt)
			}
			i = end
			continue
		}

		// Find the end of the current function
		if tokens[i].Value == "kazi" {
			// Find the end of this function
			end := i + 1
			braceCount := 0
			foundFirstBrace := false

			for end < len(tokens) {
				if tokens[end].Value == "{" {
					braceCount++
					foundFirstBrace = true
				} else if tokens[end].Value == "}" {
					braceCount--
					if braceCount == 0 && foundFirstBrace {
						end++
						break
					}
				}
				end++
			}

			// Parse this function
			function := Parse(tokens[i:end])
			if function != nil {
				functions = append(functions, function)
			}
			i = end
		} else {
			i++
		}
	}

	return ProgramNode{Functions: functions, Imports: imports}
}

// Parse converts tokens into an Abstract Syntax Tree (AST)
func Parse(tokens []lexer.Token) ast.ASTNode {
	if len(tokens) == 0 {
		return nil
	}

	// Handle function definitions
	if tokens[0].Value == "kazi" && len(tokens) >= 4 && tokens[2].Value == "(" {
		return ParseFunctionDefinition(tokens)
	}

	// Handle return statements
	if tokens[0].Value == "rudisha" {
		if len(tokens) > 1 {
			return ast.ReturnNode{
				Value: ParseExpression(tokens[1:]),
			}
		}
		return ast.ReturnNode{Value: nil}
	}

	// Handle variable declarations
	if tokens[0].Value == "namba" && len(tokens) >= 4 && tokens[2].Value == "=" {
		return ast.VariableDeclarationNode{
			Name:  tokens[1].Value,
			Value: ParseExpression(tokens[3:]),
		}
	}

	// Handle string variable declarations
	if tokens[0].Value == "maneno" && len(tokens) >= 4 && tokens[2].Value == "=" {
		return ast.StringVariableDeclarationNode{
			Name:  tokens[1].Value,
			Value: ParseExpression(tokens[3:]),
		}
	}

	// Handle array declarations
	if tokens[0].Value == "orodha" && len(tokens) >= 5 && tokens[3].Value == "=" {
		arrayLiteral := ParseArrayLiteral(tokens[4:])
		var elements []ast.ASTNode
		if arrayNode, ok := arrayLiteral.(ast.ArrayNode); ok {
			elements = arrayNode.Elements
		}
		return ast.ArrayDeclarationNode{
			Name:     tokens[2].Value,
			Type:     tokens[1].Value,
			Elements: elements,
		}
	}

	// Handle function calls
	if (tokens[0].Value == "andika" || tokens[0].Value == "ongeza" || tokens[0].Value == "ondoa" || tokens[0].Value == "urefu_orodha" || tokens[0].Value == "pata" || tokens[0].Value == "soma" || tokens[0].Value == "andika_faili" || tokens[0].Value == "unda_faili" || tokens[0].Value == "faili_ipo" || tokens[0].Value == "ondoa_faili") && len(tokens) > 1 && tokens[1].Value == "(" {
		return ast.FunctionCallNode{
			Name: tokens[0].Value,
			Args: ParseArguments(tokens[2:]),
		}
	}

	// Handle identifiers (including module function calls like math.add)
	if tokens[0].Type == lexer.TokenIdentifier {
		return ast.IdentifierNode{Value: tokens[0].Value}
	}

	// Handle numbers
	if tokens[0].Type == lexer.TokenNumber {
		return ast.NumberNode{Value: tokens[0].Value}
	}

	return nil
}

// ParseBlock parses a block of statements
func ParseBlock(tokens []lexer.Token) []ast.ASTNode {
	var statements []ast.ASTNode
	i := 0

	for i < len(tokens) {
		if tokens[i].Value == "" {
			i++
			continue
		}

		// Parse return statements
		if tokens[i].Value == "rudisha" {
			end := i + 1
			for end < len(tokens) && tokens[end].Value != "namba" && tokens[end].Value != "maneno" && tokens[end].Value != "andika" && tokens[end].Value != "orodha" && tokens[end].Value != "rudisha" {
				end++
			}
			stmt := Parse(tokens[i:end])
			if stmt != nil {
				statements = append(statements, stmt)
			}
			i = end
			continue
		}

		// Parse variable declarations
		if (tokens[i].Value == "namba" || tokens[i].Value == "maneno") && i+3 < len(tokens) && tokens[i+2].Value == "=" {
			end := i + 3
			for end < len(tokens) && tokens[end].Value != "namba" && tokens[end].Value != "maneno" && tokens[end].Value != "andika" && tokens[end].Value != "orodha" {
				end++
			}
			stmt := Parse(tokens[i:end])
			if stmt != nil {
				statements = append(statements, stmt)
			}
			i = end
			continue
		}

		// Parse array declarations
		if tokens[i].Value == "orodha" && i+4 < len(tokens) && tokens[i+3].Value == "=" {
			end := i + 4
			bracketCount := 0
			for end < len(tokens) {
				if tokens[end].Value == "[" {
					bracketCount++
				} else if tokens[end].Value == "]" {
					bracketCount--
					if bracketCount == 0 {
						end++
						break
					}
				}
				end++
			}
			stmt := Parse(tokens[i:end])
			if stmt != nil {
				statements = append(statements, stmt)
			}
			i = end
			continue
		}

		// Parse function calls
		if i+1 < len(tokens) && tokens[i+1].Value == "(" {
			end := i + 2
			parenCount := 1
			for end < len(tokens) && parenCount > 0 {
				if tokens[end].Value == "(" {
					parenCount++
				} else if tokens[end].Value == ")" {
					parenCount--
				}
				end++
			}
			stmt := Parse(tokens[i:end])
			if stmt != nil {
				statements = append(statements, stmt)
			}
			i = end
			continue
		}

		i++
	}

	return statements
}

// ParseExpression parses an expression
func ParseExpression(tokens []lexer.Token) ast.ASTNode {
	if len(tokens) == 0 {
		return nil
	}

	// Handle array literals
	if len(tokens) >= 2 && tokens[0].Value == "[" {
		return ParseArrayLiteral(tokens)
	}

	// Handle function calls like ingiza("prompt")
	if len(tokens) >= 3 && tokens[1].Value == "(" {
		if tokens[0].Value == "ingiza" {
			for i := 2; i < len(tokens); i++ {
				if tokens[i].Value == ")" {
					if i > 2 && tokens[2].Type == lexer.TokenString {
						return ast.InputNode{Prompt: tokens[2].Value}
					}
					return ast.InputNode{}
				}
			}
		} else {
			return ast.FunctionCallNode{
				Name: tokens[0].Value,
				Args: ParseArguments(tokens[2:]),
			}
		}
	}

	// Handle binary operations
	if len(tokens) >= 3 && tokens[1].Type == lexer.TokenOperator && tokens[1].Value != "=" {
		return ast.BinaryOpNode{
			Left:  ParseExpression(tokens[:1]),
			Op:    tokens[1].Value,
			Right: ParseExpression(tokens[2:3]),
		}
	}

	// Handle string literals
	if tokens[0].Type == lexer.TokenString {
		return ast.IdentifierNode{Value: tokens[0].Value}
	}

	// Handle numbers
	if tokens[0].Type == lexer.TokenNumber {
		return ast.NumberNode{Value: tokens[0].Value}
	}

	// Handle identifiers
	if tokens[0].Type == lexer.TokenIdentifier {
		return ast.IdentifierNode{Value: tokens[0].Value}
	}

	return nil
}

// ParseArguments parses function arguments
func ParseArguments(tokens []lexer.Token) []ast.ASTNode {
	var args []ast.ASTNode
	var currentArg []lexer.Token

	for _, token := range tokens {
		if token.Value == "," || token.Value == ")" {
			if len(currentArg) > 0 {
				if len(currentArg) == 1 && currentArg[0].Type == lexer.TokenString {
					args = append(args, ast.IdentifierNode{Value: currentArg[0].Value})
				} else {
					args = append(args, ParseExpression(currentArg))
				}
				currentArg = nil
			}
		} else {
			currentArg = append(currentArg, token)
		}
	}

	return args
}

// ParseFunctionDefinition parses function definitions
func ParseFunctionDefinition(tokens []lexer.Token) ast.ASTNode {
	if len(tokens) < 4 || tokens[0].Value != "kazi" {
		return nil
	}

	functionName := tokens[1].Value

	// Parse parameters between ( and )
	var parameters []ast.Parameter
	if tokens[2].Value == "(" {
		// Find closing parenthesis
		parenEnd := -1
		for i := 3; i < len(tokens); i++ {
			if tokens[i].Value == ")" {
				parenEnd = i
				break
			}
		}

		if parenEnd > 3 {
			// Parse parameters
			parameters = ParseParameters(tokens[3:parenEnd])
		}
	}

	// Find the function body
	braceStart := -1
	braceEnd := -1
	braceCount := 0

	for i := 3; i < len(tokens); i++ {
		if tokens[i].Value == "{" {
			if braceCount == 0 {
				braceStart = i + 1
			}
			braceCount++
		} else if tokens[i].Value == "}" {
			braceCount--
			if braceCount == 0 {
				braceEnd = i
				break
			}
		}
	}

	if braceStart == -1 || braceEnd == -1 {
		return nil
	}

	// Parse the function body
	body := ParseBlock(tokens[braceStart:braceEnd])

	return ast.FunctionNode{
		Name:       functionName,
		Parameters: parameters,
		Body:       body,
	}
}

// ParseParameters parses function parameters
func ParseParameters(tokens []lexer.Token) []ast.Parameter {
	var parameters []ast.Parameter
	i := 0

	for i < len(tokens) {
		// Expect: type name, type name, ...
		if i+1 < len(tokens) {
			paramType := tokens[i].Value
			paramName := tokens[i+1].Value

			parameters = append(parameters, ast.Parameter{
				Name: paramName,
				Type: paramType,
			})

			i += 2

			// Skip comma if present
			if i < len(tokens) && tokens[i].Value == "," {
				i++
			}
		} else {
			break
		}
	}

	return parameters
}

// ParseArrayLiteral parses array literals
func ParseArrayLiteral(tokens []lexer.Token) ast.ASTNode {
	if len(tokens) < 2 || tokens[0].Value != "[" {
		return nil
	}

	closingBracket := -1
	for i := 1; i < len(tokens); i++ {
		if tokens[i].Value == "]" {
			closingBracket = i
			break
		}
	}

	if closingBracket == -1 {
		return nil
	}

	var elements []ast.ASTNode
	if closingBracket > 1 {
		elements = ParseArrayElements(tokens[1:closingBracket])
	}

	return ast.ArrayNode{Elements: elements}
}

// ParseArrayElements parses comma-separated array elements
func ParseArrayElements(tokens []lexer.Token) []ast.ASTNode {
	var elements []ast.ASTNode
	var currentElement []lexer.Token

	for _, token := range tokens {
		if token.Value == "," {
			if len(currentElement) > 0 {
				element := ParseExpression(currentElement)
				if element != nil {
					elements = append(elements, element)
				}
				currentElement = nil
			}
		} else {
			currentElement = append(currentElement, token)
		}
	}

	if len(currentElement) > 0 {
		element := ParseExpression(currentElement)
		if element != nil {
			elements = append(elements, element)
		}
	}

	return elements
}