package repl

import (
	"bufio"
	"fmt"
	"kwenda/ast"
	"kwenda/interpreter"
	"kwenda/lexer"
	"kwenda/parser"
	"os"
	"strings"
)

const PROMPT = "kwenda> "
const CONTINUATION_PROMPT = "....... "

type REPL struct {
	env          *interpreter.Environment
	multilineBuffer strings.Builder
	inMultiline  bool
	braceCount   int
	history      []string
}

func NewREPL() *REPL {
	return &REPL{
		env:     interpreter.NewEnvironment(),
		history: make([]string, 0),
	}
}

func (r *REPL) Start() {
	fmt.Println("╔═══════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║              KWENDA REPL - Interactive Swahili Programming                ║")
	fmt.Println("║                       Version 1.0.0                                       ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Type 'ondoka' or 'exit' to quit")
	fmt.Println("Type 'msaada' or 'help' for help")
	fmt.Println("Type 'historia' or 'history' to see command history")
	fmt.Println("Type 'safisha' or 'clear' to clear variables")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Show prompt
		if r.inMultiline {
			fmt.Print(CONTINUATION_PROMPT)
		} else {
			fmt.Print(PROMPT)
		}

		// Read input
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		// Handle multiline input
		if r.inMultiline {
			r.multilineBuffer.WriteString(line + "\n")
			
			// Update brace count
			r.braceCount += strings.Count(line, "{")
			r.braceCount -= strings.Count(line, "}")
			
			// Check if multiline is complete
			if r.braceCount <= 0 {
				r.inMultiline = false
				completeInput := r.multilineBuffer.String()
				r.multilineBuffer.Reset()
				r.processInput(completeInput)
			}
			continue
		}

		// Check for special commands
		if r.handleCommand(trimmedLine) {
			continue
		}

		// Check if this starts a multiline block
		braceCount := strings.Count(line, "{") - strings.Count(line, "}")
		if braceCount > 0 {
			r.inMultiline = true
			r.braceCount = braceCount
			r.multilineBuffer.WriteString(line + "\n")
			continue
		}

		// Process single line
		r.processInput(line)
	}
}

func (r *REPL) handleCommand(input string) bool {
	switch input {
	case "ondoka", "exit":
		fmt.Println("Kwaheri! (Goodbye!)")
		os.Exit(0)
		return true
		
	case "msaada", "help":
		r.printHelp()
		return true
		
	case "historia", "history":
		r.printHistory()
		return true
		
	case "safisha", "clear":
		r.env = interpreter.NewEnvironment()
		fmt.Println("Variables cleared / Vigezo vimefutwa")
		return true
		
	case "":
		return true
		
	default:
		return false
	}
}

func (r *REPL) processInput(input string) {
	// Add to history
	r.history = append(r.history, input)

	// Lex the input
	tokens := lexer.Lex(input)
	
	// Check if it's an expression or statement
	if r.isExpression(tokens) {
		// Parse as expression and evaluate
		expr := parser.ParseExpression(tokens)
		if expr != nil {
			result := interpreter.Interpret(expr, r.env)
			r.printResult(result)
		} else {
			fmt.Println("Hitilafu ya kusoma / Parse error")
		}
	} else {
		// Parse as statement
		program := parser.ParseProgram(tokens)
		
		// Execute each statement
		for _, stmt := range program.Functions {
			result := interpreter.Interpret(stmt, r.env)
			
			// Only print result if it's not nil and not a control flow result
			if result != nil {
				if cfResult, ok := result.(interpreter.ControlFlowResult); ok {
					if cfResult.Type == interpreter.ControlThrow {
						if errVal, ok := cfResult.Value.(interpreter.ErrorValue); ok {
							fmt.Printf("Hitilafu / Error: %s\n", errVal.Message)
						}
					}
				}
			}
		}
	}
}

func (r *REPL) isExpression(tokens []lexer.Token) bool {
	// Simple heuristic: if it starts with a keyword like kazi, kama, etc., it's a statement
	// Otherwise, treat as expression
	if len(tokens) == 0 {
		return false
	}
	
	firstToken := tokens[0]
	keywords := []string{"kazi", "kama", "wakati", "kwa", "darasa", "rudisha", 
	                     "namba", "maneno", "boolean", "kamusi", "orodha", "aina"}
	
	for _, keyword := range keywords {
		if firstToken.Value == keyword {
			return false
		}
	}
	
	return true
}

func (r *REPL) printResult(result interface{}) {
	switch v := result.(type) {
	case nil:
		// Don't print nil results
	case string:
		fmt.Printf("\"%s\"\n", v)
	case map[string]interface{}:
		fmt.Print("{")
		first := true
		for key, value := range v {
			if !first {
				fmt.Print(", ")
			}
			fmt.Printf("%q: %v", key, value)
			first = false
		}
		fmt.Println("}")
	case []interface{}:
		fmt.Print("[")
		for i, elem := range v {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(elem)
		}
		fmt.Println("]")
	case ast.EnumValueNode:
		fmt.Printf("%s.%s\n", v.EnumName, v.Value)
	default:
		fmt.Println(v)
	}
}

func (r *REPL) printHelp() {
	help := `
╔═════════════════════════════════════════════════════════════════════════╗
║                        KWENDA REPL HELP                                 ║
╚═════════════════════════════════════════════════════════════════════════╝

COMMANDS:
  ondoka, exit          Exit the REPL
  msaada, help          Show this help message
  historia, history     Show command history
  safisha, clear        Clear all variables

USAGE:
  - Type any expression to evaluate it
  - Type statements to execute them
  - Multiline input is supported (use { } to start)
  - Variables persist between commands

EXAMPLES:
  kwenda> 2 + 2
  4

  kwenda> namba x = 10
  kwenda> x * 5
  50

  kwenda> kazi greet(name) {
  ....... andika("Habari, ", name)
  ....... }
  kwenda> greet("Dunia")
  Habari, Dunia

  kwenda> aina Status { PENDING, ACTIVE }
  kwenda> s = Status.ACTIVE
  kwenda> s
  Status.ACTIVE

FEATURES:
  ✓ Immediate expression evaluation
  ✓ Statement execution
  ✓ Persistent variables and functions
  ✓ Multiline input support
  ✓ Command history
  ✓ All Kwenda language features available

`
	fmt.Println(help)
}

func (r *REPL) printHistory() {
	if len(r.history) == 0 {
		fmt.Println("No history / Hakuna historia")
		return
	}
	
	fmt.Println("\nCommand History / Historia ya Amri:")
	fmt.Println("════════════════════════════════════")
	for i, cmd := range r.history {
		fmt.Printf("%3d: %s\n", i+1, cmd)
	}
	fmt.Println()
}
