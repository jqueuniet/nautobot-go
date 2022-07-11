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

// WritableVirtualMachineWithConfigContextStatusEnum the model 'WritableVirtualMachineWithConfigContextStatusEnum'
type WritableVirtualMachineWithConfigContextStatusEnum string

// List of WritableVirtualMachineWithConfigContextStatusEnum
const (
	ACTIVE WritableVirtualMachineWithConfigContextStatusEnum = "active"
	DECOMMISSIONING WritableVirtualMachineWithConfigContextStatusEnum = "decommissioning"
	FAILED WritableVirtualMachineWithConfigContextStatusEnum = "failed"
	OFFLINE WritableVirtualMachineWithConfigContextStatusEnum = "offline"
	PLANNED WritableVirtualMachineWithConfigContextStatusEnum = "planned"
	STAGED WritableVirtualMachineWithConfigContextStatusEnum = "staged"
)

// All allowed values of WritableVirtualMachineWithConfigContextStatusEnum enum
var AllowedWritableVirtualMachineWithConfigContextStatusEnumEnumValues = []WritableVirtualMachineWithConfigContextStatusEnum{
	"active",
	"decommissioning",
	"failed",
	"offline",
	"planned",
	"staged",
}

func (v *WritableVirtualMachineWithConfigContextStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WritableVirtualMachineWithConfigContextStatusEnum(value)
	for _, existing := range AllowedWritableVirtualMachineWithConfigContextStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WritableVirtualMachineWithConfigContextStatusEnum", value)
}

// NewWritableVirtualMachineWithConfigContextStatusEnumFromValue returns a pointer to a valid WritableVirtualMachineWithConfigContextStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWritableVirtualMachineWithConfigContextStatusEnumFromValue(v string) (*WritableVirtualMachineWithConfigContextStatusEnum, error) {
	ev := WritableVirtualMachineWithConfigContextStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WritableVirtualMachineWithConfigContextStatusEnum: valid values are %v", v, AllowedWritableVirtualMachineWithConfigContextStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WritableVirtualMachineWithConfigContextStatusEnum) IsValid() bool {
	for _, existing := range AllowedWritableVirtualMachineWithConfigContextStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WritableVirtualMachineWithConfigContextStatusEnum value
func (v WritableVirtualMachineWithConfigContextStatusEnum) Ptr() *WritableVirtualMachineWithConfigContextStatusEnum {
	return &v
}

type NullableWritableVirtualMachineWithConfigContextStatusEnum struct {
	value *WritableVirtualMachineWithConfigContextStatusEnum
	isSet bool
}

func (v NullableWritableVirtualMachineWithConfigContextStatusEnum) Get() *WritableVirtualMachineWithConfigContextStatusEnum {
	return v.value
}

func (v *NullableWritableVirtualMachineWithConfigContextStatusEnum) Set(val *WritableVirtualMachineWithConfigContextStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableVirtualMachineWithConfigContextStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableVirtualMachineWithConfigContextStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableVirtualMachineWithConfigContextStatusEnum(val *WritableVirtualMachineWithConfigContextStatusEnum) *NullableWritableVirtualMachineWithConfigContextStatusEnum {
	return &NullableWritableVirtualMachineWithConfigContextStatusEnum{value: val, isSet: true}
}

func (v NullableWritableVirtualMachineWithConfigContextStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableVirtualMachineWithConfigContextStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

