package parser

import (
	"lugha-yangu/ast"
	"lugha-yangu/lexer"
)

// ProgramNode represents the entire program
type ProgramNode struct {
	Functions []ast.ASTNode
}

// ParseProgram parses the entire program with multiple functions
func ParseProgram(tokens []lexer.Token) ProgramNode {
	var functions []ast.ASTNode
	i := 0

	for i < len(tokens) {
		// Skip empty tokens
		if tokens[i].Value == "" {
			i++
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

	return ProgramNode{Functions: functions}
}

// Parse converts tokens into an Abstract Syntax Tree (AST)
func Parse(tokens []lexer.Token) ast.ASTNode {
	if len(tokens) == 0 {
		return nil
	}

	// Handle conditional statements (e.g., kama x > 5 { ... } sivyo { ... })
	if tokens[0].Value == "kama" {
		return ParseIfStatement(tokens)
	}

	// Handle while loops (e.g., wakati x < 10 { ... })
	if tokens[0].Value == "wakati" {
		return ParseWhileStatement(tokens)
	}

	// Handle for loops (e.g., kwa i = 0; i < 10; i = i + 1 { ... })
	if tokens[0].Value == "kwa" {
		return ParseForStatement(tokens)
	}

	// Handle break statements (vunja)
	if tokens[0].Value == "vunja" {
		return ast.BreakNode{}
	}

	// Handle continue statements (endelea)
	if tokens[0].Value == "endelea" {
		return ast.ContinueNode{}
	}

	// Handle return statements (rudisha)
	if tokens[0].Value == "rudisha" {
		if len(tokens) > 1 {
			return ast.ReturnNode{Value: ParseExpression(tokens[1:])}
		}
		return ast.ReturnNode{Value: nil}
	}

	// Handle function definitions (e.g., kazi kuu() { ... } or kazi hesabu(namba x, namba y) namba { ... })
	if tokens[0].Value == "kazi" && len(tokens) >= 4 && tokens[2].Value == "(" {
		return ParseFunctionDefinition(tokens)
	}

	// Handle variable declarations (e.g., namba x = 10)
	if tokens[0].Value == "namba" && len(tokens) >= 4 && tokens[2].Value == "=" {
		return ast.VariableDeclarationNode{
			Name:  tokens[1].Value,
			Value: ParseExpression(tokens[3:]),
		}
	}

	// Handle boolean variable declarations (e.g., boolean x = kweli)
	if tokens[0].Value == "boolean" && len(tokens) >= 4 && tokens[2].Value == "=" {
		return ast.VariableDeclarationNode{
			Name:  tokens[1].Value,
			Value: ParseExpression(tokens[3:]),
		}
	}

	// Handle assignment statements (e.g., x = 10)
	if len(tokens) >= 3 && tokens[0].Type == lexer.TokenIdentifier && tokens[1].Value == "=" {
		return ast.VariableDeclarationNode{
			Name:  tokens[0].Value,
			Value: ParseExpression(tokens[2:]),
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

	// Handle logical operations (na, au)
	if len(tokens) >= 3 && (tokens[1].Value == "na" || tokens[1].Value == "au") {
		return ast.BinaryOpNode{
			Left:  ParseExpression(tokens[:1]),
			Op:    tokens[1].Value,
			Right: ParseExpression(tokens[2:]),
		}
	}

	// Handle binary operations (e.g., a + b) but NOT assignment
	if len(tokens) >= 3 && tokens[1].Type == lexer.TokenOperator && tokens[1].Value != "=" {
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

	// Handle boolean literals
	if tokens[0].Type == lexer.TokenBoolean || tokens[0].Value == "kweli" || tokens[0].Value == "uwongo" {
		return ast.BooleanNode{Value: tokens[0].Value == "kweli"}
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

		// Parse while loops (wakati ...)
		if tokens[i].Value == "wakati" {
			// Find the end of the while loop
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
			
			stmt := Parse(tokens[i:end])
			if stmt != nil {
				statements = append(statements, stmt)
			}
			i = end
			continue
		}

		// Parse for loops (kwa ...)
		if tokens[i].Value == "kwa" {
			// Find the end of the for loop
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
				} else if parenCount == 0 && (tokens[end].Value == "namba" || tokens[end].Value == "boolean" || tokens[end].Value == "andika" || tokens[end].Value == "kama" || tokens[end].Value == "wakati" || tokens[end].Value == "kwa") {
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

		// Parse boolean variable declarations (boolean x = ...)
		if tokens[i].Value == "boolean" && i+3 < len(tokens) && tokens[i+2].Value == "=" {
			// Find the end of this statement
			end := i + 3
			parenCount := 0
			for end < len(tokens) {
				if tokens[end].Value == "(" {
					parenCount++
				} else if tokens[end].Value == ")" {
					parenCount--
				} else if parenCount == 0 && (tokens[end].Value == "namba" || tokens[end].Value == "boolean" || tokens[end].Value == "andika" || tokens[end].Value == "kama" || tokens[end].Value == "wakati" || tokens[end].Value == "kwa") {
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

		// Parse assignment statements (x = ...)
		if i+2 < len(tokens) && tokens[i].Type == lexer.TokenIdentifier && tokens[i+1].Value == "=" {
			// Find the end of this statement
			end := i + 2
			parenCount := 0
			for end < len(tokens) {
				if tokens[end].Value == "(" {
					parenCount++
				} else if tokens[end].Value == ")" {
					parenCount--
				} else if parenCount == 0 && (tokens[end].Value == "namba" || tokens[end].Value == "boolean" || tokens[end].Value == "andika" || tokens[end].Value == "kama" || tokens[end].Value == "wakati" || tokens[end].Value == "kwa") {
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

		// Parse function calls (andika(...) or any function call)
		if i+1 < len(tokens) && tokens[i+1].Value == "(" {
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

		// Parse break statements (vunja)
		if tokens[i].Value == "vunja" {
			statements = append(statements, ast.BreakNode{})
			i++
			continue
		}

		// Parse continue statements (endelea)
		if tokens[i].Value == "endelea" {
			statements = append(statements, ast.ContinueNode{})
			i++
			continue
		}

		// Parse return statements (rudisha)
		if tokens[i].Value == "rudisha" {
			// Find the end of the return statement
			end := i + 1
			parenCount := 0
			for end < len(tokens) {
				if tokens[end].Value == "(" {
					parenCount++
				} else if tokens[end].Value == ")" {
					parenCount--
				} else if parenCount == 0 && (tokens[end].Value == "namba" || tokens[end].Value == "boolean" || tokens[end].Value == "andika" || tokens[end].Value == "kama" || tokens[end].Value == "wakati" || tokens[end].Value == "kwa" || tokens[end].Value == "rudisha") {
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
	
	// Handle function calls like ingiza("prompt") or any function call
	if len(tokens) >= 3 && tokens[1].Value == "(" {
		if tokens[0].Value == "ingiza" {
			// Special handling for ingiza
			for i := 2; i < len(tokens); i++ {
				if tokens[i].Value == ")" {
					if i > 2 && tokens[2].Type == lexer.TokenString {
						return ast.InputNode{Prompt: tokens[2].Value}
					}
					return ast.InputNode{}
				}
			}
		} else {
			// General function call
			return ast.FunctionCallNode{
				Name: tokens[0].Value,
				Args: ParseArguments(tokens[2:]),
			}
		}
	}
	
	// Handle logical operations (na, au)
	if len(tokens) >= 3 && (tokens[1].Value == "na" || tokens[1].Value == "au") {
		return ast.BinaryOpNode{
			Left:  ParseExpression(tokens[:1]),
			Op:    tokens[1].Value,
			Right: ParseExpression(tokens[2:3]), // Take one token for the right operand
		}
	}

	// Handle binary operations (e.g., x + y, x > 5, x == 10) but NOT assignment
	if len(tokens) >= 3 && tokens[1].Type == lexer.TokenOperator && tokens[1].Value != "=" {
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

	// Handle boolean literals
	if tokens[0].Type == lexer.TokenBoolean || tokens[0].Value == "kweli" || tokens[0].Value == "uwongo" {
		return ast.BooleanNode{Value: tokens[0].Value == "kweli"}
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
	// Look for semicolons to separate the parts
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
		// Simple for loop: kwa condition { ... }
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
		condition = ParseExpression(tokens[initEnd+1:conditionEnd])
	}
	
	if updateEnd > conditionEnd+1 {
		update = ParseExpression(tokens[conditionEnd+1:updateEnd])
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

// ParseFunctionDefinition parses function definitions with parameters and return types
func ParseFunctionDefinition(tokens []lexer.Token) ast.ASTNode {
	if len(tokens) < 4 || tokens[0].Value != "kazi" {
		return nil
	}

	functionName := tokens[1].Value

	// Find the closing parenthesis for parameters
	parenStart := -1
	parenEnd := -1
	for i := 2; i < len(tokens); i++ {
		if tokens[i].Value == "(" {
			parenStart = i
		} else if tokens[i].Value == ")" {
			parenEnd = i
			break
		}
	}

	if parenStart == -1 || parenEnd == -1 {
		return nil
	}

	// Parse parameters
	var parameters []ast.Parameter
	if parenEnd > parenStart+1 {
		parameters = ParseParameters(tokens[parenStart+1:parenEnd])
	}

	// Check for return type
	var returnType string
	bodyStart := parenEnd + 1

	// If there's a type after the closing parenthesis, it's the return type
	if bodyStart < len(tokens) && (tokens[bodyStart].Value == "namba" || tokens[bodyStart].Value == "boolean") {
		returnType = tokens[bodyStart].Value
		bodyStart++
	}

	// Find the function body
	braceStart := -1
	braceEnd := -1
	braceCount := 0

	for i := bodyStart; i < len(tokens); i++ {
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
		ReturnType: returnType,
		Body:       body,
	}
}

// ParseParameters parses function parameters (e.g., namba x, boolean y)
func ParseParameters(tokens []lexer.Token) []ast.Parameter {
	var parameters []ast.Parameter
	var currentParam []lexer.Token

	for _, token := range tokens {
		if token.Value == "," {
			if len(currentParam) >= 2 {
				// Parameter should be: type name
				paramType := currentParam[0].Value
				paramName := currentParam[1].Value
				parameters = append(parameters, ast.Parameter{
					Name: paramName,
					Type: paramType,
				})
			}
			currentParam = nil
		} else {
			currentParam = append(currentParam, token)
		}
	}

	// Handle the last parameter
	if len(currentParam) >= 2 {
		paramType := currentParam[0].Value
		paramName := currentParam[1].Value
		parameters = append(parameters, ast.Parameter{
			Name: paramName,
			Type: paramType,
		})
	}

	return parameters
}