package command

import (
	"praksa_nimbustech/server_global_error_handler/model"
	"praksa_nimbustech/server_global_error_handler/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling GetUser request...")
	if r.URL.Query().Get("error") == "true" {
		panic("Simulirana greška u servisu")
	}

	w.Header().Set("Content-Type", "application/json")

	ee3 := json.NewEncoder(w).Encode(storage.Users)
	if ee3 != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Println("GetUser done successfully!")
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling addUser request...")

	var newUser model.User
	err1 := json.NewDecoder(r.Body).Decode(&newUser)
	if err1 != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	storage.Users = append(storage.Users, newUser)
	w.Header().Set("Content-Type", "application/json")
	err2 := json.NewEncoder(w).Encode(newUser)
	if err2 != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Println("User added successfully!")
}
