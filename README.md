# Golang
# Customer Management API

This project is a simple Customer Management API built using Go, Gorilla Mux for routing, and PostgreSQL as the database. It allows users to perform CRUD (Create, Read, Update, Delete) operations on customer data via HTTP endpoints.

## Key Features

- **Retrieve Customers**: Fetch all customer records or a specific customer by ID.
- **Add Customers**: Add new customer records to the PostgreSQL database.
- **Update Customers**: Update information for an existing customer or multiple customers.
- **Delete Customers**: Remove customer records from the database.
- **Database Integration**: Uses PostgreSQL to store and manage customer data.

## Prerequisites

Before starting the project, ensure you have the following installed:

1. [Go](https://golang.org/dl/) (1.18 or higher)
2. [PostgreSQL](https://www.postgresql.org/download/)
3. [Git](https://git-scm.com/)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/customer-management-api.git
   cd customer-management-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up the PostgreSQL database:

   - Create a database named `customers_db` (or your preferred name).
   - Create a `customers` table with the following schema:
     ```sql
     CREATE TABLE customers (
         id SERIAL PRIMARY KEY,
         name VARCHAR(255) NOT NULL,
         role VARCHAR(255),
         email VARCHAR(255) UNIQUE,
         phone BIGINT,
         contacted BOOLEAN DEFAULT FALSE
     );
     ```

4. Create a `.env` file in the project root with the following environment variables:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_postgres_username
   DB_PASSWORD=your_postgres_password
   DB_NAME=customers_db
   ```

## Running the Project

1. Start the PostgreSQL server.
2. Run the Go application:
   ```bash
   go run main.go
   ```
3. The server will start on [http://localhost:3000](http://localhost:3000).

## API Endpoints

### `/customers`
- `GET`: Retrieve all customers.
- `POST`: Add a new customer.
- `PUT`: Update multiple customers.

### `/customers/{id}`
- `GET`: Retrieve a customer by ID.
- `PUT`: Update a customer by ID.
- `DELETE`: Delete a customer by ID.

## Example Request/Response

### Add Customer
**Request:**
```bash
POST /customers
Content-Type: application/json

{
    "name": "Jane Doe",
    "role": "Developer",
    "email": "jane@example.com",
    "phone": 1234567890,
    "contacted": true
}
```
**Response:**
```json
{
    "id": 4,
    "name": "Jane Doe",
    "role": "Developer",
    "email": "jane@example.com",
    "phone": 1234567890,
    "contacted": true
}
```

## Future Enhancements

- Add authentication and authorization.
- Implement pagination for customer lists.
- Add logging and monitoring.

---

If you encounter any issues or have questions, feel free to open an issue in the repository.


