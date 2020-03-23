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

// NewMessagingServiceMessagesPendingGetParams creates a new MessagingServiceMessagesPendingGetParams object
// with the default values initialized.
func NewMessagingServiceMessagesPendingGetParams() *MessagingServiceMessagesPendingGetParams {

	return &MessagingServiceMessagesPendingGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMessagingServiceMessagesPendingGetParamsWithTimeout creates a new MessagingServiceMessagesPendingGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMessagingServiceMessagesPendingGetParamsWithTimeout(timeout time.Duration) *MessagingServiceMessagesPendingGetParams {

	return &MessagingServiceMessagesPendingGetParams{

		timeout: timeout,
	}
}

// NewMessagingServiceMessagesPendingGetParamsWithContext creates a new MessagingServiceMessagesPendingGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewMessagingServiceMessagesPendingGetParamsWithContext(ctx context.Context) *MessagingServiceMessagesPendingGetParams {

	return &MessagingServiceMessagesPendingGetParams{

		Context: ctx,
	}
}

// NewMessagingServiceMessagesPendingGetParamsWithHTTPClient creates a new MessagingServiceMessagesPendingGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMessagingServiceMessagesPendingGetParamsWithHTTPClient(client *http.Client) *MessagingServiceMessagesPendingGetParams {

	return &MessagingServiceMessagesPendingGetParams{
		HTTPClient: client,
	}
}

/*MessagingServiceMessagesPendingGetParams contains all the parameters to send to the API endpoint
for the messaging service messages pending get operation typically these are written to a http.Request
*/
type MessagingServiceMessagesPendingGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the messaging service messages pending get params
func (o *MessagingServiceMessagesPendingGetParams) WithTimeout(timeout time.Duration) *MessagingServiceMessagesPendingGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the messaging service messages pending get params
func (o *MessagingServiceMessagesPendingGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the messaging service messages pending get params
func (o *MessagingServiceMessagesPendingGetParams) WithContext(ctx context.Context) *MessagingServiceMessagesPendingGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the messaging service messages pending get params
func (o *MessagingServiceMessagesPendingGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the messaging service messages pending get params
func (o *MessagingServiceMessagesPendingGetParams) WithHTTPClient(client *http.Client) *MessagingServiceMessagesPendingGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the messaging service messages pending get params
func (o *MessagingServiceMessagesPendingGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MessagingServiceMessagesPendingGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}