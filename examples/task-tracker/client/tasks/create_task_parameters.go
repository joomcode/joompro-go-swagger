// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/joomcode/joompro-go-swagger/examples/task-tracker/models"
)

// NewCreateTaskParams creates a new CreateTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTaskParams() *CreateTaskParams {
	return &CreateTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTaskParamsWithTimeout creates a new CreateTaskParams object
// with the ability to set a timeout on a request.
func NewCreateTaskParamsWithTimeout(timeout time.Duration) *CreateTaskParams {
	return &CreateTaskParams{
		timeout: timeout,
	}
}

// NewCreateTaskParamsWithContext creates a new CreateTaskParams object
// with the ability to set a context for a request.
func NewCreateTaskParamsWithContext(ctx context.Context) *CreateTaskParams {
	return &CreateTaskParams{
		Context: ctx,
	}
}

// NewCreateTaskParamsWithHTTPClient creates a new CreateTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTaskParamsWithHTTPClient(client *http.Client) *CreateTaskParams {
	return &CreateTaskParams{
		HTTPClient: client,
	}
}

/*
CreateTaskParams contains all the parameters to send to the API endpoint

	for the create task operation.

	Typically these are written to a http.Request.
*/
type CreateTaskParams struct {

	/* Body.

	   The task to create
	*/
	Body *models.Task

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTaskParams) WithDefaults() *CreateTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create task params
func (o *CreateTaskParams) WithTimeout(timeout time.Duration) *CreateTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create task params
func (o *CreateTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create task params
func (o *CreateTaskParams) WithContext(ctx context.Context) *CreateTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create task params
func (o *CreateTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create task params
func (o *CreateTaskParams) WithHTTPClient(client *http.Client) *CreateTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create task params
func (o *CreateTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create task params
func (o *CreateTaskParams) WithBody(body *models.Task) *CreateTaskParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create task params
func (o *CreateTaskParams) SetBody(body *models.Task) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
