package main

import (
	"ecovantory/internal/database"
	"ecovantory/internal/repository"
	"fmt"
	"log"
)

func main() {
	fmt.Println("=== EcoVantory System Starting ===")

	database.InitDB()

	category, err := repository.GetCategoryById(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Kategori: ", &category)
}