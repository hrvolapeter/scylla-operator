// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ColumnFamilyMetricsReadLatencyHistogramByNameGetReader is a Reader for the ColumnFamilyMetricsReadLatencyHistogramByNameGet structure.
type ColumnFamilyMetricsReadLatencyHistogramByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsReadLatencyHistogramByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsReadLatencyHistogramByNameGetOK creates a ColumnFamilyMetricsReadLatencyHistogramByNameGetOK with default headers values
func NewColumnFamilyMetricsReadLatencyHistogramByNameGetOK() *ColumnFamilyMetricsReadLatencyHistogramByNameGetOK {
	return &ColumnFamilyMetricsReadLatencyHistogramByNameGetOK{}
}

/*ColumnFamilyMetricsReadLatencyHistogramByNameGetOK handles this case with default header values.

ColumnFamilyMetricsReadLatencyHistogramByNameGetOK column family metrics read latency histogram by name get o k
*/
type ColumnFamilyMetricsReadLatencyHistogramByNameGetOK struct {
}

func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/read_latency/histogram/{name}][%d] columnFamilyMetricsReadLatencyHistogramByNameGetOK ", 200)
}

func (o *ColumnFamilyMetricsReadLatencyHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
