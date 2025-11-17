# Loop Constructs in Kwenda

This document describes the loop constructs available in the Kwenda programming language.

## While Loops (`wakati`)

While loops execute a block of code repeatedly as long as a condition is true.

### Syntax
```swahili
wakati condition {
    // statements
}
```

### Example
```swahili
kazi kuu() {
    namba i = 1
    
    wakati i <= 5 {
        andika("Namba:", i)
        i = i + 1
    }
}
```

**Output:**
```
Namba: 1
Namba: 2
Namba: 3
Namba: 4
Namba: 5
```

## For Loops (`kwa`)

For loops in Kwenda support two forms:

### Simple For Loop (Condition Only)
This form works like a while loop but uses the `kwa` keyword.

```swahili
kwa condition {
    // statements
}
```

### Example
```swahili
kazi kuu() {
    namba i = 0
    
    kwa i < 3 {
        andika("i ni:", i)
        i = i + 1
    }
}
```

**Output:**
```
i ni: 0
i ni: 1
i ni: 2
```

### Full For Loop (Init; Condition; Update)
**Note: This syntax is currently under development and may not work correctly.**

```swahili
kwa init; condition; update {
    // statements
}
```

## Working Examples

### Example 1: Counting with While Loop
```swahili
kazi kuu() {
    namba hesabu = 1
    
    andika("Kuhesabu kutoka 1 hadi 3:")
    
    wakati hesabu <= 3 {
        andika("Hesabu:", hesabu)
        hesabu = hesabu + 1
    }
    
    andika("Mwisho!")
}
```

### Example 2: Simple For Loop
```swahili
kazi kuu() {
    namba j = 5
    
    andika("Kupunguza kutoka 5:")
    
    kwa j > 0 {
        andika("j ni:", j)
        j = j - 1
    }
    
    andika("Tumefika sifuri!")
}
```

### Example 3: Nested While Loops
```swahili
kazi kuu() {
    namba i = 1
    
    wakati i <= 2 {
        namba j = 1
        
        wakati j <= 2 {
            namba bidhaa = i * j
            andika(i, "x", j, "=", bidhaa)
            j = j + 1
        }
        
        i = i + 1
    }
}
```

**Output:**
```
1 x 1 = 1
1 x 2 = 2
2 x 1 = 2
2 x 2 = 4
```

## Loop Control Statements

### Break Statement (`vunja`)

The `vunja` statement immediately exits the current loop.

#### Syntax
```swahili
vunja
```

#### Example
```swahili
kazi kuu() {
    namba i = 1
    
    wakati i <= 10 {
        kama i == 5 {
            andika("Tunavunja loop!")
            vunja
        }
        andika("i =", i)
        i = i + 1
    }
}
```

**Output:**
```
i = 1
i = 2
i = 3
i = 4
Tunavunja loop!
```

### Continue Statement (`endelea`)

The `endelea` statement skips the rest of the current iteration and continues with the next iteration.

#### Syntax
```swahili
endelea
```

#### Example
```swahili
kazi kuu() {
    namba i = 0
    
    wakati i < 5 {
        i = i + 1
        kama i == 3 {
            andika("Tunaruka i =", i)
            endelea
        }
        andika("i =", i)
    }
}
```

**Output:**
```
i = 1
i = 2
Tunaruka i = 3
i = 4
i = 5
```

### Nested Loops with Break/Continue

Break and continue statements only affect the innermost loop they are in.

```swahili
kazi kuu() {
    namba i = 1
    wakati i <= 2 {
        namba j = 1
        wakati j <= 3 {
            kama j == 2 {
                j = j + 1
                endelea  # Only affects inner loop
            }
            andika(i, ":", j)
            j = j + 1
        }
        i = i + 1
    }
}
```

**Output:**
```
1 : 1
1 : 3
2 : 1
2 : 3
```

## Current Limitations

1. **No labeled break/continue**: Break and continue only affect the innermost loop
2. **Full for loop syntax**: The `kwa init; condition; update` syntax is experimental
3. **No infinite loop protection**: Be careful with loop conditions to avoid infinite loops

## Best Practices

1. Always ensure loop conditions will eventually become false
2. Initialize loop variables before the loop
3. Update loop variables inside the loop body
4. Use meaningful variable names for loop counters
5. Test loops with simple examples first

## Keywords

- `wakati` - while loop
- `kwa` - for loop
- `vunja` - break statement
- `endelea` - continue statement
- `namba` - number variable declaration
- `andika` - print statement