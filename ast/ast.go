package ast

type ASTNode interface{}

type FunctionNode struct {
    Name   string
    Params []string
    Body   []ASTNode
}

type ReturnNode struct {
    Value ASTNode
}

type BinaryOpNode struct {
    Left  ASTNode
    Op    string
    Right ASTNode
}

type NumberNode struct {
    Value string
}

type IdentifierNode struct {
    Value string
}