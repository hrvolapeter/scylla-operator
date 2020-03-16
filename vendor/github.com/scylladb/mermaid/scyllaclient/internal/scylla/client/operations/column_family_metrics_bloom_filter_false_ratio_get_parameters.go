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

// NewColumnFamilyMetricsBloomFilterFalseRatioGetParams creates a new ColumnFamilyMetricsBloomFilterFalseRatioGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsBloomFilterFalseRatioGetParams() *ColumnFamilyMetricsBloomFilterFalseRatioGetParams {

	return &ColumnFamilyMetricsBloomFilterFalseRatioGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsBloomFilterFalseRatioGetParamsWithTimeout creates a new ColumnFamilyMetricsBloomFilterFalseRatioGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsBloomFilterFalseRatioGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsBloomFilterFalseRatioGetParams {

	return &ColumnFamilyMetricsBloomFilterFalseRatioGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsBloomFilterFalseRatioGetParamsWithContext creates a new ColumnFamilyMetricsBloomFilterFalseRatioGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsBloomFilterFalseRatioGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsBloomFilterFalseRatioGetParams {

	return &ColumnFamilyMetricsBloomFilterFalseRatioGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsBloomFilterFalseRatioGetParamsWithHTTPClient creates a new ColumnFamilyMetricsBloomFilterFalseRatioGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsBloomFilterFalseRatioGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsBloomFilterFalseRatioGetParams {

	return &ColumnFamilyMetricsBloomFilterFalseRatioGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsBloomFilterFalseRatioGetParams contains all the parameters to send to the API endpoint
for the column family metrics bloom filter false ratio get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsBloomFilterFalseRatioGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics bloom filter false ratio get params
func (o *ColumnFamilyMetricsBloomFilterFalseRatioGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsBloomFilterFalseRatioGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics bloom filter false ratio get params
func (o *ColumnFamilyMetricsBloomFilterFalseRatioGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics bloom filter false ratio get params
func (o *ColumnFamilyMetricsBloomFilterFalseRatioGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsBloomFilterFalseRatioGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics bloom filter false ratio get params
func (o *ColumnFamilyMetricsBloomFilterFalseRatioGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics bloom filter false ratio get params
func (o *ColumnFamilyMetricsBloomFilterFalseRatioGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsBloomFilterFalseRatioGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics bloom filter false ratio get params
func (o *ColumnFamilyMetricsBloomFilterFalseRatioGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsBloomFilterFalseRatioGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
