package main

import (
	"fmt"
	"go-contacts/app"
	"go-contacts/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/notice/post", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/notice/fetch", controllers.GetContactsFor).Methods("GET")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	handler := cors.Default().Handler(router)

	port := os.Getenv("PORT")
	// if port number is not defined in env variable, then only it will use port defined here

	if port == "" {
		port = "8080"
	}

	log.Println("Listening .......")

	err := http.ListenAndServe(":"+port, handler) //Launch the app, visit localhost:8000/api/*
	if err != nil {
		fmt.Print(err)
	}

}
