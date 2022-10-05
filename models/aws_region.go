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

// AwsRegion AWS region
//
// AWS region configuration
// Example: {"access_key_id":"****************L7GT","allowlist":[{"key":"tag-key","value":"Instance:Having:This:Tag"}],"denylist":[{"key":"tag:Environment","value":"development"}],"enabled":true,"id":"0","ipv4_address":"private","name":"frontend-service","region":"us-east-1","retry_timeout":1,"secret_access_key":"****************soLl"}
//
// swagger:model awsRegion
type AwsRegion struct {

	// AWS Access Key ID.
	AccessKeyID string `json:"access_key_id,omitempty"`

	// Specify the AWS filters used to filter the EC2 instances to add
	Allowlist []*AwsFilters `json:"allowlist,omitempty"`

	// Specify the AWS filters used to filter the EC2 instances to ignore
	Denylist []*AwsFilters `json:"denylist,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// Auto generated ID.
	// Read Only: true
	// Pattern: ^[^\s]+$
	ID *string `json:"id,omitempty"`

	// Select which IPv4 address the Service Discovery has to use for the backend server entry
	// Required: true
	// Enum: [private public]
	IPV4Address *string `json:"ipv4_address"`

	// name
	// Required: true
	Name *string `json:"name"`

	// region
	// Required: true
	Region *string `json:"region"`

	// Duration in seconds in-between data pulling requests to the AWS region
	// Required: true
	// Minimum: 1
	RetryTimeout *int64 `json:"retry_timeout"`

	// AWS Secret Access Key.
	SecretAccessKey string `json:"secret_access_key,omitempty"`

	// server slots base
	ServerSlotsBase *int64 `json:"server_slots_base,omitempty"`

	// server slots growth increment
	ServerSlotsGrowthIncrement int64 `json:"server_slots_growth_increment,omitempty"`

	// server slots growth type
	// Enum: [linear exponential]
	ServerSlotsGrowthType *string `json:"server_slots_growth_type,omitempty"`
}

// Validate validates this aws region
func (m *AwsRegion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowlist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDenylist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV4Address(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetryTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerSlotsGrowthType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsRegion) validateAllowlist(formats strfmt.Registry) error {
	if swag.IsZero(m.Allowlist) { // not required
		return nil
	}

	for i := 0; i < len(m.Allowlist); i++ {
		if swag.IsZero(m.Allowlist[i]) { // not required
			continue
		}

		if m.Allowlist[i] != nil {
			if err := m.Allowlist[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowlist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowlist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AwsRegion) validateDenylist(formats strfmt.Registry) error {
	if swag.IsZero(m.Denylist) { // not required
		return nil
	}

	for i := 0; i < len(m.Denylist); i++ {
		if swag.IsZero(m.Denylist[i]) { // not required
			continue
		}

		if m.Denylist[i] != nil {
			if err := m.Denylist[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("denylist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("denylist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AwsRegion) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *AwsRegion) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", *m.ID, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var awsRegionTypeIPV4AddressPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["private","public"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		awsRegionTypeIPV4AddressPropEnum = append(awsRegionTypeIPV4AddressPropEnum, v)
	}
}

const (

	// AwsRegionIPV4AddressPrivate captures enum value "private"
	AwsRegionIPV4AddressPrivate string = "private"

	// AwsRegionIPV4AddressPublic captures enum value "public"
	AwsRegionIPV4AddressPublic string = "public"
)

// prop value enum
func (m *AwsRegion) validateIPV4AddressEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, awsRegionTypeIPV4AddressPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AwsRegion) validateIPV4Address(formats strfmt.Registry) error {

	if err := validate.Required("ipv4_address", "body", m.IPV4Address); err != nil {
		return err
	}

	// value enum
	if err := m.validateIPV4AddressEnum("ipv4_address", "body", *m.IPV4Address); err != nil {
		return err
	}

	return nil
}

func (m *AwsRegion) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AwsRegion) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *AwsRegion) validateRetryTimeout(formats strfmt.Registry) error {

	if err := validate.Required("retry_timeout", "body", m.RetryTimeout); err != nil {
		return err
	}

	if err := validate.MinimumInt("retry_timeout", "body", *m.RetryTimeout, 1, false); err != nil {
		return err
	}

	return nil
}

var awsRegionTypeServerSlotsGrowthTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["linear","exponential"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		awsRegionTypeServerSlotsGrowthTypePropEnum = append(awsRegionTypeServerSlotsGrowthTypePropEnum, v)
	}
}

const (

	// AwsRegionServerSlotsGrowthTypeLinear captures enum value "linear"
	AwsRegionServerSlotsGrowthTypeLinear string = "linear"

	// AwsRegionServerSlotsGrowthTypeExponential captures enum value "exponential"
	AwsRegionServerSlotsGrowthTypeExponential string = "exponential"
)

// prop value enum
func (m *AwsRegion) validateServerSlotsGrowthTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, awsRegionTypeServerSlotsGrowthTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AwsRegion) validateServerSlotsGrowthType(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerSlotsGrowthType) { // not required
		return nil
	}

	// value enum
	if err := m.validateServerSlotsGrowthTypeEnum("server_slots_growth_type", "body", *m.ServerSlotsGrowthType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this aws region based on the context it is used
func (m *AwsRegion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowlist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDenylist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsRegion) contextValidateAllowlist(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Allowlist); i++ {

		if m.Allowlist[i] != nil {
			if err := m.Allowlist[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowlist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowlist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AwsRegion) contextValidateDenylist(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Denylist); i++ {

		if m.Denylist[i] != nil {
			if err := m.Denylist[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("denylist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("denylist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AwsRegion) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsRegion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsRegion) UnmarshalBinary(b []byte) error {
	var res AwsRegion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
