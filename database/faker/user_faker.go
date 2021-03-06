package faker

import (
	"time"

	"github.com/Jehanv60/gotoko/app/models"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Userfaker(db *gorm.DB) *models.User {
	return &models.User{
		ID:            uuid.New().String(),
		FirstName:     faker.FirstName(),
		LastName:      faker.LastName(),
		Email:         faker.Email(),
		Password:      "tester",
		RememberToken: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}
}
