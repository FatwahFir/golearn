package app

import (
	categoryController "github.com/FatwahFir/golearn/controllers/category"
	"github.com/FatwahFir/golearn/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController categoryController.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/category", categoryController.FindAll)
	router.GET("/api/category/:categoryId", categoryController.FindById)
	router.POST("/api/category", categoryController.Create)
	router.PUT("/api/category/:categoryId", categoryController.Update)
	router.DELETE("/api/category/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
