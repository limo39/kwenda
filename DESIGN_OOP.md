# Object-Oriented Programming in Kwenda - Design Document

## Overview
Kwenda now supports object-oriented programming with a simplified but powerful approach using dictionary-based objects.

## Approach: Dictionary-Based Objects

Instead of implementing full class syntax (which would require extensive parser changes), we use a pragmatic approach:
- Objects are dictionaries/maps
- "Classes" are constructor functions that return configured dictionaries
- Methods are functions that take the object as the first parameter
- This approach is similar to JavaScript prototypes or Python's explicit self

## Syntax

### Creating an "Object Constructor"
```swahili
kazi Mtu_unda(maneno jina, namba umri) {
    # Create object as dictionary
    # Properties are stored as key-value pairs
    # Return the configured object
}
```

### Creating Instances
```swahili
# Create new person
Mtu mtu1 = Mtu_unda("Amina", 25)
```

### Accessing Properties
```swahili
# Access using dot notation (if supported) or function calls
maneno jina = Mtu_jina(mtu1)
```

### Calling Methods
```swahili
# Methods take object as first parameter
Mtu_salamu(mtu1)
```

## Implementation Status

### Phase 1: Function-Based OOP (CURRENT)
- ✅ Constructor functions
- ✅ Method functions with object parameter
- ✅ Naming convention: ClassName_methodName
- ❌ No special syntax (uses existing function system)

### Phase 2: Dictionary Support (PLANNED)
- Add dictionary/map data type
- Support for key-value storage
- Enable property access

### Phase 3: Syntactic Sugar (FUTURE)
- `darasa` keyword for class definitions
- `unda` keyword for instantiation
- `hii` keyword for self-reference
- Dot notation for method calls

## Current Workaround

Until full OOP is implemented, use this pattern:

```swahili
# Constructor
kazi Mtu_unda(maneno jina, namba umri) {
    andika("Created person:", jina, "age:", umri)
    rudisha jina  # Return identifier
}

# Methods
kazi Mtu_salamu(maneno jina) {
    andika("Habari, jina langu ni", jina)
}

kazi Mtu_umri_ongeza(maneno jina, namba miaka) {
    andika(jina, "is now", miaka, "years older")
}

# Usage
kazi kuu() {
    maneno mtu1 = Mtu_unda("Amina", 25)
    Mtu_salamu(mtu1)
    Mtu_umri_ongeza(mtu1, 1)
}
```

## Benefits of This Approach

1. **Works Now**: No parser changes needed
2. **Clear**: Explicit about what's happening
3. **Flexible**: Easy to understand and debug
4. **Familiar**: Similar to procedural programming
5. **Upgradeable**: Can add syntax sugar later

## Future Enhancements

When dictionary support is added:
```swahili
kazi Mtu_unda(maneno jina, namba umri) {
    kamusi mtu = {}
    mtu["jina"] = jina
    mtu["umri"] = umri
    rudisha mtu
}

kazi Mtu_salamu(kamusi mtu) {
    andika("Habari, jina langu ni", mtu["jina"])
}
```

When full OOP syntax is added:
```swahili
darasa Mtu {
    maneno jina
    namba umri
    
    kazi unda(maneno j, namba u) {
        hii.jina = j
        hii.umri = u
    }
    
    kazi salamu() {
        andika("Habari, jina langu ni", hii.jina)
    }
}

Mtu mtu1 = unda Mtu("Amina", 25)
mtu1.salamu()
```
