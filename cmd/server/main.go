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

	categories, err := repository.GetAllCategorY()

	if err != nil {
		log.Fatal("Gagal Mengambil data: ", err)
	}

	fmt.Println("\nDaftar Kategori di Database:")
	fmt.Println("---------------------------------")
	for i, c := range categories {
		fmt.Printf("%d. [ID: %d] Nama Kategori: %s\n", i+1, c.ID, c.CategoryName)
	}
	fmt.Println(("---------------------------------"))
}