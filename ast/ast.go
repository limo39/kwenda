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

// FunctionNode represents a function definition (e.g., kazi kuu() { ... })
type FunctionNode struct {
    Name string   // Function name
    Body []ASTNode // Function body
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