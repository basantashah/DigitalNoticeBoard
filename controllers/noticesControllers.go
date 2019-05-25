package controllers

import (
	"encoding/json"
	"fmt"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
	"os"

	"gopkg.in/mail.v2"
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

	// To-Do for mail approval

	m := mail.NewMessage()
	m.SetHeader("From", "bindassbasanta@gmail.com")
	m.SetHeader("To", "sity1n117028@islingtoncollege.edu.np")
	m.SetHeader("Subject", "New notice approval!")
	m.SetBody("text/html",
		"Dear"+" "+"<b>"+notice.Department+"</b>"+"<br />"+
			"The new notice has been requested to be posted from your department, Please confirm or suggest changes to the notice by replying to this mail"+"<br />"+
			"the title of  noitce is:"+" "+"<b>"+notice.Title+"</b>"+"<br />"+
			"the subject of  notice is:"+" "+"<b>"+notice.Subject+"</b>"+"<br />"+
			"the department of  noitce is:"+" "+"<b>"+notice.Department+"</b>"+"<br />"+
			"the type of  noitce is:"+" "+"<b>"+notice.Type+"</b>"+"<br />"+
			"<br />"+
			"<b>"+"Thanks and Regards"+"</b>"+
			"<b1>"+"Student Services"+"</b1>")
	dialer := mail.NewPlainDialer("smtp.gmail.com", 587, "bindassbasanta@gmail.com", os.Getenv("PASSWORD"))
	err = dialer.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}
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

//DeleteNotice is used to delete notice but only for given user
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

	// Approval for updating notice should be in library but here for testing purpose
	m := mail.NewMessage()
	m.SetHeader("From", "bindassbasanta@gmail.com")
	m.SetHeader("To", "sity1n117028@islingtoncollege.edu.np")
	m.SetHeader("Subject", "!!ALERT!!! Updating Notice !! ALERT!!")
	m.SetBody("text/html",
		"Dear"+" "+"<b>"+notice.Department+"</b>"+"<br />"+
			"<br />"+
			"There are few changes made for the notice posted for IT department, Please ignore if it was requrested from IT department or respond to this mail if its not"+"<br />"+
			"the updated title of  noitce is:"+" "+"<b>"+notice.Title+"</b>"+"<br />"+
			"the updated subject of  notice is:"+" "+"<b>"+notice.Subject+"</b>"+"<br />"+
			"the updated department of  noitce is:"+" "+"<b>"+notice.Department+"</b>"+"<br />"+
			"the updated type of  noitce is:"+" "+"<b>"+notice.Type+"</b>"+"<br />"+
			"<br />"+
			"<b>"+"Thanks and Regards"+"</b>"+"<br />"+
			"<b1>"+"Student Services"+"</b1>")
	dialer := mail.NewPlainDialer("smtp.gmail.com", 587, "bindassbasanta@gmail.com", os.Getenv("PASSWORD"))
	err = dialer.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}
}
