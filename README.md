# 📦 Contacts Rise Assignment 

This web server API was built with [Go V.1.24](https://golang.org/), [Gin](https://gin-gonic.com/), [MongoDB](https://www.mongodb.com/), and [Prometheus](https://prometheus.io/) for monitoring.

---

## 🚀 Features

- 🧭 REST API with Gin
- 🍃 MongoDB for persistent storage
- 📈 Prometheus metrics exposed at `/metrics`
- 📄 Basic CRUD for Contact entries
  
## 🖥️ Frontend

This project includes a simple HTML, CSS, and JavaScript frontend.

## 📁 Project Structure

```bash
rise-assignment/
├── main.go             # Application entry point
├── api/                # Route handlers
├── db/                 # MongoDB integration and logic
├── static/             # Frontend static files (HTML, JS, CSS)
├── compose.yml         # Docker Compose configuration
└── Dockerfile          # Docker image definition
```

## 🛠️ Setup

### 1. Clone the repo

```bash
git clone https://github.com/danapesah/rise-assignment.git
cd rise-assignment
```

### 2. Setup docker
```bash
docker compose up
```

### 🌐 How to Use

Once the backend server is running (on [http://localhost:8080](http://localhost:8080)), the frontend can be accessed at:
`http://localhost:8080/`

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
| `id`           | `string` | ✅ Yes    | User ID created randomly by mongo|
| `first_name`   | `string` | ✅ Yes    | Contact's first name |
| `last_name`    | `string` | ✅ Yes    | Contact's last name  |
| `phone_number` | `string` | ✅ Yes    | Phone number         |
| `address`      | `string` | ✅ Yes    | Address              |


### 📊 Metrics

| Method | Endpoint    | Description                     |
|--------|-------------|---------------------------------|
| GET    | `/metrics`  | Prometheus metrics for monitoring |

---

## 📬 Testing the Endpoints

### 📥 Create a Contact 15 contacts (Create at least 10 for pagination)

- **Method**: `POST`
- **Endpoint**: `http://localhost:8080/contacts`
- **Body** (JSON):
```json
  {
    "first_name": "Albert",
    "last_name": "Einstein",
    "phone_number": "050-1111111",
    "address": "Princeton, NJ"
  }
```

```json
  {
    "first_name": "Marie",
    "last_name": "Curie",
    "phone_number": "050-2222222",
    "address": "Paris, France"
  }
```
```json
  {
    "first_name": "Isaac",
    "last_name": "Newton",
    "phone_number": "050-3333333",
    "address": "Cambridge, UK"
  }
```
```json
{
    "first_name": "Ada",
    "last_name": "Lovelace",
    "phone_number": "050-4444444",
    "address": "London, UK"
  }
```
```json
 {
    "first_name": "Leonardo",
    "last_name": "da Vinci",
    "phone_number": "050-5555555",
    "address": "Florence, Italy"
  }
```
```json
   {
    "first_name": "Nikola",
    "last_name": "Tesla",
    "phone_number": "050-6666666",
    "address": "New York, NY"
  }
```
```json
   {
    "first_name": "Katherine",
    "last_name": "Johnson",
    "phone_number": "050-7777777",
    "address": "Hampton, VA"
  }
```
```json
  {
    "first_name": "Alan",
    "last_name": "Turing",
    "phone_number": "050-8888888",
    "address": "Manchester, UK"
  }
```
```json
  {
    "first_name": "Galileo",
    "last_name": "Galilei",
    "phone_number": "050-9999999",
    "address": "Pisa, Italy"
  }
```
```json
  {
    "first_name": "Rosalind",
    "last_name": "Franklin",
    "phone_number": "050-1234567",
    "address": "London, UK"
  }
```
```json
   {
    "first_name": "Stephen",
    "last_name": "Hawking",
    "phone_number": "050-2345678",
    "address": "Cambridge, UK"
  }
```
```json
   {
    "first_name": "Thomas",
    "last_name": "Edison",
    "phone_number": "050-3456789",
    "address": "West Orange, NJ"
  }
```
```json
   {
    "first_name": "Jane",
    "last_name": "Goodall",
    "phone_number": "050-4567890",
    "address": "Bournemouth, UK"
  }
```
```json
  {
    "first_name": "Elon",
    "last_name": "Musk",
    "phone_number": "050-5678901",
    "address": "Austin, TX"
  }
```
```json
{
  "first_name": "Alan",
  "last_name": "Watts",
  "phone_number": "050-6789012",
  "address": "San Francisco, CA"
}
```

## 🧪 Testing `GET /contacts`

### For the first pagination
- **Method**: `GET`
- **Endpoint**: `http://localhost:8080/contacts?page=0`

### For the Second pagination
- **Method**: `GET`
- **Endpoint**: `http://localhost:8080/contacts?page=1`

###  All contacts that have the first name "Alan"
- **Method**: `GET`
- **Endpoint**: `http://localhost:8080/contacts?first_name=Alan`

## 🧪 Testing `DELETE /contacts/:id`
### 	Removes a contact by its ID, which you can retrieve using the GET /contacts request
- **Method**: `DELETE`
- **Endpoint**: `http://localhost:8080/contacts/<ID>`

## 🧪 Testing `PUT /contacts`
### 	Edits the contact by its ID, which you can retrieve using the GET /contacts request
- **Method**: `PUT`
- **Endpoint**: `http://localhost:8080/contacts`
- **Body** (JSON):
```json
{
  "id": <string>
  "first_name": "Alan",
  "last_name": "Kahlo",
  "phone_number": "050-7890123",
  "address": "Coyoacán, Mexico City"
}
```
Search for “Alan” again to verify the changes.
