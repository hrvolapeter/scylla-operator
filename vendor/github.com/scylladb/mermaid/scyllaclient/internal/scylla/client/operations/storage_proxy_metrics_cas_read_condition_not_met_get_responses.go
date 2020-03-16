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

// StorageProxyMetricsCasReadConditionNotMetGetReader is a Reader for the StorageProxyMetricsCasReadConditionNotMetGet structure.
type StorageProxyMetricsCasReadConditionNotMetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsCasReadConditionNotMetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsCasReadConditionNotMetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageProxyMetricsCasReadConditionNotMetGetOK creates a StorageProxyMetricsCasReadConditionNotMetGetOK with default headers values
func NewStorageProxyMetricsCasReadConditionNotMetGetOK() *StorageProxyMetricsCasReadConditionNotMetGetOK {
	return &StorageProxyMetricsCasReadConditionNotMetGetOK{}
}

/*StorageProxyMetricsCasReadConditionNotMetGetOK handles this case with default header values.

StorageProxyMetricsCasReadConditionNotMetGetOK storage proxy metrics cas read condition not met get o k
*/
type StorageProxyMetricsCasReadConditionNotMetGetOK struct {
	Payload int32
}

func (o *StorageProxyMetricsCasReadConditionNotMetGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_proxy/metrics/cas_read/condition_not_met][%d] storageProxyMetricsCasReadConditionNotMetGetOK  %+v", 200, o.Payload)
}

func (o *StorageProxyMetricsCasReadConditionNotMetGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageProxyMetricsCasReadConditionNotMetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
