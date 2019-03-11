package controllers

import (
	"emptyApi/models"
	"emptyApi/repositories"
	"emptyApi/restapi/operations"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func initDbMock(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf(err.Error())
	}
	gomrDb, err := gorm.Open("postgres", db)
	if err != nil {
		t.Errorf(err.Error())
	}
	return gomrDb, mock
}

func TestApiHandler_Get(t *testing.T) {
	a := assert.New(t)

	db, mock := initDbMock(t)

	apiHandler := &ApiHandler{repositories.NewTestRepository(func() *gorm.DB {
		return db
	})}

	sqlString := `SELECT \* FROM "tests" WHERE "tests"."id" = \$1`
	mock.ExpectQuery(sqlString).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(int64(1), []byte("name")))

	w := httptest.NewRecorder()

	response := apiHandler.Get(operations.GetParams{ID: 1}, true)

	response.WriteResponse(w, runtime.TextProducer())
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	a.NoError(mock.ExpectationsWereMet())

	a.Equal(`{"code":1,"data":{"id":1,"name":"name"}}`, string(body))
}

func TestApiHandler_Post(t *testing.T) {
	a := assert.New(t)

	db, mock := initDbMock(t)

	apiHandler := &ApiHandler{repositories.NewTestRepository(func() *gorm.DB {
		return db
	})}

	sqlString := `INSERT INTO "tests" \("name"\) VALUES \(\$1\) RETURNING "tests"."id"`
	mock.ExpectQuery(sqlString).
		WithArgs("hello test").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(int64(2)))

	w := httptest.NewRecorder()

	response := apiHandler.Post(operations.PostParams{Body: &models.Test{Name: "hello test"}}, true)

	response.WriteResponse(w, runtime.TextProducer())
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	a.NoError(mock.ExpectationsWereMet())

	a.Equal(`{"code":1,"data":{"id":2,"name":"hello test"}}`, string(body))
}
func TestApiHandler_AuthorizationError(t *testing.T) {
	a := assert.New(t)

	apiHandler := &ApiHandler{}

	result, err := apiHandler.Authorization("")
	a.Error(err)
	a.Equal(false, result)
}

func TestApiHandler_Authorization(t *testing.T) {
	a := assert.New(t)

	apiHandler := &ApiHandler{}

	token := GetResponse(apiHandler.Key(operations.KeyParams{}))

	result, err := apiHandler.Authorization(token)
	a.NoError(err)
	a.Equal(true, result)
}

func GetResponse(responder middleware.Responder) string {
	w := httptest.NewRecorder()
	responder.WriteResponse(w, runtime.TextProducer())
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
