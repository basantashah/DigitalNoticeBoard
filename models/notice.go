package models

import (
	"fmt"
	"time"

	u "github.com/basantashah/DigitalNoticeBoard/utils"

	"github.com/jinzhu/gorm"
)

// ///////////////////// //
// create database	//
// ////////////////////	//
// Notices for creating notice table on db
type Notices struct {
	gorm.Model
	Title string `json:"title"`
	// Date       time.Time `json:"date"`
	Expiry     time.Time `json:"expiry"`
	Subject    string    `json:"subject"`
	Content    string    `json:"content"`
	Department string    `json:"department"`
	Urgent     bool      `json:"urgent"`
	Status     bool      `json:"status"`
	UserID     uint      `json:"user_id"`
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (notice *Notices) Validate() (map[string]interface{}, bool) {

	if notice.Subject == "" {
		return u.Message(false, "Subject information should be on the payload"), false
	}

	if notice.Content == "" {
		return u.Message(false, "Content of notice should be on payload"), false
	}

	if notice.Department == "" {
		return u.Message(false, "Notice from which department should be on payload"), false
	}

	if notice.UserID <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

//
/* Create database, this should be called while running this program */

//
func (notice *Notices) Create() map[string]interface{} {

	if resp, ok := notice.Validate(); !ok {
		return resp
	}

	GetDB().Create(notice).Where("")

	resp := u.Message(true, "success")
	resp["notice"] = notice
	return resp
}

// To-DO

func (notice *Notices) Update() map[string]interface{} {
	// notices := make([]*Notices, 0)
	if resp, ok := notice.Validate(); !ok {
		return resp
	}
	GetDB().Update(notice).Where("")

	resp := u.Message(true, "success")
	resp["notice"] = notice
	return resp
}

func Getnotices(user uint) []*Notices {
	notices := make([]*Notices, 0)
	err := GetDB().Table("notices").Where("status = true").Find(&notices).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return notices
}

func Getyournotices(user uint) []*Notices {
	notices := make([]*Notices, 0)
	err := GetDB().Table("notices").Where("user_id = ?", user).Find(&notices).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return notices
}

// func DeleteNotice(user uint) []*Notices {
// 	notices := make([]*Notices, 0)
// 	err := GetDB().Table("notices").Where("user_id = ?", user).Find(&notices).Error
// 	err :=
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}

// 	return notices
// }
