// Code generated by go-swagger; DO NOT EDIT.

package job

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

// GetOrganizationOrganizationIDJobReader is a Reader for the GetOrganizationOrganizationIDJob structure.
type GetOrganizationOrganizationIDJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationOrganizationIDJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationOrganizationIDJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationOrganizationIDJobOK creates a GetOrganizationOrganizationIDJobOK with default headers values
func NewGetOrganizationOrganizationIDJobOK() *GetOrganizationOrganizationIDJobOK {
	return &GetOrganizationOrganizationIDJobOK{}
}

/* GetOrganizationOrganizationIDJobOK describes a response with status code 200, with default header values.

Successful response
*/
type GetOrganizationOrganizationIDJobOK struct {
	Payload *GetOrganizationOrganizationIDJobOKBody
}

func (o *GetOrganizationOrganizationIDJobOK) Error() string {
	return fmt.Sprintf("[GET /organization/{organizationId}/job][%d] getOrganizationOrganizationIdJobOK  %+v", 200, o.Payload)
}
func (o *GetOrganizationOrganizationIDJobOK) GetPayload() *GetOrganizationOrganizationIDJobOKBody {
	return o.Payload
}

func (o *GetOrganizationOrganizationIDJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrganizationOrganizationIDJobOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOrganizationOrganizationIDJobOKBody get organization organization ID job o k body
swagger:model GetOrganizationOrganizationIDJobOKBody
*/
type GetOrganizationOrganizationIDJobOKBody struct {

	// data
	Data []*models.Job `json:"data"`

	// Included resources
	// Unique: true
	Included []*GetOrganizationOrganizationIDJobOKBodyIncludedItems0 `json:"included"`
}

// Validate validates this get organization organization ID job o k body
func (o *GetOrganizationOrganizationIDJobOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetOrganizationOrganizationIDJobOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrganizationOrganizationIdJobOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetOrganizationOrganizationIDJobOKBody) validateIncluded(formats strfmt.Registry) error {
	if swag.IsZero(o.Included) { // not required
		return nil
	}

	if err := validate.UniqueItems("getOrganizationOrganizationIdJobOK"+"."+"included", "body", o.Included); err != nil {
		return err
	}

	for i := 0; i < len(o.Included); i++ {
		if swag.IsZero(o.Included[i]) { // not required
			continue
		}

		if o.Included[i] != nil {
			if err := o.Included[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrganizationOrganizationIdJobOK" + "." + "included" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get organization organization ID job o k body based on the context it is used
func (o *GetOrganizationOrganizationIDJobOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetOrganizationOrganizationIDJobOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrganizationOrganizationIdJobOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetOrganizationOrganizationIDJobOKBody) contextValidateIncluded(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Included); i++ {

		if o.Included[i] != nil {
			if err := o.Included[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrganizationOrganizationIdJobOK" + "." + "included" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationOrganizationIDJobOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationOrganizationIDJobOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrganizationOrganizationIDJobOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetOrganizationOrganizationIDJobOKBodyIncludedItems0 get organization organization ID job o k body included items0
swagger:model GetOrganizationOrganizationIDJobOKBodyIncludedItems0
*/
type GetOrganizationOrganizationIDJobOKBodyIncludedItems0 struct {

	// attributes
	Attributes interface{} `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// relationships
	Relationships interface{} `json:"relationships,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this get organization organization ID job o k body included items0
func (o *GetOrganizationOrganizationIDJobOKBodyIncludedItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization organization ID job o k body included items0 based on context it is used
func (o *GetOrganizationOrganizationIDJobOKBodyIncludedItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationOrganizationIDJobOKBodyIncludedItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationOrganizationIDJobOKBodyIncludedItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationOrganizationIDJobOKBodyIncludedItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}