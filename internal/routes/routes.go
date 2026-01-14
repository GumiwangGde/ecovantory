package routes

import (
	"ecovantory/internal/handlers"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/categories", handlers.CategoriesHandler)
	mux.HandleFunc("/category", handlers.CategoryDetailHandler)

	mux.HandleFunc("/items", handlers.ItemHandler)
	mux.HandleFunc("/item", handlers.ItemDetailHandler)

	return mux
}