# AvailableIPVrf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Rd** | Pointer to **NullableString** | Unique route distinguisher (as defined in RFC 4364) | [optional] 
**Display** | **string** | Human friendly display value | [readonly] 
**PrefixCount** | **int32** |  | [readonly] 

## Methods

### NewAvailableIPVrf

`func NewAvailableIPVrf(id string, url string, name string, display string, prefixCount int32, ) *AvailableIPVrf`

NewAvailableIPVrf instantiates a new AvailableIPVrf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableIPVrfWithDefaults

`func NewAvailableIPVrfWithDefaults() *AvailableIPVrf`

NewAvailableIPVrfWithDefaults instantiates a new AvailableIPVrf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AvailableIPVrf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AvailableIPVrf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AvailableIPVrf) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *AvailableIPVrf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AvailableIPVrf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AvailableIPVrf) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *AvailableIPVrf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AvailableIPVrf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AvailableIPVrf) SetName(v string)`

SetName sets Name field to given value.


### GetRd

`func (o *AvailableIPVrf) GetRd() string`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *AvailableIPVrf) GetRdOk() (*string, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *AvailableIPVrf) SetRd(v string)`

SetRd sets Rd field to given value.

### HasRd

`func (o *AvailableIPVrf) HasRd() bool`

HasRd returns a boolean if a field has been set.

### SetRdNil

`func (o *AvailableIPVrf) SetRdNil(b bool)`

 SetRdNil sets the value for Rd to be an explicit nil

### UnsetRd
`func (o *AvailableIPVrf) UnsetRd()`

UnsetRd ensures that no value is present for Rd, not even an explicit nil
### GetDisplay

`func (o *AvailableIPVrf) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *AvailableIPVrf) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *AvailableIPVrf) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetPrefixCount

`func (o *AvailableIPVrf) GetPrefixCount() int32`

GetPrefixCount returns the PrefixCount field if non-nil, zero value otherwise.

### GetPrefixCountOk

`func (o *AvailableIPVrf) GetPrefixCountOk() (*int32, bool)`

GetPrefixCountOk returns a tuple with the PrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCount

`func (o *AvailableIPVrf) SetPrefixCount(v int32)`

SetPrefixCount sets PrefixCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


