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
)

type Token struct {
	Type  TokenType
	Value string
}

func isSwahiliKeyword(word string) bool {
	keywords := []string{"kazi", "kama", "kwa", "rudisha", "namba", "andika"}
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

	for _, char := range input {
		if unicode.IsSpace(char) {
			// End of current token
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
		} else if char == '+' || char == '-' || char == '*' || char == '/' {
			// Handle operators
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
		} else if char == '{' || char == '}' || char == '(' || char == ')' || char == '=' || char == ';' || char == ',' {
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
		} else {
			tokens = append(tokens, Token{Type: TokenIdentifier, Value: tokenValue})
		}
	}

	return tokens
}