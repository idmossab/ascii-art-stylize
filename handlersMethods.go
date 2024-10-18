package main

import (
	"net/http"
	"strings"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path == "/" {
		// Display the homepage
		renderTemplate(w, "Home Page", &res)
		res = result{
			Symbol:  "", // Clear previous values
			Banner: "",
			Text: "",
			Err:  "",
		}
	} else {
		errorHandler(w, http.StatusNotFound)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path == "/ascii-art" {
		// Parse the form data
		if err := r.ParseForm(); err != nil {
			setError(w, r, "Failed to parse form data.")
			return
		}
		text := r.FormValue("text")
		banner := r.FormValue("banner")
		if text == "" || banner == "" { // Set error message if any field is empty
			// setError(w, r, "Text or Banner cannot be empty")
			errorHandler(w, http.StatusBadRequest)
			return
		}
		if len(text) > 700 { // Check if text exceeds 700 characters
			// setError(w, r, "Please enter less than 700 characters.")
			errorHandler(w, http.StatusBadRequest)
			return
		}
			// Validate banner value
			validBanners := map[string]bool{
				"shadow":    true,
				"standard":  true,
				"thinkertoy": true,
			}
			if !validBanners[banner] {
				errorHandler(w, http.StatusBadRequest) // Invalid banner value
				return
			}
	
		// Generate ascii-art
		artResult := artHandler(text, banner)

		// Check if special characters are present
		if artResult == "Special charactere is not allowed." {
			// Set error message for non-printable characters
			setError(w, r, "Please enter printable ASCII characters only.")
			//errorHandler(w, http.StatusBadRequest)
			return
		} else {
			// Process form values and generate ASCII art if valid
			res = result{
				Symbol: "\n" + artResult,
				Err:  "", // Clear error
				Text:  text, // Store the text
				Banner: banner, // Store the selected banner
			}
		}

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		// 404 Error for incorrect path
		errorHandler(w, http.StatusNotFound)
	}
}

// Function to handle requests to the /Js/ path
func jsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	// Check the path
	if strings.HasPrefix(r.URL.Path, "/assets/js/") {
		// If the request is directly to /Js/, return a forbidden error
		if r.URL.Path == "/assets/js/" {
			errorHandler(w, http.StatusForbidden)
			return
		}
		/*if r.URL.Path == "/Js/script.js" {
			errorHandler(w, http.StatusForbidden)
			return
		}*/
		// If the request is for a specific file, pass the request to http.FileServer
		http.StripPrefix("/assets/js/", http.FileServer(http.Dir("assets/js"))).ServeHTTP(w, r)
	} else {
		// If the path is incorrect, return a 404 error
		errorHandler(w, http.StatusNotFound)
	}
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	// Check the path
	if strings.HasPrefix(r.URL.Path, "/assets/css/") {
		// If the request is directly to /assets/css/, return a forbidden error
		if r.URL.Path == "/assets/css/" {
			errorHandler(w, http.StatusForbidden)
			return
		}
		// If the request is for a specific file, pass the request to http.FileServer
		http.StripPrefix("/assets/css/", http.FileServer(http.Dir("assets/css"))).ServeHTTP(w, r)
	} else {
		// If the path is incorrect, return a 404 error
		errorHandler(w, http.StatusNotFound)
	}
}

