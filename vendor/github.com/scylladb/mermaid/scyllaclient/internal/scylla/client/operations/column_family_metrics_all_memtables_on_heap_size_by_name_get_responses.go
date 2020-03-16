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

// ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetReader is a Reader for the ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGet structure.
type ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK creates a ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK with default headers values
func NewColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK() *ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK {
	return &ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK{}
}

/*ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK handles this case with default header values.

ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK column family metrics all memtables on heap size by name get o k
*/
type ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/all_memtables_on_heap_size/{name}][%d] columnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
