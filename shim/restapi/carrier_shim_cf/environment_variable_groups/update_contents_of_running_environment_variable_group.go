// Code generated by go-swagger; DO NOT EDIT.

package environment_variable_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateContentsOfRunningEnvironmentVariableGroupHandlerFunc turns a function with the right signature into a update contents of running environment variable group handler
type UpdateContentsOfRunningEnvironmentVariableGroupHandlerFunc func(UpdateContentsOfRunningEnvironmentVariableGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateContentsOfRunningEnvironmentVariableGroupHandlerFunc) Handle(params UpdateContentsOfRunningEnvironmentVariableGroupParams) middleware.Responder {
	return fn(params)
}

// UpdateContentsOfRunningEnvironmentVariableGroupHandler interface for that can handle valid update contents of running environment variable group params
type UpdateContentsOfRunningEnvironmentVariableGroupHandler interface {
	Handle(UpdateContentsOfRunningEnvironmentVariableGroupParams) middleware.Responder
}

// NewUpdateContentsOfRunningEnvironmentVariableGroup creates a new http.Handler for the update contents of running environment variable group operation
func NewUpdateContentsOfRunningEnvironmentVariableGroup(ctx *middleware.Context, handler UpdateContentsOfRunningEnvironmentVariableGroupHandler) *UpdateContentsOfRunningEnvironmentVariableGroup {
	return &UpdateContentsOfRunningEnvironmentVariableGroup{Context: ctx, Handler: handler}
}

/*UpdateContentsOfRunningEnvironmentVariableGroup swagger:route PUT /config/environment_variable_groups/running environmentVariableGroups updateContentsOfRunningEnvironmentVariableGroup

Updating the contents of the running environment variable group

curl --insecure -i %s/v2/config/environment_variable_groups/running -X PUT -H 'Authorization: %s' -d '%s'

*/
type UpdateContentsOfRunningEnvironmentVariableGroup struct {
	Context *middleware.Context
	Handler UpdateContentsOfRunningEnvironmentVariableGroupHandler
}

func (o *UpdateContentsOfRunningEnvironmentVariableGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateContentsOfRunningEnvironmentVariableGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
