# Const/Immutable Variables in Kwenda

## Overview

Kwenda supports immutable (constant) variables using the `thabiti` keyword (Swahili for "firm" or "constant"). Once a const variable is declared and initialized, it cannot be reassigned.

## Syntax

```swahili
thabiti <type> <name> = <value>
```

## Supported Types

Const declarations work with all Kwenda data types:

### Basic Types

```swahili
// Numbers
thabiti namba PI = 3.14159
thabiti namba MAX_SIZE = 1000

// Strings
thabiti maneno APP_NAME = "Kwenda"
thabiti maneno VERSION = "1.0.0"

// Booleans
thabiti boolean DEBUG_MODE = uwongo
thabiti boolean IS_PRODUCTION = kweli
```

### Collection Types

```swahili
// Arrays
thabiti orodha namba PRIMES = [2, 3, 5, 7, 11]

// Dictionaries
thabiti kamusi CONFIG = {"host": "localhost", "port": 8080}

// Sets
thabiti seti namba WINNING_NUMBERS = {7, 14, 21, 28}

// Tuples
thabiti tuple namba COORDINATES = (10, 20, 30)
```

## Const Protection

Attempting to reassign a const variable will result in an error:

```swahili
thabiti namba PI = 3.14159
PI = 3.14  // Error: Cannot reassign const variable 'PI'
```

### Error Message

```
Hitilafu: Hauwezi kubadilisha thamani ya kigezo thabiti 'PI' (Error: Cannot reassign const variable 'PI')
```

## Mutable vs Const Comparison

```swahili
// Mutable variable - can be changed
namba x = 10
andika("x =", x)        // Output: x = 10
x = 20
andika("x =", x)        // Output: x = 20

// Const variable - cannot be changed
thabiti namba y = 30
andika("y =", y)        // Output: y = 30
y = 40                  // Error: Cannot reassign const variable
andika("y =", y)        // Output: y = 30 (unchanged)
```

## Use Cases

### 1. Mathematical Constants

```swahili
thabiti namba PI = 3.14159
thabiti namba E = 2.71828
thabiti namba GOLDEN_RATIO = 1.618

kazi calculateCircleArea(namba radius) namba {
    rudisha PI * radius * radius
}
```

### 2. Configuration Values

```swahili
thabiti namba MAX_USERS = 1000
thabiti namba TIMEOUT_SECONDS = 30
thabiti maneno DEFAULT_LANGUAGE = "Swahili"

kazi showConfig() {
    andika("Max Users:", MAX_USERS)
    andika("Timeout:", TIMEOUT_SECONDS, "seconds")
    andika("Default Language:", DEFAULT_LANGUAGE)
}
```

### 3. Error Codes

```swahili
thabiti namba ERROR_NOT_FOUND = 404
thabiti namba ERROR_UNAUTHORIZED = 401
thabiti namba SUCCESS = 200

kazi getStatusMessage(namba code) maneno {
    kama code == SUCCESS {
        rudisha "Mafanikio (Success)"
    } sivyo kama code == ERROR_NOT_FOUND {
        rudisha "Haipatikani (Not Found)"
    } sivyo {
        rudisha "Hitilafu (Error)"
    }
}
```

### 4. Limits and Boundaries

```swahili
thabiti namba MIN_AGE = 0
thabiti namba MAX_AGE = 150
thabiti namba PASSING_GRADE = 50

kazi isValidAge(namba age) boolean {
    rudisha age >= MIN_AGE na age <= MAX_AGE
}

kazi hasPassed(namba score) boolean {
    rudisha score >= PASSING_GRADE
}
```

### 5. Multipliers and Conversions

```swahili
thabiti namba HOURS_PER_DAY = 24
thabiti namba MINUTES_PER_HOUR = 60
thabiti namba SECONDS_PER_MINUTE = 60

thabiti namba SECONDS_PER_DAY = HOURS_PER_DAY * MINUTES_PER_HOUR * SECONDS_PER_MINUTE

andika("Seconds in a day:", SECONDS_PER_DAY)  // Output: 86400
```

### 6. Const in Functions

```swahili
kazi calculateArea(namba width, namba height) namba {
    thabiti namba SCALE_FACTOR = 1.5
    rudisha width * height * SCALE_FACTOR
}
```

### 7. Const Arrays for Reference Data

```swahili
thabiti orodha namba FIBONACCI = [1, 1, 2, 3, 5, 8, 13, 21, 34]
thabiti orodha maneno WEEKDAYS = ["Jumatatu", "Jumanne", "Jumatano", "Alhamisi", "Ijumaa"]

// Note: Array reference is const, but elements can still be modified
// This is standard behavior in most programming languages
```

## Best Practices

### 1. Use UPPER_CASE for Constants

```swahili
thabiti namba MAX_SIZE = 100        // Good
thabiti namba max_size = 100        // Works, but not conventional
```

### 2. Declare Constants at the Top

```swahili
// Constants first
thabiti namba PI = 3.14159
thabiti namba E = 2.71828

// Then regular variables
namba radius = 5
namba area = PI * radius * radius
```

### 3. Use Const for Values That Should Never Change

```swahili
// Good use cases
thabiti namba MAX_RETRIES = 3
thabiti maneno API_VERSION = "v1.0"
thabiti boolean ENABLE_LOGGING = kweli

// Don't use const for values that need to change
namba currentRetries = 0        // Mutable (needs to increment)
maneno userInput = ""           // Mutable (gets user data)
```

### 4. Const for Computed Constants

```swahili
thabiti namba BASE_PRICE = 100
thabiti namba TAX_RATE = 0.15
thabiti namba TOTAL_PRICE = BASE_PRICE * (1 + TAX_RATE)
```

## Implementation Details

### AST Representation

Const variables are marked with an `IsConst` flag in the AST:

```go
type VariableDeclarationNode struct {
    Name    string
    Value   ASTNode
    IsConst bool      // true for thabiti declarations
    VarType string
}
```

### Runtime Protection

At runtime, const variables are tracked in a separate map in the environment:

```go
type Environment struct {
    Variables map[string]interface{}
    ConstVars map[string]bool         // Tracks which vars are const
    Functions map[string]ast.FunctionNode
    Parent    *Environment
}
```

When attempting to reassign a variable, the interpreter checks if it's const:

```go
func (env *Environment) Set(name string, value interface{}) {
    if env.IsConst(name) {
        fmt.Printf("Hitilafu: Hauwezi kubadilisha thamani ya kigezo thabiti '%s'\n", name)
        return
    }
    // Update variable...
}
```

### Scoping

Const variables follow the same scoping rules as regular variables:

```swahili
thabiti namba GLOBAL_CONST = 100

kazi myFunction() {
    thabiti namba LOCAL_CONST = 50
    andika(GLOBAL_CONST)    // Can access global const
    andika(LOCAL_CONST)     // Can access local const
}

andika(LOCAL_CONST)  // Error: LOCAL_CONST not defined in this scope
```

## Examples

See the following example files for comprehensive demonstrations:

- `examples/const_simple.swh` - Basic const variable tests
- `examples/const_demo.swh` - Comprehensive const usage examples

## Comparison with Other Languages

### JavaScript (ES6+)
```javascript
const PI = 3.14159;  // Similar to thabiti namba PI = 3.14159
```

### Python
```python
PI = 3.14159  # Python doesn't have true const (convention: UPPER_CASE)
```

### Rust
```rust
const PI: f64 = 3.14159;  // Similar to thabiti namba PI = 3.14159
```

### Go
```go
const PI = 3.14159  // Similar to thabiti namba PI = 3.14159
```

## Summary

- **Keyword**: `thabiti` (Swahili for "firm/constant")
- **Purpose**: Declare immutable variables that cannot be reassigned
- **Syntax**: `thabiti <type> <name> = <value>`
- **Types**: Works with all Kwenda data types
- **Protection**: Runtime error when attempting to reassign
- **Scoping**: Follows standard variable scoping rules
- **Best Practice**: Use UPPER_CASE naming convention

Const variables provide safety and clarity by preventing accidental reassignment of values that should remain constant throughout program execution.
