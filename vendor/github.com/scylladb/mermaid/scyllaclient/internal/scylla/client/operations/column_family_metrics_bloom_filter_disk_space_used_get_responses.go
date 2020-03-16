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

// ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetReader is a Reader for the ColumnFamilyMetricsBloomFilterDiskSpaceUsedGet structure.
type ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK creates a ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK with default headers values
func NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK() *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK {
	return &ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK{}
}

/*ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK handles this case with default header values.

ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK column family metrics bloom filter disk space used get o k
*/
type ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/bloom_filter_disk_space_used][%d] columnFamilyMetricsBloomFilterDiskSpaceUsedGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
