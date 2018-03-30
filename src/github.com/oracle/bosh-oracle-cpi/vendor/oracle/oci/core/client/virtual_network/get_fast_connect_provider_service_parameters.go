// Code generated by go-swagger; DO NOT EDIT.

package virtual_network

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

// NewGetFastConnectProviderServiceParams creates a new GetFastConnectProviderServiceParams object
// with the default values initialized.
func NewGetFastConnectProviderServiceParams() *GetFastConnectProviderServiceParams {
	var ()
	return &GetFastConnectProviderServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFastConnectProviderServiceParamsWithTimeout creates a new GetFastConnectProviderServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFastConnectProviderServiceParamsWithTimeout(timeout time.Duration) *GetFastConnectProviderServiceParams {
	var ()
	return &GetFastConnectProviderServiceParams{

		timeout: timeout,
	}
}

// NewGetFastConnectProviderServiceParamsWithContext creates a new GetFastConnectProviderServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFastConnectProviderServiceParamsWithContext(ctx context.Context) *GetFastConnectProviderServiceParams {
	var ()
	return &GetFastConnectProviderServiceParams{

		Context: ctx,
	}
}

// NewGetFastConnectProviderServiceParamsWithHTTPClient creates a new GetFastConnectProviderServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFastConnectProviderServiceParamsWithHTTPClient(client *http.Client) *GetFastConnectProviderServiceParams {
	var ()
	return &GetFastConnectProviderServiceParams{
		HTTPClient: client,
	}
}

/*GetFastConnectProviderServiceParams contains all the parameters to send to the API endpoint
for the get fast connect provider service operation typically these are written to a http.Request
*/
type GetFastConnectProviderServiceParams struct {

	/*ProviderServiceID
	  The OCID of the provider service.

	*/
	ProviderServiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fast connect provider service params
func (o *GetFastConnectProviderServiceParams) WithTimeout(timeout time.Duration) *GetFastConnectProviderServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fast connect provider service params
func (o *GetFastConnectProviderServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fast connect provider service params
func (o *GetFastConnectProviderServiceParams) WithContext(ctx context.Context) *GetFastConnectProviderServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fast connect provider service params
func (o *GetFastConnectProviderServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fast connect provider service params
func (o *GetFastConnectProviderServiceParams) WithHTTPClient(client *http.Client) *GetFastConnectProviderServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fast connect provider service params
func (o *GetFastConnectProviderServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProviderServiceID adds the providerServiceID to the get fast connect provider service params
func (o *GetFastConnectProviderServiceParams) WithProviderServiceID(providerServiceID string) *GetFastConnectProviderServiceParams {
	o.SetProviderServiceID(providerServiceID)
	return o
}

// SetProviderServiceID adds the providerServiceId to the get fast connect provider service params
func (o *GetFastConnectProviderServiceParams) SetProviderServiceID(providerServiceID string) {
	o.ProviderServiceID = providerServiceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFastConnectProviderServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param providerServiceId
	if err := r.SetPathParam("providerServiceId", o.ProviderServiceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}