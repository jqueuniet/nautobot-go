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

// WritableDeviceWithConfigContextStatusEnum the model 'WritableDeviceWithConfigContextStatusEnum'
type WritableDeviceWithConfigContextStatusEnum string

// List of WritableDeviceWithConfigContextStatusEnum
const (
	ACTIVE WritableDeviceWithConfigContextStatusEnum = "active"
	DECOMMISSIONING WritableDeviceWithConfigContextStatusEnum = "decommissioning"
	FAILED WritableDeviceWithConfigContextStatusEnum = "failed"
	INVENTORY WritableDeviceWithConfigContextStatusEnum = "inventory"
	OFFLINE WritableDeviceWithConfigContextStatusEnum = "offline"
	PLANNED WritableDeviceWithConfigContextStatusEnum = "planned"
	STAGED WritableDeviceWithConfigContextStatusEnum = "staged"
)

// All allowed values of WritableDeviceWithConfigContextStatusEnum enum
var AllowedWritableDeviceWithConfigContextStatusEnumEnumValues = []WritableDeviceWithConfigContextStatusEnum{
	"active",
	"decommissioning",
	"failed",
	"inventory",
	"offline",
	"planned",
	"staged",
}

func (v *WritableDeviceWithConfigContextStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WritableDeviceWithConfigContextStatusEnum(value)
	for _, existing := range AllowedWritableDeviceWithConfigContextStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WritableDeviceWithConfigContextStatusEnum", value)
}

// NewWritableDeviceWithConfigContextStatusEnumFromValue returns a pointer to a valid WritableDeviceWithConfigContextStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWritableDeviceWithConfigContextStatusEnumFromValue(v string) (*WritableDeviceWithConfigContextStatusEnum, error) {
	ev := WritableDeviceWithConfigContextStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WritableDeviceWithConfigContextStatusEnum: valid values are %v", v, AllowedWritableDeviceWithConfigContextStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WritableDeviceWithConfigContextStatusEnum) IsValid() bool {
	for _, existing := range AllowedWritableDeviceWithConfigContextStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WritableDeviceWithConfigContextStatusEnum value
func (v WritableDeviceWithConfigContextStatusEnum) Ptr() *WritableDeviceWithConfigContextStatusEnum {
	return &v
}

type NullableWritableDeviceWithConfigContextStatusEnum struct {
	value *WritableDeviceWithConfigContextStatusEnum
	isSet bool
}

func (v NullableWritableDeviceWithConfigContextStatusEnum) Get() *WritableDeviceWithConfigContextStatusEnum {
	return v.value
}

func (v *NullableWritableDeviceWithConfigContextStatusEnum) Set(val *WritableDeviceWithConfigContextStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableDeviceWithConfigContextStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableDeviceWithConfigContextStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableDeviceWithConfigContextStatusEnum(val *WritableDeviceWithConfigContextStatusEnum) *NullableWritableDeviceWithConfigContextStatusEnum {
	return &NullableWritableDeviceWithConfigContextStatusEnum{value: val, isSet: true}
}

func (v NullableWritableDeviceWithConfigContextStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableDeviceWithConfigContextStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

