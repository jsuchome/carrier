// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// RemoveManagedOrganizationFromUserURL generates an URL for the remove managed organization from user operation
type RemoveManagedOrganizationFromUserURL struct {
	GUID                    string
	ManagedOrganizationGUID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RemoveManagedOrganizationFromUserURL) WithBasePath(bp string) *RemoveManagedOrganizationFromUserURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RemoveManagedOrganizationFromUserURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *RemoveManagedOrganizationFromUserURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/users/{guid}/managed_organizations/{managed_organization_guid}"

	guid := o.GUID
	if guid != "" {
		_path = strings.Replace(_path, "{guid}", guid, -1)
	} else {
		return nil, errors.New("guid is required on RemoveManagedOrganizationFromUserURL")
	}

	managedOrganizationGUID := o.ManagedOrganizationGUID
	if managedOrganizationGUID != "" {
		_path = strings.Replace(_path, "{managed_organization_guid}", managedOrganizationGUID, -1)
	} else {
		return nil, errors.New("managedOrganizationGuid is required on RemoveManagedOrganizationFromUserURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v2"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *RemoveManagedOrganizationFromUserURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *RemoveManagedOrganizationFromUserURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *RemoveManagedOrganizationFromUserURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on RemoveManagedOrganizationFromUserURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on RemoveManagedOrganizationFromUserURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *RemoveManagedOrganizationFromUserURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}