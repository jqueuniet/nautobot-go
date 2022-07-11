# WritableCustomFieldChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Field** | **string** |  | 
**Value** | **string** |  | 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewWritableCustomFieldChoice

`func NewWritableCustomFieldChoice(id string, url string, field string, value string, display string, ) *WritableCustomFieldChoice`

NewWritableCustomFieldChoice instantiates a new WritableCustomFieldChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableCustomFieldChoiceWithDefaults

`func NewWritableCustomFieldChoiceWithDefaults() *WritableCustomFieldChoice`

NewWritableCustomFieldChoiceWithDefaults instantiates a new WritableCustomFieldChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableCustomFieldChoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableCustomFieldChoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableCustomFieldChoice) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *WritableCustomFieldChoice) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableCustomFieldChoice) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableCustomFieldChoice) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetField

`func (o *WritableCustomFieldChoice) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *WritableCustomFieldChoice) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *WritableCustomFieldChoice) SetField(v string)`

SetField sets Field field to given value.


### GetValue

`func (o *WritableCustomFieldChoice) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WritableCustomFieldChoice) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WritableCustomFieldChoice) SetValue(v string)`

SetValue sets Value field to given value.


### GetWeight

`func (o *WritableCustomFieldChoice) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *WritableCustomFieldChoice) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *WritableCustomFieldChoice) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *WritableCustomFieldChoice) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisplay

`func (o *WritableCustomFieldChoice) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableCustomFieldChoice) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableCustomFieldChoice) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


