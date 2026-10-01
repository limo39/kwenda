# REPL - Interactive Shell

Kwenda includes a powerful REPL (Read-Eval-Print Loop) for interactive programming and testing.

## Starting the REPL

```bash
kwenda repl
```

You can also use these aliases:
```bash
kwenda interactive
kwenda shell
```

## Features

### 1. Immediate Expression Evaluation
Type any expression and see the result immediately:

```swahili
kwenda> 2 + 2
4

kwenda> 10 * 5
50

kwenda> "Habari" + " Dunia"
"Habari Dunia"
```

### 2. Variables
Define and use variables across commands:

```swahili
kwenda> namba x = 10
kwenda> namba y = 20
kwenda> x + y
30

kwenda> maneno name = "Kwenda"
kwenda> name
"Kwenda"
```

### 3. Functions
Define and call functions interactively:

```swahili
kwenda> kazi greet(name) {
....... andika("Habari, ", name, "!")
....... }
kwenda> greet("Mwalimu")
Habari, Mwalimu !
```

### 4. Multiline Input
The REPL automatically detects multiline input when you use braces:

```swahili
kwenda> kazi factorial(n) {
....... kama n <= 1 {
.......     rudisha 1
....... } sivyo {
.......     rudisha n * factorial(n - 1)
....... }
....... }
kwenda> factorial(5)
120
```

### 5. Arrays and Dictionaries

```swahili
kwenda> orodha nums = [1, 2, 3, 4, 5]
kwenda> nums
[1, 2, 3, 4, 5]

kwenda> kamusi person = {"name": "Ali", "age": 25}
kwenda> person
{"name": Ali, "age": 25}
```

### 6. List Comprehensions

```swahili
kwenda> squares = [x * x kwa x katika [1, 2, 3, 4, 5]]
kwenda> squares
[1, 4, 9, 16, 25]
```

### 7. Enums

```swahili
kwenda> aina Status { PENDING, ACTIVE, COMPLETED }
kwenda> current = Status.ACTIVE
kwenda> current
Status.ACTIVE

kwenda> current == Status.ACTIVE
kweli
```

### 8. Object-Oriented Programming

```swahili
kwenda> darasa Person {
....... namba age
....... maneno name
....... 
....... kazi unda(n, a) {
.......     hii.name = n
.......     hii.age = a
....... }
....... 
....... kazi greet() {
.......     andika("Hi, I'm ", hii.name)
....... }
....... }
kwenda> p = unda Person("Amina", 30)
kwenda> p.greet()
Hi, I'm Amina
```

## REPL Commands

### Special Commands
- `ondoka` or `exit` - Exit the REPL
- `msaada` or `help` - Show help information
- `historia` or `history` - Display command history
- `safisha` or `clear` - Clear all variables and functions

### Usage Examples

#### View Command History
```swahili
kwenda> historia

Command History / Historia ya Amri:
════════════════════════════════════
  1: namba x = 10
  2: x * 5
  3: kazi add(a, b) { rudisha a + b }
  4: add(10, 20)
```

#### Clear Variables
```swahili
kwenda> namba x = 10
kwenda> x
10
kwenda> safisha
Variables cleared / Vigezo vimefutwa
kwenda> x
x
```

## Tips and Tricks

### 1. Quick Testing
Use the REPL to quickly test Kwenda expressions:

```swahili
kwenda> 2 ** 10
1024

kwenda> [1, 2, 3] + [4, 5, 6]
[1, 2, 3, 4, 5, 6]
```

### 2. Function Prototyping
Test function logic before adding to your program:

```swahili
kwenda> kazi isPrime(n) {
....... kama n < 2 {
.......     rudisha uwongo
....... }
....... namba i = 2
....... wakati i * i <= n {
.......     kama n % i == 0 {
.......         rudisha uwongo
.......     }
.......     i = i + 1
....... }
....... rudisha kweli
....... }
kwenda> isPrime(17)
kweli
```

### 3. Learning and Exploration
Experiment with language features:

```swahili
kwenda> # Test string methods
kwenda> "Habari".urefu()
6

kwenda> "HABARI".herufi_ndogo()
"habari"

kwenda> # Test math operations
kwenda> 10 / 3
3.3333333333333335

kwenda> 10 % 3
1
```

### 4. Immediate Feedback
See results instantly without writing a full program:

```swahili
kwenda> orodha nums = [1, 2, 3, 4, 5]
kwenda> sum = 0
kwenda> kwa n katika nums {
....... sum = sum + n
....... }
kwenda> sum
15
```

## Differences from Script Mode

### 1. Automatic Result Display
In the REPL, expressions are automatically evaluated and displayed:

```swahili
# In REPL - shows result automatically
kwenda> 5 + 5
10

# In script mode - need andika() to see output
namba result = 5 + 5
andika(result)  # Must explicitly print
```

### 2. Persistent State
Variables and functions persist between commands:

```swahili
kwenda> namba counter = 0
kwenda> counter = counter + 1
kwenda> counter
1
kwenda> counter = counter + 1
kwenda> counter
2
```

### 3. Interactive Debugging
Test code snippets and fix errors immediately:

```swahili
kwenda> namba x = "10"  # Oops, string instead of number
kwenda> x + 5
"105"  # Unexpected result!
kwenda> x = 10  # Fix it
kwenda> x + 5
15  # Correct!
```

## Best Practices

1. **Use for Learning**: Perfect for understanding how Kwenda features work
2. **Quick Prototyping**: Test algorithms and logic before writing full programs
3. **Interactive Debugging**: Test expressions and functions in isolation
4. **Documentation**: Try examples from documentation interactively
5. **Exploration**: Experiment with new language features safely

## Limitations

- No file I/O commands work as expected (use script mode for files)
- Long programs are better suited for script files
- History is not persisted between sessions
- Module imports should be done in script mode

## Examples Session

Here's a complete REPL session showing various features:

```swahili
$ kwenda repl

╔═══════════════════════════════════════════════════════════════════════════╗
║              KWENDA REPL - Interactive Swahili Programming                ║
║                       Version 1.0.0                                       ║
╚═══════════════════════════════════════════════════════════════════════════╝

Type 'ondoka' or 'exit' to quit
Type 'msaada' or 'help' for help

kwenda> # Define a function
kwenda> kazi fibonacci(n) {
....... kama n <= 1 {
.......     rudisha n
....... }
....... rudisha fibonacci(n - 1) + fibonacci(n - 2)
....... }

kwenda> # Test it
kwenda> fibonacci(10)
55

kwenda> # Create array with list comprehension
kwenda> fibs = [fibonacci(i) kwa i katika [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]]
kwenda> fibs
[0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55]

kwenda> # Define enum
kwenda> aina Color { RED, GREEN, BLUE }
kwenda> favorite = Color.BLUE
kwenda> favorite
Color.BLUE

kwenda> # View history
kwenda> historia

Command History / Historia ya Amri:
════════════════════════════════════
  1: kazi fibonacci(n) { ... }
  2: fibonacci(10)
  3: fibs = [fibonacci(i) kwa i katika [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]]
  4: fibs
  5: aina Color { RED, GREEN, BLUE }
  6: favorite = Color.BLUE
  7: favorite

kwenda> ondoka
Kwaheri! (Goodbye!)
```

## Bilingual Support

All commands support both Swahili and English:

| Swahili | English | Description |
|---------|---------|-------------|
| ondoka | exit | Exit REPL |
| msaada | help | Show help |
| historia | history | Show history |
| safisha | clear | Clear variables |

This makes the REPL accessible to both Swahili speakers and those learning the language!
