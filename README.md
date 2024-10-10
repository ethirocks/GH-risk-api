# GH-risk-api

## Risk Management API

The **Risk Management API** is a web-based service designed to manage risks. It provides an HTTP-based interface to create, update, and retrieve risks, with data persisted in memory for the purposes of this assignment.

---

## Project Overview

The API supports basic CRUD operations on risks, which include:

- **Create a new risk**: `POST /v1/risks`
- **Retrieve all risks**: `GET /v1/risks`
- **Retrieve a specific risk by ID**: `GET /v1/risks/{id}`
- **Update an existing risk**: `PUT /v1/risks/{id}`

---

## Risk Data Structure

Each risk contains the following fields:

| Field         | Type   | Description                                                    |
|---------------|--------|----------------------------------------------------------------|
| `ID`          | UUID   | Unique identifier for the risk (generated on creation).        |
| `State`       | string | Must be one of: `open`, `closed`, `accepted`, `investigating`. |
| `Title`       | string | A brief title describing the risk.                             |
| `Description` | string | A detailed description of the risk.                            |

---

## Features

- **CRUD Operations**:
  - Create, update, and retrieve risks.
  - List all available risks in memory.

- **Logging**:
  - Request logging for every HTTP method and path.
  - Error logging for any failed request or operation.

- **In-Memory Data Storage**:
  - All risks are stored in memory (`map[string]Risk`).
  - No external databases are used.

---

## Endpoints

### 1. `POST /v1/risks`

**Description**: Create a new risk.

**Request Body**:

```json
{
  "title": "Risk Title",
  "description": "Risk Description",
  "state": "open"
}
```

**Response**:

- `201 Created`: Returns the created risk.
- `400 Bad Request`: If the input data is invalid.

### 2. `GET /v1/risks`

**Description**: Retrieve all risks.

**Response**:

- `200 OK`: Returns a list of all risks.
- `204 No Content`: If no risks exist.

### 3. `GET /v1/risks/{id}`

**Description**: Retrieve a specific risk by ID.

**Response**:

- `200 OK`: Returns the risk if found.
- `204 No Content`: If the risk is not found.

### 4. `PUT /v1/risks/{id}`

**Description**: Update a specific risk by ID.

**Request Body**:

```json
{
  "title": "Updated Title",
  "description": "Updated Description",
  "state": "closed"
}
```

**Response**:

- `200 OK`: If the risk is updated successfully.
- `400 Bad Request`: If the input data is invalid.
- `204 No Content`: If the risk does not exist.

## Running the Project

### Prerequisites

- Go: Make sure Go is installed on your machine.

### Steps

1. Clone the repository:

```bash
git clone https://github.com/ethirajmudhaliar/GH-risk-api.git
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the server:

```bash
go run main.go
```

4. Run the test coverage:

```bash
go test ./... -cover
```

## Architecture

### 1. **main.go**

- Initializes the HTTP server and sets up the routes and middleware.
- **Router Setup**: Defines API routes (`/v1/risks`, `/v1/risks/{id}`) using the Gorilla Mux router.
- **Logging Middleware**: Logs details of incoming requests and their processing time.

### 2. **v1 Package**

Contains the business logic for handling risk operations, including:
- `CreateRisk`: Handles the creation of a new risk.
- `GetRisks`: Lists all risks stored in memory.
- `GetRiskByID`: Fetches a specific risk by its unique ID.
- `UpdateRisk`: Updates an existing risk by ID.

### 3. **common Package**

Includes in-memory storage and response helper functions:
- `RiskStorage`: In-memory storage using a map and slice for risks.
- `RespondWithJSON` & `RespondWithError`: Functions to standardize JSON responses and error handling.

### 4. **logger Package**

Provides logging capabilities for the application:
- `Info`: Logs informational messages.
- `Error`: Logs error messages.
- `LogRequest`: Logs HTTP requests, including method, URL, and processing time.

### 5. **validation Package**

Contains validation logic for the API:
- **State Validation**: Ensures that the `state` field in a risk is valid (`open`, `closed`, `accepted`, `investigating`).

---

## Testing

### Unit Tests

Tests are included for individual handlers and functions located in the corresponding `_test.go` files.

Examples:
- `create_risk_test.go`: Tests for the `CreateRisk` function.
- `get_risks_test.go`: Tests for the `GetRisks` function.
|
|
|

## API Example Usage

### Create a New Risk

```bash
curl -X POST http://localhost:8080/v1/risks \
  -H "Content-Type: application/json" \
  -d '{"title": "New Risk", "description": "Risk description", "state": "open"}'
```

### Get All Risks

```bash
curl http://localhost:8080/v1/risks
```

### Get a Specific Risk by ID

```bash
curl http://localhost:8080/v1/risks/id
```

### Get a Specific Risk by ID

```bash
curl http://localhost:8080/v1/risks/1curl -X PUT http://localhost:8080/v1/risks/id \
  -H "Content-Type: application/json" \
  -d '{"title": "Updated Risk", "description": "Updated description", "state": "closed"}'
```
