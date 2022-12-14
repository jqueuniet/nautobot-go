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

// WritablePowerFeed Mixin to add `status` choice field to model serializers.
type WritablePowerFeed struct {
	Id string `json:"id"`
	Url string `json:"url"`
	PowerPanel string `json:"power_panel"`
	Rack NullableString `json:"rack,omitempty"`
	Name string `json:"name"`
	Status WritablePowerFeedStatusEnum `json:"status"`
	Type *PowerFeedTypeChoices `json:"type,omitempty"`
	Supply *SupplyEnum `json:"supply,omitempty"`
	Phase *PhaseEnum `json:"phase,omitempty"`
	Voltage *int32 `json:"voltage,omitempty"`
	Amperage *int32 `json:"amperage,omitempty"`
	// Maximum permissible draw (percentage)
	MaxUtilization *int32 `json:"max_utilization,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Cable CircuitTerminationCable `json:"cable"`
	CablePeer map[string]interface{} `json:"cable_peer"`
	CablePeerType NullableString `json:"cable_peer_type"`
	ConnectedEndpoint map[string]interface{} `json:"connected_endpoint"`
	ConnectedEndpointType NullableString `json:"connected_endpoint_type"`
	ConnectedEndpointReachable NullableBool `json:"connected_endpoint_reachable"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	ComputedFields map[string]interface{} `json:"computed_fields"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewWritablePowerFeed instantiates a new WritablePowerFeed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritablePowerFeed(id string, url string, powerPanel string, name string, status WritablePowerFeedStatusEnum, cable CircuitTerminationCable, cablePeer map[string]interface{}, cablePeerType NullableString, connectedEndpoint map[string]interface{}, connectedEndpointType NullableString, connectedEndpointReachable NullableBool, created string, lastUpdated time.Time, computedFields map[string]interface{}, display string) *WritablePowerFeed {
	this := WritablePowerFeed{}
	this.Id = id
	this.Url = url
	this.PowerPanel = powerPanel
	this.Name = name
	this.Status = status
	this.Cable = cable
	this.CablePeer = cablePeer
	this.CablePeerType = cablePeerType
	this.ConnectedEndpoint = connectedEndpoint
	this.ConnectedEndpointType = connectedEndpointType
	this.ConnectedEndpointReachable = connectedEndpointReachable
	this.Created = created
	this.LastUpdated = lastUpdated
	this.ComputedFields = computedFields
	this.Display = display
	return &this
}

// NewWritablePowerFeedWithDefaults instantiates a new WritablePowerFeed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritablePowerFeedWithDefaults() *WritablePowerFeed {
	this := WritablePowerFeed{}
	return &this
}

// GetId returns the Id field value
func (o *WritablePowerFeed) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WritablePowerFeed) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *WritablePowerFeed) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WritablePowerFeed) SetUrl(v string) {
	o.Url = v
}

// GetPowerPanel returns the PowerPanel field value
func (o *WritablePowerFeed) GetPowerPanel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PowerPanel
}

// GetPowerPanelOk returns a tuple with the PowerPanel field value
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetPowerPanelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerPanel, true
}

// SetPowerPanel sets field value
func (o *WritablePowerFeed) SetPowerPanel(v string) {
	o.PowerPanel = v
}

// GetRack returns the Rack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritablePowerFeed) GetRack() string {
	if o == nil || o.Rack.Get() == nil {
		var ret string
		return ret
	}
	return *o.Rack.Get()
}

// GetRackOk returns a tuple with the Rack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritablePowerFeed) GetRackOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rack.Get(), o.Rack.IsSet()
}

// HasRack returns a boolean if a field has been set.
func (o *WritablePowerFeed) HasRack() bool {
	if o != nil && o.Rack.IsSet() {
		return true
	}

	return false
}

// SetRack gets a reference to the given NullableString and assigns it to the Rack field.
func (o *WritablePowerFeed) SetRack(v string) {
	o.Rack.Set(&v)
}
// SetRackNil sets the value for Rack to be an explicit nil
func (o *WritablePowerFeed) SetRackNil() {
	o.Rack.Set(nil)
}

// UnsetRack ensures that no value is present for Rack, not even an explicit nil
func (o *WritablePowerFeed) UnsetRack() {
	o.Rack.Unset()
}

// GetName returns the Name field value
func (o *WritablePowerFeed) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritablePowerFeed) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *WritablePowerFeed) GetStatus() WritablePowerFeedStatusEnum {
	if o == nil {
		var ret WritablePowerFeedStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetStatusOk() (*WritablePowerFeedStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WritablePowerFeed) SetStatus(v WritablePowerFeedStatusEnum) {
	o.Status = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WritablePowerFeed) GetType() PowerFeedTypeChoices {
	if o == nil || o.Type == nil {
		var ret PowerFeedTypeChoices
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetTypeOk() (*PowerFeedTypeChoices, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WritablePowerFeed) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given PowerFeedTypeChoices and assigns it to the Type field.
func (o *WritablePowerFeed) SetType(v PowerFeedTypeChoices) {
	o.Type = &v
}

// GetSupply returns the Supply field value if set, zero value otherwise.
func (o *WritablePowerFeed) GetSupply() SupplyEnum {
	if o == nil || o.Supply == nil {
		var ret SupplyEnum
		return ret
	}
	return *o.Supply
}

// GetSupplyOk returns a tuple with the Supply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetSupplyOk() (*SupplyEnum, bool) {
	if o == nil || o.Supply == nil {
		return nil, false
	}
	return o.Supply, true
}

// HasSupply returns a boolean if a field has been set.
func (o *WritablePowerFeed) HasSupply() bool {
	if o != nil && o.Supply != nil {
		return true
	}

	return false
}

// SetSupply gets a reference to the given SupplyEnum and assigns it to the Supply field.
func (o *WritablePowerFeed) SetSupply(v SupplyEnum) {
	o.Supply = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *WritablePowerFeed) GetPhase() PhaseEnum {
	if o == nil || o.Phase == nil {
		var ret PhaseEnum
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetPhaseOk() (*PhaseEnum, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *WritablePowerFeed) HasPhase() bool {
	if o != nil && o.Phase != nil {
		return true
	}

	return false
}

// SetPhase gets a reference to the given PhaseEnum and assigns it to the Phase field.
func (o *WritablePowerFeed) SetPhase(v PhaseEnum) {
	o.Phase = &v
}

// GetVoltage returns the Voltage field value if set, zero value otherwise.
func (o *WritablePowerFeed) GetVoltage() int32 {
	if o == nil || o.Voltage == nil {
		var ret int32
		return ret
	}
	return *o.Voltage
}

// GetVoltageOk returns a tuple with the Voltage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetVoltageOk() (*int32, bool) {
	if o == nil || o.Voltage == nil {
		return nil, false
	}
	return o.Voltage, true
}

// HasVoltage returns a boolean if a field has been set.
func (o *WritablePowerFeed) HasVoltage() bool {
	if o != nil && o.Voltage != nil {
		return true
	}

	return false
}

// SetVoltage gets a reference to the given int32 and assigns it to the Voltage field.
func (o *WritablePowerFeed) SetVoltage(v int32) {
	o.Voltage = &v
}

// GetAmperage returns the Amperage field value if set, zero value otherwise.
func (o *WritablePowerFeed) GetAmperage() int32 {
	if o == nil || o.Amperage == nil {
		var ret int32
		return ret
	}
	return *o.Amperage
}

// GetAmperageOk returns a tuple with the Amperage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetAmperageOk() (*int32, bool) {
	if o == nil || o.Amperage == nil {
		return nil, false
	}
	return o.Amperage, true
}

// HasAmperage returns a boolean if a field has been set.
func (o *WritablePowerFeed) HasAmperage() bool {
	if o != nil && o.Amperage != nil {
		return true
	}

	return false
}

// SetAmperage gets a reference to the given int32 and assigns it to the Amperage field.
func (o *WritablePowerFeed) SetAmperage(v int32) {
	o.Amperage = &v
}

// GetMaxUtilization returns the MaxUtilization field value if set, zero value otherwise.
func (o *WritablePowerFeed) GetMaxUtilization() int32 {
	if o == nil || o.MaxUtilization == nil {
		var ret int32
		return ret
	}
	return *o.MaxUtilization
}

// GetMaxUtilizationOk returns a tuple with the MaxUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetMaxUtilizationOk() (*int32, bool) {
	if o == nil || o.MaxUtilization == nil {
		return nil, false
	}
	return o.MaxUtilization, true
}

// HasMaxUtilization returns a boolean if a field has been set.
func (o *WritablePowerFeed) HasMaxUtilization() bool {
	if o != nil && o.MaxUtilization != nil {
		return true
	}

	return false
}

// SetMaxUtilization gets a reference to the given int32 and assigns it to the MaxUtilization field.
func (o *WritablePowerFeed) SetMaxUtilization(v int32) {
	o.MaxUtilization = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritablePowerFeed) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritablePowerFeed) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritablePowerFeed) SetComments(v string) {
	o.Comments = &v
}

// GetCable returns the Cable field value
func (o *WritablePowerFeed) GetCable() CircuitTerminationCable {
	if o == nil {
		var ret CircuitTerminationCable
		return ret
	}

	return o.Cable
}

// GetCableOk returns a tuple with the Cable field value
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetCableOk() (*CircuitTerminationCable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cable, true
}

// SetCable sets field value
func (o *WritablePowerFeed) SetCable(v CircuitTerminationCable) {
	o.Cable = v
}

// GetCablePeer returns the CablePeer field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *WritablePowerFeed) GetCablePeer() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CablePeer
}

// GetCablePeerOk returns a tuple with the CablePeer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritablePowerFeed) GetCablePeerOk() (map[string]interface{}, bool) {
	if o == nil || o.CablePeer == nil {
		return nil, false
	}
	return o.CablePeer, true
}

// SetCablePeer sets field value
func (o *WritablePowerFeed) SetCablePeer(v map[string]interface{}) {
	o.CablePeer = v
}

// GetCablePeerType returns the CablePeerType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WritablePowerFeed) GetCablePeerType() string {
	if o == nil || o.CablePeerType.Get() == nil {
		var ret string
		return ret
	}

	return *o.CablePeerType.Get()
}

// GetCablePeerTypeOk returns a tuple with the CablePeerType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritablePowerFeed) GetCablePeerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CablePeerType.Get(), o.CablePeerType.IsSet()
}

// SetCablePeerType sets field value
func (o *WritablePowerFeed) SetCablePeerType(v string) {
	o.CablePeerType.Set(&v)
}

// GetConnectedEndpoint returns the ConnectedEndpoint field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *WritablePowerFeed) GetConnectedEndpoint() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ConnectedEndpoint
}

// GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritablePowerFeed) GetConnectedEndpointOk() (map[string]interface{}, bool) {
	if o == nil || o.ConnectedEndpoint == nil {
		return nil, false
	}
	return o.ConnectedEndpoint, true
}

// SetConnectedEndpoint sets field value
func (o *WritablePowerFeed) SetConnectedEndpoint(v map[string]interface{}) {
	o.ConnectedEndpoint = v
}

// GetConnectedEndpointType returns the ConnectedEndpointType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WritablePowerFeed) GetConnectedEndpointType() string {
	if o == nil || o.ConnectedEndpointType.Get() == nil {
		var ret string
		return ret
	}

	return *o.ConnectedEndpointType.Get()
}

// GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritablePowerFeed) GetConnectedEndpointTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointType.Get(), o.ConnectedEndpointType.IsSet()
}

// SetConnectedEndpointType sets field value
func (o *WritablePowerFeed) SetConnectedEndpointType(v string) {
	o.ConnectedEndpointType.Set(&v)
}

// GetConnectedEndpointReachable returns the ConnectedEndpointReachable field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *WritablePowerFeed) GetConnectedEndpointReachable() bool {
	if o == nil || o.ConnectedEndpointReachable.Get() == nil {
		var ret bool
		return ret
	}

	return *o.ConnectedEndpointReachable.Get()
}

// GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritablePowerFeed) GetConnectedEndpointReachableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointReachable.Get(), o.ConnectedEndpointReachable.IsSet()
}

// SetConnectedEndpointReachable sets field value
func (o *WritablePowerFeed) SetConnectedEndpointReachable(v bool) {
	o.ConnectedEndpointReachable.Set(&v)
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritablePowerFeed) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritablePowerFeed) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *WritablePowerFeed) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritablePowerFeed) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritablePowerFeed) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritablePowerFeed) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
func (o *WritablePowerFeed) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *WritablePowerFeed) SetCreated(v string) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *WritablePowerFeed) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *WritablePowerFeed) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetComputedFields returns the ComputedFields field value
func (o *WritablePowerFeed) GetComputedFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ComputedFields
}

// GetComputedFieldsOk returns a tuple with the ComputedFields field value
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetComputedFieldsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputedFields, true
}

// SetComputedFields sets field value
func (o *WritablePowerFeed) SetComputedFields(v map[string]interface{}) {
	o.ComputedFields = v
}

// GetDisplay returns the Display field value
func (o *WritablePowerFeed) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *WritablePowerFeed) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *WritablePowerFeed) SetDisplay(v string) {
	o.Display = v
}

func (o WritablePowerFeed) MarshalJSON() ([]byte, error) {
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
		toSerialize["computed_fields"] = o.ComputedFields
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableWritablePowerFeed struct {
	value *WritablePowerFeed
	isSet bool
}

func (v NullableWritablePowerFeed) Get() *WritablePowerFeed {
	return v.value
}

func (v *NullableWritablePowerFeed) Set(val *WritablePowerFeed) {
	v.value = val
	v.isSet = true
}

func (v NullableWritablePowerFeed) IsSet() bool {
	return v.isSet
}

func (v *NullableWritablePowerFeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritablePowerFeed(val *WritablePowerFeed) *NullableWritablePowerFeed {
	return &NullableWritablePowerFeed{value: val, isSet: true}
}

func (v NullableWritablePowerFeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritablePowerFeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


