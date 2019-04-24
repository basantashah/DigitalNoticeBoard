package controllers

import (
	"encoding/json"
	"fmt"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
)

var CreateNotice = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	notice := &models.Notices{}

	err := json.NewDecoder(r.Body).Decode(notice)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body, create notice error in notice controller"))
		return
	}

	notice.UserID = user
	resp := notice.Create()
	u.Respond(w, resp)
}

var GetNoticeFor = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	fmt.Println(id)
	data := models.Getnotices(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// only get notice that user posted
var GetYourNoticesOnly = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := models.Getyournotices(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// TOp-DO
var DeleteNotice = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	// id := r.Context().Value("id").(uint)
	notice := &models.Notices{}

	err := json.NewDecoder(r.Body).Decode(&notice)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body, create notice error in notice delete"))
		return
	}
	notice.UserID = user
	// notice.ID = id
	fmt.Println("this is the notice id", notice.ID)
	resp := notice.Delete()
	u.Respond(w, resp)

}

// var UpdateNotice = func(w http.ResponseWriter, r *http.Request) {

// 	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
// 	// id := r.Context().Value("id").(uint)
// 	notice := &models.Notices{}

// 	err := json.NewDecoder(r.Body).Decode(&notice)
// 	if err != nil {
// 		u.Respond(w, u.Message(false, "Error while decoding request body, create notice error in notice update"))
// 		return
// 	}
// 	notice.UserID = user
// 	// notice.ID = id
// 	fmt.Println("this is the notice id", notice.ID)
// 	resp := notice.Update()
// 	u.Respond(w, resp)

// }

// TOp-DO

var UpdateNotice = func(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	fmt.Println(userID)
	notice := &models.Notices{}
	err := json.NewDecoder(r.Body).Decode(&notice)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body, create notice error in notice controller"))
		return
	}
	notice.UserID = userID
	resp := notice.Update()
	u.Respond(w, resp)
}
