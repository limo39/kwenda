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