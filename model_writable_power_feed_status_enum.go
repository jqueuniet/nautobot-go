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

// WritablePowerFeedStatusEnum the model 'WritablePowerFeedStatusEnum'
type WritablePowerFeedStatusEnum string

// List of WritablePowerFeedStatusEnum
//const (
//	ACTIVE WritablePowerFeedStatusEnum = "active"
//	FAILED WritablePowerFeedStatusEnum = "failed"
//	OFFLINE WritablePowerFeedStatusEnum = "offline"
//	PLANNED WritablePowerFeedStatusEnum = "planned"
//)

// All allowed values of WritablePowerFeedStatusEnum enum
var AllowedWritablePowerFeedStatusEnumEnumValues = []WritablePowerFeedStatusEnum{
	"active",
	"failed",
	"offline",
	"planned",
}

func (v *WritablePowerFeedStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WritablePowerFeedStatusEnum(value)
	for _, existing := range AllowedWritablePowerFeedStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WritablePowerFeedStatusEnum", value)
}

// NewWritablePowerFeedStatusEnumFromValue returns a pointer to a valid WritablePowerFeedStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWritablePowerFeedStatusEnumFromValue(v string) (*WritablePowerFeedStatusEnum, error) {
	ev := WritablePowerFeedStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WritablePowerFeedStatusEnum: valid values are %v", v, AllowedWritablePowerFeedStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WritablePowerFeedStatusEnum) IsValid() bool {
	for _, existing := range AllowedWritablePowerFeedStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WritablePowerFeedStatusEnum value
func (v WritablePowerFeedStatusEnum) Ptr() *WritablePowerFeedStatusEnum {
	return &v
}

type NullableWritablePowerFeedStatusEnum struct {
	value *WritablePowerFeedStatusEnum
	isSet bool
}

func (v NullableWritablePowerFeedStatusEnum) Get() *WritablePowerFeedStatusEnum {
	return v.value
}

func (v *NullableWritablePowerFeedStatusEnum) Set(val *WritablePowerFeedStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableWritablePowerFeedStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableWritablePowerFeedStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritablePowerFeedStatusEnum(val *WritablePowerFeedStatusEnum) *NullableWritablePowerFeedStatusEnum {
	return &NullableWritablePowerFeedStatusEnum{value: val, isSet: true}
}

func (v NullableWritablePowerFeedStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritablePowerFeedStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
