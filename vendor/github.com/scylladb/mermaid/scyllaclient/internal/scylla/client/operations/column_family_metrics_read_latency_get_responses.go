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

// ColumnFamilyMetricsReadLatencyGetReader is a Reader for the ColumnFamilyMetricsReadLatencyGet structure.
type ColumnFamilyMetricsReadLatencyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsReadLatencyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsReadLatencyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsReadLatencyGetOK creates a ColumnFamilyMetricsReadLatencyGetOK with default headers values
func NewColumnFamilyMetricsReadLatencyGetOK() *ColumnFamilyMetricsReadLatencyGetOK {
	return &ColumnFamilyMetricsReadLatencyGetOK{}
}

/*ColumnFamilyMetricsReadLatencyGetOK handles this case with default header values.

ColumnFamilyMetricsReadLatencyGetOK column family metrics read latency get o k
*/
type ColumnFamilyMetricsReadLatencyGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsReadLatencyGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/read_latency][%d] columnFamilyMetricsReadLatencyGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsReadLatencyGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsReadLatencyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
