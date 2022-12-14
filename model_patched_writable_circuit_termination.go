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

// PatchedWritableCircuitTermination Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritableCircuitTermination struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Circuit *string `json:"circuit,omitempty"`
	TermSide NullableTermSideEnum `json:"term_side,omitempty"`
	Site NullableString `json:"site,omitempty"`
	ProviderNetwork NullableString `json:"provider_network,omitempty"`
	PortSpeed NullableInt32 `json:"port_speed,omitempty"`
	// Upstream speed, if different from port speed
	UpstreamSpeed NullableInt32 `json:"upstream_speed,omitempty"`
	XconnectId *string `json:"xconnect_id,omitempty"`
	PpInfo *string `json:"pp_info,omitempty"`
	Description *string `json:"description,omitempty"`
	Cable *CircuitTerminationCable `json:"cable,omitempty"`
	CablePeer map[string]interface{} `json:"cable_peer,omitempty"`
	CablePeerType NullableString `json:"cable_peer_type,omitempty"`
	ConnectedEndpoint map[string]interface{} `json:"connected_endpoint,omitempty"`
	ConnectedEndpointType NullableString `json:"connected_endpoint_type,omitempty"`
	ConnectedEndpointReachable NullableBool `json:"connected_endpoint_reachable,omitempty"`
	// Human friendly display value
	Display *string `json:"display,omitempty"`
}

// NewPatchedWritableCircuitTermination instantiates a new PatchedWritableCircuitTermination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableCircuitTermination() *PatchedWritableCircuitTermination {
	this := PatchedWritableCircuitTermination{}
	return &this
}

// NewPatchedWritableCircuitTerminationWithDefaults instantiates a new PatchedWritableCircuitTermination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableCircuitTerminationWithDefaults() *PatchedWritableCircuitTermination {
	this := PatchedWritableCircuitTermination{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTermination) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTermination) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedWritableCircuitTermination) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTermination) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTermination) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedWritableCircuitTermination) SetUrl(v string) {
	o.Url = &v
}

// GetCircuit returns the Circuit field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTermination) GetCircuit() string {
	if o == nil || o.Circuit == nil {
		var ret string
		return ret
	}
	return *o.Circuit
}

// GetCircuitOk returns a tuple with the Circuit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTermination) GetCircuitOk() (*string, bool) {
	if o == nil || o.Circuit == nil {
		return nil, false
	}
	return o.Circuit, true
}

// HasCircuit returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasCircuit() bool {
	if o != nil && o.Circuit != nil {
		return true
	}

	return false
}

// SetCircuit gets a reference to the given string and assigns it to the Circuit field.
func (o *PatchedWritableCircuitTermination) SetCircuit(v string) {
	o.Circuit = &v
}

// GetTermSide returns the TermSide field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitTermination) GetTermSide() TermSideEnum {
	if o == nil || o.TermSide.Get() == nil {
		var ret TermSideEnum
		return ret
	}
	return *o.TermSide.Get()
}

// GetTermSideOk returns a tuple with the TermSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitTermination) GetTermSideOk() (*TermSideEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.TermSide.Get(), o.TermSide.IsSet()
}

// HasTermSide returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasTermSide() bool {
	if o != nil && o.TermSide.IsSet() {
		return true
	}

	return false
}

// SetTermSide gets a reference to the given NullableTermSideEnum and assigns it to the TermSide field.
func (o *PatchedWritableCircuitTermination) SetTermSide(v TermSideEnum) {
	o.TermSide.Set(&v)
}
// SetTermSideNil sets the value for TermSide to be an explicit nil
func (o *PatchedWritableCircuitTermination) SetTermSideNil() {
	o.TermSide.Set(nil)
}

// UnsetTermSide ensures that no value is present for TermSide, not even an explicit nil
func (o *PatchedWritableCircuitTermination) UnsetTermSide() {
	o.TermSide.Unset()
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitTermination) GetSite() string {
	if o == nil || o.Site.Get() == nil {
		var ret string
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitTermination) GetSiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableString and assigns it to the Site field.
func (o *PatchedWritableCircuitTermination) SetSite(v string) {
	o.Site.Set(&v)
}
// SetSiteNil sets the value for Site to be an explicit nil
func (o *PatchedWritableCircuitTermination) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *PatchedWritableCircuitTermination) UnsetSite() {
	o.Site.Unset()
}

// GetProviderNetwork returns the ProviderNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitTermination) GetProviderNetwork() string {
	if o == nil || o.ProviderNetwork.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProviderNetwork.Get()
}

// GetProviderNetworkOk returns a tuple with the ProviderNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitTermination) GetProviderNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderNetwork.Get(), o.ProviderNetwork.IsSet()
}

// HasProviderNetwork returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasProviderNetwork() bool {
	if o != nil && o.ProviderNetwork.IsSet() {
		return true
	}

	return false
}

// SetProviderNetwork gets a reference to the given NullableString and assigns it to the ProviderNetwork field.
func (o *PatchedWritableCircuitTermination) SetProviderNetwork(v string) {
	o.ProviderNetwork.Set(&v)
}
// SetProviderNetworkNil sets the value for ProviderNetwork to be an explicit nil
func (o *PatchedWritableCircuitTermination) SetProviderNetworkNil() {
	o.ProviderNetwork.Set(nil)
}

// UnsetProviderNetwork ensures that no value is present for ProviderNetwork, not even an explicit nil
func (o *PatchedWritableCircuitTermination) UnsetProviderNetwork() {
	o.ProviderNetwork.Unset()
}

// GetPortSpeed returns the PortSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitTermination) GetPortSpeed() int32 {
	if o == nil || o.PortSpeed.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PortSpeed.Get()
}

// GetPortSpeedOk returns a tuple with the PortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitTermination) GetPortSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortSpeed.Get(), o.PortSpeed.IsSet()
}

// HasPortSpeed returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasPortSpeed() bool {
	if o != nil && o.PortSpeed.IsSet() {
		return true
	}

	return false
}

// SetPortSpeed gets a reference to the given NullableInt32 and assigns it to the PortSpeed field.
func (o *PatchedWritableCircuitTermination) SetPortSpeed(v int32) {
	o.PortSpeed.Set(&v)
}
// SetPortSpeedNil sets the value for PortSpeed to be an explicit nil
func (o *PatchedWritableCircuitTermination) SetPortSpeedNil() {
	o.PortSpeed.Set(nil)
}

// UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
func (o *PatchedWritableCircuitTermination) UnsetPortSpeed() {
	o.PortSpeed.Unset()
}

// GetUpstreamSpeed returns the UpstreamSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitTermination) GetUpstreamSpeed() int32 {
	if o == nil || o.UpstreamSpeed.Get() == nil {
		var ret int32
		return ret
	}
	return *o.UpstreamSpeed.Get()
}

// GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitTermination) GetUpstreamSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpstreamSpeed.Get(), o.UpstreamSpeed.IsSet()
}

// HasUpstreamSpeed returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasUpstreamSpeed() bool {
	if o != nil && o.UpstreamSpeed.IsSet() {
		return true
	}

	return false
}

// SetUpstreamSpeed gets a reference to the given NullableInt32 and assigns it to the UpstreamSpeed field.
func (o *PatchedWritableCircuitTermination) SetUpstreamSpeed(v int32) {
	o.UpstreamSpeed.Set(&v)
}
// SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil
func (o *PatchedWritableCircuitTermination) SetUpstreamSpeedNil() {
	o.UpstreamSpeed.Set(nil)
}

// UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
func (o *PatchedWritableCircuitTermination) UnsetUpstreamSpeed() {
	o.UpstreamSpeed.Unset()
}

// GetXconnectId returns the XconnectId field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTermination) GetXconnectId() string {
	if o == nil || o.XconnectId == nil {
		var ret string
		return ret
	}
	return *o.XconnectId
}

// GetXconnectIdOk returns a tuple with the XconnectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTermination) GetXconnectIdOk() (*string, bool) {
	if o == nil || o.XconnectId == nil {
		return nil, false
	}
	return o.XconnectId, true
}

// HasXconnectId returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasXconnectId() bool {
	if o != nil && o.XconnectId != nil {
		return true
	}

	return false
}

// SetXconnectId gets a reference to the given string and assigns it to the XconnectId field.
func (o *PatchedWritableCircuitTermination) SetXconnectId(v string) {
	o.XconnectId = &v
}

// GetPpInfo returns the PpInfo field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTermination) GetPpInfo() string {
	if o == nil || o.PpInfo == nil {
		var ret string
		return ret
	}
	return *o.PpInfo
}

// GetPpInfoOk returns a tuple with the PpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTermination) GetPpInfoOk() (*string, bool) {
	if o == nil || o.PpInfo == nil {
		return nil, false
	}
	return o.PpInfo, true
}

// HasPpInfo returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasPpInfo() bool {
	if o != nil && o.PpInfo != nil {
		return true
	}

	return false
}

// SetPpInfo gets a reference to the given string and assigns it to the PpInfo field.
func (o *PatchedWritableCircuitTermination) SetPpInfo(v string) {
	o.PpInfo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTermination) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTermination) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableCircuitTermination) SetDescription(v string) {
	o.Description = &v
}

// GetCable returns the Cable field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTermination) GetCable() CircuitTerminationCable {
	if o == nil || o.Cable == nil {
		var ret CircuitTerminationCable
		return ret
	}
	return *o.Cable
}

// GetCableOk returns a tuple with the Cable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTermination) GetCableOk() (*CircuitTerminationCable, bool) {
	if o == nil || o.Cable == nil {
		return nil, false
	}
	return o.Cable, true
}

// HasCable returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasCable() bool {
	if o != nil && o.Cable != nil {
		return true
	}

	return false
}

// SetCable gets a reference to the given CircuitTerminationCable and assigns it to the Cable field.
func (o *PatchedWritableCircuitTermination) SetCable(v CircuitTerminationCable) {
	o.Cable = &v
}

// GetCablePeer returns the CablePeer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitTermination) GetCablePeer() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CablePeer
}

// GetCablePeerOk returns a tuple with the CablePeer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitTermination) GetCablePeerOk() (map[string]interface{}, bool) {
	if o == nil || o.CablePeer == nil {
		return nil, false
	}
	return o.CablePeer, true
}

// HasCablePeer returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasCablePeer() bool {
	if o != nil && o.CablePeer != nil {
		return true
	}

	return false
}

// SetCablePeer gets a reference to the given map[string]interface{} and assigns it to the CablePeer field.
func (o *PatchedWritableCircuitTermination) SetCablePeer(v map[string]interface{}) {
	o.CablePeer = v
}

// GetCablePeerType returns the CablePeerType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitTermination) GetCablePeerType() string {
	if o == nil || o.CablePeerType.Get() == nil {
		var ret string
		return ret
	}
	return *o.CablePeerType.Get()
}

// GetCablePeerTypeOk returns a tuple with the CablePeerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitTermination) GetCablePeerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CablePeerType.Get(), o.CablePeerType.IsSet()
}

// HasCablePeerType returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasCablePeerType() bool {
	if o != nil && o.CablePeerType.IsSet() {
		return true
	}

	return false
}

// SetCablePeerType gets a reference to the given NullableString and assigns it to the CablePeerType field.
func (o *PatchedWritableCircuitTermination) SetCablePeerType(v string) {
	o.CablePeerType.Set(&v)
}
// SetCablePeerTypeNil sets the value for CablePeerType to be an explicit nil
func (o *PatchedWritableCircuitTermination) SetCablePeerTypeNil() {
	o.CablePeerType.Set(nil)
}

// UnsetCablePeerType ensures that no value is present for CablePeerType, not even an explicit nil
func (o *PatchedWritableCircuitTermination) UnsetCablePeerType() {
	o.CablePeerType.Unset()
}

// GetConnectedEndpoint returns the ConnectedEndpoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitTermination) GetConnectedEndpoint() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ConnectedEndpoint
}

// GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitTermination) GetConnectedEndpointOk() (map[string]interface{}, bool) {
	if o == nil || o.ConnectedEndpoint == nil {
		return nil, false
	}
	return o.ConnectedEndpoint, true
}

// HasConnectedEndpoint returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasConnectedEndpoint() bool {
	if o != nil && o.ConnectedEndpoint != nil {
		return true
	}

	return false
}

// SetConnectedEndpoint gets a reference to the given map[string]interface{} and assigns it to the ConnectedEndpoint field.
func (o *PatchedWritableCircuitTermination) SetConnectedEndpoint(v map[string]interface{}) {
	o.ConnectedEndpoint = v
}

// GetConnectedEndpointType returns the ConnectedEndpointType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitTermination) GetConnectedEndpointType() string {
	if o == nil || o.ConnectedEndpointType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConnectedEndpointType.Get()
}

// GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitTermination) GetConnectedEndpointTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointType.Get(), o.ConnectedEndpointType.IsSet()
}

// HasConnectedEndpointType returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasConnectedEndpointType() bool {
	if o != nil && o.ConnectedEndpointType.IsSet() {
		return true
	}

	return false
}

// SetConnectedEndpointType gets a reference to the given NullableString and assigns it to the ConnectedEndpointType field.
func (o *PatchedWritableCircuitTermination) SetConnectedEndpointType(v string) {
	o.ConnectedEndpointType.Set(&v)
}
// SetConnectedEndpointTypeNil sets the value for ConnectedEndpointType to be an explicit nil
func (o *PatchedWritableCircuitTermination) SetConnectedEndpointTypeNil() {
	o.ConnectedEndpointType.Set(nil)
}

// UnsetConnectedEndpointType ensures that no value is present for ConnectedEndpointType, not even an explicit nil
func (o *PatchedWritableCircuitTermination) UnsetConnectedEndpointType() {
	o.ConnectedEndpointType.Unset()
}

// GetConnectedEndpointReachable returns the ConnectedEndpointReachable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuitTermination) GetConnectedEndpointReachable() bool {
	if o == nil || o.ConnectedEndpointReachable.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ConnectedEndpointReachable.Get()
}

// GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuitTermination) GetConnectedEndpointReachableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointReachable.Get(), o.ConnectedEndpointReachable.IsSet()
}

// HasConnectedEndpointReachable returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasConnectedEndpointReachable() bool {
	if o != nil && o.ConnectedEndpointReachable.IsSet() {
		return true
	}

	return false
}

// SetConnectedEndpointReachable gets a reference to the given NullableBool and assigns it to the ConnectedEndpointReachable field.
func (o *PatchedWritableCircuitTermination) SetConnectedEndpointReachable(v bool) {
	o.ConnectedEndpointReachable.Set(&v)
}
// SetConnectedEndpointReachableNil sets the value for ConnectedEndpointReachable to be an explicit nil
func (o *PatchedWritableCircuitTermination) SetConnectedEndpointReachableNil() {
	o.ConnectedEndpointReachable.Set(nil)
}

// UnsetConnectedEndpointReachable ensures that no value is present for ConnectedEndpointReachable, not even an explicit nil
func (o *PatchedWritableCircuitTermination) UnsetConnectedEndpointReachable() {
	o.ConnectedEndpointReachable.Unset()
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PatchedWritableCircuitTermination) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuitTermination) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PatchedWritableCircuitTermination) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PatchedWritableCircuitTermination) SetDisplay(v string) {
	o.Display = &v
}

func (o PatchedWritableCircuitTermination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Circuit != nil {
		toSerialize["circuit"] = o.Circuit
	}
	if o.TermSide.IsSet() {
		toSerialize["term_side"] = o.TermSide.Get()
	}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.ProviderNetwork.IsSet() {
		toSerialize["provider_network"] = o.ProviderNetwork.Get()
	}
	if o.PortSpeed.IsSet() {
		toSerialize["port_speed"] = o.PortSpeed.Get()
	}
	if o.UpstreamSpeed.IsSet() {
		toSerialize["upstream_speed"] = o.UpstreamSpeed.Get()
	}
	if o.XconnectId != nil {
		toSerialize["xconnect_id"] = o.XconnectId
	}
	if o.PpInfo != nil {
		toSerialize["pp_info"] = o.PpInfo
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Cable != nil {
		toSerialize["cable"] = o.Cable
	}
	if o.CablePeer != nil {
		toSerialize["cable_peer"] = o.CablePeer
	}
	if o.CablePeerType.IsSet() {
		toSerialize["cable_peer_type"] = o.CablePeerType.Get()
	}
	if o.ConnectedEndpoint != nil {
		toSerialize["connected_endpoint"] = o.ConnectedEndpoint
	}
	if o.ConnectedEndpointType.IsSet() {
		toSerialize["connected_endpoint_type"] = o.ConnectedEndpointType.Get()
	}
	if o.ConnectedEndpointReachable.IsSet() {
		toSerialize["connected_endpoint_reachable"] = o.ConnectedEndpointReachable.Get()
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedWritableCircuitTermination struct {
	value *PatchedWritableCircuitTermination
	isSet bool
}

func (v NullablePatchedWritableCircuitTermination) Get() *PatchedWritableCircuitTermination {
	return v.value
}

func (v *NullablePatchedWritableCircuitTermination) Set(val *PatchedWritableCircuitTermination) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCircuitTermination) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCircuitTermination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCircuitTermination(val *PatchedWritableCircuitTermination) *NullablePatchedWritableCircuitTermination {
	return &NullablePatchedWritableCircuitTermination{value: val, isSet: true}
}

func (v NullablePatchedWritableCircuitTermination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCircuitTermination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


