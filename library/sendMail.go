package library

import (
	"fmt"
	"os"

	"gopkg.in/mail.v2"
)

func SendMail() {
	m := mail.NewMessage()
	m.SetHeader("From", "bindassbasanta@gmail.com")
	m.SetHeader("To", "basanta.shah@islingtoncollege.edu.np", "sity1n117028@islingtoncollege.edu.np")
	// m.SetAddressHeader("Cc", "bindassbasanta@gmail.com", "Dan")
	m.SetHeader("Subject", "New notice approval!")
	m.SetBody("text/string", "<pre>resp</pre>")

	dialer := mail.NewPlainDialer("smtp.gmail.com", 587, "bindassbasanta@gmail.com", os.Getenv("PASSWORD")) /* gomail.NewPlainDialer("smtp.gmail.com", 587, "basanta.shah@islingtoncollege.edu.np", "th3Altern@tive") */
	err := dialer.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Email Sent", dialer)
}
