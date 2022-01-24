// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ssfilatov/go-swagger/examples/todo-list/models"
)

// DestroyOneReader is a Reader for the DestroyOne structure.
type DestroyOneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DestroyOneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDestroyOneNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDestroyOneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDestroyOneNoContent creates a DestroyOneNoContent with default headers values
func NewDestroyOneNoContent() *DestroyOneNoContent {
	return &DestroyOneNoContent{}
}

/*DestroyOneNoContent handles this case with default header values.

Deleted
*/
type DestroyOneNoContent struct {
}

func (o *DestroyOneNoContent) Error() string {
	return fmt.Sprintf("[DELETE /{id}][%d] destroyOneNoContent ", 204)
}

func (o *DestroyOneNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDestroyOneDefault creates a DestroyOneDefault with default headers values
func NewDestroyOneDefault(code int) *DestroyOneDefault {
	return &DestroyOneDefault{
		_statusCode: code,
	}
}

/*DestroyOneDefault handles this case with default header values.

error
*/
type DestroyOneDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the destroy one default response
func (o *DestroyOneDefault) Code() int {
	return o._statusCode
}

func (o *DestroyOneDefault) Error() string {
	return fmt.Sprintf("[DELETE /{id}][%d] destroyOne default  %+v", o._statusCode, o.Payload)
}

func (o *DestroyOneDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DestroyOneDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
