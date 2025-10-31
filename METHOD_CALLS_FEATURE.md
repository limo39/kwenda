# Method Calls with Dot Notation Feature

## Overview
Added support for calling methods using dot notation (e.g., `object.method(args)`) in Kwenda/Lugha Yangu.

## Implementation

### 1. AST Node
Added `MethodCallNode` to `ast/ast.go`:
```go
type MethodCallNode struct {
    Object ASTNode   // The object whose method is being called
    Method string    // The method name
    Args   []ASTNode // Method arguments
}
```

### 2. Parser Changes
- Added method call parsing in `Parse()` function (for statements)
- Added method call parsing in `ParseBlock()` function (for method calls within blocks)
- Added method call parsing in `ParseExpression()` function (for method calls in expressions)
- Improved binary operation parsing to handle complex expressions like `hii.count + 1`

### 3. Interpreter Changes
- Added `MethodCallNode` handler in `Interpret()` function
- Method calls:
  1. Evaluate the object
  2. Look up the class definition
  3. Find the method in the class
  4. Create a new environment for method execution
  5. Set `hii` to refer to the current instance
  6. Bind method parameters
  7. Execute method body
  8. Handle return values and control flow

### 4. Key Features
- ✅ Call methods with dot notation: `object.method()`
- ✅ Pass arguments to methods: `object.method(arg1, arg2)`
- ✅ Access `hii` (this/self) inside methods
- ✅ Modify object properties from within methods
- ✅ Method calls work in expressions
- ✅ Support for complex expressions in method bodies

## Examples

### Basic Method Call
```swahili
darasa Mtu {
    maneno jina
    
    kazi unda(maneno j) {
        hii.jina = j
    }
    
    kazi salamu() {
        andika("Habari!", hii.jina)
    }
}

kazi kuu() {
    kamusi mtu = unda Mtu("Amina")
    mtu.salamu()  # Output: Habari! Amina
}
```

### Method with Parameters
```swahili
darasa Counter {
    namba count
    
    kazi unda() {
        hii.count = 0
    }
    
    kazi ongeza(namba n) {
        hii.count = hii.count + n
    }
}

kazi kuu() {
    kamusi c = unda Counter()
    c.ongeza(5)
    andika("Count:", c["count"])  # Output: Count: 5
}
```

### Method Modifying Object State
```swahili
darasa Person {
    maneno name
    namba age
    
    kazi unda(maneno n, namba a) {
        hii.name = n
        hii.age = a
    }
    
    kazi birthday() {
        hii.age = hii.age + 1
        andika("Happy birthday!", hii.name, "is now", hii.age)
    }
}

kazi kuu() {
    kamusi p = unda Person("Alice", 25)
    p.birthday()  # Output: Happy birthday! Alice is now 26
}
```

## Testing
Created comprehensive test files:
- `tests/test_method_call.swh` - Full method call test
- `tests/test_method_simple.swh` - Simple method test
- `tests/test_method_debug.swh` - Debug test
- `tests/test_member_expr.swh` - Member access in expressions
- `tests/test_constructor.swh` - Constructor testing
- `tests/test_mixed_params.swh` - Mixed parameter types

All tests pass successfully!

## Benefits
1. **More intuitive OOP syntax** - Matches common programming language patterns
2. **Cleaner code** - `object.method()` is more readable than function-based approaches
3. **Better encapsulation** - Methods are clearly associated with their objects
4. **Familiar syntax** - Easier for developers coming from other languages

## Future Enhancements
- Method chaining (e.g., `object.method1().method2()`)
- Static methods
- Private methods
- Method overloading
