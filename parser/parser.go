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

		// Handle class definitions
		if tokens[i].Value == "darasa" {
			// Find the end of this class
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

			// Parse this class
			class := Parse(tokens[i:end])
			if class != nil {
				functions = append(functions, class)
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

	// Handle conditional statements
	if tokens[0].Value == "kama" {
		return ParseIfStatement(tokens)
	}

	// Handle while loops
	if tokens[0].Value == "wakati" {
		return ParseWhileStatement(tokens)
	}

	// Handle for loops
	if tokens[0].Value == "kwa" {
		return ParseForStatement(tokens)
	}

	// Handle break statements
	if tokens[0].Value == "vunja" {
		return ast.BreakNode{}
	}

	// Handle continue statements
	if tokens[0].Value == "endelea" {
		return ast.ContinueNode{}
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

	// Handle try-catch statements
	if tokens[0].Value == "jaribu" {
		return ParseTryStatement(tokens)
	}

	// Handle throw statements
	if tokens[0].Value == "tupa" && len(tokens) > 1 {
		var endIndex = len(tokens)
		for i := 1; i < len(tokens); i++ {
			if tokens[i].Value == "}" || tokens[i].Value == ";" {
				endIndex = i
				break
			}
		}
		return ast.ThrowNode{Message: ParseExpression(tokens[1:endIndex])}
	}

	// Handle class definitions
	if tokens[0].Value == "darasa" {
		return ParseClassDefinition(tokens)
	}

	// Handle class instantiation (unda ClassName(args))
	if tokens[0].Value == "unda" && len(tokens) >= 3 {
		return ParseNewInstance(tokens)
	}

	// Handle method calls as statements (e.g., object.method(args))
	// Must check before regular function calls
	if len(tokens) >= 4 && tokens[0].Type == lexer.TokenIdentifier && tokens[1].Value == "." && tokens[3].Value == "(" {
		var object ast.ASTNode
		if tokens[0].Value == "hii" {
			object = ast.ThisNode{}
		} else {
			object = ast.IdentifierNode{Value: tokens[0].Value}
		}
		return ast.MethodCallNode{
			Object: object,
			Method: tokens[2].Value,
			Args:   ParseArguments(tokens[4:]),
		}
	}

	// Handle function definitions
	if tokens[0].Value == "kazi" && len(tokens) >= 4 && tokens[2].Value == "(" {
		return ParseFunctionDefinition(tokens)
	}

	// Handle variable declarations
	if tokens[0].Value == "namba" && len(tokens) >= 4 && tokens[2].Value == "=" {
		return ast.VariableDeclarationNode{
			Name:  tokens[1].Value,
			Value: ParseExpression(tokens[3:]),
		}
	}

	// Handle boolean variable declarations
	if tokens[0].Value == "boolean" && len(tokens) >= 4 && tokens[2].Value == "=" {
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

	// Handle dictionary declarations (kamusi name = {...} or kamusi name = func())
	if tokens[0].Value == "kamusi" && len(tokens) >= 4 && tokens[2].Value == "=" {
		var value ast.ASTNode
		if tokens[3].Value == "{" {
			value = ParseDictionaryLiteral(tokens[3:])
		} else {
			value = ParseExpression(tokens[3:])
		}
		return ast.DictionaryDeclarationNode{
			Name:  tokens[1].Value,
			Value: value,
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

	// Handle member assignment (e.g., hii.jina = "Amina")
	if len(tokens) >= 5 && tokens[1].Value == "." && tokens[3].Value == "=" {
		var object ast.ASTNode
		if tokens[0].Value == "hii" {
			object = ast.ThisNode{}
		} else {
			object = ast.IdentifierNode{Value: tokens[0].Value}
		}
		return ast.MemberAssignmentNode{
			Object: object,
			Member: tokens[2].Value,
			Value:  ParseExpression(tokens[4:]),
		}
	}

	// Handle array/dictionary assignment (must come before array access!)
	if len(tokens) >= 6 && tokens[0].Type == lexer.TokenIdentifier && tokens[1].Value == "[" {
		bracketEnd := -1
		for i := 2; i < len(tokens); i++ {
			if tokens[i].Value == "]" {
				bracketEnd = i
				break
			}
		}
		if bracketEnd != -1 && bracketEnd+1 < len(tokens) && tokens[bracketEnd+1].Value == "=" {
			return ast.ArrayAssignmentNode{
				Array: ast.IdentifierNode{Value: tokens[0].Value},
				Index: ParseExpression(tokens[2:bracketEnd]),
				Value: ParseExpression(tokens[bracketEnd+2:]),
			}
		}
	}

	// Handle array/dictionary access
	if len(tokens) >= 4 && tokens[0].Type == lexer.TokenIdentifier && tokens[1].Value == "[" {
		for i := 2; i < len(tokens); i++ {
			if tokens[i].Value == "]" {
				return ast.ArrayAccessNode{
					Array: ast.IdentifierNode{Value: tokens[0].Value},
					Index: ParseExpression(tokens[2:i]),
				}
			}
		}
	}

	// Handle assignment statements
	if len(tokens) >= 3 && tokens[0].Type == lexer.TokenIdentifier && tokens[1].Value == "=" {
		return ast.VariableDeclarationNode{
			Name:  tokens[0].Value,
			Value: ParseExpression(tokens[2:]),
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

	// Handle method calls with dot notation (e.g., object.method(args))
	// Must check before regular function calls
	if len(tokens) >= 4 && tokens[1].Value == "." && tokens[3].Value == "(" {
		var object ast.ASTNode
		if tokens[0].Value == "hii" {
			object = ast.ThisNode{}
		} else {
			object = ast.IdentifierNode{Value: tokens[0].Value}
		}
		return ast.MethodCallNode{
			Object: object,
			Method: tokens[2].Value,
			Args:   ParseArguments(tokens[4:]),
		}
	}

	// Handle function calls (built-in and user-defined)
	if tokens[0].Type == lexer.TokenIdentifier && len(tokens) > 1 && tokens[1].Value == "(" {
		return ast.FunctionCallNode{
			Name: tokens[0].Value,
			Args: ParseArguments(tokens[2:]),
		}
	}
	
	// Legacy: Handle specific built-in function calls
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

		// Parse control flow statements (if, while, for, try)
		if tokens[i].Value == "kama" || tokens[i].Value == "wakati" || tokens[i].Value == "kwa" || tokens[i].Value == "jaribu" {
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
						// Check for else/catch/finally
						if end+1 < len(tokens) && (tokens[end+1].Value == "sivyo" || tokens[end+1].Value == "shika" || tokens[end+1].Value == "hatimaye") {
							end++
							continue
						} else {
							end++
							break
						}
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

		// Parse simple statements (break, continue, throw)
		if tokens[i].Value == "vunja" || tokens[i].Value == "endelea" {
			stmt := Parse(tokens[i : i+1])
			if stmt != nil {
				statements = append(statements, stmt)
			}
			i++
			continue
		}

		// Parse throw statements
		if tokens[i].Value == "tupa" {
			end := i + 1
			for end < len(tokens) && tokens[end].Value != "namba" && tokens[end].Value != "maneno" && tokens[end].Value != "andika" && tokens[end].Value != "kama" && tokens[end].Value != "wakati" {
				end++
			}
			stmt := Parse(tokens[i:end])
			if stmt != nil {
				statements = append(statements, stmt)
			}
			i = end
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
		if (tokens[i].Value == "namba" || tokens[i].Value == "maneno" || tokens[i].Value == "boolean") && i+3 < len(tokens) && tokens[i+2].Value == "=" {
			end := i + 3
			for end < len(tokens) {
				// Stop at keywords that start new statements
				if tokens[end].Value == "namba" || tokens[end].Value == "maneno" || tokens[end].Value == "boolean" || tokens[end].Value == "andika" || tokens[end].Value == "orodha" || tokens[end].Value == "kama" || tokens[end].Value == "wakati" || tokens[end].Value == "kwa" || tokens[end].Value == "vunja" || tokens[end].Value == "endelea" || tokens[end].Value == "rudisha" || tokens[end].Value == "tupa" {
					break
				}
				// Stop at assignment statements (identifier followed by =)
				if end+1 < len(tokens) && tokens[end].Type == lexer.TokenIdentifier && tokens[end+1].Value == "=" {
					break
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

		// Parse dictionary declarations
		if tokens[i].Value == "kamusi" && i+3 < len(tokens) && tokens[i+2].Value == "=" {
			end := i + 3
			// Check if it's a dictionary literal or function call
			if tokens[end].Value == "{" {
				// Dictionary literal
				braceCount := 0
				for end < len(tokens) {
					if tokens[end].Value == "{" {
						braceCount++
					} else if tokens[end].Value == "}" {
						braceCount--
						if braceCount == 0 {
							end++
							break
						}
					}
					end++
				}
			} else {
				// Function call or expression
				for end < len(tokens) {
					if tokens[end].Value == "namba" || tokens[end].Value == "maneno" || tokens[end].Value == "boolean" || tokens[end].Value == "kamusi" || tokens[end].Value == "andika" || tokens[end].Value == "orodha" || tokens[end].Value == "kama" || tokens[end].Value == "wakati" || tokens[end].Value == "kwa" || tokens[end].Value == "vunja" || tokens[end].Value == "endelea" || tokens[end].Value == "rudisha" || tokens[end].Value == "tupa" {
						break
					}
					if end+1 < len(tokens) && tokens[end].Type == lexer.TokenIdentifier && tokens[end+1].Value == "=" {
						break
					}
					end++
				}
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

		// Parse member assignment (e.g., obj.property = value or hii.property = value)
		if i+4 < len(tokens) && (tokens[i].Type == lexer.TokenIdentifier || tokens[i].Value == "hii") && tokens[i+1].Value == "." && tokens[i+3].Value == "=" {
			end := i + 4
			for end < len(tokens) {
				if tokens[end].Value == "namba" || tokens[end].Value == "maneno" || tokens[end].Value == "boolean" || tokens[end].Value == "kamusi" || tokens[end].Value == "andika" || tokens[end].Value == "orodha" || tokens[end].Value == "kama" || tokens[end].Value == "wakati" || tokens[end].Value == "kwa" || tokens[end].Value == "vunja" || tokens[end].Value == "endelea" || tokens[end].Value == "rudisha" || tokens[end].Value == "tupa" {
					break
				}
				// Check for next member assignment (hii.property = value) or regular assignment (identifier = value)
				if end+3 < len(tokens) && (tokens[end].Type == lexer.TokenIdentifier || tokens[end].Value == "hii") && tokens[end+1].Value == "." && tokens[end+3].Value == "=" {
					break // Next member assignment found
				}
				if end+1 < len(tokens) && tokens[end].Type == lexer.TokenIdentifier && tokens[end+1].Value == "=" {
					break // Next regular assignment found
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

		// Parse array/dictionary assignment (e.g., arr[0] = 5 or dict["key"] = value)
		if i+3 < len(tokens) && tokens[i].Type == lexer.TokenIdentifier && tokens[i+1].Value == "[" {
			bracketEnd := -1
			for j := i + 2; j < len(tokens); j++ {
				if tokens[j].Value == "]" {
					bracketEnd = j
					break
				}
			}
			if bracketEnd != -1 && bracketEnd+1 < len(tokens) && tokens[bracketEnd+1].Value == "=" {
				end := bracketEnd + 2
				for end < len(tokens) {
					if tokens[end].Value == "namba" || tokens[end].Value == "maneno" || tokens[end].Value == "boolean" || tokens[end].Value == "andika" || tokens[end].Value == "orodha" || tokens[end].Value == "kamusi" || tokens[end].Value == "kama" || tokens[end].Value == "wakati" || tokens[end].Value == "kwa" || tokens[end].Value == "vunja" || tokens[end].Value == "endelea" || tokens[end].Value == "rudisha" || tokens[end].Value == "tupa" {
						break
					}
					if end+1 < len(tokens) && tokens[end].Type == lexer.TokenIdentifier && tokens[end+1].Value == "=" {
						break
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
		}

		// Parse assignment statements (e.g., i = i + 1)
		if i+2 < len(tokens) && tokens[i].Type == lexer.TokenIdentifier && tokens[i+1].Value == "=" {
			end := i + 2
			for end < len(tokens) {
				// Stop at keywords that start new statements
				if tokens[end].Value == "namba" || tokens[end].Value == "maneno" || tokens[end].Value == "boolean" || tokens[end].Value == "andika" || tokens[end].Value == "orodha" || tokens[end].Value == "kama" || tokens[end].Value == "wakati" || tokens[end].Value == "kwa" || tokens[end].Value == "vunja" || tokens[end].Value == "endelea" || tokens[end].Value == "rudisha" || tokens[end].Value == "tupa" {
					break
				}
				// Stop at another assignment statement (identifier followed by =)
				if end+1 < len(tokens) && tokens[end].Type == lexer.TokenIdentifier && tokens[end+1].Value == "=" {
					break
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

		// Parse method calls (e.g., object.method(args))
		if i+3 < len(tokens) && tokens[i].Type == lexer.TokenIdentifier && tokens[i+1].Value == "." && tokens[i+3].Value == "(" {
			end := i + 4
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

	// Handle class instantiation (unda ClassName(args))
	if tokens[0].Value == "unda" && len(tokens) >= 3 {
		return ParseNewInstance(tokens)
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

	// Handle binary operations - need to find the operator position
	// Skip over complex expressions like array access, member access, function calls
	opIndex := -1
	parenDepth := 0
	bracketDepth := 0
	
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Value == "(" {
			parenDepth++
		} else if tokens[i].Value == ")" {
			parenDepth--
		} else if tokens[i].Value == "[" {
			bracketDepth++
		} else if tokens[i].Value == "]" {
			bracketDepth--
		} else if parenDepth == 0 && bracketDepth == 0 && tokens[i].Type == lexer.TokenOperator && tokens[i].Value != "=" {
			opIndex = i
			break
		}
	}
	
	if opIndex > 0 && opIndex < len(tokens)-1 {
		return ast.BinaryOpNode{
			Left:  ParseExpression(tokens[:opIndex]),
			Op:    tokens[opIndex].Value,
			Right: ParseExpression(tokens[opIndex+1:]),
		}
	}

	// Handle method calls with dot notation (e.g., object.method(args)) - MUST come before member access!
	if len(tokens) >= 4 && tokens[1].Value == "." && tokens[3].Value == "(" {
		var object ast.ASTNode
		if tokens[0].Value == "hii" {
			object = ast.ThisNode{}
		} else {
			object = ast.IdentifierNode{Value: tokens[0].Value}
		}
		return ast.MethodCallNode{
			Object: object,
			Method: tokens[2].Value,
			Args:   ParseArguments(tokens[4:]),
		}
	}

	// Handle member access with dot notation (e.g., hii.jina) - MUST come before 'hii' keyword check!
	if len(tokens) >= 3 && tokens[1].Value == "." {
		var object ast.ASTNode
		if tokens[0].Value == "hii" {
			object = ast.ThisNode{}
		} else {
			object = ast.IdentifierNode{Value: tokens[0].Value}
		}
		return ast.MemberAccessNode{
			Object: object,
			Member: tokens[2].Value,
		}
	}

	// Handle 'hii' keyword (this/self)
	if tokens[0].Value == "hii" {
		return ast.ThisNode{}
	}

	// Handle boolean literals
	if tokens[0].Value == "kweli" {
		return ast.BooleanNode{Value: true}
	}
	if tokens[0].Value == "uwongo" {
		return ast.BooleanNode{Value: false}
	}

	// Handle string literals
	if tokens[0].Type == lexer.TokenString {
		return ast.StringNode{Value: tokens[0].Value}
	}

	// Handle numbers
	if tokens[0].Type == lexer.TokenNumber {
		return ast.NumberNode{Value: tokens[0].Value}
	}



	// Handle array/dictionary access (e.g., arr[0] or dict["key"])
	if len(tokens) >= 4 && tokens[0].Type == lexer.TokenIdentifier && tokens[1].Value == "[" {
		for i := 2; i < len(tokens); i++ {
			if tokens[i].Value == "]" {
				return ast.ArrayAccessNode{
					Array: ast.IdentifierNode{Value: tokens[0].Value},
					Index: ParseExpression(tokens[2:i]),
				}
			}
		}
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
					args = append(args, ast.StringNode{Value: currentArg[0].Value})
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

// ParseIfStatement parses conditional statements (kama ... { ... } sivyo { ... })
func ParseIfStatement(tokens []lexer.Token) ast.ASTNode {
	if len(tokens) < 4 || tokens[0].Value != "kama" {
		return nil
	}

	// Find the condition (between "kama" and "{")
	conditionStart := 1
	conditionEnd := -1
	for i := conditionStart; i < len(tokens); i++ {
		if tokens[i].Value == "{" {
			conditionEnd = i
			break
		}
	}

	if conditionEnd == -1 {
		return nil
	}

	// Parse the condition
	condition := ParseExpression(tokens[conditionStart:conditionEnd])

	// Find the then body (between first "{" and "}")
	thenStart := conditionEnd + 1
	thenEnd := -1
	braceCount := 1
	for i := thenStart; i < len(tokens); i++ {
		if tokens[i].Value == "{" {
			braceCount++
		} else if tokens[i].Value == "}" {
			braceCount--
			if braceCount == 0 {
				thenEnd = i
				break
			}
		}
	}

	if thenEnd == -1 {
		return nil
	}

	// Parse the then body
	thenBody := ParseBlock(tokens[thenStart:thenEnd])

	// Check for else clause (sivyo)
	var elseBody []ast.ASTNode
	if thenEnd+1 < len(tokens) && tokens[thenEnd+1].Value == "sivyo" {
		// Find the else body
		elseStart := -1
		elseEnd := -1

		// Look for opening brace after "sivyo"
		for i := thenEnd + 2; i < len(tokens); i++ {
			if tokens[i].Value == "{" {
				elseStart = i + 1
				break
			}
		}

		if elseStart != -1 {
			// Find closing brace for else body
			braceCount = 1
			for i := elseStart; i < len(tokens); i++ {
				if tokens[i].Value == "{" {
					braceCount++
				} else if tokens[i].Value == "}" {
					braceCount--
					if braceCount == 0 {
						elseEnd = i
						break
					}
				}
			}

			if elseEnd != -1 {
				elseBody = ParseBlock(tokens[elseStart:elseEnd])
			}
		}
	}

	return ast.IfNode{
		Condition: condition,
		ThenBody:  thenBody,
		ElseBody:  elseBody,
	}
}

// ParseWhileStatement parses while loops (wakati condition { ... })
func ParseWhileStatement(tokens []lexer.Token) ast.ASTNode {
	if len(tokens) < 4 || tokens[0].Value != "wakati" {
		return nil
	}

	// Find the condition (between "wakati" and "{")
	conditionStart := 1
	conditionEnd := -1
	for i := conditionStart; i < len(tokens); i++ {
		if tokens[i].Value == "{" {
			conditionEnd = i
			break
		}
	}

	if conditionEnd == -1 {
		return nil
	}

	// Parse the condition
	condition := ParseExpression(tokens[conditionStart:conditionEnd])

	// Find the body (between "{" and "}")
	bodyStart := conditionEnd + 1
	bodyEnd := -1
	braceCount := 1
	for i := bodyStart; i < len(tokens); i++ {
		if tokens[i].Value == "{" {
			braceCount++
		} else if tokens[i].Value == "}" {
			braceCount--
			if braceCount == 0 {
				bodyEnd = i
				break
			}
		}
	}

	if bodyEnd == -1 {
		return nil
	}

	// Parse the body
	body := ParseBlock(tokens[bodyStart:bodyEnd])

	return ast.WhileNode{
		Condition: condition,
		Body:      body,
	}
}

// ParseForStatement parses for loops (kwa init; condition; update { ... })
func ParseForStatement(tokens []lexer.Token) ast.ASTNode {
	if len(tokens) < 4 || tokens[0].Value != "kwa" {
		return nil
	}

	// Find the parts of the for loop (init; condition; update)
	var initEnd, conditionEnd, updateEnd int = -1, -1, -1
	semicolonCount := 0

	for i := 1; i < len(tokens); i++ {
		if tokens[i].Value == ";" {
			semicolonCount++
			if semicolonCount == 1 {
				initEnd = i
			} else if semicolonCount == 2 {
				conditionEnd = i
			}
		} else if tokens[i].Value == "{" {
			updateEnd = i
			break
		}
	}

	// If we don't have semicolons, treat it as a simple for loop with just condition
	if semicolonCount == 0 {
		conditionStart := 1
		conditionEnd := -1
		for i := conditionStart; i < len(tokens); i++ {
			if tokens[i].Value == "{" {
				conditionEnd = i
				break
			}
		}

		if conditionEnd == -1 {
			return nil
		}

		condition := ParseExpression(tokens[conditionStart:conditionEnd])

		// Find the body
		bodyStart := conditionEnd + 1
		bodyEnd := -1
		braceCount := 1
		for i := bodyStart; i < len(tokens); i++ {
			if tokens[i].Value == "{" {
				braceCount++
			} else if tokens[i].Value == "}" {
				braceCount--
				if braceCount == 0 {
					bodyEnd = i
					break
				}
			}
		}

		if bodyEnd == -1 {
			return nil
		}

		body := ParseBlock(tokens[bodyStart:bodyEnd])

		return ast.ForNode{
			Init:      nil,
			Condition: condition,
			Update:    nil,
			Body:      body,
		}
	}

	// Full for loop with init; condition; update
	if initEnd == -1 || conditionEnd == -1 || updateEnd == -1 {
		return nil
	}

	// Parse init, condition, and update
	var init, condition, update ast.ASTNode

	if initEnd > 1 {
		init = ParseExpression(tokens[1:initEnd])
	}

	if conditionEnd > initEnd+1 {
		condition = ParseExpression(tokens[initEnd+1 : conditionEnd])
	}

	if updateEnd > conditionEnd+1 {
		update = ParseExpression(tokens[conditionEnd+1 : updateEnd])
	}

	// Find the body
	bodyStart := updateEnd + 1
	bodyEnd := -1
	braceCount := 1
	for i := bodyStart; i < len(tokens); i++ {
		if tokens[i].Value == "{" {
			braceCount++
		} else if tokens[i].Value == "}" {
			braceCount--
			if braceCount == 0 {
				bodyEnd = i
				break
			}
		}
	}

	if bodyEnd == -1 {
		return nil
	}

	body := ParseBlock(tokens[bodyStart:bodyEnd])

	return ast.ForNode{
		Init:      init,
		Condition: condition,
		Update:    update,
		Body:      body,
	}
}

// ParseTryStatement parses try-catch statements (jaribu { ... } shika (var) { ... } hatimaye { ... })
func ParseTryStatement(tokens []lexer.Token) ast.ASTNode {
	if len(tokens) < 4 || tokens[0].Value != "jaribu" {
		return nil
	}

	var tryBody []ast.ASTNode
	var catchVar string
	var catchBody []ast.ASTNode
	var finallyBody []ast.ASTNode

	i := 1
	// Parse try block
	if i < len(tokens) && tokens[i].Value == "{" {
		i++ // Skip opening brace
		tryStart := i
		braceCount := 1
		for i < len(tokens) && braceCount > 0 {
			if tokens[i].Value == "{" {
				braceCount++
			} else if tokens[i].Value == "}" {
				braceCount--
			}
			i++
		}
		tryEnd := i - 1
		tryBody = ParseBlock(tokens[tryStart:tryEnd])
	}

	// Parse catch block if present
	if i < len(tokens) && tokens[i].Value == "shika" {
		i++ // Skip "shika"

		// Parse catch variable
		if i < len(tokens) && tokens[i].Value == "(" {
			i++ // Skip "("
			if i < len(tokens) && tokens[i].Type == lexer.TokenIdentifier {
				catchVar = tokens[i].Value
				i++
			}
			if i < len(tokens) && tokens[i].Value == ")" {
				i++ // Skip ")"
			}
		}

		// Parse catch body
		if i < len(tokens) && tokens[i].Value == "{" {
			i++ // Skip opening brace
			catchStart := i
			braceCount := 1
			for i < len(tokens) && braceCount > 0 {
				if tokens[i].Value == "{" {
					braceCount++
				} else if tokens[i].Value == "}" {
					braceCount--
				}
				i++
			}
			catchEnd := i - 1
			catchBody = ParseBlock(tokens[catchStart:catchEnd])
		}
	}

	// Parse finally block if present
	if i < len(tokens) && tokens[i].Value == "hatimaye" {
		i++ // Skip "hatimaye"

		if i < len(tokens) && tokens[i].Value == "{" {
			i++ // Skip opening brace
			finallyStart := i
			braceCount := 1
			for i < len(tokens) && braceCount > 0 {
				if tokens[i].Value == "{" {
					braceCount++
				} else if tokens[i].Value == "}" {
					braceCount--
				}
				i++
			}
			finallyEnd := i - 1
			finallyBody = ParseBlock(tokens[finallyStart:finallyEnd])
		}
	}

	return ast.TryNode{
		TryBody:     tryBody,
		CatchVar:    catchVar,
		CatchBody:   catchBody,
		FinallyBody: finallyBody,
	}
}

// ParseDictionaryLiteral parses dictionary literals like {key: value, key2: value2}
func ParseDictionaryLiteral(tokens []lexer.Token) ast.ASTNode {
	if len(tokens) == 0 {
		return nil
	}
	
	// Handle empty dictionary {}
	if len(tokens) == 2 && tokens[0].Value == "{" && tokens[1].Value == "}" {
		return ast.DictionaryNode{Pairs: []ast.DictionaryPair{}}
	}
	
	if tokens[0].Value != "{" {
		return nil
	}

	closingBrace := -1
	braceCount := 1
	for i := 1; i < len(tokens); i++ {
		if tokens[i].Value == "{" {
			braceCount++
		} else if tokens[i].Value == "}" {
			braceCount--
			if braceCount == 0 {
				closingBrace = i
				break
			}
		}
	}

	if closingBrace == -1 {
		return nil
	}

	var pairs []ast.DictionaryPair
	if closingBrace > 1 {
		pairs = ParseDictionaryPairs(tokens[1:closingBrace])
	}

	return ast.DictionaryNode{Pairs: pairs}
}

// ParseDictionaryPairs parses comma-separated key:value pairs
func ParseDictionaryPairs(tokens []lexer.Token) []ast.DictionaryPair {
	var pairs []ast.DictionaryPair
	var currentPair []lexer.Token

	for _, token := range tokens {
		if token.Value == "," {
			if len(currentPair) > 0 {
				pair := ParseDictionaryPair(currentPair)
				if pair != nil {
					pairs = append(pairs, *pair)
				}
				currentPair = nil
			}
		} else {
			currentPair = append(currentPair, token)
		}
	}

	if len(currentPair) > 0 {
		pair := ParseDictionaryPair(currentPair)
		if pair != nil {
			pairs = append(pairs, *pair)
		}
	}

	return pairs
}

// ParseDictionaryPair parses a single key:value pair
func ParseDictionaryPair(tokens []lexer.Token) *ast.DictionaryPair {
	// Find the colon separator
	colonIndex := -1
	for i, token := range tokens {
		if token.Value == ":" {
			colonIndex = i
			break
		}
	}

	if colonIndex == -1 || colonIndex == 0 || colonIndex == len(tokens)-1 {
		return nil
	}

	// Parse key - if it's a single string token, treat it as a string literal
	var key ast.ASTNode
	if colonIndex == 1 && tokens[0].Type == lexer.TokenString {
		// String literal key
		key = ast.StringNode{Value: tokens[0].Value}
	} else {
		// Expression key
		key = ParseExpression(tokens[:colonIndex])
	}
	
	value := ParseExpression(tokens[colonIndex+1:])

	if key == nil || value == nil {
		return nil
	}

	return &ast.DictionaryPair{
		Key:   key,
		Value: value,
	}
}


// ParseClassDefinition parses class definitions (darasa ClassName { ... } or darasa Child : Parent { ... })
func ParseClassDefinition(tokens []lexer.Token) ast.ASTNode {
	if len(tokens) < 3 || tokens[0].Value != "darasa" {
		return nil
	}

	className := tokens[1].Value
	parentClass := "" // For inheritance

	// Check for inheritance syntax: darasa Child : Parent
	startIndex := 2
	if len(tokens) > 3 && tokens[2].Value == ":" {
		parentClass = tokens[3].Value
		startIndex = 4
	}

	// Find the class body
	braceStart := -1
	braceEnd := -1
	braceCount := 0

	for i := startIndex; i < len(tokens); i++ {
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

	// Parse class body (properties and methods)
	var properties []ast.PropertyNode
	var methods []ast.FunctionNode
	var constructor *ast.FunctionNode

	bodyTokens := tokens[braceStart:braceEnd]
	i := 0

	for i < len(bodyTokens) {
		if bodyTokens[i].Value == "" {
			i++
			continue
		}

		// Parse property declarations (type name)
		if (bodyTokens[i].Value == "namba" || bodyTokens[i].Value == "maneno" || bodyTokens[i].Value == "boolean" || bodyTokens[i].Value == "kamusi" || bodyTokens[i].Value == "orodha") && i+1 < len(bodyTokens) {
			propType := bodyTokens[i].Value
			propName := bodyTokens[i+1].Value
			properties = append(properties, ast.PropertyNode{
				Name: propName,
				Type: propType,
			})
			i += 2
			continue
		}

		// Parse method definitions
		if bodyTokens[i].Value == "kazi" {
			// Find the end of this method
			end := i + 1
			methodBraceCount := 0
			foundFirstBrace := false

			for end < len(bodyTokens) {
				if bodyTokens[end].Value == "{" {
					methodBraceCount++
					foundFirstBrace = true
				} else if bodyTokens[end].Value == "}" {
					methodBraceCount--
					if methodBraceCount == 0 && foundFirstBrace {
						end++
						break
					}
				}
				end++
			}

			// Parse this method
			method := ParseFunctionDefinition(bodyTokens[i:end])
			if funcNode, ok := method.(ast.FunctionNode); ok {
				// Check if it's a constructor (named "unda")
				if funcNode.Name == "unda" {
					constructor = &funcNode
				} else {
					methods = append(methods, funcNode)
				}
			}
			i = end
		} else {
			i++
		}
	}

	return ast.ClassNode{
		Name:        className,
		Parent:      parentClass,
		Properties:  properties,
		Methods:     methods,
		Constructor: constructor,
	}
}


// ParseNewInstance parses class instantiation (unda ClassName(args))
func ParseNewInstance(tokens []lexer.Token) ast.ASTNode {
	if len(tokens) < 3 || tokens[0].Value != "unda" {
		return nil
	}

	className := tokens[1].Value

	// Find arguments
	if tokens[2].Value == "(" {
		// Find closing parenthesis
		parenEnd := -1
		for i := 3; i < len(tokens); i++ {
			if tokens[i].Value == ")" {
				parenEnd = i
				break
			}
		}

		var args []ast.ASTNode
		if parenEnd > 3 {
			args = ParseArguments(tokens[3 : parenEnd+1]) // Include closing paren
		}

		return ast.NewInstanceNode{
			ClassName: className,
			Args:      args,
		}
	}

	return nil
}
