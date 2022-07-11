# Cable

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
**Status** | [**CableStatus**](CableStatus.md) |  | 
**Label** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **NullableInt32** |  | [optional] 
**LengthUnit** | Pointer to [**CableLengthUnit**](CableLengthUnit.md) |  | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewCable

`func NewCable(id string, url string, terminationAType string, terminationAId string, terminationA map[string]interface{}, terminationBType string, terminationBId string, terminationB map[string]interface{}, status CableStatus, display string, ) *Cable`

NewCable instantiates a new Cable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCableWithDefaults

`func NewCableWithDefaults() *Cable`

NewCableWithDefaults instantiates a new Cable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cable) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *Cable) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Cable) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Cable) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTerminationAType

`func (o *Cable) GetTerminationAType() string`

GetTerminationAType returns the TerminationAType field if non-nil, zero value otherwise.

### GetTerminationATypeOk

`func (o *Cable) GetTerminationATypeOk() (*string, bool)`

GetTerminationATypeOk returns a tuple with the TerminationAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAType

`func (o *Cable) SetTerminationAType(v string)`

SetTerminationAType sets TerminationAType field to given value.


### GetTerminationAId

`func (o *Cable) GetTerminationAId() string`

GetTerminationAId returns the TerminationAId field if non-nil, zero value otherwise.

### GetTerminationAIdOk

`func (o *Cable) GetTerminationAIdOk() (*string, bool)`

GetTerminationAIdOk returns a tuple with the TerminationAId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationAId

`func (o *Cable) SetTerminationAId(v string)`

SetTerminationAId sets TerminationAId field to given value.


### GetTerminationA

`func (o *Cable) GetTerminationA() map[string]interface{}`

GetTerminationA returns the TerminationA field if non-nil, zero value otherwise.

### GetTerminationAOk

`func (o *Cable) GetTerminationAOk() (*map[string]interface{}, bool)`

GetTerminationAOk returns a tuple with the TerminationA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationA

`func (o *Cable) SetTerminationA(v map[string]interface{})`

SetTerminationA sets TerminationA field to given value.


### SetTerminationANil

`func (o *Cable) SetTerminationANil(b bool)`

 SetTerminationANil sets the value for TerminationA to be an explicit nil

### UnsetTerminationA
`func (o *Cable) UnsetTerminationA()`

UnsetTerminationA ensures that no value is present for TerminationA, not even an explicit nil
### GetTerminationBType

`func (o *Cable) GetTerminationBType() string`

GetTerminationBType returns the TerminationBType field if non-nil, zero value otherwise.

### GetTerminationBTypeOk

`func (o *Cable) GetTerminationBTypeOk() (*string, bool)`

GetTerminationBTypeOk returns a tuple with the TerminationBType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBType

`func (o *Cable) SetTerminationBType(v string)`

SetTerminationBType sets TerminationBType field to given value.


### GetTerminationBId

`func (o *Cable) GetTerminationBId() string`

GetTerminationBId returns the TerminationBId field if non-nil, zero value otherwise.

### GetTerminationBIdOk

`func (o *Cable) GetTerminationBIdOk() (*string, bool)`

GetTerminationBIdOk returns a tuple with the TerminationBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationBId

`func (o *Cable) SetTerminationBId(v string)`

SetTerminationBId sets TerminationBId field to given value.


### GetTerminationB

`func (o *Cable) GetTerminationB() map[string]interface{}`

GetTerminationB returns the TerminationB field if non-nil, zero value otherwise.

### GetTerminationBOk

`func (o *Cable) GetTerminationBOk() (*map[string]interface{}, bool)`

GetTerminationBOk returns a tuple with the TerminationB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationB

`func (o *Cable) SetTerminationB(v map[string]interface{})`

SetTerminationB sets TerminationB field to given value.


### SetTerminationBNil

`func (o *Cable) SetTerminationBNil(b bool)`

 SetTerminationBNil sets the value for TerminationB to be an explicit nil

### UnsetTerminationB
`func (o *Cable) UnsetTerminationB()`

UnsetTerminationB ensures that no value is present for TerminationB, not even an explicit nil
### GetType

`func (o *Cable) GetType() CableType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Cable) GetTypeOk() (*CableType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Cable) SetType(v CableType)`

SetType sets Type field to given value.

### HasType

`func (o *Cable) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *Cable) GetStatus() CableStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cable) GetStatusOk() (*CableStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cable) SetStatus(v CableStatus)`

SetStatus sets Status field to given value.


### GetLabel

`func (o *Cable) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Cable) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Cable) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Cable) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetColor

`func (o *Cable) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Cable) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Cable) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Cable) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLength

`func (o *Cable) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *Cable) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *Cable) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *Cable) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *Cable) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *Cable) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetLengthUnit

`func (o *Cable) GetLengthUnit() CableLengthUnit`

GetLengthUnit returns the LengthUnit field if non-nil, zero value otherwise.

### GetLengthUnitOk

`func (o *Cable) GetLengthUnitOk() (*CableLengthUnit, bool)`

GetLengthUnitOk returns a tuple with the LengthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthUnit

`func (o *Cable) SetLengthUnit(v CableLengthUnit)`

SetLengthUnit sets LengthUnit field to given value.

### HasLengthUnit

`func (o *Cable) HasLengthUnit() bool`

HasLengthUnit returns a boolean if a field has been set.

### GetTags

`func (o *Cable) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Cable) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Cable) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Cable) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Cable) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Cable) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Cable) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Cable) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDisplay

`func (o *Cable) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Cable) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Cable) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


