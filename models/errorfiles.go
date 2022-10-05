// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Errorfiles errorfiles
//
// swagger:model errorfiles
type Errorfiles struct {

	// codes
	Codes []int64 `json:"codes,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this errorfiles
func (m *Errorfiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var errorfilesCodesItemsEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[200,400,401,403,404,405,407,408,410,413,425,429,500,501,502,503,504]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		errorfilesCodesItemsEnum = append(errorfilesCodesItemsEnum, v)
	}
}

func (m *Errorfiles) validateCodesItemsEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, errorfilesCodesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Errorfiles) validateCodes(formats strfmt.Registry) error {
	if swag.IsZero(m.Codes) { // not required
		return nil
	}

	for i := 0; i < len(m.Codes); i++ {

		// value enum
		if err := m.validateCodesItemsEnum("codes"+"."+strconv.Itoa(i), "body", m.Codes[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this errorfiles based on context it is used
func (m *Errorfiles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Errorfiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Errorfiles) UnmarshalBinary(b []byte) error {
	var res Errorfiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
