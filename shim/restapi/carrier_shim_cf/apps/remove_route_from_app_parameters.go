// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewRemoveRouteFromAppParams creates a new RemoveRouteFromAppParams object
// no default values defined in spec.
func NewRemoveRouteFromAppParams() RemoveRouteFromAppParams {

	return RemoveRouteFromAppParams{}
}

// RemoveRouteFromAppParams contains all the bound params for the remove route from app operation
// typically these are obtained from a http.Request
//
// swagger:parameters removeRouteFromApp
type RemoveRouteFromAppParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The guid parameter is used as a part of the request URL: '/v2/apps/:guid/routes/:route_guid'
	  Required: true
	  In: path
	*/
	GUID string
	/*The route_guid parameter is used as a part of the request URL: '/v2/apps/:guid/routes/:route_guid'
	  Required: true
	  In: path
	*/
	RouteGUID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRemoveRouteFromAppParams() beforehand.
func (o *RemoveRouteFromAppParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rGUID, rhkGUID, _ := route.Params.GetOK("guid")
	if err := o.bindGUID(rGUID, rhkGUID, route.Formats); err != nil {
		res = append(res, err)
	}

	rRouteGUID, rhkRouteGUID, _ := route.Params.GetOK("route_guid")
	if err := o.bindRouteGUID(rRouteGUID, rhkRouteGUID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindGUID binds and validates parameter GUID from path.
func (o *RemoveRouteFromAppParams) bindGUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.GUID = raw

	return nil
}

// bindRouteGUID binds and validates parameter RouteGUID from path.
func (o *RemoveRouteFromAppParams) bindRouteGUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.RouteGUID = raw

	return nil
}
