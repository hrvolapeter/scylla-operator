// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageProxyMetricsRangeHistogramGetReader is a Reader for the StorageProxyMetricsRangeHistogramGet structure.
type StorageProxyMetricsRangeHistogramGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsRangeHistogramGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsRangeHistogramGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageProxyMetricsRangeHistogramGetOK creates a StorageProxyMetricsRangeHistogramGetOK with default headers values
func NewStorageProxyMetricsRangeHistogramGetOK() *StorageProxyMetricsRangeHistogramGetOK {
	return &StorageProxyMetricsRangeHistogramGetOK{}
}

/*StorageProxyMetricsRangeHistogramGetOK handles this case with default header values.

StorageProxyMetricsRangeHistogramGetOK storage proxy metrics range histogram get o k
*/
type StorageProxyMetricsRangeHistogramGetOK struct {
}

func (o *StorageProxyMetricsRangeHistogramGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_proxy/metrics/range/histogram][%d] storageProxyMetricsRangeHistogramGetOK ", 200)
}

func (o *StorageProxyMetricsRangeHistogramGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
