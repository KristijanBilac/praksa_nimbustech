package api

import (
	"praksa_nimbustech/server_global_error_handler/command"
	"praksa_nimbustech/server_global_error_handler/error"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Router() {
	r := mux.NewRouter()

	r.Use(error.GlobalErrorHandler)

	r.HandleFunc("/api/v1/users", command.GetUser).Methods("GET")
	r.HandleFunc("/api/v1/users", command.AddUser).Methods("POST")

	fmt.Println("Server started on port 8088")
	log.Fatal(http.ListenAndServe(":8088", r))
}
