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
| `leta` | import | Import a module file |
| `jaribu` | try | Try block for error handling |
| `shika` | catch | Catch block for handling errors |
| `hatimaye` | finally | Finally block (always executes) |
| `tupa` | throw | Throw an error |
| `kamusi` | dictionary/map | Declare a dictionary variable |
| `darasa` | class | Define a class |
| `unda` | new/create | Create a class instance |
| `hii` | this/self | Reference to current instance |

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

### Comments

Kwenda supports single-line comments using the `#` character. Comments can appear:
- At the start of a line
- At the end of a line (inline comments)
- Anywhere in your code

```swahili
# This is a comment at the start of a line

kazi kuu() {
    # Comment inside a function
    namba x = 10  # Inline comment after code
    
    # Multiple consecutive comments
    # can be used to create
    # comment blocks
    
    andika("Hello")  # Print greeting
}

# Comment at the end of the file
```

### Dictionaries (Maps)

Kwenda supports dictionary/map data structures for key-value storage:

```swahili
# Create dictionary
kamusi person = {"jina": "Amina", "umri": 25, "mji": "Dar es Salaam"}

# Access values
maneno name = person["jina"]
namba age = person["umri"]

# Modify values
person["umri"] = 26

# Add new keys
person["kazi"] = "Mwalimu"

# Iterate through properties (access by key)
andika("Name:", person["jina"])
andika("Age:", person["umri"])
andika("City:", person["mji"])
andika("Job:", person["kazi"])
```

### Object-Oriented Programming

#### Class Syntax

Define classes using the `darasa` keyword:

```swahili
darasa Mtu {
    # Property declarations
    maneno jina
    namba umri
    maneno mji
    
    # Constructor (special method named 'unda')
    kazi unda(maneno j, namba u, maneno m) {
        hii.jina = j      # 'hii' refers to current instance (this/self)
        hii.umri = u
        hii.mji = m
    }
    
    # Methods
    kazi salamu() {
        andika("Habari! Jina langu ni", hii.jina)
    }
    
    kazi siku_ya_kuzaliwa() {
        hii.umri = hii.umri + 1
        andika("Happy Birthday! You are now", hii.umri)
    }
}
```

#### Creating Instances

Use the `unda` keyword to create class instances:

```swahili
# Create instances
kamusi mtu1 = unda Mtu("Amina", 25, "Dar es Salaam")
kamusi mtu2 = unda Mtu("Juma", 30, "Arusha")

# Access properties
andika("Name:", mtu1["jina"])
andika("Age:", mtu1["umri"])

# Modify properties
mtu1["umri"] = 26
```

#### The `hii` Keyword

Inside class methods, use `hii` to reference the current instance:

```swahili
darasa Counter {
    namba count
    
    kazi unda() {
        hii.count = 0
    }
    
    kazi ongeza() {
        hii.count = hii.count + 1
        andika("Count:", hii.count)
    }
}
```

### Module System

#### Importing Modules
```swahili
# Import a module file
leta "modules/math.swh"
leta "modules/strings.swh"

# Use module functions with namespace
namba result = math.ongeza_kubwa(10, 5)
maneno greeting = strings.salamu("Amina")
```

#### Creating Modules
```swahili
# File: modules/mymodule.swh
# Define functions that will be available to importers
kazi my_function(namba x) {
    rudisha x * 2
}

# Module-level variables
namba MY_CONSTANT = 42
```

#### Module Namespaces
- Each module has its own namespace
- Functions and variables are accessed using dot notation: `module.function()`
- Module names are derived from the filename (without `.swh` extension)
- Modules are cached - importing the same module multiple times loads it only once

### Error Handling

#### Try-Catch Blocks
```swahili
jaribu {
    # Code that might throw an error
    namba x = pata(arr, 100)  # Index out of bounds
} shika (error) {
    # Handle the error
    andika("Error occurred:", error)
}
```

#### Try-Catch-Finally
```swahili
jaribu {
    # Try some operation
    maneno content = soma("file.txt")
} shika (e) {
    # Handle error
    andika("Could not read file:", e)
} hatimaye {
    # Always executes, even if no error
    andika("Cleanup complete")
}
```

#### Throwing Errors
```swahili
kazi validate(namba age) {
    kama age < 0 {
        tupa "Age cannot be negative"
    }
    kama age > 150 {
        tupa "Age is unrealistic"
    }
    rudisha kweli
}
```

### Improved Error Messages

Kwenda provides detailed, bilingual error messages with context to help you debug your programs quickly:

#### Error Message Features
- **Bilingual Messages**: Errors shown in both Swahili and English
- **Contextual Information**: Detailed explanation of what went wrong
- **Helpful Suggestions**: Guidance on how to fix the error
- **Beautiful Formatting**: Professional error display with Unicode box drawing

#### Common Error Types

**Array Index Out of Bounds:**
```swahili
orodha namba nums = [1, 2, 3]
namba value = pata(nums, 10)  # Index 10 is out of bounds
```
Error output:
```
Index 10 ni nje ya mipaka ya orodha (urefu: 3)
Katika kazi 'pata': Jaribu kutumia index kati ya 0 na 2
```

**Division by Zero:**
```swahili
leta "modules/math.swh"
namba result = math.gawanya(100, 0)  # Cannot divide by zero
```
Error output:
```
╔═══════════════════════════════════════════════════════════╗
║ HITILAFU (ERROR)                                          ║
╚═══════════════════════════════════════════════════════════╝
Ujumbe: Haiwezekani kugawanya na sifuri (Cannot divide by zero)
```

**File Not Found:**
```swahili
maneno content = soma("missing.txt")  # File doesn't exist
```
Error output:
```
Hitilafu ya kusoma faili 'missing.txt': open missing.txt: no such file or directory
Katika kazi 'soma': Hakikisha faili ipo na una ruhusa ya kusoma
```

**Type Mismatch:**
```swahili
maneno not_array = "This is a string"
namba value = pata(not_array, 0)  # Wrong type
```
Error output:
```
Hii si orodha
Katika kazi 'pata': Argument ya kwanza lazima iwe orodha
```

#### Error Handling Best Practices

1. **Use try-catch for expected errors:**
```swahili
jaribu {
    maneno content = soma("config.txt")
    andika("Config loaded:", content)
} shika (error) {
    andika("Using default config due to error:", error)
    # Provide fallback behavior
}
```

2. **Validate input before operations:**
```swahili
kazi safe_divide(namba a, namba b) {
    kama b == 0 {
        tupa "Haiwezekani kugawanya na sifuri"
    }
    rudisha a / b
}
```

3. **Use finally for cleanup:**
```swahili
jaribu {
    # Open resources
    maneno data = soma("data.txt")
} shika (error) {
    andika("Error reading file:", error)
} hatimaye {
    # Always cleanup, even if error occurred
    andika("Cleanup complete")
}
```

## 📚 Examples

> **📖 For comprehensive examples and tutorials, see [EXAMPLES.md](EXAMPLES.md)**
> 
> **🚀 Quick Start**: All example files are in the `examples/` directory. Run with: `go run main.go examples/filename.swh`

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
- Nested function calls as arguments not fully supported (use intermediate variables)
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
- [ ] Lambda functions
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