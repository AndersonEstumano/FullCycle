package service

import (
	"github.com/AndersonEstumano/FullCycle/internal/database"
	"github.com/AndersonEstumano/FullCycle/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(ProductDB database.ProductDB) *ProductService {
	return &ProductService{
		ProductDB: ProductDB,
	}
}

func (ps *ProductService) CreateProduct(product *entity.Product) (*entity.Product, error) {
	_, err := ps.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := ps.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetProductsByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProductsByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}