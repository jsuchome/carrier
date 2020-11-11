// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAllUsersHandlerFunc turns a function with the right signature into a list all users handler
type ListAllUsersHandlerFunc func(ListAllUsersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAllUsersHandlerFunc) Handle(params ListAllUsersParams) middleware.Responder {
	return fn(params)
}

// ListAllUsersHandler interface for that can handle valid list all users params
type ListAllUsersHandler interface {
	Handle(ListAllUsersParams) middleware.Responder
}

// NewListAllUsers creates a new http.Handler for the list all users operation
func NewListAllUsers(ctx *middleware.Context, handler ListAllUsersHandler) *ListAllUsers {
	return &ListAllUsers{Context: ctx, Handler: handler}
}

/*ListAllUsers swagger:route GET /users users listAllUsers

List all Users

curl --insecure -i %s/v2/users -X GET -H 'Authorization: %s'

*/
type ListAllUsers struct {
	Context *middleware.Context
	Handler ListAllUsersHandler
}

func (o *ListAllUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAllUsersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}