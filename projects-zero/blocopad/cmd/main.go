package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"com.blocopad/blocopad/internal/db"
	"github.com/gorilla/mux"
)

func main() {
	serverPort := "4440" // how I haven't default I put
	os.Setenv("API_CERT_PATH", "server.crt")
	os.Setenv("API_PK_PATH", "server.key")
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

	certificatePath := ""
	if cPath, hasValue := os.LookupEnv("API_CERT_PATH"); hasValue {
		certificatePath = cPath
	} else {
		log.Panicln("Please create env vars API_CERT_PATH and API_PK_PATH!")
	}

	privateKeyPath := ""
	if pkPath, hasValue := os.LookupEnv("API_PK_PATH"); hasValue {
		privateKeyPath = pkPath
	} else {
		log.Panicln("Please create env vars API_PK_PATH and API_CERT_PATH")
	}

	db.DatabaseUrl = databaseUrl
	db.DatabasePassword = databasePassword

	router := mux.NewRouter()

	router.HandleFunc("/api/note/{id}", ReadNote).Methods("GET")
	router.HandleFunc("/api/note", WriteNote).Methods("POST")
	err := http.ListenAndServeTLS(fmt.Sprintf(":%s", serverPort), certificatePath, privateKeyPath, router)
	fmt.Println(err)

}
