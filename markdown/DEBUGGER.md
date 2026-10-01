# Debugger - Interactive Debugging

Kwenda includes an interactive debugger for step-by-step program execution and inspection.

## Starting the Debugger

```bash
kwenda debug program.swh
```

You can also use the alias:
```bash
kwenda dbg program.swh
```

## Debugger Interface

When you start the debugger, you'll see:

```
╔═══════════════════════════════════════════════════════════════════════════╗
║                    KWENDA DEBUGGER                                        ║
║                Interactive Debugging Mode                                 ║
╚═══════════════════════════════════════════════════════════════════════════╝

Type 'h' or 'help' for debugger commands

--> Current line shown here

(kwenda-dbg) 
```

## Core Commands

### Execution Control

#### Run the Program
```
(kwenda-dbg) run
(kwenda-dbg) r
```
Executes the entire program.

#### Continue Execution
```
(kwenda-dbg) continue
(kwenda-dbg) c
(kwenda-dbg) endelea
```
Continue execution until the next breakpoint.

#### Step Into
```
(kwenda-dbg) step
(kwenda-dbg) s
(kwenda-dbg) hatua
```
Execute the next line of code, stepping into functions.

#### Step Over
```
(kwenda-dbg) next
(kwenda-dbg) n
(kwenda-dbg) ifuatayo
```
Execute the next line, but don't step into function calls.

### Breakpoints

#### Set Breakpoint
```
(kwenda-dbg) break 10
(kwenda-dbg) b 10
```
Set a breakpoint at line 10.

```
(kwenda-dbg) break functionName
(kwenda-dbg) b functionName
```
Set a breakpoint at the start of a function.

#### Delete Breakpoint
```
(kwenda-dbg) delete 10
(kwenda-dbg) d 10
(kwenda-dbg) futa 10
```
Remove the breakpoint at line 10.

#### List Breakpoints
```
(kwenda-dbg) listbreaks
(kwenda-dbg) lb
```
Show all current breakpoints.

### Source Code Viewing

#### List Source
```
(kwenda-dbg) list
(kwenda-dbg) l
(kwenda-dbg) orodha
```
List source code around the current line.

```
(kwenda-dbg) list 20
(kwenda-dbg) l 20
```
List source code around line 20.

```
(kwenda-dbg) list 10 30
(kwenda-dbg) l 10 30
```
List source code from line 10 to 30.

### Variable Inspection

#### Print Variable
```
(kwenda-dbg) print myVariable
(kwenda-dbg) p myVariable
(kwenda-dbg) andika myVariable
```
Display the value of `myVariable`.

#### List All Variables
```
(kwenda-dbg) vars
(kwenda-dbg) v
(kwenda-dbg) vigezo
```
Show all variables and their values.

#### Evaluate Expression
```
(kwenda-dbg) eval x + y
(kwenda-dbg) e x + y
(kwenda-dbg) tathmini x + y
```
Evaluate and display the result of an expression.

### Watch Variables

#### Add Watch
```
(kwenda-dbg) watch counter
(kwenda-dbg) w counter
(kwenda-dbg) angalia counter
```
Add `counter` to the watch list. The debugger will show its value after each step.

#### List Watches
```
(kwenda-dbg) listwatch
(kwenda-dbg) lw
```
Show all watched variables and their current values.

#### Remove Watch
```
(kwenda-dbg) deletewatch counter
(kwenda-dbg) dw counter
```
Remove `counter` from the watch list.

### Function Inspection

#### List Functions
```
(kwenda-dbg) funcs
(kwenda-dbg) f
(kwenda-dbg) kazi
```
Display all defined functions with their signatures.

### Help and Exit

#### Show Help
```
(kwenda-dbg) help
(kwenda-dbg) h
(kwenda-dbg) msaada
```
Display all debugger commands.

#### Quit Debugger
```
(kwenda-dbg) quit
(kwenda-dbg) q
(kwenda-dbg) ondoka
```
Exit the debugger.

## Example Debugging Session

### Sample Program (debug_example.swh)
```swahili
kazi factorial(n) {
    kama n <= 1 {
        rudisha 1
    }
    rudisha n * factorial(n - 1)
}

kazi kuu() {
    namba result = factorial(5)
    andika("Factorial of 5 is: ", result)
}
```

### Debugging Session
```
$ kwenda debug debug_example.swh

╔═══════════════════════════════════════════════════════════════════════════╗
║                    KWENDA DEBUGGER                                        ║
╚═══════════════════════════════════════════════════════════════════════════╝

--> 1: kazi factorial(n) {

(kwenda-dbg) b 9
Breakpoint set at line 9

(kwenda-dbg) lb
Breakpoints:
════════════
  Line 9: enabled

(kwenda-dbg) l
Source Code:
════════════
--> 1: kazi factorial(n) {
    2:     kama n <= 1 {
    3:         rudisha 1
    4:     }
    5:     rudisha n * factorial(n - 1)
    6: }
    7: 
    8: kazi kuu() {
 B  9:     namba result = factorial(5)
   10:     andika("Factorial of 5 is: ", result)
   11: }

(kwenda-dbg) r
Running program...

(kwenda-dbg) v
Variables:
══════════
result = <not yet assigned>

(kwenda-dbg) w result
Added watch for 'result'

(kwenda-dbg) s
Step...

(kwenda-dbg) lw
Watched Variables:
══════════════════
result = 120

(kwenda-dbg) p result
result = 120

(kwenda-dbg) e result * 2
result = 240

(kwenda-dbg) funcs
Functions:
══════════
  factorial(n: namba)
  kuu()

(kwenda-dbg) c
Factorial of 5 is: 120

(kwenda-dbg) q
Exiting debugger...
```

## Breakpoint Features

### Line Breakpoints
Set breakpoints at specific line numbers:

```
(kwenda-dbg) b 15
Breakpoint set at line 15
```

### Multiple Breakpoints
Set multiple breakpoints:

```
(kwenda-dbg) b 10
(kwenda-dbg) b 20
(kwenda-dbg) b 35
(kwenda-dbg) lb
Breakpoints:
════════════
  Line 10: enabled
  Line 20: enabled
  Line 35: enabled
```

### Breakpoint Management
Delete specific breakpoints:

```
(kwenda-dbg) d 20
Breakpoint deleted at line 20
```

## Variable Watching

Watch multiple variables simultaneously:

```
(kwenda-dbg) w counter
Added watch for 'counter'

(kwenda-dbg) w total
Added watch for 'total'

(kwenda-dbg) lw
Watched Variables:
══════════════════
counter = 5
total = 150
```

Remove watches:

```
(kwenda-dbg) dw counter
Removed watch for 'counter'
```

## Expression Evaluation

Test expressions in the current context:

```
(kwenda-dbg) e x + y
result = 30

(kwenda-dbg) e x * 2
result = 20

(kwenda-dbg) e my_array[0]
result = 5
```

## Source Code Navigation

View different parts of the source:

```
# View current position
(kwenda-dbg) l

# View around line 25
(kwenda-dbg) l 25

# View lines 10-30
(kwenda-dbg) l 10 30
```

## Bilingual Commands

All commands support both Swahili and English:

| Command Type | Swahili | English |
|-------------|---------|---------|
| Help | msaada | help |
| Continue | endelea | continue |
| Step | hatua | step |
| Next | ifuatayo | next |
| List | orodha | list |
| Print | andika | print |
| Variables | vigezo | vars |
| Functions | kazi | funcs |
| Quit | ondoka | quit |
| Watch | angalia | watch |
| Evaluate | tathmini | eval |
| Delete | futa | delete |

## Tips for Effective Debugging

### 1. Strategic Breakpoints
Set breakpoints at key decision points:

```
(kwenda-dbg) b 10  # Before important calculation
(kwenda-dbg) b 25  # Inside critical loop
(kwenda-dbg) b 40  # Before function return
```

### 2. Watch Critical Variables
Monitor variables that change frequently:

```
(kwenda-dbg) w counter
(kwenda-dbg) w accumulator
(kwenda-dbg) w status
```

### 3. Use Expression Evaluation
Test hypotheses about your code:

```
(kwenda-dbg) e array.urefu()  # Check array length
(kwenda-dbg) e x % 2 == 0      # Test if even
(kwenda-dbg) e name.herufi_kubwa()  # Check uppercase
```

### 4. Inspect All Variables
When confused, list everything:

```
(kwenda-dbg) v  # See all variables
(kwenda-dbg) f  # See all functions
```

### 5. Source Navigation
Keep track of where you are:

```
(kwenda-dbg) l  # Always check current context
```

## Common Debugging Scenarios

### Finding Logic Errors
1. Set breakpoint before suspected error
2. Run to breakpoint
3. Check variable values
4. Step through execution
5. Watch for unexpected values

### Tracking Variable Changes
1. Add variable to watch list
2. Set breakpoints in relevant sections
3. Continue execution
4. Observe watch list updates

### Understanding Control Flow
1. Set breakpoints at each branch
2. Step through conditions
3. Verify correct path is taken

### Inspecting Function Behavior
1. Set breakpoint at function entry
2. Check parameter values
3. Step through function
4. Verify return value

## Limitations

- Line tracking requires source code
- Breakpoints are set by line number
- Watch expressions evaluated at stops
- Performance impact during debugging

## Best Practices

1. **Start Small**: Debug small sections at a time
2. **Use Watches**: Monitor key variables throughout execution
3. **Strategic Breakpoints**: Set breakpoints at critical points only
4. **Evaluate Freely**: Use expression evaluation to test hypotheses
5. **Clean Up**: Remove breakpoints and watches when done
6. **Document Findings**: Note issues discovered during debugging

## Future Enhancements

Planned debugger improvements:
- Conditional breakpoints
- Call stack inspection
- Memory visualization
- Time-travel debugging
- Remote debugging support

The Kwenda debugger makes finding and fixing bugs easier with its bilingual, intuitive interface!
