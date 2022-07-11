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

// CircuitTerminationProviderNetwork struct for CircuitTerminationProviderNetwork
type CircuitTerminationProviderNetwork struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Slug *string `json:"slug,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewCircuitTerminationProviderNetwork instantiates a new CircuitTerminationProviderNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCircuitTerminationProviderNetwork(id string, url string, name string, display string) *CircuitTerminationProviderNetwork {
	this := CircuitTerminationProviderNetwork{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.Display = display
	return &this
}

// NewCircuitTerminationProviderNetworkWithDefaults instantiates a new CircuitTerminationProviderNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCircuitTerminationProviderNetworkWithDefaults() *CircuitTerminationProviderNetwork {
	this := CircuitTerminationProviderNetwork{}
	return &this
}

// GetId returns the Id field value
func (o *CircuitTerminationProviderNetwork) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CircuitTerminationProviderNetwork) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CircuitTerminationProviderNetwork) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *CircuitTerminationProviderNetwork) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CircuitTerminationProviderNetwork) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CircuitTerminationProviderNetwork) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *CircuitTerminationProviderNetwork) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CircuitTerminationProviderNetwork) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CircuitTerminationProviderNetwork) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *CircuitTerminationProviderNetwork) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTerminationProviderNetwork) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *CircuitTerminationProviderNetwork) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *CircuitTerminationProviderNetwork) SetSlug(v string) {
	o.Slug = &v
}

// GetDisplay returns the Display field value
func (o *CircuitTerminationProviderNetwork) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *CircuitTerminationProviderNetwork) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *CircuitTerminationProviderNetwork) SetDisplay(v string) {
	o.Display = v
}

func (o CircuitTerminationProviderNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableCircuitTerminationProviderNetwork struct {
	value *CircuitTerminationProviderNetwork
	isSet bool
}

func (v NullableCircuitTerminationProviderNetwork) Get() *CircuitTerminationProviderNetwork {
	return v.value
}

func (v *NullableCircuitTerminationProviderNetwork) Set(val *CircuitTerminationProviderNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuitTerminationProviderNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuitTerminationProviderNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuitTerminationProviderNetwork(val *CircuitTerminationProviderNetwork) *NullableCircuitTerminationProviderNetwork {
	return &NullableCircuitTerminationProviderNetwork{value: val, isSet: true}
}

func (v NullableCircuitTerminationProviderNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuitTerminationProviderNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


