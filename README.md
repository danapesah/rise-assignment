# 📦 Contacts Rise Assignment 

This web server API was built with [Go V.1.24](https://golang.org/), [Gin](https://gin-gonic.com/), [MongoDB](https://www.mongodb.com/), and [Prometheus](https://prometheus.io/) for monitoring.

---

## 🚀 Features

- 🧭 REST API with Gin
- 🍃 MongoDB for persistent storage
- 📈 Prometheus metrics exposed at `/metrics`
- 📄 Basic CRUD for Contact entries
- 
## 🖥️ Frontend

This project includes a simple HTML, CSS, and JavaScript frontend.

## 📁 Project Structure
/app

├── main.go              # Entry point

├── api/                 # Route handlers

├── db/                  # db (mongodb) handlers

├── static/              # frontend static files

├── compose.yml

└── Dockerfile

---

## 🛠️ Setup

### 1. Clone the repo

git clone https://github.com/yourname/my-go-api.git
cd riseAssignment 

### 2. Setup docker
docker compose up

### 🌐 How to Use

Once the backend server is running (on [http://localhost:8080](http://localhost:8080)), the frontend can be accessed at:
http://localhost:8080/

## 📡 Available Endpoints

The backend exposes the following REST API endpoints:

### 🧾 Contacts

| Method | Endpoint           | Description                     |
|--------|--------------------|---------------------------------|
| GET    | `/contacts`        | Get all contacts                |
| POST   | `/contacts`        | Create a new contact            |
| PUT    | `/contacts`        | Update an existing contact      |
| DELETE | `/contacts/:id`    | Delete a contact                |


### 🔍 Query Parameters for `GET /contacts`

You can filter contacts and paginate results using these query parameters:

| Parameter     | Type     | Description                       |
|---------------|----------|-----------------------------------|
| `page`        | `int`    | Page number for pagination (Used if no other paramters are provided) |
| `first_name`  | `string` | Filter by first name              |
| `last_name`   | `string` | Filter by last name               |
| `address`     | `string` | Filter by address                 |
| `phone_number`| `string` | Filter by phone number            |

### ✉️ POST `/contacts` Parameters

To create a new contact, send a `POST` request with a **JSON body** containing the following fields:

| Field          | Type     | Required | Description          |
|----------------|----------|----------|----------------------|
| `first_name`   | `string` | ✅ Yes    | Contact's first name |
| `last_name`    | `string` | ✅ Yes    | Contact's last name  |
| `phone_number` | `string` | ✅ Yes    | Phone number         |
| `address`      | `string` | ✅ Yes    | Address              |


### ✉️ PUT `/contacts` Parameters

To update a contact, send a `PUT` request with a **JSON body** containing the following fields:

| Field          | Type     | Required | Description          |
|----------------|----------|----------|----------------------|
| `first_name`   | `string` | ✅ Yes    | Contact's first name |
| `last_name`    | `string` | ✅ Yes    | Contact's last name  |
| `phone_number` | `string` | ✅ Yes    | Phone number         |
| `address`      | `string` | ✅ Yes    | Address              |


### 📊 Metrics

| Method | Endpoint    | Description                     |
|--------|-------------|---------------------------------|
| GET    | `/metrics`  | Prometheus metrics for monitoring |

---



