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

// ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetReader is a Reader for the ColumnFamilyMetricsReadLatencyMovingAverageHistogramGet structure.
type ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsReadLatencyMovingAverageHistogramGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsReadLatencyMovingAverageHistogramGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsReadLatencyMovingAverageHistogramGetOK creates a ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetOK with default headers values
func NewColumnFamilyMetricsReadLatencyMovingAverageHistogramGetOK() *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetOK {
	return &ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetOK{}
}

/*ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetOK handles this case with default header values.

ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetOK column family metrics read latency moving average histogram get o k
*/
type ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetOK struct {
	Payload []*models.RateMovingAverageAndHistogram
}

func (o *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetOK) GetPayload() []*models.RateMovingAverageAndHistogram {
	return o.Payload
}

func (o *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsReadLatencyMovingAverageHistogramGetDefault creates a ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetDefault with default headers values
func NewColumnFamilyMetricsReadLatencyMovingAverageHistogramGetDefault(code int) *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetDefault {
	return &ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics read latency moving average histogram get default response
func (o *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsReadLatencyMovingAverageHistogramGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
