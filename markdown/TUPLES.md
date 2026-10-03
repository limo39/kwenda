# Tuple Data Structures in Kwenda

## Overview

Tuples are ordered, **immutable** collections of elements. Once created, tuple elements cannot be changed - any modification creates a new tuple. This document describes the tuple data structure implementation in Kwenda.

## Key Characteristics

- **Ordered**: Elements maintain their position
- **Immutable**: Cannot be modified after creation
- **Indexed**: Access elements by position (0-based)
- **Heterogeneous**: Can contain different types of elements
- **Fast**: Efficient for read-only operations

## Syntax

### Declaring a Tuple

```swahili
tuple namba coordinates = (10, 20, 30)
tuple maneno names = ("Ali", "Fatuma", "Hassan")
tuple namba rgb = (255, 128, 0)
```

### Empty Tuple

```swahili
tuple namba empty = ()
```

### Single Element Tuple

```swahili
tuple namba single = (42, )  // Note the comma
```

## Tuple Operations

### Basic Operations

#### Get Length
```swahili
namba size = urefu_tuple(myTuple)
```

#### Access Element by Index
```swahili
namba elem = pata_tuple(myTuple, 0)  // Get first element
```

#### Check Membership
```swahili
boolean exists = imo_tuple(myTuple, 5)  // Returns kweli or uwongo
```

### Tuple Manipulation (Creates New Tuples)

#### Concatenate Tuples
```swahili
tuple namba combined = unganisha_tuple(tuple1, tuple2)
```

#### Repeat Tuple
```swahili
tuple namba repeated = rudia_tuple(myTuple, 3)  // Repeat 3 times
```

#### Slice Tuple
```swahili
// From start index to end
tuple namba slice1 = kata_tuple(myTuple, 2)

// From start to end index
tuple namba slice2 = kata_tuple(myTuple, 1, 4)
```

#### Replace Element (Creates New Tuple)
```swahili
tuple namba modified = badilisha_tuple(myTuple, 1, 999)  // Replace index 1
```

### Searching Operations

#### Find First Index
```swahili
namba index = index_tuple(myTuple, 15)  // Returns -1 if not found
```

#### Count Occurrences
```swahili
namba count = hesabu_tuple(myTuple, 5)
```

### Conversion Operations

#### Array to Tuple
```swahili
orodha namba arr = [1, 2, 3]
tuple namba t = tuple_kwa_orodha(arr)
```

#### Tuple to Array
```swahili
orodha namba arr = orodha_kwa_tuple(myTuple)
```

#### Tuple to String
```swahili
// Without separator
maneno str1 = kiungo_tuple(myTuple)

// With separator
maneno str2 = kiungo_tuple(myTuple, ", ")
```

## Complete Function Reference

| Function | Purpose | Parameters | Returns |
|----------|---------|------------|---------|
| `urefu_tuple` | Get tuple length | (tuple) | namba |
| `pata_tuple` | Get element at index | (tuple, index) | element |
| `imo_tuple` | Check if element exists | (tuple, element) | boolean |
| `unganisha_tuple` | Concatenate tuples | (tuple1, tuple2) | tuple |
| `rudia_tuple` | Repeat tuple | (tuple, count) | tuple |
| `kata_tuple` | Slice tuple | (tuple, start [, end]) | tuple |
| `index_tuple` | Find first index | (tuple, element) | namba (-1 if not found) |
| `hesabu_tuple` | Count occurrences | (tuple, element) | namba |
| `tuple_kwa_orodha` | Convert array to tuple | (array) | tuple |
| `orodha_kwa_tuple` | Convert tuple to array | (tuple) | orodha |
| `kiungo_tuple` | Join to string | (tuple [, separator]) | maneno |
| `badilisha_tuple` | Replace element | (tuple, index, value) | tuple |

## Immutability

Tuples are **immutable**, meaning:

```swahili
tuple namba original = (1, 2, 3)

// This creates a NEW tuple, doesn't modify original
tuple namba modified = badilisha_tuple(original, 1, 999)

andika(original)   // (1, 2, 3) - unchanged
andika(modified)   // (1, 999, 3) - new tuple
```

## When to Use Tuples

### Use Tuples When:

1. **Data shouldn't change**: Coordinates, RGB colors, dates
2. **Fixed structure**: Database records, configuration settings
3. **Dictionary keys**: Immutable keys for kamusi
4. **Function returns**: Multiple return values
5. **Performance**: Read-only data that's accessed frequently

### Example Use Cases

#### Coordinates
```swahili
tuple namba point = (100, 200)
namba x = pata_tuple(point, 0)
namba y = pata_tuple(point, 1)
```

#### RGB Colors
```swahili
tuple namba red = (255, 0, 0)
tuple namba green = (0, 255, 0)
tuple namba blue = (0, 0, 255)
```

#### Date/Time
```swahili
tuple namba date = (2024, 12, 25)  // Year, Month, Day
tuple namba time = (14, 30, 0)     // Hour, Minute, Second
```

## Tuples vs Arrays vs Sets

| Feature | Tuple | Array (Orodha) | Set (Seti) |
|---------|-------|----------------|------------|
| Ordered | ✅ Yes | ✅ Yes | ❌ No |
| Mutable | ❌ No | ✅ Yes | ✅ Yes |
| Duplicates | ✅ Yes | ✅ Yes | ❌ No |
| Indexed Access | ✅ Yes | ✅ Yes | ❌ No |
| Use Case | Fixed data | Dynamic lists | Unique items |

## Examples

See `examples/tuples_demo.swh` for comprehensive examples.

## Implementation Notes

- Tuples are implemented as a special `TupleValue` type in Go
- All modification operations create new tuples
- Tuples can be converted to/from arrays when mutability is needed
- Tuple access is O(1) for indexing
- Memory efficient for read-only data

## Best Practices

1. **Use for fixed structures**: When data won't change, use tuples
2. **Convert when needed**: Use `orodha_kwa_tuple()` if you need to modify
3. **Coordinates and colors**: Perfect for (x, y) or (r, g, b) values
4. **Named access**: Extract to variables for clarity
   ```swahili
   tuple namba point = (10, 20)
   namba x = pata_tuple(point, 0)
   namba y = pata_tuple(point, 1)
   ```

5. **Return multiple values**: Use tuples for functions that return multiple results
