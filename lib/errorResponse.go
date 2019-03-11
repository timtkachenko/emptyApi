package lib

import (
	"emptyApi/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)


func ErrorResponse(err error) middleware.Responder {
	return &ResponderFunc{
		Payload:&models.Error{
			 Message:err.Error(),
		},
	}
}

// ResponderFunc wraps a func as a Responder interface
type ResponderFunc struct {
	Payload *models.Error `json:"body,omitempty"`
}

// WriteResponse writes to the response
func (fn ResponderFunc) WriteResponse(rw http.ResponseWriter, pr runtime.Producer) {
	rw.WriteHeader(http.StatusBadRequest)
	if fn.Payload != nil {
		payload := fn.Payload
		if err := pr.Produce(rw, payload); err != nil {
			panic(err)
		}
	}
}
