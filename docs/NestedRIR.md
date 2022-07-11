# NestedRIR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**AggregateCount** | **int32** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewNestedRIR

`func NewNestedRIR(id string, url string, name string, aggregateCount int32, display string, ) *NestedRIR`

NewNestedRIR instantiates a new NestedRIR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedRIRWithDefaults

`func NewNestedRIRWithDefaults() *NestedRIR`

NewNestedRIRWithDefaults instantiates a new NestedRIR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedRIR) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedRIR) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedRIR) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedRIR) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedRIR) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedRIR) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *NestedRIR) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedRIR) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedRIR) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *NestedRIR) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *NestedRIR) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *NestedRIR) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *NestedRIR) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetAggregateCount

`func (o *NestedRIR) GetAggregateCount() int32`

GetAggregateCount returns the AggregateCount field if non-nil, zero value otherwise.

### GetAggregateCountOk

`func (o *NestedRIR) GetAggregateCountOk() (*int32, bool)`

GetAggregateCountOk returns a tuple with the AggregateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateCount

`func (o *NestedRIR) SetAggregateCount(v int32)`

SetAggregateCount sets AggregateCount field to given value.


### GetDisplay

`func (o *NestedRIR) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedRIR) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedRIR) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


