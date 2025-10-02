# Maneno (Strings) katika Kwenda

Kwenda inasaidia utumizi wa maneno (strings) kwa njia rahisi na ya kufurahisha. Hii ni mwongozo kamili wa jinsi ya kutumia maneno katika lugha ya Kwenda.

## Kutangaza Vigeuzi vya Maneno

```kwenda
maneno jina = "Amina"
maneno salamu = "Habari za asubuhi"
maneno tupu = ""
```

## Operesheni za Kimsingi

### 1. Kuunganisha Maneno (String Concatenation)

```kwenda
# Kutumia opereta +
maneno jina = "Amina"
maneno ujumbe = "Habari " + jina + "!"

# Kutumia function ya unganisha
maneno ujumbe2 = unganisha("Habari ", jina, "!")
```

### 2. Kulinganisha Maneno (String Comparison)

```kwenda
maneno jina1 = "Amina"
maneno jina2 = "Bakari"

boolean ni_sawa = jina1 == jina2        # uwongo
boolean si_sawa = jina1 != jina2        # kweli
```

## Functions za Maneno

### urefu(maneno) - Kupata Urefu wa Neno

Inarudisha idadi ya herufi katika neno.

```kwenda
maneno jina = "Amina"
namba urefu = urefu(jina)  # 5
```

### unganisha(maneno...) - Kuunganisha Maneno

Inaunganisha maneno mengi kuwa neno moja.

```kwenda
maneno ujumbe = unganisha("Habari ", "za ", "asubuhi")  # "Habari za asubuhi"
```

### kata(maneno, mahali_ya_kuanza, [urefu]) - Kukata Sehemu ya Neno

Inachukua sehemu ya neno kuanzia mahali fulani.

```kwenda
maneno neno = "Habari Dunia"
maneno sehemu1 = kata(neno, 0, 6)    # "Habari"
maneno sehemu2 = kata(neno, 7)       # "Dunia"
```

### badilisha(maneno, la_zamani, la_mpya) - Kubadilisha Sehemu za Neno

Inabadilisha sehemu zote za neno kwa sehemu mpya.

```kwenda
maneno sentensi = "Habari za asubuhi"
maneno mpya = badilisha(sentensi, "asubuhi", "jioni")  # "Habari za jioni"
```

### tafuta(maneno, neno_la_kutafuta) - Kutafuta Neno

Inarudisha mahali pa neno linalotafutwa, au -1 kama halijulikani.

```kwenda
maneno sentensi = "Habari za asubuhi"
namba mahali = tafuta(sentensi, "za")     # 7
namba hakuna = tafuta(sentensi, "jioni")  # -1
```

### awali(maneno, mwanzo) - Kuangalia Mwanzo wa Neno

Inarudisha kweli kama neno linaanza na neno lililotolewa.

```kwenda
maneno sentensi = "Habari za asubuhi"
boolean inaanza = awali(sentensi, "Habari")  # kweli
```

### mwisho(maneno, mwisho) - Kuangalia Mwisho wa Neno

Inarudisha kweli kama neno linaishia na neno lililotolewa.

```kwenda
maneno sentensi = "Habari za asubuhi"
boolean inaishia = mwisho(sentensi, "asubuhi")  # kweli
```

### herufi_kubwa(maneno) - Kubadilisha kuwa Herufi Kubwa

Inabadilisha herufi zote kuwa kubwa.

```kwenda
maneno jina = "amina"
maneno kubwa = herufi_kubwa(jina)  # "AMINA"
```

### herufi_ndogo(maneno) - Kubadilisha kuwa Herufi Ndogo

Inabadilisha herufi zote kuwa ndogo.

```kwenda
maneno jina = "AMINA"
maneno ndogo = herufi_ndogo(jina)  # "amina"
```

### ondoa_nafasi(maneno) - Kuondoa Nafasi za Ziada

Inaondoa nafasi za mwanzo na mwisho wa neno.

```kwenda
maneno na_nafasi = "   Habari Dunia   "
maneno safi = ondoa_nafasi(na_nafasi)  # "Habari Dunia"
```

### gawanya_maneno(maneno, [kigawanyo]) - Kugawanya Neno

Inagawanya neno kuwa sehemu na inarudisha idadi ya sehemu.

```kwenda
maneno sentensi = "moja mbili tatu"
namba idadi = gawanya_maneno(sentensi)        # 3

maneno orodha = "a,b,c,d"
namba idadi2 = gawanya_maneno(orodha, ",")    # 4
```

## Mifano ya Matumizi

### Mfano 1: Kutengeneza Salamu

```kwenda
kazi salamu(maneno jina, maneno wakati) maneno {
    maneno ujumbe = unganisha("Habari za ", wakati, ", ", jina, "!")
    rudisha ujumbe
}

kazi kuu() {
    maneno salamu_asubuhi = salamu("Amina", "asubuhi")
    andika(salamu_asubuhi)  # "Habari za asubuhi, Amina!"
}
```

### Mfano 2: Kuthibitisha Barua Pepe

```kwenda
kazi ni_barua_pepe(maneno barua) boolean {
    namba mahali_at = tafuta(barua, "@")
    namba mahali_dot = tafuta(barua, ".")
    
    boolean ina_at = mahali_at != -1
    boolean ina_dot = mahali_dot != -1
    boolean dot_baada_at = mahali_dot > mahali_at
    
    rudisha ina_at na ina_dot na dot_baada_at
}

kazi kuu() {
    boolean sahihi = ni_barua_pepe("user@example.com")  # kweli
    andika("Barua pepe ni sahihi:", sahihi)
}
```

### Mfano 3: Kuchanganua Jina

```kwenda
kazi angalia_jina(maneno jina) {
    namba urefu_jina = urefu(jina)
    
    kama urefu_jina < 3 {
        andika("Jina ni fupi sana!")
    } sivyo {
        kama urefu_jina > 10 {
            andika("Jina ni refu sana!")
        } sivyo {
            andika("Jina ni la kawaida.")
        }
    }
    
    # Angalia kama jina linaanza na irabu
    maneno herufi_ya_kwanza = kata(jina, 0, 1)
    maneno ndogo = herufi_ndogo(herufi_ya_kwanza)
    
    boolean ni_irabu = ndogo == "a" au ndogo == "e" au ndogo == "i" au ndogo == "o" au ndogo == "u"
    
    kama ni_irabu {
        andika("Jina linaanza na irabu.")
    } sivyo {
        andika("Jina linaanza na konsonanti.")
    }
}
```

## Vidokezo vya Matumizi

1. **Kutumia Nukuu**: Maneno lazima yawe ndani ya nukuu mbili `""`
2. **Kuunganisha**: Unaweza kutumia `+` au `unganisha()` kuunganisha maneno
3. **Kulinganisha**: Tumia `==` na `!=` kulinganisha maneno
4. **Mahali**: Mahali pa herufi huanza kutoka 0
5. **Kesi**: Tumia `herufi_kubwa()` na `herufi_ndogo()` kubadilisha kesi

## Mifano ya Faili

Angalia mifano hii ya faili:
- `examples/string_basic.swh` - Msingi wa maneno
- `examples/string_manipulation.swh` - Kubadilisha maneno
- `examples/string_functions.swh` - Functions za maneno

Hii ni mwongozo kamili wa kutumia maneno katika Kwenda. Tumia mifano hii kuanza kuandika programu zako za maneno!