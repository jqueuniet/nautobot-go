# PatchedWritableFrontPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Device** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**PortTypeChoices**](PortTypeChoices.md) |  | [optional] 
**RearPort** | Pointer to **string** |  | [optional] 
**RearPortPosition** | Pointer to **int32** |  | [optional] [default to 1]
**Description** | Pointer to **string** |  | [optional] 
**Cable** | Pointer to [**CircuitTerminationCable**](CircuitTerminationCable.md) |  | [optional] 
**CablePeer** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**CablePeerType** | Pointer to **NullableString** |  | [optional] [readonly] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ComputedFields** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedWritableFrontPort

`func NewPatchedWritableFrontPort() *PatchedWritableFrontPort`

NewPatchedWritableFrontPort instantiates a new PatchedWritableFrontPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableFrontPortWithDefaults

`func NewPatchedWritableFrontPortWithDefaults() *PatchedWritableFrontPort`

NewPatchedWritableFrontPortWithDefaults instantiates a new PatchedWritableFrontPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedWritableFrontPort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedWritableFrontPort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedWritableFrontPort) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedWritableFrontPort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedWritableFrontPort) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedWritableFrontPort) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedWritableFrontPort) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedWritableFrontPort) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDevice

`func (o *PatchedWritableFrontPort) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedWritableFrontPort) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedWritableFrontPort) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedWritableFrontPort) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetName

`func (o *PatchedWritableFrontPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableFrontPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableFrontPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableFrontPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedWritableFrontPort) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedWritableFrontPort) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedWritableFrontPort) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedWritableFrontPort) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritableFrontPort) GetType() PortTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableFrontPort) GetTypeOk() (*PortTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableFrontPort) SetType(v PortTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableFrontPort) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRearPort

`func (o *PatchedWritableFrontPort) GetRearPort() string`

GetRearPort returns the RearPort field if non-nil, zero value otherwise.

### GetRearPortOk

`func (o *PatchedWritableFrontPort) GetRearPortOk() (*string, bool)`

GetRearPortOk returns a tuple with the RearPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPort

`func (o *PatchedWritableFrontPort) SetRearPort(v string)`

SetRearPort sets RearPort field to given value.

### HasRearPort

`func (o *PatchedWritableFrontPort) HasRearPort() bool`

HasRearPort returns a boolean if a field has been set.

### GetRearPortPosition

`func (o *PatchedWritableFrontPort) GetRearPortPosition() int32`

GetRearPortPosition returns the RearPortPosition field if non-nil, zero value otherwise.

### GetRearPortPositionOk

`func (o *PatchedWritableFrontPort) GetRearPortPositionOk() (*int32, bool)`

GetRearPortPositionOk returns a tuple with the RearPortPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPortPosition

`func (o *PatchedWritableFrontPort) SetRearPortPosition(v int32)`

SetRearPortPosition sets RearPortPosition field to given value.

### HasRearPortPosition

`func (o *PatchedWritableFrontPort) HasRearPortPosition() bool`

HasRearPortPosition returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableFrontPort) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableFrontPort) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableFrontPort) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableFrontPort) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCable

`func (o *PatchedWritableFrontPort) GetCable() CircuitTerminationCable`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *PatchedWritableFrontPort) GetCableOk() (*CircuitTerminationCable, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *PatchedWritableFrontPort) SetCable(v CircuitTerminationCable)`

SetCable sets Cable field to given value.

### HasCable

`func (o *PatchedWritableFrontPort) HasCable() bool`

HasCable returns a boolean if a field has been set.

### GetCablePeer

`func (o *PatchedWritableFrontPort) GetCablePeer() map[string]interface{}`

GetCablePeer returns the CablePeer field if non-nil, zero value otherwise.

### GetCablePeerOk

`func (o *PatchedWritableFrontPort) GetCablePeerOk() (*map[string]interface{}, bool)`

GetCablePeerOk returns a tuple with the CablePeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeer

`func (o *PatchedWritableFrontPort) SetCablePeer(v map[string]interface{})`

SetCablePeer sets CablePeer field to given value.

### HasCablePeer

`func (o *PatchedWritableFrontPort) HasCablePeer() bool`

HasCablePeer returns a boolean if a field has been set.

### SetCablePeerNil

`func (o *PatchedWritableFrontPort) SetCablePeerNil(b bool)`

 SetCablePeerNil sets the value for CablePeer to be an explicit nil

### UnsetCablePeer
`func (o *PatchedWritableFrontPort) UnsetCablePeer()`

UnsetCablePeer ensures that no value is present for CablePeer, not even an explicit nil
### GetCablePeerType

`func (o *PatchedWritableFrontPort) GetCablePeerType() string`

GetCablePeerType returns the CablePeerType field if non-nil, zero value otherwise.

### GetCablePeerTypeOk

`func (o *PatchedWritableFrontPort) GetCablePeerTypeOk() (*string, bool)`

GetCablePeerTypeOk returns a tuple with the CablePeerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeerType

`func (o *PatchedWritableFrontPort) SetCablePeerType(v string)`

SetCablePeerType sets CablePeerType field to given value.

### HasCablePeerType

`func (o *PatchedWritableFrontPort) HasCablePeerType() bool`

HasCablePeerType returns a boolean if a field has been set.

### SetCablePeerTypeNil

`func (o *PatchedWritableFrontPort) SetCablePeerTypeNil(b bool)`

 SetCablePeerTypeNil sets the value for CablePeerType to be an explicit nil

### UnsetCablePeerType
`func (o *PatchedWritableFrontPort) UnsetCablePeerType()`

UnsetCablePeerType ensures that no value is present for CablePeerType, not even an explicit nil
### GetTags

`func (o *PatchedWritableFrontPort) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableFrontPort) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableFrontPort) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableFrontPort) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableFrontPort) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableFrontPort) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableFrontPort) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableFrontPort) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetComputedFields

`func (o *PatchedWritableFrontPort) GetComputedFields() map[string]interface{}`

GetComputedFields returns the ComputedFields field if non-nil, zero value otherwise.

### GetComputedFieldsOk

`func (o *PatchedWritableFrontPort) GetComputedFieldsOk() (*map[string]interface{}, bool)`

GetComputedFieldsOk returns a tuple with the ComputedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFields

`func (o *PatchedWritableFrontPort) SetComputedFields(v map[string]interface{})`

SetComputedFields sets ComputedFields field to given value.

### HasComputedFields

`func (o *PatchedWritableFrontPort) HasComputedFields() bool`

HasComputedFields returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedWritableFrontPort) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedWritableFrontPort) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedWritableFrontPort) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedWritableFrontPort) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


