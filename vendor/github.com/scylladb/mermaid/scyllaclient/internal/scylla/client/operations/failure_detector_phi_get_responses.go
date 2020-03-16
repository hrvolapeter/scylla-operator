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

// FailureDetectorPhiGetReader is a Reader for the FailureDetectorPhiGet structure.
type FailureDetectorPhiGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FailureDetectorPhiGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFailureDetectorPhiGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFailureDetectorPhiGetOK creates a FailureDetectorPhiGetOK with default headers values
func NewFailureDetectorPhiGetOK() *FailureDetectorPhiGetOK {
	return &FailureDetectorPhiGetOK{}
}

/*FailureDetectorPhiGetOK handles this case with default header values.

FailureDetectorPhiGetOK failure detector phi get o k
*/
type FailureDetectorPhiGetOK struct {
	Payload string
}

func (o *FailureDetectorPhiGetOK) Error() string {
	return fmt.Sprintf("[GET /failure_detector/phi][%d] failureDetectorPhiGetOK  %+v", 200, o.Payload)
}

func (o *FailureDetectorPhiGetOK) GetPayload() string {
	return o.Payload
}

func (o *FailureDetectorPhiGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
