// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RemoveManagerFromOrganizationCreatedCode is the HTTP code returned for type RemoveManagerFromOrganizationCreated
const RemoveManagerFromOrganizationCreatedCode int = 201

/*RemoveManagerFromOrganizationCreated successful response

swagger:response removeManagerFromOrganizationCreated
*/
type RemoveManagerFromOrganizationCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RemoveManagerFromOrganizationResponseResource `json:"body,omitempty"`
}

// NewRemoveManagerFromOrganizationCreated creates RemoveManagerFromOrganizationCreated with default headers values
func NewRemoveManagerFromOrganizationCreated() *RemoveManagerFromOrganizationCreated {

	return &RemoveManagerFromOrganizationCreated{}
}

// WithPayload adds the payload to the remove manager from organization created response
func (o *RemoveManagerFromOrganizationCreated) WithPayload(payload *models.RemoveManagerFromOrganizationResponseResource) *RemoveManagerFromOrganizationCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove manager from organization created response
func (o *RemoveManagerFromOrganizationCreated) SetPayload(payload *models.RemoveManagerFromOrganizationResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveManagerFromOrganizationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
