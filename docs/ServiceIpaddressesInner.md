# ServiceIpaddressesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Family** | **int32** |  | [readonly] 
**Address** | **string** |  | 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewServiceIpaddressesInner

`func NewServiceIpaddressesInner(id string, url string, family int32, address string, display string, ) *ServiceIpaddressesInner`

NewServiceIpaddressesInner instantiates a new ServiceIpaddressesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceIpaddressesInnerWithDefaults

`func NewServiceIpaddressesInnerWithDefaults() *ServiceIpaddressesInner`

NewServiceIpaddressesInnerWithDefaults instantiates a new ServiceIpaddressesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceIpaddressesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceIpaddressesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceIpaddressesInner) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ServiceIpaddressesInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServiceIpaddressesInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServiceIpaddressesInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFamily

`func (o *ServiceIpaddressesInner) GetFamily() int32`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *ServiceIpaddressesInner) GetFamilyOk() (*int32, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *ServiceIpaddressesInner) SetFamily(v int32)`

SetFamily sets Family field to given value.


### GetAddress

`func (o *ServiceIpaddressesInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ServiceIpaddressesInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ServiceIpaddressesInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetDisplay

`func (o *ServiceIpaddressesInner) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ServiceIpaddressesInner) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ServiceIpaddressesInner) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


