# Go REST API with MongoDB

A scalable **RESTful API built using Golang and MongoDB** with a modular project structure.
This project demonstrates how to build APIs using **MongoDB, Gorilla Mux Router, and various services like Email, Excel export, and File Upload**.

---

## Features

* RESTful API using **Golang**
* **MongoDB** database integration
* **Gorilla Mux** router
* **Environment configuration using `.env`**
* **File upload support**
* **Excel export functionality**
* **Email sending service**
* Modular and scalable project architecture

---

## Tech Stack

* **Golang**
* **MongoDB**
* **Gorilla Mux Router**
* **Excelize (Excel generation)**
* **Gomail (Email service)**
* **Godotenv (.env configuration)**

---

## Project Structure

```text
go-rest-api/
│
├── cmd/
│   └── main.go
│
├── config/
│   └── db.go
│
├── models/
│   └── user.go
│
├── controllers/
│   └── user_controller.go
│
├── routes/
│   └── routes.go
│
├── services/
│   ├── email_service.go
│   ├── excel_service.go
│   └── file_service.go
│
├── uploads/
│
├── .env
└── go.mod
```

---

## Installation

### Clone the Repository

```bash
git clone https://github.com/yourusername/go-rest-api.git
cd go-rest-api
```

---

### Initialize Go Module (if not initialized)

```bash
go mod init go-rest-api
```

---

### Install Dependencies

```bash
go get go.mongodb.org/mongo-driver/mongo
go get github.com/gorilla/mux
go get github.com/joho/godotenv
go get github.com/xuri/excelize/v2
go get gopkg.in/gomail.v2
```

---

## MongoDB Setup

Install MongoDB locally or use **MongoDB Atlas**.

Create a database:

```
go_api_db
```

Collection example:

```
users
```

---

## Environment Configuration

Create a `.env` file in the root directory.

Example:

```env
PORT=8080
MONGO_URI=mongodb://localhost:27017
DB_NAME=go_api_db

EMAIL_HOST=smtp.gmail.com
EMAIL_PORT=587
EMAIL_USER=your_email@gmail.com
EMAIL_PASS=your_password
```

---

## Run the Application

```bash
go run cmd/main.go
```

Server will start at:

```
http://localhost:8080
```

---

## API Endpoints (Example)

| Method | Endpoint        | Description    |
| ------ | --------------- | -------------- |
| POST   | /api/users      | Create user    |
| GET    | /api/users      | Get all users  |
| GET    | /api/users/{id} | Get user by ID |
| PUT    | /api/users/{id} | Update user    |
| DELETE | /api/users/{id} | Delete user    |

---

## File Upload

Uploaded files are stored in:

```
uploads/
```

The `file_service.go` handles file operations.

---

## Excel Export

The `excel_service.go` generates Excel files using **Excelize**.

Example use cases:

* Export user data
* Generate reports
* Download spreadsheet

---

## Email Service

The `email_service.go` uses **Gomail** for sending emails.

Example usage:

* Send welcome emails
* Send notifications
* Send reports

---

## License

This project is licensed under the **MIT License**.

---

## Author

Built using **Golang + MongoDB + Gorilla Mux** with modular architecture.
