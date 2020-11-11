// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAllRoutesHandlerFunc turns a function with the right signature into a list all routes handler
type ListAllRoutesHandlerFunc func(ListAllRoutesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAllRoutesHandlerFunc) Handle(params ListAllRoutesParams) middleware.Responder {
	return fn(params)
}

// ListAllRoutesHandler interface for that can handle valid list all routes params
type ListAllRoutesHandler interface {
	Handle(ListAllRoutesParams) middleware.Responder
}

// NewListAllRoutes creates a new http.Handler for the list all routes operation
func NewListAllRoutes(ctx *middleware.Context, handler ListAllRoutesHandler) *ListAllRoutes {
	return &ListAllRoutes{Context: ctx, Handler: handler}
}

/*ListAllRoutes swagger:route GET /routes routes listAllRoutes

List all Routes

curl --insecure -i %s/v2/routes -X GET -H 'Authorization: %s'

*/
type ListAllRoutes struct {
	Context *middleware.Context
	Handler ListAllRoutesHandler
}

func (o *ListAllRoutes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAllRoutesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}