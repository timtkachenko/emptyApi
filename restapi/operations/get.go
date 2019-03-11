// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "emptyApi/models"
)

// GetHandlerFunc turns a function with the right signature into a get handler
type GetHandlerFunc func(GetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHandlerFunc) Handle(params GetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetHandler interface for that can handle valid get params
type GetHandler interface {
	Handle(GetParams, interface{}) middleware.Responder
}

// NewGet creates a new http.Handler for the get operation
func NewGet(ctx *middleware.Context, handler GetHandler) *Get {
	return &Get{Context: ctx, Handler: handler}
}

/*Get swagger:route GET /{id} get

test

*/
type Get struct {
	Context *middleware.Context
	Handler GetHandler
}

func (o *Get) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetOKBody get o k body
// swagger:model GetOKBody
type GetOKBody struct {

	// code
	Code models.SuccessCode `json:"code,omitempty"`

	// data
	Data *models.Test `json:"data,omitempty"`
}

// Validate validates this get o k body
func (o *GetOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOKBody) UnmarshalBinary(b []byte) error {
	var res GetOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
