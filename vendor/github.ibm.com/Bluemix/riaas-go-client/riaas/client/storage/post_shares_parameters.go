// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewPostSharesParams creates a new PostSharesParams object
// with the default values initialized.
func NewPostSharesParams() *PostSharesParams {
	var ()
	return &PostSharesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSharesParamsWithTimeout creates a new PostSharesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSharesParamsWithTimeout(timeout time.Duration) *PostSharesParams {
	var ()
	return &PostSharesParams{

		timeout: timeout,
	}
}

// NewPostSharesParamsWithContext creates a new PostSharesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSharesParamsWithContext(ctx context.Context) *PostSharesParams {
	var ()
	return &PostSharesParams{

		Context: ctx,
	}
}

// NewPostSharesParamsWithHTTPClient creates a new PostSharesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSharesParamsWithHTTPClient(client *http.Client) *PostSharesParams {
	var ()
	return &PostSharesParams{
		HTTPClient: client,
	}
}

/*PostSharesParams contains all the parameters to send to the API endpoint
for the post shares operation typically these are written to a http.Request
*/
type PostSharesParams struct {

	/*ShareTemplate*/
	ShareTemplate PostSharesBody
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post shares params
func (o *PostSharesParams) WithTimeout(timeout time.Duration) *PostSharesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post shares params
func (o *PostSharesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post shares params
func (o *PostSharesParams) WithContext(ctx context.Context) *PostSharesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post shares params
func (o *PostSharesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post shares params
func (o *PostSharesParams) WithHTTPClient(client *http.Client) *PostSharesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post shares params
func (o *PostSharesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithShareTemplate adds the shareTemplate to the post shares params
func (o *PostSharesParams) WithShareTemplate(shareTemplate PostSharesBody) *PostSharesParams {
	o.SetShareTemplate(shareTemplate)
	return o
}

// SetShareTemplate adds the shareTemplate to the post shares params
func (o *PostSharesParams) SetShareTemplate(shareTemplate PostSharesBody) {
	o.ShareTemplate = shareTemplate
}

// WithGeneration adds the generation to the post shares params
func (o *PostSharesParams) WithGeneration(generation int64) *PostSharesParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the post shares params
func (o *PostSharesParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithVersion adds the version to the post shares params
func (o *PostSharesParams) WithVersion(version string) *PostSharesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post shares params
func (o *PostSharesParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PostSharesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.ShareTemplate); err != nil {
		return err
	}

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
