package parser

import (
	"lugha-yangu/ast"
	"lugha-yangu/lexer"
)

// Parse converts tokens into an Abstract Syntax Tree (AST)
func Parse(tokens []lexer.Token) ast.ASTNode {
	if len(tokens) == 0 {
		return nil
	}

	// Handle conditional statements (e.g., kama x > 5 { ... } sivyo { ... })
	if tokens[0].Value == "kama" {
		return ParseIfStatement(tokens)
	}

	// Handle function definitions (e.g., kazi kuu() { ... })
	if tokens[0].Value == "kazi" && len(tokens) >= 4 && tokens[2].Value == "(" && tokens[3].Value == ")" {
		// Extract the function name
		functionName := tokens[1].Value

		// Find the opening and closing braces for the function body
		bodyStart := -1
		bodyEnd := -1
		braceCount := 0

		for i, token := range tokens {
			if token.Value == "{" {
				if braceCount == 0 {
					bodyStart = i + 1
				}
				braceCount++
			} else if token.Value == "}" {
				braceCount--
				if braceCount == 0 {
					bodyEnd = i
					break
				}
			}
		}

		if bodyStart == -1 || bodyEnd == -1 {
			return nil // Invalid function body
		}

		// Parse the function body
		body := ParseBlock(tokens[bodyStart:bodyEnd])

		return ast.FunctionNode{
			Name: functionName,
			Body: body,
		}
	}

	// Handle variable declarations (e.g., namba x = 10)
	if tokens[0].Value == "namba" && len(tokens) >= 4 && tokens[2].Value == "=" {
		return ast.VariableDeclarationNode{
			Name:  tokens[1].Value,
			Value: ParseExpression(tokens[3:]),
		}
	}

	// Handle user input (ingiza)
	if tokens[0].Value == "ingiza" {
		if len(tokens) > 1 && tokens[1].Type == lexer.TokenString {
			return ast.InputNode{Prompt: tokens[1].Value}
		}
		return ast.InputNode{}
	}

	// Handle function calls (e.g., andika(x, y))
	if tokens[0].Type == lexer.TokenIdentifier && len(tokens) > 1 && tokens[1].Value == "(" {
		return ast.FunctionCallNode{
			Name: tokens[0].Value,
			Args: ParseArguments(tokens[2:]),
		}
	}
	
	// Handle function calls with keyword (e.g., andika(x, y))
	if tokens[0].Value == "andika" && len(tokens) > 1 && tokens[1].Value == "(" {
		return ast.FunctionCallNode{
			Name: tokens[0].Value,
			Args: ParseArguments(tokens[2:]),
		}
	}

	// Handle binary operations (e.g., a + b)
	if len(tokens) >= 3 && tokens[1].Type == lexer.TokenOperator {
		return ast.BinaryOpNode{
			Left:  ParseExpression(tokens[:1]),
			Op:    tokens[1].Value,
			Right: ParseExpression(tokens[2:]),
		}
	}

	// Handle numbers
	if tokens[0].Type == lexer.TokenNumber {
		return ast.NumberNode{Value: tokens[0].Value}
	}

	// Handle identifiers
	if tokens[0].Type == lexer.TokenIdentifier {
		return ast.IdentifierNode{Value: tokens[0].Value}
	}

	// Default case: return nil for unknown nodes
	return nil
}

// ParseBlock parses a block of statements (e.g., inside a function body)
func ParseBlock(tokens []lexer.Token) []ast.ASTNode {
	var statements []ast.ASTNode
	i := 0

	for i < len(tokens) {
		// Skip any whitespace or empty tokens
		if tokens[i].Value == "" {
			i++
			continue
		}
		


		// Parse conditional statements (kama ...)
		if tokens[i].Value == "kama" {
			// Find the end of the if statement (including else if present)
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
						// Check if there's a "sivyo" after this closing brace
						if end+1 < len(tokens) && tokens[end+1].Value == "sivyo" {
							// Find the end of the else block
							end += 2 // Skip "sivyo"
							// Look for opening brace of else block
							for end < len(tokens) && tokens[end].Value != "{" {
								end++
							}
							if end < len(tokens) {
								end++ // Skip opening brace
								braceCount = 1
								for end < len(tokens) && braceCount > 0 {
									if tokens[end].Value == "{" {
										braceCount++
									} else if tokens[end].Value == "}" {
										braceCount--
									}
									end++
								}
							}
							break
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

		// Parse variable declarations (namba x = ...)
		if tokens[i].Value == "namba" && i+3 < len(tokens) && tokens[i+2].Value == "=" {
			// Find the end of this statement (look for next keyword or end of tokens)
			end := i + 3
			parenCount := 0
			for end < len(tokens) {
				if tokens[end].Value == "(" {
					parenCount++
				} else if tokens[end].Value == ")" {
					parenCount--
				} else if parenCount == 0 && (tokens[end].Value == "namba" || tokens[end].Value == "andika" || tokens[end].Value == "kama") {
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

		// Parse function calls (andika(...))
		if tokens[i].Value == "andika" && i+1 < len(tokens) && tokens[i+1].Value == "(" {
			// Find the matching closing parenthesis
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

		// If we reach here, we didn't match any specific pattern
		// Try to parse as a general statement
		stmt := Parse(tokens[i:i+1])
		if stmt != nil {
			statements = append(statements, stmt)
		}
		i++
	}

	return statements
}

// ParseExpression parses an expression (e.g., a + b, 42, x)
func ParseExpression(tokens []lexer.Token) ast.ASTNode {
	if len(tokens) == 0 {
		return nil
	}
	
	// Handle function calls like ingiza("prompt")
	if len(tokens) >= 3 && tokens[0].Value == "ingiza" && tokens[1].Value == "(" {
		// Find the closing parenthesis
		for i := 2; i < len(tokens); i++ {
			if tokens[i].Value == ")" {
				if i > 2 && tokens[2].Type == lexer.TokenString {
					return ast.InputNode{Prompt: tokens[2].Value}
				}
				return ast.InputNode{}
			}
		}
	}
	
	// Handle binary operations (e.g., x + y, x > 5, x == 10)
	if len(tokens) >= 3 && tokens[1].Type == lexer.TokenOperator {
		return ast.BinaryOpNode{
			Left:  ParseExpression(tokens[:1]),
			Op:    tokens[1].Value,
			Right: ParseExpression(tokens[2:3]), // Take one token for the right operand
		}
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

// ParseArguments parses function arguments (e.g., x, y)
func ParseArguments(tokens []lexer.Token) []ast.ASTNode {
	var args []ast.ASTNode
	var currentArg []lexer.Token

	for _, token := range tokens {
		if token.Value == "," || token.Value == ")" {
			if len(currentArg) > 0 {
				// Handle string literals directly
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