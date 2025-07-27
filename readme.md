# billing & Payroll Management API

### 1. Run the Project
Without Docker
```
go run main.go
```
##### or
with docker
```
docker-compose build --no-cache
docker-compose up
```

### 2. Technology Stack
* Golang (1.21+)

* Fiber (HTTP Framework)

* SQLX + PostgreSQL

* JWT v4

* Gomock + Testify (Testing)

* Docker / Docker Compose

### 📄 3.  API Logging
All API requests and responses are logged to the audit log table using auditlog middleware.

### 4. Env Example
```
APP_NAME = Payroll Payslip
PORT = 8083

DB_HOST = 
DB_USER = 
DB_PASSWORD = 
DB_NAME = 
DB_PORT = 

JWT_SECRET_KEY = 
JWT_LIFESPAN = 
```