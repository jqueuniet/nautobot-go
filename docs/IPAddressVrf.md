# IPAddressVrf

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

### NewIPAddressVrf

`func NewIPAddressVrf(id string, url string, name string, display string, prefixCount int32, ) *IPAddressVrf`

NewIPAddressVrf instantiates a new IPAddressVrf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressVrfWithDefaults

`func NewIPAddressVrfWithDefaults() *IPAddressVrf`

NewIPAddressVrfWithDefaults instantiates a new IPAddressVrf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPAddressVrf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPAddressVrf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPAddressVrf) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *IPAddressVrf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IPAddressVrf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IPAddressVrf) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *IPAddressVrf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPAddressVrf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPAddressVrf) SetName(v string)`

SetName sets Name field to given value.


### GetRd

`func (o *IPAddressVrf) GetRd() string`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *IPAddressVrf) GetRdOk() (*string, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *IPAddressVrf) SetRd(v string)`

SetRd sets Rd field to given value.

### HasRd

`func (o *IPAddressVrf) HasRd() bool`

HasRd returns a boolean if a field has been set.

### SetRdNil

`func (o *IPAddressVrf) SetRdNil(b bool)`

 SetRdNil sets the value for Rd to be an explicit nil

### UnsetRd
`func (o *IPAddressVrf) UnsetRd()`

UnsetRd ensures that no value is present for Rd, not even an explicit nil
### GetDisplay

`func (o *IPAddressVrf) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *IPAddressVrf) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *IPAddressVrf) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetPrefixCount

`func (o *IPAddressVrf) GetPrefixCount() int32`

GetPrefixCount returns the PrefixCount field if non-nil, zero value otherwise.

### GetPrefixCountOk

`func (o *IPAddressVrf) GetPrefixCountOk() (*int32, bool)`

GetPrefixCountOk returns a tuple with the PrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCount

`func (o *IPAddressVrf) SetPrefixCount(v int32)`

SetPrefixCount sets PrefixCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


