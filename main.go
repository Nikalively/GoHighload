package main

import (
	"gohighload/handlers"
	"gohighload/metrics"
	"gohighload/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Use(utils.RateLimitMiddleware)
	r.Use(metrics.MetricsMiddleware)

	r.HandleFunc("/api/users", handlers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/api/users/{id}", handlers.GetUserHandler).Methods("GET")
	r.HandleFunc("/api/users", handlers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/api/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/api/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/metrics", handlers.MetricsHandler).Methods("GET")

	utils.Logger.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
