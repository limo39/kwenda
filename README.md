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

### Mathematical Rounding Operations

Kwenda provides four built-in functions for rounding numbers, each with different behavior:

#### Basic Usage

```swahili
# Floor - round down to nearest integer
namba x = chini(4.7)    # x = 4
namba y = chini(-2.3)   # y = -3 (rounds toward negative infinity)

# Ceiling - round up to nearest integer  
namba a = juu(4.2)      # a = 5
namba b = juu(-2.7)     # b = -2 (rounds toward positive infinity)

# Round - round to nearest integer
namba p = zunguka(4.5)  # p = 5 (rounds away from zero for .5)
namba q = zunguka(4.3)  # q = 4

# Truncate - remove decimal part (round toward zero)
namba m = kata_desimali(4.9)   # m = 4
namba n = kata_desimali(-4.9)  # n = -4 (rounds toward zero)
```

#### Key Differences for Negative Numbers

The main difference between these functions becomes apparent with negative numbers:

```swahili
namba value = -2.7

andika("chini(-2.7) =", chini(value))           # -3 (rounds DOWN)
andika("kata_desimali(-2.7) =", kata_desimali(value))  # -2 (rounds toward ZERO)
andika("juu(-2.7) =", juu(value))               # -2 (rounds UP)
andika("zunguka(-2.7) =", zunguka(value))       # -3 (rounds to nearest)
```

**Comparison Table:**

| Function | 4.7 | -2.7 | Description |
|----------|-----|------|-------------|
| `chini` (floor) | 4 | -3 | Always rounds toward negative infinity |
| `juu` (ceil) | 5 | -2 | Always rounds toward positive infinity |
| `zunguka` (round) | 5 | -3 | Rounds to nearest integer |
| `kata_desimali` (truncate) | 4 | -2 | Always rounds toward zero |

#### Using with Math Module Constants

```swahili
leta "modules/math.swh"

kazi kuu() {
    andika("PI =", math.PI)                  # 3.14159265359
    andika("chini(PI) =", chini(math.PI))    # 3
    andika("juu(PI) =", juu(math.PI))        # 4
    andika("zunguka(PI) =", zunguka(math.PI)) # 3
    
    andika("E =", math.E)                    # 2.71828182846
    andika("chini(E) =", chini(math.E))      # 2
    andika("juu(E) =", juu(math.E))          # 3
}
```

#### IEEE 754 Rounding Behavior

The `zunguka` (round) function follows IEEE 754 standard rounding:
- Values like 4.5 round to 5 (away from zero)
- The rounding is consistent and predictable for all edge cases

### Mathematical Root Operations

Kwenda provides three built-in functions for calculating roots, each optimized for different use cases:

#### Basic Usage

```swahili
# Square root - mzizi_mraba(number)
namba x = mzizi_mraba(16)     # x = 4
namba y = mzizi_mraba(2)      # y = 1.4142135623730951
namba z = mzizi_mraba(100)    # z = 10

# Cube root - mzizi_mchemraba(number)
namba a = mzizi_mchemraba(27)    # a = 3
namba b = mzizi_mchemraba(-8)    # b = -2 (works with negatives!)
namba c = mzizi_mchemraba(10)    # c = 2.154434690031884

# Nth root - mzizi(number, root)
namba p = mzizi(16, 2)    # p = 4 (square root)
namba q = mzizi(8, 3)     # q = 2 (cube root)
namba r = mzizi(32, 5)    # r = 2 (5th root)
namba s = mzizi(-8, 3)    # s = -2 (odd roots of negatives work)
```

#### Handling Negative Numbers

Root functions have different behaviors with negative numbers:

```swahili
# Square root - ONLY non-negative numbers
namba valid = mzizi_mraba(25)      # 5 - OK
# namba error = mzizi_mraba(-25)   # ERROR! Cannot calculate

# Cube root - ANY number (positive or negative)
namba pos = mzizi_mchemraba(27)    # 3
namba neg = mzizi_mchemraba(-27)   # -3 - Works fine!

# Nth root - depends on whether root is even or odd
namba odd_root = mzizi(-8, 3)      # -2 - Odd root, works
# namba even_root = mzizi(-16, 4)  # ERROR! Even root of negative
```

**Error Handling:**

```swahili
jaribu {
    namba x = mzizi_mraba(-4)  # Will throw error
} shika (error) {
    andika("Cannot calculate square root of negative number")
}

jaribu {
    namba y = mzizi(-16, 4)    # Will throw error (even root)
} shika (error) {
    andika("Cannot calculate even root of negative number")
}

# This works fine - odd root of negative
namba z = mzizi(-27, 3)        # -3 (no error)
```

#### Comparison of Root Functions

**When to use each function:**

| Function | Use Case | Negative Numbers | Performance |
|----------|----------|------------------|-------------|
| `mzizi_mraba` | Square roots only | ❌ Not allowed | Fastest |
| `mzizi_mchemraba` | Cube roots only | ✅ Allowed | Fast |
| `mzizi` | Any root (2nd, 3rd, 4th, etc.) | ⚠️ Odd roots only | Flexible |

```swahili
# These are equivalent:
namba a = mzizi_mraba(16)      # 4
namba b = mzizi(16, 2)         # 4 (same result)

# These are equivalent:
namba c = mzizi_mchemraba(8)   # 2
namba d = mzizi(8, 3)          # 2 (same result)

# Only mzizi can do 5th, 7th, 10th roots, etc:
namba e = mzizi(32, 5)         # 2 (5th root)
namba f = mzizi(128, 7)        # 2 (7th root)
```

#### Using with Math Module Constants

```swahili
leta "modules/math.swh"

kazi kuu() {
    andika("PI =", math.PI)                        # 3.14159265359
    andika("Square root of PI =", mzizi_mraba(math.PI))      # 1.772...
    andika("Cube root of PI =", mzizi_mchemraba(math.PI))    # 1.464...
    andika("4th root of PI =", mzizi(math.PI, 4))            # 1.331...
    
    andika("E =", math.E)                          # 2.71828182846
    andika("Square root of E =", mzizi_mraba(math.E))        # 1.648...
}
```

#### Combining Roots with Rounding

```swahili
# Calculate and round in one expression
namba sqrt2 = mzizi_mraba(2)           # 1.4142135623730951
andika("Rounded:", zunguka(sqrt2))     # 1

# Or directly
andika("Floor of sqrt(10):", chini(mzizi_mraba(10)))     # 3
andika("Ceil of cbrt(10):", juu(mzizi_mchemraba(10)))    # 3

# Using in calculations
namba a = mzizi_mraba(16) + mzizi_mraba(9)   # 4 + 3 = 7
namba b = mzizi_mraba(16) * mzizi_mraba(9)   # 4 * 3 = 12
```

#### Special Cases

```swahili
# Root of 0 is always 0
andika(mzizi_mraba(0))        # 0
andika(mzizi_mchemraba(0))    # 0

# Root of 1 is always 1
andika(mzizi_mraba(1))        # 1
andika(mzizi_mchemraba(1))    # 1
andika(mzizi(1, 5))           # 1
andika(mzizi(1, 100))         # 1

# Perfect roots return integers
andika(mzizi_mraba(25))       # 5 (not 5.0)
andika(mzizi_mchemraba(8))    # 2 (not 2.0)

# Non-perfect roots return floats
andika(mzizi_mraba(2))        # 1.4142135623730951
```

### Exponential & Logarithm Operations

Kwenda provides four built-in functions for exponential and logarithmic calculations, essential for scientific computing, growth calculations, and data analysis:

#### Basic Usage

```swahili
# Exponential - exp(x) returns e^x
namba a = exp(0)      # a = 1
namba b = exp(1)      # b = 2.718281828... (which is e)
namba c = exp(2)      # c = 7.389...
namba d = exp(-1)     # d = 0.367... (which is 1/e)

# Natural logarithm - log_asili(x) returns ln(x)
namba p = log_asili(1)      # p = 0
namba q = log_asili(2)      # q = 0.693...
namba r = log_asili(10)     # r = 2.302...
namba s = log_asili(100)    # s = 4.605...

# Base-10 logarithm - log10(x)
namba m = log10(1)      # m = 0
namba n = log10(10)     # n = 1
namba o = log10(100)    # o = 2
namba p = log10(1000)   # p = 3

# Base-2 logarithm - log2(x)
namba w = log2(1)       # w = 0
namba x = log2(2)       # x = 1
namba y = log2(8)       # y = 3 (because 2^3 = 8)
namba z = log2(16)      # z = 4 (because 2^4 = 16)
```

#### Inverse Relationships

Exponential and logarithm functions are inverses of each other:

```swahili
# exp and log_asili are inverses
andika(exp(log_asili(5)))       # 5
andika(log_asili(exp(3)))       # 3

# This means:
# If y = exp(x), then x = log_asili(y)
# If y = log_asili(x), then x = exp(y)

# Example: solving exponential equations
namba x = log_asili(10)         # What power of e gives 10?
andika("e^", x, "=", exp(x))    # Verify: e^x = 10
```

#### Logarithm Properties

```swahili
# Property 1: log(a * b) = log(a) + log(b)
namba a = 8
namba b = 16
andika(log_asili(a * b))                    # 4.852...
andika(log_asili(a) + log_asili(b))         # 4.852... (same)

# Property 2: log(a / b) = log(a) - log(b)
andika(log_asili(a) - log_asili(b))         # -0.693...

# Property 3: log(a^n) = n * log(a)
namba n = 3
andika(log_asili(a * a * a))                # 6.238...
andika(n * log_asili(a))                    # 6.238... (same)
```

#### Error Handling

Logarithms only work with positive numbers:

```swahili
# Valid: positive numbers
namba x = log_asili(5)      # 1.609...
namba y = log10(100)        # 2
namba z = log2(8)           # 3

# Invalid: zero or negative numbers
jaribu {
    namba bad1 = log_asili(0)       # ERROR
} shika (error) {
    andika("Cannot take log of zero")
}

jaribu {
    namba bad2 = log10(-5)          # ERROR
} shika (error) {
    andika("Cannot take log of negative")
}

# Exponential overflow
jaribu {
    namba big = exp(710)            # ERROR (too large)
} shika (error) {
    andika("Result would overflow")
}
```

#### Using with Math Constants

```swahili
leta "modules/math.swh"

kazi kuu() {
    # E is Euler's number (base of natural log)
    andika("E =", math.E)                   # 2.71828...
    andika("exp(1) =", exp(1))              # 2.71828... (same as E)
    andika("log_asili(E) =", log_asili(math.E))  # 1
    
    # PI with exponential/log
    andika("PI =", math.PI)                 # 3.14159...
    andika("exp(PI) =", exp(math.PI))       # 23.140...
    andika("log_asili(PI) =", log_asili(math.PI))  # 1.144...
}
```

#### Practical Applications

**Compound Growth:**
```swahili
# Continuous compound interest: A = P * e^(rt)
namba principal = 1000
namba rate = 0.05              # 5% annual rate
namba time = 10                # 10 years

namba amount = principal * exp(rate * time)
andika("After 10 years:", amount)  # $1648.72
```

**Half-Life Calculations:**
```swahili
# Half-life: time for quantity to reduce to half
namba half_life = log_asili(0.5)
andika("Time to halve:", half_life)  # -0.693...

# Double-time: time for quantity to double
namba double_time = log_asili(2)
andika("Time to double:", double_time)  # 0.693...
```

**Powers of 2 (Computer Science):**
```swahili
# How many bits needed to represent a number?
namba value = 1024
namba bits = juu(log2(value))
andika("Bits needed for", value, ":", bits)  # 10 bits
```

**Orders of Magnitude:**
```swahili
# How many digits in a number?
namba num = 1000000
namba digits = juu(log10(num)) + 1
andika("Digits in", num, ":", digits)  # 7 digits
```

#### Change of Base Formula

Convert between different logarithm bases:

```swahili
# log_b(x) = log_asili(x) / log_asili(b)

# Example: Calculate log base 5 of 125
namba x = 125
namba base = 5
namba result = log_asili(x) / log_asili(base)
andika("log_5(125) =", result)  # 3 (because 5^3 = 125)

# Verify with our log functions:
# log2(x) = log_asili(x) / log_asili(2)
namba test = 8
andika("log2(8) direct:", log2(test))           # 3
andika("Using formula:", log_asili(test) / log_asili(2))  # 3 (same)
```

#### Combining with Other Math Functions

```swahili
# Exponential with rounding
namba exp_val = exp(2.5)
andika("exp(2.5) =", exp_val)               # 12.182...
andika("Rounded:", zunguka(exp_val))        # 12

# Logarithm with roots
namba num = 16
andika("log2(16) =", log2(num))             # 4
andika("sqrt(16) =", mzizi_mraba(num))      # 4
andika("log2(sqrt(16)) =", log2(mzizi_mraba(num)))  # 2

# Exponential growth with roots
namba growth = exp(log_asili(2))            # 2
andika("e^(ln(2)) =", growth)
```

#### Comparison of Functions

| Function | Domain | Range | Use Case |
|----------|--------|-------|----------|
| `exp(x)` | All numbers | Positive only | Growth, compound interest |
| `log_asili(x)` | Positive only | All numbers | Natural processes, calculus |
| `log10(x)` | Positive only | All numbers | Orders of magnitude, pH, decibels |
| `log2(x)` | Positive only | All numbers | Computer science, information theory |

**Key Relationships:**
- `exp(log_asili(x)) = x` (for x > 0)
- `log_asili(exp(x)) = x` (for all x)
- `log10(x) = log_asili(x) / log_asili(10)`
- `log2(x) = log_asili(x) / log_asili(2)`

### Trigonometric Functions

Kwenda provides comprehensive trigonometric functions for angle calculations, including basic trig functions (sin, cos, tan), their inverses (asin, acos, atan, atan2), and angle conversion utilities (radians, degrees).

**Important:** All trigonometric functions work with angles in **radians**, not degrees. Use the `radians()` and `degrees()` functions to convert between the two.

#### Angle Conversion

```swahili
# Convert degrees to radians
namba angle_rad = radians(90)      # π/2 ≈ 1.5708
andika("90° = ", angle_rad, " radians")

# Convert radians to degrees
namba angle_deg = degrees(1.5708)  # ≈ 90
andika("1.5708 radians = ", angle_deg, "°")

# Common conversions
andika("0° = ", radians(0), " rad")       # 0
andika("30° = ", radians(30), " rad")     # π/6 ≈ 0.5236
andika("45° = ", radians(45), " rad")     # π/4 ≈ 0.7854
andika("60° = ", radians(60), " rad")     # π/3 ≈ 1.0472
andika("90° = ", radians(90), " rad")     # π/2 ≈ 1.5708
andika("180° = ", radians(180), " rad")   # π ≈ 3.1416
andika("360° = ", radians(360), " rad")   # 2π ≈ 6.2832
```

#### Basic Trigonometric Functions

```swahili
# Sine - sin(angle_in_radians)
andika("sin(0) = ", sin(0))                       # 0
andika("sin(30°) = ", sin(radians(30)))           # 0.5
andika("sin(45°) = ", sin(radians(45)))           # √2/2 ≈ 0.7071
andika("sin(90°) = ", sin(radians(90)))           # 1

# Cosine - cos(angle_in_radians)
andika("cos(0) = ", cos(0))                       # 1
andika("cos(60°) = ", cos(radians(60)))           # 0.5
andika("cos(45°) = ", cos(radians(45)))           # √2/2 ≈ 0.7071
andika("cos(90°) = ", cos(radians(90)))           # 0 (approximately)

# Tangent - tan(angle_in_radians)
andika("tan(0) = ", tan(0))                       # 0
andika("tan(45°) = ", tan(radians(45)))           # 1
andika("tan(30°) = ", tan(radians(30)))           # √3/3 ≈ 0.5774
andika("tan(60°) = ", tan(radians(60)))           # √3 ≈ 1.7321
```

#### Inverse Trigonometric Functions

Inverse trig functions return angles in **radians**:

```swahili
# Arc sine - asin(x) returns angle where sin(angle) = x
# Domain: -1 ≤ x ≤ 1, Range: -π/2 to π/2
namba angle1 = asin(0.5)                          # π/6 ≈ 0.5236 radians (30°)
namba angle2 = asin(1)                            # π/2 ≈ 1.5708 radians (90°)
andika("asin(0.5) = ", degrees(angle1), "°")      # 30°

# Arc cosine - acos(x) returns angle where cos(angle) = x
# Domain: -1 ≤ x ≤ 1, Range: 0 to π
namba angle3 = acos(0.5)                          # π/3 ≈ 1.0472 radians (60°)
namba angle4 = acos(0)                            # π/2 ≈ 1.5708 radians (90°)
andika("acos(0.5) = ", degrees(angle3), "°")      # 60°

# Arc tangent - atan(x) returns angle where tan(angle) = x
# Domain: all real numbers, Range: -π/2 to π/2
namba angle5 = atan(1)                            # π/4 ≈ 0.7854 radians (45°)
namba angle6 = atan(0)                            # 0 radians (0°)
andika("atan(1) = ", degrees(angle5), "°")        # 45°

# Two-argument arc tangent - atan2(y, x) for proper quadrant handling
# Returns angle from origin to point (x, y)
namba angle7 = atan2(1, 1)                        # π/4 ≈ 0.7854 (Quadrant I)
namba angle8 = atan2(1, -1)                       # 3π/4 ≈ 2.3562 (Quadrant II)
namba angle9 = atan2(-1, -1)                      # -3π/4 ≈ -2.3562 (Quadrant III)
namba angle10 = atan2(-1, 1)                      # -π/4 ≈ -0.7854 (Quadrant IV)
```

#### Pythagorean Identity

The fundamental relationship sin²(x) + cos²(x) = 1:

```swahili
namba angle = radians(30)
namba s = sin(angle)
namba c = cos(angle)

# Verify the identity
namba sum = s*s + c*c
andika("sin²(30°) + cos²(30°) = ", sum)  # Always equals 1
```

#### Practical Applications

**Right Triangle Calculations:**

```swahili
# Given: right triangle with hypotenuse = 10, angle = 30°
namba hypotenuse = 10
namba angle = radians(30)

# Calculate sides
namba opposite = hypotenuse * sin(angle)          # 5
namba adjacent = hypotenuse * cos(angle)          # 8.66

andika("Opposite side = ", opposite)
andika("Adjacent side = ", adjacent)

# Verify: opposite² + adjacent² = hypotenuse²
namba check = opposite*opposite + adjacent*adjacent
andika("Verification: ", check, " ≈ ", hypotenuse*hypotenuse)
```

**Finding Angle Between Two Points:**

```swahili
# Point coordinates
namba x = 4
namba y = 3

# Find angle from origin to point (x, y)
namba angle_rad = atan2(y, x)
namba angle_deg = degrees(angle_rad)

andika("Point (", x, ", ", y, ")")
andika("Angle from origin: ", angle_deg, "°")

# Calculate distance
namba distance = mzizi_mraba(x*x + y*y)  # 5
andika("Distance: ", distance)
```

**Navigation and Bearings:**

```swahili
# Calculate direction to travel from point A to point B
namba dx = 100  # East displacement
namba dy = 100  # North displacement

namba bearing_rad = atan2(dy, dx)
namba bearing_deg = degrees(bearing_rad)

andika("Bearing: ", bearing_deg, "° from East")

# Convert to compass bearing (from North)
namba compass_bearing = 90 - bearing_deg
andika("Compass bearing: ", compass_bearing, "° from North")
```

#### Error Handling

```swahili
# asin and acos have restricted domains
jaribu {
    namba bad = asin(2)  # ERROR: must be in [-1, 1]
} shika kosa {
    andika("Domain error: asin requires input between -1 and 1")
}

jaribu {
    namba bad = acos(1.5)  # ERROR: must be in [-1, 1]
} shika kosa {
    andika("Domain error: acos requires input between -1 and 1")
}

# tan is undefined at π/2 + nπ
jaribu {
    namba bad = tan(radians(90))  # May produce very large number or error
} shika kosa {
    andika("tan(90°) is undefined (approaches infinity)")
}
```

#### Combining with Other Math Functions

```swahili
# Rounding trig results
namba angle = radians(30)
andika("sin(30°) rounded:", zunguka(sin(angle)))  # 1

# Using with roots
namba a = sin(radians(30))  # 0.5
namba b = cos(radians(30))  # 0.866
namba hyp = mzizi_mraba(a*a + b*b)  # 1 (Pythagorean theorem)

# Converting between trig functions
# tan(x) = sin(x) / cos(x)
namba angle2 = radians(45)
namba tan_val = sin(angle2) / cos(angle2)
andika("tan(45°) = sin(45°)/cos(45°) = ", tan_val)  # 1
```

#### Comparison Table

| Function | Input | Output | Domain | Use Case |
|----------|-------|--------|--------|----------|
| `sin(x)` | Radians | Ratio [-1, 1] | All numbers | Height, oscillation |
| `cos(x)` | Radians | Ratio [-1, 1] | All numbers | Distance, projection |
| `tan(x)` | Radians | Any number | x ≠ π/2 + nπ | Slope, gradient |
| `asin(x)` | Ratio [-1, 1] | Radians [-π/2, π/2] | [-1, 1] | Finding angles |
| `acos(x)` | Ratio [-1, 1] | Radians [0, π] | [-1, 1] | Finding angles |
| `atan(x)` | Any number | Radians [-π/2, π/2] | All numbers | Finding angles |
| `atan2(y, x)` | Any numbers | Radians [-π, π] | All numbers | Direction, quadrants |
| `radians(deg)` | Degrees | Radians | All numbers | Angle conversion |
| `degrees(rad)` | Radians | Degrees | All numbers | Angle conversion |

**Key Relationships:**
- `sin(asin(x)) = x` (for -1 ≤ x ≤ 1)
- `cos(acos(x)) = x` (for -1 ≤ x ≤ 1)
- `tan(atan(x)) = x` (for all x)
- `sin²(x) + cos²(x) = 1` (Pythagorean identity)
- `tan(x) = sin(x) / cos(x)` (when cos(x) ≠ 0)
- Use `atan2(y, x)` instead of `atan(y/x)` to handle all quadrants correctly

### Hyperbolic Functions

Kwenda provides hyperbolic functions for advanced mathematical applications including calculus, physics, and engineering. These functions are analogs of trigonometric functions but relate to hyperbolas instead of circles.

#### Basic Hyperbolic Functions

```swahili
# Hyperbolic sine - sinh(x) = (e^x - e^-x) / 2
andika("sinh(0) = ", sinh(0))           # 0
andika("sinh(1) = ", sinh(1))           # 1.1752...
andika("sinh(2) = ", sinh(2))           # 3.6269...
andika("sinh(-1) = ", sinh(-1))         # -1.1752... (odd function)

# Hyperbolic cosine - cosh(x) = (e^x + e^-x) / 2
andika("cosh(0) = ", cosh(0))           # 1
andika("cosh(1) = ", cosh(1))           # 1.5431...
andika("cosh(2) = ", cosh(2))           # 3.7622...
andika("cosh(-1) = ", cosh(-1))         # 1.5431... (even function)

# Hyperbolic tangent - tanh(x) = sinh(x) / cosh(x)
andika("tanh(0) = ", tanh(0))           # 0
andika("tanh(1) = ", tanh(1))           # 0.7616...
andika("tanh(2) = ", tanh(2))           # 0.9640...
andika("tanh(10) = ", tanh(10))         # ~1 (approaches 1 as x → ∞)
andika("tanh(-10) = ", tanh(-10))       # ~-1 (approaches -1 as x → -∞)
```

#### Inverse Hyperbolic Functions

```swahili
# Inverse hyperbolic sine - asinh(x) = ln(x + sqrt(x² + 1))
# Domain: all real numbers
andika("asinh(0) = ", asinh(0))         # 0
andika("asinh(1) = ", asinh(1))         # 0.8814...
andika("asinh(2) = ", asinh(2))         # 1.4436...
andika("asinh(-1) = ", asinh(-1))       # -0.8814...

# Inverse hyperbolic cosine - acosh(x) = ln(x + sqrt(x² - 1))
# Domain: x >= 1
andika("acosh(1) = ", acosh(1))         # 0
andika("acosh(2) = ", acosh(2))         # 1.3170...
andika("acosh(5) = ", acosh(5))         # 2.2924...

# Inverse hyperbolic tangent - atanh(x) = 0.5 * ln((1+x)/(1-x))
# Domain: -1 < x < 1 (strictly between -1 and 1)
andika("atanh(0) = ", atanh(0))         # 0
andika("atanh(0.5) = ", atanh(0.5))     # 0.5493...
andika("atanh(-0.5) = ", atanh(-0.5))   # -0.5493...
andika("atanh(0.9) = ", atanh(0.9))     # 1.4722...
```

#### Inverse Relationships

```swahili
# Hyperbolic functions and their inverses
namba x = 2
andika("sinh(", x, ") = ", sinh(x))
andika("asinh(sinh(", x, ")) = ", asinh(sinh(x)))  # Returns original x

namba y = 1.5
andika("asinh(", y, ") = ", asinh(y))
andika("sinh(asinh(", y, ")) = ", sinh(asinh(y)))  # Returns original y

# Similar relationships for cosh/acosh and tanh/atanh
```

#### Hyperbolic Identities

**Fundamental Identity:** cosh²(x) - sinh²(x) = 1

```swahili
namba val = 2
namba sinh_val = sinh(val)
namba cosh_val = cosh(val)
namba identity = cosh_val * cosh_val - sinh_val * sinh_val
andika("cosh²(x) - sinh²(x) = ", identity)  # Always equals 1
```

**Relationship:** tanh(x) = sinh(x) / cosh(x)

```swahili
namba x = 1.5
namba tanh_direct = tanh(x)
namba tanh_computed = sinh(x) / cosh(x)
andika("tanh(x) = ", tanh_direct)
andika("sinh(x)/cosh(x) = ", tanh_computed)  # Same value
```

**Exponential Definitions:**

```swahili
# Verify definitions using exponential functions
namba x = 1
andika("sinh(x) = ", sinh(x))
andika("(e^x - e^-x)/2 = ", (exp(x) - exp(-x)) / 2)  # Same

andika("cosh(x) = ", cosh(x))
andika("(e^x + e^-x)/2 = ", (exp(x) + exp(-x)) / 2)  # Same
```

#### Symmetry Properties

```swahili
namba x = 2

# sinh is an odd function: sinh(-x) = -sinh(x)
andika("sinh(-x) = ", sinh(-x))
andika("-sinh(x) = ", -sinh(x))  # Equal

# cosh is an even function: cosh(-x) = cosh(x)
andika("cosh(-x) = ", cosh(-x))
andika("cosh(x) = ", cosh(x))    # Equal

# tanh is an odd function: tanh(-x) = -tanh(x)
andika("tanh(-x) = ", tanh(-x))
andika("-tanh(x) = ", -tanh(x))  # Equal
```

#### Practical Applications

**Catenary Curve (Hanging Cable):**

```swahili
# A cable hanging between two points forms a catenary curve
# Formula: y = a * cosh(x/a)
namba cable_a = 10  # Cable parameter
namba cable_x = 5   # Horizontal position
namba cable_y = cable_a * cosh(cable_x / cable_a)

andika("Cable height at x=", cable_x, ": y=", cable_y)
```

**Special Relativity (Rapidity):**

```swahili
# In special relativity, velocity relates to rapidity via tanh
# β = v/c = tanh(rapidity)
namba rapidity = 1
namba beta = tanh(rapidity)
andika("Velocity as fraction of light speed: ", beta)  # 0.7616...
```

**Hyperbolic Geometry:**

```swahili
# Hyperbolic distance in Poincaré disk model
namba r = 0.5  # Radial distance
namba distance = 2 * atanh(r)
andika("Hyperbolic distance: ", distance)
```

#### Error Handling

```swahili
# acosh requires x >= 1
jaribu {
    namba bad = acosh(0.5)  # ERROR: must be >= 1
} shika kosa {
    andika("Domain error: acosh requires input >= 1")
}

# atanh requires -1 < x < 1 (strictly)
jaribu {
    namba bad1 = atanh(1)   # ERROR: must be strictly < 1
} shika kosa {
    andika("Domain error: atanh requires -1 < x < 1")
}

jaribu {
    namba bad2 = atanh(-1)  # ERROR: must be strictly > -1
} shika kosa {
    andika("Domain error: atanh requires -1 < x < 1")
}

jaribu {
    namba bad3 = atanh(1.5) # ERROR: outside domain
} shika kosa {
    andika("Domain error: atanh input outside (-1, 1)")
}

# sinh and cosh can overflow for large |x|
jaribu {
    namba big = sinh(1000)  # ERROR: result too large
} shika kosa {
    andika("Overflow: result approaches infinity")
}
```

#### Combining with Other Functions

```swahili
# With rounding functions
andika("juu(sinh(1)) = ", juu(sinh(1)))          # 2
andika("chini(cosh(2)) = ", chini(cosh(2)))      # 3

# With exponential/logarithm
andika("exp(asinh(1)) = ", exp(asinh(1)))        # 2.4142...
andika("log_asili(cosh(2)) = ", log_asili(cosh(2)))  # 1.3250...

# With trigonometric (Gudermannian function)
# gd(x) = 2*atan(tanh(x/2))
namba x = 1
namba gud = 2 * atan(tanh(x / 2))
andika("Gudermannian of ", x, " = ", gud)
```

#### Asymptotic Behavior

```swahili
# tanh approaches ±1 for large |x|
andika("tanh(5) = ", tanh(5))     # 0.9999...
andika("tanh(10) = ", tanh(10))   # ~1.0
andika("tanh(-5) = ", tanh(-5))   # -0.9999...

# For large |x|, sinh(x) ≈ cosh(x) ≈ 0.5 * e^|x|
namba large_x = 5
andika("sinh(x) = ", sinh(large_x))
andika("cosh(x) = ", cosh(large_x))
andika("Ratio cosh/sinh = ", cosh(large_x) / sinh(large_x))  # ~1
```

#### Comparison Table

| Function | Domain | Range | Formula | Use Case |
|----------|--------|-------|---------|----------|
| `sinh(x)` | All numbers | All numbers | (e^x - e^-x)/2 | Wave equations, special relativity |
| `cosh(x)` | All numbers | [1, ∞) | (e^x + e^-x)/2 | Catenary curves, hanging cables |
| `tanh(x)` | All numbers | (-1, 1) | sinh(x)/cosh(x) | Activation functions, relativity |
| `asinh(x)` | All numbers | All numbers | ln(x + √(x²+1)) | Inverse of sinh |
| `acosh(x)` | [1, ∞) | [0, ∞) | ln(x + √(x²-1)) | Inverse of cosh |
| `atanh(x)` | (-1, 1) | All numbers | 0.5·ln((1+x)/(1-x)) | Inverse of tanh, logit function |

**Key Relationships:**
- `sinh(asinh(x)) = x` (for all x)
- `cosh(acosh(x)) = x` (for x ≥ 1)
- `tanh(atanh(x)) = x` (for -1 < x < 1)
- `cosh²(x) - sinh²(x) = 1` (fundamental hyperbolic identity)
- `tanh(x) = sinh(x) / cosh(x)`
- `sinh(x) = (e^x - e^-x) / 2`
- `cosh(x) = (e^x + e^-x) / 2`
- sinh is odd: `sinh(-x) = -sinh(x)`
- cosh is even: `cosh(-x) = cosh(x)`
- `tanh(x)` approaches ±1 as `x` approaches ±∞

### Number Properties & Utilities

Kwenda provides 14 built-in functions for checking number properties and performing common mathematical utilities. These functions help with number classification, finding extremes, and performing mathematical operations.

#### Number Property Checkers

These functions check properties of numbers and return boolean values:

```swahili
# Parity checks
andika("ni_shufwa(4) = ", ni_shufwa(4))         # kweli (even)
andika("ni_witiri(5) = ", ni_witiri(5))         # kweli (odd)

# Sign checks
andika("ni_chanya(5) = ", ni_chanya(5))         # kweli (positive)
andika("ni_hasi(-5) = ", ni_hasi(-5))           # kweli (negative)
andika("ni_sifuri(0) = ", ni_sifuri(0))         # kweli (zero)

# Prime check
andika("ni_namba_kuu(7) = ", ni_namba_kuu(7))   # kweli (7 is prime)
andika("ni_namba_kuu(4) = ", ni_namba_kuu(4))   # uwongo (4 is not prime)

# Perfect square check
andika("ni_mraba_kamili(9) = ", ni_mraba_kamili(9))      # kweli (9 = 3²)
andika("ni_mraba_kamili(10) = ", ni_mraba_kamili(10))    # uwongo (10 is not perfect square)
```

#### Utility Functions

These functions perform common operations on numbers:

```swahili
# Absolute value
andika("kiwango(-5) = ", kiwango(-5))           # 5

# Sign function - returns -1, 0, or 1
andika("ishara(-5) = ", ishara(-5))             # -1
andika("ishara(0) = ", ishara(0))               # 0
andika("ishara(5) = ", ishara(5))               # 1

# Maximum and minimum (support multiple arguments)
andika("max(3, 7, 2, 9) = ", max(3, 7, 2, 9))   # 9
andika("min(3, 7, 2, 9) = ", min(3, 7, 2, 9))   # 2
```

#### Mathematical Functions

```swahili
# Greatest Common Divisor
andika("kigawanyaji_kikuu(12, 8) = ", kigawanyaji_kikuu(12, 8))   # 4

# Least Common Multiple
andika("kigawanyaji_ndogo(12, 8) = ", kigawanyaji_ndogo(12, 8))   # 24

# Factorial
andika("ukweli(5) = ", ukweli(5))               # 120
andika("ukweli(0) = ", ukweli(0))               # 1

# Power
andika("tweza(2, 8) = ", tweza(2, 8))           # 256
andika("tweza(2, -1) = ", tweza(2, -1))         # 0.5
```

#### Practical Applications

**Checking Number Properties:**

```swahili
# Check if a number is prime and even
namba n = 7
kama (ni_namba_kuu(n)) && ni_shufwa(n) {
    andika(n, " is prime and even (impossible!)")
} sivyo {
    andika(n, " is not both prime and even")
}

# Find all primes up to 20
dor = 2
wakati (dor <= 20) {
    kama (ni_namba_kuu(dor)) {
        andika(dor, " is prime")
    }
    dor = dor + 1
}
```

**Fraction Simplification:**

```swahili
# Simplify a fraction using GCD
namba numerator = 36
namba denominator = 24
namba gcd_val = kigawanyaji_kikuu(numerator, denominator)

andika("Original: ", numerator, "/", denominator)
andika("Simplified: ", numerator / gcd_val, "/", denominator / gcd_val)  # 3/2
```

**Finding Common Multiples:**

```swahili
# Find LCM for scheduling
namba task1_interval = 6    # Every 6 days
namba task2_interval = 9    # Every 9 days
namba common_day = kigawanyaji_ndogo(task1_interval, task2_interval)

andika("Both tasks repeat every ", common_day, " days")  # 18 days
```

**Powers of Numbers:**

```swahili
# Calculate powers (useful for bit operations)
andika("2^10 = ", tweza(2, 10))                 # 1024 (max for 10 bits)
andika("Memory size (bytes) = ", tweza(2, 20)) # 1048576 (1 MB)

# Calculate roots using negative exponents
andika("Square root of 16 = ", tweza(16, 0.5))  # 4
andika("Cube root of 27 = ", tweza(27, 1/3))    # 3
```

**Extremes and Comparisons:**

```swahili
# Find min and max values
namba scores = [85, 92, 78, 88, 95]
andika("Highest score: ", max(85, 92, 78, 88, 95))    # 95
andika("Lowest score: ", min(85, 92, 78, 88, 95))     # 78

# Find absolute differences
namba a = -15
namba b = 8
andika("Distance between ", a, " and ", b, ": ", kiwango(a - b))  # 23
```

#### Error Handling

```swahili
# Factorial only works with non-negative integers
jaribu {
    namba bad = ukweli(-5)  # ERROR
} shika kosa {
    andika("Cannot compute factorial of negative number")
}

# Power function checks for overflow
jaribu {
    namba big = tweza(10, 1000)  # ERROR - result too large
} shika kosa {
    andika("Result would overflow")
}
```

#### Comparison Table

| Function | Swahili Name | Input | Output | Use Case |
|----------|--------------|-------|--------|----------|
| `ni_shufwa` | is_even | Number | Boolean | Check if divisible by 2 |
| `ni_witiri` | is_odd | Number | Boolean | Check if not divisible by 2 |
| `ni_chanya` | is_positive | Number | Boolean | Check if > 0 |
| `ni_hasi` | is_negative | Number | Boolean | Check if < 0 |
| `ni_sifuri` | is_zero | Number | Boolean | Check if == 0 |
| `ni_namba_kuu` | is_prime | Number | Boolean | Check primality (trial division) |
| `ni_mraba_kamili` | is_perfect_square | Number | Boolean | Check if perfect square |
| `kiwango` | abs | Number | Number | Absolute value |
| `ishara` | sign | Number | -1/0/1 | Sign of number |
| `max` | maximum | Numbers... | Number | Largest of multiple values |
| `min` | minimum | Numbers... | Number | Smallest of multiple values |
| `kigawanyaji_kikuu` | gcd | Two numbers | Number | Greatest common divisor |
| `kigawanyaji_ndogo` | lcm | Two numbers | Number | Least common multiple |
| `ukweli` | factorial | Non-neg int | Number | n! product |
| `tweza` | power | Base, exponent | Number | Base raised to exponent |

**Key Characteristics:**
- Property checkers return boolean (true/false)
- All functions preserve number types (int returns int when appropriate)
- `max` and `min` support any number of arguments
- `kigawanyaji_kikuu(0, n) = n` (GCD definition)
- Primes detected via trial division up to √n
- Factorial limited to prevent overflow
- Power function handles fractional exponents (e.g., `tweza(4, 0.5)` = 2)

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


## 🔧 Advanced Operations

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

