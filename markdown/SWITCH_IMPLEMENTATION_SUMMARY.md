# Switch Statement Implementation Summary

## Overview
Pattern matching support has been added to Kwenda using switch/case statements with Swahili keywords: `swichi` (switch), `ikiwa` (case), and `kawaida` (default).

## Keywords
- **swichi** - Switch statement
- **ikiwa** - Case clause (literally "if it is")
- **kawaida** - Default case (literally "normally/usually")

## Architecture

### 1. AST Nodes (ast/ast.go)

#### SwitchNode
Represents a complete switch statement:
```go
type SwitchNode struct {
    Expression ASTNode     // The value being matched
    Cases      []CaseNode  // List of case clauses
    Default    []ASTNode   // Optional default case body
}
```

#### CaseNode
Represents a single case clause:
```go
type CaseNode struct {
    Values []ASTNode // List of values to match (supports multiple)
    Body   []ASTNode // Statements to execute if matched
}
```

### 2. Lexer Updates (lexer/lexer.go)

Added three new keywords to the keyword list:
- `swichi`
- `ikiwa`
- `kawaida`

These are recognized during lexical analysis and tokenized as KEYWORD type tokens.

### 3. Parser Implementation (parser/parser.go)

#### Main Parse Function
Added switch statement detection:
```go
if tokens[0].Value == "swichi" {
    return ParseSwitchStatement(tokens)
}
```

#### ParseSwitchStatement
Core parsing logic:
1. Validates that first token is "swichi"
2. Extracts expression between "swichi" and "{"
3. Finds the switch body boundaries
4. Iterates through body to find "ikiwa" and "kawaida" clauses
5. Delegates case parsing to ParseCaseClause
6. Constructs and returns SwitchNode

#### ParseCaseClause
Parses individual case clauses:
1. Validates "ikiwa" keyword
2. Extracts comma-separated values before "{"
3. Parses each value as an expression
4. Extracts case body between "{" and "}"
5. Returns CaseNode and next position

#### Key Features
- **Multiple values per case**: Values are comma-separated
- **Nested brace tracking**: Proper handling of nested blocks
- **Expression parsing**: Each case value is parsed as a full expression

### 4. Interpreter Implementation (interpreter/interpreter.go)

#### SwitchNode Handler
Execution logic:
```go
case ast.SwitchNode:
    switchExpr := Interpret(n.Expression, env)
    
    // Try each case
    for _, caseClause := range n.Cases {
        matched := false
        for _, caseValue := range caseClause.Values {
            if compareValues(switchExpr, caseVal) {
                matched = true
                break
            }
        }
        
        if matched {
            // Execute case body
            // No fall-through - exit after first match
            return result
        }
    }
    
    // Execute default if no match
    if len(n.Default) > 0 {
        // Execute default body
    }
```

#### compareValues Function
Value comparison helper:
- Handles nil cases
- Supports int, float64, string, bool, EnumValue
- Performs type-aware equality checking
- Handles numeric type coercion (int ↔ float64)

### 5. Control Flow Support

#### Return Statements
Switch cases can use `rudisha` to return from the containing function:
```go
if cf.Type == ControlReturn {
    return result  // Propagate return up
}
```

#### Break Statements
Switch cases can use `vunja` to exit early:
```go
if cf.Type == ControlBreak {
    return cf.Value  // Exit switch
}
```

#### No Fall-through
After executing a matching case, the switch automatically exits. This is safer than traditional C-style fall-through behavior.

### 6. REPL Integration (repl/repl.go)

Added "swichi" to the keyword list for expression detection, ensuring switch statements are properly recognized as statements rather than expressions in the REPL.

## Syntax Examples

### Basic Switch
```kwenda
swichi x {
    ikiwa 1 {
        andika("One")
    }
    ikiwa 2 {
        andika("Two")
    }
    kawaida {
        andika("Other")
    }
}
```

### Multiple Values
```kwenda
swichi day {
    ikiwa 1, 2, 3, 4, 5 {
        andika("Weekday")
    }
    ikiwa 6, 7 {
        andika("Weekend")
    }
}
```

### With Strings
```kwenda
swichi color {
    ikiwa "red" {
        andika("Stop")
    }
    ikiwa "yellow" {
        andika("Caution")
    }
    ikiwa "green" {
        andika("Go")
    }
}
```

### In Functions
```kwenda
kazi getGrade(namba score) {
    swichi score {
        ikiwa 10, 9 {
            rudisha "A"
        }
        ikiwa 8, 7 {
            rudisha "B"
        }
        kawaida {
            rudisha "F"
        }
    }
}
```

### Nested Switches
```kwenda
swichi category {
    ikiwa 1 {
        swichi subcategory {
            ikiwa 1 {
                andika("Software")
            }
            ikiwa 2 {
                andika("Hardware")
            }
        }
    }
    ikiwa 2 {
        andika("Books")
    }
}
```

## Supported Types

The switch statement works with:
1. **Numbers** (namba) - both int and float64
2. **Strings** (maneno)
3. **Booleans** (boolean) - kweli/uwongo
4. **Enums** (aina) - custom enumeration types

## Design Decisions

### 1. No Fall-through
**Decision**: Cases don't fall through to the next case
**Rationale**: 
- Safer and less error-prone
- More intuitive for beginners
- Aligns with modern languages (Swift, Rust, Kotlin)
- C-style fall-through is a common source of bugs

### 2. Multiple Values Per Case
**Decision**: Allow comma-separated values in one case
**Rationale**:
- Reduces code duplication
- More readable than multiple identical cases
- Common pattern for grouping related values

### 3. Optional Default Case
**Decision**: Default case (kawaida) is optional
**Rationale**:
- Some switches exhaustively cover all possibilities
- Forcing default would add unnecessary boilerplate
- Developer can choose when to handle unexpected values

### 4. Expression-based Matching
**Decision**: Switch on any expression result
**Rationale**:
- Maximum flexibility
- Can switch on function results, calculations, etc.
- Not limited to simple variables

### 5. Type-aware Comparison
**Decision**: Use compareValues helper for equality
**Rationale**:
- Handles numeric type coercion (int vs float64)
- Supports all Kwenda data types
- Extensible for future types

## Error Handling

The implementation handles:
- Missing expression after "swichi"
- Missing braces around switch body
- Missing braces around case bodies
- Missing values in case clauses
- Nested braces in case bodies

Parse errors result in nil return, allowing the parser to try other statement types.

## Performance Characteristics

- **Time Complexity**: O(n) where n is the number of cases
- **Space Complexity**: O(1) for the switch itself, O(m) for case bodies where m is total statements
- **Optimization Potential**: Could add jump table optimization for dense integer cases in the future

## Testing

Test files:
- `examples/switch_demo.swh` - Comprehensive examples (12 scenarios)
- `examples/switch_simple.swh` - Basic functionality tests

Test coverage:
- ✅ Number matching
- ✅ String matching
- ✅ Boolean matching
- ✅ Multiple values per case
- ✅ Default case
- ✅ Switch without default
- ✅ Return in switch
- ✅ Functions with switch
- ✅ Nested switches
- ✅ Menu systems
- ✅ Calculator operations
- ✅ State machines

## Documentation

- `markdown/PATTERN_MATCHING.md` - Complete user guide with examples and best practices
- `markdown/SWITCH_IMPLEMENTATION_SUMMARY.md` - This technical summary
- `README.md` - Updated with switch keywords and feature list

## Future Enhancements

Potential improvements:
1. **Pattern matching on types** - Match different data types
2. **Guard clauses** - Add conditions to cases (e.g., `ikiwa x kama x > 5`)
3. **Destructuring** - Match array/tuple elements
4. **Range matching** - Match numeric ranges
5. **Jump table optimization** - For dense integer switches
6. **Exhaustiveness checking** - Warn if not all enum values covered

## Comparison with Other Languages

### Go
```go
switch x {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Other")
}
```

### Kwenda
```kwenda
swichi x {
    ikiwa 1 {
        andika("One")
    }
    ikiwa 2 {
        andika("Two")
    }
    kawaida {
        andika("Other")
    }
}
```

Key differences:
- Kwenda uses Swahili keywords
- Both have no fall-through by default
- Both support multiple values per case
- Similar syntax structure

## Integration Points

The switch statement integrates with:
- **Variables**: Can switch on any variable
- **Functions**: Can switch on function results
- **Control flow**: Works with return, break
- **Error handling**: Can be used in try-catch blocks
- **Enums**: Natural fit for enum value matching
- **REPL**: Full support in interactive mode

## Summary

The switch statement implementation provides:
- ✅ Clean, intuitive syntax using Swahili keywords
- ✅ Support for all Kwenda data types
- ✅ Multiple values per case
- ✅ No fall-through (safer)
- ✅ Optional default case
- ✅ Integration with control flow
- ✅ Comprehensive documentation and examples
- ✅ Efficient O(n) evaluation

This feature makes Kwenda more expressive and practical for real-world programming while maintaining its educational mission with clear, accessible Swahili syntax.
