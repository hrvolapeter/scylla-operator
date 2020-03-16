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

// FindConfigShutdownAnnounceInMsReader is a Reader for the FindConfigShutdownAnnounceInMs structure.
type FindConfigShutdownAnnounceInMsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigShutdownAnnounceInMsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigShutdownAnnounceInMsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigShutdownAnnounceInMsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigShutdownAnnounceInMsOK creates a FindConfigShutdownAnnounceInMsOK with default headers values
func NewFindConfigShutdownAnnounceInMsOK() *FindConfigShutdownAnnounceInMsOK {
	return &FindConfigShutdownAnnounceInMsOK{}
}

/*FindConfigShutdownAnnounceInMsOK handles this case with default header values.

Config value
*/
type FindConfigShutdownAnnounceInMsOK struct {
	Payload int64
}

func (o *FindConfigShutdownAnnounceInMsOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigShutdownAnnounceInMsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigShutdownAnnounceInMsDefault creates a FindConfigShutdownAnnounceInMsDefault with default headers values
func NewFindConfigShutdownAnnounceInMsDefault(code int) *FindConfigShutdownAnnounceInMsDefault {
	return &FindConfigShutdownAnnounceInMsDefault{
		_statusCode: code,
	}
}

/*FindConfigShutdownAnnounceInMsDefault handles this case with default header values.

unexpected error
*/
type FindConfigShutdownAnnounceInMsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config shutdown announce in ms default response
func (o *FindConfigShutdownAnnounceInMsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigShutdownAnnounceInMsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigShutdownAnnounceInMsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigShutdownAnnounceInMsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
