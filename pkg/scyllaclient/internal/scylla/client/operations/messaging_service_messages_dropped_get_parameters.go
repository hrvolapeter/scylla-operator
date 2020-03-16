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

// NewMessagingServiceMessagesDroppedGetParams creates a new MessagingServiceMessagesDroppedGetParams object
// with the default values initialized.
func NewMessagingServiceMessagesDroppedGetParams() *MessagingServiceMessagesDroppedGetParams {

	return &MessagingServiceMessagesDroppedGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMessagingServiceMessagesDroppedGetParamsWithTimeout creates a new MessagingServiceMessagesDroppedGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMessagingServiceMessagesDroppedGetParamsWithTimeout(timeout time.Duration) *MessagingServiceMessagesDroppedGetParams {

	return &MessagingServiceMessagesDroppedGetParams{

		timeout: timeout,
	}
}

// NewMessagingServiceMessagesDroppedGetParamsWithContext creates a new MessagingServiceMessagesDroppedGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewMessagingServiceMessagesDroppedGetParamsWithContext(ctx context.Context) *MessagingServiceMessagesDroppedGetParams {

	return &MessagingServiceMessagesDroppedGetParams{

		Context: ctx,
	}
}

// NewMessagingServiceMessagesDroppedGetParamsWithHTTPClient creates a new MessagingServiceMessagesDroppedGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMessagingServiceMessagesDroppedGetParamsWithHTTPClient(client *http.Client) *MessagingServiceMessagesDroppedGetParams {

	return &MessagingServiceMessagesDroppedGetParams{
		HTTPClient: client,
	}
}

/*MessagingServiceMessagesDroppedGetParams contains all the parameters to send to the API endpoint
for the messaging service messages dropped get operation typically these are written to a http.Request
*/
type MessagingServiceMessagesDroppedGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the messaging service messages dropped get params
func (o *MessagingServiceMessagesDroppedGetParams) WithTimeout(timeout time.Duration) *MessagingServiceMessagesDroppedGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the messaging service messages dropped get params
func (o *MessagingServiceMessagesDroppedGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the messaging service messages dropped get params
func (o *MessagingServiceMessagesDroppedGetParams) WithContext(ctx context.Context) *MessagingServiceMessagesDroppedGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the messaging service messages dropped get params
func (o *MessagingServiceMessagesDroppedGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the messaging service messages dropped get params
func (o *MessagingServiceMessagesDroppedGetParams) WithHTTPClient(client *http.Client) *MessagingServiceMessagesDroppedGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the messaging service messages dropped get params
func (o *MessagingServiceMessagesDroppedGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MessagingServiceMessagesDroppedGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
