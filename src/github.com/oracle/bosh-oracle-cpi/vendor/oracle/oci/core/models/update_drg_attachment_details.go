// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateDrgAttachmentDetails update drg attachment details
// swagger:model UpdateDrgAttachmentDetails
type UpdateDrgAttachmentDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	//
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`
}

// Validate validates this update drg attachment details
func (m *UpdateDrgAttachmentDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateDrgAttachmentDetails) validateDisplayName(formats strfmt.Registry) error {

	if swag.IsZero(m.DisplayName) { // not required
		return nil
	}

	if err := validate.MinLength("displayName", "body", string(m.DisplayName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", string(m.DisplayName), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateDrgAttachmentDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateDrgAttachmentDetails) UnmarshalBinary(b []byte) error {
	var res UpdateDrgAttachmentDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
