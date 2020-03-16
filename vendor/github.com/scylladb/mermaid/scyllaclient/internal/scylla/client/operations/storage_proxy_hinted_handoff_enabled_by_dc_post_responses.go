// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageProxyHintedHandoffEnabledByDcPostReader is a Reader for the StorageProxyHintedHandoffEnabledByDcPost structure.
type StorageProxyHintedHandoffEnabledByDcPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyHintedHandoffEnabledByDcPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyHintedHandoffEnabledByDcPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageProxyHintedHandoffEnabledByDcPostOK creates a StorageProxyHintedHandoffEnabledByDcPostOK with default headers values
func NewStorageProxyHintedHandoffEnabledByDcPostOK() *StorageProxyHintedHandoffEnabledByDcPostOK {
	return &StorageProxyHintedHandoffEnabledByDcPostOK{}
}

/*StorageProxyHintedHandoffEnabledByDcPostOK handles this case with default header values.

StorageProxyHintedHandoffEnabledByDcPostOK storage proxy hinted handoff enabled by dc post o k
*/
type StorageProxyHintedHandoffEnabledByDcPostOK struct {
}

func (o *StorageProxyHintedHandoffEnabledByDcPostOK) Error() string {
	return fmt.Sprintf("[POST /storage_proxy/hinted_handoff_enabled_by_dc][%d] storageProxyHintedHandoffEnabledByDcPostOK ", 200)
}

func (o *StorageProxyHintedHandoffEnabledByDcPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
