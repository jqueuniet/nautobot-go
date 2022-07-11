# WritableCable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**TerminationAType** | **string** |  | 
**TerminationAId** | **string** |  | 
**TerminationA** | **map[string]interface{}** |  | [readonly] 
**TerminationBType** | **string** |  | 
**TerminationBId** | **string** |  | 
**TerminationB** | **map[string]interface{}** |  | [readonly] 
**Type** | Pointer to [**CableType**](CableType.md) |  | [optional] 
**Status** | [**WritableCableStatusEnum**](WritableCableStatusEnum.md) |  | 
**Label** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **NullableInt32** |  | [optional] 
**LengthUnit** | Pointer to [**PatchedWritableCableLengthUnit**](PatchedWritableCableLengthUnit.md) |  | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ComputedFields** | **map[string]interface{}** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewWritableCable

`func NewWritableCable(id string, url string, terminationAType string, terminationAId string, terminationA map[string]interface{}, terminationBType string, terminationBId string, terminationB map[string]interface{}, status WritableCableStatusEnum, computedFields map[string]interface{}, display string, ) *WritableCable`

NewWritableCable instantiates a new WritableCable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableCableWithDefaults

`func NewWritableCableWithDefaults() *WritableCable`

NewWritableCableWithDefaults instantiates a new WritableCable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableCable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableCable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableCable) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *WritableCable) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableCable) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableCable) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTerminationAType

`func (o *WritableCable) GetTerminationAType() string`

GetTerminationAType returns the TerminationAType field if non-nil, zero value otherwise.

### GetTerminationATypeOk

`func (o *WritableCable) GetTerminationATypeOk() (*string, bool)`

GetTerminationATypeOk returns a tuple with the TerminationAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAType

`func (o *WritableCable) SetTerminationAType(v string)`

SetTerminationAType sets TerminationAType field to given value.


### GetTerminationAId

`func (o *WritableCable) GetTerminationAId() string`

GetTerminationAId returns the TerminationAId field if non-nil, zero value otherwise.

### GetTerminationAIdOk

`func (o *WritableCable) GetTerminationAIdOk() (*string, bool)`

GetTerminationAIdOk returns a tuple with the TerminationAId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAId

`func (o *WritableCable) SetTerminationAId(v string)`

SetTerminationAId sets TerminationAId field to given value.


### GetTerminationA

`func (o *WritableCable) GetTerminationA() map[string]interface{}`

GetTerminationA returns the TerminationA field if non-nil, zero value otherwise.

### GetTerminationAOk

`func (o *WritableCable) GetTerminationAOk() (*map[string]interface{}, bool)`

GetTerminationAOk returns a tuple with the TerminationA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationA

`func (o *WritableCable) SetTerminationA(v map[string]interface{})`

SetTerminationA sets TerminationA field to given value.


### SetTerminationANil

`func (o *WritableCable) SetTerminationANil(b bool)`

 SetTerminationANil sets the value for TerminationA to be an explicit nil

### UnsetTerminationA
`func (o *WritableCable) UnsetTerminationA()`

UnsetTerminationA ensures that no value is present for TerminationA, not even an explicit nil
### GetTerminationBType

`func (o *WritableCable) GetTerminationBType() string`

GetTerminationBType returns the TerminationBType field if non-nil, zero value otherwise.

### GetTerminationBTypeOk

`func (o *WritableCable) GetTerminationBTypeOk() (*string, bool)`

GetTerminationBTypeOk returns a tuple with the TerminationBType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBType

`func (o *WritableCable) SetTerminationBType(v string)`

SetTerminationBType sets TerminationBType field to given value.


### GetTerminationBId

`func (o *WritableCable) GetTerminationBId() string`

GetTerminationBId returns the TerminationBId field if non-nil, zero value otherwise.

### GetTerminationBIdOk

`func (o *WritableCable) GetTerminationBIdOk() (*string, bool)`

GetTerminationBIdOk returns a tuple with the TerminationBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBId

`func (o *WritableCable) SetTerminationBId(v string)`

SetTerminationBId sets TerminationBId field to given value.


### GetTerminationB

`func (o *WritableCable) GetTerminationB() map[string]interface{}`

GetTerminationB returns the TerminationB field if non-nil, zero value otherwise.

### GetTerminationBOk

`func (o *WritableCable) GetTerminationBOk() (*map[string]interface{}, bool)`

GetTerminationBOk returns a tuple with the TerminationB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationB

`func (o *WritableCable) SetTerminationB(v map[string]interface{})`

SetTerminationB sets TerminationB field to given value.


### SetTerminationBNil

`func (o *WritableCable) SetTerminationBNil(b bool)`

 SetTerminationBNil sets the value for TerminationB to be an explicit nil

### UnsetTerminationB
`func (o *WritableCable) UnsetTerminationB()`

UnsetTerminationB ensures that no value is present for TerminationB, not even an explicit nil
### GetType

`func (o *WritableCable) GetType() CableType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableCable) GetTypeOk() (*CableType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableCable) SetType(v CableType)`

SetType sets Type field to given value.

### HasType

`func (o *WritableCable) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *WritableCable) GetStatus() WritableCableStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableCable) GetStatusOk() (*WritableCableStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableCable) SetStatus(v WritableCableStatusEnum)`

SetStatus sets Status field to given value.


### GetLabel

`func (o *WritableCable) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritableCable) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritableCable) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritableCable) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetColor

`func (o *WritableCable) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WritableCable) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WritableCable) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *WritableCable) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLength

`func (o *WritableCable) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *WritableCable) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *WritableCable) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *WritableCable) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *WritableCable) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *WritableCable) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetLengthUnit

`func (o *WritableCable) GetLengthUnit() PatchedWritableCableLengthUnit`

GetLengthUnit returns the LengthUnit field if non-nil, zero value otherwise.

### GetLengthUnitOk

`func (o *WritableCable) GetLengthUnitOk() (*PatchedWritableCableLengthUnit, bool)`

GetLengthUnitOk returns a tuple with the LengthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthUnit

`func (o *WritableCable) SetLengthUnit(v PatchedWritableCableLengthUnit)`

SetLengthUnit sets LengthUnit field to given value.

### HasLengthUnit

`func (o *WritableCable) HasLengthUnit() bool`

HasLengthUnit returns a boolean if a field has been set.

### GetTags

`func (o *WritableCable) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableCable) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableCable) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableCable) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableCable) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableCable) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableCable) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableCable) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetComputedFields

`func (o *WritableCable) GetComputedFields() map[string]interface{}`

GetComputedFields returns the ComputedFields field if non-nil, zero value otherwise.

### GetComputedFieldsOk

`func (o *WritableCable) GetComputedFieldsOk() (*map[string]interface{}, bool)`

GetComputedFieldsOk returns a tuple with the ComputedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFields

`func (o *WritableCable) SetComputedFields(v map[string]interface{})`

SetComputedFields sets ComputedFields field to given value.


### GetDisplay

`func (o *WritableCable) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableCable) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableCable) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


