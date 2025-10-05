# Lugha Yangu - A Swahili Programming Language

**Lugha Yangu** (meaning "My Language" in Swahili) is a fully-featured, educational programming language with native Swahili syntax. Built in Go, it's designed to make programming more accessible to Swahili speakers by using familiar keywords and concepts while providing modern programming capabilities.

## 🌟 Features

- **Native Swahili Syntax**: All keywords are in Swahili
- **Interactive Input/Output**: Built-in support for user interaction
- **Arithmetic Operations**: Basic mathematical operations
- **Variable Declarations**: Type-safe variable handling
- **String Manipulation**: Comprehensive string operations and functions
- **Function Definitions**: Support for custom functions with parameters and return values
- **Boolean Data Type**: Native boolean support with `kweli`/`uwongo`
- **Conditional Statements**: If/else logic with `kama`/`sivyo`
- **Loop Constructs**: While loops with `wakati` and for loops with `kwa`
- **Loop Control**: Break and continue statements with `vunja`/`endelea`
- **Logical Operations**: AND/OR operations with `na`/`au`
- **Educational Focus**: Perfect for learning programming concepts in Swahili

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

The interpreter will execute the program specified in `main.go`. You can change the input file to run different examples from the `examples/` directory.

## 📝 Language Syntax

### Keywords
| Swahili | English | Description |
|---------|---------|-------------|
| `kazi` | function | Define a function |
| `kuu` | main | Main function (entry point) |
| `namba` | number | Declare a number variable |
| `orodha` | array/list | Declare an array variable |
| `ingiza` | input | Get user input |
| `andika` | print | Print output |
| `rudisha` | return | Return a value |
| `ongeza` | add | Add element to array |
| `ondoa` | remove | Remove element from array |
| `urefu_orodha` | array_length | Get array length |
| `pata` | get | Get element at index |
| `soma` | read | Read file content |
| `andika_faili` | write_file | Write content to file |
| `unda_faili` | create_file | Create empty file |
| `faili_ipo` | file_exists | Check if file exists |
| `ondoa_faili` | delete_file | Delete file |
| `kama` | if | Conditional statement |
| `sivyo` | else | Alternative condition |
| `wakati` | while | While loop |
| `kwa` | for | For loop |
| `vunja` | break | Break out of loop |
| `endelea` | continue | Continue to next iteration |
| `boolean` | boolean | Declare a boolean variable |
| `kweli` | true | Boolean true value |
| `uwongo` | false | Boolean false value |
| `na` | and | Logical AND operator |
| `au` | or | Logical OR operator |
| `maneno` | string | Declare a string variable |
| `urefu` | length | Get string length |
| `unganisha` | concatenate | Join strings together |
| `kata` | substring | Extract part of string |
| `badilisha` | replace | Replace text in string |
| `tafuta` | find | Find position of substring |
| `awali` | starts_with | Check if string starts with text |
| `mwisho` | ends_with | Check if string ends with text |
| `herufi_kubwa` | uppercase | Convert to uppercase |
| `herufi_ndogo` | lowercase | Convert to lowercase |
| `ondoa_nafasi` | trim | Remove whitespace |
| `gawanya_maneno` | split | Split string into parts |
| `rudisha` | return | Return a value from function |

### Basic Syntax

#### Variable Declaration
```swahili
namba x = 10
namba y = ingiza("Ingiza namba:")
boolean iko_kweli = kweli
boolean si_kweli = uwongo
maneno jina = "Amina"
maneno salamu = "Habari"
orodha namba namba_za_kwanza = [1, 2, 3, 4, 5]  # Array declaration
```

#### Function Definition
```swahili
kazi kuu() {
    // Main function (entry point)
}

kazi jumla(namba x, namba y) namba {
    // Function with parameters and return type
    rudisha x + y
}

kazi salamu(namba umri) {
    // Function with parameters, no return type
    andika("Habari!")
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

#### Array Operations
```swahili
orodha namba arr = [1, 2, 3]        # Create array
namba urefu = urefu_orodha(arr)      # Get length
namba kipengele = pata(arr, 0)       # Get element at index 0
ongeza(arr, 4)                       # Add element to end
ondoa(arr, 1)                        # Remove element at index 1
andika("Orodha:", arr)               # Print array: [1, 3, 4]
```

#### File I/O Operations
```swahili
# File creation and writing
unda_faili("data.txt")                    # Create empty file
andika_faili("data.txt", "Hello World")   # Write content (overwrite)
andika_faili("data.txt", "\nNew line", kweli)  # Append content

# File reading
maneno maudhui = soma("data.txt")         # Read file content
andika("Content:", maudhui)               # Display content

# File management
boolean ipo = faili_ipo("data.txt")       # Check if file exists
ondoa_faili("data.txt")                   # Delete file
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

#### Boolean Operations
```swahili
boolean a = kweli
boolean b = uwongo

boolean c = a na b    // Logical AND (false)
boolean d = a au b    // Logical OR (true)
boolean e = a == kweli // Boolean comparison (true)
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

#### String Manipulation

##### String Variables and Concatenation
```swahili
maneno jina = "Amina"
maneno mji = "Dar es Salaam"
maneno ujumbe = "Habari " + jina + " kutoka " + mji

# Using unganisha function
maneno salamu = unganisha("Habari za ", "asubuhi", ", ", jina, "!")
```

##### String Functions
```swahili
maneno neno = "Habari Dunia"

# Get string length
namba urefu_neno = urefu(neno)  # 12

# Convert case
maneno kubwa = herufi_kubwa(neno)    # "HABARI DUNIA"
maneno ndogo = herufi_ndogo(neno)    # "habari dunia"

# Substring operations
maneno sehemu = kata(neno, 0, 6)     # "Habari"
maneno mwisho = kata(neno, 7)        # "Dunia"

# Find and replace
namba mahali = tafuta(neno, "Dunia") # 7
maneno mpya = badilisha(neno, "Dunia", "Tanzania") # "Habari Tanzania"

# Prefix and suffix checks
boolean inaanza = awali(neno, "Habari")  # kweli
boolean inaishia = mwisho(neno, "Dunia") # kweli

# Trim whitespace
maneno na_nafasi = "   Karibu   "
maneno safi = ondoa_nafasi(na_nafasi)    # "Karibu"

# Word counting
maneno sentensi = "Hii ni sentensi yenye maneno kadhaa"
namba idadi = gawanya_maneno(sentensi)   # 6
```

##### String Comparison
```swahili
maneno neno1 = "Habari"
maneno neno2 = "Habari"
boolean ni_sawa = neno1 == neno2  # kweli
boolean si_sawa = neno1 != "Mambo"  # kweli
```

## 📚 Examples

> **Note**: All examples are available in the `examples/` directory. You can run any example by changing the file path in `main.go`.

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

### Example 8: Functions with Parameters and Return Values
```swahili
kazi jumla(namba a, namba b) namba {
    namba jibu = a + b
    rudisha jibu
}

kazi salamu(namba umri) {
    kama umri >= 18 {
        andika("Habari za asubuhi, mtu mzima!")
    } sivyo {
        andika("Habari za asubuhi, kijana!")
    }
}

kazi ni_kubwa(namba x, namba y) boolean {
    rudisha x > y
}

kazi kuu() {
    namba jibu = jumla(10, 5)
    andika("10 + 5 =", jibu)
    
    salamu(25)
    
    boolean kubwa = ni_kubwa(8, 3)
    andika("8 > 3:", kubwa)
}
```

**Output:**
```
10 + 5 = 15
Habari za asubuhi, mtu mzima!
8 > 3: true
```

### Example 7: Boolean Data Type
```swahili
kazi kuu() {
    boolean iko_jua = kweli
    boolean mvua = uwongo
    namba joto = 25
    
    andika("Hali ya hewa:")
    andika("Jua:", iko_jua)
    andika("Mvua:", mvua)
    
    kama iko_jua na joto > 20 {
        andika("Siku nzuri!")
    } sivyo {
        andika("Hali mbaya")
    }
    
    boolean hali_nzuri = iko_jua au joto > 30
    andika("Hali nzuri:", hali_nzuri)
}
```

**Output:**
```
Hali ya hewa:
Jua: true
Mvua: false
Siku nzuri!
Hali nzuri: true
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

### Example 9: File I/O Operations
```swahili
kazi kuu() {
    # Create and manage files
    maneno jina_faili = "data.txt"
    unda_faili(jina_faili)
    
    # Write data to file
    orodha namba namba_zangu = [10, 20, 30]
    andika_faili(jina_faili, "NAMBA ZA MUHIMU:\n")
    
    namba i = 0
    namba urefu = urefu_orodha(namba_zangu)
    wakati i < urefu {
        namba namba_ya_sasa = pata(namba_zangu, i)
        andika_faili(jina_faili, "Namba ", kweli)
        andika_faili(jina_faili, namba_ya_sasa, kweli)
        andika_faili(jina_faili, "\n", kweli)
        i = i + 1
    }
    
    # Read and display file content
    maneno maudhui = soma(jina_faili)
    andika("Maudhui ya faili:")
    andika(maudhui)
    
    # Check if file exists and clean up
    boolean ipo = faili_ipo(jina_faili)
    andika("Faili ipo:", ipo)
    ondoa_faili(jina_faili)
    andika("Faili imeondolewa")
}
```

### Example 10: String Manipulation
```swahili
kazi kuu() {
    # Basic string operations
    maneno jina = "Amina"
    maneno mji = "Dar es Salaam"
    maneno ujumbe = "Habari " + jina + " kutoka " + mji
    
    andika("Basic concatenation:", ujumbe)
    
    namba urefu_ujumbe = urefu(ujumbe)
    andika("Length of message:", urefu_ujumbe)
    
    # String manipulation functions
    maneno kubwa = herufi_kubwa(ujumbe)
    maneno ndogo = herufi_ndogo(ujumbe)
    
    andika("Uppercase:", kubwa)
    andika("Lowercase:", ndogo)
    
    # Substring operations
    maneno sehemu = kata(ujumbe, 0, 6)
    andika("First 6 characters:", sehemu)
    
    # Find and replace
    namba mahali = tafuta(ujumbe, "Amina")
    andika("Position of 'Amina':", mahali)
    
    maneno mpya = badilisha(ujumbe, "Amina", "Fatuma")
    andika("After replacement:", mpya)
    
    # Complex string building
    maneno salamu_kamili = unganisha("Habari za ", "asubuhi", ", ", jina, "!")
    andika("Complex greeting:", salamu_kamili)
}
```

**Output:**
```
Basic concatenation: Habari Amina kutoka Dar es Salaam
Length of message: 33
Uppercase: HABARI AMINA KUTOKA DAR ES SALAAM
Lowercase: habari amina kutoka dar es salaam
First 6 characters: Habari
Position of 'Amina': 7
After replacement: Habari Fatuma kutoka Dar es Salaam
Complex greeting: Habari za asubuhi, Amina!
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
│   ├── simple_nested_break.swh # Nested loops with break/continue
│   ├── boolean_basic.swh      # Basic boolean operations
│   ├── boolean_conditionals.swh # Booleans with conditionals
│   ├── boolean_loops.swh      # Booleans with loops
│   ├── functions_basic.swh    # Basic function examples
│   ├── functions_advanced.swh # Advanced function features
│   ├── functions_comprehensive.swh # Comprehensive function demo
│   ├── string_basic.swh       # Basic string operations
│   ├── string_manipulation.swh # String manipulation functions
│   ├── string_functions.swh   # String functions with user-defined functions
│   └── string_comprehensive.swh # Comprehensive string demo
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
- **Booleans**: True/false values (`kweli`/`uwongo`)
- **Strings**: Text values with comprehensive manipulation functions

### Operations
- **Arithmetic**: `+`, `-`, `*`, `/`
- **Comparison**: `==`, `!=`, `<`, `<=`, `>`, `>=`
- **Logical**: `na` (AND), `au` (OR)
- **Input**: `ingiza()` with optional prompt
- **Output**: `andika()` with multiple arguments
- **Assignment**: `=` operator
- **File I/O**: `soma()`, `andika_faili()`, `unda_faili()`, `faili_ipo()`, `ondoa_faili()`
- **Array Operations**: `ongeza()`, `ondoa()`, `urefu_orodha()`, `pata()`

### Control Flow
- **Functions**: `kazi` keyword for function definitions with parameters and return types
- **Function Calls**: Support for user-defined functions with arguments
- **Return Statements**: `rudisha` keyword for returning values
- **Conditionals**: `kama`/`sivyo` for if/else statements
- **Loops**: `wakati` for while loops, `kwa` for for loops
- **Loop Control**: `vunja` for break, `endelea` for continue
- **Nested Logic**: Support for nested conditional and loop statements
- **Main execution**: Automatic execution of `kuu()` function

## 🚧 Current Limitations

- Only integer arithmetic (no floating-point)
- No arrays or complex data structures
- Single-file programs only
- Limited operator precedence (use parentheses for complex expressions)
- No function overloading
- No recursive function optimization
- No labeled break/continue (only affects innermost loop)

## 🔮 Future Enhancements

### ✅ Completed Features
- [x] Conditional statements (`kama`/`sivyo` for if/else) ✅
- [x] Loop constructs (`wakati` for while, `kwa` for for) ✅
- [x] Break and continue statements (`vunja`/`endelea` for break/continue) ✅
- [x] Boolean data type (`kweli`/`uwongo` for true/false) ✅
- [x] Function parameters and return values ✅
- [x] Logical operators (`na`/`au` for AND/OR) ✅
- [x] String data type and manipulation functions ✅

### 🚀 Planned Features
- [x] Array/list data structures ✅
- [x] File I/O operations (`soma`/`andika_faili` for read/write) ✅
- [ ] Error handling with try/catch (`jaribu`/`shika`)
- [ ] Multi-file support and imports
- [ ] Standard library functions
- [ ] Floating-point arithmetic
- [ ] Comments support (`#` or `//`)
- [ ] Better error messages with line numbers

## 🤝 Contributing

Contributions are welcome! Areas where help is needed:

1. **Language Features**: Add new keywords and constructs
2. **Error Handling**: Improve error messages and debugging
3. **Documentation**: Expand examples and tutorials
4. **Testing**: Add comprehensive test cases
5. **Performance**: Optimize interpreter performance

## 📄 License

This project is open source. Feel free to use, modify, and distribute.

## 📚 Additional Documentation

- **[FUNCTIONS.md](FUNCTIONS.md)**: Comprehensive guide to function parameters and return values
- **[LOOPS.md](LOOPS.md)**: Detailed documentation on loop constructs and control flow
- **[BOOLEANS.md](BOOLEANS.md)**: Complete guide to boolean data types and logical operations
- **[STRINGS.md](STRINGS.md)**: Complete guide to string manipulation and functions

## 🎓 Educational Use

Kwenda is perfect for:
- Teaching programming concepts in Swahili
- Computer science education in East Africa
- Learning programming fundamentals
- Understanding interpreter design and implementation
- Cultural preservation through technology

## 🌍 Language Philosophy

Kwenda believes that programming should be accessible in one's native language. By using Swahili keywords and concepts, we aim to:
- Lower the barrier to entry for programming
- Preserve and promote local languages in technology
- Make computer science education more inclusive
- Demonstrate that programming concepts are universal

## 🙏 Acknowledgments

- Inspired by the need for programming languages in local languages
- Built with Go's excellent parsing and compilation tools
- Designed for educational purposes and community learning
- Special thanks to the Swahili-speaking developer community

---

**Karibu kwenye ulimwengu wa programu kwa Kiswahili!** 
*(Welcome to the world of programming in Swahili!)*

**Lugha Yangu ni zaidi ya lugha ya programu - ni daraja kuelekea teknolojia kwa wote.**
*(Lugha Yangu is more than a programming language - it's a bridge to technology for everyone.)*