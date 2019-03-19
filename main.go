package main

import (
	"go-contacts/app"
	"go-contacts/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/notice/post", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/notice/fetch", controllers.GetContactsFor).Methods("GET")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Listening .......")

	corsObj := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(corsObj)(router)))

}
