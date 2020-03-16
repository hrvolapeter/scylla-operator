// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageServiceRebuildPostReader is a Reader for the StorageServiceRebuildPost structure.
type StorageServiceRebuildPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceRebuildPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceRebuildPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceRebuildPostOK creates a StorageServiceRebuildPostOK with default headers values
func NewStorageServiceRebuildPostOK() *StorageServiceRebuildPostOK {
	return &StorageServiceRebuildPostOK{}
}

/*StorageServiceRebuildPostOK handles this case with default header values.

StorageServiceRebuildPostOK storage service rebuild post o k
*/
type StorageServiceRebuildPostOK struct {
}

func (o *StorageServiceRebuildPostOK) Error() string {
	return fmt.Sprintf("[POST /storage_service/rebuild][%d] storageServiceRebuildPostOK ", 200)
}

func (o *StorageServiceRebuildPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
