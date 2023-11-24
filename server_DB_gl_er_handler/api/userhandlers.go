package api

import (
	"awesomeProject/server_DB_gl_er_handler/command"
	"awesomeProject/server_DB_gl_er_handler/error"
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

	fmt.Println("Server started on port 8084")
	log.Fatal(http.ListenAndServe(":8084", r))
}
