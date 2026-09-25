# List Comprehensions in Kwenda

List comprehensions provide a concise way to create new arrays by applying an expression to each element of an existing array, optionally filtering the elements based on a condition.

## Syntax

### Basic Form (without filter)
```
[expression kwa variable katika iterable]
```

### With Filter
```
[expression kwa variable katika iterable kama condition]
```

## Keywords

- **kwa** - "for" (loop keyword)
- **katika** - "in" (membership keyword)
- **kama** - "if" (condition keyword)

## Examples

### 1. Identity Comprehension
Create a copy of an array:
```kwenda
orodha namba nums = [1, 2, 3, 4, 5]
orodha namba copy = [x kwa x katika nums]
// Result: [1, 2, 3, 4, 5]
```

### 2. Transform Each Element
Double each number:
```kwenda
orodha namba nums = [1, 2, 3, 4, 5]
orodha namba doubled = [x * 2 kwa x katika nums]
// Result: [2, 4, 6, 8, 10]
```

Square each number:
```kwenda
orodha namba nums = [1, 2, 3, 4, 5]
orodha namba squared = [x * x kwa x katika nums]
// Result: [1, 4, 9, 16, 25]
```

### 3. Filter Elements
Get only numbers greater than 2:
```kwenda
orodha namba nums = [1, 2, 3, 4, 5]
orodha namba filtered = [x kwa x katika nums kama x > 2]
// Result: [3, 4, 5]
```

Get numbers less than or equal to 3:
```kwenda
orodha namba nums = [1, 2, 3, 4, 5]
orodha namba small = [x kwa x katika nums kama x <= 3]
// Result: [1, 2, 3]
```

### 4. Filter and Transform
Double only numbers greater than 2:
```kwenda
orodha namba nums = [1, 2, 3, 4, 5]
orodha namba result = [x * 2 kwa x katika nums kama x > 2]
// Result: [6, 8, 10]
```

### 5. Complex Expressions
Apply polynomial expression (x² + x):
```kwenda
orodha namba nums = [1, 2, 3, 4, 5]
orodha namba result = [x * x + x kwa x katika nums]
// Result: [2, 6, 12, 20, 30]
```

### 6. Working with Ranges
```kwenda
orodha namba range = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
orodha namba result = [x * 3 kwa x katika range kama x > 5]
// Result: [18, 21, 24, 27, 30]
```

### 7. Negative Numbers
```kwenda
orodha namba nums = [1, 2, 3, 4, 5]
orodha namba negated = [x * -1 kwa x katika nums]
// Result: [-1, -2, -3, -4, -5]
```

### 8. Empty Results
If no elements match the filter condition, an empty array is returned:
```kwenda
orodha namba nums = [1, 2, 3, 4, 5]
orodha namba empty = [x kwa x katika nums kama x > 10]
// Result: []
```

## How It Works

1. **Iteration**: The comprehension iterates over each element in the iterable array
2. **Variable Binding**: Each element is bound to the specified variable name
3. **Condition Check** (optional): If a condition is provided with `kama`, only elements that satisfy the condition proceed
4. **Expression Evaluation**: The expression is evaluated for each element (or each filtered element)
5. **Result Collection**: All evaluated values are collected into a new array

## Use Cases

List comprehensions are particularly useful for:

1. **Data transformation** - Converting values in one format to another
2. **Filtering** - Selecting specific elements from an array
3. **Mathematical operations** - Applying calculations to each element
4. **Data cleaning** - Removing unwanted values while processing
5. **Creating derived data** - Generating new arrays based on existing ones

## Comparison with Loops

### Using a loop:
```kwenda
orodha namba nums = [1, 2, 3, 4, 5]
orodha namba doubled = []
wakati i = 0, i < urefu(nums), i = i + 1 {
    weka_katika(doubled, nums[i] * 2)
}
```

### Using list comprehension:
```kwenda
orodha namba nums = [1, 2, 3, 4, 5]
orodha namba doubled = [x * 2 kwa x katika nums]
```

The list comprehension is more concise, readable, and expressive.

## Technical Notes

- List comprehensions create a new array; they don't modify the original
- The loop variable (`x` in the examples) is scoped to the comprehension only
- Comprehensions can be nested (though this is not yet demonstrated)
- The iterable must be an array; other types will result in an empty array
- Expressions can reference the loop variable multiple times (e.g., `x * x + x`)

## Related Features

- **Arrays**: `orodha` keyword for array declarations
- **Loops**: `wakati` (while) and loop control (`vunja`, `endelea`)
- **Conditionals**: `kama` (if) for filtering conditions
- **Array Functions**: `urefu()`, `pata()`, `weka_katika()`, etc.
