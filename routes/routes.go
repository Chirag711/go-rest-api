package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Chirag711/go-rest-api/controllers"
	"github.com/gorilla/mux"
)

// -------------------- REGISTER ALL ROUTES --------------------

func RegisterUserRoutes(r *mux.Router) {

	// Health Check
	r.HandleFunc("/health", HealthCheck).Methods("GET")

	// API Versioning
	api := r.PathPrefix("/api/v1").Subrouter()

	// -------------------- USER ROUTES --------------------

	api.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	api.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	api.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	api.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	// -------------------- FILE ROUTE --------------------

	api.HandleFunc("/upload", controllers.UploadFile).Methods("POST")

	// -------------------- EXCEL EXPORT ROUTE --------------------

	api.HandleFunc("/export/users", controllers.ExportUsers).Methods("GET")
}

// -------------------- HEALTH CHECK --------------------

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status":  "success",
		"message": "Server is running",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
