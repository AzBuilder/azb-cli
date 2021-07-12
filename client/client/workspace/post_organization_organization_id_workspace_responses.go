// Code generated by go-swagger; DO NOT EDIT.

package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"azb/client/models"
)

// PostOrganizationOrganizationIDWorkspaceReader is a Reader for the PostOrganizationOrganizationIDWorkspace structure.
type PostOrganizationOrganizationIDWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrganizationOrganizationIDWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostOrganizationOrganizationIDWorkspaceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOrganizationOrganizationIDWorkspaceCreated creates a PostOrganizationOrganizationIDWorkspaceCreated with default headers values
func NewPostOrganizationOrganizationIDWorkspaceCreated() *PostOrganizationOrganizationIDWorkspaceCreated {
	return &PostOrganizationOrganizationIDWorkspaceCreated{}
}

/* PostOrganizationOrganizationIDWorkspaceCreated describes a response with status code 201, with default header values.

Successful response
*/
type PostOrganizationOrganizationIDWorkspaceCreated struct {
	Payload *PostOrganizationOrganizationIDWorkspaceCreatedBody
}

func (o *PostOrganizationOrganizationIDWorkspaceCreated) Error() string {
	return fmt.Sprintf("[POST /organization/{organizationId}/workspace][%d] postOrganizationOrganizationIdWorkspaceCreated  %+v", 201, o.Payload)
}
func (o *PostOrganizationOrganizationIDWorkspaceCreated) GetPayload() *PostOrganizationOrganizationIDWorkspaceCreatedBody {
	return o.Payload
}

func (o *PostOrganizationOrganizationIDWorkspaceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostOrganizationOrganizationIDWorkspaceCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostOrganizationOrganizationIDWorkspaceBody post organization organization ID workspace body
swagger:model PostOrganizationOrganizationIDWorkspaceBody
*/
type PostOrganizationOrganizationIDWorkspaceBody struct {

	// data
	Data *models.Workspace `json:"data,omitempty"`
}

// Validate validates this post organization organization ID workspace body
func (o *PostOrganizationOrganizationIDWorkspaceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOrganizationOrganizationIDWorkspaceBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post organization organization ID workspace body based on the context it is used
func (o *PostOrganizationOrganizationIDWorkspaceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOrganizationOrganizationIDWorkspaceBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOrganizationOrganizationIDWorkspaceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOrganizationOrganizationIDWorkspaceBody) UnmarshalBinary(b []byte) error {
	var res PostOrganizationOrganizationIDWorkspaceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostOrganizationOrganizationIDWorkspaceCreatedBody post organization organization ID workspace created body
swagger:model PostOrganizationOrganizationIDWorkspaceCreatedBody
*/
type PostOrganizationOrganizationIDWorkspaceCreatedBody struct {

	// data
	Data *models.Workspace `json:"data,omitempty"`
}

// Validate validates this post organization organization ID workspace created body
func (o *PostOrganizationOrganizationIDWorkspaceCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOrganizationOrganizationIDWorkspaceCreatedBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postOrganizationOrganizationIdWorkspaceCreated" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post organization organization ID workspace created body based on the context it is used
func (o *PostOrganizationOrganizationIDWorkspaceCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOrganizationOrganizationIDWorkspaceCreatedBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postOrganizationOrganizationIdWorkspaceCreated" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOrganizationOrganizationIDWorkspaceCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOrganizationOrganizationIDWorkspaceCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostOrganizationOrganizationIDWorkspaceCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
