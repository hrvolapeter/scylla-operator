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

// ColumnFamilyMetricsReadLatencyByNameGetReader is a Reader for the ColumnFamilyMetricsReadLatencyByNameGet structure.
type ColumnFamilyMetricsReadLatencyByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsReadLatencyByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsReadLatencyByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsReadLatencyByNameGetOK creates a ColumnFamilyMetricsReadLatencyByNameGetOK with default headers values
func NewColumnFamilyMetricsReadLatencyByNameGetOK() *ColumnFamilyMetricsReadLatencyByNameGetOK {
	return &ColumnFamilyMetricsReadLatencyByNameGetOK{}
}

/*ColumnFamilyMetricsReadLatencyByNameGetOK handles this case with default header values.

ColumnFamilyMetricsReadLatencyByNameGetOK column family metrics read latency by name get o k
*/
type ColumnFamilyMetricsReadLatencyByNameGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsReadLatencyByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/read_latency/{name}][%d] columnFamilyMetricsReadLatencyByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsReadLatencyByNameGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsReadLatencyByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
