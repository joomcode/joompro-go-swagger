// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"github.com/ssfilatov/go-swagger/examples/external-types/fred"
)

// MyTuple Demonstrates references to some external type in the context of a tuple.
//
// Notice that "additionalItems" is not a construct that pass swagger validation,
// but is supported by go-swagger.
//
// swagger:model MyTuple
type MyTuple struct {

	// p0
	// Required: true
	P0 *MyType `json:"-"` // custom serializer

	// Second element of the tuple, defined as follows.
	//
	// P1 *fred.MyAlternateType `json:"-"` // custom serializer
	//
	// Required: true
	P1 *fred.MyAlternateType `json:"-"` // custom serializer

	// Additional items to a tuple, from an external type.
	// This defines the following field in the tuple
	//
	// MyTupleItems []map[string]fred.MyAlternateType
	//
	MyTupleItems []map[string]fred.MyAlternateType `json:"-"`
}

// UnmarshalJSON unmarshals this tuple type from a JSON array
func (m *MyTuple) UnmarshalJSON(raw []byte) error {
	// stage 1, get the array but just the array
	var stage1 []json.RawMessage
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&stage1); err != nil {
		return err
	}

	// stage 2: hydrates struct members with tuple elements
	var lastIndex int

	if len(stage1) > 0 {
		var dataP0 MyType
		buf = bytes.NewBuffer(stage1[0])
		dec := json.NewDecoder(buf)
		dec.UseNumber()
		if err := dec.Decode(&dataP0); err != nil {
			return err
		}
		m.P0 = &dataP0

		lastIndex = 0

	}
	if len(stage1) > 1 {
		var dataP1 fred.MyAlternateType
		buf = bytes.NewBuffer(stage1[1])
		dec := json.NewDecoder(buf)
		dec.UseNumber()
		if err := dec.Decode(&dataP1); err != nil {
			return err
		}
		m.P1 = &dataP1

		lastIndex = 1

	}

	// stage 3: hydrates AdditionalItems
	if len(stage1) > lastIndex+1 {
		for _, val := range stage1[lastIndex+1:] {
			var toadd map[string]fred.MyAlternateType
			buf = bytes.NewBuffer(val)
			dec := json.NewDecoder(buf)
			dec.UseNumber()
			if err := dec.Decode(&toadd); err != nil {
				return err
			}
			m.MyTupleItems = append(m.MyTupleItems, toadd)
		}
	}
	return nil
}

// MarshalJSON marshals this tuple type into a JSON array
func (m MyTuple) MarshalJSON() ([]byte, error) {
	data := []interface{}{
		m.P0, m.P1,
	}

	for _, v := range m.MyTupleItems {
		data = append(data, v)
	}
	return json.Marshal(data)
}

// Validate validates this my tuple
func (m *MyTuple) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateP0(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateP1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMyTupleItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MyTuple) validateP0(formats strfmt.Registry) error {

	if err := validate.Required("0", "body", m.P0); err != nil {
		return err
	}

	if m.P0 != nil {
		if err := m.P0.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("0")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("0")
			}
			return err
		}
	}

	return nil
}

func (m *MyTuple) validateP1(formats strfmt.Registry) error {

	if err := validate.Required("1", "body", m.P1); err != nil {
		return err
	}

	if m.P1 != nil {
		if err := m.P1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("1")
			}
			return err
		}
	}

	return nil
}

func (m *MyTuple) validateMyTupleItems(formats strfmt.Registry) error {

	for i := range m.MyTupleItems {

		for k := range m.MyTupleItems[i] {

			if val, ok := m.MyTupleItems[i][k]; ok {
				if err := val.Validate(formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName(strconv.Itoa(i+2) + "." + k)
					} else if ce, ok := err.(*errors.CompositeError); ok {
						return ce.ValidateName(strconv.Itoa(i+2) + "." + k)
					}
					return err
				}
			}

		}

	}

	return nil
}

// ContextValidate validate this my tuple based on the context it is used
func (m *MyTuple) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateP0(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MyTuple) contextValidateP0(ctx context.Context, formats strfmt.Registry) error {

	if m.P0 != nil {
		if err := m.P0.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("0")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("0")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MyTuple) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MyTuple) UnmarshalBinary(b []byte) error {
	var res MyTuple
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
