# ObjectPermissionUsersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Username** | **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewObjectPermissionUsersInner

`func NewObjectPermissionUsersInner(id string, url string, username string, display string, ) *ObjectPermissionUsersInner`

NewObjectPermissionUsersInner instantiates a new ObjectPermissionUsersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectPermissionUsersInnerWithDefaults

`func NewObjectPermissionUsersInnerWithDefaults() *ObjectPermissionUsersInner`

NewObjectPermissionUsersInnerWithDefaults instantiates a new ObjectPermissionUsersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectPermissionUsersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectPermissionUsersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectPermissionUsersInner) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ObjectPermissionUsersInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ObjectPermissionUsersInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ObjectPermissionUsersInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *ObjectPermissionUsersInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ObjectPermissionUsersInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ObjectPermissionUsersInner) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDisplay

`func (o *ObjectPermissionUsersInner) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ObjectPermissionUsersInner) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ObjectPermissionUsersInner) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


