package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	u "github.com/basantashah/DigitalNoticeBoard/utils"

	"github.com/basantashah/DigitalNoticeBoard/models"
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

// var DeleteNotice = func(w http.ResponseWriter, r *http.Request) {
// 	id := r.Context().Value("user").(uint)
// 	fmt.Println(id)
// 	data := models.DeleteNotice(id)
// 	resp := u.Message(true, "success")
// 	resp["data"] = data
// 	u.Respond(w, resp)

// }

// only get notice that user posted
var GetYourNoticesOnly = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := models.Getyournotices(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// TOp-DO
var UpdateNotice = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	notice := &models.Notices{}

	err := json.NewDecoder(r.Body).Decode(notice)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body, create notice error in notice controller"))
		return
	}

	notice.UserID = user
	resp := notice.Update()
	u.Respond(w, resp)

}
