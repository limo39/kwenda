# Set Data Structures in Kwenda

## Overview

Sets are unordered collections of unique elements. This document describes the set data structure implementation in Kwenda.

## Syntax

### Declaring a Set

```swahili
seti namba mySet = {1, 2, 3, 4, 5}
seti maneno names = {"Ali", "Fatuma", "Hassan"}
```

### Empty Set

```swahili
seti namba emptySet = {}
```

## Set Operations

### Basic Operations

#### Add Element
```swahili
ongeza_seti(mySet, 10)  // Returns new size
```

#### Remove Element
```swahili
ondoa_seti(mySet, 2)  // Returns new size
```

#### Check Membership
```swahili
boolean exists = imo_seti(mySet, 5)  // Returns kweli or uwongo
```

#### Get Size
```swahili
namba size = ukubwa_seti(mySet)
```

### Mathematical Set Operations

#### Union (Muungano)
Returns a set containing all elements from both sets.
```swahili
seti namba result = muungano(set1, set2)
```

####Intersection (Makutano)
Returns a set containing only elements in both sets.
```swahili
seti namba result = makutano(set1, set2)
```

#### Difference (Tofauti)
Returns elements in first set but not in second.
```swahili
seti namba result = tofauti(set1, set2)
```

#### Symmetric Difference (Tofauti Simetrik)
Returns elements in either set but not in both.
```swahili
seti namba result = tofauti_simetrik(set1, set2)
```

### Set Comparison

#### Subset Check
```swahili
boolean isSub = ni_subeti(set1, set2)  // Is set1 a subset of set2?
```

#### Superset Check
```swahili
boolean isSuper = ni_supereti(set1, set2)  // Is set1 a superset of set2?
```

### Set Utilities

#### Clear Set
```swahili
tupu_seti(mySet)  // Removes all elements
```

#### Copy Set
```swahili
seti namba copy = nakili_seti(mySet)
```

#### Convert Array to Set
```swahili
orodha namba arr = [1, 2, 2, 3, 3, 3]
seti namba uniqueSet = seti_kwa_orodha(arr)  // {1, 2, 3}
```

#### Convert Set to Array
```swahili
orodha namba arr = orodha_kwa_seti(mySet)
```

## Complete Function Reference

| Function | Purpose | Parameters | Returns |
|----------|---------|------------|---------|
| `ongeza_seti` | Add element to set | (set, element) | New size |
| `ondoa_seti` | Remove element from set | (set, element) | New size |
| `imo_seti` | Check if element is in set | (set, element) | boolean |
| `ukubwa_seti` | Get set size | (set) | namba |
| `muungano` | Union of two sets | (set1, set2) | seti |
| `makutano` | Intersection of two sets | (set1, set2) | seti |
| `tofauti` | Difference of two sets | (set1, set2) | seti |
| `tofauti_simetrik` | Symmetric difference | (set1, set2) | seti |
| `ni_subeti` | Check if subset | (set1, set2) | boolean |
| `ni_supereti` | Check if superset | (set1, set2) | boolean |
| `tupu_seti` | Clear all elements | (set) | namba (0) |
| `nakili_seti` | Copy a set | (set) | seti |
| `seti_kwa_orodha` | Convert array to set | (array) | seti |
| `orodha_kwa_seti` | Convert set to array | (set) | orodha |

## Examples

See `examples/sets_demo.swh` for comprehensive examples.

## Implementation Notes

- Sets are implemented internally as Go maps with keys representing elements
- Elements are stored as string keys for consistency
- Sets automatically handle duplicates (adding an existing element has no effect)
- Set order is not guaranteed (sets are unordered by definition)
- Conversion functions help bridge between arrays and sets

## Use Cases

1. **Removing duplicates from arrays**
2. **Tracking unique visitors/users**
3. **Finding common elements between collections**
4. **Set-based filtering and validation**
5. **Graph algorithms (adjacency lists)**

