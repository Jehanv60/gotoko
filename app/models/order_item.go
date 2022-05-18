package models

import (
	"time"
)

type OrderItem struct {
	ID        string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Order     Order
	OrderID   string `gorm:"size:36;index"`
	Product   Product
	ProductID string `gorm:"size:36;index"`
	Qty       int

	Sku  string `gorm:"size:36;index"`
	Name string `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
