package debugger

import (
	"bufio"
	"fmt"
	"kwenda/ast"
	"kwenda/interpreter"
	"kwenda/lexer"
	"kwenda/parser"
	"os"
	"strconv"
	"strings"
)

type Breakpoint struct {
	Line    int
	Enabled bool
}

type Debugger struct {
	sourceLines  []string
	breakpoints  map[int]*Breakpoint
	env          *interpreter.Environment
	currentLine  int
	stepping     bool
	stepOver     bool
	running      bool
	watchVars    []string
}

func NewDebugger(source string, env *interpreter.Environment) *Debugger {
	return &Debugger{
		sourceLines: strings.Split(source, "\n"),
		breakpoints: make(map[int]*Breakpoint),
		env:         env,
		currentLine: 0,
		watchVars:   make([]string, 0),
	}
}

func (d *Debugger) Start() {
	fmt.Println("╔═══════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    KWENDA DEBUGGER                                        ║")
	fmt.Println("║                Interactive Debugging Mode                                 ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Type 'h' or 'help' for debugger commands")
	fmt.Println()

	d.showCurrentLine()
	d.commandLoop()
}

func (d *Debugger) commandLoop() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("(kwenda-dbg) ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := parts[0]

		switch command {
		case "h", "help", "msaada":
			d.printHelp()

		case "l", "list", "orodha":
			d.listSource(parts)

		case "b", "break", "breakpoint":
			d.setBreakpoint(parts)

		case "d", "delete", "futa":
			d.deleteBreakpoint(parts)

		case "lb", "listbreaks":
			d.listBreakpoints()

		case "c", "continue", "endelea":
			d.continueExecution()

		case "s", "step", "hatua":
			d.step()

		case "n", "next", "ifuatayo":
			d.next()

		case "p", "print", "andika":
			d.printVariable(parts)

		case "w", "watch", "angalia":
			d.addWatch(parts)

		case "lw", "listwatch":
			d.listWatches()

		case "dw", "deletewatch":
			d.deleteWatch(parts)

		case "e", "eval", "tathmini":
			d.evalExpression(parts)

		case "v", "vars", "vigezo":
			d.listVariables()

		case "f", "funcs", "kazi":
			d.listFunctions()

		case "q", "quit", "ondoka":
			fmt.Println("Exiting debugger / Kuondoka...")
			os.Exit(0)

		case "r", "run", "endesha":
			d.run()

		default:
			fmt.Printf("Unknown command: %s (Type 'help' for commands)\n", command)
		}
	}
}

func (d *Debugger) printHelp() {
	help := `
╔════════════════════════════════════════════════════════════════════════╗
║                      DEBUGGER COMMANDS                                 ║
╚════════════════════════════════════════════════════════════════════════╝

EXECUTION:
  r, run              Run the program
  c, continue         Continue execution until next breakpoint
  s, step             Step into (execute one line)
  n, next             Step over (execute next line)
  q, quit             Exit debugger

BREAKPOINTS:
  b <line>            Set breakpoint at line number
  b <func>            Set breakpoint at function
  d <line>            Delete breakpoint at line
  lb                  List all breakpoints

SOURCE:
  l                   List current source lines
  l <start> <end>     List source lines from start to end
  l <line>            List source around line

INSPECTION:
  p <var>             Print variable value
  v, vars             List all variables
  f, funcs            List all functions
  e <expr>            Evaluate expression

WATCHES:
  w <var>             Add variable to watch list
  lw                  List watched variables
  dw <var>            Remove variable from watch list

ALIASES (Swahili):
  msaada = help       endelea = continue      hatua = step
  ifuatayo = next     orodha = list           andika = print
  vigezo = vars       kazi = funcs            ondoka = quit
  angalia = watch     tathmini = eval         futa = delete

EXAMPLES:
  b 10                Set breakpoint at line 10
  s                   Step one line
  p myvar             Print value of myvar
  w counter           Watch variable 'counter'
  e x + y             Evaluate expression x + y

`
	fmt.Println(help)
}

func (d *Debugger) showCurrentLine() {
	if d.currentLine >= 0 && d.currentLine < len(d.sourceLines) {
		fmt.Printf("\n--> %4d: %s\n\n", d.currentLine+1, d.sourceLines[d.currentLine])
	}
}

func (d *Debugger) listSource(parts []string) {
	start := 0
	end := len(d.sourceLines)

	if len(parts) == 2 {
		// List around a specific line
		line, err := strconv.Atoi(parts[1])
		if err == nil {
			start = max(0, line-6)
			end = min(len(d.sourceLines), line+5)
		}
	} else if len(parts) == 3 {
		// List range
		s, err1 := strconv.Atoi(parts[1])
		e, err2 := strconv.Atoi(parts[2])
		if err1 == nil && err2 == nil {
			start = max(0, s-1)
			end = min(len(d.sourceLines), e)
		}
	} else {
		// List around current line
		start = max(0, d.currentLine-5)
		end = min(len(d.sourceLines), d.currentLine+10)
	}

	fmt.Println("\nSource Code:")
	fmt.Println("════════════")
	for i := start; i < end; i++ {
		marker := "    "
		if i == d.currentLine {
			marker = "-->"
		} else if _, exists := d.breakpoints[i+1]; exists {
			marker = " B "
		}
		fmt.Printf("%s %4d: %s\n", marker, i+1, d.sourceLines[i])
	}
	fmt.Println()
}

func (d *Debugger) setBreakpoint(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Usage: b <line_number>")
		return
	}

	line, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("Invalid line number: %s\n", parts[1])
		return
	}

	if line < 1 || line > len(d.sourceLines) {
		fmt.Printf("Line %d is out of range (1-%d)\n", line, len(d.sourceLines))
		return
	}

	d.breakpoints[line] = &Breakpoint{Line: line, Enabled: true}
	fmt.Printf("Breakpoint set at line %d\n", line)
}

func (d *Debugger) deleteBreakpoint(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Usage: d <line_number>")
		return
	}

	line, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("Invalid line number: %s\n", parts[1])
		return
	}

	if _, exists := d.breakpoints[line]; exists {
		delete(d.breakpoints, line)
		fmt.Printf("Breakpoint deleted at line %d\n", line)
	} else {
		fmt.Printf("No breakpoint at line %d\n", line)
	}
}

func (d *Debugger) listBreakpoints() {
	if len(d.breakpoints) == 0 {
		fmt.Println("No breakpoints set")
		return
	}

	fmt.Println("\nBreakpoints:")
	fmt.Println("════════════")
	for line, bp := range d.breakpoints {
		status := "enabled"
		if !bp.Enabled {
			status = "disabled"
		}
		fmt.Printf("  Line %d: %s\n", line, status)
	}
	fmt.Println()
}

func (d *Debugger) printVariable(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Usage: p <variable_name>")
		return
	}

	varName := parts[1]
	value := d.env.Get(varName)
	
	if value == nil {
		fmt.Printf("Variable '%s' not found\n", varName)
		return
	}

	d.printValue(varName, value)
}

func (d *Debugger) printValue(name string, value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("%s = \"%s\"\n", name, v)
	case map[string]interface{}:
		fmt.Printf("%s = {", name)
		first := true
		for key, val := range v {
			if !first {
				fmt.Print(", ")
			}
			fmt.Printf("%q: %v", key, val)
			first = false
		}
		fmt.Println("}")
	case []interface{}:
		fmt.Printf("%s = [", name)
		for i, elem := range v {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(elem)
		}
		fmt.Println("]")
	case ast.EnumValueNode:
		fmt.Printf("%s = %s.%s\n", name, v.EnumName, v.Value)
	default:
		fmt.Printf("%s = %v\n", name, v)
	}
}

func (d *Debugger) listVariables() {
	if len(d.env.Variables) == 0 {
		fmt.Println("No variables defined")
		return
	}

	fmt.Println("\nVariables:")
	fmt.Println("══════════")
	for name, value := range d.env.Variables {
		d.printValue(name, value)
	}
	fmt.Println()
}

func (d *Debugger) listFunctions() {
	if len(d.env.Functions) == 0 {
		fmt.Println("No functions defined")
		return
	}

	fmt.Println("\nFunctions:")
	fmt.Println("══════════")
	for name, fn := range d.env.Functions {
		fmt.Printf("  %s(", name)
		for i, param := range fn.Parameters {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s: %s", param.Name, param.Type)
		}
		fmt.Println(")")
	}
	fmt.Println()
}

func (d *Debugger) addWatch(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Usage: w <variable_name>")
		return
	}

	varName := parts[1]
	d.watchVars = append(d.watchVars, varName)
	fmt.Printf("Added watch for '%s'\n", varName)
}

func (d *Debugger) listWatches() {
	if len(d.watchVars) == 0 {
		fmt.Println("No watches set")
		return
	}

	fmt.Println("\nWatched Variables:")
	fmt.Println("══════════════════")
	for _, varName := range d.watchVars {
		value := d.env.Get(varName)
		if value != nil {
			d.printValue(varName, value)
		} else {
			fmt.Printf("%s = <not defined>\n", varName)
		}
	}
	fmt.Println()
}

func (d *Debugger) deleteWatch(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Usage: dw <variable_name>")
		return
	}

	varName := parts[1]
	for i, v := range d.watchVars {
		if v == varName {
			d.watchVars = append(d.watchVars[:i], d.watchVars[i+1:]...)
			fmt.Printf("Removed watch for '%s'\n", varName)
			return
		}
	}
	fmt.Printf("No watch found for '%s'\n", varName)
}

func (d *Debugger) evalExpression(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Usage: e <expression>")
		return
	}

	exprString := strings.Join(parts[1:], " ")
	tokens := lexer.Lex(exprString)
	expr := parser.ParseExpression(tokens)
	
	if expr != nil {
		result := interpreter.Interpret(expr, d.env)
		d.printValue("result", result)
	} else {
		fmt.Println("Error parsing expression")
	}
}

func (d *Debugger) step() {
	fmt.Println("Step execution not fully implemented (requires line tracking)")
	fmt.Println("Use 'run' to execute the program")
}

func (d *Debugger) next() {
	fmt.Println("Next execution not fully implemented (requires line tracking)")
	fmt.Println("Use 'run' to execute the program")
}

func (d *Debugger) continueExecution() {
	fmt.Println("Continue execution not fully implemented")
	fmt.Println("Use 'run' to execute the program")
}

func (d *Debugger) run() {
	fmt.Println("Running program...")
	fmt.Println("(Full execution - breakpoint support requires AST instrumentation)")
	fmt.Println()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
