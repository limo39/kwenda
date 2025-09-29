# Lugha Yangu - A Swahili Programming Language

**Lugha Yangu** (meaning "My Language" in Swahili) is a simple, educational programming language with Swahili syntax. Built in Go, it's designed to make programming more accessible to Swahili speakers by using familiar keywords and concepts.

## 🌟 Features

- **Native Swahili Syntax**: All keywords are in Swahili
- **Interactive Input/Output**: Built-in support for user interaction
- **Arithmetic Operations**: Basic mathematical operations
- **Variable Declarations**: Type-safe variable handling
- **Function Definitions**: Support for custom functions
- **Conditional Statements**: If/else logic with `kama`/`sivyo`
- **Loop Constructs**: While loops with `wakati` and for loops with `kwa`
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
| `wakati` | while | While loop |
| `kwa` | for | For loop |
| `vunja` | break | Break out of loop |
| `endelea` | continue | Continue to next iteration |

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

#### Loop Constructs

##### While Loops
```swahili
namba i = 1
wakati i <= 5 {
    andika("Namba:", i)
    i = i + 1
}
```

##### For Loops (Full Syntax)
```swahili
kwa i = 0; i < 5; i = i + 1 {
    andika("Iteration:", i)
}
```

##### For Loops (Simple Condition)
```swahili
namba x = 10
kwa x > 0 {
    andika("x ni:", x)
    x = x - 1
}
```

#### Loop Control Statements

##### Break Statement (`vunja`)
```swahili
wakati i <= 10 {
    kama i == 5 {
        vunja  # Exit the loop
    }
    andika("i =", i)
    i = i + 1
}
```

##### Continue Statement (`endelea`)
```swahili
wakati i < 5 {
    i = i + 1
    kama i == 3 {
        endelea  # Skip to next iteration
    }
    andika("i =", i)
}
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

### Example 4: While Loop
```swahili
kazi kuu() {
    namba i = 1
    
    andika("Mfuatano wa namba kutoka 1 hadi 5:")
    
    wakati i <= 5 {
        andika("Namba:", i)
        i = i + 1
    }
    
    andika("Mwisho wa loop!")
}
```

**Output:**
```
Mfuatano wa namba kutoka 1 hadi 5:
Namba: 1
Namba: 2
Namba: 3
Namba: 4
Namba: 5
Mwisho wa loop!
```

### Example 5: For Loop
```swahili
kazi kuu() {
    andika("For loop kutoka 0 hadi 4:")
    
    kwa i = 0; i < 5; i = i + 1 {
        andika("Iteration:", i)
    }
    
    andika("For loop simple:")
    namba j = 3
    
    kwa j > 0 {
        andika("j ni:", j)
        j = j - 1
    }
}
```

**Output:**
```
For loop kutoka 0 hadi 4:
Iteration: 0
Iteration: 1
Iteration: 2
Iteration: 3
Iteration: 4
For loop simple:
j ni: 3
j ni: 2
j ni: 1
```

### Example 6: Break and Continue
```swahili
kazi kuu() {
    andika("Break example:")
    namba i = 1
    wakati i <= 10 {
        kama i == 5 {
            andika("Breaking at i =", i)
            vunja
        }
        andika("i =", i)
        i = i + 1
    }
    
    andika("Continue example:")
    namba j = 0
    wakati j < 5 {
        j = j + 1
        kama j == 3 {
            andika("Skipping j =", j)
            endelea
        }
        andika("j =", j)
    }
}
```

**Output:**
```
Break example:
i = 1
i = 2
i = 3
i = 4
Breaking at i = 5
Continue example:
j = 1
j = 2
Skipping j = 3
j = 4
j = 5
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
│   ├── conditionals_demo.swh  # Comprehensive demo
│   ├── while_loop.swh         # While loop examples
│   ├── for_loop.swh           # For loop examples
│   ├── nested_loops.swh       # Nested loop examples
│   ├── break_example.swh      # Break statement examples
│   ├── continue_example.swh   # Continue statement examples
│   └── simple_nested_break.swh # Nested loops with break/continue
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
- **Loops**: `wakati` for while loops, `kwa` for for loops
- **Loop Control**: `vunja` for break, `endelea` for continue
- **Nested Logic**: Support for nested conditional and loop statements
- **Main execution**: Automatic execution of `kuu()` function

## 🚧 Current Limitations

- Only integer arithmetic (no floating-point)
- No arrays or complex data structures
- Single-file programs only
- No function parameters (except main)
- No boolean data type (uses integers for conditions)
- No labeled break/continue (only affects innermost loop)

## 🔮 Future Enhancements

- [x] Conditional statements (`kama`/`sivyo` for if/else) ✅
- [x] Loop constructs (`wakati` for while, `kwa` for for) ✅
- [x] Break and continue statements (`vunja`/`endelea` for break/continue) ✅
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