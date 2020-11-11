// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewListServiceBrokerDeleteEventsExperimentalParams creates a new ListServiceBrokerDeleteEventsExperimentalParams object
// no default values defined in spec.
func NewListServiceBrokerDeleteEventsExperimentalParams() ListServiceBrokerDeleteEventsExperimentalParams {

	return ListServiceBrokerDeleteEventsExperimentalParams{}
}

// ListServiceBrokerDeleteEventsExperimentalParams contains all the bound params for the list service broker delete events experimental operation
// typically these are obtained from a http.Request
//
// swagger:parameters listServiceBrokerDeleteEventsExperimental
type ListServiceBrokerDeleteEventsExperimentalParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListServiceBrokerDeleteEventsExperimentalParams() beforehand.
func (o *ListServiceBrokerDeleteEventsExperimentalParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}