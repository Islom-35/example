package domain

import "time"

type ProductRespository interface {
	Save(product *Product) error
	Get(ID *int) (Product, error)
	Update(ID *int, inp *UpdateProductInput) error
	FindAll(page, limit int) ([]*Product, error)
	Remove(ID int) error
}

type Product struct {
	Id         int    `gorm:"primarykey" json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Created_at time.Time
}
