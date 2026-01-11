package main

import (
	"ecovantory/internal/database"
	"ecovantory/internal/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("=== EcoVantory System Starting ===")

	database.InitDB()

	router := routes.SetupRoutes()

	port := ":8080"
	fmt.Printf("The server is running on http://localhost%s\n", port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("Failed running server: ", err)
	}
}