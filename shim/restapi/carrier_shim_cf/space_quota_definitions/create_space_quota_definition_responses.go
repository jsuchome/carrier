// Code generated by go-swagger; DO NOT EDIT.

package space_quota_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// CreateSpaceQuotaDefinitionCreatedCode is the HTTP code returned for type CreateSpaceQuotaDefinitionCreated
const CreateSpaceQuotaDefinitionCreatedCode int = 201

/*CreateSpaceQuotaDefinitionCreated successful response

swagger:response createSpaceQuotaDefinitionCreated
*/
type CreateSpaceQuotaDefinitionCreated struct {

	/*
	  In: Body
	*/
	Payload *models.CreateSpaceQuotaDefinitionResponse `json:"body,omitempty"`
}

// NewCreateSpaceQuotaDefinitionCreated creates CreateSpaceQuotaDefinitionCreated with default headers values
func NewCreateSpaceQuotaDefinitionCreated() *CreateSpaceQuotaDefinitionCreated {

	return &CreateSpaceQuotaDefinitionCreated{}
}

// WithPayload adds the payload to the create space quota definition created response
func (o *CreateSpaceQuotaDefinitionCreated) WithPayload(payload *models.CreateSpaceQuotaDefinitionResponse) *CreateSpaceQuotaDefinitionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create space quota definition created response
func (o *CreateSpaceQuotaDefinitionCreated) SetPayload(payload *models.CreateSpaceQuotaDefinitionResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSpaceQuotaDefinitionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}