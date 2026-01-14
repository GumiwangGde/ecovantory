package handlers

import (
	"ecovantory/internal/models"
	"ecovantory/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

func ItemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			categories, err := service.GetAllItemsService()
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories)

		case http.MethodPost:
			var input struct {
				ItemName string `json:"item_name"`
				ItemCode    string  `json:"item_code"`
				Description string  `json:"description"`
				BasePrice   float64 `json:"base_price"`
				SellPrice   float64 `json:"sell_price"`
				UnitId      uint    `json:"unit_id"`
				CategoryIds []uint  `json:"category_ids"`
			}

			err := json.NewDecoder(r.Body).Decode(&input)
			if err != nil {
				http.Error(w, "input invalid", http.StatusBadRequest)
				return
			}

			newItem := models.Item{
				ItemName:    input.ItemName,
				ItemCode:    input.ItemCode,
				Description: input.Description,
				BasePrice:   input.BasePrice,
				SellPrice:   input.SellPrice,
				UnitID:      input.UnitId,
			}

			err = service.CreateItemService(&newItem, input.CategoryIds)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"message": "item Created"})

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func ItemDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	switch r.Method {
		case http.MethodGet:
			item, err := service.GetItemByIdService(uint(id))
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)

		case http.MethodPut:
			var input struct {
				ItemName    string  `json:"item_name"`
				ItemCode    string  `json:"item_code"`
				Description string  `json:"description"`
				BasePrice   float64 `json:"base_price"`
				SellPrice   float64 `json:"sell_price"`
				UnitId      uint    `json:"unit_id"`
				CategoryIds []uint  `json:"category_ids"`
			}

			err := json.NewDecoder(r.Body).Decode(&input)
			if err != nil {
				http.Error(w, "json format invalid", http.StatusBadRequest)
				return
			}

			updatedData := models.Item{
				ItemName:    input.ItemName,
				ItemCode:    input.ItemCode,
				Description: input.Description,
				BasePrice:   input.BasePrice,
				SellPrice:   input.SellPrice,
				UnitID:      input.UnitId,
			}

			err = service.UpdateItemService(uint(id), &updatedData, input.CategoryIds)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(map[string]string{"message": "item updated"})

		case http.MethodDelete:
			err:= service.DeleteItemService(uint(id))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(map[string]string{"message": "item deleted"})
	}
}