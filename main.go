package main

import (
	"fmt"
	"go-contacts/app"
	"go-contacts/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// func init() {
// 	gotenv.Load()
// }

// func handler(w http.ResponseWriter, req *http.Request) {
// 	// ...
// 	enableCors(&w)
// 	// ...
// }

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }

// func setupResponse(w *http.ResponseWriter, req *http.Request) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// }

// func indexHandler(w http.ResponseWriter, req *http.Request) {
// 	setupResponse(&w, req)
// 	if (*req).Method == "OPTIONS" {
// 		return
// 	}

// 	// process the request...
// }

func main() {

	router := mux.NewRouter()
	// enableCors(&w)

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/notice/post", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/notice/fetch", controllers.GetContactsFor).Methods("GET") //  user/2/notice

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler
	/* change this port if this port is occupied  */
	//
	handler := cors.Default().Handler(router)
	port := os.Getenv("PORT")
	// if port number is not defined in env variable, then only it will use port defined here
	if port == "" {
		port = "8080"
	}

	/*
		Uncomment this if you want to see port number while running api
	*/
	fmt.Println(port)

	// err := http.ListenAndServe(":"+port, router, handler) //Launch the app, visit localhost:8000/api/*
	// if err != nil {
	// 	fmt.Print(err)
	// }
	http.ListenAndServe(":8080", handler)
}
