//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/FatwahFir/golearn/app"
	categoryController "github.com/FatwahFir/golearn/controllers/category"
	"github.com/FatwahFir/golearn/middleware"
	categoryRepository "github.com/FatwahFir/golearn/repository/category"
	categoryService "github.com/FatwahFir/golearn/services/category"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	categoryRepository.NewCategoryRepository,
	wire.Bind(new(categoryRepository.CategoryRepository), new(*categoryRepository.CategoryRepositoryImpl)),
	categoryService.NewCategoryService,
	wire.Bind(new(categoryService.CategoryService), new(*categoryService.CategoryServiceImpl)),
	categoryController.NewCategoryController,
	wire.Bind(new(categoryController.CategoryController), new(*categoryController.CategoryControllerImpl)),
)

func ProvideValidatorOptions() []validator.Option {
	return []validator.Option{}
}

func InitializeServer() *http.Server {

	wire.Build(
		app.NewDb,
		ProvideValidatorOptions,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil

}
