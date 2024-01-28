package domain

import (
	"errors"
	"time"
)

var (
	ErrProductNotFound = errors.New("ProductNotFound")
)

type Product struct {
	Id         int    `gorm:"primarykey" json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Created_at time.Time
}

type GetPaginationInput struct {
	Page  uint `json:"page"`
	Limit uint `json:"limit"`
}

type UpdateProductInput struct {
	Name  *string `json:"name"`
	Price *int    `json:"price"`
}

type ProductRespository interface {
	Save(product *Product) error
	Get(ID *int) (Product, error)
	Update(ID *int, inp *UpdateProductInput) error
	FindAll(page, limit int)([]*Product,error)
	Remove(ID int) error
}
