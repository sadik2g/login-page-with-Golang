package main

import (
	"fmt"
	"log"
	"net/http"
	"complexloginapp/internal/handlers"
	"complexloginapp/pkg/database"
)

func main() {
	// Initialize the database
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/dashboard", handlers.DashboardHandler)

	port := 8080
	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Server is running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
