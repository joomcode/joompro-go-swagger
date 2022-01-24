// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ssfilatov/go-swagger/examples/composed-auth/models"
)

// GetItemsOKCode is the HTTP code returned for type GetItemsOK
const GetItemsOKCode int = 200

/*GetItemsOK multiple items

swagger:response getItemsOK
*/
type GetItemsOK struct {

	/*
	  In: Body
	*/
	Payload []models.Item `json:"body,omitempty"`
}

// NewGetItemsOK creates GetItemsOK with default headers values
func NewGetItemsOK() *GetItemsOK {

	return &GetItemsOK{}
}

// WithPayload adds the payload to the get items o k response
func (o *GetItemsOK) WithPayload(payload []models.Item) *GetItemsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get items o k response
func (o *GetItemsOK) SetPayload(payload []models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetItemsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]models.Item, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetItemsDefault other error response

swagger:response getItemsDefault
*/
type GetItemsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetItemsDefault creates GetItemsDefault with default headers values
func NewGetItemsDefault(code int) *GetItemsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetItemsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get items default response
func (o *GetItemsDefault) WithStatusCode(code int) *GetItemsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get items default response
func (o *GetItemsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get items default response
func (o *GetItemsDefault) WithPayload(payload *models.Error) *GetItemsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get items default response
func (o *GetItemsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetItemsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
