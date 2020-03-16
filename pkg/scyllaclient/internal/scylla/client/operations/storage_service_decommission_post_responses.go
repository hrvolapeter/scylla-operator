// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// StorageServiceDecommissionPostReader is a Reader for the StorageServiceDecommissionPost structure.
type StorageServiceDecommissionPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceDecommissionPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceDecommissionPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceDecommissionPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceDecommissionPostOK creates a StorageServiceDecommissionPostOK with default headers values
func NewStorageServiceDecommissionPostOK() *StorageServiceDecommissionPostOK {
	return &StorageServiceDecommissionPostOK{}
}

/*StorageServiceDecommissionPostOK handles this case with default header values.

StorageServiceDecommissionPostOK storage service decommission post o k
*/
type StorageServiceDecommissionPostOK struct {
}

func (o *StorageServiceDecommissionPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceDecommissionPostDefault creates a StorageServiceDecommissionPostDefault with default headers values
func NewStorageServiceDecommissionPostDefault(code int) *StorageServiceDecommissionPostDefault {
	return &StorageServiceDecommissionPostDefault{
		_statusCode: code,
	}
}

/*StorageServiceDecommissionPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceDecommissionPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service decommission post default response
func (o *StorageServiceDecommissionPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceDecommissionPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceDecommissionPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceDecommissionPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
