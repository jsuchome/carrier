// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RemoveAuditedOrganizationFromUserCreatedCode is the HTTP code returned for type RemoveAuditedOrganizationFromUserCreated
const RemoveAuditedOrganizationFromUserCreatedCode int = 201

/*RemoveAuditedOrganizationFromUserCreated successful response

swagger:response removeAuditedOrganizationFromUserCreated
*/
type RemoveAuditedOrganizationFromUserCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RemoveAuditedOrganizationFromUserResponseResource `json:"body,omitempty"`
}

// NewRemoveAuditedOrganizationFromUserCreated creates RemoveAuditedOrganizationFromUserCreated with default headers values
func NewRemoveAuditedOrganizationFromUserCreated() *RemoveAuditedOrganizationFromUserCreated {

	return &RemoveAuditedOrganizationFromUserCreated{}
}

// WithPayload adds the payload to the remove audited organization from user created response
func (o *RemoveAuditedOrganizationFromUserCreated) WithPayload(payload *models.RemoveAuditedOrganizationFromUserResponseResource) *RemoveAuditedOrganizationFromUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove audited organization from user created response
func (o *RemoveAuditedOrganizationFromUserCreated) SetPayload(payload *models.RemoveAuditedOrganizationFromUserResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveAuditedOrganizationFromUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
