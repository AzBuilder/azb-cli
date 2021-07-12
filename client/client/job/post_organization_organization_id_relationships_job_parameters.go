// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPostOrganizationOrganizationIDRelationshipsJobParams creates a new PostOrganizationOrganizationIDRelationshipsJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostOrganizationOrganizationIDRelationshipsJobParams() *PostOrganizationOrganizationIDRelationshipsJobParams {
	return &PostOrganizationOrganizationIDRelationshipsJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostOrganizationOrganizationIDRelationshipsJobParamsWithTimeout creates a new PostOrganizationOrganizationIDRelationshipsJobParams object
// with the ability to set a timeout on a request.
func NewPostOrganizationOrganizationIDRelationshipsJobParamsWithTimeout(timeout time.Duration) *PostOrganizationOrganizationIDRelationshipsJobParams {
	return &PostOrganizationOrganizationIDRelationshipsJobParams{
		timeout: timeout,
	}
}

// NewPostOrganizationOrganizationIDRelationshipsJobParamsWithContext creates a new PostOrganizationOrganizationIDRelationshipsJobParams object
// with the ability to set a context for a request.
func NewPostOrganizationOrganizationIDRelationshipsJobParamsWithContext(ctx context.Context) *PostOrganizationOrganizationIDRelationshipsJobParams {
	return &PostOrganizationOrganizationIDRelationshipsJobParams{
		Context: ctx,
	}
}

// NewPostOrganizationOrganizationIDRelationshipsJobParamsWithHTTPClient creates a new PostOrganizationOrganizationIDRelationshipsJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostOrganizationOrganizationIDRelationshipsJobParamsWithHTTPClient(client *http.Client) *PostOrganizationOrganizationIDRelationshipsJobParams {
	return &PostOrganizationOrganizationIDRelationshipsJobParams{
		HTTPClient: client,
	}
}

/* PostOrganizationOrganizationIDRelationshipsJobParams contains all the parameters to send to the API endpoint
   for the post organization organization ID relationships job operation.

   Typically these are written to a http.Request.
*/
type PostOrganizationOrganizationIDRelationshipsJobParams struct {

	/* OrganizationID.

	   organization Identifier
	*/
	OrganizationID string

	// Relationship.
	Relationship PostOrganizationOrganizationIDRelationshipsJobBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post organization organization ID relationships job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOrganizationOrganizationIDRelationshipsJobParams) WithDefaults() *PostOrganizationOrganizationIDRelationshipsJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post organization organization ID relationships job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOrganizationOrganizationIDRelationshipsJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post organization organization ID relationships job params
func (o *PostOrganizationOrganizationIDRelationshipsJobParams) WithTimeout(timeout time.Duration) *PostOrganizationOrganizationIDRelationshipsJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post organization organization ID relationships job params
func (o *PostOrganizationOrganizationIDRelationshipsJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post organization organization ID relationships job params
func (o *PostOrganizationOrganizationIDRelationshipsJobParams) WithContext(ctx context.Context) *PostOrganizationOrganizationIDRelationshipsJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post organization organization ID relationships job params
func (o *PostOrganizationOrganizationIDRelationshipsJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post organization organization ID relationships job params
func (o *PostOrganizationOrganizationIDRelationshipsJobParams) WithHTTPClient(client *http.Client) *PostOrganizationOrganizationIDRelationshipsJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post organization organization ID relationships job params
func (o *PostOrganizationOrganizationIDRelationshipsJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the post organization organization ID relationships job params
func (o *PostOrganizationOrganizationIDRelationshipsJobParams) WithOrganizationID(organizationID string) *PostOrganizationOrganizationIDRelationshipsJobParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the post organization organization ID relationships job params
func (o *PostOrganizationOrganizationIDRelationshipsJobParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithRelationship adds the relationship to the post organization organization ID relationships job params
func (o *PostOrganizationOrganizationIDRelationshipsJobParams) WithRelationship(relationship PostOrganizationOrganizationIDRelationshipsJobBody) *PostOrganizationOrganizationIDRelationshipsJobParams {
	o.SetRelationship(relationship)
	return o
}

// SetRelationship adds the relationship to the post organization organization ID relationships job params
func (o *PostOrganizationOrganizationIDRelationshipsJobParams) SetRelationship(relationship PostOrganizationOrganizationIDRelationshipsJobBody) {
	o.Relationship = relationship
}

// WriteToRequest writes these params to a swagger request
func (o *PostOrganizationOrganizationIDRelationshipsJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Relationship); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}