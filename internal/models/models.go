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
type User struct {
	ID uint `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
	Permission string `json:"permission"`
	BaseModel
}

type Item struct {
	ID uint `json:"item_id"`
	UnitId uint `json:"unit_id"`
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
	TransactionId uint `json:"transaction_id"`
	ItemId uint `json:"item_id"`
	MutationType string `json:"mutation_type"`
	Quantity int `json:"quantity"`
	BalanceAfter int `json:"balance_after"`
	BaseModel
}

type Transaction struct {
	ID uint `json:"transaction_id"`
	UserId uint `json:"user_id"`
	Invoice string `json:"invoice"`
	TotalPrice float64 `json:"total_price"`
	BaseModel
}

type TransactionDetail struct {
	ID uint `json:"transaction_detail_id"`
	TransactionId uint `json:"transaction_id"`
	ItemId uint `json:"item_id"`
	Quantity int `json:"quantity"`
	SellPrice float64 `json:"sell_price"`
	SubTotal float64 `json:"sub_total"`
	BaseModel
}

type Log struct {
	ID uint `json:"log_id"`
	UserId uint `json:"user_id"`
	TransactionId uint `json:"transaction_id"`
	Action string `json:"action"`
	BaseModel
}
