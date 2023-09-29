package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/FatwahFir/golearn/app"
	categoryController "github.com/FatwahFir/golearn/controllers/category"
	"github.com/FatwahFir/golearn/helpers"
	"github.com/FatwahFir/golearn/middleware"
	categoryRepository "github.com/FatwahFir/golearn/repository/category"
	categoryService "github.com/FatwahFir/golearn/services/category"
	"github.com/go-playground/validator/v10"
)

func main() {

	validator := validator.New()
	db := app.NewDb()
	categoryRepository := categoryRepository.NewCategoryRepository()
	categoryService := categoryService.NewCategoryService(categoryRepository, db, validator)
	categoryController := categoryController.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()

	helpers.PanicIfError(err)

}
