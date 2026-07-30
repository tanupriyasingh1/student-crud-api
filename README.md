# Student CRUD API (Golang + PostgreSQL)

A simple REST API to manage student records — built with Go and PostgreSQL.

## Features
- Create a student
- View all students
- Update a student
- Delete a student

## Tech Stack
- Golang
- PostgreSQL
- net/http (standard library, no framework)

## Setup

1. Install PostgreSQL and create a database:
```sql
CREATE DATABASE studentdb;
```

2. Update the connection string in `main.go` with your own Postgres username/password:
```go
connStr := "host=localhost port=5432 user=postgres password=yourpassword dbname=studentdb sslmode=disable"
```

3. Install dependencies:
```bash
go mod tidy
```

4. Run the server:
```bash
go run main.go
```

Server starts on `http://localhost:8080`

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /students | Get all students |
| POST | /students | Create a student |
| PUT | /students?id=1 | Update a student |
| DELETE | /students?id=1 | Delete a student |

## Example Requests

**Create:**
```bash
curl -X POST http://localhost:8080/students -d '{"name":"tanu","age":21,"grade":"A"}'
```

**View all:**
```bash
curl http://localhost:8080/students
```

**Update:**
```bash
curl -X PUT "http://localhost:8080/students?id=1" -d '{"name":"tanu","age":22,"grade":"A+"}'
```

**Delete:**
```bash
curl -X DELETE "http://localhost:8080/students?id=1"
```
