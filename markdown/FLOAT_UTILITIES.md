# Float Utilities in Kwenda

Kwenda provides comprehensive floating-point utilities for precise numerical computations, comparisons, and manipulations.

## Table of Contents
1. [Rounding Functions](#rounding-functions)
2. [Modulo and Remainder](#modulo-and-remainder)
3. [Sign Manipulation](#sign-manipulation)
4. [Approximate Comparisons](#approximate-comparisons)
5. [Type Checking](#type-checking)
6. [Value Clamping](#value-clamping)
7. [Part Extraction](#part-extraction)
8. [Interpolation](#interpolation)

---

## Rounding Functions

### pindika (Round to Nearest)
Rounds a number to the nearest integer.

**Syntax:**
```kwenda
pindika(x)
```

**Parameters:**
- `x` (number): The value to round

**Returns:**
- Integer: The nearest integer to x

**Examples:**
```kwenda
andika(pindika(3.4))    // Output: 3
andika(pindika(3.5))    // Output: 4
andika(pindika(3.6))    // Output: 4
andika(pindika(-2.5))   // Output: -2
```

**Notes:**
- Uses "round half away from zero" strategy
- `pindika(3.5)` = 4, `pindika(-3.5)` = -4

---

### sakafu (Floor)
Rounds a number down to the nearest integer.

**Syntax:**
```kwenda
sakafu(x)
```

**Parameters:**
- `x` (number): The value to floor

**Returns:**
- Integer: The largest integer ≤ x

**Examples:**
```kwenda
andika(sakafu(3.9))     // Output: 3
andika(sakafu(3.1))     // Output: 3
andika(sakafu(-2.1))    // Output: -3
andika(sakafu(-2.9))    // Output: -3
```

**Notes:**
- Always rounds toward negative infinity
- `sakafu(-2.1)` = -3 (not -2)

---

### dari (Ceiling)
Rounds a number up to the nearest integer.

**Syntax:**
```kwenda
dari(x)
```

**Parameters:**
- `x` (number): The value to ceil

**Returns:**
- Integer: The smallest integer ≥ x

**Examples:**
```kwenda
andika(dari(3.1))      // Output: 4
andika(dari(3.9))      // Output: 4
andika(dari(-2.9))     // Output: -2
andika(dari(-2.1))     // Output: -2
```

**Notes:**
- Always rounds toward positive infinity
- `dari(-2.9)` = -2 (not -3)

---

### kata (Truncate)
Removes the decimal part of a number, rounding toward zero.

**Syntax:**
```kwenda
kata(x)
```

**Parameters:**
- `x` (number): The value to truncate

**Returns:**
- Integer: The integer part of x

**Examples:**
```kwenda
andika(kata(3.9))      // Output: 3
andika(kata(3.1))      // Output: 3
andika(kata(-2.9))     // Output: -2
andika(kata(-2.1))     // Output: -2
```

**Notes:**
- Rounds toward zero
- `kata(-2.9)` = -2, `sakafu(-2.9)` = -3

---

## Modulo and Remainder

### bakiza (Modulo)
Computes the floating-point modulo (remainder of division).

**Syntax:**
```kwenda
bakiza(x, y)
```

**Parameters:**
- `x` (number): Dividend
- `y` (number): Divisor

**Returns:**
- The IEEE 754 floating-point remainder of x/y

**Examples:**
```kwenda
andika(bakiza(7.5, 2))     // Output: 1.5
andika(bakiza(10, 3))      // Output: 1
andika(bakiza(-10, 3))     // Output: -1
andika(bakiza(5.5, 2.5))   // Output: 0.5
```

**Notes:**
- Throws error if y = 0
- Result has same sign as x
- Equivalent to `x - floor(x/y) * y`

---

### salio (Remainder)
Computes the IEEE 754 remainder operation.

**Syntax:**
```kwenda
salio(x, y)
```

**Parameters:**
- `x` (number): Dividend
- `y` (number): Divisor

**Returns:**
- The IEEE 754 remainder of x/y

**Examples:**
```kwenda
andika(salio(7.5, 2.5))    // Output: 0
andika(salio(10, 3))       // Output: 1
andika(salio(-10, 3))      // Output: -1
```

**Notes:**
- Different from `bakiza` for certain values
- Uses IEEE 754 remainder definition
- Throws error if y = 0

---

## Sign Manipulation

### kopanja (Copysign)
Returns a value with the magnitude of x and the sign of y.

**Syntax:**
```kwenda
kopanja(magnitude, sign)
```

**Parameters:**
- `magnitude` (number): The magnitude value
- `sign` (number): The sign source

**Returns:**
- A number with magnitude of x and sign of y

**Examples:**
```kwenda
andika(kopanja(5, 1))      // Output: 5
andika(kopanja(5, -1))     // Output: -5
andika(kopanja(-5, 1))     // Output: 5
andika(kopanja(3.14, -2))  // Output: -3.14
```

**Use Cases:**
- Ensuring a value has a specific sign
- Flipping signs conditionally
- Mathematical transformations

---

## Approximate Comparisons

### karibia (Approximately Equal)
Checks if two floating-point numbers are approximately equal.

**Syntax:**
```kwenda
karibia(x, y)
karibia(x, y, epsilon)
```

**Parameters:**
- `x` (number): First value
- `y` (number): Second value
- `epsilon` (number, optional): Tolerance (default: 1e-9)

**Returns:**
- Boolean: `true` if |x - y| < epsilon

**Examples:**
```kwenda
andika(karibia(1.0, 1.0000000001))      // Output: true
andika(karibia(1.0, 1.1))               // Output: false
andika(karibia(0.1, 0.10001, 0.001))    // Output: true
andika(karibia(0.1, 0.10001, 0.00001))  // Output: false
```

**Use Cases:**
- Comparing floating-point results
- Avoiding precision errors
- Testing numerical algorithms

**Notes:**
- Default epsilon is 1e-9
- Custom epsilon must be positive
- Essential for floating-point comparisons

---

## Type Checking

### ni_kamili (Is Whole Number)
Checks if a number is a whole number (integer).

**Syntax:**
```kwenda
ni_kamili(x)
```

**Parameters:**
- `x` (number): The value to check

**Returns:**
- Boolean: `true` if x has no fractional part

**Examples:**
```kwenda
andika(ni_kamili(5.0))     // Output: true
andika(ni_kamili(5.5))     // Output: false
andika(ni_kamili(-3.0))    // Output: true
andika(ni_kamili(0.1))     // Output: false
```

---

### ni_usawa (Is Finite)
Checks if a number is finite (not infinite and not NaN).

**Syntax:**
```kwenda
ni_usawa(x)
```

**Parameters:**
- `x` (number): The value to check

**Returns:**
- Boolean: `true` if x is finite

**Examples:**
```kwenda
andika(ni_usawa(10))         // Output: true
andika(ni_usawa(3.14))       // Output: true
```

**Notes:**
- Returns `false` for infinity
- Returns `false` for NaN
- Useful for validation

---

### ni_bila_kikomo (Is Infinite)
Checks if a number is infinite.

**Syntax:**
```kwenda
ni_bila_kikomo(x)
```

**Parameters:**
- `x` (number): The value to check

**Returns:**
- Boolean: `true` if x is ±∞

**Examples:**
```kwenda
namba result = 1 / 0
andika(ni_bila_kikomo(result))  // Would be true for infinity
```

---

### ni_sio_namba (Is NaN)
Checks if a value is NaN (Not a Number).

**Syntax:**
```kwenda
ni_sio_namba(x)
```

**Parameters:**
- `x` (number): The value to check

**Returns:**
- Boolean: `true` if x is NaN

**Examples:**
```kwenda
andika(ni_sio_namba(5))        // Output: false
andika(ni_sio_namba(3.14))     // Output: false
```

---

## Value Clamping

### upeo (Clamp)
Restricts a value to a specified range [min, max].

**Syntax:**
```kwenda
upeo(value, min, max)
```

**Parameters:**
- `value` (number): The value to clamp
- `min` (number): Minimum bound
- `max` (number): Maximum bound

**Returns:**
- The clamped value within [min, max]

**Examples:**
```kwenda
andika(upeo(5, 0, 10))      // Output: 5
andika(upeo(-5, 0, 10))     // Output: 0
andika(upeo(15, 0, 10))     // Output: 10
andika(upeo(7.5, 2.5, 8.5)) // Output: 7.5
```

**Use Cases:**
- Constraining user input
- Keeping values in valid ranges
- Game development (health, score limits)
- UI boundaries

**Notes:**
- Returns `min` if value < min
- Returns `max` if value > max
- Returns `value` if min ≤ value ≤ max
- Throws error if min > max

---

## Part Extraction

### sehemu_desimali (Decimal Part)
Extracts the fractional/decimal part of a number.

**Syntax:**
```kwenda
sehemu_desimali(x)
```

**Parameters:**
- `x` (number): The input number

**Returns:**
- Float: The fractional part of x

**Examples:**
```kwenda
andika(sehemu_desimali(3.14159))   // Output: 0.14159
andika(sehemu_desimali(-2.718))    // Output: -0.718
andika(sehemu_desimali(5.0))       // Output: 0
```

**Notes:**
- Result has same sign as input
- `sehemu_desimali(-2.5)` = -0.5 (not 0.5)
- Returns 0 for whole numbers

---

### sehemu_kamili (Integer Part)
Extracts the integer part of a number.

**Syntax:**
```kwenda
sehemu_kamili(x)
```

**Parameters:**
- `x` (number): The input number

**Returns:**
- Integer: The integer part of x

**Examples:**
```kwenda
andika(sehemu_kamili(3.14159))   // Output: 3
andika(sehemu_kamili(-2.718))    // Output: -2
andika(sehemu_kamili(5.0))       // Output: 5
```

**Notes:**
- Equivalent to `kata(x)` but optimized
- Truncates toward zero

---

## Interpolation

### mzunguko (Linear Interpolation / Lerp)
Performs linear interpolation between two values.

**Syntax:**
```kwenda
mzunguko(start, end)          // Midpoint (t=0.5)
mzunguko(start, end, t)       // Custom interpolation
```

**Parameters:**
- `start` (number): Starting value
- `end` (number): Ending value
- `t` (number, optional): Interpolation factor (default: 0.5)

**Returns:**
- Interpolated value: `start + t * (end - start)`

**Examples:**
```kwenda
andika(mzunguko(0, 10))         // Output: 5 (midpoint)
andika(mzunguko(0, 10, 0.5))    // Output: 5
andika(mzunguko(0, 10, 0.25))   // Output: 2.5
andika(mzunguko(0, 10, 0.75))   // Output: 7.5
andika(mzunguko(100, 200, 0.3)) // Output: 130
```

**Use Cases:**
- Animation easing
- Color blending
- Smooth transitions
- Game development

**Notes:**
- `t = 0` returns `start`
- `t = 1` returns `end`
- `t = 0.5` returns midpoint
- Works with any t value (can extrapolate)

---

## Complete Example

```kwenda
andika("=== Float Utilities Demo ===")

andika("\nRounding:")
namba x = 3.7
andika("pindika(3.7) =", pindika(x))
andika("sakafu(3.7) =", sakafu(x))
andika("dari(3.7) =", dari(x))
andika("kata(3.7) =", kata(x))

andika("\nModulo:")
andika("bakiza(10, 3) =", bakiza(10, 3))

andika("\nSign control:")
andika("kopanja(5, -1) =", kopanja(5, -1))

andika("\nComparison:")
namba a = 0.1 + 0.2
namba b = 0.3
andika("0.1 + 0.2 == 0.3?", karibia(a, b))

andika("\nType checking:")
andika("ni_kamili(5.0) =", ni_kamili(5.0))
andika("ni_kamili(5.1) =", ni_kamili(5.1))

andika("\nClamping:")
namba score = 150
andika("upeo(150, 0, 100) =", upeo(score, 0, 100))

andika("\nPart extraction:")
namba pi = 3.14159
andika("sehemu_kamili(pi) =", sehemu_kamili(pi))
andika("sehemu_desimali(pi) =", sehemu_desimali(pi))

andika("\nInterpolation:")
andika("mzunguko(0, 100, 0.25) =", mzunguko(0, 100, 0.25))
```

---

## Summary Table

| Function | Purpose | Example |
|----------|---------|---------|
| `pindika(x)` | Round to nearest | `pindika(3.5)` → 4 |
| `sakafu(x)` | Round down | `sakafu(3.9)` → 3 |
| `dari(x)` | Round up | `dari(3.1)` → 4 |
| `kata(x)` | Truncate | `kata(-3.9)` → -3 |
| `bakiza(x,y)` | Modulo | `bakiza(10,3)` → 1 |
| `salio(x,y)` | Remainder | `salio(10,3)` → 1 |
| `kopanja(x,y)` | Copy sign | `kopanja(5,-1)` → -5 |
| `karibia(x,y)` | Approx equal | `karibia(1.0, 1.0001)` → true |
| `ni_kamili(x)` | Is integer | `ni_kamili(5.0)` → true |
| `ni_usawa(x)` | Is finite | `ni_usawa(10)` → true |
| `ni_bila_kikomo(x)` | Is infinite | Check for ±∞ |
| `ni_sio_namba(x)` | Is NaN | Check for NaN |
| `upeo(v,min,max)` | Clamp | `upeo(15,0,10)` → 10 |
| `sehemu_desimali(x)` | Decimal part | `sehemu_desimali(3.14)` → 0.14 |
| `sehemu_kamili(x)` | Integer part | `sehemu_kamili(3.14)` → 3 |
| `mzunguko(s,e,t)` | Interpolate | `mzunguko(0,10,0.5)` → 5 |

---

## Best Practices

1. **Use `karibia()` for float comparisons**
   ```kwenda
   // Bad
   kama (a == b) { ... }
   
   // Good
   kama (karibia(a, b)) { ... }
   ```

2. **Validate input ranges with `upeo()`**
   ```kwenda
   namba age = soma()
   age = upeo(age, 0, 150)  // Ensure valid age
   ```

3. **Check for special values**
   ```kwenda
   kama (ni_usawa(result)) {
       // Safe to use result
   } sivyo {
       andika("Error: Invalid result")
   }
   ```

4. **Use appropriate rounding**
   - `pindika()` for general rounding
   - `sakafu()` for array indices
   - `dari()` for ceiling division
   - `kata()` for integer conversion

---

## Version Information

Float utilities are part of Kwenda's standard library and are available in all versions.
