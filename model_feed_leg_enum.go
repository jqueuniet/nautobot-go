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

// FeedLegEnum the model 'FeedLegEnum'
type FeedLegEnum string

// List of FeedLegEnum
const (
	//A FeedLegEnum = "A"
	B FeedLegEnum = "B"
	C FeedLegEnum = "C"
)

// All allowed values of FeedLegEnum enum
var AllowedFeedLegEnumEnumValues = []FeedLegEnum{
	"A",
	"B",
	"C",
}

func (v *FeedLegEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeedLegEnum(value)
	for _, existing := range AllowedFeedLegEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FeedLegEnum", value)
}

// NewFeedLegEnumFromValue returns a pointer to a valid FeedLegEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeedLegEnumFromValue(v string) (*FeedLegEnum, error) {
	ev := FeedLegEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeedLegEnum: valid values are %v", v, AllowedFeedLegEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeedLegEnum) IsValid() bool {
	for _, existing := range AllowedFeedLegEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeedLegEnum value
func (v FeedLegEnum) Ptr() *FeedLegEnum {
	return &v
}

type NullableFeedLegEnum struct {
	value *FeedLegEnum
	isSet bool
}

func (v NullableFeedLegEnum) Get() *FeedLegEnum {
	return v.value
}

func (v *NullableFeedLegEnum) Set(val *FeedLegEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedLegEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedLegEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedLegEnum(val *FeedLegEnum) *NullableFeedLegEnum {
	return &NullableFeedLegEnum{value: val, isSet: true}
}

func (v NullableFeedLegEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedLegEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
