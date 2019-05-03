package library

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail() {
	m := gomail.NewMessage()
	m.SetHeader("From", "bindassbasanta@gmail.com")
	m.SetHeader("To", "bindassbasanta@gmail.com")
	m.SetAddressHeader("Cc", "basanta.shah@islingtoncollege.edu.np", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "basantashah1993@gmail.com", os.Getenv("PASSWORD"))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("mail sent")
}
