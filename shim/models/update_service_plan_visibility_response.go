// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateServicePlanVisibilityResponse update service plan visibility response
//
// swagger:model updateServicePlanVisibilityResponse
type UpdateServicePlanVisibilityResponse struct {

	// The organization Guid
	OrganizationGUID string `json:"organization_guid,omitempty"`

	// The organization Url
	OrganizationURL string `json:"organization_url,omitempty"`

	// The service Plan Guid
	ServicePlanGUID string `json:"service_plan_guid,omitempty"`

	// The service Plan Url
	ServicePlanURL string `json:"service_plan_url,omitempty"`
}

// Validate validates this update service plan visibility response
func (m *UpdateServicePlanVisibilityResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateServicePlanVisibilityResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateServicePlanVisibilityResponse) UnmarshalBinary(b []byte) error {
	var res UpdateServicePlanVisibilityResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
