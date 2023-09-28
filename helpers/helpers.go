package helpers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/FatwahFir/golearn/model/entity"
	categoryReponse "github.com/FatwahFir/golearn/model/response/category"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		PanicIfError(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}

func ToCategoryResponse(category entity.Category) categoryReponse.CategoryResponse {
	categoryResp := categoryReponse.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}

	return categoryResp
}

func ToCategoryResponses(categories []entity.Category) []categoryReponse.CategoryResponse {

	var categoryResponses []categoryReponse.CategoryResponse

	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}

func ReadRequest(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteResponse(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
