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
- **Extensible**: Modular design for easy extension.

---


## Project Structure

```
quan-lang/ 
├── array/ # Array utilities 
├── debug/ # Debug utilities 
├── env/ # Environment (variable/function scope) 
├── expression/ # AST node definitions 
├── helper/ # Helper functions 
├── intepreter/ # Interpreter logic 
├── lexer/ # Lexer (tokenizer) 
├── paraser/ # Parser (AST builder) 
├── quan-lang/ # Language entry point 
├── token/ # Token definitions 
├── go.mod 
├── main.go 
└── readme.md
```
