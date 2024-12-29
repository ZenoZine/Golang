package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type Customer struct {
	ID        uint32
	Name      string
	Role      string
	Email     string
	Phone     uint64
	Contacted bool
}

var db *sql.DB
var database = map[string]Customer{
	"1": {
		ID:        1,
		Name:      "John Doe",
		Role:      "Manager",
		Email:     "john@example.com",
		Phone:     1234567890,
		Contacted: false,
	},
	"2": {
		ID:        2,
		Name:      "Bob Smith",
		Role:      "Engineer",
		Email:     "bob@example.com",
		Phone:     9876543210,
		Contacted: false,
	},
	"3": {
		ID:        3,
		Name:      "Alice Johnson",
		Role:      "Sales Rep",
		Email:     "alice@example.com",
		Phone:     5555555555,
		Contacted: true,
	},
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newCustomer Customer
	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newCustomer)

	_, err := db.Exec("INSERT INTO customers (id, name, role, email, phone, contacted) VALUES ($1, $2, $3, $4, $5, $6)",
		newCustomer.ID, newCustomer.Name, newCustomer.Role, newCustomer.Email, newCustomer.Phone, newCustomer.Contacted)
	if err != nil {
		http.Error(w, "Unable to add customer", http.StatusInternalServerError)
		return
	}
	// 6. Return data type
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCustomer)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	query := "DELETE FROM customers WHERE id = $1"
	result, err := db.Exec(query, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete customer"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Customer not found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"error": "Customer deleted successfully"})
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	var customer Customer
	query := "SELECT id, name, role, email, phone, contacted FROM customers WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&customer.ID, &customer.Name, &customer.Role, &customer.Email, &customer.Phone, &customer.Contacted)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Customer not found"})
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch customer"})
		return
	}

	json.NewEncoder(w).Encode(customer)
}

func getCustomers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := "SELECT id, name, role, email, phone, contacted FROM customers"
	rows, err := db.Query(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch customers"})
		return
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Role, &customer.Email, &customer.Phone, &customer.Contacted); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to parse customer data"})
			return
		}
		customers = append(customers, customer)
	}

	json.NewEncoder(w).Encode(customers)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	var updatedCustomer Customer
	reqBody, err := io.ReadAll(r.Body)
	if err != nil || json.Unmarshal(reqBody, &updatedCustomer) != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	query := "UPDATE customers SET name = $1, role = $2, email = $3, phone = $4, contacted = $5 WHERE id = $6"
	result, err := db.Exec(query, updatedCustomer.Name, updatedCustomer.Role, updatedCustomer.Email, updatedCustomer.Phone, updatedCustomer.Contacted, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update customer"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Customer not found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCustomer)
}

func updateCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updates map[string]Customer
	reqBody, err := io.ReadAll(r.Body)
	if err != nil || json.Unmarshal(reqBody, &updates) != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	tx, err := db.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to start transaction"})
		return
	}

	for id, updatedCustomer := range updates {
		query := "UPDATE customers SET name = $1, role = $2, email = $3, phone = $4, contacted = $5 WHERE id = $6"
		_, err := tx.Exec(query, updatedCustomer.Name, updatedCustomer.Role, updatedCustomer.Email, updatedCustomer.Phone, updatedCustomer.Contacted, id)
		if err != nil {
			tx.Rollback()
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Failed to update customer with ID %s", id)})
			return
		}
	}

	if err := tx.Commit(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to commit transaction"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Customers updated successfully"})
}

func main() {
	if loadErr := godotenv.Load(); loadErr != nil {
		log.Fatalf("Error loading .env file: %v", loadErr)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	if dsn == "" {
		log.Fatalf("Database connection string is empty")
	}

	var err error
	db, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Creating the router
	router := mux.NewRouter().StrictSlash(true)

	// Handler functions
	router.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCustomers(w, r)
		case http.MethodPost:
			addCustomer(w, r)
		case http.MethodPut:
			updateCustomers(w, r)
		default:
			http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
		}
	})

	router.HandleFunc("/customers/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCustomer(w, r)
		case http.MethodPut:
			updateCustomer(w, r)
		case http.MethodDelete:
			deleteCustomer(w, r)
		default:
			http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
		}
	})

	srv := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	go func() {
		fmt.Println("Server starting on port 3000...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutting down server...")

	if err := db.Close(); err != nil {
		log.Fatalf("Database close failed %v", err)
	}

	if err := srv.Close(); err != nil {
		log.Fatalf("Server close failed: %v", err)
	}
}
