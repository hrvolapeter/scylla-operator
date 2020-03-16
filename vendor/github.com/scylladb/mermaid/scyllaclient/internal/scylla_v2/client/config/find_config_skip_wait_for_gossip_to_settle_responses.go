// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/scylla_v2/models"
)

// FindConfigSkipWaitForGossipToSettleReader is a Reader for the FindConfigSkipWaitForGossipToSettle structure.
type FindConfigSkipWaitForGossipToSettleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigSkipWaitForGossipToSettleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigSkipWaitForGossipToSettleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigSkipWaitForGossipToSettleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigSkipWaitForGossipToSettleOK creates a FindConfigSkipWaitForGossipToSettleOK with default headers values
func NewFindConfigSkipWaitForGossipToSettleOK() *FindConfigSkipWaitForGossipToSettleOK {
	return &FindConfigSkipWaitForGossipToSettleOK{}
}

/*FindConfigSkipWaitForGossipToSettleOK handles this case with default header values.

Config value
*/
type FindConfigSkipWaitForGossipToSettleOK struct {
	Payload int64
}

func (o *FindConfigSkipWaitForGossipToSettleOK) Error() string {
	return fmt.Sprintf("[GET /config/skip_wait_for_gossip_to_settle][%d] findConfigSkipWaitForGossipToSettleOK  %+v", 200, o.Payload)
}

func (o *FindConfigSkipWaitForGossipToSettleOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigSkipWaitForGossipToSettleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigSkipWaitForGossipToSettleDefault creates a FindConfigSkipWaitForGossipToSettleDefault with default headers values
func NewFindConfigSkipWaitForGossipToSettleDefault(code int) *FindConfigSkipWaitForGossipToSettleDefault {
	return &FindConfigSkipWaitForGossipToSettleDefault{
		_statusCode: code,
	}
}

/*FindConfigSkipWaitForGossipToSettleDefault handles this case with default header values.

unexpected error
*/
type FindConfigSkipWaitForGossipToSettleDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config skip wait for gossip to settle default response
func (o *FindConfigSkipWaitForGossipToSettleDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigSkipWaitForGossipToSettleDefault) Error() string {
	return fmt.Sprintf("[GET /config/skip_wait_for_gossip_to_settle][%d] find_config_skip_wait_for_gossip_to_settle default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigSkipWaitForGossipToSettleDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigSkipWaitForGossipToSettleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
