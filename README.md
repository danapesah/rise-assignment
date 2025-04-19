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

├── static/              # db (mongodb) handlers

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



