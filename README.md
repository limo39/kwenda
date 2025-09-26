# Lugha Yangu - A Swahili Programming Language

**Lugha Yangu** (meaning "My Language" in Swahili) is a simple, educational programming language with Swahili syntax. Built in Go, it's designed to make programming more accessible to Swahili speakers by using familiar keywords and concepts.

## 🌟 Features

- **Native Swahili Syntax**: All keywords are in Swahili
- **Interactive Input/Output**: Built-in support for user interaction
- **Arithmetic Operations**: Basic mathematical operations
- **Variable Declarations**: Type-safe variable handling
- **Function Definitions**: Support for custom functions
- **Conditional Statements**: If/else logic with `kama`/`sivyo`
- **Educational Focus**: Perfect for learning programming concepts

## 🚀 Quick Start

### Prerequisites
- Go 1.23.3 or higher

### Installation
```bash
git clone <repository-url>
cd lugha-yangu
go build
```

### Running a Program
```bash
go run main.go
```

The interpreter will execute the program in `examples/example1.swh` by default.

## 📝 Language Syntax

### Keywords
| Swahili | English | Description |
|---------|---------|-------------|
| `kazi` | function | Define a function |
| `kuu` | main | Main function (entry point) |
| `namba` | number | Declare a number variable |
| `ingiza` | input | Get user input |
| `andika` | print | Print output |
| `rudisha` | return | Return a value |
| `kama` | if | Conditional statement |
| `sivyo` | else | Alternative condition |

### Basic Syntax

#### Variable Declaration
```swahili
namba x = 10
namba y = ingiza("Ingiza namba:")
```

#### Function Definition
```swahili
kazi kuu() {
    // Your code here
}
```

#### Input/Output
```swahili
namba x = ingiza("Ingiza namba ya kwanza:")  // Input with prompt
andika("Jibu ni:", x)                        // Print output
```

#### Arithmetic Operations
```swahili
namba jibu = x + y    // Addition
namba tofauti = x - y // Subtraction
namba bidhaa = x * y  // Multiplication
namba mgawanyo = x / y // Division
```

#### Conditional Statements
```swahili
kama x > 10 {
    andika("Kubwa kuliko 10")
} sivyo {
    andika("Ndogo au sawa na 10")
}
```

#### Comparison Operators
```swahili
x == y    // Equal to
x != y    // Not equal to
x < y     // Less than
x <= y    // Less than or equal to
x > y     // Greater than
x >= y    // Greater than or equal to
```

## 📚 Examples

### Example 1: Simple Calculator
```swahili
kazi kuu() {
    namba x = ingiza("Ingiza namba ya kwanza:")
    namba y = ingiza("Ingiza namba ya pili:")
    
    namba jibu = x + y
    andika("Jibu ni:", jibu)
}
```

**Output:**
```
Ingiza namba ya kwanza: 10
Ingiza namba ya pili: 15
Jibu ni: 25
```

### Example 2: Conditional Logic
```swahili
kazi kuu() {
    namba umri = ingiza("Ingiza umri wako:")
    
    kama umri >= 18 {
        andika("Wewe ni mtu mzima")
    } sivyo {
        andika("Wewe ni mtoto")
    }
    
    kama umri == 21 {
        andika("Hongera! Umefika umri wa miaka 21!")
    }
}
```

**Output:**
```
Ingiza umri wako: 25
Wewe ni mtu mzima
```

### Example 3: Nested Conditionals
```swahili
kazi kuu() {
    namba alama = ingiza("Ingiza alama yako:")
    
    kama alama >= 90 {
        andika("A - Bora sana!")
    } sivyo {
        kama alama >= 80 {
            andika("B - Nzuri")
        } sivyo {
            kama alama >= 70 {
                andika("C - Wastani")
            } sivyo {
                andika("D au F - Jitahidi zaidi")
            }
        }
    }
}
```

## 🏗️ Project Structure

```
lugha-yangu/
├── main.go              # Entry point
├── lexer/
│   └── lexer.go        # Tokenization
├── parser/
│   └── parser.go       # Syntax analysis
├── ast/
│   └── ast.go          # Abstract Syntax Tree definitions
├── interpreter/
│   └── interpreter.go  # Code execution
├── environment/
│   └── environment.go  # Variable environment
├── examples/
│   ├── example1.swh           # Basic calculator
│   ├── example2.swh           # Alternative syntax
│   ├── example3.swh           # Minimal example
│   ├── conditional.swh        # Conditional statements
│   ├── nested_if.swh          # Nested conditionals
│   └── conditionals_demo.swh  # Comprehensive demo
└── README.md
```

## 🔧 Architecture

The interpreter follows a traditional architecture:

1. **Lexer** (`lexer/lexer.go`): Converts source code into tokens
2. **Parser** (`parser/parser.go`): Builds an Abstract Syntax Tree (AST)
3. **AST** (`ast/ast.go`): Defines node types for the syntax tree
4. **Interpreter** (`interpreter/interpreter.go`): Executes the AST
5. **Environment** (`environment/environment.go`): Manages variable scope

## 🎯 Supported Operations

### Data Types
- **Numbers**: Integer values
- **Strings**: Text literals (for prompts and output)

### Operations
- **Arithmetic**: `+`, `-`, `*`, `/`
- **Comparison**: `==`, `!=`, `<`, `<=`, `>`, `>=`
- **Input**: `ingiza()` with optional prompt
- **Output**: `andika()` with multiple arguments
- **Assignment**: `=` operator

### Control Flow
- **Functions**: `kazi` keyword for function definitions
- **Conditionals**: `kama`/`sivyo` for if/else statements
- **Nested Logic**: Support for nested conditional statements
- **Main execution**: Automatic execution of `kuu()` function

## 🚧 Current Limitations

- Only integer arithmetic (no floating-point)
- No loops (for/while)
- No arrays or complex data structures
- Single-file programs only
- No function parameters (except main)
- No boolean data type (uses integers for conditions)

## 🔮 Future Enhancements

- [x] Conditional statements (`kama`/`sivyo` for if/else) ✅
- [ ] Loop constructs (`wakati` for while, `kwa` for for)
- [ ] Boolean data type (`kweli`/`uwongo` for true/false)
- [ ] Function parameters and return values
- [ ] String manipulation functions
- [ ] File I/O operations
- [ ] Error handling and better error messages
- [ ] Multi-file support
- [ ] Standard library functions

## 🤝 Contributing

Contributions are welcome! Areas where help is needed:

1. **Language Features**: Add new keywords and constructs
2. **Error Handling**: Improve error messages and debugging
3. **Documentation**: Expand examples and tutorials
4. **Testing**: Add comprehensive test cases
5. **Performance**: Optimize interpreter performance

## 📄 License

This project is open source. Feel free to use, modify, and distribute.

## 🙏 Acknowledgments

- Inspired by the need for programming languages in local languages
- Built with Go's excellent parsing and compilation tools
- Designed for educational purposes and community learning

---

**Karibu kwenye ulimwengu wa programu kwa Kiswahili!** 
*(Welcome to the world of programming in Swahili!)*