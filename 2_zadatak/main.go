package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var userList = []User{
	{ID: 1, Name: "Petar Petrovic"},
	{ID: 2, Name: "Jovan Pavlovic"},
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userList)
}

func main() {
	// Endpoint za health check
	http.HandleFunc("/api/v1/health-check", healthCheckHandler)

	// Endpoint za korisnike
	http.HandleFunc("/api/v1/users", usersHandler)

	http.ListenAndServe(":8080", nil)
}
