# Mathematical Constants in Kwenda

Kwenda provides a set of pre-defined mathematical constants in the math module for common calculations. All constants are defined with high precision (11+ decimal places).

## Available Constants

To use any constant, import the math module:

```swahili
leta "modules/math.swh"
```

Then access constants using `math.CONSTANT_NAME`.

### Primary Constants

#### PI (π)
```swahili
namba value = math.PI  # 3.14159265359
```

**Description:** Pi (π) is the ratio of a circle's circumference to its diameter. One of the most fundamental constants in mathematics.

**Uses:**
- Circle and sphere calculations
- Trigonometric functions
- Angular measurements
- Wave physics

**Example:**
```swahili
namba radius = 5
namba circumference = 2 * math.PI * radius    # 31.4159...
namba area = math.PI * radius * radius         # 78.5398...
```

#### E (Euler's Number)
```swahili
namba value = math.E  # 2.71828182846
```

**Description:** E is Euler's number, the base of natural logarithms and the foundation of exponential growth/decay.

**Uses:**
- Exponential and logarithmic functions
- Compound interest calculations
- Radioactive decay
- Growth models
- Natural phenomena

**Example:**
```swahili
namba compound_amount = principal * exp(rate * time)  # Uses e internally
```

### Related PI Constants

#### PI_2 (π/2)
```swahili
namba quarter_rotation = math.PI_2  # 1.57079632679
```

**Description:** Quarter rotation (90 degrees in radians).

**Uses:**
- Right angle calculations
- Trigonometric identities
- Quadrant boundaries

#### PI_4 (π/4)
```swahili
namba eighth_rotation = math.PI_4  # 0.78539816339
```

**Description:** Eighth rotation (45 degrees in radians).

**Uses:**
- 45-degree angles
- Isosceles right triangles
- Common angle in calculations

#### TAU (2π)
```swahili
namba full_rotation = math.TAU  # 6.28318530718
```

**Description:** Tau represents a full rotation (360 degrees in radians). Some mathematicians prefer tau over pi for rotational calculations.

**Uses:**
- Full circular rotations
- Frequency calculations
- Wave period calculations

**Example:**
```swahili
namba revolutions = 3
namba radians = revolutions * math.TAU  # 3 full rotations
```

### Angle Conversion Constants

#### DEG_TO_RAD
```swahili
namba radians = degrees * math.DEG_TO_RAD
```

**Value:** 0.01745329249 (= π/180)

**Example:**
```swahili
namba angle_deg = 90
namba angle_rad = angle_deg * math.DEG_TO_RAD  # 1.5708... radians
```

#### RAD_TO_DEG
```swahili
namba degrees = radians * math.RAD_TO_DEG
```

**Value:** 57.2957795131 (= 180/π)

**Example:**
```swahili
namba angle_rad = math.PI / 4
namba angle_deg = angle_rad * math.RAD_TO_DEG  # 45 degrees
```

### Root Constants

#### SQRT_2 (√2)
```swahili
namba value = math.SQRT_2  # 1.41421356237
```

**Description:** Square root of 2.

**Uses:**
- Diagonal of unit square
- Hypotenuse calculations
- Signal processing

**Example:**
```swahili
namba hypotenuse = 1 * math.SQRT_2  # For unit square
```

#### SQRT_3 (√3)
```swahili
namba value = math.SQRT_3  # 1.73205080757
```

**Description:** Square root of 3.

**Uses:**
- Equilateral triangle calculations
- 60-degree angle calculations

#### SQRT_5 (√5)
```swahili
namba value = math.SQRT_5  # 2.23606797750
```

**Description:** Square root of 5.

**Uses:**
- Golden ratio related calculations
- Fibonacci sequences
- Geometric constructions

### Golden Ratio

#### PHI (φ)
```swahili
namba ratio = math.PHI  # 1.61803398875
```

**Description:** The golden ratio, often found in nature and aesthetics.

**Uses:**
- Art and architecture proportions
- Fibonacci sequences
- Natural pattern analysis
- Geometric design

**Example:**
```swahili
namba width = 100
namba height = width / math.PHI  # Golden rectangle proportions
```

### Logarithmic Constants

#### LN_2 (ln(2))
```swahili
namba value = math.LN_2  # 0.69314718056
```

**Description:** Natural logarithm of 2.

**Uses:**
- Half-life calculations
- Binary logarithm conversions
- Information theory

#### LN_10 (ln(10))
```swahili
namba value = math.LN_10  # 2.30258509299
```

**Description:** Natural logarithm of 10.

**Uses:**
- Logarithm base conversions
- pH calculations
- Decibel calculations

### Advanced Constants

#### EULER_GAMMA (γ)
```swahili
namba value = math.EULER_GAMMA  # 0.57721566490
```

**Description:** The Euler-Mascheroni constant, appears in number theory and analysis.

**Uses:**
- Advanced mathematical analysis
- Series calculations
- Number theory

## Common Usage Patterns

### Circle Calculations
```swahili
leta "modules/math.swh"

kazi kuu() {
    namba radius = 10
    
    # Circumference
    namba circumference = 2 * math.PI * radius
    
    # Area
    namba area = math.PI * radius * radius
    
    # Semicircle area
    namba semicircle = (math.PI * radius * radius) / 2
    
    andika("Circumference:", circumference)
    andika("Area:", area)
    andika("Semicircle:", semicircle)
}
```

### Sphere Calculations
```swahili
namba radius = 5

# Surface area
namba surface = 4 * math.PI * radius * radius

# Volume
namba volume = (4.0 / 3.0) * math.PI * radius * radius * radius
```

### Exponential Growth
```swahili
namba initial = 100
namba growth_rate = 0.1
namba time = 5

# Population growth: N(t) = N0 * e^(rt)
namba population = initial * exp(growth_rate * time)
```

### Exponential Decay
```swahili
namba initial = 100
namba decay_rate = 0.05
namba time = 10

# Radioactive decay: N(t) = N0 * e^(-rt)
namba remaining = initial * exp(0 - decay_rate * time)
```

### Angle Conversion
```swahili
# Degrees to radians
namba angle_deg = 45
namba angle_rad = angle_deg * math.DEG_TO_RAD

# Radians to degrees
namba angle_rad2 = math.PI / 3
namba angle_deg2 = angle_rad2 * math.RAD_TO_DEG
```

### Trigonometric Calculations
```swahili
# 30-60-90 triangle
namba angle_30 = math.PI / 6
namba angle_60 = math.PI / 3
namba angle_90 = math.PI / 2

andika("sin(30°) =", sin(angle_30))
andika("sin(60°) =", sin(angle_60))
andika("sin(90°) =", sin(angle_90))
```

### Half-Life Calculation
```swahili
namba half_life = 5730  # Carbon-14 half-life in years
namba time = 11460     # 2 half-lives
namba initial = 100

# Decay formula: N(t) = N0 * 0.5^(t / half_life)
# Or: N(t) = N0 * e^(-t * ln(2) / half_life)
namba remaining = initial * exp(0 - math.LN_2 * time / half_life)
```

### Distance and Pythagorean Theorem
```swahili
namba a = 3
namba b = 4

# Using Pythagorean theorem
namba c = mzizi_mraba(a * a + b * b)  # 5
namba diagonal = mzizi_mraba(2) * side  # Diagonal of square using SQRT_2
```

### Frequency and Period
```swahili
namba frequency = 50  # Hz
namba period = math.TAU / frequency  # Time for full cycle
```

## Precision and Accuracy

All constants in Kwenda are stored with 11 decimal places of precision:

- **PI:** 3.14159265359
- **E:** 2.71828182846
- **SQRT_2:** 1.41421356237
- **PHI:** 1.61803398875

This precision is sufficient for most engineering and scientific calculations.

## Mathematical Relationships

Some constants are related:

- `PI_2 = PI / 2`
- `PI_4 = PI / 4`
- `TAU = 2 * PI`
- `DEG_TO_RAD = PI / 180`
- `RAD_TO_DEG = 180 / PI`
- `LN_2 = log_asili(2)`
- `LN_10 = log_asili(10)`
- `PHI ≈ (1 + SQRT_5) / 2` (golden ratio formula)

## Performance

Accessing constants is extremely fast—they are pre-computed values, not functions. There is no computational overhead.

## See Also

- [MATH.md](MATH.md) - Mathematical functions (sin, cos, exp, log, etc.)
- [RANDOM_NUMBERS.md](RANDOM_NUMBERS.md) - Random number generation
- [README.md](README.md) - Main language documentation
