package main

import (
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
	router.HandleFunc("/api/user/changepassword", controllers.ChangePassword).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/notice/post", controllers.CreateNotice).Methods("POST")
	router.HandleFunc("/api/notice/update", controllers.UpdateNotice).Methods("POST")
	router.HandleFunc("/api/notice/delete", controllers.DeleteNotice).Methods("POST")
	router.HandleFunc("/api/notice/fetch", controllers.GetNoticeFor).Methods("GET")
	router.HandleFunc("/api/notice/yournotice", controllers.GetYourNoticesOnly).Methods("GET")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware
	router.Use(app.LoggingMiddleware) //log all the post and get with response

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	log.Println("Listening .......")

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*",
		},
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), corsOpts.Handler(router))

}
