# Go CLI Calculator

A powerful command-line calculator built in Go to master core language concepts, CLI design, and clean architecture.

---

## Goal

This project is not just a calculator — it's a **learning vehicle** to deeply understand Go by building something practical and extensible.

---

## What I Will Master

### Go Fundamentals

* Variables, constants, and types
* Functions and modular design
* Error handling (idiomatic Go)
* Structs and methods

### CLI Development

* Reading user input (stdin)
* Command parsing
* Interactive CLI loops
* Handling invalid inputs gracefully

### Project Architecture

* Package-based structure
* Separation of concerns
* Clean and maintainable code organization

### Standard Library Mastery

* `fmt` → input/output
* `bufio` → efficient input reading
* `strconv` → string ↔ number conversion
* `strings` → parsing expressions

### Concepts (Phase 2+)

* Interfaces for extensibility
* Unit testing (`testing` package)
* Error wrapping and custom errors
* Logging (optional enhancement)

---

## What You Will Learn (Deep Skills)

* How Go handles **compiled binaries**
* Writing **idiomatic Go code**
* Designing **CLI tools like real-world utilities**
* Building **extensible systems (add more operations later)**
* Thinking in **packages instead of files**

---

## Project Structure

```
go-cli-calculator/
│
├── go.mod
├── main.go
│
├── cmd/
│   └── calculator/
│       └── main.go        # Entry point for CLI
│
├── internal/
│   ├── parser/
│   │   └── parser.go      # Parse user input (e.g. "2 + 3")
│   │
│   ├── evaluator/
│   │   └── evaluator.go   # Logic to compute expressions
│   │
│   └── cli/
│       └── cli.go         # CLI loop & user interaction
│
├── pkg/
│   └── operations/
│       ├── add.go
│       ├── sub.go
│       ├── mul.go
│       └── div.go         # Reusable operations
│
└── README.md
```

---

## Features (Planned)

### Phase 1

* Basic arithmetic: `+ - * /`
* Interactive CLI mode
* Input validation
* Error handling

### Phase 2

* Expression parsing (`2 + 3 * 4`)
* Operator precedence
* Parentheses support

### Phase 3

* History tracking
* Commands like:

  * `history`
  * `clear`
  * `exit`

### Phase 4 (Advanced)

* Plugin-like operations (extendable)
* Configurable precision
* Support for floats & scientific calculations

---

## How to Run

```bash
go run main.go
```

Or build binary:

```bash
go build -o calc
./calc
```

---

## Future Improvements

* Add unit tests for parser & evaluator
* Benchmark performance
* Convert into a reusable Go package
* Add flags (`--expression "2+3"`)

---

## Philosophy

> "Build small. Understand deeply. Scale later."

This project focuses on:

* Writing **clean Go**
* Understanding **why things work**
* Building **real developer intuition**

---

## Final Outcome

By the end of this project, you will:

* Be comfortable writing Go CLI tools
* Understand Go project structuring
* Be ready to build:

  * CLI apps
  * APIs
  * System tools

---

## 🧑‍💻 Author

Built as part of a journey to master Go through real-world projects.

---
