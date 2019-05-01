// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GcpStackDriverCollectorAttributeV2 gcp stack driver collector attribute v2
// swagger:model GcpStackDriverCollectorAttributeV2
type GcpStackDriverCollectorAttributeV2 struct {

	// period
	Period int32 `json:"period,omitempty"`
}

// Name gets the name of this subtype
func (m *GcpStackDriverCollectorAttributeV2) Name() string {
	return "GcpStackDriverCollectorAttributeV2"
}

// SetName sets the name of this subtype
func (m *GcpStackDriverCollectorAttributeV2) SetName(val string) {

}

// Period gets the period of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *GcpStackDriverCollectorAttributeV2) UnmarshalJSON(raw []byte) error {
	var data struct {

		// period
		Period int32 `json:"period,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Name string `json:"name"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result GcpStackDriverCollectorAttributeV2

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.Period = data.Period

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m GcpStackDriverCollectorAttributeV2) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// period
		Period int32 `json:"period,omitempty"`
	}{

		Period: m.Period,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Name string `json:"name"`
	}{

		Name: m.Name(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this gcp stack driver collector attribute v2
func (m *GcpStackDriverCollectorAttributeV2) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GcpStackDriverCollectorAttributeV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpStackDriverCollectorAttributeV2) UnmarshalBinary(b []byte) error {
	var res GcpStackDriverCollectorAttributeV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}