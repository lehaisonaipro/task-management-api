# Task Management API

This is a Task Management API built with Golang and MongoDB, providing features for managing tasks assigned to employees.

## Setup and Running with Docker

### Prerequisites
- Docker
- Docker Compose

### Clone the Repository
```sh
git clone https://github.com/lehaisonaipro/task-management-api.git
cd task-management-api
```

### Build and Run the API with Docker
```sh
docker-compose up --build
```
- The API will be available at [http://localhost:8888](http://localhost:8888).

### Docker Compose File (docker-compose.yml)
```yaml
version: '3.8'

services:
    api:
        build: .
        ports:
            - "8888:8888"
        environment:
            - MONGO_URI=mongodb://mongo:27017/task_db
        depends_on:
            - mongo
        restart: always

    mongo:
        image: mongo:latest
        container_name: task_management_db
        ports:
            - "27017:27017"
        volumes:
            - mongo_data:/data/db

volumes:
    mongo_data:
```

## 🏗️ Building and Running Manually

### Install Dependencies
```sh
go mod tidy
```

### Run the Application
```sh
go run main.go
```
- The API will be available at [http://localhost:8888](http://localhost:8888).

## 📌 API Endpoints

All responses are in JSON format.

### 1️⃣ User Registration

**Endpoint:** `POST /api/users/register`

**Request:**
```json
{
    "name": "John Doe",
    "email": "johndoe@example.com",
    "password": "securepassword",
    "role": "employee"
}
```

**Response:**
```json
{
    "message": "User registered successfully",
    "user_id": "12345"
}
```

### 2️⃣ User Login

**Endpoint:** `POST /api/users/login`

**Request:**
```json
{
    "email": "johndoe@example.com",
    "password": "securepassword"
}
```

**Response:**
```json
{
    "token": "jwt-token"
}
```

### 3️⃣ Create a Task

**Endpoint:** `POST /api/tasks`

**Authentication:** Employer Only (Bearer Token Required)

**Request:**
```json
{
    "title": "Complete report",
    "description": "Finish the project report by next Monday",
    "status": "pending",
    "due_date": "2025-03-10"
}
```

**Response:**
```json
{
    "message": "Task created successfully",
    "task_id": "task123"
}
```

### 4️⃣ Assign Task to Employee

**Endpoint:** `PUT /api/tasks/{task_id}/assign`

**Authentication:** Employer Only

**Request:**
```json
{
    "assigned_to": "employee123"
}
```

**Response:**
```json
{
    "message": "Task assigned successfully"
}
```

### 5️⃣ View Tasks by Employee

**Endpoint:** `GET /api/tasks?assigned_to={employee_id}`

**Authentication:** Employer or Assigned Employee

**Response:**
```json
[
    {
        "id": "task1",
        "title": "Complete report",
        "status": "in_progress",
        "assigned_to": "employee123"
    },
    {
        "id": "task2",
        "title": "Submit invoice",
        "status": "pending",
        "assigned_to": "employee123"
    }
]
```

### 6️⃣ View All Tasks (Employer)

**Endpoint:** `GET /api/tasks`

**Authentication:** Employer Only

**Response:**
```json
[
    {
        "id": "task1",
        "title": "Complete report",
        "status": "pending",
        "assigned_to": "employee123"
    },
    {
        "id": "task2",
        "title": "Review PR",
        "status": "in_progress",
        "assigned_to": "employee456"
    }
]
```

## ⚙️ Config File

Create a `config.yaml` file.

## 🔍 Testing the API

Use Postman or cURL to test the endpoints.

### Example cURL for login:
```sh
curl -X POST http://localhost:8888/api/users/login \
         -H "Content-Type: application/json" \
         -d '{"email": "johndoe@example.com", "password": "securepassword"}'
```