# Type Annotations and Type Checking in Kwenda

## Overview

Kwenda supports optional type annotations to provide better code safety and documentation. You can add type information to variables, function parameters, and return values. Type checking can be performed before execution to catch type-related errors early.

## Type System

### Basic Types

Kwenda supports the following basic types:

| Type | Keyword | Description | Example |
|------|---------|-------------|---------|
| Number | `namba` | Integer and floating-point numbers | `42`, `3.14` |
| String | `maneno` | Text strings | `"Habari"` |
| Boolean | `boolean` | True/false values | `kweli`, `uwongo` |
| Array | `orodha` | Ordered collections | `[1, 2, 3]` |
| Dictionary | `kamusi` | Key-value maps | `{"key": "value"}` |
| Set | `seti` | Unique element collections | `{1, 2, 3}` |
| Tuple | `tuple` | Immutable collections | `(1, 2, 3)` |
| Function | `kazi` | Function values | `lambda() { ... }` |

### Special Types

- **`any`** - Accepts any type (no type checking)
- **`void`** - No return value (for functions)

## Variable Type Annotations

### Explicit Type Declarations

```kwenda
// Number variables
namba x = 10
namba price = 99.99

// String variables
maneno jina = "Amina"
maneno message = "Habari!"

// Boolean variables
boolean isActive = kweli
boolean hasPermission = uwongo

// Array variables
orodha namba numbers = [1, 2, 3, 4, 5]
orodha maneno names = ["Ali", "Fatuma"]

// Dictionary variables
kamusi userInfo = {"name": "Juma", "age": 25}

// Set variables
seti namba uniqueNumbers = {1, 2, 3}

// Tuple variables
tuple namba coordinates = (10, 20, 30)
```

### Type Inference

When no type is explicitly specified, Kwenda infers the type from the value:

```kwenda
namba x = 42          // Type: namba
maneno name = "Ali"   // Type: maneno
boolean flag = kweli  // Type: boolean
```

## Function Type Annotations

### Parameter Types

Specify types for function parameters:

```kwenda
kazi greet(maneno name) {
    andika("Habari", name)
}

kazi add(namba a, namba b) {
    rudisha a + b
}

kazi display(maneno message, namba count, boolean show) {
    // Function body
}
```

### Return Types

Specify the return type after parameters:

```kwenda
kazi add(namba a, namba b) namba {
    rudisha a + b
}

kazi getName() maneno {
    rudisha "Amina"
}

kazi isValid(namba x) boolean {
    rudisha x > 0
}

kazi getNumbers() orodha {
    rudisha [1, 2, 3, 4, 5]
}
```

### Function Without Return Type

If a function doesn't return a value, you can omit the return type:

```kwenda
kazi printMessage(maneno msg) {
    andika(msg)
}

// Or explicitly specify void
kazi doSomething() void {
    andika("Doing something...")
}
```

## Type Checking Modes

Kwenda offers two type checking modes:

### 1. Non-Strict Mode (`--check`)

Performs basic type checking but allows:
- Undeclared variables
- Type mismatches when using `any` type
- Missing type annotations

```bash
kwenda --check program.swh
```

**Use when:**
- Developing prototypes
- Learning the language
- Working with dynamic code

### 2. Strict Mode (`--strict`)

Enforces strict type safety:
- All variables must be declared
- Type mismatches are errors
- Function signatures must match exactly
- Conditions must be boolean

```bash
kwenda --strict program.swh
```

**Use when:**
- Building production code
- Working in teams
- Preventing runtime errors
- Maintaining large codebases

## Type Checking Examples

### Example 1: Correct Typing

```kwenda
kazi multiply(namba a, namba b) namba {
    rudisha a * b
}

namba result = multiply(5, 10)
andika("Result:", result)
```

**Type check:** ✓ Passes (all types match)

### Example 2: Type Mismatch

```kwenda
kazi add(namba a, namba b) namba {
    rudisha a + b
}

maneno text = "hello"
namba result = add(text, 5)  // Error!
```

**Error:** `Expected 'namba', got 'maneno'`

### Example 3: Return Type Mismatch

```kwenda
kazi getName() maneno {
    rudisha 42  // Error!
}
```

**Error:** `Function return type mismatch. Expected 'maneno', got 'namba'`

### Example 4: Wrong Argument Count

```kwenda
kazi greet(maneno name, namba age) {
    andika("Hello", name, age)
}

greet("Ali")  // Error!
```

**Error:** `Function expects 2 arguments, got 1`

### Example 5: Boolean Condition

```kwenda
namba x = 10

kama x {  // Error in strict mode!
    andika("True")
}
```

**Error:** `Condition of 'kama' must be type 'boolean'`

**Correct version:**
```kwenda
kama x > 0 {  // Returns boolean
    andika("Positive")
}
```

## Benefits of Type Annotations

### 1. Early Error Detection

Catch type errors before running code:

```kwenda
kazi divide(namba a, namba b) namba {
    rudisha a / b
}

// Type checker catches this before execution
maneno result = divide("10", "2")  // Error!
```

### 2. Better Documentation

Types document what values functions expect:

```kwenda
// Clear what this function does
kazi calculateArea(namba width, namba height) namba {
    rudisha width * height
}
```

### 3. IDE Support

Type information enables:
- Autocomplete suggestions
- Parameter hints
- Refactoring tools
- Error highlighting

### 4. Team Collaboration

Types make code more maintainable:
- Clear function contracts
- Easier to understand code
- Safer refactoring
- Reduced bugs

## Best Practices

### 1. Always Annotate Function Signatures

```kwenda
// Good
kazi calculateTotal(namba price, namba quantity) namba {
    rudisha price * quantity
}

// Less clear
kazi calculateTotal(price, quantity) {
    rudisha price * quantity
}
```

### 2. Use Specific Types

```kwenda
// Good
kazi getNames() orodha maneno {
    rudisha ["Ali", "Fatuma"]
}

// Less specific
kazi getNames() orodha {
    rudisha ["Ali", "Fatuma"]
}
```

### 3. Be Consistent

Choose either strict or non-strict mode for your project and stick with it.

### 4. Annotate Public APIs

Always add types to functions that are called from other modules:

```kwenda
// Public API - well typed
kazi processUser(maneno name, namba age, boolean active) kamusi {
    rudisha {"name": name, "age": age, "active": active}
}
```

### 5. Use Type Checking in CI/CD

Add type checking to your build pipeline:

```bash
#!/bin/bash
# build.sh

echo "Running type checker..."
kwenda --strict src/main.swh || exit 1

echo "Running tests..."
kwenda test || exit 1

echo "Build successful!"
```

## Advanced Type Features

### Complex Function Signatures

```kwenda
kazi processData(
    maneno filename,
    namba limit,
    boolean includeHeaders
) orodha {
    // Process and return data
    rudisha []
}
```

### Functions Returning Functions

```kwenda
kazi createMultiplier(namba factor) kazi {
    rudisha lambda(namba x) namba {
        rudisha x * factor
    }
}

kazi double = createMultiplier(2)
namba result = double(5)  // 10
```

### Array Type Annotations

```kwenda
// Array of numbers
orodha namba scores = [85, 90, 78, 92]

// Array of strings
orodha maneno names = ["Ali", "Fatuma", "Juma"]

// Mixed arrays (use without type)
orodha mixed = [1, "text", kweli]
```

## Common Type Errors and Solutions

### Error 1: Type Mismatch in Assignment

```kwenda
namba x = 10
x = "hello"  // Error!
```

**Solution:** Keep consistent types
```kwenda
namba x = 10
x = 20  // OK
```

### Error 2: Wrong Parameter Type

```kwenda
kazi add(namba a, namba b) namba {
    rudisha a + b
}

add("5", "10")  // Error!
```

**Solution:** Convert to correct type or use correct type
```kwenda
add(5, 10)  // OK
```

### Error 3: Missing Return Statement

```kwenda
kazi getValue() namba {
    namba x = 10
    // Missing rudisha!
}  // Error in strict mode
```

**Solution:** Always return a value
```kwenda
kazi getValue() namba {
    namba x = 10
    rudisha x
}
```

### Error 4: Void Function Returning Value

```kwenda
kazi printMessage() void {
    rudisha "Done"  // Error!
}
```

**Solution:** Either return nothing or change return type
```kwenda
kazi printMessage() {
    andika("Done")
}

// Or
kazi getMessage() maneno {
    rudisha "Done"
}
```

## Type Checking Workflow

### Development Workflow

1. **Write Code** - Add type annotations as you code
2. **Quick Check** - Run `kwenda --check` for basic validation
3. **Strict Check** - Run `kwenda --strict` before committing
4. **Execute** - Run `kwenda program.swh` to execute

### Example Workflow

```bash
# Step 1: Write your program
vim myprogram.swh

# Step 2: Quick type check
kwenda --check myprogram.swh

# Step 3: Fix any errors, then strict check
kwenda --strict myprogram.swh

# Step 4: Run the program
kwenda myprogram.swh
```

## Integration with Tools

### Git Pre-commit Hook

```bash
#!/bin/bash
# .git/hooks/pre-commit

echo "Running type checker..."
for file in $(git diff --cached --name-only | grep ".swh$"); do
    kwenda --strict "$file"
    if [ $? -ne 0 ]; then
        echo "Type checking failed for $file"
        exit 1
    fi
done
```

### Makefile Integration

```makefile
.PHONY: check strict run test

check:
	kwenda --check src/*.swh

strict:
	kwenda --strict src/*.swh

run: check
	kwenda src/main.swh

test: strict
	./run_tests.sh
```

## Summary

Type annotations in Kwenda provide:

- ✅ **Optional** - Use as much or as little as you want
- ✅ **Flexible** - Two modes: `--check` and `--strict`
- ✅ **Clear** - Self-documenting code
- ✅ **Safe** - Catch errors before runtime
- ✅ **Bilingual** - Errors in Swahili and English
- ✅ **Practical** - Works with existing code

Use `kwenda --check` for development and `kwenda --strict` for production code to maximize safety and maintainability.
