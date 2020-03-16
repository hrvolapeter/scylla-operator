// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageProxyCounterWriteRPCTimeoutPostReader is a Reader for the StorageProxyCounterWriteRPCTimeoutPost structure.
type StorageProxyCounterWriteRPCTimeoutPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyCounterWriteRPCTimeoutPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyCounterWriteRPCTimeoutPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageProxyCounterWriteRPCTimeoutPostOK creates a StorageProxyCounterWriteRPCTimeoutPostOK with default headers values
func NewStorageProxyCounterWriteRPCTimeoutPostOK() *StorageProxyCounterWriteRPCTimeoutPostOK {
	return &StorageProxyCounterWriteRPCTimeoutPostOK{}
}

/*StorageProxyCounterWriteRPCTimeoutPostOK handles this case with default header values.

StorageProxyCounterWriteRPCTimeoutPostOK storage proxy counter write Rpc timeout post o k
*/
type StorageProxyCounterWriteRPCTimeoutPostOK struct {
}

func (o *StorageProxyCounterWriteRPCTimeoutPostOK) Error() string {
	return fmt.Sprintf("[POST /storage_proxy/counter_write_rpc_timeout][%d] storageProxyCounterWriteRpcTimeoutPostOK ", 200)
}

func (o *StorageProxyCounterWriteRPCTimeoutPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
