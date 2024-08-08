package main

import (
	"log"
	"net/http"

	"github.com/username/basic_curd_golang/database"

	"github.com/username/basic_curd_golang/handlers"
)

func main() {
	dataSourceName := "host=localhost port=5432 user=postgres password=12345 dbname=golangcurd sslmode=disable"
	database.InitDB(dataSourceName)

	http.HandleFunc("/users", handlers.GetUsers)
	http.HandleFunc("/user", handlers.GetUser)
	http.HandleFunc("/user/create", handlers.CreateUser)
	http.HandleFunc("/user/update", handlers.UpdateUser)
	// http.HandleFunc("/user/delete", handlers.DeleteUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
