// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageServiceBulkLoadByPathPostReader is a Reader for the StorageServiceBulkLoadByPathPost structure.
type StorageServiceBulkLoadByPathPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceBulkLoadByPathPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceBulkLoadByPathPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceBulkLoadByPathPostOK creates a StorageServiceBulkLoadByPathPostOK with default headers values
func NewStorageServiceBulkLoadByPathPostOK() *StorageServiceBulkLoadByPathPostOK {
	return &StorageServiceBulkLoadByPathPostOK{}
}

/*StorageServiceBulkLoadByPathPostOK handles this case with default header values.

StorageServiceBulkLoadByPathPostOK storage service bulk load by path post o k
*/
type StorageServiceBulkLoadByPathPostOK struct {
}

func (o *StorageServiceBulkLoadByPathPostOK) Error() string {
	return fmt.Sprintf("[POST /storage_service/bulk_load/{path}][%d] storageServiceBulkLoadByPathPostOK ", 200)
}

func (o *StorageServiceBulkLoadByPathPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
