package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/FatwahFir/golearn/helpers"
	"github.com/FatwahFir/golearn/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	//tanpa dependency injection
	// validator := validator.New()
	// db := app.NewDb()
	// categoryRepository := categoryRepository.NewCategoryRepository()
	// categoryService := categoryService.NewCategoryService(categoryRepository, db, validator)
	// categoryController := categoryController.NewCategoryController(categoryService)

	// router := app.NewRouter(categoryController)
	// authMiddleware := middleware.NewAuthMiddleware(router)

	// server := NewServer(authMiddleware)

	server := InitializeServer()
	err := server.ListenAndServe()

	helpers.PanicIfError(err)

}
