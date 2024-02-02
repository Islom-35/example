package adapters

import (
	"example/internal/domain"

	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRespository {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) Save(product *domain.Product) error {

	err := p.db.Create(&product)

	return err.Error
}

func (p *productRepo) Get(ID *int) (domain.Product, error) {
	var product *domain.Product
	result := p.db.First(&product, &ID)

	return *product, result.Error
}

func (p *productRepo) Update(ID *int, inp *domain.UpdateProductInput) error {
	product, err := p.Get(ID)
	if err != nil {
		return err
	}	

	product.Name = *inp.Name
	product.Price = *inp.Price

	result := p.db.Save(&product)

	return result.Error
}

func (p *productRepo) FindAll(page, limit int) ([]*domain.Product, error) {
	var products []*domain.Product

	offset := (page - 1) * limit
	result := p.db.Order("id asc").Limit(limit).Offset(offset).Find(&products)
	return products, result.Error
}

func (p *productRepo) Remove(ID int) error {
	post, err := p.Get(&ID)
	if err != nil {
		return err
	}
	result := p.db.Delete(&post)

	return result.Error
}
