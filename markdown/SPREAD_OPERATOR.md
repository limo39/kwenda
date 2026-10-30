# Spread Operator (`...`)

The spread operator is a powerful feature in Kwenda that allows you to expand (or "spread") the elements of arrays and dictionaries inline. It uses the `...` syntax and provides a clean way to combine, copy, and manipulate collections.

## Overview

The spread operator (`...`) unpacks elements from arrays or key-value pairs from dictionaries into a new collection. It's particularly useful for:
- Combining multiple arrays or dictionaries
- Creating shallow copies
- Adding elements around existing collections
- Merging configurations with override capabilities

## Array Spreading

### Basic Syntax

```swahili
orodha type_name variable = [...array_expression]
```

### Examples

#### 1. Combining Arrays

```swahili
orodha namba arr1 = [1, 2, 3]
orodha namba arr2 = [4, 5, 6]
orodha namba combined = [...arr1, ...arr2]
andika("Combined:", combined)  # [1, 2, 3, 4, 5, 6]
```

#### 2. Array with Additional Elements

```swahili
orodha namba middle = [2, 3, 4]
orodha namba extended = [1, ...middle, 5, 6]
andika("Extended:", extended)  # [1, 2, 3, 4, 5, 6]
```

#### 3. Shallow Copy

```swahili
orodha namba original = [10, 20, 30]
orodha namba copy = [...original]
andika("Copy:", copy)  # [10, 20, 30]
```

#### 4. Multiple Spreads

```swahili
orodha namba a = [1]
orodha namba b = [2, 3]
orodha namba c = [4]
orodha namba sequence = [...a, ...b, ...c]
andika("Sequence:", sequence)  # [1, 2, 3, 4]
```

#### 5. Empty Array Spread

```swahili
orodha namba empty = []
orodha namba withEmpty = [1, ...empty, 2]
andika("Result:", withEmpty)  # [1, 2]
```

#### 6. Spread at Different Positions

```swahili
orodha namba nums = [5, 6, 7]

# Spread at beginning
orodha namba atStart = [...nums, 8, 9]
andika("At start:", atStart)  # [5, 6, 7, 8, 9]

# Spread at end
orodha namba atEnd = [1, 2, ...nums]
andika("At end:", atEnd)  # [1, 2, 5, 6, 7]

# Spread in middle
orodha namba inMiddle = [1, ...nums, 10]
andika("In middle:", inMiddle)  # [1, 5, 6, 7, 10]
```

## Dictionary Spreading

### Basic Syntax

```swahili
kamusi variable = {...dictionary_expression}
```

### Examples

#### 1. Merging Dictionaries

```swahili
kamusi person = {"name": "Amina", "age": 25}
kamusi location = {"city": "Dar es Salaam", "country": "Tanzania"}
kamusi profile = {...person, ...location}
andika("Profile:", profile)
# {"name": "Amina", "age": 25, "city": "Dar es Salaam", "country": "Tanzania"}
```

#### 2. Property Override

When the same key exists in multiple dictionaries, the last value wins:

```swahili
kamusi defaults = {"theme": "light", "fontSize": 12, "language": "en"}
kamusi userPrefs = {"fontSize": 16, "language": "sw"}
kamusi settings = {...defaults, ...userPrefs}
andika("Settings:", settings)
# {"theme": "light", "fontSize": 16, "language": "sw"}
```

#### 3. Adding Properties

```swahili
kamusi baseConfig = {"debug": uwongo}
kamusi extendedConfig = {...baseConfig, "port": 3000, "host": "localhost"}
andika("Config:", extendedConfig)
# {"debug": uwongo, "port": 3000, "host": "localhost"}
```

#### 4. Dictionary Cloning

```swahili
kamusi source = {"id": 123, "status": "active"}
kamusi cloned = {...source}
andika("Cloned:", cloned)
# {"id": 123, "status": "active"}
```

#### 5. Configuration Merging

```swahili
kamusi prodDefaults = {"env": "production", "cache": kweli, "timeout": 5000}
kamusi customConfig = {"timeout": 10000, "maxRetries": 3}
kamusi finalConfig = {...prodDefaults, ...customConfig}
andika("Final:", finalConfig)
# {"env": "production", "cache": kweli, "timeout": 10000, "maxRetries": 3}
```

## Use Cases

### 1. Combining Data from Multiple Sources

```swahili
# Merging user data from different sources
kamusi basicInfo = {"name": "Fatuma", "email": "fatuma@example.com"}
kamusi preferences = {"theme": "dark", "language": "sw"}
kamusi metadata = {"lastLogin": "2024-01-15", "role": "admin"}

kamusi completeUser = {...basicInfo, ...preferences, ...metadata}
andika("User:", completeUser)
```

### 2. Creating Variations

```swahili
# Base configuration with variations
orodha namba baseFeatures = [1, 2, 3]
orodha namba premiumFeatures = [...baseFeatures, 4, 5, 6]
orodha namba enterpriseFeatures = [...premiumFeatures, 7, 8, 9]

andika("Base:", baseFeatures)       # [1, 2, 3]
andika("Premium:", premiumFeatures) # [1, 2, 3, 4, 5, 6]
andika("Enterprise:", enterpriseFeatures) # [1, 2, 3, 4, 5, 6, 7, 8, 9]
```

### 3. Building Function Arguments

```swahili
orodha namba baseArgs = [100, 200]
orodha namba extraArgs = [300, 400]
orodha namba allArgs = [...baseArgs, ...extraArgs]
# Use allArgs to pass to a function
```

### 4. Configuration Override Pattern

```swahili
# Common pattern: default config + user config
kamusi defaultSettings = {
    "notifications": kweli,
    "sound": kweli,
    "vibration": uwongo
}

kamusi userSettings = {
    "sound": uwongo,
    "theme": "blue"
}

kamusi finalSettings = {...defaultSettings, ...userSettings}
# Result: defaults with user overrides applied
```

## Important Notes

### Shallow Copy Behavior

The spread operator creates a shallow copy. For nested structures, only the top level is copied:

```swahili
orodha namba nested = [[1, 2], [3, 4]]
orodha namba copy = [...nested]
# copy contains references to the same inner arrays
```

### Performance Considerations

- Spreading large arrays or dictionaries creates new collections
- For very large datasets, consider alternative approaches
- Multiple spreads are processed left to right

### Current Limitations

1. **Dictionary Declaration Parsing**: When multiple dictionary declarations with spread operators appear in sequence in the same file, the parser may have difficulty determining statement boundaries. 

**Workaround**: Use functions to return dictionaries or split declarations across different files.

```swahili
# May have issues:
kamusi d1 = {"a": 1}
kamusi d2 = {"b": 2}
kamusi merged = {...d1, ...d2}

# Workaround - use function:
kazi getSetting() {
    kamusi d1 = {"a": 1}
    kamusi d2 = {"b": 2}
    rudisha {...d1, ...d2}
}
```

## Best Practices

1. **Use for Immutability**: Create new collections instead of modifying existing ones
2. **Configuration Merging**: Ideal for combining default and user settings
3. **Data Composition**: Build complex structures from simpler parts
4. **Avoid Deep Nesting**: Remember it's a shallow operation
5. **Clear Intent**: Use spread when the intent is to combine or copy

## Related Features

- **List Comprehensions**: `[x kwa x katika iterable]` for transformed arrays
- **Dictionary Comprehensions**: `{key: value kwa item katika iterable}`
- **Array Slicing**: `array[start:end:step]` for extracting subarrays
- **Const Variables**: `thabiti orodha namba arr = [...source]` for immutable copies

## Examples in Action

See the `examples/` directory for complete working examples:
- `examples/spread_arrays.swh` - 10 array spreading test cases
- `examples/spread_dicts.swh` - 10 dictionary spreading test cases  
- `examples/spread_demo.swh` - Comprehensive demonstration with 12 scenarios

## Summary

The spread operator (`...`) is a versatile feature that makes working with arrays and dictionaries more intuitive and expressive. It enables functional programming patterns, immutable data transformations, and clean configuration management in Kwenda programs.
