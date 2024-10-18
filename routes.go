package main

import (
    "net/http"
)

// All Routes
func registerRoutes() {
    http.HandleFunc("/", getHandler)            
    http.HandleFunc("/ascii-art", postHandler) 
    http.HandleFunc("/assets/js/", jsHandler)        
}
