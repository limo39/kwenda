# Object-Oriented Programming in Kwenda

## Overview

Kwenda supports object-oriented programming using a **function-based OOP pattern**. While we don't yet have `class` syntax, you can achieve OOP principles using constructor functions and method functions.

## Core Concepts

### 1. Constructor Functions
Constructor functions create and initialize "objects" (represented by identifiers or data structures).

**Naming Convention**: `ClassName_unda` or `ClassName_new`

```swahili
kazi Mtu_unda(maneno jina, namba umri) {
    andika("[Mtu] Created:", jina, "age:", umri)
    rudisha jina  # Return identifier
}
```

### 2. Method Functions
Methods are functions that operate on objects, taking the object as a parameter.

**Naming Convention**: `ClassName_methodName`

```swahili
kazi Mtu_salamu(maneno jina) {
    andika("Habari! Jina langu ni", jina)
}

kazi Mtu_siku_ya_kuzaliwa(maneno jina) {
    andika("🎉", jina, "ana sherehe ya siku ya kuzaliwa!")
}
```

### 3. Creating Instances

```swahili
# Create instances using constructor
maneno mtu1 = Mtu_unda("Amina", 25)
maneno mtu2 = Mtu_unda("Juma", 30)
```

### 4. Calling Methods

```swahili
# Call methods by passing the instance
Mtu_salamu(mtu1)  # Output: Habari! Jina langu ni Amina
Mtu_salamu(mtu2)  # Output: Habari! Jina langu ni Juma
```

## Complete Example: Person Class

```swahili
############################################
# "Class": Mtu (Person)
############################################

# Constructor
kazi Mtu_unda(maneno jina, namba umri, maneno mji) {
    andika("[Mtu] Created:", jina, "age:", umri, "from:", mji)
    rudisha jina
}

# Method: Greet
kazi Mtu_salamu(maneno jina) {
    andika("Habari! Jina langu ni", jina)
}

# Method: Introduce
kazi Mtu_jitambulishe(maneno jina, namba umri, maneno mji) {
    andika("Jina langu ni", jina)
    andika("Nina umri wa miaka", umri)
    andika("Ninatoka", mji)
}

# Method: Celebrate birthday
kazi Mtu_siku_ya_kuzaliwa(maneno jina) {
    andika("🎉", jina, "ana sherehe ya siku ya kuzaliwa!")
}

############################################
# Usage
############################################

kazi kuu() {
    # Create instances
    maneno mtu1 = Mtu_unda("Amina", 25, "Dar es Salaam")
    maneno mtu2 = Mtu_unda("Juma", 30, "Arusha")
    
    # Call methods
    Mtu_salamu(mtu1)
    Mtu_salamu(mtu2)
    
    Mtu_jitambulishe(mtu1, 25, "Dar es Salaam")
    Mtu_siku_ya_kuzaliwa(mtu2)
}
```

**Output:**
```
[Mtu] Created: Amina age: 25 from: Dar es Salaam
[Mtu] Created: Juma age: 30 from: Arusha
Habari! Jina langu ni Amina
Habari! Jina langu ni Juma
Jina langu ni Amina
Nina umri wa miaka 25
Ninatoka Dar es Salaam
🎉 Juma ana sherehe ya siku ya kuzaliwa!
```

## Example: Car Class

```swahili
############################################
# "Class": Gari (Car)
############################################

# Constructor
kazi Gari_unda(maneno aina, namba mwaka, maneno rangi) {
    andika("[Gari] Created:", aina, "year:", mwaka, "color:", rangi)
    rudisha aina
}

# Method: Start engine
kazi Gari_washa(maneno aina) {
    andika("🚗", aina, "- Injini imewashwa! Vroom vroom!")
}

# Method: Drive
kazi Gari_endesha(maneno aina, namba umbali) {
    andika("🚗", aina, "- Inaendesha kilomita", umbali)
}

# Method: Stop
kazi Gari_simama(maneno aina) {
    andika("🚗", aina, "- Imesimama. Injini imezimwa.")
}

############################################
# Usage
############################################

kazi kuu() {
    # Create car instances
    maneno gari1 = Gari_unda("Toyota", 2020, "Nyekundu")
    maneno gari2 = Gari_unda("Honda", 2022, "Bluu")
    
    # Use car 1
    Gari_washa(gari1)
    Gari_endesha(gari1, 50)
    Gari_simama(gari1)
    
    # Use car 2
    Gari_washa(gari2)
    Gari_endesha(gari2, 100)
    Gari_simama(gari2)
}
```

**Output:**
```
[Gari] Created: Toyota year: 2020 color: Nyekundu
[Gari] Created: Honda year: 2022 color: Bluu
🚗 Toyota - Injini imewashwa! Vroom vroom!
🚗 Toyota - Inaendesha kilomita 50
🚗 Toyota - Imesimama. Injini imezimwa.
🚗 Honda - Injini imewashwa! Vroom vroom!
🚗 Honda - Inaendesha kilomita 100
🚗 Honda - Imesimama. Injini imezimwa.
```

## OOP Principles in Kwenda

### 1. Encapsulation
Group related functions together using naming conventions:
- `ClassName_unda` - Constructor
- `ClassName_methodName` - Methods

### 2. Abstraction
Hide implementation details inside functions:
```swahili
kazi Gari_washa(maneno aina) {
    # Internal implementation hidden
    andika("🚗", aina, "- Injini imewashwa!")
}
```

### 3. Multiple Instances
Create multiple independent instances:
```swahili
maneno mtu1 = Mtu_unda("Amina", 25)
maneno mtu2 = Mtu_unda("Juma", 30)
maneno mtu3 = Mtu_unda("Fatuma", 28)
```

### 4. Method Chaining (Simulated)
While not automatic, you can design methods to enable sequential operations:
```swahili
Gari_washa(gari1)
Gari_endesha(gari1, 50)
Gari_simama(gari1)
```

## Best Practices

### 1. Consistent Naming
Always use the pattern: `ClassName_methodName`
```swahili
# Good
kazi Mtu_salamu(maneno jina) { ... }
kazi Gari_washa(maneno aina) { ... }

# Avoid
kazi salamu_mtu(maneno jina) { ... }
kazi washa(maneno aina) { ... }
```

### 2. Constructor Logging
Log object creation for debugging:
```swahili
kazi Mtu_unda(maneno jina, namba umri) {
    andika("[Mtu] Created:", jina, "age:", umri)
    rudisha jina
}
```

### 3. Parameter Validation
Validate inputs in constructors:
```swahili
kazi Mtu_unda(maneno jina, namba umri) {
    kama umri < 0 {
        andika("[Error] Age cannot be negative")
        rudisha ""
    }
    kama umri > 150 {
        andika("[Error] Age is unrealistic")
        rudisha ""
    }
    andika("[Mtu] Created:", jina, "age:", umri)
    rudisha jina
}
```

### 4. Method Documentation
Add comments to explain what methods do:
```swahili
# Method: Calculate and display person's birth year
kazi Mtu_mwaka_wa_kuzaliwa(maneno jina, namba umri) {
    namba mwaka_sasa = 2024
    namba mwaka_kuzaliwa = mwaka_sasa - umri
    andika(jina, "alizaliwa mwaka", mwaka_kuzaliwa)
}
```

## Limitations

### Current Limitations
1. **No True State**: Objects don't maintain internal state (use global variables as workaround)
2. **No Inheritance**: Cannot extend classes
3. **No Polymorphism**: No method overriding
4. **No Private Members**: All methods are public
5. **Manual Method Calls**: Must explicitly pass object to methods

### Workarounds

**State Management** (using global variables):
```swahili
# Global state for account
namba Akaunti_salio = 0

kazi Akaunti_unda(maneno jina, namba salio_awali) {
    Akaunti_salio = salio_awali
    andika("[Akaunti] Created for:", jina)
    rudisha jina
}

kazi Akaunti_weka(maneno jina, namba kiasi) {
    Akaunti_salio = Akaunti_salio + kiasi
    andika("Deposited:", kiasi, "New balance:", Akaunti_salio)
}
```

## Future Enhancements

### Planned Features
1. **Dictionary/Map Support**: Enable true object properties
2. **Class Syntax**: `darasa` keyword for class definitions
3. **Instance Creation**: `unda` keyword for instantiation
4. **Self Reference**: `hii` keyword for this/self
5. **Dot Notation**: `object.method()` syntax

### Future Syntax (Planned)
```swahili
darasa Mtu {
    maneno jina
    namba umri
    
    kazi unda(maneno j, namba u) {
        hii.jina = j
        hii.umri = u
    }
    
    kazi salamu() {
        andika("Habari! Jina langu ni", hii.jina)
    }
}

# Usage
Mtu mtu1 = unda Mtu("Amina", 25)
mtu1.salamu()
```

## Complete Working Example

See `examples/oop_demo.swh` for a comprehensive demonstration including:
- Person class with multiple methods
- Car class with state simulation
- Bank account class with balance management
- Multiple instances and method calls

Run it with:
```bash
go run main.go examples/oop_demo.swh
```

## Conclusion

While Kwenda doesn't yet have traditional class syntax, the function-based OOP pattern provides:
- ✅ Encapsulation through naming conventions
- ✅ Multiple instances
- ✅ Method organization
- ✅ Clear, understandable code
- ✅ Easy to learn and use

This approach is similar to procedural programming with OOP principles, making it accessible while we work toward full OOP syntax support.
