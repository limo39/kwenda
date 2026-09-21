# Advanced Special Functions in Kwenda

Kwenda provides advanced mathematical special functions used in scientific computing, statistics, and engineering applications.

## Gamma Function

The **Gamma function** is a generalization of the factorial function to real and complex numbers.

For positive integers: `Γ(n) = (n-1)!`

### Syntax
```kwenda
gamma(x)
```

### Parameters
- `x` (number): The input value

### Returns
- A floating-point number representing Γ(x)

### Examples
```kwenda
andika(gamma(5))      // Output: 24 (which is 4! = 24)
andika(gamma(1))      // Output: 1
andika(gamma(0.5))    // Output: 1.7724538509055159 (≈ √π)
```

### Notes
- The Gamma function has poles at negative integers (0, -1, -2, ...)
- For non-negative integers n, Γ(n+1) = n!
- Used in probability distributions, combinatorics, and analysis

---

## Error Function (erf)

The **Error Function** is used in statistics and probability theory. It represents the probability integral of a normal distribution.

### Syntax
```kwenda
erf(x)
```

### Parameters
- `x` (number): The input value

### Returns
- A value between -1 and 1 representing the error function

### Examples
```kwenda
andika(erf(0))    // Output: 0
andika(erf(1))    // Output: 0.8427007929497149
andika(erf(2))    // Output: 0.9953222650189527
```

### Domain
- Input: All real numbers
- Output: [-1, 1]

### Notes
- erf(-x) = -erf(x) (odd function)
- erf(0) = 0
- erf(∞) = 1
- Used extensively in statistics for normal distribution calculations

---

## Complementary Error Function (erfc)

The **Complementary Error Function** is defined as: `erfc(x) = 1 - erf(x)`

### Syntax
```kwenda
erfc(x)
```

### Parameters
- `x` (number): The input value

### Returns
- A value between 0 and 2 representing 1 - erf(x)

### Examples
```kwenda
andika(erfc(0))    // Output: 1
andika(erfc(1))    // Output: 0.15729920705028513
andika(erfc(2))    // Output: 0.004677734981047265
```

### Notes
- More accurate for large x values where erf(x) is very close to 1
- erfc(x) = 1 - erf(x)
- Often used in statistical calculations and signal processing

---

## Bessel Functions of the First Kind

Bessel functions of the first kind (J_n) are solutions to Bessel's differential equation, used in oscillatory phenomena and wave equations.

### J₀ Function

#### Syntax
```kwenda
j0(x)
```

#### Parameters
- `x` (number): The input value

#### Returns
- The Bessel function J₀(x)

#### Examples
```kwenda
andika(j0(0))    // Output: 1
andika(j0(1))    // Output: 0.7651976865579666
andika(j0(2))    // Output: 0.2238907791411228
```

### J₁ Function

#### Syntax
```kwenda
j1(x)
```

#### Parameters
- `x` (number): The input value

#### Returns
- The Bessel function J₁(x)

#### Examples
```kwenda
andika(j1(0))    // Output: 0
andika(j1(1))    // Output: 0.4400505857449335
andika(j1(2))    // Output: 0.5767052694876283
```

### Notes
- Used in circular wave propagation, vibrating membranes, heat conduction in cylindrical coordinates
- J₀(0) = 1, J₁(0) = 0
- Oscillatory behavior with decreasing amplitude
- Domain: All real numbers

---

## Bessel Functions of the Second Kind

Bessel functions of the second kind (Y_n) are also solutions to Bessel's differential equation, linearly independent from J_n.

### Y₀ Function

#### Syntax
```kwenda
y0(x)
```

#### Parameters
- `x` (number): The input value (must be positive)

#### Returns
- The Bessel function Y₀(x)

#### Examples
```kwenda
andika(y0(1))    // Output: 0.08825696421567697
andika(y0(2))    // Output: 0.5103756726497451
```

### Y₁ Function

#### Syntax
```kwenda
y1(x)
```

#### Parameters
- `x` (number): The input value (must be positive)

#### Returns
- The Bessel function Y₁(x)

#### Examples
```kwenda
andika(y1(1))    // Output: -0.7812128213002887
andika(y1(2))    // Output: -0.10703243154093756
```

### Notes
- **Domain restriction**: Y₀ and Y₁ are only defined for positive x values
- Will throw an error if x ≤ 0
- Also called Neumann functions (N_n)
- Linearly independent from J_n, forming a complete solution space
- Used for boundary value problems with cylindrical symmetry

---

## Complete Test Example

Here's a comprehensive example testing all advanced special functions:

```kwenda
andika("=== Advanced Special Functions ===")
andika("\nGamma Function:")
andika("gamma(5) =", gamma(5))
andika("gamma(0.5) =", gamma(0.5))

andika("\nError Functions:")
andika("erf(1) =", erf(1))
andika("erfc(1) =", erfc(1))

andika("\nBessel J Functions:")
andika("j0(0) =", j0(0))
andika("j1(1) =", j1(1))

andika("\nBessel Y Functions:")
andika("y0(1) =", y0(1))
andika("y1(2) =", y1(2))
```

---

## Applications

### Gamma Function
- Combinatorics: Computing generalizations of factorials
- Probability: Beta distributions, Gamma distributions
- Complex analysis: Extending factorials to real and complex numbers

### Error Function
- Statistics: Normal distribution calculations
- Signal processing: Gaussian filter properties
- Physics: Heat diffusion, wave propagation

### Bessel Functions
- Wave equations: Circular membranes, cylindrical waveguides
- Heat conduction: Temperature distribution in cylinders
- Vibrations: Cylindrical resonators, seismic waves
- Signal processing: Filter design for circular/polar coordinates
- Quantum mechanics: Angular momentum eigenfunctions
- Optics: Diffraction patterns, Airy disk calculations

---

## Numerical Accuracy

All functions use Go's standard `math` package implementations:
- High precision: 64-bit floating-point (double precision)
- Suitable for scientific and engineering applications
- Results match established mathematical libraries

---

## Error Handling

- **Gamma function**: May return infinity at poles (negative integers)
- **Error functions**: Handle all real inputs correctly
- **Bessel J functions**: Handle all real inputs
- **Bessel Y functions**: Throw error for x ≤ 0 (undefined in that domain)

---

## Version Information

- Advanced special functions added in version 1.0+
- Based on Go's `math` package (math.Gamma, math.Erf, math.Erfc, math.J0, math.J1, math.Y0, math.Y1)
