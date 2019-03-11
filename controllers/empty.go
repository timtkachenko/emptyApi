package controllers

import (
	"emptyApi/lib"
	"emptyApi/repositories"
	"emptyApi/restapi/operations"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"
	"log"
)

const key = "test"

type ApiHandler struct {
	*repositories.TestRepository
}

func (ah *ApiHandler) Health(params operations.HealthParams) middleware.Responder {
	return operations.NewHealthOK().WithPayload("")
}

func (ah *ApiHandler) Key(params operations.KeyParams) middleware.Responder {
	tokennew := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": "user",
	})
	tokenString, err := tokennew.SignedString([]byte(key))
	if err != nil {
		return lib.ErrorResponse(err)
	}

	return operations.NewKeyOK().WithPayload(tokenString)
}

func (ah *ApiHandler) Get(params operations.GetParams, principal interface{}) middleware.Responder {
	test, err := ah.FindByID(params.ID);
	if err != nil {
		return lib.ErrorResponse(err)
	}

	return operations.NewGetOK().WithPayload(&operations.GetOKBody{Code: 1, Data: test})
}

func (ah *ApiHandler) Post(params operations.PostParams, principal interface{}) middleware.Responder {
	if err := ah.Store(params.Body); err != nil {
		return lib.ErrorResponse(err)
	}

	return operations.NewPostOK().WithPayload(&operations.PostOKBody{Code: 1, Data: params.Body})
}

func (ah *ApiHandler) Authorization(reqToken string) (interface{}, error) {
	token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err == nil && token.Valid {
		log.Println("valid token")
		return true, nil
	}
	return false, fmt.Errorf("invalid token")
}
