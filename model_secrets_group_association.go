/*
API Documentation

Source of truth and network automation platform

API version: 1.3.10b1 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SecretsGroupAssociation Serializer for `SecretsGroupAssociation` objects.
type SecretsGroupAssociation struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Group NestedSecretsGroup `json:"group"`
	AccessType AccessTypeEnum `json:"access_type"`
	SecretType SecretTypeEnum `json:"secret_type"`
	Secret NestedSecret `json:"secret"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewSecretsGroupAssociation instantiates a new SecretsGroupAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretsGroupAssociation(id string, url string, group NestedSecretsGroup, accessType AccessTypeEnum, secretType SecretTypeEnum, secret NestedSecret, display string) *SecretsGroupAssociation {
	this := SecretsGroupAssociation{}
	this.Id = id
	this.Url = url
	this.Group = group
	this.AccessType = accessType
	this.SecretType = secretType
	this.Secret = secret
	this.Display = display
	return &this
}

// NewSecretsGroupAssociationWithDefaults instantiates a new SecretsGroupAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretsGroupAssociationWithDefaults() *SecretsGroupAssociation {
	this := SecretsGroupAssociation{}
	return &this
}

// GetId returns the Id field value
func (o *SecretsGroupAssociation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecretsGroupAssociation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecretsGroupAssociation) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *SecretsGroupAssociation) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SecretsGroupAssociation) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SecretsGroupAssociation) SetUrl(v string) {
	o.Url = v
}

// GetGroup returns the Group field value
func (o *SecretsGroupAssociation) GetGroup() NestedSecretsGroup {
	if o == nil {
		var ret NestedSecretsGroup
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *SecretsGroupAssociation) GetGroupOk() (*NestedSecretsGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *SecretsGroupAssociation) SetGroup(v NestedSecretsGroup) {
	o.Group = v
}

// GetAccessType returns the AccessType field value
func (o *SecretsGroupAssociation) GetAccessType() AccessTypeEnum {
	if o == nil {
		var ret AccessTypeEnum
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *SecretsGroupAssociation) GetAccessTypeOk() (*AccessTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *SecretsGroupAssociation) SetAccessType(v AccessTypeEnum) {
	o.AccessType = v
}

// GetSecretType returns the SecretType field value
func (o *SecretsGroupAssociation) GetSecretType() SecretTypeEnum {
	if o == nil {
		var ret SecretTypeEnum
		return ret
	}

	return o.SecretType
}

// GetSecretTypeOk returns a tuple with the SecretType field value
// and a boolean to check if the value has been set.
func (o *SecretsGroupAssociation) GetSecretTypeOk() (*SecretTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretType, true
}

// SetSecretType sets field value
func (o *SecretsGroupAssociation) SetSecretType(v SecretTypeEnum) {
	o.SecretType = v
}

// GetSecret returns the Secret field value
func (o *SecretsGroupAssociation) GetSecret() NestedSecret {
	if o == nil {
		var ret NestedSecret
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *SecretsGroupAssociation) GetSecretOk() (*NestedSecret, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *SecretsGroupAssociation) SetSecret(v NestedSecret) {
	o.Secret = v
}

// GetDisplay returns the Display field value
func (o *SecretsGroupAssociation) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *SecretsGroupAssociation) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *SecretsGroupAssociation) SetDisplay(v string) {
	o.Display = v
}

func (o SecretsGroupAssociation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["group"] = o.Group
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

type NullableSecretsGroupAssociation struct {
	value *SecretsGroupAssociation
	isSet bool
}

func (v NullableSecretsGroupAssociation) Get() *SecretsGroupAssociation {
	return v.value
}

func (v *NullableSecretsGroupAssociation) Set(val *SecretsGroupAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretsGroupAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretsGroupAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretsGroupAssociation(val *SecretsGroupAssociation) *NullableSecretsGroupAssociation {
	return &NullableSecretsGroupAssociation{value: val, isSet: true}
}

func (v NullableSecretsGroupAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretsGroupAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


