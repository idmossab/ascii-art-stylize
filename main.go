package main

import (
	"html/template"
	"net/http"
	"log"
)

type result struct {
	Symbol  string
	Err  string
	Text  string // Store the text value
	Banner string // Store the selected banner value
}
type Email struct{
	MsgValid string
	MsgError string
}
var (
	templates = template.Must(template.ParseGlob("templates/*.html"))
	res       result
	msg Email
)

func main() {
	/*http.HandleFunc("/", getHandler)
	http.HandleFunc("/ascii-art", postHandler)
	http.HandleFunc("/Js/", jsHandler)
    http.HandleFunc("/assets/js/", jsHandler) 
    http.HandleFunc("/assets/css/", cssHandler)     
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)*/

	server := createServer()

    registerRoutes()

    // Start server
    err := server.StartServer() 
     if err != nil {
            // If there's an error, store it and show it in the error page
			log.Fatal("An error occurred while starting the server:", err)
        }
}

func renderTemplate(w http.ResponseWriter, title string, result *result) {
	err := templates.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"Title":  title,
		"Result": result, // Pass the result to the template
		"Msg":msg,
	})
	// Server error
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
	}
}
