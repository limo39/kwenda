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

## Current Limitations

1. **No break/continue statements**: Loops cannot be terminated early or skip iterations
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
- `namba` - number variable declaration
- `andika` - print statement