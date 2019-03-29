// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/skycoin/hardware-wallet-daemon/src/models"
)

// NewPostIntermediatePinMatrixParams creates a new PostIntermediatePinMatrixParams object
// with the default values initialized.
func NewPostIntermediatePinMatrixParams() *PostIntermediatePinMatrixParams {
	var ()
	return &PostIntermediatePinMatrixParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIntermediatePinMatrixParamsWithTimeout creates a new PostIntermediatePinMatrixParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIntermediatePinMatrixParamsWithTimeout(timeout time.Duration) *PostIntermediatePinMatrixParams {
	var ()
	return &PostIntermediatePinMatrixParams{

		timeout: timeout,
	}
}

// NewPostIntermediatePinMatrixParamsWithContext creates a new PostIntermediatePinMatrixParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIntermediatePinMatrixParamsWithContext(ctx context.Context) *PostIntermediatePinMatrixParams {
	var ()
	return &PostIntermediatePinMatrixParams{

		Context: ctx,
	}
}

// NewPostIntermediatePinMatrixParamsWithHTTPClient creates a new PostIntermediatePinMatrixParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIntermediatePinMatrixParamsWithHTTPClient(client *http.Client) *PostIntermediatePinMatrixParams {
	var ()
	return &PostIntermediatePinMatrixParams{
		HTTPClient: client,
	}
}

/*PostIntermediatePinMatrixParams contains all the parameters to send to the API endpoint
for the post intermediate pin matrix operation typically these are written to a http.Request
*/
type PostIntermediatePinMatrixParams struct {

	/*PinMatrixRequest
	  PinMatrixRequest is request data for /api/v1/intermediate/pinMatrix

	*/
	PinMatrixRequest *models.PinMatrixRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post intermediate pin matrix params
func (o *PostIntermediatePinMatrixParams) WithTimeout(timeout time.Duration) *PostIntermediatePinMatrixParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post intermediate pin matrix params
func (o *PostIntermediatePinMatrixParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post intermediate pin matrix params
func (o *PostIntermediatePinMatrixParams) WithContext(ctx context.Context) *PostIntermediatePinMatrixParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post intermediate pin matrix params
func (o *PostIntermediatePinMatrixParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post intermediate pin matrix params
func (o *PostIntermediatePinMatrixParams) WithHTTPClient(client *http.Client) *PostIntermediatePinMatrixParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post intermediate pin matrix params
func (o *PostIntermediatePinMatrixParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPinMatrixRequest adds the pinMatrixRequest to the post intermediate pin matrix params
func (o *PostIntermediatePinMatrixParams) WithPinMatrixRequest(pinMatrixRequest *models.PinMatrixRequest) *PostIntermediatePinMatrixParams {
	o.SetPinMatrixRequest(pinMatrixRequest)
	return o
}

// SetPinMatrixRequest adds the pinMatrixRequest to the post intermediate pin matrix params
func (o *PostIntermediatePinMatrixParams) SetPinMatrixRequest(pinMatrixRequest *models.PinMatrixRequest) {
	o.PinMatrixRequest = pinMatrixRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostIntermediatePinMatrixParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PinMatrixRequest != nil {
		if err := r.SetBodyParam(o.PinMatrixRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
