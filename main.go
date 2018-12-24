package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/GoldchainCrypto/go-contacts/app"
	"github.com/gorilla/mux"
)

func main() {
	// ///////////////////// //
	// Installing router mux //
	// ///////////////////// //

	router := mux.NewRouter()

	router.Use(app.JwtAuthentication) //Adding JWT for middleware

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
