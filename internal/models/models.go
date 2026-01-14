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
	ID uint `gorm:"primaryKey" json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
	Permission string `json:"permission"`
	BaseModel
}

type Item struct {
	ID uint `gorm:"primaryKey" json:"item_id"`
	UnitID uint `json:"unit_id"`
	Unit Unit `gorm:"foreignKey:UnitID" json:"unit"`
	ItemCode string `json:"item_code"`
	ItemName string `json:"item_name"`
	Description string `json:"description"`
	BasePrice float64 `json:"base_price"`
	SellPrice float64 `json:"sell_price"`
	Categories []Category `gorm:"many2many:item_categories;" json:"categories"`
	BaseModel
}

type ItemCategory struct {
	ID uint `gorm:"primaryKey" json:"item_category_id"`
	ItemID uint `json:"item_id"`
	CategoryId uint `json:"category_id"`
	BaseModel
}

type Category struct {
	ID uint `gorm:"primaryKey" json:"category_id"`
	CategoryName string `json:"category_name"`
	BaseModel
}

type Unit struct {
	ID uint `gorm:"primaryKey" json:"unit_id"`
	UnitName string `json:"unit_name"`
	BaseModel
}

type Stock struct {
	ID uint `gorm:"primaryKey" json:"stock_id"`
	ItemID uint `json:"item_id"`
	Item Item `gorm:"foreignKey:ItemID" json:"item"`
	PhysicalStatus string `json:"physical_status"`
	PositionStatus int `json:"position_status"`
	Quantity int `json:"quantity"`
	BaseModel
}

type StockMutations struct {
	ID uint `gorm:"primaryKey" json:"stock_mutations_id"`
	TransactionId uint `json:"transaction_id"`
	Transaction Transaction `gorm:"foreignKey:TransactionId" json:"transaction"`
	ItemID uint `json:"item_id"`
	Item Item `gorm:"foreignKey:ItemID" json:"item"`
	MutationType string `json:"mutation_type"`
	Quantity int `json:"quantity"`
	BalanceAfter int `json:"balance_after"`
	BaseModel
}

type Transaction struct {
	ID uint `gorm:"primaryKey" json:"transaction_id"`
	UserID uint `json:"user_id"`
	User User `gorm:"foreignKey:UserID" json:"user"`
	Invoice string `json:"invoice"`
	TotalPrice float64 `json:"total_price"`
	BaseModel
}

type TransactionDetail struct {
	ID uint `gorm:"primaryKey" json:"transaction_detail_id"`
	TransactionId uint `json:"transaction_id"`
	ItemID uint `json:"item_id"`
	Item Item `gorm:"foreignKey:ItemID" json:"item"`
	Quantity int `json:"quantity"`
	SellPrice float64 `json:"sell_price"`
	SubTotal float64 `json:"sub_total"`
	BaseModel
}

type Log struct {
	ID uint `gorm:"primaryKey" json:"log_id"`
	UserID uint `json:"user_id"`
	TransactionId uint `json:"transaction_id"`
	Action string `json:"action"`
	User User `gorm:"foreignKey:UserID" json:"user"`
	BaseModel
}
