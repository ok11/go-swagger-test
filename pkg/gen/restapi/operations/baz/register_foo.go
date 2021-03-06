// Code generated by go-swagger; DO NOT EDIT.

package baz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// RegisterFooHandlerFunc turns a function with the right signature into a register foo handler
type RegisterFooHandlerFunc func(RegisterFooParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RegisterFooHandlerFunc) Handle(params RegisterFooParams) middleware.Responder {
	return fn(params)
}

// RegisterFooHandler interface for that can handle valid register foo params
type RegisterFooHandler interface {
	Handle(RegisterFooParams) middleware.Responder
}

// NewRegisterFoo creates a new http.Handler for the register foo operation
func NewRegisterFoo(ctx *middleware.Context, handler RegisterFooHandler) *RegisterFoo {
	return &RegisterFoo{Context: ctx, Handler: handler}
}

/*RegisterFoo swagger:route POST /foos baz registerFoo

an operation to create a foo

Registers a new foo

*/
type RegisterFoo struct {
	Context *middleware.Context
	Handler RegisterFooHandler
}

func (o *RegisterFoo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRegisterFooParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
