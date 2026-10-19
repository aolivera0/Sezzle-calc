# Sezzle Calculator

A full-stack calculator application built as a technical assessment project. It uses a React and TypeScript frontend and a Go backend that exposes a REST API for arithmetic operations. Developed by Angel H. Oivera Pinzón. Link of repository: [GitHub](https://github.com/aolivera0/Sezzle-calc).

## Features

- Addition, subtraction, multiplication, division
- Exponentiation, square root, percentage
- Division-by-zero and negative square-root validation
- Invalid JSON and unsupported-operation handling
- JSON-based REST API (`POST /api/calculate`)
- React frontend with input validation, loading state, and error display
- Responsive design (basic mobile support)
- Backend unit tests and code coverage reporting
- Frontend unit tests with Vitest and React Testing Library

## Technology Stack

### Frontend
- React 19
- TypeScript
- Vite
- Vitest + React Testing Library
- React Compiler (via Babel preset)

### Backend
- Go 1.24
- Standard library `net/http` server
- `encoding/json` for request/response handling
- Go standard `testing` package

## Project Structure

```text
Sezzle-calc/
├── backend/
│   ├── main.go
│   ├── go.mod
│   ├── internal/
│   │   ├── calculator/
│   │   │   ├── calculator.go
│   │   │   └── calculator_test.go
│   │   └── httpapi/
│   │       ├── handler.go
│   │       └── handler_test.go
│   ├── coverage.out      (generated, gitignored)
│   └── coverage.html     (generated, gitignored)
├── frontend/
│   ├── index.html
│   ├── package.json
│   ├── vite.config.ts
│   ├── vitest.config.ts
│   ├── tsconfig.json
│   ├── tsconfig.app.json
│   ├── tsconfig.node.json
│   ├── eslint.config.js
│   └── src/
│       ├── main.tsx
│       ├── App.tsx
│       ├── App.test.tsx
│       ├── App.css
│       ├── index.css
│       └── test/
│           └── setup.ts
├── README.md
└── .gitignore
```

## Prerequisites

- Go 1.24+
- Node.js 18+ and npm

## Backend — Setup and Run

From the repository root:

```bash
cd backend
go run main.go
```

The API server starts on `http://localhost:8080`.

### Run backend tests

```bash
cd backend
go test -v ./...
```

### Generate coverage report

```bash
cd backend
go test -coverpkg=./... -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

HTML coverage report:

```bash
cd backend
go tool cover -html=coverage.out -o coverage.html
```

## Frontend — Setup and Run

From the repository root:

```bash
cd frontend
npm install
npm run dev
```

The Vite dev server starts on `http://localhost:5173` and proxies `/api` requests to the Go backend at `http://localhost:8080` (see `vite.config.ts`).

### Run frontend tests

```bash
cd frontend
npm test
```

### Build for production

```bash
cd frontend
npm run build
```

## Backend API

### Calculate

`POST /api/calculate`

**Request body:**

```json
{
  "operation": "add",
  "a": 10,
  "b": 5
}
```

**Successful response (`200 OK`):**

```json
{
  "result": 15
}
```

**Supported operation values:**

| Operation | Description | Notes |
|---|---|---|
| `add` | Adds two numbers | `a + b` |
| `subtract` | Subtracts the second number from the first | `a - b` |
| `multiply` | Multiplies two numbers | `a × b` |
| `divide` | Divides the first number by the second | `a / b` |
| `power` | Raises the first number to the power of the second | `a ^ b` |
| `sqrt` | Square root of the first number | `√a` (field `b` ignored) |
| `percentage` | `b` percent of `a` | `a × b / 100` |

**Error responses:**

| Status | Meaning |
|---|---|
| `400 Bad Request` | Invalid JSON, unsupported operation, division by zero, or square root of a negative number |
| `405 Method Not Allowed` | HTTP method other than POST |
| `500 Internal Server Error` | Failed to encode the response |

### CURL Examples

Addition:

```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"add","a":10,"b":5}'
```

Division:

```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"divide","a":10,"b":2}'
```

Division by zero (error):

```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"divide","a":5,"b":0}'
```

Square root of negative number (error):

```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"sqrt","a":-3,"b":0}'
```

Unsupported operation (error):

```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"log","a":10,"b":2}'
```

## Testing

Backend tests are written using Go's standard testing package and cover each operation plus edge cases (division by zero, square root of a negative number, invalid JSON, unsupported operation, wrong HTTP method, empty body).

Frontend tests use Vitest with React Testing Library and mock the global `fetch` to verify successful results, API errors, and input validation without sending real requests.

## Design Decisions and Assumptions

- **Single calculate endpoint:** All operations are handled by one `POST /api/calculate` endpoint that selects behavior based on the `operation` field, keeping the API surface small and easy to document.
- **`float64` for all numbers:** Go's `float64` is used throughout the backend to support integers and decimals uniformly.
- **`b` ignored for `sqrt`:** The square-root operation only requires `a`; `b` is accepted in the request but unused, and the frontend disables the `b` input when `sqrt` is selected.
- **Percentage semantics:** `percentage` computes `b` percent of `a` (`a × b / 100`).
- **No CORS middleware in development:** The Vite dev server proxies `/api` to the Go backend, avoiding CORS configuration in development. 
- **React Compiler enabled:** The frontend uses the React Compiler Babel preset for automatic memoization.
- **Error surfacing:** Backend errors return JSON `{"error": "..."}` with appropriate HTTP status; the frontend displays them in an `aria-live` polite region for accessibility.
- **Mobile support:** A `@media (max-width: 520px)` rule stacks the input fields into a single column.

## Prompts Used

The following prompts were used during development:

1. "If I want to build an API RESTful for a calculaltor with Go but I do not know the language, what should I learn first? explain me how the language works and guide me through the process."

2. "Help me plan a full-stack calculator for a technical assessment using React with TypeScript for the frontend and Go for the backend. Explain the responsibilities of each layer and guide me through the implementation step by step rather than providing the entire solution at once."

3. "Help me configure Vitest and React Testing Library for my Vite application. Write tests for rendering the calculator, displaying successful results, handling API errors, and preventing requests when inputs are empty."

4. "Help me write a README with setup instructions, how to run frontend and backend, API examples, design decisions, and prompts used."
