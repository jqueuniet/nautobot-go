/*
API Documentation

Source of truth and network automation platform

API version: 1.3.7 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// NestedSecretsGroupAssociation Returns a nested representation of an object on read, but accepts either the nested representation or the primary key value on write operations.
type NestedSecretsGroupAssociation struct {
	Id string `json:"id"`
	Url string `json:"url"`
	AccessType AccessTypeEnum `json:"access_type"`
	SecretType SecretTypeEnum `json:"secret_type"`
	Secret NestedSecret `json:"secret"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewNestedSecretsGroupAssociation instantiates a new NestedSecretsGroupAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedSecretsGroupAssociation(id string, url string, accessType AccessTypeEnum, secretType SecretTypeEnum, secret NestedSecret, display string) *NestedSecretsGroupAssociation {
	this := NestedSecretsGroupAssociation{}
	this.Id = id
	this.Url = url
	this.AccessType = accessType
	this.SecretType = secretType
	this.Secret = secret
	this.Display = display
	return &this
}

// NewNestedSecretsGroupAssociationWithDefaults instantiates a new NestedSecretsGroupAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedSecretsGroupAssociationWithDefaults() *NestedSecretsGroupAssociation {
	this := NestedSecretsGroupAssociation{}
	return &this
}

// GetId returns the Id field value
func (o *NestedSecretsGroupAssociation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedSecretsGroupAssociation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedSecretsGroupAssociation) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedSecretsGroupAssociation) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedSecretsGroupAssociation) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedSecretsGroupAssociation) SetUrl(v string) {
	o.Url = v
}

// GetAccessType returns the AccessType field value
func (o *NestedSecretsGroupAssociation) GetAccessType() AccessTypeEnum {
	if o == nil {
		var ret AccessTypeEnum
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *NestedSecretsGroupAssociation) GetAccessTypeOk() (*AccessTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *NestedSecretsGroupAssociation) SetAccessType(v AccessTypeEnum) {
	o.AccessType = v
}

// GetSecretType returns the SecretType field value
func (o *NestedSecretsGroupAssociation) GetSecretType() SecretTypeEnum {
	if o == nil {
		var ret SecretTypeEnum
		return ret
	}

	return o.SecretType
}

// GetSecretTypeOk returns a tuple with the SecretType field value
// and a boolean to check if the value has been set.
func (o *NestedSecretsGroupAssociation) GetSecretTypeOk() (*SecretTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretType, true
}

// SetSecretType sets field value
func (o *NestedSecretsGroupAssociation) SetSecretType(v SecretTypeEnum) {
	o.SecretType = v
}

// GetSecret returns the Secret field value
func (o *NestedSecretsGroupAssociation) GetSecret() NestedSecret {
	if o == nil {
		var ret NestedSecret
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *NestedSecretsGroupAssociation) GetSecretOk() (*NestedSecret, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *NestedSecretsGroupAssociation) SetSecret(v NestedSecret) {
	o.Secret = v
}

// GetDisplay returns the Display field value
func (o *NestedSecretsGroupAssociation) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedSecretsGroupAssociation) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedSecretsGroupAssociation) SetDisplay(v string) {
	o.Display = v
}

func (o NestedSecretsGroupAssociation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["access_type"] = o.AccessType
	}
	if true {
		toSerialize["secret_type"] = o.SecretType
	}
	if true {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableNestedSecretsGroupAssociation struct {
	value *NestedSecretsGroupAssociation
	isSet bool
}

func (v NullableNestedSecretsGroupAssociation) Get() *NestedSecretsGroupAssociation {
	return v.value
}

func (v *NullableNestedSecretsGroupAssociation) Set(val *NestedSecretsGroupAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedSecretsGroupAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedSecretsGroupAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedSecretsGroupAssociation(val *NestedSecretsGroupAssociation) *NullableNestedSecretsGroupAssociation {
	return &NullableNestedSecretsGroupAssociation{value: val, isSet: true}
}

func (v NullableNestedSecretsGroupAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedSecretsGroupAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


