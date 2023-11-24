package command

import (
	"praksa_nimbustech/server/model"
	"praksa_nimbustech/server/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	// postavljanje COuntent type u JSON format
	w.Header().Set("Content-Type", "application/json")
	// prevod liste users iz go u json format
	json.NewEncoder(w).Encode(storage.Users)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling addUser request...")
	var newUser model.User
	//dodavanje novog korisnika (čitanog iz tela HTTP zahteva) u listu korisnika (users).
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Dodaj novog korisnika u listu
	storage.Users = append(storage.Users, newUser)

	//Zatim šalje JSON reprezentaciju novog korisnika kao odgovor na HTTP zahtev.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
	fmt.Println("User added successfully!")
}
