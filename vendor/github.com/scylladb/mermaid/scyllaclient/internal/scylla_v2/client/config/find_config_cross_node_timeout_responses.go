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

// FindConfigCrossNodeTimeoutReader is a Reader for the FindConfigCrossNodeTimeout structure.
type FindConfigCrossNodeTimeoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCrossNodeTimeoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigCrossNodeTimeoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigCrossNodeTimeoutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCrossNodeTimeoutOK creates a FindConfigCrossNodeTimeoutOK with default headers values
func NewFindConfigCrossNodeTimeoutOK() *FindConfigCrossNodeTimeoutOK {
	return &FindConfigCrossNodeTimeoutOK{}
}

/*FindConfigCrossNodeTimeoutOK handles this case with default header values.

Config value
*/
type FindConfigCrossNodeTimeoutOK struct {
	Payload bool
}

func (o *FindConfigCrossNodeTimeoutOK) Error() string {
	return fmt.Sprintf("[GET /config/cross_node_timeout][%d] findConfigCrossNodeTimeoutOK  %+v", 200, o.Payload)
}

func (o *FindConfigCrossNodeTimeoutOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigCrossNodeTimeoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCrossNodeTimeoutDefault creates a FindConfigCrossNodeTimeoutDefault with default headers values
func NewFindConfigCrossNodeTimeoutDefault(code int) *FindConfigCrossNodeTimeoutDefault {
	return &FindConfigCrossNodeTimeoutDefault{
		_statusCode: code,
	}
}

/*FindConfigCrossNodeTimeoutDefault handles this case with default header values.

unexpected error
*/
type FindConfigCrossNodeTimeoutDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config cross node timeout default response
func (o *FindConfigCrossNodeTimeoutDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCrossNodeTimeoutDefault) Error() string {
	return fmt.Sprintf("[GET /config/cross_node_timeout][%d] find_config_cross_node_timeout default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigCrossNodeTimeoutDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigCrossNodeTimeoutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
