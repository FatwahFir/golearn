package categoryService

import (
	"context"
	"database/sql"

	"github.com/FatwahFir/golearn/exception"
	"github.com/FatwahFir/golearn/helpers"
	"github.com/FatwahFir/golearn/model/entity"
	categoryRequest "github.com/FatwahFir/golearn/model/request/category"
	categoryReponse "github.com/FatwahFir/golearn/model/response/category"
	categoryRepository "github.com/FatwahFir/golearn/repository/category"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository categoryRepository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(repository categoryRepository.CategoryRepository, db *sql.DB, validator *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		CategoryRepository: repository,
		DB:                 db,
		Validate:           validator,
	}
}

func (service CategoryServiceImpl) Create(ctx context.Context, request categoryRequest.CategoryCreateRequest) categoryReponse.CategoryResponse {

	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	helpers.PanicIfError(err)

	category := entity.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	// return categoryReponse.CategoryResponse(catgory)
	return helpers.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Update(ctx context.Context, request categoryRequest.CategoryUpdateRequest) categoryReponse.CategoryResponse {

	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	helpers.PanicIfError(err)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	// return categoryReponse.CategoryResponse(catgory)
	return helpers.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	helpers.PanicIfError(err)

	category, err := service.CategoryRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service CategoryServiceImpl) FindById(ctx context.Context, categoryId int) categoryReponse.CategoryResponse {
	tx, err := service.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	helpers.PanicIfError(err)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helpers.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) FindAll(ctx context.Context) []categoryReponse.CategoryResponse {
	tx, err := service.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	helpers.PanicIfError(err)

	categories := service.CategoryRepository.FindAll(ctx, tx)
	helpers.PanicIfError(err)

	return helpers.ToCategoryResponses(categories)

}
