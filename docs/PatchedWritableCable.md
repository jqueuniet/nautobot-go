# PatchedWritableCable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**TerminationAType** | Pointer to **string** |  | [optional] 
**TerminationAId** | Pointer to **string** |  | [optional] 
**TerminationA** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**TerminationBType** | Pointer to **string** |  | [optional] 
**TerminationBId** | Pointer to **string** |  | [optional] 
**TerminationB** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Type** | Pointer to [**CableType**](CableType.md) |  | [optional] 
**Status** | Pointer to [**WritableCableStatusEnum**](WritableCableStatusEnum.md) |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **NullableInt32** |  | [optional] 
**LengthUnit** | Pointer to [**PatchedWritableCableLengthUnit**](PatchedWritableCableLengthUnit.md) |  | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ComputedFields** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedWritableCable

`func NewPatchedWritableCable() *PatchedWritableCable`

NewPatchedWritableCable instantiates a new PatchedWritableCable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableCableWithDefaults

`func NewPatchedWritableCableWithDefaults() *PatchedWritableCable`

NewPatchedWritableCableWithDefaults instantiates a new PatchedWritableCable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedWritableCable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedWritableCable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedWritableCable) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedWritableCable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedWritableCable) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedWritableCable) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedWritableCable) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedWritableCable) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTerminationAType

`func (o *PatchedWritableCable) GetTerminationAType() string`

GetTerminationAType returns the TerminationAType field if non-nil, zero value otherwise.

### GetTerminationATypeOk

`func (o *PatchedWritableCable) GetTerminationATypeOk() (*string, bool)`

GetTerminationATypeOk returns a tuple with the TerminationAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAType

`func (o *PatchedWritableCable) SetTerminationAType(v string)`

SetTerminationAType sets TerminationAType field to given value.

### HasTerminationAType

`func (o *PatchedWritableCable) HasTerminationAType() bool`

HasTerminationAType returns a boolean if a field has been set.

### GetTerminationAId

`func (o *PatchedWritableCable) GetTerminationAId() string`

GetTerminationAId returns the TerminationAId field if non-nil, zero value otherwise.

### GetTerminationAIdOk

`func (o *PatchedWritableCable) GetTerminationAIdOk() (*string, bool)`

GetTerminationAIdOk returns a tuple with the TerminationAId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAId

`func (o *PatchedWritableCable) SetTerminationAId(v string)`

SetTerminationAId sets TerminationAId field to given value.

### HasTerminationAId

`func (o *PatchedWritableCable) HasTerminationAId() bool`

HasTerminationAId returns a boolean if a field has been set.

### GetTerminationA

`func (o *PatchedWritableCable) GetTerminationA() map[string]interface{}`

GetTerminationA returns the TerminationA field if non-nil, zero value otherwise.

### GetTerminationAOk

`func (o *PatchedWritableCable) GetTerminationAOk() (*map[string]interface{}, bool)`

GetTerminationAOk returns a tuple with the TerminationA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationA

`func (o *PatchedWritableCable) SetTerminationA(v map[string]interface{})`

SetTerminationA sets TerminationA field to given value.

### HasTerminationA

`func (o *PatchedWritableCable) HasTerminationA() bool`

HasTerminationA returns a boolean if a field has been set.

### SetTerminationANil

`func (o *PatchedWritableCable) SetTerminationANil(b bool)`

 SetTerminationANil sets the value for TerminationA to be an explicit nil

### UnsetTerminationA
`func (o *PatchedWritableCable) UnsetTerminationA()`

UnsetTerminationA ensures that no value is present for TerminationA, not even an explicit nil
### GetTerminationBType

`func (o *PatchedWritableCable) GetTerminationBType() string`

GetTerminationBType returns the TerminationBType field if non-nil, zero value otherwise.

### GetTerminationBTypeOk

`func (o *PatchedWritableCable) GetTerminationBTypeOk() (*string, bool)`

GetTerminationBTypeOk returns a tuple with the TerminationBType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBType

`func (o *PatchedWritableCable) SetTerminationBType(v string)`

SetTerminationBType sets TerminationBType field to given value.

### HasTerminationBType

`func (o *PatchedWritableCable) HasTerminationBType() bool`

HasTerminationBType returns a boolean if a field has been set.

### GetTerminationBId

`func (o *PatchedWritableCable) GetTerminationBId() string`

GetTerminationBId returns the TerminationBId field if non-nil, zero value otherwise.

### GetTerminationBIdOk

`func (o *PatchedWritableCable) GetTerminationBIdOk() (*string, bool)`

GetTerminationBIdOk returns a tuple with the TerminationBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBId

`func (o *PatchedWritableCable) SetTerminationBId(v string)`

SetTerminationBId sets TerminationBId field to given value.

### HasTerminationBId

`func (o *PatchedWritableCable) HasTerminationBId() bool`

HasTerminationBId returns a boolean if a field has been set.

### GetTerminationB

`func (o *PatchedWritableCable) GetTerminationB() map[string]interface{}`

GetTerminationB returns the TerminationB field if non-nil, zero value otherwise.

### GetTerminationBOk

`func (o *PatchedWritableCable) GetTerminationBOk() (*map[string]interface{}, bool)`

GetTerminationBOk returns a tuple with the TerminationB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationB

`func (o *PatchedWritableCable) SetTerminationB(v map[string]interface{})`

SetTerminationB sets TerminationB field to given value.

### HasTerminationB

`func (o *PatchedWritableCable) HasTerminationB() bool`

HasTerminationB returns a boolean if a field has been set.

### SetTerminationBNil

`func (o *PatchedWritableCable) SetTerminationBNil(b bool)`

 SetTerminationBNil sets the value for TerminationB to be an explicit nil

### UnsetTerminationB
`func (o *PatchedWritableCable) UnsetTerminationB()`

UnsetTerminationB ensures that no value is present for TerminationB, not even an explicit nil
### GetType

`func (o *PatchedWritableCable) GetType() CableType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableCable) GetTypeOk() (*CableType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableCable) SetType(v CableType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableCable) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritableCable) GetStatus() WritableCableStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableCable) GetStatusOk() (*WritableCableStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableCable) SetStatus(v WritableCableStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableCable) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedWritableCable) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedWritableCable) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedWritableCable) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedWritableCable) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetColor

`func (o *PatchedWritableCable) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PatchedWritableCable) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PatchedWritableCable) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PatchedWritableCable) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLength

`func (o *PatchedWritableCable) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *PatchedWritableCable) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *PatchedWritableCable) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *PatchedWritableCable) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *PatchedWritableCable) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *PatchedWritableCable) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetLengthUnit

`func (o *PatchedWritableCable) GetLengthUnit() PatchedWritableCableLengthUnit`

GetLengthUnit returns the LengthUnit field if non-nil, zero value otherwise.

### GetLengthUnitOk

`func (o *PatchedWritableCable) GetLengthUnitOk() (*PatchedWritableCableLengthUnit, bool)`

GetLengthUnitOk returns a tuple with the LengthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthUnit

`func (o *PatchedWritableCable) SetLengthUnit(v PatchedWritableCableLengthUnit)`

SetLengthUnit sets LengthUnit field to given value.

### HasLengthUnit

`func (o *PatchedWritableCable) HasLengthUnit() bool`

HasLengthUnit returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableCable) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableCable) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableCable) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableCable) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableCable) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableCable) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableCable) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableCable) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetComputedFields

`func (o *PatchedWritableCable) GetComputedFields() map[string]interface{}`

GetComputedFields returns the ComputedFields field if non-nil, zero value otherwise.

### GetComputedFieldsOk

`func (o *PatchedWritableCable) GetComputedFieldsOk() (*map[string]interface{}, bool)`

GetComputedFieldsOk returns a tuple with the ComputedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFields

`func (o *PatchedWritableCable) SetComputedFields(v map[string]interface{})`

SetComputedFields sets ComputedFields field to given value.

### HasComputedFields

`func (o *PatchedWritableCable) HasComputedFields() bool`

HasComputedFields returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedWritableCable) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedWritableCable) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedWritableCable) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedWritableCable) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


