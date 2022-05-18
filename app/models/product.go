package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID               string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	ParentID         string `gorm:"size:36;index"`
	User             User
	UserID           string `gorm:"size:36;index"`
	ProductImages    []ProductImage
	Categories       []Category `gorm:"many2many:product_categories;"`
	Sku              string     `gorm:"size:100;index"`
	Name             string     `gorm:"size:255"`
	Slug             string     `gorm:"size:255"`
	Stock            int
	ShortDescription string `gorm:"type:text"`
	Description      string `gorm:"type:text"`
	Status           int    `gorm:"default:0"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}
