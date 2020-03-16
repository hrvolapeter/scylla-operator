// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageServiceRescheduleFailedDeletionsPostReader is a Reader for the StorageServiceRescheduleFailedDeletionsPost structure.
type StorageServiceRescheduleFailedDeletionsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceRescheduleFailedDeletionsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceRescheduleFailedDeletionsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceRescheduleFailedDeletionsPostOK creates a StorageServiceRescheduleFailedDeletionsPostOK with default headers values
func NewStorageServiceRescheduleFailedDeletionsPostOK() *StorageServiceRescheduleFailedDeletionsPostOK {
	return &StorageServiceRescheduleFailedDeletionsPostOK{}
}

/*StorageServiceRescheduleFailedDeletionsPostOK handles this case with default header values.

StorageServiceRescheduleFailedDeletionsPostOK storage service reschedule failed deletions post o k
*/
type StorageServiceRescheduleFailedDeletionsPostOK struct {
}

func (o *StorageServiceRescheduleFailedDeletionsPostOK) Error() string {
	return fmt.Sprintf("[POST /storage_service/reschedule_failed_deletions][%d] storageServiceRescheduleFailedDeletionsPostOK ", 200)
}

func (o *StorageServiceRescheduleFailedDeletionsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
