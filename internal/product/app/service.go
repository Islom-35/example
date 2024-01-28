package app

import (
	"example/internal/product/domain"
	"log"
	"time"
)

type ProductService interface {
	Create(product domain.Product) error
	Get(ID int) (domain.Product, error)
	Update(ID int, inp domain.UpdateProductInput) error
	FindAll(page, limit int) ([]domain.Product, error)
	Remove(ID int) error
}

func NewProductService(repo domain.ProductRespository) ProductService {
	return &productService{
		repo: repo,
	}
}

type productService struct {
	repo domain.ProductRespository
}

func (p *productService) Create(product domain.Product) error {
	product.Created_at = time.Now()
	if err := p.repo.Save(product); err != nil {
		log.Println(">>>")
		return err
	}
	return nil
}

func (p *productService) Get(ID int) (domain.Product, error) {
	return p.repo.Get(ID)
}

func (p *productService) Update(ID int, inp domain.UpdateProductInput) error {
	return p.repo.Update(ID, inp)
}

func (p *productService) FindAll(page, limit int) ([]domain.Product, error) {
	return p.repo.FindAll(page, limit)
}

func (p *productService) Remove(ID int) error {
	return p.repo.Remove(ID)
}
