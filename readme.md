# Quan-Lang

Quan-Lang is a simple interpreted programming language implemented in Go. It is designed for learning and experimenting with language design, interpreters, and compilers. Quan-Lang features a custom lexer, parser, and interpreter, and supports variables, functions, arithmetic operations, and conditionals.

---

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Example](#example)
- [Lexer Example](#lexer-example)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- **Lexer**: Converts source code into tokens.
- **Parser**: Builds an abstract syntax tree (AST) from tokens.
- **Interpreter**: Evaluates the AST and executes code.
- **Variables**: Assignment and usage.
- **Functions**: User-defined functions with parameters and return values.
- **Conditionals**: `if`/`else` statements.
- **Arithmetic**: Supports `+`, `-`, `*`, `/`, `%`, `^`, and comparison operators.
- **Block Scoping**: Functions and conditionals have their own scope.
- **Objects**: Object literals and property access.
- **Arrays**: Array literals, indexing, and array utility functions.
- **Floats**: Native support for floating-point numbers and arithmetic.
- **EmptyReturn**: Return without value from a function.
- **Null**: Null value.
- **Template Strings**: JS-like template strings with `${}` expressions.
- **Debug Options**: Built-in debug utilities and options for tracing/interpreter output.
- **Extensible**: Modular design for easy extension.
- **WebAssembly**: Support WebAssembly, this engine can run from browser
- **New APIs**: Support fetch(), toJson(), toMap()
- **Parser**: int(), float(), string(), bool()

---

## Project Structure

```
quan-lang/
├── array/         # Array utilities
├── debug/         # Debug utilities
├── env/           # Environment (variable/function scope)
├── expression/    # AST node definitions
├── helper/        # Helper functions
├── intepreter/    # Interpreter logic
├── lexer/         # Lexer (tokenizer)
├── paraser/       # Parser (AST builder)
├── quan-lang/     # Language entry point
├── token/         # Token definitions
├── go.mod
├── main.go
├── readme.md
├── sample-program.quan
```

---

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/quan-lang.git
   cd quan-lang
   ```

2. **Install Go (if not already installed):**
   [Download Go](https://golang.org/dl/)

3. **Run the interpreter:**
   ```sh
   go run main.go
   ```

---

## Usage

You can write Quan-Lang code as a string and execute it using the provided API, or run a `.quan` script file.

---

## Example

```go
program := `
    fn calculateInterest(principal, rate, time) {
        return principal * rate * time / 100;
    }
    interest = calculateInterest(loanAmount, 3, 1);
`

env, _ := lang.Execuate(program, &env.Env{
    Vars: map[string]interface{}{"loanAmount": 100000},
})

fmt.Printf("Interest: %f\n", env.Vars["interest"])
```

---

## Array, Float, and Debug Example

```go
program := `
    arr = [1, 2, 3.5, 4];
    sum = 0.0;
    for (i = 0; i < len(arr); i = i + 1) {
        sum = sum + arr[i];
    }
    debug("Sum of array:", sum);
`
env, _ := lang.Execuate(program, &env.Env{
    Debug: true, // Enable debug output
})

fmt.Printf("Sum: %f\n", env.Vars["sum"])
```

---

## Lexer Example

The lexer scans the input string and produces a slice of tokens.  
Here is a simplified excerpt from [`lexer/lexer.go`](lexer/lexer.go):


---

## Contributing

Contributions are welcome! Please open issues or submit pull requests for improvements and bug fixes.

---

## License

This project is licensed under the MIT License.


## Builds
Buid binary
```shell
go build main.go 
```

WebAssembliy
```shell
GOOS=js GOARCH=wasm go build -o qlang.wasm main.go     
```




## Maintainer and Creator
Name: jianshangquan
Github: [Account](https://github.com/jianshangquan)