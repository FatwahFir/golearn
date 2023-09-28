package categoryService

import (
	"context"

	categoryRequest "github.com/FatwahFir/golearn/model/request/category"
	categoryReponse "github.com/FatwahFir/golearn/model/response/category"
)

type CategoryService interface {
	Create(ctx context.Context, request categoryRequest.CategoryCreateRequest) categoryReponse.CategoryResponse
	Update(ctx context.Context, request categoryRequest.CategoryUpdateRequest) categoryReponse.CategoryResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, categoryId int) categoryReponse.CategoryResponse
	FindAll(ctx context.Context) []categoryReponse.CategoryResponse
}
