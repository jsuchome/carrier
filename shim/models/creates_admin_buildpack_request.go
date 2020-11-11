// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreatesAdminBuildpackRequest creates admin buildpack request
//
// swagger:model createsAdminBuildpackRequest
type CreatesAdminBuildpackRequest struct {

	// Whether or not the buildpack will be used for staging
	Enabled GenericObject `json:"enabled,omitempty"`

	// The name of the uploaded buildpack file
	Filename GenericObject `json:"filename,omitempty"`

	// Whether or not the buildpack is locked to prevent updates
	Locked GenericObject `json:"locked,omitempty"`

	// The name of the buildpack. To be used by app buildpack field. (only alphanumeric characters)
	Name string `json:"name,omitempty"`

	// The order in which the buildpacks are checked during buildpack auto-detection.
	Position GenericObject `json:"position,omitempty"`
}

// Validate validates this creates admin buildpack request
func (m *CreatesAdminBuildpackRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatesAdminBuildpackRequest) validateEnabled(formats strfmt.Registry) error {

	if swag.IsZero(m.Enabled) { // not required
		return nil
	}

	if err := m.Enabled.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("enabled")
		}
		return err
	}

	return nil
}

func (m *CreatesAdminBuildpackRequest) validateFilename(formats strfmt.Registry) error {

	if swag.IsZero(m.Filename) { // not required
		return nil
	}

	if err := m.Filename.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("filename")
		}
		return err
	}

	return nil
}

func (m *CreatesAdminBuildpackRequest) validateLocked(formats strfmt.Registry) error {

	if swag.IsZero(m.Locked) { // not required
		return nil
	}

	if err := m.Locked.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("locked")
		}
		return err
	}

	return nil
}

func (m *CreatesAdminBuildpackRequest) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := m.Position.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("position")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreatesAdminBuildpackRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatesAdminBuildpackRequest) UnmarshalBinary(b []byte) error {
	var res CreatesAdminBuildpackRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}