package controllers

import (
	"encoding/json"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	contact := &models.Notices{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserID = user
	resp := contact.Create()
	u.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	id := r.Context().Value("user").(uint)
	data := models.Getnotices(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
