package main

import (
	"praksa_nimbustech/server_DB_gl_er_handler/api"
	"praksa_nimbustech/server_DB_gl_er_handler/storage"
)

func main() {
	storage.InitialMigration()
	api.Router()

}
