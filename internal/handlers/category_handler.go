package handlers

import (
	"ecovantory/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			categories, err := service.GetAllCategoriesService()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories)

		case http.MethodPost:
			var input struct {
				Name string `json:"name"`
			}

			err := json.NewDecoder(r.Body).Decode(&input)
			if err != nil {
				http.Error(w, "Invalid json format", http.StatusBadRequest)
				return
			}

			err = service.CreateCategoryService(input.Name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"message": "category created"})

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CategoryDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID must be a number", http.StatusBadRequest)
		return
	}

	switch r.Method {
		case http.MethodGet: 
			category, err := service.GetCategoryByIDService(uint(id))
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category)

		case http.MethodPut:
			var input struct {
				Name string `json:"name"`
			}
			err := json.NewDecoder(r.Body).Decode(&input)
			if err != nil {
				http.Error(w, "json format invalid", http.StatusInternalServerError)
				return
			}

			err = service.UpdateCategoryService(uint(id), input.Name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(map[string]string{"message": "category updated"})

		case http.MethodDelete:
			err = service.DeleteCategoryService(uint(id))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "category created"})

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
