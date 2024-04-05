package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/AndersonEstumano/FullCycle/internal/entity"
	"github.com/AndersonEstumano/FullCycle/internal/service"
)

type CategoryHandler struct {
	CategoryService *service.CategoryService
}

func NewCategoryHandler(categoryService *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		CategoryService: categoryService,
	}
}

func (ch *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := ch.CategoryService.CreateCategory(category.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)

}

func (ch *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := ch.CategoryService.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func (ch *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	category, err := ch.CategoryService.GetCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

