// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/joomcode/joompro-go-swagger/examples/task-tracker/models"
)

// DeleteTaskReader is a Reader for the DeleteTask structure.
type DeleteTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteTaskNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTaskNoContent creates a DeleteTaskNoContent with default headers values
func NewDeleteTaskNoContent() *DeleteTaskNoContent {
	return &DeleteTaskNoContent{}
}

/*
	DeleteTaskNoContent describes a response with status code 204, with default header values.

Task deleted
*/
type DeleteTaskNoContent struct {
}

// IsSuccess returns true when this delete task no content response returns a 2xx status code
func (o *DeleteTaskNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete task no content response returns a 3xx status code
func (o *DeleteTaskNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete task no content response returns a 4xx status code
func (o *DeleteTaskNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete task no content response returns a 5xx status code
func (o *DeleteTaskNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete task no content response returns a 4xx status code
func (o *DeleteTaskNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteTaskNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tasks/{id}][%d] deleteTaskNoContent ", 204)
}

func (o *DeleteTaskNoContent) String() string {
	return fmt.Sprintf("[DELETE /tasks/{id}][%d] deleteTaskNoContent ", 204)
}

func (o *DeleteTaskNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTaskDefault creates a DeleteTaskDefault with default headers values
func NewDeleteTaskDefault(code int) *DeleteTaskDefault {
	return &DeleteTaskDefault{
		_statusCode: code,
	}
}

/*
	DeleteTaskDefault describes a response with status code -1, with default header values.

Error response
*/
type DeleteTaskDefault struct {
	_statusCode int
	XErrorCode  string

	Payload *models.Error
}

// Code gets the status code for the delete task default response
func (o *DeleteTaskDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete task default response returns a 2xx status code
func (o *DeleteTaskDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete task default response returns a 3xx status code
func (o *DeleteTaskDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete task default response returns a 4xx status code
func (o *DeleteTaskDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete task default response returns a 5xx status code
func (o *DeleteTaskDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete task default response returns a 4xx status code
func (o *DeleteTaskDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteTaskDefault) Error() string {
	return fmt.Sprintf("[DELETE /tasks/{id}][%d] deleteTask default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTaskDefault) String() string {
	return fmt.Sprintf("[DELETE /tasks/{id}][%d] deleteTask default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTaskDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Error-Code
	hdrXErrorCode := response.GetHeader("X-Error-Code")

	if hdrXErrorCode != "" {
		o.XErrorCode = hdrXErrorCode
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
