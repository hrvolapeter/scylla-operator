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

// NewColumnFamilyMetricsCompressionRatioByNameGetParams creates a new ColumnFamilyMetricsCompressionRatioByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsCompressionRatioByNameGetParams() *ColumnFamilyMetricsCompressionRatioByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCompressionRatioByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsCompressionRatioByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsCompressionRatioByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsCompressionRatioByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsCompressionRatioByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCompressionRatioByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsCompressionRatioByNameGetParamsWithContext creates a new ColumnFamilyMetricsCompressionRatioByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsCompressionRatioByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsCompressionRatioByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCompressionRatioByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsCompressionRatioByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsCompressionRatioByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsCompressionRatioByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsCompressionRatioByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCompressionRatioByNameGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsCompressionRatioByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics compression ratio by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsCompressionRatioByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics compression ratio by name get params
func (o *ColumnFamilyMetricsCompressionRatioByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsCompressionRatioByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics compression ratio by name get params
func (o *ColumnFamilyMetricsCompressionRatioByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics compression ratio by name get params
func (o *ColumnFamilyMetricsCompressionRatioByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsCompressionRatioByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics compression ratio by name get params
func (o *ColumnFamilyMetricsCompressionRatioByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics compression ratio by name get params
func (o *ColumnFamilyMetricsCompressionRatioByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsCompressionRatioByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics compression ratio by name get params
func (o *ColumnFamilyMetricsCompressionRatioByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics compression ratio by name get params
func (o *ColumnFamilyMetricsCompressionRatioByNameGetParams) WithName(name string) *ColumnFamilyMetricsCompressionRatioByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics compression ratio by name get params
func (o *ColumnFamilyMetricsCompressionRatioByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsCompressionRatioByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
