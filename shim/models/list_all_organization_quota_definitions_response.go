// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListAllOrganizationQuotaDefinitionsResponse list all organization quota definitions response
//
// swagger:model listAllOrganizationQuotaDefinitionsResponse
type ListAllOrganizationQuotaDefinitionsResponse struct {

	// The instance Memory Limit
	InstanceMemoryLimit int64 `json:"instance_memory_limit,omitempty"`

	// The memory Limit
	MemoryLimit int64 `json:"memory_limit,omitempty"`

	// The name
	Name string `json:"name,omitempty"`

	// The non Basic Services Allowed
	NonBasicServicesAllowed bool `json:"non_basic_services_allowed,omitempty"`

	// The total Routes
	TotalRoutes int64 `json:"total_routes,omitempty"`

	// The total Services
	TotalServices int64 `json:"total_services,omitempty"`

	// The trial Db Allowed
	TrialDbAllowed bool `json:"trial_db_allowed,omitempty"`
}

// Validate validates this list all organization quota definitions response
func (m *ListAllOrganizationQuotaDefinitionsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListAllOrganizationQuotaDefinitionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListAllOrganizationQuotaDefinitionsResponse) UnmarshalBinary(b []byte) error {
	var res ListAllOrganizationQuotaDefinitionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}