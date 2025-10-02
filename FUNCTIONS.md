# Functions in Kwenda

This document describes function definitions, parameters, and return values in the Kwenda programming language.

## Function Definition

### Basic Syntax
```swahili
kazi function_name(parameter_list) return_type {
    // function body
    rudisha value  // optional return statement
}
```

### Components
- `kazi` - keyword for function definition
- `function_name` - name of the function
- `parameter_list` - comma-separated list of parameters (optional)
- `return_type` - type of value returned (optional)
- `rudisha` - keyword for return statement

## Function Parameters

### Parameter Syntax
```swahili
type parameter_name
```

### Supported Parameter Types
- `namba` - number parameter
- `boolean` - boolean parameter

### Examples
```swahili
# Function with no parameters
kazi salamu() {
    andika("Habari!")
}

# Function with one parameter
kazi salamu_mtu(namba umri) {
    andika("Habari, umri wako ni", umri)
}

# Function with multiple parameters
kazi jumla(namba a, namba b) {
    namba jibu = a + b
    andika("Jumla:", jibu)
}
```

## Return Values

### Return Types
Functions can return:
- `namba` - number value
- `boolean` - boolean value
- No return type (void function)

### Return Statement
```swahili
rudisha value
```

### Examples
```swahili
# Function returning a number
kazi ongeza(namba x, namba y) namba {
    rudisha x + y
}

# Function returning a boolean
kazi ni_kubwa(namba a, namba b) boolean {
    rudisha a > b
}

# Function with no return value
kazi andika_salamu() {
    andika("Habari za asubuhi!")
    # No return statement needed
}
```

## Function Calls

### Calling Functions
```swahili
function_name(arguments)
```

### Using Return Values
```swahili
# Assign return value to variable
namba jibu = ongeza(5, 3)

# Use return value directly
andika("Jibu ni:", ongeza(10, 20))

# Use in expressions
namba kubwa = ongeza(5, 3) * 2
```

## Complete Examples

### Example 1: Basic Arithmetic Functions
```swahili
kazi ongeza(namba a, namba b) namba {
    rudisha a + b
}

kazi toa(namba a, namba b) namba {
    rudisha a - b
}

kazi kuu() {
    namba x = 10
    namba y = 5
    
    namba jumla = ongeza(x, y)
    namba tofauti = toa(x, y)
    
    andika("Jumla:", jumla)      # Output: Jumla: 15
    andika("Tofauti:", tofauti)  # Output: Tofauti: 5
}
```

### Example 2: Boolean Functions
```swahili
kazi ni_sawa(namba a, namba b) boolean {
    rudisha a == b
}

kazi ni_kubwa_kuliko_kumi(namba n) boolean {
    rudisha n > 10
}

kazi kuu() {
    boolean sawa = ni_sawa(5, 5)
    boolean kubwa = ni_kubwa_kuliko_kumi(15)
    
    andika("5 == 5:", sawa)      # Output: 5 == 5: true
    andika("15 > 10:", kubwa)    # Output: 15 > 10: true
}
```

### Example 3: Conditional Returns
```swahili
kazi max(namba a, namba b) namba {
    kama a > b {
        rudisha a
    } sivyo {
        rudisha b
    }
}

kazi angalia_umri(namba umri) {
    kama umri < 18 {
        andika("Wewe ni mtoto")
        rudisha
    }
    
    kama umri < 60 {
        andika("Wewe ni mtu mzima")
    } sivyo {
        andika("Wewe ni mzee")
    }
}

kazi kuu() {
    namba mkubwa = max(15, 8)
    andika("Mkubwa:", mkubwa)    # Output: Mkubwa: 15
    
    angalia_umri(25)             # Output: Wewe ni mtu mzima
}
```

### Example 4: Functions with Loops
```swahili
kazi hesabu_jumla(namba mwanzo, namba mwisho) namba {
    namba jumla = 0
    namba i = mwanzo
    
    wakati i <= mwisho {
        jumla = jumla + i
        i = i + 1
    }
    
    rudisha jumla
}

kazi kuu() {
    namba jibu = hesabu_jumla(1, 5)
    andika("Jumla ya 1 hadi 5:", jibu)  # Output: Jumla ya 1 hadi 5: 15
}
```

### Example 5: Nested Function Calls
```swahili
kazi zidisha(namba a, namba b) namba {
    rudisha a * b
}

kazi ongeza(namba a, namba b) namba {
    rudisha a + b
}

kazi hesabu_wastani(namba a, namba b, namba c) namba {
    namba jumla = ongeza(ongeza(a, b), c)
    rudisha jumla / 3
}

kazi kuu() {
    namba wastani = hesabu_wastani(10, 20, 30)
    andika("Wastani:", wastani)  # Output: Wastani: 20
}
```

## Function Scope

### Variable Scope
- Parameters are local to the function
- Variables declared inside functions are local
- Global variables are not supported

### Example
```swahili
kazi test_scope(namba x) namba {
    namba y = x * 2  # y is local to this function
    rudisha y
}

kazi kuu() {
    namba a = 5
    namba jibu = test_scope(a)
    andika("Jibu:", jibu)
    # y is not accessible here
}
```

## Best Practices

### 1. Use Descriptive Names
```swahili
# Good
kazi hesabu_umri_kwa_miaka(namba miaka) namba {
    rudisha miaka
}

# Poor
kazi h(namba m) namba {
    rudisha m
}
```

### 2. Keep Functions Small
```swahili
# Good - single responsibility
kazi ni_namba_ya_kawaida(namba n) boolean {
    rudisha n > 0
}

kazi hesabu_kodi(namba mshahara) namba {
    rudisha mshahara * 0.3
}
```

### 3. Use Return Types
```swahili
# Good - clear return type
kazi hesabu_jumla(namba a, namba b) namba {
    rudisha a + b
}

# Acceptable for void functions
kazi andika_ripoti() {
    andika("Ripoti imeandikwa")
}
```

### 4. Validate Parameters
```swahili
kazi gawanya(namba a, namba b) namba {
    kama b == 0 {
        andika("Hitilafu: Haiwezi kugawanya na sifuri!")
        rudisha 0
    }
    rudisha a / b
}
```

## Common Patterns

### 1. Helper Functions
```swahili
kazi ni_namba_ya_kawaida(namba n) boolean {
    rudisha n > 0
}

kazi hesabu_kodi(namba mshahara) namba {
    kama ni_namba_ya_kawaida(mshahara) {
        rudisha mshahara * 0.3
    }
    rudisha 0
}
```

### 2. Calculation Functions
```swahili
kazi hesabu_eneo_mstatili(namba urefu, namba upana) namba {
    rudisha urefu * upana
}

kazi hesabu_mzunguko_mstatili(namba urefu, namba upana) namba {
    rudisha 2 * (urefu + upana)
}
```

### 3. Validation Functions
```swahili
kazi ni_umri_halali(namba umri) boolean {
    rudisha umri >= 0 na umri <= 150
}

kazi ni_alama_halali(namba alama) boolean {
    rudisha alama >= 0 na alama <= 100
}
```

## Keywords Summary

- `kazi` - Function definition keyword
- `rudisha` - Return statement keyword
- `namba` - Number parameter/return type
- `boolean` - Boolean parameter/return type
- `kuu` - Main function name (entry point)

## Current Limitations

1. **No function overloading**: Cannot have multiple functions with the same name
2. **No default parameters**: All parameters must be provided
3. **No variable arguments**: Fixed number of parameters only
4. **No recursion optimization**: Deep recursion may cause issues
5. **Limited types**: Only `namba` and `boolean` types supported