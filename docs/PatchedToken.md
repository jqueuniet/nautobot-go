# PatchedToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**WriteEnabled** | Pointer to **bool** | Permit create/update/delete operations using this key | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedToken

`func NewPatchedToken() *PatchedToken`

NewPatchedToken instantiates a new PatchedToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTokenWithDefaults

`func NewPatchedTokenWithDefaults() *PatchedToken`

NewPatchedTokenWithDefaults instantiates a new PatchedToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedToken) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedToken) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedToken) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedToken) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedToken) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedToken) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedToken) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedToken) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedToken) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedToken) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedToken) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedToken) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpires

`func (o *PatchedToken) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *PatchedToken) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *PatchedToken) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *PatchedToken) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *PatchedToken) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *PatchedToken) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetKey

`func (o *PatchedToken) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PatchedToken) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PatchedToken) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PatchedToken) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetWriteEnabled

`func (o *PatchedToken) GetWriteEnabled() bool`

GetWriteEnabled returns the WriteEnabled field if non-nil, zero value otherwise.

### GetWriteEnabledOk

`func (o *PatchedToken) GetWriteEnabledOk() (*bool, bool)`

GetWriteEnabledOk returns a tuple with the WriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteEnabled

`func (o *PatchedToken) SetWriteEnabled(v bool)`

SetWriteEnabled sets WriteEnabled field to given value.

### HasWriteEnabled

`func (o *PatchedToken) HasWriteEnabled() bool`

HasWriteEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


