package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/AndersonEstumano/FullCycle/internal/database"
	"github.com/AndersonEstumano/FullCycle/internal/service"
	"github.com/AndersonEstumano/FullCycle/internal/webserver"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)
func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/fullcycle")
	if err != nil {
		panic(err)
	}

	defer db.Close()

		categoryDB := database.NewCategoryDB(db)
		categoryService := service.NewCategoryService(*categoryDB)

		productDB := database.NewProductDB(db)
		productService := service.NewProductService(*productDB)

		categoryHandler := webserver.NewCategoryHandler(categoryService)
		productHandler := webserver.NewProductHandler(productService)

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

		fmt.Println("Server is running on port :8080")
		http.ListenAndServe(":8080", c)
	}
