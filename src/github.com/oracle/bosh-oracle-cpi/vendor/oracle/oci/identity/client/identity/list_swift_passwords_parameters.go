package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListSwiftPasswordsParams creates a new ListSwiftPasswordsParams object
// with the default values initialized.
func NewListSwiftPasswordsParams() *ListSwiftPasswordsParams {
	var ()
	return &ListSwiftPasswordsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSwiftPasswordsParamsWithTimeout creates a new ListSwiftPasswordsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSwiftPasswordsParamsWithTimeout(timeout time.Duration) *ListSwiftPasswordsParams {
	var ()
	return &ListSwiftPasswordsParams{

		timeout: timeout,
	}
}

// NewListSwiftPasswordsParamsWithContext creates a new ListSwiftPasswordsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSwiftPasswordsParamsWithContext(ctx context.Context) *ListSwiftPasswordsParams {
	var ()
	return &ListSwiftPasswordsParams{

		Context: ctx,
	}
}

// NewListSwiftPasswordsParamsWithHTTPClient creates a new ListSwiftPasswordsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSwiftPasswordsParamsWithHTTPClient(client *http.Client) *ListSwiftPasswordsParams {
	var ()
	return &ListSwiftPasswordsParams{
		HTTPClient: client,
	}
}

/*ListSwiftPasswordsParams contains all the parameters to send to the API endpoint
for the list swift passwords operation typically these are written to a http.Request
*/
type ListSwiftPasswordsParams struct {

	/*UserID
	  The OCID of the user.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list swift passwords params
func (o *ListSwiftPasswordsParams) WithTimeout(timeout time.Duration) *ListSwiftPasswordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list swift passwords params
func (o *ListSwiftPasswordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list swift passwords params
func (o *ListSwiftPasswordsParams) WithContext(ctx context.Context) *ListSwiftPasswordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list swift passwords params
func (o *ListSwiftPasswordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list swift passwords params
func (o *ListSwiftPasswordsParams) WithHTTPClient(client *http.Client) *ListSwiftPasswordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list swift passwords params
func (o *ListSwiftPasswordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the list swift passwords params
func (o *ListSwiftPasswordsParams) WithUserID(userID string) *ListSwiftPasswordsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the list swift passwords params
func (o *ListSwiftPasswordsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *ListSwiftPasswordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
