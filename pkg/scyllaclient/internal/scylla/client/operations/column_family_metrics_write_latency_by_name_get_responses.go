// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// ColumnFamilyMetricsWriteLatencyByNameGetReader is a Reader for the ColumnFamilyMetricsWriteLatencyByNameGet structure.
type ColumnFamilyMetricsWriteLatencyByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsWriteLatencyByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsWriteLatencyByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsWriteLatencyByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsWriteLatencyByNameGetOK creates a ColumnFamilyMetricsWriteLatencyByNameGetOK with default headers values
func NewColumnFamilyMetricsWriteLatencyByNameGetOK() *ColumnFamilyMetricsWriteLatencyByNameGetOK {
	return &ColumnFamilyMetricsWriteLatencyByNameGetOK{}
}

/*ColumnFamilyMetricsWriteLatencyByNameGetOK handles this case with default header values.

ColumnFamilyMetricsWriteLatencyByNameGetOK column family metrics write latency by name get o k
*/
type ColumnFamilyMetricsWriteLatencyByNameGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsWriteLatencyByNameGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsWriteLatencyByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsWriteLatencyByNameGetDefault creates a ColumnFamilyMetricsWriteLatencyByNameGetDefault with default headers values
func NewColumnFamilyMetricsWriteLatencyByNameGetDefault(code int) *ColumnFamilyMetricsWriteLatencyByNameGetDefault {
	return &ColumnFamilyMetricsWriteLatencyByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsWriteLatencyByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsWriteLatencyByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics write latency by name get default response
func (o *ColumnFamilyMetricsWriteLatencyByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsWriteLatencyByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsWriteLatencyByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsWriteLatencyByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
