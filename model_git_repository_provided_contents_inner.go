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

// GitRepositoryProvidedContentsInner - struct for GitRepositoryProvidedContentsInner
type GitRepositoryProvidedContentsInner struct {
	BlankEnum *BlankEnum
	ProvidedContentsEnum *ProvidedContentsEnum
}

// BlankEnumAsGitRepositoryProvidedContentsInner is a convenience function that returns BlankEnum wrapped in GitRepositoryProvidedContentsInner
func BlankEnumAsGitRepositoryProvidedContentsInner(v *BlankEnum) GitRepositoryProvidedContentsInner {
	return GitRepositoryProvidedContentsInner{
		BlankEnum: v,
	}
}

// ProvidedContentsEnumAsGitRepositoryProvidedContentsInner is a convenience function that returns ProvidedContentsEnum wrapped in GitRepositoryProvidedContentsInner
func ProvidedContentsEnumAsGitRepositoryProvidedContentsInner(v *ProvidedContentsEnum) GitRepositoryProvidedContentsInner {
	return GitRepositoryProvidedContentsInner{
		ProvidedContentsEnum: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GitRepositoryProvidedContentsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BlankEnum
	err = newStrictDecoder(data).Decode(&dst.BlankEnum)
	if err == nil {
		jsonBlankEnum, _ := json.Marshal(dst.BlankEnum)
		if string(jsonBlankEnum) == "{}" { // empty struct
			dst.BlankEnum = nil
		} else {
			match++
		}
	} else {
		dst.BlankEnum = nil
	}

	// try to unmarshal data into ProvidedContentsEnum
	err = newStrictDecoder(data).Decode(&dst.ProvidedContentsEnum)
	if err == nil {
		jsonProvidedContentsEnum, _ := json.Marshal(dst.ProvidedContentsEnum)
		if string(jsonProvidedContentsEnum) == "{}" { // empty struct
			dst.ProvidedContentsEnum = nil
		} else {
			match++
		}
	} else {
		dst.ProvidedContentsEnum = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.ProvidedContentsEnum = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GitRepositoryProvidedContentsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GitRepositoryProvidedContentsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GitRepositoryProvidedContentsInner) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.ProvidedContentsEnum != nil {
		return json.Marshal(&src.ProvidedContentsEnum)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GitRepositoryProvidedContentsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.ProvidedContentsEnum != nil {
		return obj.ProvidedContentsEnum
	}

	// all schemas are nil
	return nil
}

type NullableGitRepositoryProvidedContentsInner struct {
	value *GitRepositoryProvidedContentsInner
	isSet bool
}

func (v NullableGitRepositoryProvidedContentsInner) Get() *GitRepositoryProvidedContentsInner {
	return v.value
}

func (v *NullableGitRepositoryProvidedContentsInner) Set(val *GitRepositoryProvidedContentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGitRepositoryProvidedContentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGitRepositoryProvidedContentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitRepositoryProvidedContentsInner(val *GitRepositoryProvidedContentsInner) *NullableGitRepositoryProvidedContentsInner {
	return &NullableGitRepositoryProvidedContentsInner{value: val, isSet: true}
}

func (v NullableGitRepositoryProvidedContentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitRepositoryProvidedContentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


