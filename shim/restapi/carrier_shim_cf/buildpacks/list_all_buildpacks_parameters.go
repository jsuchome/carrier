// Code generated by go-swagger; DO NOT EDIT.

package buildpacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewListAllBuildpacksParams creates a new ListAllBuildpacksParams object
// no default values defined in spec.
func NewListAllBuildpacksParams() ListAllBuildpacksParams {

	return ListAllBuildpacksParams{}
}

// ListAllBuildpacksParams contains all the bound params for the list all buildpacks operation
// typically these are obtained from a http.Request
//
// swagger:parameters listAllBuildpacks
type ListAllBuildpacksParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListAllBuildpacksParams() beforehand.
func (o *ListAllBuildpacksParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
