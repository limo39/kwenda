# Boolean Data Type in Kwenda

This document describes the Boolean data type and logical operations in the Kwenda programming language.

## Boolean Values

Kwenda supports Boolean values using Swahili keywords:

- `kweli` - represents `true`
- `uwongo` - represents `false`

## Boolean Variable Declaration

### Syntax
```swahili
boolean variable_name = value
```

### Examples
```swahili
boolean iko_kweli = kweli
boolean si_kweli = uwongo
boolean hali_nzuri = kweli
```

## Logical Operators

### AND Operator (`na`)
The `na` operator returns `kweli` only if both operands are true.

```swahili
boolean a = kweli
boolean b = uwongo
boolean result = a na b  # Result: uwongo
```

**Truth Table:**
| Left | Right | Result |
|------|-------|--------|
| kweli | kweli | kweli |
| kweli | uwongo | uwongo |
| uwongo | kweli | uwongo |
| uwongo | uwongo | uwongo |

### OR Operator (`au`)
The `au` operator returns `kweli` if at least one operand is true.

```swahili
boolean a = kweli
boolean b = uwongo
boolean result = a au b  # Result: kweli
```

**Truth Table:**
| Left | Right | Result |
|------|-------|--------|
| kweli | kweli | kweli |
| kweli | uwongo | kweli |
| uwongo | kweli | kweli |
| uwongo | uwongo | uwongo |

## Boolean Comparisons

Booleans can be compared using equality operators:

```swahili
boolean x = kweli
boolean y = uwongo

boolean same = x == kweli      # Result: kweli
boolean different = y != uwongo # Result: uwongo
boolean equal = x == y         # Result: uwongo
```

## Booleans in Conditionals

Booleans work naturally with conditional statements:

```swahili
kazi kuu() {
    boolean iko_jua = kweli
    boolean mvua = uwongo
    namba joto = 25
    
    kama iko_jua na joto > 20 {
        andika("Siku nzuri ya kwenda nje!")
    } sivyo {
        andika("Bora ukae ndani")
    }
    
    kama mvua au joto < 15 {
        andika("Chukua jaketi")
    }
}
```

## Booleans in Loops

Boolean variables can control loop execution:

### While Loops
```swahili
kazi kuu() {
    boolean endelea = kweli
    namba hesabu = 0
    
    wakati endelea {
        hesabu = hesabu + 1
        andika("Hesabu:", hesabu)
        
        kama hesabu >= 5 {
            endelea = uwongo
        }
    }
}
```

### For Loops
```swahili
kazi kuu() {
    boolean imeisha = uwongo
    namba i = 0
    
    kwa imeisha == uwongo {
        i = i + 1
        andika("i =", i)
        
        kama i >= 3 {
            imeisha = kweli
        }
    }
}
```

## Type Conversion

### Numbers to Booleans
- `0` converts to `uwongo` (false)
- Any non-zero number converts to `kweli` (true)

### Booleans to Numbers
- `kweli` converts to `1`
- `uwongo` converts to `0`

### Strings to Booleans
- Empty string `""` converts to `uwongo`
- Any non-empty string converts to `kweli`

## Complex Boolean Expressions

You can combine multiple Boolean operations:

```swahili
kazi kuu() {
    boolean iko_jua = kweli
    boolean hakuna_mvua = kweli
    namba joto = 25
    namba upepo = 5
    
    # Complex condition
    boolean hali_nzuri = iko_jua na hakuna_mvua na joto > 20
    
    kama hali_nzuri {
        andika("Hali ya hewa ni nzuri kabisa!")
    }
    
    # Another complex condition
    boolean chukua_jaketi = joto < 15 au upepo > 10
    
    kama chukua_jaketi {
        andika("Chukua jaketi")
    }
}
```

## Best Practices

1. **Use descriptive names**: Choose Boolean variable names that clearly indicate what they represent
   ```swahili
   boolean iko_jua = kweli        # Good
   boolean x = kweli              # Poor
   ```

2. **Avoid double negatives**: Use positive Boolean names when possible
   ```swahili
   boolean iko_tayari = kweli     # Good
   boolean si_tayari = uwongo     # Less clear
   ```

3. **Use Booleans for flags**: Boolean variables are perfect for controlling program flow
   ```swahili
   boolean endelea_loop = kweli
   boolean imemaliza = uwongo
   ```

4. **Combine with conditions**: Use Boolean variables to make complex conditions more readable
   ```swahili
   boolean hali_nzuri = joto > 20 na hakuna_mvua
   kama hali_nzuri {
       # Do something
   }
   ```

## Common Patterns

### Toggle Pattern
```swahili
boolean hali = kweli
hali = hali == uwongo  # Toggle the boolean value
```

### Flag Pattern
```swahili
boolean imepatikana = uwongo

wakati imepatikana == uwongo {
    # Search for something
    kama /* found condition */ {
        imepatikana = kweli
    }
}
```

### State Machine Pattern
```swahili
boolean imeanza = uwongo
boolean inafanya_kazi = uwongo
boolean imemaliza = uwongo

kama imeanza == uwongo {
    imeanza = kweli
    andika("Kuanza...")
}

kama imeanza na inafanya_kazi == uwongo {
    inafanya_kazi = kweli
    andika("Inafanya kazi...")
}
```

## Keywords Summary

- `boolean` - Boolean variable declaration
- `kweli` - Boolean true value
- `uwongo` - Boolean false value
- `na` - Logical AND operator
- `au` - Logical OR operator
- `==` - Equality comparison
- `!=` - Inequality comparison