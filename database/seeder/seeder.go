package seeder

import (
	"github.com/Jehanv60/gotoko/database/faker"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func Registerseeder(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: faker.Userfaker(db)},
		{Seeder: faker.Productfaker(db)},
	}
}

func DBseed(db *gorm.DB) error {
	for _, seeder := range Registerseeder(db) {
		err := db.Debug().Create(seeder.Seeder).Error
		if err != nil {
			return err
		}
	}
	return nil
}
