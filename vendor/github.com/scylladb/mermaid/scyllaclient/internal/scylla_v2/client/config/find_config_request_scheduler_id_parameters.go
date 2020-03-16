// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindConfigRequestSchedulerIDParams creates a new FindConfigRequestSchedulerIDParams object
// with the default values initialized.
func NewFindConfigRequestSchedulerIDParams() *FindConfigRequestSchedulerIDParams {

	return &FindConfigRequestSchedulerIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigRequestSchedulerIDParamsWithTimeout creates a new FindConfigRequestSchedulerIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigRequestSchedulerIDParamsWithTimeout(timeout time.Duration) *FindConfigRequestSchedulerIDParams {

	return &FindConfigRequestSchedulerIDParams{

		timeout: timeout,
	}
}

// NewFindConfigRequestSchedulerIDParamsWithContext creates a new FindConfigRequestSchedulerIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigRequestSchedulerIDParamsWithContext(ctx context.Context) *FindConfigRequestSchedulerIDParams {

	return &FindConfigRequestSchedulerIDParams{

		Context: ctx,
	}
}

// NewFindConfigRequestSchedulerIDParamsWithHTTPClient creates a new FindConfigRequestSchedulerIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigRequestSchedulerIDParamsWithHTTPClient(client *http.Client) *FindConfigRequestSchedulerIDParams {

	return &FindConfigRequestSchedulerIDParams{
		HTTPClient: client,
	}
}

/*FindConfigRequestSchedulerIDParams contains all the parameters to send to the API endpoint
for the find config request scheduler id operation typically these are written to a http.Request
*/
type FindConfigRequestSchedulerIDParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config request scheduler id params
func (o *FindConfigRequestSchedulerIDParams) WithTimeout(timeout time.Duration) *FindConfigRequestSchedulerIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config request scheduler id params
func (o *FindConfigRequestSchedulerIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config request scheduler id params
func (o *FindConfigRequestSchedulerIDParams) WithContext(ctx context.Context) *FindConfigRequestSchedulerIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config request scheduler id params
func (o *FindConfigRequestSchedulerIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config request scheduler id params
func (o *FindConfigRequestSchedulerIDParams) WithHTTPClient(client *http.Client) *FindConfigRequestSchedulerIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config request scheduler id params
func (o *FindConfigRequestSchedulerIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigRequestSchedulerIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
