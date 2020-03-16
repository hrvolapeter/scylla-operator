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

// NewFindConfigCommitlogSyncParams creates a new FindConfigCommitlogSyncParams object
// with the default values initialized.
func NewFindConfigCommitlogSyncParams() *FindConfigCommitlogSyncParams {

	return &FindConfigCommitlogSyncParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigCommitlogSyncParamsWithTimeout creates a new FindConfigCommitlogSyncParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigCommitlogSyncParamsWithTimeout(timeout time.Duration) *FindConfigCommitlogSyncParams {

	return &FindConfigCommitlogSyncParams{

		timeout: timeout,
	}
}

// NewFindConfigCommitlogSyncParamsWithContext creates a new FindConfigCommitlogSyncParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigCommitlogSyncParamsWithContext(ctx context.Context) *FindConfigCommitlogSyncParams {

	return &FindConfigCommitlogSyncParams{

		Context: ctx,
	}
}

// NewFindConfigCommitlogSyncParamsWithHTTPClient creates a new FindConfigCommitlogSyncParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigCommitlogSyncParamsWithHTTPClient(client *http.Client) *FindConfigCommitlogSyncParams {

	return &FindConfigCommitlogSyncParams{
		HTTPClient: client,
	}
}

/*FindConfigCommitlogSyncParams contains all the parameters to send to the API endpoint
for the find config commitlog sync operation typically these are written to a http.Request
*/
type FindConfigCommitlogSyncParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config commitlog sync params
func (o *FindConfigCommitlogSyncParams) WithTimeout(timeout time.Duration) *FindConfigCommitlogSyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config commitlog sync params
func (o *FindConfigCommitlogSyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config commitlog sync params
func (o *FindConfigCommitlogSyncParams) WithContext(ctx context.Context) *FindConfigCommitlogSyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config commitlog sync params
func (o *FindConfigCommitlogSyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config commitlog sync params
func (o *FindConfigCommitlogSyncParams) WithHTTPClient(client *http.Client) *FindConfigCommitlogSyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config commitlog sync params
func (o *FindConfigCommitlogSyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigCommitlogSyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
