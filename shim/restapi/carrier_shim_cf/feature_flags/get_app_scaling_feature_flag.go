// Code generated by go-swagger; DO NOT EDIT.

package feature_flags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAppScalingFeatureFlagHandlerFunc turns a function with the right signature into a get app scaling feature flag handler
type GetAppScalingFeatureFlagHandlerFunc func(GetAppScalingFeatureFlagParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAppScalingFeatureFlagHandlerFunc) Handle(params GetAppScalingFeatureFlagParams) middleware.Responder {
	return fn(params)
}

// GetAppScalingFeatureFlagHandler interface for that can handle valid get app scaling feature flag params
type GetAppScalingFeatureFlagHandler interface {
	Handle(GetAppScalingFeatureFlagParams) middleware.Responder
}

// NewGetAppScalingFeatureFlag creates a new http.Handler for the get app scaling feature flag operation
func NewGetAppScalingFeatureFlag(ctx *middleware.Context, handler GetAppScalingFeatureFlagHandler) *GetAppScalingFeatureFlag {
	return &GetAppScalingFeatureFlag{Context: ctx, Handler: handler}
}

/*GetAppScalingFeatureFlag swagger:route GET /config/feature_flags/app_scaling featureFlags getAppScalingFeatureFlag

Get the App Scaling feature flag

curl --insecure -i %s/v2/config/feature_flags/app_scaling -X GET -H 'Authorization: %s'

*/
type GetAppScalingFeatureFlag struct {
	Context *middleware.Context
	Handler GetAppScalingFeatureFlagHandler
}

func (o *GetAppScalingFeatureFlag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAppScalingFeatureFlagParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
