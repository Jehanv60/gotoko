package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID            string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	UserID        string `gorm:"size:36;index"`
	User          User
	OrderItems    []OrderItem
	OrderCustomer *OrderCustomer
	Code          string `gorm:"size:50;index"`
	Status        int
	OrderDate     time.Time
	PaymentDue    time.Time
	PaymentStatus string `gorm:"size:50;index"`
	PaymentToken  string `gorm:"size:100;index"`

	Note                string `gorm:"type:text"`
	ShippingCourier     string `gorm:"size:100"`
	ShippingServiceName string `gorm:"size:100"`
	ApprovedBy          string `gorm:"size:36"`
	ApprovedAt          time.Time
	CancelledBy         string `gorm:"size:36"`
	CancelledAt         time.Time
	CancellationNote    string `gorm:"size:255"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}
