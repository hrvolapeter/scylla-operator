// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageServiceSlowQueryPostReader is a Reader for the StorageServiceSlowQueryPost structure.
type StorageServiceSlowQueryPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceSlowQueryPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceSlowQueryPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceSlowQueryPostOK creates a StorageServiceSlowQueryPostOK with default headers values
func NewStorageServiceSlowQueryPostOK() *StorageServiceSlowQueryPostOK {
	return &StorageServiceSlowQueryPostOK{}
}

/*StorageServiceSlowQueryPostOK handles this case with default header values.

StorageServiceSlowQueryPostOK storage service slow query post o k
*/
type StorageServiceSlowQueryPostOK struct {
}

func (o *StorageServiceSlowQueryPostOK) Error() string {
	return fmt.Sprintf("[POST /storage_service/slow_query][%d] storageServiceSlowQueryPostOK ", 200)
}

func (o *StorageServiceSlowQueryPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
