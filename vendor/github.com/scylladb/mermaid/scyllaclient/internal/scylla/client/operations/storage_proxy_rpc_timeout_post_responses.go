// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageProxyRPCTimeoutPostReader is a Reader for the StorageProxyRPCTimeoutPost structure.
type StorageProxyRPCTimeoutPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyRPCTimeoutPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyRPCTimeoutPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageProxyRPCTimeoutPostOK creates a StorageProxyRPCTimeoutPostOK with default headers values
func NewStorageProxyRPCTimeoutPostOK() *StorageProxyRPCTimeoutPostOK {
	return &StorageProxyRPCTimeoutPostOK{}
}

/*StorageProxyRPCTimeoutPostOK handles this case with default header values.

StorageProxyRPCTimeoutPostOK storage proxy Rpc timeout post o k
*/
type StorageProxyRPCTimeoutPostOK struct {
}

func (o *StorageProxyRPCTimeoutPostOK) Error() string {
	return fmt.Sprintf("[POST /storage_proxy/rpc_timeout][%d] storageProxyRpcTimeoutPostOK ", 200)
}

func (o *StorageProxyRPCTimeoutPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
