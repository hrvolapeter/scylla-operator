// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ColumnFamilyMetricsCoordinatorScanGetReader is a Reader for the ColumnFamilyMetricsCoordinatorScanGet structure.
type ColumnFamilyMetricsCoordinatorScanGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsCoordinatorScanGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsCoordinatorScanGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsCoordinatorScanGetOK creates a ColumnFamilyMetricsCoordinatorScanGetOK with default headers values
func NewColumnFamilyMetricsCoordinatorScanGetOK() *ColumnFamilyMetricsCoordinatorScanGetOK {
	return &ColumnFamilyMetricsCoordinatorScanGetOK{}
}

/*ColumnFamilyMetricsCoordinatorScanGetOK handles this case with default header values.

ColumnFamilyMetricsCoordinatorScanGetOK column family metrics coordinator scan get o k
*/
type ColumnFamilyMetricsCoordinatorScanGetOK struct {
}

func (o *ColumnFamilyMetricsCoordinatorScanGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/coordinator/scan][%d] columnFamilyMetricsCoordinatorScanGetOK ", 200)
}

func (o *ColumnFamilyMetricsCoordinatorScanGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
