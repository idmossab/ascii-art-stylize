package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

// Function to send an email
func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path == "/send-email" {
		// Parse the form data
		if err := r.ParseForm(); err != nil {
			setError(w, r, "Failed to parse form data.")
			return
		}
		// Get form values
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		// Construct the email body
		body := fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s", name, email, message)

		// Send the email using SMTP
		err := sendEmail(body)
		if err != nil {
			log.Println("Error sending email:", err)
			msgErrorEmail(w, r, "Failed to send email")
			return
		}

		msgValidEmail(w, r, "Email sent successfully!")
	} else {
		// 404 Error for incorrect path
		errorHandler(w, http.StatusNotFound)
	}
}


// Helper function to send the email
func sendEmail( body string) error {
	from:="mossabbs888@gmail.com"
	appPassword := "klzo ruxm dmkc bail" // Replace with your generated app password
	to := []string{"mossabbs888@gmail.com"} // Receiver's email address
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	// Construct the message
	msg := []byte("To: " + to[0] + "\r\n" +
		"Subject: Contact Form Submission\r\n" +
		"\r\n" + body + "\r\n")

	// Authentication for Gmail
	auth := smtp.PlainAuth("", from, appPassword, smtpHost)
	// Send the email
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
}

