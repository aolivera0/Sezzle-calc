# Sezzle Calculator

A full-stack calculator application built as a technical assessment project. It uses a React and TypeScript frontend and a Go backend that exposes a REST API for arithmetic operations.

## Features

- Addition
- Subtraction
- Multiplication
- Division
- Division-by-zero validation
- JSON-based REST API
- HTTP error handling
- Backend unit tests and code coverage reporting

## Technology Stack

### Frontend
- React
- TypeScript
- Vite

### Backend
- Go
- Standard library HTTP server
- JSON request and response handling
- Go testing package

## Project Structure

```text
Sezzle-calc/
├── backend/
│   ├── internal/
│   │   ├── calculator/
│   │   └── httpapi/
│   ├── go.mod
│   ├── main.go
│   ├── coverage.out
│   └── coverage.html
├── frontend/
│   ├── src/
│   ├── package.json
│   └── vite.config.ts
└── README.md
```

## Backend API

### Calculate

`POST /api/calculate`

Request body:

```json
{
  "operation": "add",
  "a": 10,
  "b": 5
}
```

Successful response (`200 OK`):

```json
{
  "result": 15
}
```

Supported operation values:

| Operation | Description |
|---|---|
| `add` | Adds two numbers |
| `subtract` | Subtracts the second number from the first |
| `multiply` | Multiplies two numbers |
| `divide` | Divides the first number by the second |

Error responses:

| Status | Meaning |
|---|---|
| `400 Bad Request` | Invalid JSON, unsupported operation, or division by zero |
| `405 Method Not Allowed` | HTTP method other than POST |

## Testing

Backend tests are written using Go's standard testing package.

Run the backend tests:

```bash
cd backend
go test -v ./...
```

Generate a combined coverage report:

```bash
go test -coverpkg=./... -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

Generate an HTML coverage report:

```bash
go tool cover -html=coverage.out -o coverage.html
```
