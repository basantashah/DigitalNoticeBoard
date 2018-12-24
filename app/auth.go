package app


import (
	"net/http"
	u "lens/utils"
	"strings"
	"go-contacts/models"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"context"
	"fmt"
)

// ////////////////// //
/* JWT authentication */
// ////////////////// //

var JwtAuthentication = fun (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWritter, r http.Request){
		// these are the end points where we don't need access tokens
		authenticationNotRequired := []string{"/api/user/new", "/api/user/login"}

		// The current request path
		requestPath := r.URL.Path 
		// ////////////////////////////////////////////////////////// //
		/* check and serve if any request needs authentication or not */
		// ////////////////////////////////////////////////////////// //

		for _, value := range authenticationNotRequired {
			if value == requestPath {
				next.ServeHTTP(w,r)
				return
			}
		}
		response := make(map[string] interface{})
		tokenHeader := r.Header.Get("Authorization")
		
		// check if token is missing 
		if tokenHeader == "" {
			response = u.Message(false, "Error: missing authentication token")
			w.WriterHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		// The general format of token is , Bearer and then authentication token, we need to check if the received token is in actual format
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriterHeader(http.StatusForbidden)
			w.Header().Add("Content-type", "application/json")
			u.Respond(w, response)
			return
		}
		// Taken the token only and remove the part bearer using split function
		tokenPart := splitted[1]
		actualToken := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, actualToken, func(token *jwt.Token) (interface{}, error){
			return []byte(os.Getenv("token_password")), nil
		})

		// Malformed token
		if err != nil {
			response = u.Message(false, "Malformed authentication token")
			w.WriterHeader(http.StatusForbidden)
			w.Header().Add("Content-type", "application/json")
			u.Respond(w, response)
			return
		}
		// Incorrect token
		if !token.Valid {
			response = u.Message(false, "Invalid token")
			w.WriterHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json" )
			u.Respond(w, response)
			return
		}


		// //////////////////// //
		/* NO ERROR and Proceed */
		// //////////////////// //
		fmt.Sprintf("User %", actualToken)  // Used for monitoring
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}