package main

import (
	"awesomeProject/server_DB_gl_er_handler/api"
	"awesomeProject/server_DB_gl_er_handler/storage"
)

func main() {
	storage.InitialMigration()
	api.Router()

}
