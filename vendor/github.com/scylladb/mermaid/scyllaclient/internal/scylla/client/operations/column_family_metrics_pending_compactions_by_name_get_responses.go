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

// ColumnFamilyMetricsPendingCompactionsByNameGetReader is a Reader for the ColumnFamilyMetricsPendingCompactionsByNameGet structure.
type ColumnFamilyMetricsPendingCompactionsByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsPendingCompactionsByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsPendingCompactionsByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsPendingCompactionsByNameGetOK creates a ColumnFamilyMetricsPendingCompactionsByNameGetOK with default headers values
func NewColumnFamilyMetricsPendingCompactionsByNameGetOK() *ColumnFamilyMetricsPendingCompactionsByNameGetOK {
	return &ColumnFamilyMetricsPendingCompactionsByNameGetOK{}
}

/*ColumnFamilyMetricsPendingCompactionsByNameGetOK handles this case with default header values.

ColumnFamilyMetricsPendingCompactionsByNameGetOK column family metrics pending compactions by name get o k
*/
type ColumnFamilyMetricsPendingCompactionsByNameGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsPendingCompactionsByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/pending_compactions/{name}][%d] columnFamilyMetricsPendingCompactionsByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsPendingCompactionsByNameGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsPendingCompactionsByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
