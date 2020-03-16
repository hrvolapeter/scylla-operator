// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CollectdByPluginidPostReader is a Reader for the CollectdByPluginidPost structure.
type CollectdByPluginidPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollectdByPluginidPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCollectdByPluginidPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCollectdByPluginidPostOK creates a CollectdByPluginidPostOK with default headers values
func NewCollectdByPluginidPostOK() *CollectdByPluginidPostOK {
	return &CollectdByPluginidPostOK{}
}

/*CollectdByPluginidPostOK handles this case with default header values.

CollectdByPluginidPostOK collectd by pluginid post o k
*/
type CollectdByPluginidPostOK struct {
}

func (o *CollectdByPluginidPostOK) Error() string {
	return fmt.Sprintf("[POST /collectd/{pluginid}][%d] collectdByPluginidPostOK ", 200)
}

func (o *CollectdByPluginidPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
