# gradecalc

A lightweight, zero-dependency Go application and REST API for calculating university semester averages (GPA) and accumulated academic credits.

Designed with clean architecture and domain-driven separation: the core calculation engine is an isolated package that powers both a command-line interface (CLI) and an HTTP REST API.

---

## Features

* **3 Evaluation Schemas**:
  * **Type 1**: Exam only.
  * **Type 2**: Exam + TD (Directed work) weighted by a configurable `split`.
  * **Type 3**: Exam + TD + TP (Practical work).
* **Credit Tracking**: Automatically awards module credits when a module grade $\ge 10.0 / 20$.
* **Dual Interface**:
  * **CLI**: Supports file-based and interactive terminal prompt modes.
  * **HTTP API**: High-performance REST endpoint (`POST /calculate`) built with standard library Go 1.22+ routing.
* **Standard Library Only**: Zero external dependencies (`net/http`, `encoding/json`, `flag`, `io`).
* **Stream-Based I/O**: Efficient `io.Reader` streaming for memory-efficient deserialization.

---

## Project Structure

```text
.
├── internal/
│   └── calculator/       # Pure domain package (business logic & types)
│       ├── types.go      # Data models (Modules, Module, Grade)
│       ├── helpers.go    # Computation math & stream decoding (FromJSON)
│       └── calculator_test.go # Integration and unit tests
├── cmd/
│   ├── cli/              # Command-line interface application
│   │   ├── main.go       # Flag parsing and CLI flow
│   │   └── helpers.go    # Interactive terminal prompting & output formatting
│   └── api/              # HTTP REST API server
│       ├── main.go       # Server initialization (port 8080)
│       ├── handler.go    # HTTP POST /calculate handler
│       └── types.go      # HTTP response structures
├── Taskfile.yml          # Task runner commands
├── test.json             # Sample module configuration
└── go.mod
```

---

## Configuration Schema (`test.json`)

Modules and grading rules are configured via JSON:

```json
{
  "split": 0.6,
  "modules": [
    {
      "name": "analysis",
      "coff": 4,
      "credit": 6,
      "type": 2,
      "grades": {
        "exam": 10.0,
        "td": 14.0
      }
    },
    {
      "name": "algorithm",
      "coff": 4,
      "credit": 6,
      "type": 3,
      "grades": {
        "exam": 14.75,
        "td": 16.75,
        "tp": 16.5
      }
    }
  ]
}
```

### Grade Formulas:
* **Type 1**: $\text{Grade} = \text{Exam}$
* **Type 2**: $\text{Grade} = (\text{Exam} \times \text{split}) + (\text{TD} \times (1 - \text{split}))$
* **Type 3**: $\text{Grade} = (\text{Exam} \times \text{split}) + \left(\text{TD} \times \frac{1 - \text{split}}{2}\right) + \left(\text{TP} \times \frac{1 - \text{split}}{2}\right)$

---

## Quickstart

### 1. Command-Line Interface (CLI)

#### Run with a JSON file:
```bash
go run ./cmd/cli -file test.json
```
Output:
```text
grade: 10.853333, credit: 16
```

#### Run in Interactive Mode:
Prompts you to enter grades manually module by module in your terminal:
```bash
go run ./cmd/cli -file test.json -interactive
```
Example session:
```text
analysis
  TD: 12
  Exam: 14
algorithm
  TP: 15
  TD: 16
  Exam: 14
grade: 13.800000, credit: 10
```

---

### 2. HTTP REST API

#### Start the server:
```bash
go run ./cmd/api
```
The server will start listening on `:8080`.

#### Send a calculation request:
```bash
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  --data-binary @test.json
```

#### Response (`200 OK`):
```json
{
  "grade": 10.853333,
  "credit": 16
}
```

---

## Running Tests

Run the test suite using the standard Go testing tool:

```bash
go test -v ./...
```

---

## License

This project is licensed under the MIT License.
