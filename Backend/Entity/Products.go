package Entity

import(
	"gorm.io/gorm"
)

type Products struct{
	gorm.Model

	NameProduct	string `valid:"required~NameProduct is required"`
	Description string `valid:"required~Description is required"`
	Price       uint `valid:"required~Peice is required"`
	Stock       uint   `valid:"required~Stock is required"`
}