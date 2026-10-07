# String Interpolation in Kwenda

## Overview

String interpolation allows you to embed variables and expressions directly into strings using f-string syntax. This feature makes string formatting more readable and convenient.

## Syntax

```swahili
f"text {expression} more text"
```

- Prefix the string with `f`
- Wrap expressions in curly braces `{}`
- Expressions are evaluated and converted to strings

## Basic Usage

### Simple Variables

```swahili
maneno jina = "Amina"
namba umri = 25

andika(f"Jina langu ni {jina}")
// Output: Jina langu ni Amina

andika(f"Nina miaka {umri}")
// Output: Nina miaka 25

andika(f"Habari, mimi ni {jina}, umri wangu ni {umri}")
// Output: Habari, mimi ni Amina, umri wangu ni 25
```

## Expressions

### Arithmetic Operations

```swahili
namba x = 10
namba y = 20

andika(f"Jumla: {x + y}")           // Jumla: 30
andika(f"Tofauti: {x - y}")         // Tofauti: -10
andika(f"Bidhaa: {x * y}")          // Bidhaa: 200
andika(f"Gawanya: {y / x}")         // Gawanya: 2
andika(f"Wastani: {(x + y) / 2}")   // Wastani: 15
```

### Comparison Operations

```swahili
namba a = 15
namba b = 10

andika(f"{a} > {b} = {a > b}")      // 15 > 10 = kweli
andika(f"{a} == {b} = {a == b}")    // 15 == 10 = uwongo
```

### Function Calls

```swahili
maneno neno = "habari"

andika(f"Urefu: {urefu(neno)}")                    // Urefu: 6
andika(f"Herufi kubwa: {herufi_kubwa(neno)}")      // Herufi kubwa: HABARI
andika(f"Herufi ndogo: {herufi_ndogo(neno)}")      // Herufi ndogo: habari
```

## Data Types

### Booleans

```swahili
boolean ni_kweli = kweli
boolean ni_uwongo = uwongo

andika(f"Thamani: {ni_kweli}")      // Thamani: kweli
andika(f"Thamani: {ni_uwongo}")     // Thamani: uwongo
```

### Arrays

```swahili
orodha namba namba = [1, 2, 3, 4, 5]

andika(f"Orodha: {namba}")          // Orodha: [1, 2, 3, 4, 5]
andika(f"Urefu: {urefu_orodha(namba)}")  // Urefu: 5
andika(f"Kwanza: {namba[0]}")       // Kwanza: 1
andika(f"Slice: {namba[1:3]}")      // Slice: [2, 3]
```

### Tuples

```swahili
tuple namba point = (10, 20)

andika(f"Point: {point}")           // Point: (10, 20)
andika(f"X: {pata_tuple(point, 0)}")  // X: 10
```

### Sets

```swahili
seti namba set_data = {1, 2, 3}

andika(f"Set: {set_data}")          // Set: {1, 2, 3}
andika(f"Ukubwa: {ukubwa_seti(set_data)}")  // Ukubwa: 3
```

## Number Formatting

### Integers

```swahili
namba idadi = 42

andika(f"Idadi: {idadi}")           // Idadi: 42
andika(f"Mara mbili: {idadi * 2}")  // Mara mbili: 84
```

### Floats

```swahili
namba pi = 3.14159

andika(f"Pi: {pi}")                 // Pi: 3.14159
andika(f"Duara: {2 * pi * 5}")      // Duara: 31.4159
```

Floats are automatically formatted to remove unnecessary trailing zeros.

## Practical Examples

### Receipt/Invoice

```swahili
maneno bidhaa = "Kitabu"
namba bei = 15000
namba idadi = 3

andika(f"=== RISITI ===")
andika(f"Bidhaa: {bidhaa}")
andika(f"Bei: TSh {bei}")
andika(f"Idadi: {idadi}")
andika(f"Jumla: TSh {bei * idadi}")

// Output:
// === RISITI ===
// Bidhaa: Kitabu
// Bei: TSh 15000
// Idadi: 3
// Jumla: TSh 45000
```

### User Profile

```swahili
maneno jina = "Hassan"
namba umri = 30
maneno mji = "Arusha"

andika(f"Jina: {jina}, Umri: {umri}, Mji: {mji}")
// Output: Jina: Hassan, Umri: 30, Mji: Arusha
```

### Progress Display

```swahili
namba imekamilika = 75
namba jumla = 100

andika(f"Maendeleo: {imekamilika}/{jumla} ({imekamilika * 100 / jumla}%)")
// Output: Maendeleo: 75/100 (75%)
```

### Temperature Conversion

```swahili
namba celsius = 25
namba fahrenheit = (celsius * 9 / 5) + 32

andika(f"{celsius}°C = {fahrenheit}°F")
// Output: 25°C = 77°F
```

### Mathematical Formulas

```swahili
namba r = 5

andika(f"Radius: {r}")
andika(f"Diameter: {r * 2}")
andika(f"Circumference: {2 * 3.14 * r}")
andika(f"Area: {3.14 * r * r}")

// Output:
// Radius: 5
// Diameter: 10
// Circumference: 31.4
// Area: 78.5
```

## Complex Expressions

### Multiple Variables

```swahili
namba a = 10
namba b = 20
namba c = 30

andika(f"a={a}, b={b}, c={c}, sum={a + b + c}")
// Output: a=10, b=20, c=30, sum=60
```

### Nested Calculations

```swahili
namba pato = 1000
namba gharama = 750

andika(f"Pato: {pato}, Gharama: {gharama}, Faida: {pato - gharama}")
andika(f"Asilimia ya faida: {(pato - gharama) * 100 / pato}%")

// Output:
// Pato: 1000, Gharama: 750, Faida: 250
// Asilimia ya faida: 25%
```

## Comparison with Concatenation

### Without Interpolation (Old Way)

```swahili
maneno jina = "Amina"
namba umri = 25

// Using concatenation
andika("Jina: " + jina + ", Umri: " + umri)  // Error: can't concatenate string and number
```

### With Interpolation (New Way)

```swahili
andika(f"Jina: {jina}, Umri: {umri}")  // Works perfectly!
```

## Best Practices

### 1. Use for Readability

❌ **Don't** do this:
```swahili
andika("Jumla ya " + x + " na " + y + " ni " + (x + y))  // Error-prone
```

✅ **Do** this:
```swahili
andika(f"Jumla ya {x} na {y} ni {x + y}")  // Clear and readable
```

### 2. Complex Expressions

For complex expressions, calculate first:

```swahili
// Good - clear and maintainable
namba wastani = (a + b + c) / 3
andika(f"Wastani: {wastani}")

// Also OK - direct calculation
andika(f"Wastani: {(a + b + c) / 3}")
```

### 3. Keep It Simple

If the expression is very complex, consider breaking it down:

```swahili
namba faida_halisi = pato - gharama - kodi
namba asilimia = faida_halisi * 100 / pato

andika(f"Faida: {faida_halisi} ({asilimia}%)")
```

## Limitations

### Current Implementation

- ✅ Variables
- ✅ Arithmetic expressions
- ✅ Comparison expressions
- ✅ Function calls
- ✅ Array/tuple/set access
- ✅ All data types
- ❌ Escaped braces `{{` and `}}` (future feature)
- ❌ Format specifiers (future feature)

## Format Specifiers (Future)

Planned features:

```swahili
// Decimal places
f"Pi: {pi:.2f}"                     // Pi: 3.14

// Padding
f"Number: {num:5d}"                 // Number:    42

// Alignment
f"Left: {text:<10}"                 // Left: text      
f"Right: {text:>10}"                // Right:      text
```

## Error Handling

### Invalid Expressions

If an expression can't be evaluated, it will be shown as-is or cause an error:

```swahili
// Undefined variable
andika(f"Value: {undefined_var}")   // Error: variable not found
```

### Type Errors

Most types convert to strings automatically:

```swahili
namba x = 10
maneno y = "20"

andika(f"Sum: {x + y}")  // May cause error depending on type rules
```

## Performance

- Interpolation is evaluated at runtime
- Each `{}` expression is evaluated separately
- For frequently used strings, consider pre-calculating

## See Also

- **STRINGS.md** - String functions
- **unganisha()** - String concatenation
- **FORMAT_SPECIFIERS.md** - Future formatting options (planned)

## Examples

See `examples/string_interpolation_demo.swh` for comprehensive examples.
