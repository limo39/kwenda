# Class Inheritance Feature

## Overview
Added support for class inheritance in Kwenda/Lugha Yangu, allowing classes to inherit properties and methods from parent classes.

## Syntax
Use colon (`:`) to specify inheritance:
```swahili
darasa ChildClass : ParentClass {
    # Child class properties and methods
}
```

## Implementation

### 1. AST Changes
Updated `ClassNode` in `ast/ast.go`:
```go
type ClassNode struct {
    Name       string         // Class name
    Parent     string         // Parent class name (for inheritance)
    Properties []PropertyNode // Class properties
    Methods    []FunctionNode // Class methods
    Constructor *FunctionNode // Constructor method (optional)
}
```

### 2. Parser Changes
Updated `ParseClassDefinition` in `parser/parser.go`:
- Detects inheritance syntax: `darasa Child : Parent`
- Parses parent class name
- Stores parent reference in ClassNode

### 3. Interpreter Changes
Added helper functions in `interpreter/interpreter.go`:
- `collectInheritedProperties()` - Collects properties from entire inheritance chain
- `collectInheritedMethods()` - Collects methods with override support
- `findMethodInClass()` - Finds methods in class or parent chain

Updated handlers:
- `NewInstanceNode` - Initializes inherited properties
- `MethodCallNode` - Searches for methods in inheritance chain

## Features

### ✅ Property Inheritance
Child classes inherit all properties from parent classes:
```swahili
darasa Animal {
    maneno name
    namba age
}

darasa Dog : Animal {
    maneno breed  # Dog has name, age, and breed
}
```

### ✅ Method Inheritance
Child classes inherit all methods from parent classes:
```swahili
darasa Animal {
    kazi speak() {
        andika("Animal sound")
    }
}

darasa Dog : Animal {
    # Dog inherits speak() method
}

kazi kuu() {
    kamusi dog = unda Dog()
    dog.speak()  # Works! Calls inherited method
}
```

### ✅ Method Overriding
Child classes can override parent methods:
```swahili
darasa Animal {
    kazi speak() {
        andika("Generic sound")
    }
}

darasa Dog : Animal {
    kazi speak() {
        andika("Woof!")  # Overrides parent method
    }
}

kazi kuu() {
    kamusi dog = unda Dog()
    dog.speak()  # Output: Woof!
}
```

### ✅ Multi-level Inheritance
Supports inheritance chains:
```swahili
darasa Animal {
    maneno name
}

darasa Mammal : Animal {
    namba legs
}

darasa Dog : Mammal {
    maneno breed
    # Dog has: name (from Animal), legs (from Mammal), breed (own)
}
```

## Examples

### Basic Inheritance
```swahili
darasa Parent {
    namba x
    
    kazi unda(namba val) {
        hii.x = val
    }
    
    kazi show() {
        andika("x =", hii.x)
    }
}

darasa Child : Parent {
    namba y
    
    kazi unda(namba a, namba b) {
        hii.x = a
        hii.y = b
    }
}

kazi kuu() {
    kamusi c = unda Child(10, 20)
    c.show()  # Inherited method works!
}
```

### Method Overriding
```swahili
darasa Animal {
    maneno name
    
    kazi unda(maneno n) {
        hii.name = n
    }
    
    kazi speak() {
        andika(hii.name, "makes a sound")
    }
}

darasa Dog : Animal {
    kazi speak() {
        andika(hii.name, "says: Woof!")
    }
}

darasa Cat : Animal {
    kazi speak() {
        andika(hii.name, "says: Meow!")
    }
}

kazi kuu() {
    kamusi dog = unda Dog("Buddy")
    kamusi cat = unda Cat("Whiskers")
    
    dog.speak()  # Output: Buddy says: Woof!
    cat.speak()  # Output: Whiskers says: Meow!
}
```

### Polymorphism Example
```swahili
darasa Shape {
    kazi area() namba {
        rudisha 0
    }
}

darasa Rectangle : Shape {
    namba width
    namba height
    
    kazi unda(namba w, namba h) {
        hii.width = w
        hii.height = h
    }
    
    kazi area() namba {
        rudisha hii.width * hii.height
    }
}

darasa Circle : Shape {
    namba radius
    
    kazi unda(namba r) {
        hii.radius = r
    }
    
    kazi area() namba {
        rudisha 3 * hii.radius * hii.radius
    }
}

kazi kuu() {
    kamusi rect = unda Rectangle(5, 10)
    kamusi circle = unda Circle(7)
    
    andika("Rectangle area:", rect.area())
    andika("Circle area:", circle.area())
}
```

## Benefits
1. **Code Reuse** - Inherit common functionality from parent classes
2. **Polymorphism** - Override methods for specialized behavior
3. **Hierarchical Organization** - Model real-world relationships
4. **Maintainability** - Changes to parent class affect all children
5. **Extensibility** - Easy to add new specialized classes

## Testing
Created comprehensive test files:
- `tests/test_inheritance.swh` - Full inheritance test with animals
- `tests/test_inheritance_simple.swh` - Simple parent-child test

All tests demonstrate:
- Property inheritance
- Method inheritance
- Method overriding
- Multi-level inheritance

## Future Enhancements
- `super` keyword to call parent methods
- Abstract classes/methods
- Multiple inheritance
- Interface/trait support
- Protected/private members
- Constructor chaining (automatic parent constructor calls)
