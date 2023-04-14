package main

import (
	"fmt"
	"net/http"
	"os"

	"com.blocopad/blocopad/internal/db"
	"github.com/gorilla/mux"
)

func main() {
	serverPort := "4440" // how I haven't default I put

	if port, hasValue := os.LookupEnv("API_PORT"); hasValue {
		serverPort = port
	}
	databaseUrl := "localhost:6379"
	if dbUrl, hasValue := os.LookupEnv("API_DB_URL"); hasValue {
		databaseUrl = dbUrl
	}
	databasePassword := ""
	if dbPassword, hasValue := os.LookupEnv("API_DB_PASSWORD"); hasValue {
		databasePassword = dbPassword
	}

	db.DatabaseUrl = databaseUrl
	db.DatabasePassword = databasePassword

	router := mux.NewRouter()

	router.HandleFunc("/api/note/{id}", ReadNote).Methods("GET")
	router.HandleFunc("/api/note", WriteNote).Methods("POST")
	err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), router)
	fmt.Println(err)

}
