package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/AndersonEstumano/FullCycle/internal/webserver"
)

func NewRouter(categoryHandler *webserver.CategoryHandler, productHandler *webserver.ProductHandler) *chi.Mux {
	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Post("/category", categoryHandler.CreateCategory)
	c.Get("/category", categoryHandler.GetCategories)
	c.Get("/category/{id}", categoryHandler.GetCategory)

	c.Post("/products", productHandler.CreateProduct)
	c.Get("/products", productHandler.GetProducts)
	c.Get("/products/{id}", productHandler.GetProduct)
	c.Get("/products/category/{id}", productHandler.GetProductsByCategoryID)

	return c
}
