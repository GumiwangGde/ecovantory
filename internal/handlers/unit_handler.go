package handlers

import (
	"ecovantory/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

func UnitHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			units, err := service.GetAllUnitService()
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(units)

		case http.MethodPost:
			var input struct {
				Name string `json:"unit_name"`
			}

			err := json.NewDecoder(r.Body).Decode(&input)
			if err != nil {
				http.Error(w, "invalid json format", http.StatusBadRequest)
			}

			err = service.CreateUnitService(input.Name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"message": "unit created"})

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func UnitDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID must be a number", http.StatusBadRequest)
	}

	switch r.Method {
		case http.MethodGet:
			unit, err := service.GetUnitByIdService(uint(id))
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(unit)

		case http.MethodPut:
			var input struct {
				Name string `json:"unit_name"`
			}

			err := json.NewDecoder(r.Body).Decode(&input)
			if err != nil {
				http.Error(w, "format json invalid", http.StatusBadRequest)
				return
			}

			err = service.UpdateUnitService(uint(id), input.Name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(map[string]string{"message": "unit updated"})

		case http.MethodDelete:
			err := service.DeleteUnitService(uint(id))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "unit deleted"})

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}