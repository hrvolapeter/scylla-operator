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

// FindConfigInternodeSendBuffSizeInBytesReader is a Reader for the FindConfigInternodeSendBuffSizeInBytes structure.
type FindConfigInternodeSendBuffSizeInBytesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigInternodeSendBuffSizeInBytesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigInternodeSendBuffSizeInBytesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigInternodeSendBuffSizeInBytesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigInternodeSendBuffSizeInBytesOK creates a FindConfigInternodeSendBuffSizeInBytesOK with default headers values
func NewFindConfigInternodeSendBuffSizeInBytesOK() *FindConfigInternodeSendBuffSizeInBytesOK {
	return &FindConfigInternodeSendBuffSizeInBytesOK{}
}

/*FindConfigInternodeSendBuffSizeInBytesOK handles this case with default header values.

Config value
*/
type FindConfigInternodeSendBuffSizeInBytesOK struct {
	Payload int64
}

func (o *FindConfigInternodeSendBuffSizeInBytesOK) Error() string {
	return fmt.Sprintf("[GET /config/internode_send_buff_size_in_bytes][%d] findConfigInternodeSendBuffSizeInBytesOK  %+v", 200, o.Payload)
}

func (o *FindConfigInternodeSendBuffSizeInBytesOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigInternodeSendBuffSizeInBytesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigInternodeSendBuffSizeInBytesDefault creates a FindConfigInternodeSendBuffSizeInBytesDefault with default headers values
func NewFindConfigInternodeSendBuffSizeInBytesDefault(code int) *FindConfigInternodeSendBuffSizeInBytesDefault {
	return &FindConfigInternodeSendBuffSizeInBytesDefault{
		_statusCode: code,
	}
}

/*FindConfigInternodeSendBuffSizeInBytesDefault handles this case with default header values.

unexpected error
*/
type FindConfigInternodeSendBuffSizeInBytesDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config internode send buff size in bytes default response
func (o *FindConfigInternodeSendBuffSizeInBytesDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigInternodeSendBuffSizeInBytesDefault) Error() string {
	return fmt.Sprintf("[GET /config/internode_send_buff_size_in_bytes][%d] find_config_internode_send_buff_size_in_bytes default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigInternodeSendBuffSizeInBytesDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigInternodeSendBuffSizeInBytesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
