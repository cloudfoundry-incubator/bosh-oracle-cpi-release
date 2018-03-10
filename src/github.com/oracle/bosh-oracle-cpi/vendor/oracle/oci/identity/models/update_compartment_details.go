package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateCompartmentDetails update compartment details
// swagger:model UpdateCompartmentDetails
type UpdateCompartmentDetails struct {

	// The description you assign to the compartment. Does not have to be unique, and it's changeable.
	// Max Length: 400
	// Min Length: 1
	Description string `json:"description,omitempty"`
}

// Validate validates this update compartment details
func (m *UpdateCompartmentDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateCompartmentDetails) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(m.Description), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 400); err != nil {
		return err
	}

	return nil
}