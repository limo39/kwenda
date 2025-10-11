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
		"kazi", "kama", "sivyo", "kwa", "wakati", "rudisha", "namba", "andika", "ingiza",
		"kweli", "uwongo", "na", "au", "vunja", "endelea", "boolean", "maneno",
		// Array keywords
		"orodha", "ongeza", "ondoa", "urefu_orodha", "pata",
		// File I/O keywords
		"soma", "andika_faili", "unda_faili", "faili_ipo", "ondoa_faili",
		// Import/Module keywords
		"leta", "kutoka", "moduli", "umma",
		// Error handling keywords
		"jaribu", "shika", "hatimaye", "tupa",
		// String manipulation functions
		"unganisha", "kata", "badilisha", "tafuta", "awali", "mwisho",
		"herufi_kubwa", "herufi_ndogo", "ondoa_nafasi", "gawanya_maneno",
	}
	for _, kw := range keywords {
		if kw == word {
			return true
		}
	}
	return false
}

func isNumber(word string) bool {
	if len(word) == 0 {
		return false
	}
	hasDecimal := false
	for i, ch := range word {
		if ch == '.' {
			if hasDecimal || i == 0 || i == len(word)-1 {
				return false // Multiple decimals, or decimal at start/end
			}
			hasDecimal = true
		} else if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}

func Lex(input string) []Token {
	var tokens []Token
	var currentToken strings.Builder
	var inString bool
	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		char := runes[i]
		
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
				} else if isNumber(tokenValue) {
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
				} else if isNumber(tokenValue) {
					tokens = append(tokens, Token{Type: TokenNumber, Value: tokenValue})
				} else {
					tokens = append(tokens, Token{Type: TokenIdentifier, Value: tokenValue})
				}
				currentToken.Reset()
			}
			
			// Handle multi-character operators like ==, !=, <=, >=
			if i+1 < len(runes) && runes[i+1] == '=' {
				if char == '=' {
					tokens = append(tokens, Token{Type: TokenOperator, Value: "=="})
				} else if char == '!' {
					tokens = append(tokens, Token{Type: TokenOperator, Value: "!="})
				} else if char == '<' {
					tokens = append(tokens, Token{Type: TokenOperator, Value: "<="})
				} else if char == '>' {
					tokens = append(tokens, Token{Type: TokenOperator, Value: ">="})
				}
				// Skip the next character since we consumed it
				i++
			} else {
				tokens = append(tokens, Token{Type: TokenOperator, Value: string(char)})
			}
		} else if char == '{' || char == '}' || char == '(' || char == ')' || char == '[' || char == ']' || char == ';' || char == ',' {
			// Handle punctuation
			if currentToken.Len() > 0 {
				tokenValue := currentToken.String()
				if isSwahiliKeyword(tokenValue) {
					tokens = append(tokens, Token{Type: TokenKeyword, Value: tokenValue})
				} else if isNumber(tokenValue) {
					tokens = append(tokens, Token{Type: TokenNumber, Value: tokenValue})
				} else {
					tokens = append(tokens, Token{Type: TokenIdentifier, Value: tokenValue})
				}
				currentToken.Reset()
			}
			tokens = append(tokens, Token{Type: TokenPunctuation, Value: string(char)})
		} else if char == '#' {
			// Handle comments (ignore the rest of the line)
			// First, finalize any current token
			if currentToken.Len() > 0 {
				tokenValue := currentToken.String()
				if isSwahiliKeyword(tokenValue) {
					tokens = append(tokens, Token{Type: TokenKeyword, Value: tokenValue})
				} else if isNumber(tokenValue) {
					tokens = append(tokens, Token{Type: TokenNumber, Value: tokenValue})
				} else {
					tokens = append(tokens, Token{Type: TokenIdentifier, Value: tokenValue})
				}
				currentToken.Reset()
			}
			// Skip the rest of the line
			for i+1 < len(runes) && runes[i+1] != '\n' {
				i++
			}
		} else if char == '.' {
			// Handle decimal point - could be part of a number or module access
			if currentToken.Len() > 0 {
				tokenValue := currentToken.String()
				// If current token is a number and next char is a digit, it's a decimal
				if unicode.IsDigit(rune(tokenValue[0])) && i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
					currentToken.WriteRune(char)
				} else {
					// Otherwise, end current token and treat . as punctuation
					if isSwahiliKeyword(tokenValue) {
						tokens = append(tokens, Token{Type: TokenKeyword, Value: tokenValue})
					} else if isNumber(tokenValue) {
						tokens = append(tokens, Token{Type: TokenNumber, Value: tokenValue})
					} else {
						tokens = append(tokens, Token{Type: TokenIdentifier, Value: tokenValue})
					}
					currentToken.Reset()
					currentToken.WriteRune(char)
				}
			} else {
				currentToken.WriteRune(char)
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
		} else if isNumber(tokenValue) {
			tokens = append(tokens, Token{Type: TokenNumber, Value: tokenValue})
		} else if tokenValue == "kweli" || tokenValue == "uwongo" {
			tokens = append(tokens, Token{Type: TokenBoolean, Value: tokenValue})
		} else {
			tokens = append(tokens, Token{Type: TokenIdentifier, Value: tokenValue})
		}
	}

	return tokens
}