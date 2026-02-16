# REST API

A REST API built with Go for managing student records using SQLite database with graceful shutdown support.

## Features

- Create, retrieve, and list student records
- SQLite database storage
- Request validation using go-playground/validator
- YAML-based configuration
- Graceful server shutdown
- Structured JSON responses

## Tech Stack

- Go 1.25.6
- SQLite3 (github.com/mattn/go-sqlite3)
- go-playground/validator/v10 for request validation
- cleanenv for configuration management
- Native Go HTTP server

## Project Structure

```
.
├── cmd/
│   └── rest-api/
│       └── main.go              # Application entry point
├── config/
│   └── local.yaml               # Configuration file
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration loader
│   ├── http/
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go   # Student HTTP handlers
│   ├── storage/
│   │   ├── storage.go           # Storage interface
│   │   └── sqlite/
│   │       └── sqlite.go        # SQLite implementation
│   ├── type/
│   │   └── types.go             # Data types
│   └── utils/
│       └── response/
│           └── response.go      # HTTP response utilities
├── storage/                      # Database files directory
└── go.mod                        # Go module dependencies
```

## Prerequisites

- Go 1.25.6 or higher
- SQLite3

## Setup

1. Clone the repository:

```bash
git clone <repository-url>
cd rest-api
```

2. Install dependencies:

```bash
go mod download
```

3. Configure the application:

Set the `CONFIG_PATH` environment variable or use the `-config` flag:

```bash
export CONFIG_PATH=config/local.yaml
```

Alternatively, use the flag:

```bash
go run cmd/rest-api/main.go -config config/local.yaml
```

## Configuration

Edit `config/local.yaml`:

```yaml
env: "dev"
storage_path: "storage/storage.db"
http_server:
  address: ":8000"
```

**Environment Variables:**

- `CONFIG_PATH` - Path to configuration file (optional if using `-config` flag)
- `ENV` - Environment name (can be set via config)

## Running the Application

```bash
go run cmd/rest-api/main.go -config config/local.yaml
```

The server will start on `http://localhost:8000`

## API Endpoints

### Create Student

```
POST /api/students
```

**Request Body:**

```json
{
  "Name": "John Doe",
  "Email": "john@example.com",
  "Age": 20
}
```

**Response (201):**

```json
{
  "id": 1
}
```

**Error Response (400):**

```json
{
  "status": "Error",
  "error": "field Name is required field"
}
```

### Get Student by ID

```
GET /api/students/{id}
```

**Response (200):**

```json
{
  "Id": 1,
  "Name": "John Doe",
  "Email": "john@example.com",
  "Age": 20
}
```

### Get All Students

```
GET /api/students/
```

**Response (200):**

```json
[
  {
    "Id": 1,
    "Name": "John Doe",
    "Email": "john@example.com",
    "Age": 20
  },
  {
    "Id": 2,
    "Name": "Jane Smith",
    "Email": "jane@example.com",
    "Age": 22
  }
]
```

## Example cURL Requests

Create a student:

```bash
curl -X POST http://localhost:8000/api/students \
  -H "Content-Type: application/json" \
  -d '{"Name":"John Doe","Email":"john@example.com","Age":20}'
```

Get student by ID:

```bash
curl http://localhost:8000/api/students/1
```

Get all students:

```bash
curl http://localhost:8000/api/students/
```

## Demo

![Demo](assets/demo.gif)

## Building for Production

```bash
go build -o rest-api cmd/rest-api/main.go
./rest-api -config config/local.yaml
```

## Database Schema

The application automatically creates the following table:

```sql
CREATE TABLE IF NOT EXISTS students (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT,
  email TEXT,
  age INTEGER
)
```

## License

MIT

(ai generated)
