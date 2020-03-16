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

// CacheServiceMetricsCounterHitsMovingAvrageGetReader is a Reader for the CacheServiceMetricsCounterHitsMovingAvrageGet structure.
type CacheServiceMetricsCounterHitsMovingAvrageGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsCounterHitsMovingAvrageGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceMetricsCounterHitsMovingAvrageGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceMetricsCounterHitsMovingAvrageGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceMetricsCounterHitsMovingAvrageGetOK creates a CacheServiceMetricsCounterHitsMovingAvrageGetOK with default headers values
func NewCacheServiceMetricsCounterHitsMovingAvrageGetOK() *CacheServiceMetricsCounterHitsMovingAvrageGetOK {
	return &CacheServiceMetricsCounterHitsMovingAvrageGetOK{}
}

/*CacheServiceMetricsCounterHitsMovingAvrageGetOK handles this case with default header values.

CacheServiceMetricsCounterHitsMovingAvrageGetOK cache service metrics counter hits moving avrage get o k
*/
type CacheServiceMetricsCounterHitsMovingAvrageGetOK struct {
	Payload *models.RateMovingAverage
}

func (o *CacheServiceMetricsCounterHitsMovingAvrageGetOK) GetPayload() *models.RateMovingAverage {
	return o.Payload
}

func (o *CacheServiceMetricsCounterHitsMovingAvrageGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RateMovingAverage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheServiceMetricsCounterHitsMovingAvrageGetDefault creates a CacheServiceMetricsCounterHitsMovingAvrageGetDefault with default headers values
func NewCacheServiceMetricsCounterHitsMovingAvrageGetDefault(code int) *CacheServiceMetricsCounterHitsMovingAvrageGetDefault {
	return &CacheServiceMetricsCounterHitsMovingAvrageGetDefault{
		_statusCode: code,
	}
}

/*CacheServiceMetricsCounterHitsMovingAvrageGetDefault handles this case with default header values.

internal server error
*/
type CacheServiceMetricsCounterHitsMovingAvrageGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service metrics counter hits moving avrage get default response
func (o *CacheServiceMetricsCounterHitsMovingAvrageGetDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceMetricsCounterHitsMovingAvrageGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceMetricsCounterHitsMovingAvrageGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceMetricsCounterHitsMovingAvrageGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
