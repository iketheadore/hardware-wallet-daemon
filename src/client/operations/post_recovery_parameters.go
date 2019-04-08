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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostRecoveryParams creates a new PostRecoveryParams object
// with the default values initialized.
func NewPostRecoveryParams() *PostRecoveryParams {
	var ()
	return &PostRecoveryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRecoveryParamsWithTimeout creates a new PostRecoveryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRecoveryParamsWithTimeout(timeout time.Duration) *PostRecoveryParams {
	var ()
	return &PostRecoveryParams{

		timeout: timeout,
	}
}

// NewPostRecoveryParamsWithContext creates a new PostRecoveryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRecoveryParamsWithContext(ctx context.Context) *PostRecoveryParams {
	var ()
	return &PostRecoveryParams{

		Context: ctx,
	}
}

// NewPostRecoveryParamsWithHTTPClient creates a new PostRecoveryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRecoveryParamsWithHTTPClient(client *http.Client) *PostRecoveryParams {
	var ()
	return &PostRecoveryParams{
		HTTPClient: client,
	}
}

/*PostRecoveryParams contains all the parameters to send to the API endpoint
for the post recovery operation typically these are written to a http.Request
*/
type PostRecoveryParams struct {

	/*DryRun
	  perform dry-run recovery workflow (for safe mnemonic validation)

	*/
	DryRun *bool
	/*UsePassphrase
	  ask for passphrase before starting operation

	*/
	UsePassphrase *bool
	/*WordCount
	  mnemonic seed length

	*/
	WordCount int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post recovery params
func (o *PostRecoveryParams) WithTimeout(timeout time.Duration) *PostRecoveryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post recovery params
func (o *PostRecoveryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post recovery params
func (o *PostRecoveryParams) WithContext(ctx context.Context) *PostRecoveryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post recovery params
func (o *PostRecoveryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post recovery params
func (o *PostRecoveryParams) WithHTTPClient(client *http.Client) *PostRecoveryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post recovery params
func (o *PostRecoveryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDryRun adds the dryRun to the post recovery params
func (o *PostRecoveryParams) WithDryRun(dryRun *bool) *PostRecoveryParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the post recovery params
func (o *PostRecoveryParams) SetDryRun(dryRun *bool) {
	o.DryRun = dryRun
}

// WithUsePassphrase adds the usePassphrase to the post recovery params
func (o *PostRecoveryParams) WithUsePassphrase(usePassphrase *bool) *PostRecoveryParams {
	o.SetUsePassphrase(usePassphrase)
	return o
}

// SetUsePassphrase adds the usePassphrase to the post recovery params
func (o *PostRecoveryParams) SetUsePassphrase(usePassphrase *bool) {
	o.UsePassphrase = usePassphrase
}

// WithWordCount adds the wordCount to the post recovery params
func (o *PostRecoveryParams) WithWordCount(wordCount int64) *PostRecoveryParams {
	o.SetWordCount(wordCount)
	return o
}

// SetWordCount adds the wordCount to the post recovery params
func (o *PostRecoveryParams) SetWordCount(wordCount int64) {
	o.WordCount = wordCount
}

// WriteToRequest writes these params to a swagger request
func (o *PostRecoveryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DryRun != nil {

		// form param dry-run
		var frDryRun bool
		if o.DryRun != nil {
			frDryRun = *o.DryRun
		}
		fDryRun := swag.FormatBool(frDryRun)
		if fDryRun != "" {
			if err := r.SetFormParam("dry-run", fDryRun); err != nil {
				return err
			}
		}

	}

	if o.UsePassphrase != nil {

		// form param use-passphrase
		var frUsePassphrase bool
		if o.UsePassphrase != nil {
			frUsePassphrase = *o.UsePassphrase
		}
		fUsePassphrase := swag.FormatBool(frUsePassphrase)
		if fUsePassphrase != "" {
			if err := r.SetFormParam("use-passphrase", fUsePassphrase); err != nil {
				return err
			}
		}

	}

	// form param word-count
	frWordCount := o.WordCount
	fWordCount := swag.FormatInt64(frWordCount)
	if fWordCount != "" {
		if err := r.SetFormParam("word-count", fWordCount); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}