package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string
	GithubID uint32
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Vehicle struct {
	gorm.Model

	Vendor    string
	ModelType string
	User      *User `gorm:"ForeignKey:ID,constraint:OnDelete:CASCADE;"`
}

type Part struct {
	gorm.Model

	Name     string
	LifeDays int32
	LifeKM   int32
	Vehicle  *Vehicle `gorm:"ForeignKey:ID,constraint:OnDelete:CASCADE;"`
}

type PartInstance struct {
	gorm.Model

	UsedDays int32
	UsedKM   int32

	Part *Part
	Vehicle
}
