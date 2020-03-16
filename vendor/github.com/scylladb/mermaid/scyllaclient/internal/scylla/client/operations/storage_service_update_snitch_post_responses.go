// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageServiceUpdateSnitchPostReader is a Reader for the StorageServiceUpdateSnitchPost structure.
type StorageServiceUpdateSnitchPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceUpdateSnitchPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceUpdateSnitchPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceUpdateSnitchPostOK creates a StorageServiceUpdateSnitchPostOK with default headers values
func NewStorageServiceUpdateSnitchPostOK() *StorageServiceUpdateSnitchPostOK {
	return &StorageServiceUpdateSnitchPostOK{}
}

/*StorageServiceUpdateSnitchPostOK handles this case with default header values.

StorageServiceUpdateSnitchPostOK storage service update snitch post o k
*/
type StorageServiceUpdateSnitchPostOK struct {
}

func (o *StorageServiceUpdateSnitchPostOK) Error() string {
	return fmt.Sprintf("[POST /storage_service/update_snitch][%d] storageServiceUpdateSnitchPostOK ", 200)
}

func (o *StorageServiceUpdateSnitchPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
