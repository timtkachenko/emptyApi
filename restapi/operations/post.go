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

// PostHandlerFunc turns a function with the right signature into a post handler
type PostHandlerFunc func(PostParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostHandlerFunc) Handle(params PostParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostHandler interface for that can handle valid post params
type PostHandler interface {
	Handle(PostParams, interface{}) middleware.Responder
}

// NewPost creates a new http.Handler for the post operation
func NewPost(ctx *middleware.Context, handler PostHandler) *Post {
	return &Post{Context: ctx, Handler: handler}
}

/*Post swagger:route POST / post

test

*/
type Post struct {
	Context *middleware.Context
	Handler PostHandler
}

func (o *Post) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostParams()

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

// PostOKBody post o k body
// swagger:model PostOKBody
type PostOKBody struct {

	// code
	Code models.SuccessCode `json:"code,omitempty"`

	// data
	Data *models.Test `json:"data,omitempty"`
}

// Validate validates this post o k body
func (o *PostOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOKBody) UnmarshalBinary(b []byte) error {
	var res PostOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
