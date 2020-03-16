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

// ColumnFamilyMaximumCompactionByNamePostReader is a Reader for the ColumnFamilyMaximumCompactionByNamePost structure.
type ColumnFamilyMaximumCompactionByNamePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMaximumCompactionByNamePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMaximumCompactionByNamePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMaximumCompactionByNamePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMaximumCompactionByNamePostOK creates a ColumnFamilyMaximumCompactionByNamePostOK with default headers values
func NewColumnFamilyMaximumCompactionByNamePostOK() *ColumnFamilyMaximumCompactionByNamePostOK {
	return &ColumnFamilyMaximumCompactionByNamePostOK{}
}

/*ColumnFamilyMaximumCompactionByNamePostOK handles this case with default header values.

ColumnFamilyMaximumCompactionByNamePostOK column family maximum compaction by name post o k
*/
type ColumnFamilyMaximumCompactionByNamePostOK struct {
	Payload string
}

func (o *ColumnFamilyMaximumCompactionByNamePostOK) GetPayload() string {
	return o.Payload
}

func (o *ColumnFamilyMaximumCompactionByNamePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMaximumCompactionByNamePostDefault creates a ColumnFamilyMaximumCompactionByNamePostDefault with default headers values
func NewColumnFamilyMaximumCompactionByNamePostDefault(code int) *ColumnFamilyMaximumCompactionByNamePostDefault {
	return &ColumnFamilyMaximumCompactionByNamePostDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMaximumCompactionByNamePostDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMaximumCompactionByNamePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family maximum compaction by name post default response
func (o *ColumnFamilyMaximumCompactionByNamePostDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMaximumCompactionByNamePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMaximumCompactionByNamePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMaximumCompactionByNamePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
