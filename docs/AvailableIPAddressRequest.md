# AvailableIPAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**IPAddressStatusChoices**](IPAddressStatusChoices.md) |  | 

## Methods

### NewAvailableIPAddressRequest

`func NewAvailableIPAddressRequest(status IPAddressStatusChoices, ) *AvailableIPAddressRequest`

NewAvailableIPAddressRequest instantiates a new AvailableIPAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableIPAddressRequestWithDefaults

`func NewAvailableIPAddressRequestWithDefaults() *AvailableIPAddressRequest`

NewAvailableIPAddressRequestWithDefaults instantiates a new AvailableIPAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AvailableIPAddressRequest) GetStatus() IPAddressStatusChoices`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AvailableIPAddressRequest) GetStatusOk() (*IPAddressStatusChoices, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AvailableIPAddressRequest) SetStatus(v IPAddressStatusChoices)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


