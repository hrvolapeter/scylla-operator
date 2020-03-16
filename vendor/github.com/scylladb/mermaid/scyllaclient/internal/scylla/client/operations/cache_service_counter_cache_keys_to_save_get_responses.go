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

// CacheServiceCounterCacheKeysToSaveGetReader is a Reader for the CacheServiceCounterCacheKeysToSaveGet structure.
type CacheServiceCounterCacheKeysToSaveGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceCounterCacheKeysToSaveGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceCounterCacheKeysToSaveGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCacheServiceCounterCacheKeysToSaveGetOK creates a CacheServiceCounterCacheKeysToSaveGetOK with default headers values
func NewCacheServiceCounterCacheKeysToSaveGetOK() *CacheServiceCounterCacheKeysToSaveGetOK {
	return &CacheServiceCounterCacheKeysToSaveGetOK{}
}

/*CacheServiceCounterCacheKeysToSaveGetOK handles this case with default header values.

CacheServiceCounterCacheKeysToSaveGetOK cache service counter cache keys to save get o k
*/
type CacheServiceCounterCacheKeysToSaveGetOK struct {
	Payload int32
}

func (o *CacheServiceCounterCacheKeysToSaveGetOK) Error() string {
	return fmt.Sprintf("[GET /cache_service/counter_cache_keys_to_save][%d] cacheServiceCounterCacheKeysToSaveGetOK  %+v", 200, o.Payload)
}

func (o *CacheServiceCounterCacheKeysToSaveGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *CacheServiceCounterCacheKeysToSaveGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
