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

// StorageServiceSavedCachesLocationGetReader is a Reader for the StorageServiceSavedCachesLocationGet structure.
type StorageServiceSavedCachesLocationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceSavedCachesLocationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceSavedCachesLocationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceSavedCachesLocationGetOK creates a StorageServiceSavedCachesLocationGetOK with default headers values
func NewStorageServiceSavedCachesLocationGetOK() *StorageServiceSavedCachesLocationGetOK {
	return &StorageServiceSavedCachesLocationGetOK{}
}

/*StorageServiceSavedCachesLocationGetOK handles this case with default header values.

StorageServiceSavedCachesLocationGetOK storage service saved caches location get o k
*/
type StorageServiceSavedCachesLocationGetOK struct {
	Payload string
}

func (o *StorageServiceSavedCachesLocationGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_service/saved_caches/location][%d] storageServiceSavedCachesLocationGetOK  %+v", 200, o.Payload)
}

func (o *StorageServiceSavedCachesLocationGetOK) GetPayload() string {
	return o.Payload
}

func (o *StorageServiceSavedCachesLocationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
