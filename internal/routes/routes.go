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

	mux.HandleFunc("/units", handlers.UnitHandler)
	mux.HandleFunc("/unit", handlers.UnitDetailHandler)

	return mux
}