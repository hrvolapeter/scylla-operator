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

// NewFindConfigInternodeRecvBuffSizeInBytesParams creates a new FindConfigInternodeRecvBuffSizeInBytesParams object
// with the default values initialized.
func NewFindConfigInternodeRecvBuffSizeInBytesParams() *FindConfigInternodeRecvBuffSizeInBytesParams {

	return &FindConfigInternodeRecvBuffSizeInBytesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigInternodeRecvBuffSizeInBytesParamsWithTimeout creates a new FindConfigInternodeRecvBuffSizeInBytesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigInternodeRecvBuffSizeInBytesParamsWithTimeout(timeout time.Duration) *FindConfigInternodeRecvBuffSizeInBytesParams {

	return &FindConfigInternodeRecvBuffSizeInBytesParams{

		timeout: timeout,
	}
}

// NewFindConfigInternodeRecvBuffSizeInBytesParamsWithContext creates a new FindConfigInternodeRecvBuffSizeInBytesParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigInternodeRecvBuffSizeInBytesParamsWithContext(ctx context.Context) *FindConfigInternodeRecvBuffSizeInBytesParams {

	return &FindConfigInternodeRecvBuffSizeInBytesParams{

		Context: ctx,
	}
}

// NewFindConfigInternodeRecvBuffSizeInBytesParamsWithHTTPClient creates a new FindConfigInternodeRecvBuffSizeInBytesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigInternodeRecvBuffSizeInBytesParamsWithHTTPClient(client *http.Client) *FindConfigInternodeRecvBuffSizeInBytesParams {

	return &FindConfigInternodeRecvBuffSizeInBytesParams{
		HTTPClient: client,
	}
}

/*FindConfigInternodeRecvBuffSizeInBytesParams contains all the parameters to send to the API endpoint
for the find config internode recv buff size in bytes operation typically these are written to a http.Request
*/
type FindConfigInternodeRecvBuffSizeInBytesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config internode recv buff size in bytes params
func (o *FindConfigInternodeRecvBuffSizeInBytesParams) WithTimeout(timeout time.Duration) *FindConfigInternodeRecvBuffSizeInBytesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config internode recv buff size in bytes params
func (o *FindConfigInternodeRecvBuffSizeInBytesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config internode recv buff size in bytes params
func (o *FindConfigInternodeRecvBuffSizeInBytesParams) WithContext(ctx context.Context) *FindConfigInternodeRecvBuffSizeInBytesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config internode recv buff size in bytes params
func (o *FindConfigInternodeRecvBuffSizeInBytesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config internode recv buff size in bytes params
func (o *FindConfigInternodeRecvBuffSizeInBytesParams) WithHTTPClient(client *http.Client) *FindConfigInternodeRecvBuffSizeInBytesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config internode recv buff size in bytes params
func (o *FindConfigInternodeRecvBuffSizeInBytesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigInternodeRecvBuffSizeInBytesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
