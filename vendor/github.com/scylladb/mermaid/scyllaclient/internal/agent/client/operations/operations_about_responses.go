// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/agent/models"
)

// OperationsAboutReader is a Reader for the OperationsAbout structure.
type OperationsAboutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperationsAboutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOperationsAboutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOperationsAboutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOperationsAboutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOperationsAboutOK creates a OperationsAboutOK with default headers values
func NewOperationsAboutOK() *OperationsAboutOK {
	return &OperationsAboutOK{}
}

/*OperationsAboutOK handles this case with default header values.

File system details
*/
type OperationsAboutOK struct {
	Payload *models.FileSystemDetails
}

func (o *OperationsAboutOK) Error() string {
	return fmt.Sprintf("[POST /rclone/operations/about][%d] operationsAboutOK  %+v", 200, o.Payload)
}

func (o *OperationsAboutOK) GetPayload() *models.FileSystemDetails {
	return o.Payload
}

func (o *OperationsAboutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileSystemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperationsAboutNotFound creates a OperationsAboutNotFound with default headers values
func NewOperationsAboutNotFound() *OperationsAboutNotFound {
	return &OperationsAboutNotFound{}
}

/*OperationsAboutNotFound handles this case with default header values.

Not found
*/
type OperationsAboutNotFound struct {
	Payload *models.ErrorResponse
}

func (o *OperationsAboutNotFound) Error() string {
	return fmt.Sprintf("[POST /rclone/operations/about][%d] operationsAboutNotFound  %+v", 404, o.Payload)
}

func (o *OperationsAboutNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *OperationsAboutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperationsAboutInternalServerError creates a OperationsAboutInternalServerError with default headers values
func NewOperationsAboutInternalServerError() *OperationsAboutInternalServerError {
	return &OperationsAboutInternalServerError{}
}

/*OperationsAboutInternalServerError handles this case with default header values.

Server error
*/
type OperationsAboutInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *OperationsAboutInternalServerError) Error() string {
	return fmt.Sprintf("[POST /rclone/operations/about][%d] operationsAboutInternalServerError  %+v", 500, o.Payload)
}

func (o *OperationsAboutInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *OperationsAboutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
