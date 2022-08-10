# AvailablePrefixRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrefixLength** | **int32** |  | 
**Status** | [**AvailablePrefixRequestStatusEnum**](AvailablePrefixRequestStatusEnum.md) |  | 

## Methods

### NewAvailablePrefixRequest

`func NewAvailablePrefixRequest(prefixLength int32, status AvailablePrefixRequestStatusEnum, ) *AvailablePrefixRequest`

NewAvailablePrefixRequest instantiates a new AvailablePrefixRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailablePrefixRequestWithDefaults

`func NewAvailablePrefixRequestWithDefaults() *AvailablePrefixRequest`

NewAvailablePrefixRequestWithDefaults instantiates a new AvailablePrefixRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefixLength

`func (o *AvailablePrefixRequest) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *AvailablePrefixRequest) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *AvailablePrefixRequest) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.


### GetStatus

`func (o *AvailablePrefixRequest) GetStatus() AvailablePrefixRequestStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AvailablePrefixRequest) GetStatusOk() (*AvailablePrefixRequestStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AvailablePrefixRequest) SetStatus(v AvailablePrefixRequestStatusEnum)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


