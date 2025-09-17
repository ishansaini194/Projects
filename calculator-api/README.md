# Calculator API

A stateless HTTP API written in Go for basic arithmetic operations (add, sub, mul, div).  
This project demonstrates production-ready practices: validation, structured logging, graceful shutdown, panic recovery, and OpenAPI specification.

---

## Features
- Stateless (no DB or in-memory state)
- JSON input/output
- Input validation (division by zero, invalid ops)
- Structured logging with `slog`
- Graceful shutdown & panic recovery
- OpenAPI 3 spec included

---

## API

### POST `/calculate`

**Request**
```json
{
  "op": "add",
  "a": 2,
  "b": 3
}
