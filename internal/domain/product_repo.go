package domain

import (
	"gorm.io/gorm"
)

type ProductRespository interface {
	Save(product *Product) error
	Get(ID *int) (Product, error)
	Update(ID *int, inp *UpdateProductInput) error
	FindAll(page, limit int) ([]*Product, error)
	Remove(ID int) error
}

type UpdateProductInput struct {
	Name  *string `json:"name"`
	Price *int    `json:"price"`
}

// Product represents a product model
type Product struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null"`
	Price int    `json:"price" gorm:"not null"`
}
