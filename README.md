# Simple Go HTTP API

A basic HTTP API built with Go that demonstrates fundamental web server concepts including JSON responses, query parameters, request body parsing, and HTTP method validation.

## What the Service Does

This is a simple HTTP server with multiple endpoints that return JSON responses:

- **`/`** - Welcome message
- **`/health`** - Returns the service health status
- **`/status`** - Returns service status with uptime information
- **`/hello`** - Returns a personalized greeting from query parameters (GET)
- **`/greet`** - Returns a personalized greeting from JSON body (POST)

## Features

- ✅ JSON request and response handling
- ✅ Query parameter parsing
- ✅ Request body parsing with validation
- ✅ HTTP method validation (GET/POST)
- ✅ Proper error handling with status codes
- ✅ Reusable error response function
- ✅ Service uptime tracking
- ✅ Clean code organization with separate handler file

## Project Structure

```
simple-go-http-api/
├── main.go       # Main server setup and routing
├── handlers.go   # HTTP handler functions
├── api.http      # REST Client test file
└── README.md     # This file
```

## How to Run It

### Prerequisites
- Go installed on your system (version 1.16 or higher)

### Steps

1. Make sure you have both `main.go` and `handlers.go` in the same directory

2. Run the server:
   ```bash
   go run .
   ```
   or 
   ```
   go run main.go handlers.go
   ```

3. You should see:
   ```
   Server starting on http://localhost:8080/
   ```

4. The server is now running at `http://localhost:8080`

## API Endpoints

### Root Endpoint

**Request:**
```bash
GET http://localhost:8080/
```

**Response:**
```json
{
  "message": "Welcome to the API!\n"
}
```

**Status Code:** 200 OK

---

### Health Check Endpoint

**Request:**
```bash
GET http://localhost:8080/health
```

**Response:**
```json
{
  "status": "healthy"
}
```

**Status Code:** 200 OK

---

### Status Endpoint (with Uptime)

**Request:**
```bash
GET http://localhost:8080/status
```

**Response:**
```json
{
  "service": "running",
  "uptime": "2h15m30.5s"
}
```

**Status Code:** 200 OK

The uptime shows how long the service has been running since it started.

---

### Hello Endpoint (Query Parameters - GET)

**Success Request:**
```bash
GET http://localhost:8080/hello?name=Alice
```

**Response:**
```json
{
  "message": "Hello, Alice!\n"
}
```

**Status Code:** 200 OK

**Error Request (missing name):**
```bash
GET http://localhost:8080/hello
```

**Response:**
```json
{
  "error": "Name is required"
}
```

**Status Code:** 400 Bad Request

---

### Greet Endpoint (JSON Body - POST)

**Success Request:**
```bash
POST http://localhost:8080/greet
Content-Type: application/json

{
  "name": "Alice"
}
```

**Response:**
```json
{
  "message": "Hello Alice"
}
```

**Status Code:** 200 OK

**Error Request (missing name):**
```bash
POST http://localhost:8080/greet
Content-Type: application/json

{
  "name": ""
}
```

**Response:**
```json
{
  "error": "Name is required"
}
```

**Status Code:** 400 Bad Request

**Error Request (invalid JSON):**
```bash
POST http://localhost:8080/greet
Content-Type: application/json

{invalid json}
```

**Response:**
```json
{
  "error": "Invalid JSON body"
}
```

**Status Code:** 400 Bad Request

**Error Request (wrong HTTP method):**
```bash
GET http://localhost:8080/greet
```

**Response:**
```json
{
  "error": "Only POST method is allowed"
}
```

**Status Code:** 405 Method Not Allowed

---

## Testing the API

### Option 1: Using the Browser (GET requests only)

- Root: `http://localhost:8080/`
- Health check: `http://localhost:8080/health`
- Status with uptime: `http://localhost:8080/status`
- Hello with name: `http://localhost:8080/hello?name=YourName`

### Option 2: Using curl (Command Line)

**GET requests:**
```bash
curl http://localhost:8080/hello?name=Alice
```

**POST requests:**
```bash
curl -X POST http://localhost:8080/greet \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice"}'
```

### Option 3: Using VS Code REST Client Extension

1. Install the "REST Client" extension in VS Code
2. Open the `api.http` file
3. Click "Send Request" above any request
4. View the response in the split panel

The `api.http` file includes test cases for all endpoints with various scenarios.

### Option 4: Using Postman

1. Download Postman from https://www.postman.com/downloads/
2. Create requests for each endpoint
3. Set the appropriate HTTP method (GET or POST)
4. Add request body for POST requests

## Response Format

All endpoints return JSON responses with appropriate headers:
- `Content-Type: application/json`

## Error Handling

The API uses standard HTTP status codes:

| Status Code | Meaning | When Used |
|-------------|---------|-----------|
| 200 | OK | Request succeeded |
| 400 | Bad Request | Invalid input (missing parameters, invalid JSON) |
| 405 | Method Not Allowed | Wrong HTTP method used |
| 500 | Internal Server Error | Server error (not currently implemented) |

All errors return a consistent JSON format:
```json
{
  "error": "Error message here"
}
```

## Code Structure

### Main Components

**`main.go`**
- Server initialization
- Route registration
- Port configuration

**`handlers.go`**
- Request handler functions
- JSON encoding/decoding
- Error handling logic
- Data structures (User, StatusResponse, errorResponse)

**`api.http`**
- Test cases for all endpoints
- Example requests for development

### Key Functions

- `writeError()` - Reusable function for sending error responses
- `getRoot()` - Handles root endpoint
- `getHealth()` - Health check endpoint
- `getStatus()` - Service status with uptime
- `getHello()` - Greeting from query parameters
- `getGreet()` - Greeting from JSON body with POST validation

## Stopping the Server

Press `Ctrl + C` in the terminal where the server is running.

## API Summary Table

| Endpoint | Method | Parameters | Content-Type | Success Response | Error Codes |
|----------|--------|------------|--------------|------------------|-------------|
| `/` | GET | None | application/json | `{"message": "Welcome..."}` | - |
| `/health` | GET | None | application/json | `{"status": "healthy"}` | - |
| `/status` | GET | None | application/json | `{"service": "running", "uptime": "..."}` | - |
| `/hello` | GET | `name` (query) | application/json | `{"message": "Hello, {name}!"}` | 400 |
| `/greet` | POST | `{"name": "..."}` (body) | application/json | `{"message": "Hello {name}"}` | 400, 405 |

## Learning Concepts Demonstrated

This project demonstrates:
- Setting up an HTTP server in Go
- Creating RESTful API endpoints
- Handling different HTTP methods (GET, POST)
- Parsing query parameters
- Parsing JSON request bodies
- Encoding JSON responses
- Error handling with proper status codes
- Code organization with multiple files
- Using structs for structured data
- Request validation
- Testing APIs with REST Client

## Next Steps

Potential enhancements:
- Add more endpoints (UPDATE, DELETE)
- Connect to a database
- Add authentication/authorization
- Add request logging middleware
- Add unit tests
- Add environment configuration
- Implement graceful shutdown
- Add rate limiting
