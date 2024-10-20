package main

import (
    "net/http"
)

// All Routes
func registerRoutes() {
    http.HandleFunc("/", getHandler)            
    http.HandleFunc("/ascii-art", postHandler) 
    http.HandleFunc("/assets/js/", jsHandler) 
    http.HandleFunc("/assets/css/", cssHandler)
    http.HandleFunc("/assets/images/", imgHandler)
    http.HandleFunc("/send-email", sendEmailHandler)        
        
       
}
