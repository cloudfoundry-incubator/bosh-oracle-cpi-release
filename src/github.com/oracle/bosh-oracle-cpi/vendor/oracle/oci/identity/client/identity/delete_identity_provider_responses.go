package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/oci/identity/models"
)

// DeleteIdentityProviderReader is a Reader for the DeleteIdentityProvider structure.
type DeleteIdentityProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIdentityProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteIdentityProviderNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteIdentityProviderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteIdentityProviderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteIdentityProviderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteIdentityProviderConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteIdentityProviderPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteIdentityProviderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteIdentityProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteIdentityProviderNoContent creates a DeleteIdentityProviderNoContent with default headers values
func NewDeleteIdentityProviderNoContent() *DeleteIdentityProviderNoContent {
	return &DeleteIdentityProviderNoContent{}
}

/*DeleteIdentityProviderNoContent handles this case with default header values.

The identity provider is being deleted.
*/
type DeleteIdentityProviderNoContent struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string
}

func (o *DeleteIdentityProviderNoContent) Error() string {
	return fmt.Sprintf("[DELETE /identityProviders/{identityProviderId}][%d] deleteIdentityProviderNoContent ", 204)
}

func (o *DeleteIdentityProviderNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	return nil
}

// NewDeleteIdentityProviderBadRequest creates a DeleteIdentityProviderBadRequest with default headers values
func NewDeleteIdentityProviderBadRequest() *DeleteIdentityProviderBadRequest {
	return &DeleteIdentityProviderBadRequest{}
}

/*DeleteIdentityProviderBadRequest handles this case with default header values.

Bad Request
*/
type DeleteIdentityProviderBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteIdentityProviderBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /identityProviders/{identityProviderId}][%d] deleteIdentityProviderBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIdentityProviderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityProviderForbidden creates a DeleteIdentityProviderForbidden with default headers values
func NewDeleteIdentityProviderForbidden() *DeleteIdentityProviderForbidden {
	return &DeleteIdentityProviderForbidden{}
}

/*DeleteIdentityProviderForbidden handles this case with default header values.

Forbidden
*/
type DeleteIdentityProviderForbidden struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteIdentityProviderForbidden) Error() string {
	return fmt.Sprintf("[DELETE /identityProviders/{identityProviderId}][%d] deleteIdentityProviderForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIdentityProviderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityProviderNotFound creates a DeleteIdentityProviderNotFound with default headers values
func NewDeleteIdentityProviderNotFound() *DeleteIdentityProviderNotFound {
	return &DeleteIdentityProviderNotFound{}
}

/*DeleteIdentityProviderNotFound handles this case with default header values.

Not Found
*/
type DeleteIdentityProviderNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteIdentityProviderNotFound) Error() string {
	return fmt.Sprintf("[DELETE /identityProviders/{identityProviderId}][%d] deleteIdentityProviderNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIdentityProviderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityProviderConflict creates a DeleteIdentityProviderConflict with default headers values
func NewDeleteIdentityProviderConflict() *DeleteIdentityProviderConflict {
	return &DeleteIdentityProviderConflict{}
}

/*DeleteIdentityProviderConflict handles this case with default header values.

Conflict
*/
type DeleteIdentityProviderConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteIdentityProviderConflict) Error() string {
	return fmt.Sprintf("[DELETE /identityProviders/{identityProviderId}][%d] deleteIdentityProviderConflict  %+v", 409, o.Payload)
}

func (o *DeleteIdentityProviderConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityProviderPreconditionFailed creates a DeleteIdentityProviderPreconditionFailed with default headers values
func NewDeleteIdentityProviderPreconditionFailed() *DeleteIdentityProviderPreconditionFailed {
	return &DeleteIdentityProviderPreconditionFailed{}
}

/*DeleteIdentityProviderPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type DeleteIdentityProviderPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteIdentityProviderPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /identityProviders/{identityProviderId}][%d] deleteIdentityProviderPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteIdentityProviderPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityProviderInternalServerError creates a DeleteIdentityProviderInternalServerError with default headers values
func NewDeleteIdentityProviderInternalServerError() *DeleteIdentityProviderInternalServerError {
	return &DeleteIdentityProviderInternalServerError{}
}

/*DeleteIdentityProviderInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteIdentityProviderInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteIdentityProviderInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /identityProviders/{identityProviderId}][%d] deleteIdentityProviderInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIdentityProviderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityProviderDefault creates a DeleteIdentityProviderDefault with default headers values
func NewDeleteIdentityProviderDefault(code int) *DeleteIdentityProviderDefault {
	return &DeleteIdentityProviderDefault{
		_statusCode: code,
	}
}

/*DeleteIdentityProviderDefault handles this case with default header values.

An error has occurred.

*/
type DeleteIdentityProviderDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the delete identity provider default response
func (o *DeleteIdentityProviderDefault) Code() int {
	return o._statusCode
}

func (o *DeleteIdentityProviderDefault) Error() string {
	return fmt.Sprintf("[DELETE /identityProviders/{identityProviderId}][%d] DeleteIdentityProvider default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteIdentityProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}