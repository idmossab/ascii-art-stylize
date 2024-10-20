package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)
var(
	name string
	email string
	message string
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
		name = r.FormValue("name")
		email = r.FormValue("email")
		message = r.FormValue("message")
		fmt.Println(name)
		// Construct the email body
		body := fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s", name, email, message)

		// Send the email using SMTP
		err := sendEmail("mossabbs888@gmail.com", "MOSsab@06", body)
		if err != nil {
			log.Println("Error sending email:", err)
			http.Error(w, "Failed to send email", http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "Email sent successfully!")
	} else {
		// 404 Error for incorrect path
		errorHandler(w, http.StatusNotFound)
	}
}

// Helper function to send the email
func sendEmail(from, password, body string) error {
	to := []string{"mossabbs888@gmail.com"} // Receiver's email address
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	fmt.Println(body)
	// Construct the message
	msg := []byte("To: " + to[0] + "\r\n" +
		"Subject: Contact Form Submission\r\n" +
		"\r\n" + body + "\r\n")

	// Authentication for Gmail
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send the email
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
}

