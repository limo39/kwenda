# Kwenda - A Swahili Programming Language

**Kwenda** (meaning "Go" or "Move Forward" in Swahili) is a fully-featured, educational programming language with native Swahili syntax. Built in Go, it's designed to make programming more accessible to Swahili speakers by using familiar keywords and concepts while providing modern programming capabilities.

## 🌟 Features

### Core Language Features
- **Native Swahili Syntax**: All keywords and concepts in Swahili
- **Interactive I/O**: Built-in support for user input and output with `ingiza` and `andika`
- **Comments**: Single-line comments using `#`

### Data Types & Variables
- **Numbers**: Integer and floating-point support with `namba`
- **Strings**: Text manipulation with `maneno`
- **Booleans**: True/false values with `kweli`/`uwongo`
- **Arrays**: Dynamic lists with `orodha` keyword
- **Dictionaries**: Key-value maps with `kamusi` keyword
- **Type-safe Declarations**: Explicit type declarations for variables

### Operators & Expressions
- **Arithmetic**: `+`, `-`, `*`, `/`, `%` (modulo)
- **Comparison**: `==`, `!=`, `<`, `<=`, `>`, `>=`
- **Logical**: `na` (AND), `au` (OR)
- **Assignment**: Variable and member assignment

### Control Flow
- **Conditionals**: If/else statements with `kama`/`sivyo`
- **While Loops**: Iteration with `wakati`
- **For Loops**: Flexible looping with `kwa`
- **Loop Control**: Break (`vunja`) and continue (`endelea`)

### Functions & Modules
- **Function Definitions**: Custom functions with `kazi` keyword
- **Parameters & Return Values**: Type-safe function signatures
- **Module System**: Multi-file support with `leta` imports
- **Module Namespaces**: Organized code with dot notation access

### Object-Oriented Programming
- **Classes**: Define classes with `darasa` keyword
- **Inheritance**: Class inheritance with colon syntax (`darasa Child : Parent`)
- **Constructors**: Initialize objects with `unda` method
- **Properties**: Class properties with type declarations
- **Methods**: Class methods with `kazi` keyword
- **Method Calls**: Dot notation for method invocation (`object.method()`)
- **Method Overriding**: Child classes can override parent methods
- **Member Access**: Dot notation for properties (`object.property` or `hii.property`)
- **Instance Reference**: `hii` keyword for this/self reference

### Data Structures
- **Array Operations**: Add, remove, access, and get length
- **Dictionary Operations**: Create, access, modify key-value pairs
- **String Functions**: Length, substring, replace, find, case conversion, trim, split

### Error Handling
- **Try-Catch-Finally**: Robust error handling with `jaribu`/`shika`/`hatimaye`
- **Throw Errors**: Custom error throwing with `tupa`
- **Bilingual Error Messages**: Errors in both Swahili and English
- **Contextual Error Info**: Detailed error context and suggestions

### File I/O
- **Read Files**: Load file contents with `soma`
- **Write Files**: Save data with `andika_faili`
- **Create Files**: Initialize new files with `unda_faili`
- **File Management**: Check existence and delete files

### Educational Focus
- **Accessible Syntax**: Programming in native Swahili language
- **Clear Error Messages**: Helpful debugging information
- **Comprehensive Examples**: Learn by example
- **Modern Features**: Real-world programming capabilities

## 🚀 Quick Start

### Prerequisites
- Go 1.23.3 or higher

### Installation
```bash
git clone <repository-url>
cd kwenda
go build
```

### Running a Program
```bash
# Run a Kwenda program
./kwenda program.swh

# Or using go run
go run main.go program.swh
```

### Getting Help
```bash
# Display help and usage information
./kwenda --help

# Display version information
./kwenda --version

# Quick usage hint
./kwenda
```

The interpreter will execute the specified `.swh` file. You can run examples from the `examples/` directory.

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
| `leta` | import | Import a module file |
| `jaribu` | try | Try block for error handling |
| `shika` | catch | Catch block for handling errors |
| `hatimaye` | finally | Finally block (always executes) |
| `tupa` | throw | Throw an error |
| `kamusi` | dictionary/map | Declare a dictionary variable |
| `darasa` | class | Define a class |
| `unda` | new/create | Create a class instance |
| `hii` | this/self | Reference to current instance |
| `chini` | floor | Round number down to nearest integer |
| `juu` | ceil | Round number up to nearest integer |
| `zunguka` | round | Round number to nearest integer |
| `kata_desimali` | truncate | Remove decimal part (round toward zero) |
| `mzizi_mraba` | sqrt | Calculate square root of a number |
| `mzizi_mchemraba` | cbrt | Calculate cube root of a number |
| `mzizi` | nth_root | Calculate nth root of a number |
| `exp` | exp | Calculate e raised to the power x (e^x) |
| `log_asili` | ln/log | Calculate natural logarithm (base e) |
| `log10` | log10 | Calculate base-10 logarithm |
| `log2` | log2 | Calculate base-2 logarithm |
| `sin` | sin | Calculate sine of angle (in radians) |
| `cos` | cos | Calculate cosine of angle (in radians) |
| `tan` | tan | Calculate tangent of angle (in radians) |
| `asin` | asin/arcsin | Calculate arc sine (inverse sine), returns radians |
| `acos` | acos/arccos | Calculate arc cosine (inverse cosine), returns radians |
| `atan` | atan/arctan | Calculate arc tangent (inverse tangent), returns radians |
| `atan2` | atan2 | Two-argument arc tangent for proper quadrant handling |
| `radians` | radians | Convert degrees to radians |
| `degrees` | degrees | Convert radians to degrees |
| `sinh` | sinh | Calculate hyperbolic sine |
| `cosh` | cosh | Calculate hyperbolic cosine |
| `tanh` | tanh | Calculate hyperbolic tangent |
| `asinh` | asinh/arcsinh | Calculate inverse hyperbolic sine |
| `acosh` | acosh/arccosh | Calculate inverse hyperbolic cosine (domain: x >= 1) |
| `atanh` | atanh/arctanh | Calculate inverse hyperbolic tangent (domain: -1 < x < 1) |

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
kamusi person = {"jina": "Amina", "umri": 25}    # Dictionary declaration
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

## 📐 Mathematical Operations

For comprehensive documentation of all mathematical functions, including rounding, roots, exponential, logarithm, and trigonometric operations, please refer to **[MATH.md](MATH.md)**.

---

## ?? Examples

> **?? For comprehensive examples and tutorials, see [EXAMPLES.md](EXAMPLES.md)**
### Hello World
```swahili
kazi kuu() {
    andika("Habari Dunia!")
}
```

### Variables and Functions
```swahili
kazi jumla(namba a, namba b) namba {
    rudisha a + b
}

kazi kuu() {
    namba x = 10
    namba y = 5
    namba jibu = jumla(x, y)
    andika("10 + 5 =", jibu)
}
```

### Classes and Objects
```swahili
darasa Mtu {
    maneno jina
    namba umri
    
    kazi unda(maneno j, namba u) {
        hii.jina = j
        hii.umri = u
    }
}

kazi kuu() {
    kamusi mtu = unda Mtu("Amina", 25)
    andika("Name:", mtu["jina"])
}
```

**For more examples, see [EXAMPLES.md](EXAMPLES.md)**

## 🏗️ Project Structure

```
kwenda/
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
│   ├── complete_loops_demo.swh # Complete loops with break/continue
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
│   ├── string_comprehensive.swh # Comprehensive string demo
│   ├── error_handling_simple.swh # Simple error handling
│   ├── error_handling_basic.swh # Basic try/catch examples
│   ├── simple_try.swh         # Simple try/catch test
│   └── multi_file_demo.swh    # Multi-file module demo
├── modules/
│   ├── math.swh               # Math utility functions
│   └── strings.swh            # String utility functions
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
- **Mathematical Rounding**: `chini()` (floor), `juu()` (ceil), `zunguka()` (round), `kata_desimali()` (truncate)
- **Mathematical Roots**: `mzizi_mraba()` (square root), `mzizi_mchemraba()` (cube root), `mzizi()` (nth root)

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

- Limited operator precedence (use parentheses for complex expressions)
- No function overloading
- No recursive function optimization
- No labeled break/continue (only affects innermost loop)
- Module functions must be called with module prefix (e.g., `math.add()`)
- No circular import detection
- Array parameters to user-defined functions have limited support
- Floating-point precision follows IEEE 754 standard (may have rounding artifacts)
- Limited standard library (growing)

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
- [x] Error handling with try/catch (`jaribu`/`shika`/`hatimaye`) ✅
- [x] Multi-file support and imports (`leta`) ✅
- [x] Module system with namespaces ✅
- [x] Standard library modules (math, strings, arrays) ✅
- [x] Floating-point arithmetic ✅
- [x] Comments support with `#` ✅
- [x] Improved error messages with context ✅
- [x] Object-oriented programming (function-based pattern) ✅
- [x] Dictionary/map data structures (`kamusi` keyword) ✅
- [x] Class syntax with `darasa` keyword ✅
- [x] Instance creation with `unda` keyword ✅
- [x] Self-reference with `hii` keyword ✅
- [x] Dot notation for member access ✅
- [x] Method calls with dot notation (e.g., `object.method()`) ✅
- [x] Class inheritance ✅
- [x] Lambda functions ✅
- [x] Built-in mathematical rounding functions (floor, ceil, round, truncate) ✅
- [x] Built-in mathematical root functions (square root, cube root, nth root) ✅
- [ ] List comprehensions

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
- **[OOP.md](OOP.md)**: Complete guide to object-oriented programming patterns
- **[DICTIONARY_SUMMARY.md](DICTIONARY_SUMMARY.md)**: Dictionary/map implementation details
- **[OOP_SUMMARY.md](OOP_SUMMARY.md)**: OOP implementation summary and examples

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

**Kwenda ni zaidi ya lugha ya programu - ni daraja kuelekea teknolojia kwa wote.**
*(Kwenda is more than a programming language - it's a bridge to technology for everyone.)*

---

## 📐 Mathematical & Advanced Operations

For comprehensive documentation of all mathematical and advanced operations, please see **[MATH.md](MATH.md)**.

This includes:
- **Rounding Operations**: floor, ceiling, round, truncate
- **Root Operations**: square root, cube root, nth root
- **Exponential & Logarithm**: exp, natural log, log base 10, log base 2
- **Trigonometric Functions**: sin, cos, tan, asin, acos, atan, angle conversions
- **Bitwise Operations**: AND, OR, XOR, NOT, left shift, right shift
- **Modular Arithmetic**: modulo, remainder, modular exponentiation
- **Random Number Generation**: random floats, random integers, seeding

### Bitwise Operations

Kwenda provides six bitwise operations for manipulating individual bits in integers. These are useful for low-level programming, optimization, flags, and cryptography.

#### Basic Bitwise Operations

```swahili
# Bitwise AND - na_kidogo(a, b)
# Returns 1 where both operands have 1, else 0
namba result1 = na_kidogo(12, 10)      # 1100 & 1010 = 1000 = 8
namba result2 = na_kidogo(15, 7)       # 1111 & 0111 = 0111 = 7

# Bitwise OR - au_kidogo(a, b)
# Returns 1 where at least one operand has 1
namba result3 = au_kidogo(12, 10)      # 1100 | 1010 = 1110 = 14
namba result4 = au_kidogo(8, 4)        # 1000 | 0100 = 1100 = 12

# Bitwise XOR - ama_kidogo(a, b)
# Returns 1 where operands differ, else 0
namba result5 = ama_kidogo(12, 10)     # 1100 ^ 1010 = 0110 = 6
namba result6 = ama_kidogo(15, 15)     # 1111 ^ 1111 = 0000 = 0

# Bitwise NOT - si_kidogo(a)
# Inverts all bits
namba result7 = si_kidogo(0)           # ~0 = -1
namba result8 = si_kidogo(5)           # ~5 = -6

# Left Shift - gezo_kushoto(a, n)
# Shifts bits left by n positions (equivalent to multiplying by 2^n)
namba result9 = gezo_kushoto(5, 1)     # 0101 << 1 = 1010 = 10
namba result10 = gezo_kushoto(3, 2)    # 0011 << 2 = 1100 = 12
namba result11 = gezo_kushoto(1, 8)    # 0001 << 8 = 100000000 = 256

# Right Shift - gezo_kulia(a, n)
# Shifts bits right by n positions (equivalent to dividing by 2^n)
namba result12 = gezo_kulia(10, 1)     # 1010 >> 1 = 0101 = 5
namba result13 = gezo_kulia(12, 2)     # 1100 >> 2 = 0011 = 3
namba result14 = gezo_kulia(256, 8)    # >> 8 = 1
```

#### Practical Applications

**Flag Management:**
```swahili
# Using bits as boolean flags
namba flags = 0
namba READ = gezo_kushoto(1, 0)        # Bit 0
namba WRITE = gezo_kushoto(1, 1)       # Bit 1
namba EXECUTE = gezo_kushoto(1, 2)     # Bit 2

# Set flags (OR operation)
flags = au_kidogo(flags, READ)         # Enable read
flags = au_kidogo(flags, WRITE)        # Enable write

# Check if flag is set (AND operation)
kama na_kidogo(flags, READ) == READ {
    andika("Read permission granted")
}

# Clear a flag (AND with NOT)
flags = na_kidogo(flags, si_kidogo(WRITE))  # Disable write
```

**Bit Counting:**
```swahili
# Count number of set bits (population count)
namba num = 15
namba count = 0

wakati num > 0 {
    kama na_kidogo(num, 1) == 1 {
        count = count + 1
    }
    num = gezo_kulia(num, 1)
}
andika("Set bits:", count)  # 4 (15 = 1111 in binary)
```

**Swapping Values Without Temp Variable:**
```swahili
namba a = 5
namba b = 7

# XOR swap
a = ama_kidogo(a, b)
b = ama_kidogo(a, b)
a = ama_kidogo(a, b)

andika("a =", a, ", b =", b)  # a = 7, b = 5
```

**Bitwise Comparison Table:**
| Operation | Example | Result | Binary |
|-----------|---------|--------|--------|
| AND | na_kidogo(12, 10) | 8 | 1000 |
| OR | au_kidogo(12, 10) | 14 | 1110 |
| XOR | ama_kidogo(12, 10) | 6 | 0110 |
| NOT | si_kidogo(5) | -6 | ...11111010 |
| Left Shift | gezo_kushoto(5, 1) | 10 | 1010 |
| Right Shift | gezo_kulia(10, 1) | 5 | 0101 |

### Modular Arithmetic

Modular arithmetic is essential for working with remainders, cryptography, and cyclic calculations. Kwenda provides three modular operations.

#### Basic Modular Operations

```swahili
# Modulo - modulo(a, b)
# Returns remainder after division (handles floats)
namba r1 = modulo(10, 3)               # 1
namba r2 = modulo(20, 6)               # 2
namba r3 = modulo(7, 7)                # 0
namba r4 = modulo(15, 4)               # 3
namba r5 = modulo(100, 10)             # 0

# Remainder - baki(a, b)
# Returns remainder for integers
namba rem1 = baki(10, 3)               # 1
namba rem2 = baki(25, 5)               # 0
namba rem3 = baki(17, 5)               # 2
namba rem4 = baki(100, 7)              # 2

# Modular Power - tweza_modulo(base, exponent, modulus)
# Calculates (base^exponent) mod modulus efficiently
namba mp1 = tweza_modulo(2, 3, 5)      # 2^3 mod 5 = 8 mod 5 = 3
namba mp2 = tweza_modulo(3, 4, 5)      # 3^4 mod 5 = 81 mod 5 = 1
namba mp3 = tweza_modulo(7, 2, 11)     # 7^2 mod 11 = 49 mod 11 = 5
namba mp4 = tweza_modulo(2, 10, 1000)  # 2^10 mod 1000 = 1024 mod 1000 = 24
```

#### Practical Applications

**Cyclic Sequences:**
```swahili
# Generate a repeating pattern
kwa i = 0; i < 20; i = i + 1 {
    namba pattern = modulo(i, 3)
    andika("Index", i, "-> Pattern", pattern)  # 0,1,2,0,1,2,...
}
```

**Day of Week Calculation:**
```swahili
# Calculate day of week (0-6) for scheduling
namba days_since_epoch = 5000
namba day_of_week = modulo(days_since_epoch, 7)
orodha siku = ["Jumapili", "Jumatatu", "Jumatano", "Kamis", "Ijumaa", "Jumamosi", "Jumapili"]
andika("Day:", siku(day_of_week))
```

**Modular Inverse (via Extended Euclidean Algorithm):**
```swahili
# Find modular multiplicative inverse
# If a * x ≡ 1 (mod m), then x is the inverse of a
# For prime moduli, use: tweza_modulo(a, m-2, m)

namba a = 3
namba m = 7
namba inverse = tweza_modulo(a, m - 2, m)
andika("Modular inverse of", a, "mod", m, "is", inverse)  # 5
andika("Verification:", modulo(a * inverse, m))  # 1
```

**Cryptographic Applications:**
```swahili
# RSA-style modular exponentiation (simplified)
namba message = 42
namba exponent = 17
namba modulus = 1000000007

namba encrypted = tweza_modulo(message, exponent, modulus)
andika("Encrypted:", encrypted)
```

### Rounding Functions

Four rounding functions provide different strategies for converting floating-point numbers to integers, each with distinct behavior for negative numbers.

#### Rounding Operations

```swahili
# Round - pindika(x)
# Rounds to nearest integer
namba a1 = pindika(3.4)                # 3
namba a2 = pindika(3.5)                # 4
namba a3 = pindika(3.6)                # 4
namba a4 = pindika(-2.7)               # -3

# Floor - sakafu(x)
# Rounds down toward negative infinity
namba b1 = sakafu(3.9)                 # 3
namba b2 = sakafu(3.1)                 # 3
namba b3 = sakafu(-2.3)                # -3
namba b4 = sakafu(-2.9)                # -3

# Ceiling - dari(x)
# Rounds up toward positive infinity
namba c1 = dari(3.1)                   # 4
namba c2 = dari(3.9)                   # 4
namba c3 = dari(-2.1)                  # -2
namba c4 = dari(-2.9)                  # -2

# Truncate - kata(x)
# Rounds toward zero (removes decimal part)
namba d1 = kata(3.9)                   # 3
namba d2 = kata(3.1)                   # 3
namba d3 = kata(-2.7)                  # -2
namba d4 = kata(-2.1)                  # -2
```

#### Comparison Table

| Value | pindika | sakafu | dari | kata |
|-------|---------|--------|------|------|
| 3.2 | 3 | 3 | 4 | 3 |
| 3.7 | 4 | 3 | 4 | 3 |
| -2.3 | -2 | -3 | -2 | -2 |
| -2.7 | -3 | -3 | -2 | -2 |

#### Practical Applications

**Currency Rounding:**
```swahili
namba price = 19.95
namba quantity = 3
namba total = price * quantity  # 59.85

namba rounded_total = pindika(total * 100) / 100  # 59.85
andika("Total to charge: TSh", rounded_total)
```

**Grid Positioning:**
```swahili
namba x_float = 3.7
namba y_float = 2.2

# Snap to grid (floor)
namba grid_x = sakafu(x_float)         # 3
namba grid_y = sakafu(y_float)         # 2
andika("Snapped to grid: (", grid_x, ", ", grid_y, ")")

# Snap to nearest grid (round)
namba snap_x = pindika(x_float)        # 4
namba snap_y = pindika(y_float)        # 2
andika("Snapped nearest: (", snap_x, ", ", snap_y, ")")
```

**Data Compression - Quantization:**
```swahili
# Reduce precision to 1 decimal place
namba sensor_value = 3.14159
namba quantized = pindika(sensor_value * 10) / 10  # 3.1
andika("Quantized value:", quantized)
```

### Random Number Generation

Generate random numbers with optional seeding for reproducible results. Three modes: float [0,1), integer [0,n), and integer range [min,max].

#### Basic Random Operations

```swahili
# Random float [0.0, 1.0)
namba rand1 = nasibu()                 # 0.372... 
namba rand2 = nasibu()                 # 0.814...
namba rand3 = nasibu()                 # 0.521...

# Random integer [0, n)
namba rand_int1 = nasibu(10)           # 0-9
namba rand_int2 = nasibu(100)          # 0-99
namba rand_int3 = nasibu(6) + 1        # 1-6 (dice roll)

# Random integer [min, max]
namba rand_range1 = nasibu(1, 10)      # 1-10
namba rand_range2 = nasibu(50, 100)    # 50-100
namba rand_range3 = nasibu(-10, 10)    # -10 to 10
```

#### Seeding for Reproducibility

```swahili
# Set seed for reproducible sequences
weka_mbegu(42)

namba val1 = nasibu(100)
namba val2 = nasibu(100)
andika("First run:", val1, val2)

# Reset seed to same value - get same sequence
weka_mbegu(42)
namba val3 = nasibu(100)
namba val4 = nasibu(100)
andika("Second run:", val3, val4)  # Same as first run!
```

#### Practical Applications

**Dice Simulator:**
```swahili
kazi kufa_sita() {
    rudisha nasibu(1, 6)  # Roll 1-6
}

kazi kufa_kumi_ishirini() {
    rudisha nasibu(1, 20)  # Roll 1-20
}

namba roll1 = kufa_sita()
namba roll2 = kufa_sita()
andika("Rolling two six-sided dice:", roll1, "+", roll2, "=", roll1 + roll2)
```

**Shuffle Array:**
```swahili
kazi kweneza_orodha(orodha arr) {
    namba i = urefu_orodha(arr) - 1
    
    wakati i > 0 {
        namba j = nasibu(0, i)
        
        # Swap arr[i] and arr[j]
        namba temp = pata(arr, i)
        arr = ondoa(arr, i)
        ongeza(arr, temp)
        
        i = i - 1
    }
}
```

**Random Color Generator:**
```swahili
kazi rangi_nasibu() {
    namba r = nasibu(0, 256)
    namba g = nasibu(0, 256)
    namba b = nasibu(0, 256)
    
    rudisha "#" + maneno_namba(r) + maneno_namba(g) + maneno_namba(b)
}
```

**Sampling from List:**
```swahili
orodha fruits = ["apple", "banana", "orange", "mango"]

kazi sampuli(orodha list) {
    namba index = nasibu(urefu_orodha(list))
    rudisha pata(list, index)
}

andika("Random fruit:", sampuli(fruits))
```

#### Error Handling

```swahili
# Error: nasibu with invalid range
jaribu {
    namba bad = nasibu(10, 10)  # min = max, invalid
} shika (error) {
    andika("Error: min must be less than max")
}

# Error: nasibu with non-positive number (1 arg)
jaribu {
    namba bad = nasibu(0)       # max must be positive
} shika (error) {
    andika("Error: max must be positive")
}
```

#### Comparison: Random Function Modes

| Function | Usage | Range | Type |
|----------|-------|-------|------|
| nasibu() | Random float | [0.0, 1.0) | float |
| nasibu(n) | Random 0 to n | [0, n) | integer |
| nasibu(min, max) | Random in range | [min, max] | integer |
| weka_mbegu(seed) | Set seed | - | void |


