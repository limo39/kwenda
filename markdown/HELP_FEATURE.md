# Help Feature Documentation

## Overview
Added comprehensive `--help` and `--version` command-line flags to the Kwenda interpreter.

## Usage

### Display Help
```bash
kwenda --help
# or
kwenda -h
```

Shows comprehensive help information including:
- Usage instructions
- Description of Kwenda
- Example commands
- Basic syntax
- Complete keyword reference
- Feature list
- Documentation links

### Display Version
```bash
kwenda --version
# or
kwenda -v
```

Shows version information including:
- Version number
- Go version used to build
- License information
- Project tagline

### No Arguments
```bash
kwenda
```

Shows brief usage message with pointer to help:
```
Usage: kwenda <filename.swh>
Try 'kwenda --help' for more information.
```

## Help Content

### Sections Included

1. **Header** - ASCII art banner with project name
2. **Usage** - Command syntax and flags
3. **Description** - Brief overview of Kwenda
4. **Examples** - Common usage patterns
5. **Basic Syntax** - Simple code example
6. **Keywords** - Complete list of Swahili keywords with English translations
7. **Features** - Checklist of language capabilities
8. **Documentation** - Links to additional resources
9. **Version** - Current version information

### Keywords Reference

The help includes all Kwenda keywords:

| Swahili | English | Purpose |
|---------|---------|---------|
| kazi | function | Function declaration |
| kuu | main | Main function name |
| andika | write/print | Output to console |
| rudisha | return | Return value |
| namba | number | Number type |
| maneno | words/string | String type |
| boolean | boolean | Boolean type |
| kamusi | dictionary | Dictionary type |
| orodha | list | Array/list type |
| kama | if | If statement |
| sivyo | else | Else statement |
| wakati | while | While loop |
| kwa | for | For loop |
| darasa | class | Class declaration |
| unda | create | Instantiate object |
| hii | this | This/self reference |
| lambda | lambda | Anonymous function |
| leta | bring/import | Import module |

### Features Listed

✓ Variables and data types (numbers, strings, booleans)  
✓ Functions with parameters and return values  
✓ Control flow (if/else, while, for loops)  
✓ Arrays and dictionaries  
✓ Object-oriented programming (classes, inheritance)  
✓ Lambda/anonymous functions  
✓ Closures  
✓ Module system  
✓ Error handling (try/catch)  
✓ Standard library functions  

## Implementation

### Files Modified
- `main.go` - Added `printHelp()` and `printVersion()` functions

### Functions Added

#### `printHelp()`
Displays comprehensive help information with:
- Formatted ASCII art header
- Multi-section documentation
- Keyword reference table
- Feature checklist
- Usage examples

#### `printVersion()`
Displays version information:
- Version number (1.0.0)
- Go version
- License
- Project tagline

### Command-Line Parsing
Updated `main()` function to check for flags before attempting to load files:
1. Check if arguments provided
2. If `--help` or `-h`, show help and exit
3. If `--version` or `-v`, show version and exit
4. Otherwise, treat as filename and execute

## Examples

### Getting Started
```bash
# See all available options
kwenda --help

# Check version
kwenda --version

# Run a program
kwenda examples/hello.swh
```

### Help Output Preview
```
╔═══════════════════════════════════════════════════════════════════════════╗
║                    KWENDA - Swahili Programming Language                  ║
║                    "Move Forward with Technology"                         ║
╚═══════════════════════════════════════════════════════════════════════════╝

USAGE:
    kwenda <filename.swh>              Run a Kwenda program
    kwenda --help                      Show this help message
    kwenda --version                   Show version information
...
```

### Version Output Preview
```
Kwenda Programming Language
Version: 1.0.0
Built with: Go 1.23.3
License: MIT

Kwenda - "Move Forward" in Swahili
A bridge to technology for everyone.
```

## Benefits

1. **User-Friendly** - New users can quickly learn basic usage
2. **Self-Documenting** - Complete keyword reference built-in
3. **Professional** - Standard CLI conventions (--help, --version)
4. **Educational** - Shows Swahili-to-English translations
5. **Accessible** - No need to search external documentation for basics

## Future Enhancements

Potential additions:
- `--examples` flag to list available example programs
- `--keywords` flag for just the keyword reference
- `--interactive` or `--repl` for interactive mode
- `--syntax` flag for syntax highlighting guide
- Localized help in Swahili (`--msaada`)

## Testing

All flags tested and working:
- ✅ `kwenda --help` - Shows full help
- ✅ `kwenda -h` - Shows full help (short form)
- ✅ `kwenda --version` - Shows version info
- ✅ `kwenda -v` - Shows version info (short form)
- ✅ `kwenda` - Shows usage hint
- ✅ `kwenda file.swh` - Normal execution still works

## Conclusion

The help feature makes Kwenda more accessible and professional, providing users with immediate access to documentation and usage information without leaving the command line.
