/*
API Documentation

Source of truth and network automation platform

API version: 1.3.7 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// SupplyEnum the model 'SupplyEnum'
type SupplyEnum string

// List of SupplyEnum
const (
	AC SupplyEnum = "ac"
	DC SupplyEnum = "dc"
)

// All allowed values of SupplyEnum enum
var AllowedSupplyEnumEnumValues = []SupplyEnum{
	"ac",
	"dc",
}

func (v *SupplyEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SupplyEnum(value)
	for _, existing := range AllowedSupplyEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SupplyEnum", value)
}

// NewSupplyEnumFromValue returns a pointer to a valid SupplyEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSupplyEnumFromValue(v string) (*SupplyEnum, error) {
	ev := SupplyEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SupplyEnum: valid values are %v", v, AllowedSupplyEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SupplyEnum) IsValid() bool {
	for _, existing := range AllowedSupplyEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SupplyEnum value
func (v SupplyEnum) Ptr() *SupplyEnum {
	return &v
}

type NullableSupplyEnum struct {
	value *SupplyEnum
	isSet bool
}

func (v NullableSupplyEnum) Get() *SupplyEnum {
	return v.value
}

func (v *NullableSupplyEnum) Set(val *SupplyEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplyEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplyEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplyEnum(val *SupplyEnum) *NullableSupplyEnum {
	return &NullableSupplyEnum{value: val, isSet: true}
}

func (v NullableSupplyEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplyEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

