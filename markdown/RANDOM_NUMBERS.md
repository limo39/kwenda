# Random Number Generation in Kwenda

Kwenda provides built-in random number generation functions for creating random values with optional seeding for reproducible sequences.

## Overview

Random number generation is useful for:
- Games and simulations
- Randomized algorithms
- Testing and demos
- Probabilistic applications
- Sampling from data

## Functions

### nasibu() - Random Float

Generate a random floating-point number in the range [0.0, 1.0).

```swahili
namba rand = nasibu()  # e.g., 0.372814, 0.823401, 0.156739
```

**Returns:** Float between 0.0 (inclusive) and 1.0 (exclusive)

### nasibu(n) - Random Integer [0, n)

Generate a random integer from 0 up to (but not including) n.

```swahili
namba die = nasibu(6)           # 0-5
namba card = nasibu(52)         # 0-51
namba percent = nasibu(101)     # 0-100
```

**Parameters:** 
- `n` - positive integer (maximum, exclusive)

**Returns:** Integer in range [0, n)

**Errors:** If n ≤ 0, throws error

### nasibu(min, max) - Random Integer [min, max]

Generate a random integer within an inclusive range.

```swahili
namba die_roll = nasibu(1, 6)           # 1-6 (dice)
namba coin_flip = nasibu(0, 1)          # 0 or 1 (heads/tails)
namba temp = nasibu(-50, 50)            # -50 to 50 (temperature)
namba year = nasibu(1900, 2024)         # 1900-2024
```

**Parameters:**
- `min` - minimum value (inclusive)
- `max` - maximum value (inclusive)

**Returns:** Integer in range [min, max]

**Errors:** If min ≥ max, throws error

### weka_mbegu(seed) - Set Random Seed

Set the random seed for reproducible sequences. Call this before generating random numbers to get the same sequence every time.

```swahili
weka_mbegu(42)

namba val1 = nasibu(100)  # e.g., 75
namba val2 = nasibu(100)  # e.g., 11
namba val3 = nasibu(100)  # e.g., 60

# Reset seed to get same sequence
weka_mbegu(42)
namba val1_again = nasibu(100)  # 75 (same as before)
namba val2_again = nasibu(100)  # 11 (same as before)
namba val3_again = nasibu(100)  # 60 (same as before)
```

**Parameters:**
- `seed` - integer seed value

**Returns:** nil

**Note:** Different seed values produce different sequences

## Examples

### Example 1: Simple Random Floats

```swahili
kazi kuu() {
    andika("Random decimals:")
    andika(nasibu())
    andika(nasibu())
    andika(nasibu())
}
```

Output:
```
Random decimals:
0.372814
0.823401
0.156739
```

### Example 2: Dice Rolls

```swahili
kazi kufa_sita() {
    rudisha nasibu(1, 6)
}

kazi kuu() {
    namba die1 = kufa_sita()
    namba die2 = kufa_sita()
    andika("Dice roll:", die1, "+", die2, "=", die1 + die2)
}
```

Output:
```
Dice roll: 4 + 5 = 9
```

### Example 3: Coin Flip Simulation

```swahili
kazi kuu() {
    namba heads = 0
    namba tails = 0
    
    namba i = 0
    wakati i < 10 {
        namba flip = nasibu(0, 1)
        kama flip == 0 {
            andika("Heads")
            heads = heads + 1
        } sivyo {
            andika("Tails")
            tails = tails + 1
        }
        i = i + 1
    }
    
    andika("Results: Heads =", heads, ", Tails =", tails)
}
```

### Example 4: Random Selection from Array

```swahili
kazi kuu() {
    orodha maneno colors = ["red", "blue", "green", "yellow"]
    
    namba i = 0
    wakati i < 5 {
        namba index = nasibu(0, 3)  # 0-3 for 4 colors
        andika("Random color index:", index)
        i = i + 1
    }
}
```

### Example 5: Reproducible Sequences

```swahili
kazi kuu() {
    andika("Test 1:")
    weka_mbegu(999)
    andika(nasibu(100))
    andika(nasibu(100))
    andika(nasibu(100))
    
    andika("")
    andika("Test 2 (same seed):")
    weka_mbegu(999)
    andika(nasibu(100))
    andika(nasibu(100))
    andika(nasibu(100))
}
```

Output:
```
Test 1:
45
82
31

Test 2 (same seed):
45
82
31
```

### Example 6: Random Walk

```swahili
kazi kuu() {
    namba x = 0
    namba y = 0
    
    andika("Starting at (0, 0)")
    
    namba step = 0
    wakati step < 5 {
        namba dx = nasibu(-1, 1)
        namba dy = nasibu(-1, 1)
        x = x + dx
        y = y + dy
        andika("Step", step + 1, "-> (", x, ",", y, ")")
        step = step + 1
    }
}
```

### Example 7: Lottery Number Generator

```swahili
kazi kuu() {
    andika("Your lottery numbers (pick 6 from 1-49):")
    
    namba i = 0
    wakati i < 6 {
        namba number = nasibu(1, 49)
        andika("  Number", i + 1, ":", number)
        i = i + 1
    }
}
```

### Example 8: Weighted Probability

```swahili
kazi kuu() {
    namba common = 0
    namba rare = 0
    namba legendary = 0
    
    namba i = 0
    wakati i < 100 {
        namba roll = nasibu(0, 99)
        
        # 70% common, 25% rare, 5% legendary
        kama roll < 70 {
            common = common + 1
        } sivyo_kama roll < 95 {
            rare = rare + 1
        } sivyo {
            legendary = legendary + 1
        }
        
        i = i + 1
    }
    
    andika("Common:", common)
    andika("Rare:", rare)
    andika("Legendary:", legendary)
}
```

### Example 9: Password Generator

```swahili
kazi kuu() {
    andika("Generated passwords:")
    
    namba i = 0
    wakati i < 5 {
        namba code = nasibu(100000, 999999)
        andika("  Password", i + 1, ":", code)
        i = i + 1
    }
}
```

### Example 10: Shuffling with Random

```swahili
kazi kuu() {
    andika("Random indexes from 0-4 (simulate card draw without replacement):")
    
    namba i = 0
    wakati i < 5 {
        namba index = nasibu(0, 4)
        andika("  Index:", index)
        i = i + 1
    }
}
```

## Practical Applications

### Games
- Dice rolls for board games
- Random NPC behavior
- Loot drop tables
- Random map generation

### Simulations
- Monte Carlo simulations
- Physics simulations (particle systems)
- Crowd behavior
- Weather simulation

### Testing
- Random test data generation
- Stress testing with varied inputs
- Fuzz testing
- A/B testing randomization

### Statistics
- Sampling from populations
- Bootstrap resampling
- Permutation testing
- Cross-validation

### Data Science
- Train/test split randomization
- Shuffling datasets
- Stratified sampling
- Random forest initialization

## Common Patterns

### Generate Random Boolean

```swahili
boolean yes_or_no = nasibu(0, 1) == 0
```

### Generate Random Percentage

```swahili
namba percentage = nasibu(0, 100)  # 0-100
```

### Random Color (RGB)

```swahili
namba r = nasibu(0, 255)
namba g = nasibu(0, 255)
namba b = nasibu(0, 255)
andika("RGB(", r, ",", g, ",", b, ")")
```

### Random Selection from List

```swahili
orodha maneno items = ["apple", "banana", "cherry"]
namba index = nasibu(0, 2)  # For 3 items, use 0-2
andika("Selected:", items)  # Would need array indexing
```

### Range Distribution

```swahili
namba age = nasibu(18, 100)
namba score = nasibu(0, 100)
namba year = nasibu(1900, 2024)
```

## Reproducibility

For testing and demos, use `weka_mbegu()` to ensure consistent results:

```swahili
# Testing with fixed values
weka_mbegu(12345)
namba test_val1 = nasibu(100)
namba test_val2 = nasibu(100)

# Later, regenerate same values
weka_mbegu(12345)
namba test_val1_again = nasibu(100)  # Same as test_val1
namba test_val2_again = nasibu(100)  # Same as test_val2
```

## Performance

All random functions have O(1) time complexity. They are fast and efficient for high-volume generation.

## Notes

- `nasibu()` uses Go's `math/rand` package internally
- Random sequences can be made reproducible with `weka_mbegu()`
- For cryptographic randomness, these functions are not suitable (not cryptographically secure)
- Different platforms may have different random sequences with the same seed

## See Also

- [MATH.md](MATH.md) - Other mathematical functions
- [README.md](README.md) - Main language documentation
