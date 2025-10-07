package ast

// ASTNode represents a node in the Abstract Syntax Tree
type ASTNode interface{}

// NumberNode represents a numeric literal
type NumberNode struct {
    Value string
}

// IdentifierNode represents a variable or function name
type IdentifierNode struct {
    Value string
}

// BinaryOpNode represents a binary operation (e.g., a + b)
type BinaryOpNode struct {
    Left  ASTNode
    Op    string
    Right ASTNode
}

// ReturnNode represents a return statement
type ReturnNode struct {
    Value ASTNode
}

// InputNode represents a user input operation
type InputNode struct {
    Prompt string // Optional prompt message
}

// FunctionCallNode represents a function call (e.g., andika(x, y))
type FunctionCallNode struct {
    Name string   // Function name
    Args []ASTNode // Function arguments
}

// VariableDeclarationNode represents a variable declaration (e.g., namba x = 10)
type VariableDeclarationNode struct {
    Name  string // Variable name
    Value ASTNode // Variable value
}

// Parameter represents a function parameter
type Parameter struct {
    Name string // Parameter name
    Type string // Parameter type (namba, boolean)
}

// FunctionNode represents a function definition (e.g., kazi kuu() { ... })
type FunctionNode struct {
    Name       string      // Function name
    Parameters []Parameter // Function parameters
    ReturnType string      // Return type (optional)
    Body       []ASTNode   // Function body
}

// IfNode represents a conditional statement (e.g., kama x > 5 { ... } sivyo { ... })
type IfNode struct {
    Condition ASTNode   // The condition to evaluate
    ThenBody  []ASTNode // Statements to execute if condition is true
    ElseBody  []ASTNode // Statements to execute if condition is false (optional)
}

// WhileNode represents a while loop (e.g., wakati x < 10 { ... })
type WhileNode struct {
    Condition ASTNode   // The condition to evaluate
    Body      []ASTNode // Statements to execute while condition is true
}

// ForNode represents a for loop (e.g., kwa i = 0; i < 10; i = i + 1 { ... })
type ForNode struct {
    Init      ASTNode   // Initialization statement (e.g., i = 0)
    Condition ASTNode   // Loop condition (e.g., i < 10)
    Update    ASTNode   // Update statement (e.g., i = i + 1)
    Body      []ASTNode // Statements to execute in each iteration
}

// BreakNode represents a break statement (vunja)
type BreakNode struct {
    // No additional fields needed
}

// ContinueNode represents a continue statement (endelea)
type ContinueNode struct {
    // No additional fields needed
}

// BooleanNode represents a boolean literal (kweli/uwongo)
type BooleanNode struct {
    Value bool // true for kweli, false for uwongo
}

// StringNode represents a string literal
type StringNode struct {
    Value string // The string value without quotes
}

// StringVariableDeclarationNode represents a string variable declaration (e.g., maneno x = "habari")
type StringVariableDeclarationNode struct {
    Name  string  // Variable name
    Value ASTNode // Variable value
}

// ArrayNode represents an array literal (e.g., [1, 2, 3])
type ArrayNode struct {
    Elements []ASTNode // Array elements
}

// ArrayDeclarationNode represents an array variable declaration (e.g., orodha namba = [1, 2, 3])
type ArrayDeclarationNode struct {
    Name     string  // Variable name
    Type     string  // Element type (namba, maneno, etc.)
    Elements []ASTNode // Initial elements
}

// ArrayAccessNode represents array element access (e.g., arr[0])
type ArrayAccessNode struct {
    Array ASTNode // The array being accessed
    Index ASTNode // The index expression
}

// ArrayAssignmentNode represents array element assignment (e.g., arr[0] = 5)
type ArrayAssignmentNode struct {
    Array ASTNode // The array being modified
    Index ASTNode // The index expression
    Value ASTNode // The new value
}

// FileReadNode represents reading from a file (e.g., soma("file.txt"))
type FileReadNode struct {
    Filename ASTNode // The filename to read from
}

// FileWriteNode represents writing to a file (e.g., andika_faili("file.txt", "content"))
type FileWriteNode struct {
    Filename ASTNode // The filename to write to
    Content  ASTNode // The content to write
    Append   bool    // Whether to append or overwrite
}

// TryNode represents a try-catch block (e.g., jaribu { ... } shika (hitilafu) { ... })
type TryNode struct {
    TryBody     []ASTNode // Statements to try executing
    CatchVar    string    // Variable name for the caught error
    CatchBody   []ASTNode // Statements to execute if error occurs
    FinallyBody []ASTNode // Statements to execute regardless (optional)
}

// ThrowNode represents throwing an error (e.g., tupa "Error message")
type ThrowNode struct {
    Message ASTNode // The error message to throw
}