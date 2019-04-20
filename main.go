package main

import (
	"fmt"
	"go-contacts/app"
	"go-contacts/controllers"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/notice/post", controllers.CreateNotice).Methods("POST")
	router.HandleFunc("/api/notice/fetch", controllers.GetNoticeFor).Methods("GET")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware
	router.Use(LoggingMiddleware)

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
			"*", //or you can your header key values which you are using in your application

		},
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), corsOpts.Handler(router))
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		fmt.Println()
		log.Println("..............")
		// if os.Getenv("DUMP_REQUEST") == "true" {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		// }
		fmt.Println(string(requestDump))

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
