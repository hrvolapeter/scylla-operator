// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewStorageProxyMetricsReadUnavailablesGetParams creates a new StorageProxyMetricsReadUnavailablesGetParams object
// with the default values initialized.
func NewStorageProxyMetricsReadUnavailablesGetParams() *StorageProxyMetricsReadUnavailablesGetParams {

	return &StorageProxyMetricsReadUnavailablesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyMetricsReadUnavailablesGetParamsWithTimeout creates a new StorageProxyMetricsReadUnavailablesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyMetricsReadUnavailablesGetParamsWithTimeout(timeout time.Duration) *StorageProxyMetricsReadUnavailablesGetParams {

	return &StorageProxyMetricsReadUnavailablesGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyMetricsReadUnavailablesGetParamsWithContext creates a new StorageProxyMetricsReadUnavailablesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyMetricsReadUnavailablesGetParamsWithContext(ctx context.Context) *StorageProxyMetricsReadUnavailablesGetParams {

	return &StorageProxyMetricsReadUnavailablesGetParams{

		Context: ctx,
	}
}

// NewStorageProxyMetricsReadUnavailablesGetParamsWithHTTPClient creates a new StorageProxyMetricsReadUnavailablesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyMetricsReadUnavailablesGetParamsWithHTTPClient(client *http.Client) *StorageProxyMetricsReadUnavailablesGetParams {

	return &StorageProxyMetricsReadUnavailablesGetParams{
		HTTPClient: client,
	}
}

/*StorageProxyMetricsReadUnavailablesGetParams contains all the parameters to send to the API endpoint
for the storage proxy metrics read unavailables get operation typically these are written to a http.Request
*/
type StorageProxyMetricsReadUnavailablesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy metrics read unavailables get params
func (o *StorageProxyMetricsReadUnavailablesGetParams) WithTimeout(timeout time.Duration) *StorageProxyMetricsReadUnavailablesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy metrics read unavailables get params
func (o *StorageProxyMetricsReadUnavailablesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy metrics read unavailables get params
func (o *StorageProxyMetricsReadUnavailablesGetParams) WithContext(ctx context.Context) *StorageProxyMetricsReadUnavailablesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy metrics read unavailables get params
func (o *StorageProxyMetricsReadUnavailablesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy metrics read unavailables get params
func (o *StorageProxyMetricsReadUnavailablesGetParams) WithHTTPClient(client *http.Client) *StorageProxyMetricsReadUnavailablesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy metrics read unavailables get params
func (o *StorageProxyMetricsReadUnavailablesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyMetricsReadUnavailablesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
