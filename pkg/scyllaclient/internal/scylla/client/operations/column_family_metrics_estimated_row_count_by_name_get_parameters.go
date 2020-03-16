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

// NewColumnFamilyMetricsEstimatedRowCountByNameGetParams creates a new ColumnFamilyMetricsEstimatedRowCountByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsEstimatedRowCountByNameGetParams() *ColumnFamilyMetricsEstimatedRowCountByNameGetParams {
	var ()
	return &ColumnFamilyMetricsEstimatedRowCountByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsEstimatedRowCountByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsEstimatedRowCountByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsEstimatedRowCountByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsEstimatedRowCountByNameGetParams {
	var ()
	return &ColumnFamilyMetricsEstimatedRowCountByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsEstimatedRowCountByNameGetParamsWithContext creates a new ColumnFamilyMetricsEstimatedRowCountByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsEstimatedRowCountByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsEstimatedRowCountByNameGetParams {
	var ()
	return &ColumnFamilyMetricsEstimatedRowCountByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsEstimatedRowCountByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsEstimatedRowCountByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsEstimatedRowCountByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsEstimatedRowCountByNameGetParams {
	var ()
	return &ColumnFamilyMetricsEstimatedRowCountByNameGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsEstimatedRowCountByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics estimated row count by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsEstimatedRowCountByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics estimated row count by name get params
func (o *ColumnFamilyMetricsEstimatedRowCountByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsEstimatedRowCountByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics estimated row count by name get params
func (o *ColumnFamilyMetricsEstimatedRowCountByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics estimated row count by name get params
func (o *ColumnFamilyMetricsEstimatedRowCountByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsEstimatedRowCountByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics estimated row count by name get params
func (o *ColumnFamilyMetricsEstimatedRowCountByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics estimated row count by name get params
func (o *ColumnFamilyMetricsEstimatedRowCountByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsEstimatedRowCountByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics estimated row count by name get params
func (o *ColumnFamilyMetricsEstimatedRowCountByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics estimated row count by name get params
func (o *ColumnFamilyMetricsEstimatedRowCountByNameGetParams) WithName(name string) *ColumnFamilyMetricsEstimatedRowCountByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics estimated row count by name get params
func (o *ColumnFamilyMetricsEstimatedRowCountByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsEstimatedRowCountByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
