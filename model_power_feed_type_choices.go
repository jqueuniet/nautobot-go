/*
API Documentation

Source of truth and network automation platform

API version: 1.3.10b1 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PowerFeedTypeChoices the model 'PowerFeedTypeChoices'
type PowerFeedTypeChoices string

// List of PowerFeedTypeChoices
const (
	PRIMARY PowerFeedTypeChoices = "primary"
	REDUNDANT PowerFeedTypeChoices = "redundant"
)

// All allowed values of PowerFeedTypeChoices enum
var AllowedPowerFeedTypeChoicesEnumValues = []PowerFeedTypeChoices{
	"primary",
	"redundant",
}

func (v *PowerFeedTypeChoices) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerFeedTypeChoices(value)
	for _, existing := range AllowedPowerFeedTypeChoicesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerFeedTypeChoices", value)
}

// NewPowerFeedTypeChoicesFromValue returns a pointer to a valid PowerFeedTypeChoices
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerFeedTypeChoicesFromValue(v string) (*PowerFeedTypeChoices, error) {
	ev := PowerFeedTypeChoices(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerFeedTypeChoices: valid values are %v", v, AllowedPowerFeedTypeChoicesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerFeedTypeChoices) IsValid() bool {
	for _, existing := range AllowedPowerFeedTypeChoicesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerFeedTypeChoices value
func (v PowerFeedTypeChoices) Ptr() *PowerFeedTypeChoices {
	return &v
}

type NullablePowerFeedTypeChoices struct {
	value *PowerFeedTypeChoices
	isSet bool
}

func (v NullablePowerFeedTypeChoices) Get() *PowerFeedTypeChoices {
	return v.value
}

func (v *NullablePowerFeedTypeChoices) Set(val *PowerFeedTypeChoices) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFeedTypeChoices) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFeedTypeChoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFeedTypeChoices(val *PowerFeedTypeChoices) *NullablePowerFeedTypeChoices {
	return &NullablePowerFeedTypeChoices{value: val, isSet: true}
}

func (v NullablePowerFeedTypeChoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFeedTypeChoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

