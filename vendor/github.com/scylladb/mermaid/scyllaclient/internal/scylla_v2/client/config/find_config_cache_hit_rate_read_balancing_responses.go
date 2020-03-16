// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/scylla_v2/models"
)

// FindConfigCacheHitRateReadBalancingReader is a Reader for the FindConfigCacheHitRateReadBalancing structure.
type FindConfigCacheHitRateReadBalancingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCacheHitRateReadBalancingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigCacheHitRateReadBalancingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigCacheHitRateReadBalancingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCacheHitRateReadBalancingOK creates a FindConfigCacheHitRateReadBalancingOK with default headers values
func NewFindConfigCacheHitRateReadBalancingOK() *FindConfigCacheHitRateReadBalancingOK {
	return &FindConfigCacheHitRateReadBalancingOK{}
}

/*FindConfigCacheHitRateReadBalancingOK handles this case with default header values.

Config value
*/
type FindConfigCacheHitRateReadBalancingOK struct {
	Payload bool
}

func (o *FindConfigCacheHitRateReadBalancingOK) Error() string {
	return fmt.Sprintf("[GET /config/cache_hit_rate_read_balancing][%d] findConfigCacheHitRateReadBalancingOK  %+v", 200, o.Payload)
}

func (o *FindConfigCacheHitRateReadBalancingOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigCacheHitRateReadBalancingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCacheHitRateReadBalancingDefault creates a FindConfigCacheHitRateReadBalancingDefault with default headers values
func NewFindConfigCacheHitRateReadBalancingDefault(code int) *FindConfigCacheHitRateReadBalancingDefault {
	return &FindConfigCacheHitRateReadBalancingDefault{
		_statusCode: code,
	}
}

/*FindConfigCacheHitRateReadBalancingDefault handles this case with default header values.

unexpected error
*/
type FindConfigCacheHitRateReadBalancingDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config cache hit rate read balancing default response
func (o *FindConfigCacheHitRateReadBalancingDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCacheHitRateReadBalancingDefault) Error() string {
	return fmt.Sprintf("[GET /config/cache_hit_rate_read_balancing][%d] find_config_cache_hit_rate_read_balancing default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigCacheHitRateReadBalancingDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigCacheHitRateReadBalancingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
