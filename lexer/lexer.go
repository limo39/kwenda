package lexer

import (
	"strings"
	"unicode"
)

type TokenType string

const (
	TokenIdentifier TokenType = "IDENTIFIER"
	TokenNumber     TokenType = "NUMBER"
	TokenKeyword    TokenType = "KEYWORD"
	TokenOperator   TokenType = "OPERATOR"
	TokenPunctuation TokenType = "PUNCTUATION"
	TokenString     TokenType = "STRING"
	TokenBoolean    TokenType = "BOOLEAN"
)

type Token struct {
	Type  TokenType
	Value string
}

func isSwahiliKeyword(word string) bool {
	keywords := []string{
		"kazi", "kama", "sivyo", "kwa", "rudisha", "namba", "andika", "ingiza",
		"kweli", "uwongo", "na", "au",
	}
	for _, kw := range keywords {
		if kw == word {
			return true
		}
	}
	return false
}

func Lex(input string) []Token {
	var tokens []Token
	var currentToken strings.Builder
	var inString bool

	for _, char := range input {
		if char == '"' {
			// Handle string literals
			if inString {
				tokens = append(tokens, Token{Type: TokenString, Value: currentToken.String()})
				currentToken.Reset()
				inString = false
			} else {
				inString = true
			}
		} else if inString {
			currentToken.WriteRune(char)
		} else if unicode.IsSpace(char) {
			// End of current token
			if currentToken.Len() > 0 {
				tokenValue := currentToken.String()
				if isSwahiliKeyword(tokenValue) {
					tokens = append(tokens, Token{Type: TokenKeyword, Value: tokenValue})
				} else if unicode.IsDigit(rune(tokenValue[0])) {
					tokens = append(tokens, Token{Type: TokenNumber, Value: tokenValue})
				} else if tokenValue == "kweli" || tokenValue == "uwongo" {
					tokens = append(tokens, Token{Type: TokenBoolean, Value: tokenValue})
				} else {
					tokens = append(tokens, Token{Type: TokenIdentifier, Value: tokenValue})
				}
				currentToken.Reset()
			}
		} else if char == '+' || char == '-' || char == '*' || char == '/' || char == '=' || char == '!' || char == '<' || char == '>' {
			// Handle operators and comparisons
			if currentToken.Len() > 0 {
				tokenValue := currentToken.String()
				if isSwahiliKeyword(tokenValue) {
					tokens = append(tokens, Token{Type: TokenKeyword, Value: tokenValue})
				} else if unicode.IsDigit(rune(tokenValue[0])) {
					tokens = append(tokens, Token{Type: TokenNumber, Value: tokenValue})
				} else {
					tokens = append(tokens, Token{Type: TokenIdentifier, Value: tokenValue})
				}
				currentToken.Reset()
			}
			tokens = append(tokens, Token{Type: TokenOperator, Value: string(char)})
		} else if char == '{' || char == '}' || char == '(' || char == ')' || char == ';' || char == ',' {
			// Handle punctuation
			if currentToken.Len() > 0 {
				tokenValue := currentToken.String()
				if isSwahiliKeyword(tokenValue) {
					tokens = append(tokens, Token{Type: TokenKeyword, Value: tokenValue})
				} else if unicode.IsDigit(rune(tokenValue[0])) {
					tokens = append(tokens, Token{Type: TokenNumber, Value: tokenValue})
				} else {
					tokens = append(tokens, Token{Type: TokenIdentifier, Value: tokenValue})
				}
				currentToken.Reset()
			}
			tokens = append(tokens, Token{Type: TokenPunctuation, Value: string(char)})
		} else if char == '#' {
			// Handle comments (ignore the rest of the line)
			for char != '\n' {
				break
			}
		} else {
			// Build the current token
			currentToken.WriteRune(char)
		}
	}

	// Handle the last token if any
	if currentToken.Len() > 0 {
		tokenValue := currentToken.String()
		if isSwahiliKeyword(tokenValue) {
			tokens = append(tokens, Token{Type: TokenKeyword, Value: tokenValue})
		} else if unicode.IsDigit(rune(tokenValue[0])) {
			tokens = append(tokens, Token{Type: TokenNumber, Value: tokenValue})
		} else if tokenValue == "kweli" || tokenValue == "uwongo" {
			tokens = append(tokens, Token{Type: TokenBoolean, Value: tokenValue})
		} else {
			tokens = append(tokens, Token{Type: TokenIdentifier, Value: tokenValue})
		}
	}

	return tokens
}