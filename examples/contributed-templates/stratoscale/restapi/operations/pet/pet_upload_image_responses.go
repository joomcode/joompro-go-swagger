// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ssfilatov/go-swagger/examples/contributed-templates/stratoscale/models"
)

// PetUploadImageOKCode is the HTTP code returned for type PetUploadImageOK
const PetUploadImageOKCode int = 200

/*
PetUploadImageOK successful operation

swagger:response petUploadImageOK
*/
type PetUploadImageOK struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewPetUploadImageOK creates PetUploadImageOK with default headers values
func NewPetUploadImageOK() *PetUploadImageOK {

	return &PetUploadImageOK{}
}

// WithPayload adds the payload to the pet upload image o k response
func (o *PetUploadImageOK) WithPayload(payload *models.APIResponse) *PetUploadImageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pet upload image o k response
func (o *PetUploadImageOK) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PetUploadImageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
