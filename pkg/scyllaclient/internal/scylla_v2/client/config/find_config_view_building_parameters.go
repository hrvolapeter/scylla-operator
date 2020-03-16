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

// NewFindConfigViewBuildingParams creates a new FindConfigViewBuildingParams object
// with the default values initialized.
func NewFindConfigViewBuildingParams() *FindConfigViewBuildingParams {

	return &FindConfigViewBuildingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigViewBuildingParamsWithTimeout creates a new FindConfigViewBuildingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigViewBuildingParamsWithTimeout(timeout time.Duration) *FindConfigViewBuildingParams {

	return &FindConfigViewBuildingParams{

		timeout: timeout,
	}
}

// NewFindConfigViewBuildingParamsWithContext creates a new FindConfigViewBuildingParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigViewBuildingParamsWithContext(ctx context.Context) *FindConfigViewBuildingParams {

	return &FindConfigViewBuildingParams{

		Context: ctx,
	}
}

// NewFindConfigViewBuildingParamsWithHTTPClient creates a new FindConfigViewBuildingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigViewBuildingParamsWithHTTPClient(client *http.Client) *FindConfigViewBuildingParams {

	return &FindConfigViewBuildingParams{
		HTTPClient: client,
	}
}

/*FindConfigViewBuildingParams contains all the parameters to send to the API endpoint
for the find config view building operation typically these are written to a http.Request
*/
type FindConfigViewBuildingParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config view building params
func (o *FindConfigViewBuildingParams) WithTimeout(timeout time.Duration) *FindConfigViewBuildingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config view building params
func (o *FindConfigViewBuildingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config view building params
func (o *FindConfigViewBuildingParams) WithContext(ctx context.Context) *FindConfigViewBuildingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config view building params
func (o *FindConfigViewBuildingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config view building params
func (o *FindConfigViewBuildingParams) WithHTTPClient(client *http.Client) *FindConfigViewBuildingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config view building params
func (o *FindConfigViewBuildingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigViewBuildingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
