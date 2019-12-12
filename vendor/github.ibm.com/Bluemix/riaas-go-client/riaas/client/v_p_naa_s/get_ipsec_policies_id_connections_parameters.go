// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

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

// NewGetIpsecPoliciesIDConnectionsParams creates a new GetIpsecPoliciesIDConnectionsParams object
// with the default values initialized.
func NewGetIpsecPoliciesIDConnectionsParams() *GetIpsecPoliciesIDConnectionsParams {
	var ()
	return &GetIpsecPoliciesIDConnectionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIpsecPoliciesIDConnectionsParamsWithTimeout creates a new GetIpsecPoliciesIDConnectionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIpsecPoliciesIDConnectionsParamsWithTimeout(timeout time.Duration) *GetIpsecPoliciesIDConnectionsParams {
	var ()
	return &GetIpsecPoliciesIDConnectionsParams{

		timeout: timeout,
	}
}

// NewGetIpsecPoliciesIDConnectionsParamsWithContext creates a new GetIpsecPoliciesIDConnectionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIpsecPoliciesIDConnectionsParamsWithContext(ctx context.Context) *GetIpsecPoliciesIDConnectionsParams {
	var ()
	return &GetIpsecPoliciesIDConnectionsParams{

		Context: ctx,
	}
}

// NewGetIpsecPoliciesIDConnectionsParamsWithHTTPClient creates a new GetIpsecPoliciesIDConnectionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIpsecPoliciesIDConnectionsParamsWithHTTPClient(client *http.Client) *GetIpsecPoliciesIDConnectionsParams {
	var ()
	return &GetIpsecPoliciesIDConnectionsParams{
		HTTPClient: client,
	}
}

/*GetIpsecPoliciesIDConnectionsParams contains all the parameters to send to the API endpoint
for the get ipsec policies ID connections operation typically these are written to a http.Request
*/
type GetIpsecPoliciesIDConnectionsParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The policy identifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ipsec policies ID connections params
func (o *GetIpsecPoliciesIDConnectionsParams) WithTimeout(timeout time.Duration) *GetIpsecPoliciesIDConnectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ipsec policies ID connections params
func (o *GetIpsecPoliciesIDConnectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ipsec policies ID connections params
func (o *GetIpsecPoliciesIDConnectionsParams) WithContext(ctx context.Context) *GetIpsecPoliciesIDConnectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ipsec policies ID connections params
func (o *GetIpsecPoliciesIDConnectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ipsec policies ID connections params
func (o *GetIpsecPoliciesIDConnectionsParams) WithHTTPClient(client *http.Client) *GetIpsecPoliciesIDConnectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ipsec policies ID connections params
func (o *GetIpsecPoliciesIDConnectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get ipsec policies ID connections params
func (o *GetIpsecPoliciesIDConnectionsParams) WithGeneration(generation int64) *GetIpsecPoliciesIDConnectionsParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get ipsec policies ID connections params
func (o *GetIpsecPoliciesIDConnectionsParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the get ipsec policies ID connections params
func (o *GetIpsecPoliciesIDConnectionsParams) WithID(id string) *GetIpsecPoliciesIDConnectionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get ipsec policies ID connections params
func (o *GetIpsecPoliciesIDConnectionsParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get ipsec policies ID connections params
func (o *GetIpsecPoliciesIDConnectionsParams) WithVersion(version string) *GetIpsecPoliciesIDConnectionsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get ipsec policies ID connections params
func (o *GetIpsecPoliciesIDConnectionsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetIpsecPoliciesIDConnectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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
