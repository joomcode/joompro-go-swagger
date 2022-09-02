// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"github.com/ssfilatov/go-swagger/examples/external-types/fred"
	custom "github.com/ssfilatov/go-swagger/examples/external-types/fred"
)

// Zzz This demonstrates variations in generated code, depending on how properties are declared.
//
// Some properties are directly based on some external type and some other define collections (slices, maps)
// of these external types.
//
// Notice the use of pointers for required properties, but not for slices or maps.
//
// In addition, it demonstrates how pointer generation may be controlled with the nullable hint or the x-nullable extension.
//
//	type Zzz struct {
//		Beta []MyOtherType `json:"beta"`
//		Delta MyInteger `json:"delta,omitempty"`
//		Epsilon []custom.MyAlternateType `json:"epsilon"`
//		Gamma fred.MyAlternateInteger `json:"gamma,omitempty"`
//		Meta MyType `json:"meta,omitempty"`
//
//		NullableBeta []*MyOtherType `json:"nullableBeta"`
//		NullableDelta *MyInteger `json:"nullableDelta,omitempty"`
//		NullableEpsilon []*custom.MyAlternateType `json:"nullableEpsilon"`
//		NullableGamma *fred.MyAlternateInteger `json:"nullableGamma,omitempty"`
//		NullableMeta MyType `json:"nullableMeta,omitempty"`
//
//		ReqBeta []MyOtherType `json:"reqBeta"`
//		ReqDelta *MyInteger `json:"reqDelta"`
//		ReqEpsilon []custom.MyAlternateType `json:"reqEpsilon"`
//		ReqGamma *fred.MyAlternateInteger `json:"reqGamma"`
//		ReqMeta *MyType `json:"reqMeta"`
//	}
//
// swagger:model Zzz
type Zzz struct {

	// This property defines an array of external types (in the same package).
	//
	// []MyOtherType
	//
	// The maxItems validation is generated and the external validation is called for every item.
	//
	// Max Items: 15
	Beta []MyOtherType `json:"beta"`

	// A type is provided (integer), and the implementation is an external type in the same package.
	//
	// The maximum validation remains documentary and is ignored by the generated code.
	//
	// Maximum: 15
	Delta MyInteger `json:"delta,omitempty"`

	// epsilon
	Epsilon []custom.MyAlternateType `json:"epsilon"`

	// Property defined as an external type from package "fred"
	//
	Gamma fred.MyAlternateInteger `json:"gamma,omitempty"`

	// meta
	Meta MyType `json:"meta,omitempty"`

	// This property defines an array of external types (in the same package).
	//
	// []MyOtherType
	//
	// The maxItems validation is generated and the external validation is called for every item.
	//
	// Max Items: 15
	NullableBeta []*MyOtherType `json:"nullableBeta"`

	// A type is provided (integer), and the implementation is an external type in the same package.
	//
	// The maximum validation remains documentary and is ignored by the generated code.
	//
	// NullableDelta *MyInteger
	//
	// Maximum: 15
	NullableDelta *MyInteger `json:"nullableDelta,omitempty"`

	// In this example, items are made nullable.
	//
	// NullableEpsilon []*custom.MyAlternateType `json:"nullableEpsilon"`
	//
	NullableEpsilon []*custom.MyAlternateType `json:"nullableEpsilon"`

	// Property defined as an external type from package "fred", with a nullable hint for the
	// external type.
	//
	// NullableGamma *fred.MyAlternateInteger `json:"nullableGamma,omitempty"`
	//
	NullableGamma *fred.MyAlternateInteger `json:"nullableGamma,omitempty"`

	// nullable meta
	NullableMeta MyType `json:"nullableMeta,omitempty"`

	// req beta
	// Required: true
	ReqBeta []MyOtherType `json:"reqBeta"`

	// req delta
	// Required: true
	ReqDelta *MyInteger `json:"reqDelta"`

	// req epsilon
	// Required: true
	ReqEpsilon []custom.MyAlternateType `json:"reqEpsilon"`

	// req gamma
	// Required: true
	ReqGamma *fred.MyAlternateInteger `json:"reqGamma"`

	// req meta
	// Required: true
	ReqMeta *MyType `json:"reqMeta"`
}

// Validate validates this zzz
func (m *Zzz) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEpsilon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGamma(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNullableBeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNullableDelta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNullableEpsilon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNullableGamma(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNullableMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReqBeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReqDelta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReqEpsilon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReqGamma(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReqMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Zzz) validateBeta(formats strfmt.Registry) error {
	if swag.IsZero(m.Beta) { // not required
		return nil
	}

	iBetaSize := int64(len(m.Beta))

	if err := validate.MaxItems("beta", "body", iBetaSize, 15); err != nil {
		return err
	}

	for i := 0; i < len(m.Beta); i++ {

		if err := m.Beta[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("beta" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("beta" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Zzz) validateDelta(formats strfmt.Registry) error {
	if swag.IsZero(m.Delta) { // not required
		return nil
	}

	if err := m.Delta.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("delta")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("delta")
		}
		return err
	}

	return nil
}

func (m *Zzz) validateEpsilon(formats strfmt.Registry) error {
	if swag.IsZero(m.Epsilon) { // not required
		return nil
	}

	for i := 0; i < len(m.Epsilon); i++ {

		if err := m.Epsilon[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("epsilon" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("epsilon" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Zzz) validateGamma(formats strfmt.Registry) error {
	if swag.IsZero(m.Gamma) { // not required
		return nil
	}

	if err := m.Gamma.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("gamma")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("gamma")
		}
		return err
	}

	return nil
}

func (m *Zzz) validateMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if err := m.Meta.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("meta")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("meta")
		}
		return err
	}

	return nil
}

func (m *Zzz) validateNullableBeta(formats strfmt.Registry) error {
	if swag.IsZero(m.NullableBeta) { // not required
		return nil
	}

	iNullableBetaSize := int64(len(m.NullableBeta))

	if err := validate.MaxItems("nullableBeta", "body", iNullableBetaSize, 15); err != nil {
		return err
	}

	for i := 0; i < len(m.NullableBeta); i++ {
		if swag.IsZero(m.NullableBeta[i]) { // not required
			continue
		}

		if m.NullableBeta[i] != nil {
			if err := m.NullableBeta[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nullableBeta" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nullableBeta" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Zzz) validateNullableDelta(formats strfmt.Registry) error {
	if swag.IsZero(m.NullableDelta) { // not required
		return nil
	}

	if m.NullableDelta != nil {
		if err := m.NullableDelta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nullableDelta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nullableDelta")
			}
			return err
		}
	}

	return nil
}

func (m *Zzz) validateNullableEpsilon(formats strfmt.Registry) error {
	if swag.IsZero(m.NullableEpsilon) { // not required
		return nil
	}

	for i := 0; i < len(m.NullableEpsilon); i++ {
		if swag.IsZero(m.NullableEpsilon[i]) { // not required
			continue
		}

		if m.NullableEpsilon[i] != nil {
			if err := m.NullableEpsilon[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nullableEpsilon" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nullableEpsilon" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Zzz) validateNullableGamma(formats strfmt.Registry) error {
	if swag.IsZero(m.NullableGamma) { // not required
		return nil
	}

	if m.NullableGamma != nil {
		if err := m.NullableGamma.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nullableGamma")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nullableGamma")
			}
			return err
		}
	}

	return nil
}

func (m *Zzz) validateNullableMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.NullableMeta) { // not required
		return nil
	}

	if err := m.NullableMeta.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("nullableMeta")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("nullableMeta")
		}
		return err
	}

	return nil
}

func (m *Zzz) validateReqBeta(formats strfmt.Registry) error {

	if err := validate.Required("reqBeta", "body", m.ReqBeta); err != nil {
		return err
	}

	for i := 0; i < len(m.ReqBeta); i++ {

		if err := m.ReqBeta[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reqBeta" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reqBeta" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Zzz) validateReqDelta(formats strfmt.Registry) error {

	if err := validate.Required("reqDelta", "body", m.ReqDelta); err != nil {
		return err
	}

	if err := validate.Required("reqDelta", "body", m.ReqDelta); err != nil {
		return err
	}

	if m.ReqDelta != nil {
		if err := m.ReqDelta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reqDelta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reqDelta")
			}
			return err
		}
	}

	return nil
}

func (m *Zzz) validateReqEpsilon(formats strfmt.Registry) error {

	if err := validate.Required("reqEpsilon", "body", m.ReqEpsilon); err != nil {
		return err
	}

	for i := 0; i < len(m.ReqEpsilon); i++ {

		if err := m.ReqEpsilon[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reqEpsilon" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reqEpsilon" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Zzz) validateReqGamma(formats strfmt.Registry) error {

	if err := validate.Required("reqGamma", "body", m.ReqGamma); err != nil {
		return err
	}

	if m.ReqGamma != nil {
		if err := m.ReqGamma.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reqGamma")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reqGamma")
			}
			return err
		}
	}

	return nil
}

func (m *Zzz) validateReqMeta(formats strfmt.Registry) error {

	if err := validate.Required("reqMeta", "body", m.ReqMeta); err != nil {
		return err
	}

	if m.ReqMeta != nil {
		if err := m.ReqMeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reqMeta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reqMeta")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this zzz based on the context it is used
func (m *Zzz) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNullableMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReqMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Zzz) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Meta.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("meta")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("meta")
		}
		return err
	}

	return nil
}

func (m *Zzz) contextValidateNullableMeta(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NullableMeta.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("nullableMeta")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("nullableMeta")
		}
		return err
	}

	return nil
}

func (m *Zzz) contextValidateReqMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.ReqMeta != nil {
		if err := m.ReqMeta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reqMeta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reqMeta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Zzz) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Zzz) UnmarshalBinary(b []byte) error {
	var res Zzz
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
