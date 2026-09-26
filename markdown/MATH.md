# Kwenda Mathematics Reference Guide

Complete documentation of all mathematical operations and functions available in the Kwenda Swahili programming language.

## 📚 Table of Contents

1. [Rounding Operations](#rounding-operations)
2. [Root Operations](#root-operations)
3. [Exponential & Logarithm Operations](#exponential--logarithm-operations)
4. [Trigonometric Functions](#trigonometric-functions)
5. [Bitwise Operations](#bitwise-operations)
6. [Modular Arithmetic](#modular-arithmetic)
7. [Advanced Rounding](#advanced-rounding)
8. [Random Number Generation](#random-number-generation)

---

## Rounding Operations

Kwenda provides four built-in functions for rounding numbers, each with different behavior:

### Basic Usage

```swahili
# Floor - round down to nearest integer
namba x = chini(4.7)    # x = 4
namba y = chini(-2.3)   # y = -3 (rounds toward negative infinity)

# Ceiling - round up to nearest integer  
namba a = juu(4.2)      # a = 5
namba b = juu(-2.7)     # b = -2 (rounds toward positive infinity)

# Round - round to nearest integer
namba p = zunguka(4.5)  # p = 5 (rounds away from zero for .5)
namba q = zunguka(4.3)  # q = 4

# Truncate - remove decimal part (round toward zero)
namba m = kata_desimali(4.9)   # m = 4
namba n = kata_desimali(-4.9)  # n = -4 (rounds toward zero)
```

### Key Differences for Negative Numbers

The main difference between these functions becomes apparent with negative numbers:

```swahili
namba value = -2.7

andika("chini(-2.7) =", chini(value))           # -3 (rounds DOWN)
andika("kata_desimali(-2.7) =", kata_desimali(value))  # -2 (rounds toward ZERO)
andika("juu(-2.7) =", juu(value))               # -2 (rounds UP)
andika("zunguka(-2.7) =", zunguka(value))       # -3 (rounds to nearest)
```

**Comparison Table:**

| Function | 4.7 | -2.7 | Description |
|----------|-----|------|-------------|
| `chini` (floor) | 4 | -3 | Always rounds toward negative infinity |
| `juu` (ceil) | 5 | -2 | Always rounds toward positive infinity |
| `zunguka` (round) | 5 | -3 | Rounds to nearest integer |
| `kata_desimali` (truncate) | 4 | -2 | Always rounds toward zero |

### Using with Math Module Constants

```swahili
leta "modules/math.swh"

kazi kuu() {
    andika("PI =", math.PI)                  # 3.14159265359
    andika("chini(PI) =", chini(math.PI))    # 3
    andika("juu(PI) =", juu(math.PI))        # 4
    andika("zunguka(PI) =", zunguka(math.PI)) # 3
    
    andika("E =", math.E)                    # 2.71828182846
    andika("chini(E) =", chini(math.E))      # 2
    andika("juu(E) =", juu(math.E))          # 3
}
```

### IEEE 754 Rounding Behavior

The `zunguka` (round) function follows IEEE 754 standard rounding:
- Values like 4.5 round to 5 (away from zero)
- The rounding is consistent and predictable for all edge cases

---

## Root Operations

Kwenda provides three built-in functions for calculating roots, each optimized for different use cases:

### Basic Usage

```swahili
# Square root - mzizi_mraba(number)
namba x = mzizi_mraba(16)     # x = 4
namba y = mzizi_mraba(2)      # y = 1.4142135623730951
namba z = mzizi_mraba(100)    # z = 10

# Cube root - mzizi_mchemraba(number)
namba a = mzizi_mchemraba(27)    # a = 3
namba b = mzizi_mchemraba(-8)    # b = -2 (works with negatives!)
namba c = mzizi_mchemraba(10)    # c = 2.154434690031884

# Nth root - mzizi(number, root)
namba p = mzizi(16, 2)    # p = 4 (square root)
namba q = mzizi(8, 3)     # q = 2 (cube root)
namba r = mzizi(32, 5)    # r = 2 (5th root)
namba s = mzizi(-8, 3)    # s = -2 (odd roots of negatives work)
```

### Handling Negative Numbers

Root functions have different behaviors with negative numbers:

```swahili
# Square root - ONLY non-negative numbers
namba valid = mzizi_mraba(25)      # 5 - OK
# namba error = mzizi_mraba(-25)   # ERROR! Cannot calculate

# Cube root - ANY number (positive or negative)
namba pos = mzizi_mchemraba(27)    # 3
namba neg = mzizi_mchemraba(-27)   # -3 - Works fine!

# Nth root - depends on whether root is even or odd
namba odd_root = mzizi(-8, 3)      # -2 - Odd root, works
# namba even_root = mzizi(-16, 4)  # ERROR! Even root of negative
```

### Error Handling

```swahili
jaribu {
    namba x = mzizi_mraba(-4)  # Will throw error
} shika (error) {
    andika("Cannot calculate square root of negative number")
}

jaribu {
    namba y = mzizi(-16, 4)    # Will throw error (even root)
} shika (error) {
    andika("Cannot calculate even root of negative number")
}

# This works fine - odd root of negative
namba z = mzizi(-27, 3)        # -3 (no error)
```

### Comparison of Root Functions

**When to use each function:**

| Function | Use Case | Negative Numbers | Performance |
|----------|----------|------------------|-------------|
| `mzizi_mraba` | Square roots only | ❌ Not allowed | Fastest |
| `mzizi_mchemraba` | Cube roots only | ✅ Allowed | Fast |
| `mzizi` | Any root (2nd, 3rd, 4th, etc.) | ⚠️ Odd roots only | Flexible |

```swahili
# These are equivalent:
namba a = mzizi_mraba(16)      # 4
namba b = mzizi(16, 2)         # 4 (same result)

# These are equivalent:
namba c = mzizi_mchemraba(8)   # 2
namba d = mzizi(8, 3)          # 2 (same result)

# Only mzizi can do 5th, 7th, 10th roots, etc:
namba e = mzizi(32, 5)         # 2 (5th root)
namba f = mzizi(128, 7)        # 2 (7th root)
```

### Using with Math Module Constants

```swahili
leta "modules/math.swh"

kazi kuu() {
    andika("PI =", math.PI)                        # 3.14159265359
    andika("Square root of PI =", mzizi_mraba(math.PI))      # 1.772...
    andika("Cube root of PI =", mzizi_mchemraba(math.PI))    # 1.464...
    andika("4th root of PI =", mzizi(math.PI, 4))            # 1.331...
    
    andika("E =", math.E)                          # 2.71828182846
    andika("Square root of E =", mzizi_mraba(math.E))        # 1.648...
}
```

### Combining Roots with Rounding

```swahili
# Calculate and round in one expression
namba sqrt2 = mzizi_mraba(2)           # 1.4142135623730951
andika("Rounded:", zunguka(sqrt2))     # 1

# Or directly
andika("Floor of sqrt(10):", chini(mzizi_mraba(10)))     # 3
andika("Ceil of cbrt(10):", juu(mzizi_mchemraba(10)))    # 3

# Using in calculations
namba a = mzizi_mraba(16) + mzizi_mraba(9)   # 4 + 3 = 7
namba b = mzizi_mraba(16) * mzizi_mraba(9)   # 4 * 3 = 12
```

### Special Cases

```swahili
# Root of 0 is always 0
andika(mzizi_mraba(0))        # 0
andika(mzizi_mchemraba(0))    # 0

# Root of 1 is always 1
andika(mzizi_mraba(1))        # 1
andika(mzizi_mchemraba(1))    # 1
andika(mzizi(1, 5))           # 1
andika(mzizi(1, 100))         # 1

# Perfect roots return integers
andika(mzizi_mraba(25))       # 5 (not 5.0)
andika(mzizi_mchemraba(8))    # 2 (not 2.0)

# Non-perfect roots return floats
andika(mzizi_mraba(2))        # 1.4142135623730951
```

---

## Exponential & Logarithm Operations

Kwenda provides four built-in functions for exponential and logarithmic calculations, essential for scientific computing, growth calculations, and data analysis:

### Basic Usage

```swahili
# Exponential - exp(x) returns e^x
namba a = exp(0)      # a = 1
namba b = exp(1)      # b = 2.718281828... (which is e)
namba c = exp(2)      # c = 7.389...
namba d = exp(-1)     # d = 0.367... (which is 1/e)

# Natural logarithm - log_asili(x) returns ln(x)
namba p = log_asili(1)      # p = 0
namba q = log_asili(2)      # q = 0.693...
namba r = log_asili(10)     # r = 2.302...
namba s = log_asili(100)    # s = 4.605...

# Base-10 logarithm - log10(x)
namba m = log10(1)      # m = 0
namba n = log10(10)     # n = 1
namba o = log10(100)    # o = 2
namba p = log10(1000)   # p = 3

# Base-2 logarithm - log2(x)
namba w = log2(1)       # w = 0
namba x = log2(2)       # x = 1
namba y = log2(8)       # y = 3 (because 2^3 = 8)
namba z = log2(16)      # z = 4 (because 2^4 = 16)
```

### Inverse Relationships

Exponential and logarithm functions are inverses of each other:

```swahili
# exp and log_asili are inverses
andika(exp(log_asili(5)))       # 5
andika(log_asili(exp(3)))       # 3

# This means:
# If y = exp(x), then x = log_asili(y)
# If y = log_asili(x), then x = exp(y)

# Example: solving exponential equations
namba x = log_asili(10)         # What power of e gives 10?
andika("e^", x, "=", exp(x))    # Verify: e^x = 10
```

### Logarithm Properties

```swahili
# Property 1: log(a * b) = log(a) + log(b)
namba a = 8
namba b = 16
andika(log_asili(a * b))                    # 4.852...
andika(log_asili(a) + log_asili(b))         # 4.852... (same)

# Property 2: log(a / b) = log(a) - log(b)
andika(log_asili(a) - log_asili(b))         # -0.693...

# Property 3: log(a^n) = n * log(a)
namba n = 3
andika(log_asili(a * a * a))                # 6.238...
andika(n * log_asili(a))                    # 6.238... (same)
```

### Error Handling

Logarithms only work with positive numbers:

```swahili
# Valid: positive numbers
namba x = log_asili(5)      # 1.609...
namba y = log10(100)        # 2
namba z = log2(8)           # 3

# Invalid: zero or negative numbers
jaribu {
    namba bad1 = log_asili(0)       # ERROR
} shika (error) {
    andika("Cannot take log of zero")
}

jaribu {
    namba bad2 = log10(-5)          # ERROR
} shika (error) {
    andika("Cannot take log of negative")
}

# Exponential overflow
jaribu {
    namba big = exp(710)            # ERROR (too large)
} shika (error) {
    andika("Result would overflow")
}
```

### Using with Math Constants

```swahili
leta "modules/math.swh"

kazi kuu() {
    # E is Euler's number (base of natural log)
    andika("E =", math.E)                   # 2.71828...
    andika("exp(1) =", exp(1))              # 2.71828... (same as E)
    andika("log_asili(E) =", log_asili(math.E))  # 1
    
    # PI with exponential/log
    andika("PI =", math.PI)                 # 3.14159...
    andika("exp(PI) =", exp(math.PI))       # 23.140...
    andika("log_asili(PI) =", log_asili(math.PI))  # 1.144...
}
```

### Practical Applications

**Compound Growth:**
```swahili
# Continuous compound interest: A = P * e^(rt)
namba principal = 1000
namba rate = 0.05              # 5% annual rate
namba time = 10                # 10 years

namba amount = principal * exp(rate * time)
andika("After 10 years:", amount)  # $1648.72
```

**Half-Life Calculations:**
```swahili
# Half-life: time for quantity to reduce to half
namba half_life = log_asili(0.5)
andika("Time to halve:", half_life)  # -0.693...

# Double-time: time for quantity to double
namba double_time = log_asili(2)
andika("Time to double:", double_time)  # 0.693...
```

**Powers of 2 (Computer Science):**
```swahili
# How many bits needed to represent a number?
namba value = 1024
namba bits = juu(log2(value))
andika("Bits needed for", value, ":", bits)  # 10 bits
```

**Orders of Magnitude:**
```swahili
# How many digits in a number?
namba num = 1000000
namba digits = juu(log10(num)) + 1
andika("Digits in", num, ":", digits)  # 7 digits
```

### Change of Base Formula

Convert between different logarithm bases:

```swahili
# log_b(x) = log_asili(x) / log_asili(b)

# Example: Calculate log base 5 of 125
namba x = 125
namba base = 5
namba result = log_asili(x) / log_asili(base)
andika("log_5(125) =", result)  # 3 (because 5^3 = 125)

# Verify with our log functions:
# log2(x) = log_asili(x) / log_asili(2)
namba test = 8
andika("log2(8) direct:", log2(test))           # 3
andika("Using formula:", log_asili(test) / log_asili(2))  # 3 (same)
```

### Combining with Other Math Functions

```swahili
# Exponential with rounding
namba exp_val = exp(2.5)
andika("exp(2.5) =", exp_val)               # 12.182...
andika("Rounded:", zunguka(exp_val))        # 12

# Logarithm with roots
namba num = 16
andika("log2(16) =", log2(num))             # 4
andika("sqrt(16) =", mzizi_mraba(num))      # 4
andika("log2(sqrt(16)) =", log2(mzizi_mraba(num)))  # 2

# Exponential growth with roots
namba growth = exp(log_asili(2))            # 2
andika("e^(ln(2)) =", growth)
```

### Comparison of Functions

| Function | Domain | Range | Use Case |
|----------|--------|-------|----------|
| `exp(x)` | All numbers | Positive only | Growth, compound interest |
| `log_asili(x)` | Positive only | All numbers | Natural processes, calculus |
| `log10(x)` | Positive only | All numbers | Orders of magnitude, pH, decibels |
| `log2(x)` | Positive only | All numbers | Computer science, information theory |

**Key Relationships:**
- `exp(log_asili(x)) = x` (for x > 0)
- `log_asili(exp(x)) = x` (for all x)
- `log10(x) = log_asili(x) / log_asili(10)`
- `log2(x) = log_asili(x) / log_asili(2)`

---

## Trigonometric Functions

Kwenda provides comprehensive trigonometric functions for angle calculations, including basic trig functions (sin, cos, tan), their inverses (asin, acos, atan, atan2), and angle conversion utilities (radians, degrees).

**Important:** All trigonometric functions work with angles in **radians**, not degrees. Use the `radians()` and `degrees()` functions to convert between the two.

### Angle Conversion

```swahili
# Convert degrees to radians
namba angle_rad = radians(90)      # π/2 ≈ 1.5708
andika("90° = ", angle_rad, " radians")

# Convert radians to degrees
namba angle_deg = degrees(1.5708)  # ≈ 90
andika("1.5708 radians = ", angle_deg, "°")

# Common conversions
andika("0° = ", radians(0), " rad")       # 0
andika("30° = ", radians(30), " rad")     # π/6 ≈ 0.5236
andika("45° = ", radians(45), " rad")     # π/4 ≈ 0.7854
andika("60° = ", radians(60), " rad")     # π/3 ≈ 1.0472
andika("90° = ", radians(90), " rad")     # π/2 ≈ 1.5708
andika("180° = ", radians(180), " rad")   # π ≈ 3.1416
andika("360° = ", radians(360), " rad")   # 2π ≈ 6.2832
```

### Basic Trigonometric Functions

```swahili
# Sine - sin(angle_in_radians)
andika("sin(0) = ", sin(0))                       # 0
andika("sin(30°) = ", sin(radians(30)))           # 0.5
andika("sin(45°) = ", sin(radians(45)))           # √2/2 ≈ 0.7071
andika("sin(90°) = ", sin(radians(90)))           # 1

# Cosine - cos(angle_in_radians)
andika("cos(0) = ", cos(0))                       # 1
andika("cos(60°) = ", cos(radians(60)))           # 0.5
andika("cos(45°) = ", cos(radians(45)))           # √2/2 ≈ 0.7071
andika("cos(90°) = ", cos(radians(90)))           # 0 (approximately)

# Tangent - tan(angle_in_radians)
andika("tan(0) = ", tan(0))                       # 0
andika("tan(45°) = ", tan(radians(45)))           # 1
andika("tan(30°) = ", tan(radians(30)))           # √3/3 ≈ 0.5774
andika("tan(60°) = ", tan(radians(60)))           # √3 ≈ 1.7321
```

### Inverse Trigonometric Functions

Inverse trig functions return angles in **radians**:

```swahili
# Arc sine - asin(x) returns angle where sin(angle) = x
# Domain: -1 ≤ x ≤ 1, Range: -π/2 to π/2
namba angle1 = asin(0.5)                          # π/6 ≈ 0.5236 radians (30°)
namba angle2 = asin(1)                            # π/2 ≈ 1.5708 radians (90°)
andika("asin(0.5) = ", degrees(angle1), "°")      # 30°

# Arc cosine - acos(x) returns angle where cos(angle) = x
# Domain: -1 ≤ x ≤ 1, Range: 0 to π
namba angle3 = acos(0.5)                          # π/3 ≈ 1.0472 radians (60°)
namba angle4 = acos(0)                            # π/2 ≈ 1.5708 radians (90°)
andika("acos(0.5) = ", degrees(angle3), "°")      # 60°

# Arc tangent - atan(x) returns angle where tan(angle) = x
# Domain: all real numbers, Range: -π/2 to π/2
namba angle5 = atan(1)                            # π/4 ≈ 0.7854 radians (45°)
namba angle6 = atan(0)                            # 0 radians (0°)
andika("atan(1) = ", degrees(angle5), "°")        # 45°

# Two-argument arc tangent - atan2(y, x) for proper quadrant handling
# Returns angle from origin to point (x, y)
namba angle7 = atan2(1, 1)                        # π/4 ≈ 0.7854 (Quadrant I)
namba angle8 = atan2(1, -1)                       # 3π/4 ≈ 2.3562 (Quadrant II)
namba angle9 = atan2(-1, -1)                      # -3π/4 ≈ -2.3562 (Quadrant III)
namba angle10 = atan2(-1, 1)                      # -π/4 ≈ -0.7854 (Quadrant IV)
```

### Pythagorean Identity

The fundamental relationship sin²(x) + cos²(x) = 1:

```swahili
namba angle = radians(30)
namba s = sin(angle)
namba c = cos(angle)

# Verify the identity
namba sum = s*s + c*c
andika("sin²(30°) + cos²(30°) = ", sum)  # Always equals 1
```

### Practical Applications

**Right Triangle Calculations:**

```swahili
# Given: right triangle with hypotenuse = 10, angle = 30°
namba hypotenuse = 10
namba angle = radians(30)

# Calculate sides
namba opposite = hypotenuse * sin(angle)          # 5
namba adjacent = hypotenuse * cos(angle)          # 8.66

andika("Opposite side = ", opposite)
andika("Adjacent side = ", adjacent)

# Verify: opposite² + adjacent² = hypotenuse²
namba check = opposite*opposite + adjacent*adjacent
andika("Verification: ", check, " ≈ ", hypotenuse*hypotenuse)
```

**Finding Angle Between Two Points:**

```swahili
# Point coordinates
namba x = 4
namba y = 3

# Find angle from origin to point (x, y)
namba angle_rad = atan2(y, x)
namba angle_deg = degrees(angle_rad)

andika("Point (", x, ", ", y, ")")
andika("Angle from origin: ", angle_deg, "°")

# Calculate distance
namba distance = mzizi_mraba(x*x + y*y)  # 5
andika("Distance: ", distance)
```

**Navigation and Bearings:**

```swahili
# Calculate direction to travel from point A to point B
namba dx = 100  # East displacement
namba dy = 100  # North displacement

namba bearing_rad = atan2(dy, dx)
namba bearing_deg = degrees(bearing_rad)

andika("Bearing: ", bearing_deg, "° from East")

# Convert to compass bearing (from North)
namba compass_bearing = 90 - bearing_deg
andika("Compass bearing: ", compass_bearing, "° from North")
```

### Error Handling

```swahili
# asin and acos have restricted domains
jaribu {
    namba bad = asin(2)  # ERROR: must be in [-1, 1]
} shika kosa {
    andika("Domain error: asin requires input between -1 and 1")
}

jaribu {
    namba bad = acos(1.5)  # ERROR: must be in [-1, 1]
} shika kosa {
    andika("Domain error: acos requires input between -1 and 1")
}

# tan is undefined at π/2 + nπ
jaribu {
    namba bad = tan(radians(90))  # May produce very large number or error
} shika kosa {
    andika("tan(90°) is undefined (approaches infinity)")
}
```

---

## Bitwise Operations

Kwenda provides six bitwise operations for manipulating individual bits in integers. These are useful for low-level programming, optimization, flags, and cryptography.

### Basic Bitwise Operations

```swahili
# Bitwise AND - na_kidogo(a, b)
# Returns 1 where both operands have 1, else 0
namba result1 = na_kidogo(12, 10)      # 1100 & 1010 = 1000 = 8
namba result2 = na_kidogo(15, 7)       # 1111 & 0111 = 0111 = 7

# Bitwise OR - au_kidogo(a, b)
# Returns 1 where at least one operand has 1
namba result3 = au_kidogo(12, 10)      # 1100 | 1010 = 1110 = 14
namba result4 = au_kidogo(8, 4)        # 1000 | 0100 = 1100 = 12

# Bitwise XOR - ama_kidogo(a, b)
# Returns 1 where operands differ, else 0
namba result5 = ama_kidogo(12, 10)     # 1100 ^ 1010 = 0110 = 6
namba result6 = ama_kidogo(15, 15)     # 1111 ^ 1111 = 0000 = 0

# Bitwise NOT - si_kidogo(a)
# Inverts all bits
namba result7 = si_kidogo(0)           # ~0 = -1
namba result8 = si_kidogo(5)           # ~5 = -6

# Left Shift - gezo_kushoto(a, n)
# Shifts bits left by n positions (equivalent to multiplying by 2^n)
namba result9 = gezo_kushoto(5, 1)     # 0101 << 1 = 1010 = 10
namba result10 = gezo_kushoto(3, 2)    # 0011 << 2 = 1100 = 12
namba result11 = gezo_kushoto(1, 8)    # 0001 << 8 = 100000000 = 256

# Right Shift - gezo_kulia(a, n)
# Shifts bits right by n positions (equivalent to dividing by 2^n)
namba result12 = gezo_kulia(10, 1)     # 1010 >> 1 = 0101 = 5
namba result13 = gezo_kulia(12, 2)     # 1100 >> 2 = 0011 = 3
namba result14 = gezo_kulia(256, 8)    # >> 8 = 1
```

### Practical Applications

**Flag Management:**
```swahili
# Using bits as boolean flags
namba flags = 0
namba READ = gezo_kushoto(1, 0)        # Bit 0
namba WRITE = gezo_kushoto(1, 1)       # Bit 1
namba EXECUTE = gezo_kushoto(1, 2)     # Bit 2

# Set flags (OR operation)
flags = au_kidogo(flags, READ)         # Enable read
flags = au_kidogo(flags, WRITE)        # Enable write

# Check if flag is set (AND operation)
kama na_kidogo(flags, READ) == READ {
    andika("Read permission granted")
}

# Clear a flag (AND with NOT)
flags = na_kidogo(flags, si_kidogo(WRITE))  # Disable write
```

**Bit Counting:**
```swahili
# Count number of set bits (population count)
namba num = 15
namba count = 0

wakati num > 0 {
    kama na_kidogo(num, 1) == 1 {
        count = count + 1
    }
    num = gezo_kulia(num, 1)
}
andika("Set bits:", count)  # 4 (15 = 1111 in binary)
```

**Swapping Values Without Temp Variable:**
```swahili
namba a = 5
namba b = 7

# XOR swap
a = ama_kidogo(a, b)
b = ama_kidogo(a, b)
a = ama_kidogo(a, b)

andika("a =", a, ", b =", b)  # a = 7, b = 5
```

**Bitwise Comparison Table:**
| Operation | Example | Result | Binary |
|-----------|---------|--------|--------|
| AND | na_kidogo(12, 10) | 8 | 1000 |
| OR | au_kidogo(12, 10) | 14 | 1110 |
| XOR | ama_kidogo(12, 10) | 6 | 0110 |
| NOT | si_kidogo(5) | -6 | ...11111010 |
| Left Shift | gezo_kushoto(5, 1) | 10 | 1010 |
| Right Shift | gezo_kulia(10, 1) | 5 | 0101 |

---

## Modular Arithmetic

Modular arithmetic is essential for working with remainders, cryptography, and cyclic calculations. Kwenda provides three modular operations.

### Basic Modular Operations

```swahili
# Modulo - modulo(a, b)
# Returns remainder after division (handles floats)
namba r1 = modulo(10, 3)               # 1
namba r2 = modulo(20, 6)               # 2
namba r3 = modulo(7, 7)                # 0
namba r4 = modulo(15, 4)               # 3
namba r5 = modulo(100, 10)             # 0

# Remainder - baki(a, b)
# Returns remainder for integers
namba rem1 = baki(10, 3)               # 1
namba rem2 = baki(25, 5)               # 0
namba rem3 = baki(17, 5)               # 2
namba rem4 = baki(100, 7)              # 2

# Modular Power - tweza_modulo(base, exponent, modulus)
# Calculates (base^exponent) mod modulus efficiently
namba mp1 = tweza_modulo(2, 3, 5)      # 2^3 mod 5 = 8 mod 5 = 3
namba mp2 = tweza_modulo(3, 4, 5)      # 3^4 mod 5 = 81 mod 5 = 1
namba mp3 = tweza_modulo(7, 2, 11)     # 7^2 mod 11 = 49 mod 11 = 5
namba mp4 = tweza_modulo(2, 10, 1000)  # 2^10 mod 1000 = 1024 mod 1000 = 24
```

### Practical Applications

**Cyclic Sequences:**
```swahili
# Generate a repeating pattern
kwa i = 0; i < 20; i = i + 1 {
    namba pattern = modulo(i, 3)
    andika("Index", i, "-> Pattern", pattern)  # 0,1,2,0,1,2,...
}
```

**Day of Week Calculation:**
```swahili
# Calculate day of week (0-6) for scheduling
namba days_since_epoch = 5000
namba day_of_week = modulo(days_since_epoch, 7)
orodha siku = ["Jumapili", "Jumatatu", "Jumatano", "Kamis", "Ijumaa", "Jumamosi", "Jumapili"]
andika("Day:", siku(day_of_week))
```

**Modular Inverse (via Extended Euclidean Algorithm):**
```swahili
# Find modular multiplicative inverse
# If a * x ≡ 1 (mod m), then x is the inverse of a
# For prime moduli, use: tweza_modulo(a, m-2, m)

namba a = 3
namba m = 7
namba inverse = tweza_modulo(a, m - 2, m)
andika("Modular inverse of", a, "mod", m, "is", inverse)  # 5
andika("Verification:", modulo(a * inverse, m))  # 1
```

**Cryptographic Applications:**
```swahili
# RSA-style modular exponentiation (simplified)
namba message = 42
namba exponent = 17
namba modulus = 1000000007

namba encrypted = tweza_modulo(message, exponent, modulus)
andika("Encrypted:", encrypted)
```

---

## Advanced Rounding

Four rounding functions provide different strategies for converting floating-point numbers to integers, each with distinct behavior for negative numbers.

### Rounding Operations

```swahili
# Round - pindika(x)
# Rounds to nearest integer
namba a1 = pindika(3.4)                # 3
namba a2 = pindika(3.5)                # 4
namba a3 = pindika(3.6)                # 4
namba a4 = pindika(-2.7)               # -3

# Floor - sakafu(x)
# Rounds down toward negative infinity
namba b1 = sakafu(3.9)                 # 3
namba b2 = sakafu(3.1)                 # 3
namba b3 = sakafu(-2.3)                # -3
namba b4 = sakafu(-2.9)                # -3

# Ceiling - dari(x)
# Rounds up toward positive infinity
namba c1 = dari(3.1)                   # 4
namba c2 = dari(3.9)                   # 4
namba c3 = dari(-2.1)                  # -2
namba c4 = dari(-2.9)                  # -2

# Truncate - kata(x)
# Rounds toward zero (removes decimal part)
namba d1 = kata(3.9)                   # 3
namba d2 = kata(3.1)                   # 3
namba d3 = kata(-2.7)                  # -2
namba d4 = kata(-2.1)                  # -2
```

### Comparison Table

| Value | pindika | sakafu | dari | kata |
|-------|---------|--------|------|------|
| 3.2 | 3 | 3 | 4 | 3 |
| 3.7 | 4 | 3 | 4 | 3 |
| -2.3 | -2 | -3 | -2 | -2 |
| -2.7 | -3 | -3 | -2 | -2 |

### Practical Applications

**Currency Rounding:**
```swahili
namba price = 19.95
namba quantity = 3
namba total = price * quantity  # 59.85

namba rounded_total = pindika(total * 100) / 100  # 59.85
andika("Total to charge: TSh", rounded_total)
```

**Grid Positioning:**
```swahili
namba x_float = 3.7
namba y_float = 2.2

# Snap to grid (floor)
namba grid_x = sakafu(x_float)         # 3
namba grid_y = sakafu(y_float)         # 2
andika("Snapped to grid: (", grid_x, ", ", grid_y, ")")

# Snap to nearest grid (round)
namba snap_x = pindika(x_float)        # 4
namba snap_y = pindika(y_float)        # 2
andika("Snapped nearest: (", snap_x, ", ", snap_y, ")")
```

**Data Compression - Quantization:**
```swahili
# Reduce precision to 1 decimal place
namba sensor_value = 3.14159
namba quantized = pindika(sensor_value * 10) / 10  # 3.1
andika("Quantized value:", quantized)
```

---

## Random Number Generation

Generate random numbers with optional seeding for reproducible results. Three modes: float [0,1), integer [0,n), and integer range [min,max].

### Basic Random Operations

```swahili
# Random float [0.0, 1.0)
namba rand1 = nasibu()                 # 0.372... 
namba rand2 = nasibu()                 # 0.814...
namba rand3 = nasibu()                 # 0.521...

# Random integer [0, n)
namba rand_int1 = nasibu(10)           # 0-9
namba rand_int2 = nasibu(100)          # 0-99
namba rand_int3 = nasibu(6) + 1        # 1-6 (dice roll)

# Random integer [min, max]
namba rand_range1 = nasibu(1, 10)      # 1-10
namba rand_range2 = nasibu(50, 100)    # 50-100
namba rand_range3 = nasibu(-10, 10)    # -10 to 10
```

### Seeding for Reproducibility

```swahili
# Set seed for reproducible sequences
weka_mbegu(42)

namba val1 = nasibu(100)
namba val2 = nasibu(100)
andika("First run:", val1, val2)

# Reset seed to same value - get same sequence
weka_mbegu(42)
namba val3 = nasibu(100)
namba val4 = nasibu(100)
andika("Second run:", val3, val4)  # Same as first run!
```

### Practical Applications

**Dice Simulator:**
```swahili
kazi kufa_sita() {
    rudisha nasibu(1, 6)  # Roll 1-6
}

kazi kufa_kumi_ishirini() {
    rudisha nasibu(1, 20)  # Roll 1-20
}

namba roll1 = kufa_sita()
namba roll2 = kufa_sita()
andika("Rolling two six-sided dice:", roll1, "+", roll2, "=", roll1 + roll2)
```

**Shuffle Array:**
```swahili
kazi kweneza_orodha(orodha arr) {
    namba i = urefu_orodha(arr) - 1
    
    wakati i > 0 {
        namba j = nasibu(0, i)
        
        # Swap arr[i] and arr[j]
        namba temp = pata(arr, i)
        arr = ondoa(arr, i)
        ongeza(arr, temp)
        
        i = i - 1
    }
}
```

**Random Color Generator:**
```swahili
kazi rangi_nasibu() {
    namba r = nasibu(0, 256)
    namba g = nasibu(0, 256)
    namba b = nasibu(0, 256)
    
    rudisha "#" + maneno_namba(r) + maneno_namba(g) + maneno_namba(b)
}
```

**Sampling from List:**
```swahili
orodha fruits = ["apple", "banana", "orange", "mango"]

kazi sampuli(orodha list) {
    namba index = nasibu(urefu_orodha(list))
    rudisha pata(list, index)
}

andika("Random fruit:", sampuli(fruits))
```

### Error Handling

```swahili
# Error: nasibu with invalid range
jaribu {
    namba bad = nasibu(10, 10)  # min = max, invalid
} shika (error) {
    andika("Error: min must be less than max")
}

# Error: nasibu with non-positive number (1 arg)
jaribu {
    namba bad = nasibu(0)       # max must be positive
} shika (error) {
    andika("Error: max must be positive")
}
```

### Comparison: Random Function Modes

| Function | Usage | Range | Type |
|----------|-------|-------|------|
| nasibu() | Random float | [0.0, 1.0) | float |
| nasibu(n) | Random 0 to n | [0, n) | integer |
| nasibu(min, max) | Random in range | [min, max] | integer |
| weka_mbegu(seed) | Set seed | - | void |

---

**For more information on general language features, see [README.md](README.md)**
