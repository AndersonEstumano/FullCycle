package service

import (
	"github.com/AndersonEstumano/FullCycle/internal/database"
	"github.com/AndersonEstumano/FullCycle/internal/entity"
)

type CategoryService struct {
	CategoryDB database.CategoryDB
}

func NewCategoryService(CategoryDB database.CategoryDB) *CategoryService {
	return &CategoryService{
		CategoryDB: CategoryDB,
	}
}

func (cs *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)
	_, err := cs.CategoryDB.CreateCategory(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (cs *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := cs.CategoryDB.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (cs *CategoryService) GetCategory(id string) (*entity.Category, error) {
	category, err := cs.CategoryDB.GetCategory(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}