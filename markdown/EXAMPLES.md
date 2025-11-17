# Kwenda Programming Examples

This document contains comprehensive examples demonstrating all features of the Kwenda programming language.

## Table of Contents
- [Basic Examples](#basic-examples)
- [Control Flow](#control-flow)
- [Functions](#functions)
- [Arrays](#arrays)
- [Strings](#strings)
- [File I/O](#file-io)
- [Error Handling](#error-handling)
- [Dictionaries](#dictionaries)
- [Object-Oriented Programming](#object-oriented-programming)
- [Modules](#modules)

## Basic Examples

### Example 1: Simple Calculator
```swahili
kazi kuu() {
    namba x = ingiza("Ingiza namba ya kwanza:")
    namba y = ingiza("Ingiza namba ya pili:")
    
    namba jibu = x + y
    andika("Jibu ni:", jibu)
}
```

**Output:**
```
Ingiza namba ya kwanza: 10
Ingiza namba ya pili: 15
Jibu ni: 25
```

### Example 2: Variables and Types
```swahili
kazi kuu() {
    namba umri = 25
    maneno jina = "Amina"
    boolean ni_mwanafunzi = kweli
    orodha namba namba_za_kwanza = [1, 2, 3, 4, 5]
    kamusi person = {"jina": "Juma", "umri": 30}
    
    andika("Jina:", jina)
    andika("Umri:", umri)
    andika("Ni mwanafunzi:", ni_mwanafunzi)
    andika("Namba:", namba_za_kwanza)
    andika("Person:", person)
}
```

## Control Flow

### Example 3: Conditional Statements
```swahili
kazi kuu() {
    namba umri = ingiza("Ingiza umri wako:")
    
    kama umri >= 18 {
        andika("Wewe ni mtu mzima")
    } sivyo {
        andika("Wewe ni mtoto")
    }
    
    kama umri == 21 {
        andika("Hongera! Umefika umri wa miaka 21!")
    }
}
```

### Example 4: Nested Conditionals
```swahili
kazi kuu() {
    namba alama = ingiza("Ingiza alama yako:")
    
    kama alama >= 90 {
        andika("A - Bora sana!")
    } sivyo {
        kama alama >= 80 {
            andika("B - Nzuri")
        } sivyo {
            kama alama >= 70 {
                andika("C - Wastani")
            } sivyo {
                andika("D au F - Jitahidi zaidi")
            }
        }
    }
}
```

### Example 5: While Loop
```swahili
kazi kuu() {
    namba i = 1
    
    andika("Mfuatano wa namba kutoka 1 hadi 5:")
    
    wakati i <= 5 {
        andika("Namba:", i)
        i = i + 1
    }
    
    andika("Mwisho wa loop!")
}
```

### Example 6: For Loop
```swahili
kazi kuu() {
    andika("For loop kutoka 0 hadi 4:")
    
    kwa i = 0; i < 5; i = i + 1 {
        andika("Iteration:", i)
    }
    
    andika("For loop simple:")
    namba j = 3
    
    kwa j > 0 {
        andika("j ni:", j)
        j = j - 1
    }
}
```

### Example 7: Break and Continue
```swahili
kazi kuu() {
    andika("Break example:")
    namba i = 1
    wakati i <= 10 {
        kama i == 5 {
            andika("Breaking at i =", i)
            vunja
        }
        andika("i =", i)
        i = i + 1
    }
    
    andika("Continue example:")
    namba j = 0
    wakati j < 5 {
        j = j + 1
        kama j == 3 {
            andika("Skipping j =", j)
            endelea
        }
        andika("j =", j)
    }
}
```

## Functions

### Example 8: Functions with Parameters
```swahili
kazi jumla(namba a, namba b) namba {
    namba jibu = a + b
    rudisha jibu
}

kazi salamu(namba umri) {
    kama umri >= 18 {
        andika("Habari za asubuhi, mtu mzima!")
    } sivyo {
        andika("Habari za asubuhi, kijana!")
    }
}

kazi ni_kubwa(namba x, namba y) boolean {
    rudisha x > y
}

kazi kuu() {
    namba jibu = jumla(10, 5)
    andika("10 + 5 =", jibu)
    
    salamu(25)
    
    boolean kubwa = ni_kubwa(8, 3)
    andika("8 > 3:", kubwa)
}
```

## Arrays

### Example 9: Array Operations
```swahili
kazi kuu() {
    # Create array
    orodha namba namba_zangu = [10, 20, 30, 40, 50]
    
    # Get length
    namba urefu = urefu_orodha(namba_zangu)
    andika("Length:", urefu)
    
    # Access elements
    namba kwanza = pata(namba_zangu, 0)
    andika("First element:", kwanza)
    
    # Add element
    ongeza(namba_zangu, 60)
    andika("After adding 60:", namba_zangu)
    
    # Remove element
    ondoa(namba_zangu, 2)
    andika("After removing index 2:", namba_zangu)
}
```

## Strings

### Example 10: String Manipulation
```swahili
kazi kuu() {
    # Basic string operations
    maneno jina = "Amina"
    maneno mji = "Dar es Salaam"
    maneno ujumbe = "Habari " + jina + " kutoka " + mji
    
    andika("Basic concatenation:", ujumbe)
    
    namba urefu_ujumbe = urefu(ujumbe)
    andika("Length of message:", urefu_ujumbe)
    
    # String manipulation functions
    maneno kubwa = herufi_kubwa(ujumbe)
    maneno ndogo = herufi_ndogo(ujumbe)
    
    andika("Uppercase:", kubwa)
    andika("Lowercase:", ndogo)
    
    # Substring operations
    maneno sehemu = kata(ujumbe, 0, 6)
    andika("First 6 characters:", sehemu)
    
    # Find and replace
    namba mahali = tafuta(ujumbe, "Amina")
    andika("Position of 'Amina':", mahali)
    
    maneno mpya = badilisha(ujumbe, "Amina", "Fatuma")
    andika("After replacement:", mpya)
}
```

## File I/O

### Example 11: File Operations
```swahili
kazi kuu() {
    # Create and manage files
    maneno jina_faili = "data.txt"
    unda_faili(jina_faili)
    
    # Write data to file
    andika_faili(jina_faili, "NAMBA ZA MUHIMU:\n")
    andika_faili(jina_faili, "Namba 10\n", kweli)
    andika_faili(jina_faili, "Namba 20\n", kweli)
    
    # Read and display file content
    maneno maudhui = soma(jina_faili)
    andika("Maudhui ya faili:")
    andika(maudhui)
    
    # Check if file exists and clean up
    boolean ipo = faili_ipo(jina_faili)
    andika("Faili ipo:", ipo)
    ondoa_faili(jina_faili)
    andika("Faili imeondolewa")
}
```

## Error Handling

### Example 12: Try-Catch-Finally
```swahili
leta "modules/math.swh"

kazi kuu() {
    andika("=== Error Handling Demo ===")
    
    # Test 1: Caught division by zero
    andika("Test 1: Division by zero (caught)")
    jaribu {
        namba result = math.gawanya(100, 0)
        andika("Result:", result)
    } shika (error) {
        andika("✓ Error caught:", error)
    }
    
    # Test 2: Try-catch-finally
    andika("Test 2: Try-catch-finally")
    jaribu {
        andika("Inside try block")
        tupa "Test error"
    } shika (e) {
        andika("Inside catch:", e)
    } hatimaye {
        andika("Finally block - always executes")
    }
    
    # Test 3: File error handling
    andika("Test 3: File not found (caught)")
    jaribu {
        maneno content = soma("nonexistent.txt")
        andika("Content:", content)
    } shika (error) {
        andika("✓ Error caught:", error)
    }
    
    andika("Program continues after error handling")
}
```

## Dictionaries

### Example 13: Dictionary Operations
```swahili
kazi kuu() {
    # Create dictionary
    kamusi person = {"jina": "Amina", "umri": 25, "mji": "Dar es Salaam"}
    
    # Access values
    andika("Name:", person["jina"])
    andika("Age:", person["umri"])
    
    # Modify values
    person["umri"] = 26
    
    # Add new keys
    person["kazi"] = "Mwalimu"
    
    # Print dictionary
    andika("Person:", person)
}
```

## Object-Oriented Programming

### Example 14: Class Syntax
```swahili
# Define a class with darasa keyword
darasa Mtu {
    maneno jina
    namba umri
    maneno mji
    
    # Constructor method
    kazi unda(maneno j, namba u, maneno m) {
        hii.jina = j
        hii.umri = u
        hii.mji = m
    }
    
    # Method
    kazi salamu() {
        andika("Habari! Jina langu ni", hii.jina)
        andika("Nina umri wa miaka", hii.umri)
    }
    
    kazi siku_ya_kuzaliwa() {
        hii.umri = hii.umri + 1
        andika("🎉 Happy Birthday! You are now", hii.umri)
    }
}

kazi kuu() {
    # Create instances using unda keyword
    kamusi mtu1 = unda Mtu("Amina", 25, "Dar es Salaam")
    kamusi mtu2 = unda Mtu("Juma", 30, "Arusha")
    
    # Access properties
    andika("Person 1:", mtu1["jina"], "-", mtu1["umri"], "years old")
    andika("Person 2:", mtu2["jina"], "-", mtu2["umri"], "years old")
}
```

### Example 15: Multiple Classes
```swahili
darasa Gari {
    maneno aina
    namba mwaka
    maneno rangi
    namba kilomita
    
    kazi unda(maneno a, namba m, maneno r) {
        hii.aina = a
        hii.mwaka = m
        hii.rangi = r
        hii.kilomita = 0
    }
    
    kazi endesha(namba umbali) {
        hii.kilomita = hii.kilomita + umbali
        andika("🚗", hii.aina, "drove", umbali, "km. Total:", hii.kilomita, "km")
    }
}

kazi kuu() {
    kamusi gari1 = unda Gari("Toyota", 2020, "Nyekundu")
    kamusi gari2 = unda Gari("Honda", 2022, "Bluu")
    
    andika("Car 1:", gari1["aina"], "(", gari1["mwaka"], ")")
    andika("Car 2:", gari2["aina"], "(", gari2["mwaka"], ")")
}
```

## Modules

### Example 16: Multi-File Programs
```swahili
# File: modules/math.swh
kazi ongeza(namba a, namba b) {
    rudisha a + b
}

kazi zidisha(namba a, namba b) {
    rudisha a * b
}

namba PI = 3

# File: main.swh
leta "modules/math.swh"

kazi kuu() {
    andika("=== Multi-file Demo ===")
    
    # Use math module functions
    namba a = 10
    namba b = 5
    namba jumla = math.ongeza(a, b)
    namba bidhaa = math.zidisha(a, b)
    
    andika("Math operations:")
    andika("  ", a, "+", b, "=", jumla)
    andika("  ", a, "*", b, "=", bidhaa)
    andika("  PI =", math.PI)
}
```

## Running Examples

All examples are available in the `examples/` directory. Run them with:

```bash
go run main.go examples/example_name.swh
```

For more detailed documentation, see:
- [README.md](README.md) - Main documentation
- [OOP.md](OOP.md) - Object-oriented programming guide
- [DICTIONARY_SUMMARY.md](DICTIONARY_SUMMARY.md) - Dictionary implementation
- [CLASS_SYNTAX_SUMMARY.md](CLASS_SYNTAX_SUMMARY.md) - Class syntax details
