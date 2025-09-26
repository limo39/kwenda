# Masharti (Conditional Statements) katika Kwenda

Kwenda sasa inasaidia masharti kwa kutumia maneno ya Kiswahili:

## Muundo wa Msingi

```swahili
kama sharti {
    # Vitendo vya kufanya kama sharti ni kweli
} sivyo {
    # Vitendo vya kufanya kama sharti si kweli
}
```

## Vielelezo vya Ulinganishaji

- `==` - sawa na
- `!=` - si sawa na  
- `<` - ndogo kuliko
- `<=` - ndogo kuliko au sawa na
- `>` - kubwa kuliko
- `>=` - kubwa kuliko au sawa na

## Mifano

### Sharti la Msingi
```swahili
kazi kuu() {
    namba umri = ingiza("Ingiza umri wako:")
    
    kama umri >= 18 {
        andika("Wewe ni mtu mzima")
    } sivyo {
        andika("Wewe ni mtoto")
    }
}
```

### Masharti ya Ndani (Nested)
```swahili
kazi kuu() {
    namba alama = ingiza("Ingiza alama yako:")
    
    kama alama >= 90 {
        andika("A - Bora sana!")
    } sivyo {
        kama alama >= 80 {
            andika("B - Nzuri")
        } sivyo {
            andika("C au chini - Jitahidi zaidi")
        }
    }
}
```

### Masharti Mengi
```swahili
kazi kuu() {
    namba x = 10
    
    kama x > 5 {
        andika("Kubwa kuliko 5")
    }
    
    kama x == 10 {
        andika("Ni kumi hasa!")
    }
    
    kama x < 20 {
        andika("Ndogo kuliko 20")
    }
}
```

## Mifano ya Matumizi

Angalia mifano hii:
- `examples/conditional.swh` - Mfano wa msingi
- `examples/nested_if.swh` - Masharti ya ndani
- `examples/conditionals_demo.swh` - Mfano mkamilifu