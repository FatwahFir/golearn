package categoryController

import (
	"net/http"
	"strconv"

	"github.com/FatwahFir/golearn/helpers"
	categoryRequest "github.com/FatwahFir/golearn/model/request/category"
	"github.com/FatwahFir/golearn/model/response"
	categoryService "github.com/FatwahFir/golearn/services/category"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService categoryService.CategoryService
}

func NewCategoryController(service categoryService.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: service,
	}
}

func (controller CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryCreateRequest := categoryRequest.CategoryCreateRequest{}
	helpers.ReadRequest(request, &categoryCreateRequest)

	categoryReponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)

	res := response.WebResponse{
		StatusCode: 200,
		Status:     "OK",
		Data:       categoryReponse,
	}

	helpers.WriteResponse(writer, res)
}

func (controller CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryUpdateRequest := categoryRequest.CategoryUpdateRequest{}
	helpers.ReadRequest(request, &categoryUpdateRequest)

	id, err := strconv.Atoi(params.ByName("categoryId"))
	helpers.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryReponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)

	res := response.WebResponse{
		StatusCode: 200,
		Status:     "OK",
		Data:       categoryReponse,
	}

	helpers.WriteResponse(writer, res)
}

func (controller CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	helpers.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)

	res := response.WebResponse{
		StatusCode: 200,
		Status:     "OK",
	}

	helpers.WriteResponse(writer, res)
}

func (controller CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("categoryId"))
	helpers.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)

	res := response.WebResponse{
		StatusCode: 200,
		Status:     "OK",
		Data:       categoryResponse,
	}

	helpers.WriteResponse(writer, res)
}

func (controller CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	catgeoryResponses := controller.CategoryService.FindAll(request.Context())

	res := response.WebResponse{
		StatusCode: 200,
		Status:     "OK",
		Data:       catgeoryResponses,
	}

	helpers.WriteResponse(writer, res)
}
