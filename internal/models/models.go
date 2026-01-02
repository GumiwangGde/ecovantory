package models

import (
	"time"
)

// BaseModel standarisation for all tabel
type BaseModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// Model for database
type Item struct {
	ID uint `json:"item_id"`
	ItemCode string `json:"item_code"`
	ItemName string `json:"item_name"`
	Description string `json:"description"`
	BasePrice float64 `json:"base_price"`
	SellPrice float64 `json:"sell_price"`
	BaseModel
}

type ItemCategory struct {
	ID uint `json:"item_category_id"`
	ItemId uint `json:"item_id"`
	CategoryId uint `json:"category_id"`
	BaseModel
}

type Category struct {
	ID uint `json:"category_id"`
	CategoryName string `json:"category_name"`
	BaseModel
}

type Unit struct {
	ID uint `json:"unit_id"`
	UnitName string `json:"unit_name"`
	BaseModel
}

type Stock struct {
	ID uint `json:"stock_id"`
	ItemId uint `json:"item_id"`
	PhysicalStatus string `json:"physical_status"`
	PositionStatus int `json:"position_status"`
	Quantity int `json:"quantity"`
	BaseModel
}

type StockMutations struct {
	ID uint `json:"stock_mutations_id"`
	
}

