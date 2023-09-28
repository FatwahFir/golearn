package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/FatwahFir/golearn/app"
	categoryController "github.com/FatwahFir/golearn/controllers/category"
	"github.com/FatwahFir/golearn/exception"
	"github.com/FatwahFir/golearn/helpers"
	categoryRepository "github.com/FatwahFir/golearn/repository/category"
	categoryService "github.com/FatwahFir/golearn/services/category"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	validator := validator.New()
	db := app.NewDb()
	categoryRepository := categoryRepository.NewCategoryRepository()
	categoryService := categoryService.NewCategoryService(categoryRepository, db, validator)
	categoryController := categoryController.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/category", categoryController.FindAll)
	router.GET("/api/category/:categoryId", categoryController.FindById)
	router.POST("/api/category", categoryController.Create)
	router.PUT("/api/category/:categoryId", categoryController.Update)
	router.DELETE("/api/category/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()

	helpers.PanicIfError(err)

}
