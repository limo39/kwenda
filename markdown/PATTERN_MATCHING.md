# Pattern Matching (Switch Statements) in Kwenda

## Overview

Kwenda supports pattern matching through switch statements, allowing you to match a value against multiple patterns and execute different code blocks based on the match. This provides a cleaner alternative to multiple if-else statements.

## Syntax

### Basic Switch Statement

```kwenda
swichi expression {
    ikiwa value1 {
        // code to execute if expression == value1
    }
    ikiwa value2 {
        // code to execute if expression == value2
    }
    kawaida {
        // code to execute if no match (default case)
    }
}
```

### Keywords

- **swichi** - switch statement keyword
- **ikiwa** - case keyword (means "if it is")
- **kawaida** - default case keyword (means "normally/usually")

## Features

### 1. Multiple Values in One Case

You can match multiple values in a single case clause:

```kwenda
namba day = 6
swichi day {
    ikiwa 1, 2, 3, 4, 5 {
        andika("Weekday")
    }
    ikiwa 6, 7 {
        andika("Weekend")
    }
}
```

### 2. Supported Data Types

Switch statements work with:
- **Numbers** (namba) - integers and floats
- **Strings** (maneno)
- **Booleans** (boolean)
- **Enums** (aina)

### 3. No Fall-through

Unlike some languages (like C), Kwenda switch statements do NOT fall through to the next case. Once a case matches and executes, the switch exits automatically.

```kwenda
namba x = 1
swichi x {
    ikiwa 1 {
        andika("One")
        // Exits here automatically
    }
    ikiwa 2 {
        andika("Two")  // This won't execute
    }
}
```

### 4. Default Case (Optional)

The `kawaida` clause is optional and executes when no case matches:

```kwenda
namba x = 100
swichi x {
    ikiwa 1 {
        andika("One")
    }
    ikiwa 2 {
        andika("Two")
    }
    kawaida {
        andika("Other number")  // This executes
    }
}
```

### 5. Early Return

You can use `rudisha` (return) inside switch cases:

```kwenda
kazi getDayType(namba day) {
    swichi day {
        ikiwa 1, 2, 3, 4, 5 {
            rudisha "Weekday"
        }
        ikiwa 6, 7 {
            rudisha "Weekend"
        }
        kawaida {
            rudisha "Invalid day"
        }
    }
}
```

### 6. Break Statement

You can use `vunja` (break) to exit a switch early:

```kwenda
swichi x {
    ikiwa 1 {
        andika("Processing...")
        kama error {
            vunja  // Exit switch
        }
        andika("Done")
    }
}
```

## Examples

### Example 1: Days of the Week

```kwenda
namba day = 3
swichi day {
    ikiwa 1 {
        andika("Jumatatu (Monday)")
    }
    ikiwa 2 {
        andika("Jumanne (Tuesday)")
    }
    ikiwa 3 {
        andika("Jumatano (Wednesday)")
    }
    ikiwa 4 {
        andika("Alhamisi (Thursday)")
    }
    ikiwa 5 {
        andika("Ijumaa (Friday)")
    }
    ikiwa 6 {
        andika("Jumamosi (Saturday)")
    }
    ikiwa 7 {
        andika("Jumapili (Sunday)")
    }
    kawaida {
        andika("Invalid day")
    }
}
```

### Example 2: Traffic Light Colors

```kwenda
maneno color = "yellow"
swichi color {
    ikiwa "red" {
        andika("Stop!")
    }
    ikiwa "yellow" {
        andika("Slow down")
    }
    ikiwa "green" {
        andika("Go")
    }
    kawaida {
        andika("Unknown color")
    }
}
```

### Example 3: Grading System

```kwenda
kazi getGrade(namba score) {
    namba grade = 0
    
    kama score >= 90 {
        grade = 5
    } sivyo kama score >= 80 {
        grade = 4
    } sivyo kama score >= 70 {
        grade = 3
    } sivyo kama score >= 60 {
        grade = 2
    } sivyo {
        grade = 1
    }
    
    swichi grade {
        ikiwa 5 {
            rudisha "A - Excellent!"
        }
        ikiwa 4 {
            rudisha "B - Very good!"
        }
        ikiwa 3 {
            rudisha "C - Good"
        }
        ikiwa 2 {
            rudisha "D - Average"
        }
        kawaida {
            rudisha "F - Fail"
        }
    }
}

andika(getGrade(85))  // Output: B - Very good!
```

### Example 4: Menu System

```kwenda
kazi processMenu(namba choice) {
    swichi choice {
        ikiwa 1 {
            andika("Creating new file...")
        }
        ikiwa 2 {
            andika("Opening file...")
        }
        ikiwa 3 {
            andika("Saving file...")
        }
        ikiwa 4 {
            andika("Printing...")
        }
        ikiwa 5 {
            andika("Exiting...")
        }
        kawaida {
            andika("Invalid choice! Please select 1-5")
        }
    }
}
```

### Example 5: Calculator

```kwenda
kazi calculate(namba a, maneno operator, namba b) {
    swichi operator {
        ikiwa "+" {
            rudisha a + b
        }
        ikiwa "-" {
            rudisha a - b
        }
        ikiwa "*" {
            rudisha a * b
        }
        ikiwa "/" {
            kama b == 0 {
                rudisha "Error: Division by zero"
            }
            rudisha a / b
        }
        kawaida {
            rudisha "Unknown operator"
        }
    }
}

andika(calculate(10, "+", 5))   // 15
andika(calculate(10, "*", 5))   // 50
```

### Example 6: Seasons by Month

```kwenda
namba month = 7
swichi month {
    ikiwa 12, 1, 2 {
        andika("Winter")
    }
    ikiwa 3, 4, 5 {
        andika("Spring")
    }
    ikiwa 6, 7, 8 {
        andika("Summer")
    }
    ikiwa 9, 10, 11 {
        andika("Autumn")
    }
    kawaida {
        andika("Invalid month")
    }
}
```

### Example 7: HTTP Status Codes

```kwenda
kazi getStatusMessage(namba code) {
    swichi code {
        ikiwa 200 {
            rudisha "OK - Success"
        }
        ikiwa 201 {
            rudisha "Created"
        }
        ikiwa 400 {
            rudisha "Bad Request"
        }
        ikiwa 401 {
            rudisha "Unauthorized"
        }
        ikiwa 403 {
            rudisha "Forbidden"
        }
        ikiwa 404 {
            rudisha "Not Found"
        }
        ikiwa 500 {
            rudisha "Internal Server Error"
        }
        kawaida {
            rudisha "Unknown Status Code"
        }
    }
}
```

### Example 8: Boolean Switch

```kwenda
boolean isStudent = kweli
swichi isStudent {
    ikiwa kweli {
        andika("Student discount applied")
    }
    ikiwa uwongo {
        andika("Regular price")
    }
}
```

### Example 9: Nested Switch Statements

```kwenda
namba category = 1
namba subcategory = 2

swichi category {
    ikiwa 1 {
        andika("Category: Software")
        swichi subcategory {
            ikiwa 1 {
                andika("  - Text Editor")
            }
            ikiwa 2 {
                andika("  - Browser")
            }
            kawaida {
                andika("  - Other")
            }
        }
    }
    ikiwa 2 {
        andika("Category: Hardware")
    }
    kawaida {
        andika("Unknown category")
    }
}
```

## Comparison with If-Else

### Using if-else:
```kwenda
kama x == 1 {
    andika("One")
} sivyo kama x == 2 {
    andika("Two")
} sivyo kama x == 3 {
    andika("Three")
} sivyo {
    andika("Other")
}
```

### Using switch:
```kwenda
swichi x {
    ikiwa 1 {
        andika("One")
    }
    ikiwa 2 {
        andika("Two")
    }
    ikiwa 3 {
        andika("Three")
    }
    kawaida {
        andika("Other")
    }
}
```

The switch statement is cleaner and more maintainable for multiple value comparisons.

## Best Practices

1. **Use switch for multiple equality checks** - When comparing a single value against many possibilities, switch is clearer than if-else chains.

2. **Include a default case** - Always include a `kawaida` clause to handle unexpected values, even if it just logs an error.

3. **Keep cases simple** - Each case should perform a focused task. For complex logic, call a function.

4. **Group related values** - Use comma-separated values to group related cases (e.g., weekdays vs weekends).

5. **Consider readability** - If you only have 2-3 cases, an if-else might be more readable.

## Common Patterns

### State Machine
```kwenda
maneno state = "PENDING"
swichi state {
    ikiwa "PENDING" {
        andika("Processing request...")
    }
    ikiwa "APPROVED" {
        andika("Request approved!")
    }
    ikiwa "REJECTED" {
        andika("Request rejected")
    }
    kawaida {
        andika("Unknown state")
    }
}
```

### Command Dispatcher
```kwenda
kazi executeCommand(maneno cmd) {
    swichi cmd {
        ikiwa "start" {
            startProcess()
        }
        ikiwa "stop" {
            stopProcess()
        }
        ikiwa "restart" {
            stopProcess()
            startProcess()
        }
        ikiwa "status" {
            showStatus()
        }
        kawaida {
            andika("Unknown command:", cmd)
        }
    }
}
```

### Type Handling
```kwenda
kazi processValue(namba type, value) {
    swichi type {
        ikiwa 1 {
            // Handle as number
            andika("Number:", value)
        }
        ikiwa 2 {
            // Handle as string
            andika("String:", value)
        }
        ikiwa 3 {
            // Handle as boolean
            andika("Boolean:", value)
        }
        kawaida {
            andika("Unknown type")
        }
    }
}
```

## Performance Notes

- Switch statements use efficient value comparison
- No performance penalty compared to if-else chains
- Multiple values in one case are efficiently evaluated

## Summary

Switch statements in Kwenda provide:
- ✅ Clean syntax for multiple value comparisons
- ✅ Support for numbers, strings, booleans, and enums
- ✅ Multiple values per case
- ✅ No fall-through behavior (safer)
- ✅ Optional default case
- ✅ Works with return and break statements
- ✅ Can be nested

Use `swichi/ikiwa/kawaida` to write cleaner, more maintainable code when comparing a value against multiple possibilities.
