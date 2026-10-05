# Array Slicing in Kwenda

## Overview

Array slicing allows you to extract sub-arrays using a concise Python-style syntax. This document describes slice operations in Kwenda.

## Syntax

```swahili
orodha[start:end:step]
```

Where:
- **start**: Starting index (inclusive)
- **end**: Ending index (exclusive)
- **step**: Increment between elements (default: 1)

All three parameters are optional!

## Basic Slicing

### Standard Slice `[start:end]`

```swahili
orodha namba arr = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

orodha namba slice1 = arr[2:5]   // [2, 3, 4]
orodha namba slice2 = arr[0:3]   // [0, 1, 2]
orodha namba slice3 = arr[5:9]   // [5, 6, 7, 8]
```

**Note**: The end index is **exclusive** (not included in result).

### Omitting Parameters

#### From Beginning `[:end]`
```swahili
orodha namba first_five = arr[:5]   // [0, 1, 2, 3, 4]
```

#### To End `[start:]`
```swahili
orodha namba from_five = arr[5:]    // [5, 6, 7, 8, 9]
```

#### Full Copy `[:]`
```swahili
orodha namba copy = arr[:]          // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
```

## Negative Indices

Negative numbers count from the end of the array:
- `-1` is the last element
- `-2` is second to last
- etc.

```swahili
orodha namba arr = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

orodha namba last_three = arr[-3:]      // [7, 8, 9]
orodha namba up_to_last = arr[:-2]      // [0, 1, 2, 3, 4, 5, 6, 7]
orodha namba middle = arr[-7:-3]        // [3, 4, 5, 6]
```

## Step Values

The step parameter controls the increment:

### Every Nth Element

```swahili
orodha namba arr = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

orodha namba every_second = arr[::2]    // [0, 2, 4, 6, 8]
orodha namba every_third = arr[::3]     // [0, 3, 6, 9]
orodha namba range_step = arr[1:8:2]    // [1, 3, 5, 7]
```

### Reverse with Negative Step

```swahili
// Reverse entire array
orodha namba reversed = arr[::-1]       // [9, 8, 7, 6, 5, 4, 3, 2, 1, 0]

// Reverse a portion
orodha namba rev_part = arr[7:2:-1]     // [7, 6, 5, 4, 3]

// Reverse every other element
orodha namba rev_step = arr[::-2]       // [9, 7, 5, 3, 1]
```

## Complete Syntax Reference

| Syntax | Description | Example Result |
|--------|-------------|----------------|
| `arr[2:5]` | Elements 2, 3, 4 | `[2, 3, 4]` |
| `arr[:5]` | First 5 elements | `[0, 1, 2, 3, 4]` |
| `arr[5:]` | From index 5 to end | `[5, 6, 7, 8, 9]` |
| `arr[:]` | Full copy | `[0, 1, 2, 3, ...]` |
| `arr[-3:]` | Last 3 elements | `[7, 8, 9]` |
| `arr[:-2]` | All except last 2 | `[0, 1, ..., 7]` |
| `arr[::2]` | Every 2nd element | `[0, 2, 4, 6, 8]` |
| `arr[1::2]` | Every 2nd from index 1 | `[1, 3, 5, 7, 9]` |
| `arr[::-1]` | Reverse | `[9, 8, 7, ..., 0]` |
| `arr[5:2:-1]` | Reverse from 5 to 3 | `[5, 4, 3]` |

## Practical Examples

### Get First/Last N Elements

```swahili
orodha namba data = [10, 20, 30, 40, 50, 60, 70, 80]

// First 3 elements
orodha namba first_three = data[:3]     // [10, 20, 30]

// Last 3 elements  
orodha namba last_three = data[-3:]     // [60, 70, 80]
```

### Split Array in Half

```swahili
namba urefu = urefu_orodha(data)
namba nusu = urefu / 2

orodha namba first_half = data[:nusu]
orodha namba second_half = data[nusu:]
```

### Remove First and Last Elements

```swahili
orodha namba trimmed = data[1:-1]       // [20, 30, 40, 50, 60, 70]
```

### Get Even/Odd Indexed Elements

```swahili
// Even indices (0, 2, 4, 6, ...)
orodha namba even_idx = data[::2]       // [10, 30, 50, 70]

// Odd indices (1, 3, 5, 7, ...)
orodha namba odd_idx = data[1::2]       // [20, 40, 60, 80]
```

### Reverse Array

```swahili
orodha namba reversed = data[::-1]      // [80, 70, 60, 50, 40, 30, 20, 10]
```

### Create Array Copy

```swahili
orodha namba original = [1, 2, 3, 4, 5]
orodha namba copy = original[:]

// Modifications to copy don't affect original
copy[0] = 999
andika(original)  // [1, 2, 3, 4, 5] - unchanged
andika(copy)      // [999, 2, 3, 4, 5]
```

## Edge Cases

### Out of Bounds Indices

Out of bounds indices are safely clamped:

```swahili
orodha namba arr = [1, 2, 3]

orodha namba safe = arr[1:10]           // [2, 3] - end clamped to length
orodha namba safe2 = arr[-10:]          // [1, 2, 3] - start clamped to 0
```

### Empty Slices

When start >= end with positive step:

```swahili
orodha namba empty = arr[2:1]           // [] - empty array
orodha namba empty2 = arr[5:3]          // [] - empty array
```

### Zero Step Error

Step cannot be zero:

```swahili
// This will throw an error:
// orodha namba invalid = arr[::0]
// Error: "Hatua haiwezi kuwa sifuri (Step cannot be zero)"
```

## Slicing with Variables

You can use variables for slice parameters:

```swahili
orodha namba arr = [100, 200, 300, 400, 500]

namba start = 1
namba end = 4
namba step = 1

orodha namba result = arr[start:end:step]   // [200, 300, 400]
```

## Performance Notes

- Slicing creates a **new array** (copy of elements)
- O(n) time complexity where n is the slice length
- Original array is never modified
- Memory is allocated for the new array

## Comparison with Other Data Structures

| Feature | Array Slice | Tuple Slice | Set |
|---------|-------------|-------------|-----|
| Syntax | `arr[:]` | N/A | N/A |
| Creates Copy | ✅ Yes | Use `kata_tuple()` | N/A |
| Mutable | ✅ Yes | ❌ No | ✅ Yes |
| Ordered | ✅ Yes | ✅ Yes | ❌ No |

## Best Practices

1. **Use `[:]` for copying**: Clean and idiomatic
   ```swahili
   orodha namba copy = original[:]
   ```

2. **Negative indices for end-relative**: More readable
   ```swahili
   orodha namba last_five = arr[-5:]
   ```

3. **Step for patterns**: Use step instead of loops
   ```swahili
   orodha namba evens = arr[::2]
   ```

4. **Variables for dynamic slicing**: When indices are computed
   ```swahili
   namba halfway = urefu_orodha(arr) / 2
   orodha namba first_half = arr[:halfway]
   ```

## Common Patterns

### Pagination
```swahili
namba page_size = 10
namba page_num = 2
namba start = page_num * page_size
namba end = start + page_size

orodha namba page_items = all_items[start:end]
```

### Moving Window
```swahili
namba window_size = 3
kwa namba i = 0; i < urefu_orodha(data) - window_size; i = i + 1 {
    orodha namba window = data[i:i+window_size]
    andika("Window:", window)
}
```

### Remove Duplicates from Ends
```swahili
// If first and last are duplicates of boundary
orodha namba cleaned = data[1:-1]
```

## See Also

- **arrays.swh** - Array examples
- **kata_tuple()** - Tuple slicing function
- **TUPLES.md** - Immutable collections

## Examples

See `examples/slice_demo.swh` for comprehensive examples.
