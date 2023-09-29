package categoryTest

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/FatwahFir/golearn/app"
	categoryController "github.com/FatwahFir/golearn/controllers/category"
	"github.com/FatwahFir/golearn/helpers"
	"github.com/FatwahFir/golearn/middleware"
	"github.com/FatwahFir/golearn/model/entity"
	categoryRepository "github.com/FatwahFir/golearn/repository/category"
	categoryService "github.com/FatwahFir/golearn/services/category"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go_learn_test")
	helpers.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func truncateDB(db *sql.DB) {
	db.Exec("TRUNCATE categories")
}

func setupRouter(db *sql.DB) http.Handler {
	validator := validator.New()
	categoryRepository := categoryRepository.NewCategoryRepository()
	categoryService := categoryService.NewCategoryService(categoryRepository, db, validator)
	categoryController := categoryController.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCatgeorySuccess(t *testing.T) {

	db := setupTestDB()
	truncateDB(db)
	router := setupRouter(db)

	reqBody := strings.NewReader(`{
		"name": "Gadget"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/category", reqBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "this-api-key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, http.StatusOK, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])

}

func TestCreateCatgeoryFailed(t *testing.T) {

	db := setupTestDB()
	truncateDB(db)
	router := setupRouter(db)

	reqBody := strings.NewReader(`{
		"name": ""
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/category", reqBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "this-api-key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, http.StatusBadRequest, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])

}

func TestUpdateCatgeorySuccess(t *testing.T) {

	db := setupTestDB()
	truncateDB(db)
	router := setupRouter(db)

	tx, _ := db.Begin()
	repository := categoryRepository.NewCategoryRepository()
	category := repository.Save(context.Background(), tx, entity.Category{
		Name: "Not Gadget",
	})
	tx.Commit()

	reqBody := strings.NewReader(`{
		"name": "Gadget"
	}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/category/"+strconv.Itoa(category.Id), reqBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "this-api-key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, http.StatusOK, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))

}

func TestUpdateCatgeoryFailed(t *testing.T) {

	db := setupTestDB()
	truncateDB(db)
	router := setupRouter(db)

	tx, _ := db.Begin()
	repository := categoryRepository.NewCategoryRepository()
	category := repository.Save(context.Background(), tx, entity.Category{
		Name: "Not Gadget",
	})
	tx.Commit()

	reqBody := strings.NewReader(`{
		"name": ""
	}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/category/"+strconv.Itoa(category.Id), reqBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "this-api-key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, http.StatusBadRequest, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])

}

func TestGetByIdCatgeorySuccess(t *testing.T) {

	db := setupTestDB()
	truncateDB(db)
	router := setupRouter(db)

	tx, _ := db.Begin()
	repository := categoryRepository.NewCategoryRepository()
	category := repository.Save(context.Background(), tx, entity.Category{
		Name: "Gadget",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/category/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "this-api-key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, http.StatusOK, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
}

func TestGetByIdCatgeoryNotFound(t *testing.T) {

	db := setupTestDB()
	truncateDB(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/category/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "this-api-key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, http.StatusNotFound, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateDB(db)
	router := setupRouter(db)

	tx, _ := db.Begin()
	repository := categoryRepository.NewCategoryRepository()
	category := repository.Save(context.Background(), tx, entity.Category{
		Name: "Gadget",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/category/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "this-api-key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, http.StatusOK, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateDB(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/category/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "this-api-key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, http.StatusNotFound, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestListCatgeorySuccess(t *testing.T) {

	db := setupTestDB()
	truncateDB(db)
	router := setupRouter(db)

	tx, _ := db.Begin()
	repository := categoryRepository.NewCategoryRepository()
	category := repository.Save(context.Background(), tx, entity.Category{
		Name: "Gadget",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/category", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "this-api-key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, http.StatusOK, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	categories := responseBody["data"].([]interface{})

	categoryres := categories[0].(map[string]interface{})

	assert.Equal(t, category.Id, int(categoryres["id"].(float64)))
	assert.Equal(t, category.Name, categoryres["name"])
}

func TestUnauthorized(t *testing.T) {

	db := setupTestDB()
	truncateDB(db)
	router := setupRouter(db)

	tx, _ := db.Begin()
	repository := categoryRepository.NewCategoryRepository()
	repository.Save(context.Background(), tx, entity.Category{
		Name: "Gadget",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/category", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "false-key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()

	assert.Equal(t, http.StatusUnauthorized, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusUnauthorized, int(responseBody["status_code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
