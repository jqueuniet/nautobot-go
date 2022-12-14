/*
API Documentation

Source of truth and network automation platform

API version: 1.3.10b1 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// PowerFeed Mixin to add `status` choice field to model serializers.
type PowerFeed struct {
	Id         string             `json:"id"`
	Url        string             `json:"url"`
	PowerPanel NestedPowerPanel   `json:"power_panel"`
	Rack       NullableDeviceRack `json:"rack,omitempty"`
	Name       string             `json:"name"`
	Status     PowerFeedStatus    `json:"status"`
	Type       *PowerFeedType     `json:"type,omitempty"`
	Supply     *PowerFeedSupply   `json:"supply,omitempty"`
	Phase      *PowerFeedPhase    `json:"phase,omitempty"`
	Voltage    *int32             `json:"voltage,omitempty"`
	Amperage   *int32             `json:"amperage,omitempty"`
	// Maximum permissible draw (percentage)
	MaxUtilization             *int32                  `json:"max_utilization,omitempty"`
	Comments                   *string                 `json:"comments,omitempty"`
	Cable                      CircuitTerminationCable `json:"cable"`
	CablePeer                  map[string]interface{}  `json:"cable_peer"`
	CablePeerType              NullableString          `json:"cable_peer_type"`
	ConnectedEndpoint          map[string]interface{}  `json:"connected_endpoint"`
	ConnectedEndpointType      NullableString          `json:"connected_endpoint_type"`
	ConnectedEndpointReachable NullableBool            `json:"connected_endpoint_reachable"`
	Tags                       []TagSerializerField    `json:"tags,omitempty"`
	CustomFields               map[string]interface{}  `json:"custom_fields,omitempty"`
	Created                    string                  `json:"created"`
	LastUpdated                time.Time               `json:"last_updated"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewPowerFeed instantiates a new PowerFeed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerFeed(id string, url string, powerPanel NestedPowerPanel, name string, status PowerFeedStatus, cable CircuitTerminationCable, cablePeer map[string]interface{}, cablePeerType NullableString, connectedEndpoint map[string]interface{}, connectedEndpointType NullableString, connectedEndpointReachable NullableBool, created string, lastUpdated time.Time, display string) *PowerFeed {
	this := PowerFeed{}
	this.Id = id
	this.Url = url
	this.PowerPanel = powerPanel
	this.Name = name
	this.Status = status
	powerFeedValue := "primary"
	powerFreeLabel := "Primary"
	var type_ = PowerFeedType{Value: &powerFeedValue, Label: &powerFreeLabel}
	this.Type = &type_
	powerFeedSupplyValue := "ac"
	powerFeedSupplyLabel := "AC"
	var supply = PowerFeedSupply{Value: &powerFeedSupplyValue, Label: &powerFeedSupplyLabel}
	this.Supply = &supply
	powerFeedPhaseValue := "single-phase"
	powerFeedPhaseLabel := "Single phase"
	var phase = PowerFeedPhase{Value: &powerFeedPhaseValue, Label: &powerFeedPhaseLabel}
	this.Phase = &phase
	this.Cable = cable
	this.CablePeer = cablePeer
	this.CablePeerType = cablePeerType
	this.ConnectedEndpoint = connectedEndpoint
	this.ConnectedEndpointType = connectedEndpointType
	this.ConnectedEndpointReachable = connectedEndpointReachable
	this.Created = created
	this.LastUpdated = lastUpdated
	this.Display = display
	return &this
}

// NewPowerFeedWithDefaults instantiates a new PowerFeed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerFeedWithDefaults() *PowerFeed {
	this := PowerFeed{}
	powerFeedValue := "primary"
	powerFreeLabel := "Primary"
	var type_ = PowerFeedType{Value: &powerFeedValue, Label: &powerFreeLabel}
	this.Type = &type_
	powerFeedSupplyValue := "ac"
	powerFeedSupplyLabel := "AC"
	var supply = PowerFeedSupply{Value: &powerFeedSupplyValue, Label: &powerFeedSupplyLabel}
	this.Supply = &supply
	powerFeedPhaseValue := "single-phase"
	powerFeedPhaseLabel := "Single phase"
	var phase = PowerFeedPhase{Value: &powerFeedPhaseValue, Label: &powerFeedPhaseLabel}
	this.Phase = &phase
	return &this
}

// GetId returns the Id field value
func (o *PowerFeed) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PowerFeed) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *PowerFeed) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PowerFeed) SetUrl(v string) {
	o.Url = v
}

// GetPowerPanel returns the PowerPanel field value
func (o *PowerFeed) GetPowerPanel() NestedPowerPanel {
	if o == nil {
		var ret NestedPowerPanel
		return ret
	}

	return o.PowerPanel
}

// GetPowerPanelOk returns a tuple with the PowerPanel field value
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetPowerPanelOk() (*NestedPowerPanel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerPanel, true
}

// SetPowerPanel sets field value
func (o *PowerFeed) SetPowerPanel(v NestedPowerPanel) {
	o.PowerPanel = v
}

// GetRack returns the Rack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerFeed) GetRack() DeviceRack {
	if o == nil || o.Rack.Get() == nil {
		var ret DeviceRack
		return ret
	}
	return *o.Rack.Get()
}

// GetRackOk returns a tuple with the Rack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerFeed) GetRackOk() (*DeviceRack, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rack.Get(), o.Rack.IsSet()
}

// HasRack returns a boolean if a field has been set.
func (o *PowerFeed) HasRack() bool {
	if o != nil && o.Rack.IsSet() {
		return true
	}

	return false
}

// SetRack gets a reference to the given NullableDeviceRack and assigns it to the Rack field.
func (o *PowerFeed) SetRack(v DeviceRack) {
	o.Rack.Set(&v)
}

// SetRackNil sets the value for Rack to be an explicit nil
func (o *PowerFeed) SetRackNil() {
	o.Rack.Set(nil)
}

// UnsetRack ensures that no value is present for Rack, not even an explicit nil
func (o *PowerFeed) UnsetRack() {
	o.Rack.Unset()
}

// GetName returns the Name field value
func (o *PowerFeed) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PowerFeed) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *PowerFeed) GetStatus() PowerFeedStatus {
	if o == nil {
		var ret PowerFeedStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetStatusOk() (*PowerFeedStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PowerFeed) SetStatus(v PowerFeedStatus) {
	o.Status = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PowerFeed) GetType() PowerFeedType {
	if o == nil || o.Type == nil {
		var ret PowerFeedType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetTypeOk() (*PowerFeedType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PowerFeed) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given PowerFeedType and assigns it to the Type field.
func (o *PowerFeed) SetType(v PowerFeedType) {
	o.Type = &v
}

// GetSupply returns the Supply field value if set, zero value otherwise.
func (o *PowerFeed) GetSupply() PowerFeedSupply {
	if o == nil || o.Supply == nil {
		var ret PowerFeedSupply
		return ret
	}
	return *o.Supply
}

// GetSupplyOk returns a tuple with the Supply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetSupplyOk() (*PowerFeedSupply, bool) {
	if o == nil || o.Supply == nil {
		return nil, false
	}
	return o.Supply, true
}

// HasSupply returns a boolean if a field has been set.
func (o *PowerFeed) HasSupply() bool {
	if o != nil && o.Supply != nil {
		return true
	}

	return false
}

// SetSupply gets a reference to the given PowerFeedSupply and assigns it to the Supply field.
func (o *PowerFeed) SetSupply(v PowerFeedSupply) {
	o.Supply = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *PowerFeed) GetPhase() PowerFeedPhase {
	if o == nil || o.Phase == nil {
		var ret PowerFeedPhase
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetPhaseOk() (*PowerFeedPhase, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *PowerFeed) HasPhase() bool {
	if o != nil && o.Phase != nil {
		return true
	}

	return false
}

// SetPhase gets a reference to the given PowerFeedPhase and assigns it to the Phase field.
func (o *PowerFeed) SetPhase(v PowerFeedPhase) {
	o.Phase = &v
}

// GetVoltage returns the Voltage field value if set, zero value otherwise.
func (o *PowerFeed) GetVoltage() int32 {
	if o == nil || o.Voltage == nil {
		var ret int32
		return ret
	}
	return *o.Voltage
}

// GetVoltageOk returns a tuple with the Voltage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetVoltageOk() (*int32, bool) {
	if o == nil || o.Voltage == nil {
		return nil, false
	}
	return o.Voltage, true
}

// HasVoltage returns a boolean if a field has been set.
func (o *PowerFeed) HasVoltage() bool {
	if o != nil && o.Voltage != nil {
		return true
	}

	return false
}

// SetVoltage gets a reference to the given int32 and assigns it to the Voltage field.
func (o *PowerFeed) SetVoltage(v int32) {
	o.Voltage = &v
}

// GetAmperage returns the Amperage field value if set, zero value otherwise.
func (o *PowerFeed) GetAmperage() int32 {
	if o == nil || o.Amperage == nil {
		var ret int32
		return ret
	}
	return *o.Amperage
}

// GetAmperageOk returns a tuple with the Amperage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetAmperageOk() (*int32, bool) {
	if o == nil || o.Amperage == nil {
		return nil, false
	}
	return o.Amperage, true
}

// HasAmperage returns a boolean if a field has been set.
func (o *PowerFeed) HasAmperage() bool {
	if o != nil && o.Amperage != nil {
		return true
	}

	return false
}

// SetAmperage gets a reference to the given int32 and assigns it to the Amperage field.
func (o *PowerFeed) SetAmperage(v int32) {
	o.Amperage = &v
}

// GetMaxUtilization returns the MaxUtilization field value if set, zero value otherwise.
func (o *PowerFeed) GetMaxUtilization() int32 {
	if o == nil || o.MaxUtilization == nil {
		var ret int32
		return ret
	}
	return *o.MaxUtilization
}

// GetMaxUtilizationOk returns a tuple with the MaxUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetMaxUtilizationOk() (*int32, bool) {
	if o == nil || o.MaxUtilization == nil {
		return nil, false
	}
	return o.MaxUtilization, true
}

// HasMaxUtilization returns a boolean if a field has been set.
func (o *PowerFeed) HasMaxUtilization() bool {
	if o != nil && o.MaxUtilization != nil {
		return true
	}

	return false
}

// SetMaxUtilization gets a reference to the given int32 and assigns it to the MaxUtilization field.
func (o *PowerFeed) SetMaxUtilization(v int32) {
	o.MaxUtilization = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PowerFeed) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PowerFeed) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PowerFeed) SetComments(v string) {
	o.Comments = &v
}

// GetCable returns the Cable field value
func (o *PowerFeed) GetCable() CircuitTerminationCable {
	if o == nil {
		var ret CircuitTerminationCable
		return ret
	}

	return o.Cable
}

// GetCableOk returns a tuple with the Cable field value
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetCableOk() (*CircuitTerminationCable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cable, true
}

// SetCable sets field value
func (o *PowerFeed) SetCable(v CircuitTerminationCable) {
	o.Cable = v
}

// GetCablePeer returns the CablePeer field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *PowerFeed) GetCablePeer() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CablePeer
}

// GetCablePeerOk returns a tuple with the CablePeer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerFeed) GetCablePeerOk() (map[string]interface{}, bool) {
	if o == nil || o.CablePeer == nil {
		return nil, false
	}
	return o.CablePeer, true
}

// SetCablePeer sets field value
func (o *PowerFeed) SetCablePeer(v map[string]interface{}) {
	o.CablePeer = v
}

// GetCablePeerType returns the CablePeerType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PowerFeed) GetCablePeerType() string {
	if o == nil || o.CablePeerType.Get() == nil {
		var ret string
		return ret
	}

	return *o.CablePeerType.Get()
}

// GetCablePeerTypeOk returns a tuple with the CablePeerType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerFeed) GetCablePeerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CablePeerType.Get(), o.CablePeerType.IsSet()
}

// SetCablePeerType sets field value
func (o *PowerFeed) SetCablePeerType(v string) {
	o.CablePeerType.Set(&v)
}

// GetConnectedEndpoint returns the ConnectedEndpoint field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *PowerFeed) GetConnectedEndpoint() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ConnectedEndpoint
}

// GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerFeed) GetConnectedEndpointOk() (map[string]interface{}, bool) {
	if o == nil || o.ConnectedEndpoint == nil {
		return nil, false
	}
	return o.ConnectedEndpoint, true
}

// SetConnectedEndpoint sets field value
func (o *PowerFeed) SetConnectedEndpoint(v map[string]interface{}) {
	o.ConnectedEndpoint = v
}

// GetConnectedEndpointType returns the ConnectedEndpointType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PowerFeed) GetConnectedEndpointType() string {
	if o == nil || o.ConnectedEndpointType.Get() == nil {
		var ret string
		return ret
	}

	return *o.ConnectedEndpointType.Get()
}

// GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerFeed) GetConnectedEndpointTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointType.Get(), o.ConnectedEndpointType.IsSet()
}

// SetConnectedEndpointType sets field value
func (o *PowerFeed) SetConnectedEndpointType(v string) {
	o.ConnectedEndpointType.Set(&v)
}

// GetConnectedEndpointReachable returns the ConnectedEndpointReachable field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *PowerFeed) GetConnectedEndpointReachable() bool {
	if o == nil || o.ConnectedEndpointReachable.Get() == nil {
		var ret bool
		return ret
	}

	return *o.ConnectedEndpointReachable.Get()
}

// GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerFeed) GetConnectedEndpointReachableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointReachable.Get(), o.ConnectedEndpointReachable.IsSet()
}

// SetConnectedEndpointReachable sets field value
func (o *PowerFeed) SetConnectedEndpointReachable(v bool) {
	o.ConnectedEndpointReachable.Set(&v)
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PowerFeed) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PowerFeed) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *PowerFeed) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PowerFeed) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PowerFeed) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PowerFeed) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
func (o *PowerFeed) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *PowerFeed) SetCreated(v string) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *PowerFeed) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *PowerFeed) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetDisplay returns the Display field value
func (o *PowerFeed) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *PowerFeed) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *PowerFeed) SetDisplay(v string) {
	o.Display = v
}

func (o PowerFeed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["power_panel"] = o.PowerPanel
	}
	if o.Rack.IsSet() {
		toSerialize["rack"] = o.Rack.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Supply != nil {
		toSerialize["supply"] = o.Supply
	}
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	if o.Voltage != nil {
		toSerialize["voltage"] = o.Voltage
	}
	if o.Amperage != nil {
		toSerialize["amperage"] = o.Amperage
	}
	if o.MaxUtilization != nil {
		toSerialize["max_utilization"] = o.MaxUtilization
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if true {
		toSerialize["cable"] = o.Cable
	}
	if o.CablePeer != nil {
		toSerialize["cable_peer"] = o.CablePeer
	}
	if true {
		toSerialize["cable_peer_type"] = o.CablePeerType.Get()
	}
	if o.ConnectedEndpoint != nil {
		toSerialize["connected_endpoint"] = o.ConnectedEndpoint
	}
	if true {
		toSerialize["connected_endpoint_type"] = o.ConnectedEndpointType.Get()
	}
	if true {
		toSerialize["connected_endpoint_reachable"] = o.ConnectedEndpointReachable.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePowerFeed struct {
	value *PowerFeed
	isSet bool
}

func (v NullablePowerFeed) Get() *PowerFeed {
	return v.value
}

func (v *NullablePowerFeed) Set(val *PowerFeed) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFeed) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFeed(val *PowerFeed) *NullablePowerFeed {
	return &NullablePowerFeed{value: val, isSet: true}
}

func (v NullablePowerFeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
