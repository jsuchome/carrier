// Code generated by go-swagger; DO NOT EDIT.

package service_plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteServicePlansHandlerFunc turns a function with the right signature into a delete service plans handler
type DeleteServicePlansHandlerFunc func(DeleteServicePlansParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteServicePlansHandlerFunc) Handle(params DeleteServicePlansParams) middleware.Responder {
	return fn(params)
}

// DeleteServicePlansHandler interface for that can handle valid delete service plans params
type DeleteServicePlansHandler interface {
	Handle(DeleteServicePlansParams) middleware.Responder
}

// NewDeleteServicePlans creates a new http.Handler for the delete service plans operation
func NewDeleteServicePlans(ctx *middleware.Context, handler DeleteServicePlansHandler) *DeleteServicePlans {
	return &DeleteServicePlans{Context: ctx, Handler: handler}
}

/*DeleteServicePlans swagger:route DELETE /service_plans/{guid} servicePlans deleteServicePlans

Delete a Particular Service Plans

curl --insecure -i %s/v2/service_plans/{guid} -X DELETE -H 'Authorization: %s'

*/
type DeleteServicePlans struct {
	Context *middleware.Context
	Handler DeleteServicePlansHandler
}

func (o *DeleteServicePlans) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteServicePlansParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
