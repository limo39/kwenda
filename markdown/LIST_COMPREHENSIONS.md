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

## Nested List Comprehensions

Nested list comprehensions allow you to create multi-dimensional arrays by placing one list comprehension inside another.

### Syntax
```
[[inner_expression kwa inner_variable katika inner_iterable] kwa outer_variable katika outer_iterable]
```

### How Nested Comprehensions Work

The outer comprehension iterates over its iterable, and for each iteration, the inner comprehension creates a new array. The result is a 2D array (array of arrays).

### Examples

#### 1. Multiplication Table
```kwenda
orodha namba range = [1, 2, 3, 4, 5]
orodha namba table = [[x * y kwa y katika range] kwa x katika range]
// Result: [[1, 2, 3, 4, 5], [2, 4, 6, 8, 10], [3, 6, 9, 12, 15], ...]
```

#### 2. Coordinate Grid
```kwenda
orodha namba coords = [0, 1, 2]
// Encode (x,y) pairs as x*10 + y
orodha namba grid = [[x * 10 + y kwa y katika coords] kwa x katika coords]
// Result: [[0, 1, 2], [10, 11, 12], [20, 21, 22]]
```

#### 3. Addition Table
```kwenda
orodha namba nums = [1, 2, 3]
orodha namba sums = [[x + y kwa y katika nums] kwa x katika nums]
// Result: [[2, 3, 4], [3, 4, 5], [4, 5, 6]]
```

### Filters in Nested Comprehensions

You can apply filters to both the inner and outer comprehensions:

#### Inner Filter Only
```kwenda
orodha namba nums = [1, 2, 3, 4]
// Only include y values > 2
orodha namba filtered = [[x + y kwa y katika nums kama y > 2] kwa x katika nums]
// Result: [[4, 5], [5, 6], [6, 7], [7, 8]]
// Each row only has values where y=3,4
```

#### Outer Filter Only
```kwenda
orodha namba nums = [1, 2, 3, 4]
// Only include rows where x > 2
orodha namba filtered = [[x + y kwa y katika nums] kwa x katika nums kama x > 2]
// Result: [[4, 5, 6, 7], [5, 6, 7, 8]]
// Only rows for x=3,4
```

#### Both Inner and Outer Filters
```kwenda
orodha namba nums = [1, 2, 3, 4]
// Only rows where x > 2 AND columns where y > 1
orodha namba both = [[x + y kwa y katika nums kama y > 1] kwa x katika nums kama x > 2]
// Result: [[5, 6, 7], [6, 7, 8]]
// Rows for x=3,4 and columns for y=2,3,4
```

### Practical Use Cases for Nested Comprehensions

#### 1. Distance Matrix
Calculate distances between all point pairs:
```kwenda
orodha namba points = [0, 1, 2, 3]
orodha namba distances = [[x - y kwa y katika points] kwa x katika points]
// Creates a symmetric distance matrix
```

#### 2. Asymmetric Dimensions
Create matrices with different row and column sizes:
```kwenda
orodha namba rows = [1, 2, 3]
orodha namba cols = [10, 20, 30, 40]
orodha namba matrix = [[r + c kwa c katika cols] kwa r katika rows]
// Result: 3x4 matrix
```

#### 3. Mathematical Operations
Compute mathematical relationships:
```kwenda
orodha namba base = [2, 3, 4]
// For each x, compute x² * y for all y
orodha namba powers = [[x * x * y kwa y katika base] kwa x katika base]
// Result: [[8, 12, 16], [18, 27, 36], [32, 48, 64]]
```

### Comparison: Nested Loops vs Nested Comprehensions

**Traditional Nested Loops (verbose):**
```kwenda
orodha namba result = []
wakati i = 0, i < 3, i = i + 1 {
    orodha namba row = []
    wakati j = 0, j < 3, j = j + 1 {
        ongeza(row, i * j)
    }
    ongeza(result, row)
}
```

**Nested List Comprehension (concise):**
```kwenda
orodha namba nums = [0, 1, 2]
orodha namba result = [[i * j kwa j katika nums] kwa i katika nums]
```

### Benefits of Nested Comprehensions

1. **Concise**: One line instead of multiple nested loops
2. **Readable**: Clear intent - what, not how
3. **Less Error-Prone**: No manual index management
4. **Composable**: Easy to add filters at any level
5. **Functional Style**: Emphasizes data transformation

### Understanding Execution Order

For `[[expr kwa inner_var katika inner_iter] kwa outer_var katika outer_iter]`:

1. Outer loop iterates over `outer_iter`
2. For each outer value:
   - Inner loop iterates over `inner_iter`
   - Inner comprehension creates an array
3. All inner arrays are collected into the result

Example trace:
```kwenda
[[x + y kwa y katika [1, 2]] kwa x katika [10, 20]]

// Step 1: x = 10
//   y = 1 → 10 + 1 = 11
//   y = 2 → 10 + 2 = 12
//   Inner result: [11, 12]

// Step 2: x = 20
//   y = 1 → 20 + 1 = 21
//   y = 2 → 20 + 2 = 22
//   Inner result: [21, 22]

// Final result: [[11, 12], [21, 22]]
```

### Tips and Best Practices

1. **Start Simple**: Test inner comprehension alone first, then nest it
2. **Use Meaningful Names**: `row`/`col`, `x`/`y` help clarify structure
3. **Consider Readability**: Very deeply nested comprehensions may be hard to read
4. **Debug Incrementally**: Build outer array first, then add inner comprehension
5. **Filter Early**: Apply filters to reduce unnecessary computations

### Examples in Action

See the comprehensive demo file: `examples/nested_list_comprehensions_demo.swh`
