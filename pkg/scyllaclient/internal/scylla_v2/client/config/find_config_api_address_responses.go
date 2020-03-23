// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigAPIAddressReader is a Reader for the FindConfigAPIAddress structure.
type FindConfigAPIAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigAPIAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigAPIAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigAPIAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigAPIAddressOK creates a FindConfigAPIAddressOK with default headers values
func NewFindConfigAPIAddressOK() *FindConfigAPIAddressOK {
	return &FindConfigAPIAddressOK{}
}

/*FindConfigAPIAddressOK handles this case with default header values.

Config value
*/
type FindConfigAPIAddressOK struct {
	Payload string
}

func (o *FindConfigAPIAddressOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigAPIAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigAPIAddressDefault creates a FindConfigAPIAddressDefault with default headers values
func NewFindConfigAPIAddressDefault(code int) *FindConfigAPIAddressDefault {
	return &FindConfigAPIAddressDefault{
		_statusCode: code,
	}
}

/*FindConfigAPIAddressDefault handles this case with default header values.

unexpected error
*/
type FindConfigAPIAddressDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config api address default response
func (o *FindConfigAPIAddressDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigAPIAddressDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigAPIAddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigAPIAddressDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}