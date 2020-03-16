// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageServiceDrainGetReader is a Reader for the StorageServiceDrainGet structure.
type StorageServiceDrainGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceDrainGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceDrainGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceDrainGetOK creates a StorageServiceDrainGetOK with default headers values
func NewStorageServiceDrainGetOK() *StorageServiceDrainGetOK {
	return &StorageServiceDrainGetOK{}
}

/*StorageServiceDrainGetOK handles this case with default header values.

StorageServiceDrainGetOK storage service drain get o k
*/
type StorageServiceDrainGetOK struct {
	Payload string
}

func (o *StorageServiceDrainGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_service/drain][%d] storageServiceDrainGetOK  %+v", 200, o.Payload)
}

func (o *StorageServiceDrainGetOK) GetPayload() string {
	return o.Payload
}

func (o *StorageServiceDrainGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
