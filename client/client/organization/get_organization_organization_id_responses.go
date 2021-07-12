// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"azb/client/models"
)

// GetOrganizationOrganizationIDReader is a Reader for the GetOrganizationOrganizationID structure.
type GetOrganizationOrganizationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationOrganizationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationOrganizationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationOrganizationIDOK creates a GetOrganizationOrganizationIDOK with default headers values
func NewGetOrganizationOrganizationIDOK() *GetOrganizationOrganizationIDOK {
	return &GetOrganizationOrganizationIDOK{}
}

/* GetOrganizationOrganizationIDOK describes a response with status code 200, with default header values.

Successful response
*/
type GetOrganizationOrganizationIDOK struct {
	Payload *GetOrganizationOrganizationIDOKBody
}

func (o *GetOrganizationOrganizationIDOK) Error() string {
	return fmt.Sprintf("[GET /organization/{organizationId}][%d] getOrganizationOrganizationIdOK  %+v", 200, o.Payload)
}
func (o *GetOrganizationOrganizationIDOK) GetPayload() *GetOrganizationOrganizationIDOKBody {
	return o.Payload
}

func (o *GetOrganizationOrganizationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrganizationOrganizationIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOrganizationOrganizationIDOKBody get organization organization ID o k body
swagger:model GetOrganizationOrganizationIDOKBody
*/
type GetOrganizationOrganizationIDOKBody struct {

	// data
	Data *models.Organization `json:"data,omitempty"`

	// Included resources
	// Unique: true
	Included []*GetOrganizationOrganizationIDOKBodyIncludedItems0 `json:"included"`
}

// Validate validates this get organization organization ID o k body
func (o *GetOrganizationOrganizationIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIncluded(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationOrganizationIDOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationOrganizationIdOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationOrganizationIDOKBody) validateIncluded(formats strfmt.Registry) error {
	if swag.IsZero(o.Included) { // not required
		return nil
	}

	if err := validate.UniqueItems("getOrganizationOrganizationIdOK"+"."+"included", "body", o.Included); err != nil {
		return err
	}

	for i := 0; i < len(o.Included); i++ {
		if swag.IsZero(o.Included[i]) { // not required
			continue
		}

		if o.Included[i] != nil {
			if err := o.Included[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrganizationOrganizationIdOK" + "." + "included" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get organization organization ID o k body based on the context it is used
func (o *GetOrganizationOrganizationIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateIncluded(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationOrganizationIDOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationOrganizationIdOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationOrganizationIDOKBody) contextValidateIncluded(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Included); i++ {

		if o.Included[i] != nil {
			if err := o.Included[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrganizationOrganizationIdOK" + "." + "included" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationOrganizationIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationOrganizationIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrganizationOrganizationIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetOrganizationOrganizationIDOKBodyIncludedItems0 get organization organization ID o k body included items0
swagger:model GetOrganizationOrganizationIDOKBodyIncludedItems0
*/
type GetOrganizationOrganizationIDOKBodyIncludedItems0 struct {

	// attributes
	Attributes interface{} `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// relationships
	Relationships interface{} `json:"relationships,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this get organization organization ID o k body included items0
func (o *GetOrganizationOrganizationIDOKBodyIncludedItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization organization ID o k body included items0 based on context it is used
func (o *GetOrganizationOrganizationIDOKBodyIncludedItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationOrganizationIDOKBodyIncludedItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationOrganizationIDOKBodyIncludedItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationOrganizationIDOKBodyIncludedItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
