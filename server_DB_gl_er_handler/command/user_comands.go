package command

import (
	"praksa_nimbustech/server_DB_gl_er_handler/model"
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB
var err error

const DNS = "postgres://yeieycvd:xxxxxxxxxxxxxxx@trumpet.db.elephantsql.com/yeieycvd"

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("error") == "true" {
		panic("Simulirana greška u servisu")
	}
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	fmt.Println("Handling GetUser request...")

	w.Header().Set("Content-Type", "application/json")

	var users []model.User

	if err1 := DB.Find(&users).Error; err1 != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err2 := json.NewEncoder(w).Encode(users); err2 != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Println("GetUser done successfully!")
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	fmt.Println("Handling addUser request...")
	w.Header().Set("Content-Type", "application/json")

	var newUser model.User

	defer r.Body.Close()

	if err3 := json.NewDecoder(r.Body).Decode(&newUser); err3 != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Println("Before DB.Create(&newUser) !!!!!")

	DB.Create(&newUser)

	if err4 := json.NewEncoder(w).Encode(newUser); err4 != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Println("User added successfully!")
}
