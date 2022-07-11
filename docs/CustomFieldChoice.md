# CustomFieldChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Field** | [**NestedCustomField**](NestedCustomField.md) |  | 
**Value** | **string** |  | 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewCustomFieldChoice

`func NewCustomFieldChoice(id string, url string, field NestedCustomField, value string, display string, ) *CustomFieldChoice`

NewCustomFieldChoice instantiates a new CustomFieldChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldChoiceWithDefaults

`func NewCustomFieldChoiceWithDefaults() *CustomFieldChoice`

NewCustomFieldChoiceWithDefaults instantiates a new CustomFieldChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFieldChoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFieldChoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFieldChoice) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *CustomFieldChoice) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CustomFieldChoice) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CustomFieldChoice) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetField

`func (o *CustomFieldChoice) GetField() NestedCustomField`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *CustomFieldChoice) GetFieldOk() (*NestedCustomField, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *CustomFieldChoice) SetField(v NestedCustomField)`

SetField sets Field field to given value.


### GetValue

`func (o *CustomFieldChoice) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldChoice) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldChoice) SetValue(v string)`

SetValue sets Value field to given value.


### GetWeight

`func (o *CustomFieldChoice) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CustomFieldChoice) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CustomFieldChoice) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CustomFieldChoice) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisplay

`func (o *CustomFieldChoice) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CustomFieldChoice) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CustomFieldChoice) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


