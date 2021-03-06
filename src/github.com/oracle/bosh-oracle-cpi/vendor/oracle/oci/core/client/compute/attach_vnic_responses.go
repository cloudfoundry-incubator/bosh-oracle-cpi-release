// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "oracle/oci/core/models"
)

// AttachVnicReader is a Reader for the AttachVnic structure.
type AttachVnicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachVnicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAttachVnicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAttachVnicBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAttachVnicUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAttachVnicNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewAttachVnicConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAttachVnicInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAttachVnicDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAttachVnicOK creates a AttachVnicOK with default headers values
func NewAttachVnicOK() *AttachVnicOK {
	return &AttachVnicOK{}
}

/*AttachVnicOK handles this case with default header values.

The VNIC is being attached.
*/
type AttachVnicOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.VnicAttachment
}

func (o *AttachVnicOK) Error() string {
	return fmt.Sprintf("[POST /vnicAttachments/][%d] attachVnicOK  %+v", 200, o.Payload)
}

func (o *AttachVnicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.VnicAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachVnicBadRequest creates a AttachVnicBadRequest with default headers values
func NewAttachVnicBadRequest() *AttachVnicBadRequest {
	return &AttachVnicBadRequest{}
}

/*AttachVnicBadRequest handles this case with default header values.

Bad Request
*/
type AttachVnicBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *AttachVnicBadRequest) Error() string {
	return fmt.Sprintf("[POST /vnicAttachments/][%d] attachVnicBadRequest  %+v", 400, o.Payload)
}

func (o *AttachVnicBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachVnicUnauthorized creates a AttachVnicUnauthorized with default headers values
func NewAttachVnicUnauthorized() *AttachVnicUnauthorized {
	return &AttachVnicUnauthorized{}
}

/*AttachVnicUnauthorized handles this case with default header values.

Unauthorized
*/
type AttachVnicUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *AttachVnicUnauthorized) Error() string {
	return fmt.Sprintf("[POST /vnicAttachments/][%d] attachVnicUnauthorized  %+v", 401, o.Payload)
}

func (o *AttachVnicUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachVnicNotFound creates a AttachVnicNotFound with default headers values
func NewAttachVnicNotFound() *AttachVnicNotFound {
	return &AttachVnicNotFound{}
}

/*AttachVnicNotFound handles this case with default header values.

Not Found
*/
type AttachVnicNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *AttachVnicNotFound) Error() string {
	return fmt.Sprintf("[POST /vnicAttachments/][%d] attachVnicNotFound  %+v", 404, o.Payload)
}

func (o *AttachVnicNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachVnicConflict creates a AttachVnicConflict with default headers values
func NewAttachVnicConflict() *AttachVnicConflict {
	return &AttachVnicConflict{}
}

/*AttachVnicConflict handles this case with default header values.

Conflict
*/
type AttachVnicConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *AttachVnicConflict) Error() string {
	return fmt.Sprintf("[POST /vnicAttachments/][%d] attachVnicConflict  %+v", 409, o.Payload)
}

func (o *AttachVnicConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachVnicInternalServerError creates a AttachVnicInternalServerError with default headers values
func NewAttachVnicInternalServerError() *AttachVnicInternalServerError {
	return &AttachVnicInternalServerError{}
}

/*AttachVnicInternalServerError handles this case with default header values.

Internal Server Error
*/
type AttachVnicInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *AttachVnicInternalServerError) Error() string {
	return fmt.Sprintf("[POST /vnicAttachments/][%d] attachVnicInternalServerError  %+v", 500, o.Payload)
}

func (o *AttachVnicInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachVnicDefault creates a AttachVnicDefault with default headers values
func NewAttachVnicDefault(code int) *AttachVnicDefault {
	return &AttachVnicDefault{
		_statusCode: code,
	}
}

/*AttachVnicDefault handles this case with default header values.

An error has occurred.
*/
type AttachVnicDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the attach vnic default response
func (o *AttachVnicDefault) Code() int {
	return o._statusCode
}

func (o *AttachVnicDefault) Error() string {
	return fmt.Sprintf("[POST /vnicAttachments/][%d] AttachVnic default  %+v", o._statusCode, o.Payload)
}

func (o *AttachVnicDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
