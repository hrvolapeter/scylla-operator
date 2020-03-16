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

// NewFindConfigConsistentRangemovementParams creates a new FindConfigConsistentRangemovementParams object
// with the default values initialized.
func NewFindConfigConsistentRangemovementParams() *FindConfigConsistentRangemovementParams {

	return &FindConfigConsistentRangemovementParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigConsistentRangemovementParamsWithTimeout creates a new FindConfigConsistentRangemovementParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigConsistentRangemovementParamsWithTimeout(timeout time.Duration) *FindConfigConsistentRangemovementParams {

	return &FindConfigConsistentRangemovementParams{

		timeout: timeout,
	}
}

// NewFindConfigConsistentRangemovementParamsWithContext creates a new FindConfigConsistentRangemovementParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigConsistentRangemovementParamsWithContext(ctx context.Context) *FindConfigConsistentRangemovementParams {

	return &FindConfigConsistentRangemovementParams{

		Context: ctx,
	}
}

// NewFindConfigConsistentRangemovementParamsWithHTTPClient creates a new FindConfigConsistentRangemovementParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigConsistentRangemovementParamsWithHTTPClient(client *http.Client) *FindConfigConsistentRangemovementParams {

	return &FindConfigConsistentRangemovementParams{
		HTTPClient: client,
	}
}

/*FindConfigConsistentRangemovementParams contains all the parameters to send to the API endpoint
for the find config consistent rangemovement operation typically these are written to a http.Request
*/
type FindConfigConsistentRangemovementParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config consistent rangemovement params
func (o *FindConfigConsistentRangemovementParams) WithTimeout(timeout time.Duration) *FindConfigConsistentRangemovementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config consistent rangemovement params
func (o *FindConfigConsistentRangemovementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config consistent rangemovement params
func (o *FindConfigConsistentRangemovementParams) WithContext(ctx context.Context) *FindConfigConsistentRangemovementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config consistent rangemovement params
func (o *FindConfigConsistentRangemovementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config consistent rangemovement params
func (o *FindConfigConsistentRangemovementParams) WithHTTPClient(client *http.Client) *FindConfigConsistentRangemovementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config consistent rangemovement params
func (o *FindConfigConsistentRangemovementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigConsistentRangemovementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
