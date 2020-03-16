// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/scylla/models"
)

// FailureDetectorSimpleStatesGetReader is a Reader for the FailureDetectorSimpleStatesGet structure.
type FailureDetectorSimpleStatesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FailureDetectorSimpleStatesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFailureDetectorSimpleStatesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFailureDetectorSimpleStatesGetOK creates a FailureDetectorSimpleStatesGetOK with default headers values
func NewFailureDetectorSimpleStatesGetOK() *FailureDetectorSimpleStatesGetOK {
	return &FailureDetectorSimpleStatesGetOK{}
}

/*FailureDetectorSimpleStatesGetOK handles this case with default header values.

FailureDetectorSimpleStatesGetOK failure detector simple states get o k
*/
type FailureDetectorSimpleStatesGetOK struct {
	Payload []*models.Mapper
}

func (o *FailureDetectorSimpleStatesGetOK) Error() string {
	return fmt.Sprintf("[GET /failure_detector/simple_states][%d] failureDetectorSimpleStatesGetOK  %+v", 200, o.Payload)
}

func (o *FailureDetectorSimpleStatesGetOK) GetPayload() []*models.Mapper {
	return o.Payload
}

func (o *FailureDetectorSimpleStatesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
