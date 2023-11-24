package api

import (
	"praksa_nimbustech/server/command"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Router() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/users", command.GetUser).Methods("GET")
	r.HandleFunc("/api/v1/users", command.AddUser).Methods("POST")

	fmt.Println("Server started on port 8083")
	log.Fatal(http.ListenAndServe(":8083", r))
}
