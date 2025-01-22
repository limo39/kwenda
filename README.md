# Kwenda - A Swahili-Based Programming Language

**Kwenda** (which means "GO" in Swahili) is a simple programming language with Swahili syntax. It is implemented in Go and designed to be easy to learn and use, especially for Swahili speakers. The language supports basic features like variables, functions, and arithmetic operations.

---

## Features

- **Swahili Syntax**: Keywords and constructs are in Swahili (e.g., `kazi` for `func`, `rudisha` for `return`).
- **Simple and Lightweight**: Designed for educational purposes and small projects.
- **Extensible**: Easy to add new features and improve the language.

---

## Example Program

Here’s an example program written in kwenda:

```swahili
kazi kuongeza(a, b) {
    rudisha a + b
}

kazi kuu() {
    namba x = 10
    namba y = 20
    namba jibu = kuongeza(x, y)
    andika(jibu) // Prints the result
}
