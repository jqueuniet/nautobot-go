# ObjectChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Time** | **time.Time** |  | [readonly] 
**User** | [**JobResultUser**](JobResultUser.md) |  | 
**UserName** | **string** |  | [readonly] 
**RequestId** | **string** |  | [readonly] 
**Action** | [**ObjectChangeAction**](ObjectChangeAction.md) |  | 
**ChangedObjectType** | **string** |  | [readonly] 
**ChangedObjectId** | **string** |  | 
**ChangedObject** | **map[string]interface{}** |  | [readonly] 
**ObjectData** | **map[string]interface{}** |  | [readonly] 

## Methods

### NewObjectChange

`func NewObjectChange(id string, url string, time time.Time, user JobResultUser, userName string, requestId string, action ObjectChangeAction, changedObjectType string, changedObjectId string, changedObject map[string]interface{}, objectData map[string]interface{}, ) *ObjectChange`

NewObjectChange instantiates a new ObjectChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectChangeWithDefaults

`func NewObjectChangeWithDefaults() *ObjectChange`

NewObjectChangeWithDefaults instantiates a new ObjectChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectChange) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectChange) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectChange) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ObjectChange) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ObjectChange) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ObjectChange) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTime

`func (o *ObjectChange) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ObjectChange) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ObjectChange) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetUser

`func (o *ObjectChange) GetUser() JobResultUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ObjectChange) GetUserOk() (*JobResultUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ObjectChange) SetUser(v JobResultUser)`

SetUser sets User field to given value.


### GetUserName

`func (o *ObjectChange) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ObjectChange) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ObjectChange) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetRequestId

`func (o *ObjectChange) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ObjectChange) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ObjectChange) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetAction

`func (o *ObjectChange) GetAction() ObjectChangeAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ObjectChange) GetActionOk() (*ObjectChangeAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ObjectChange) SetAction(v ObjectChangeAction)`

SetAction sets Action field to given value.


### GetChangedObjectType

`func (o *ObjectChange) GetChangedObjectType() string`

GetChangedObjectType returns the ChangedObjectType field if non-nil, zero value otherwise.

### GetChangedObjectTypeOk

`func (o *ObjectChange) GetChangedObjectTypeOk() (*string, bool)`

GetChangedObjectTypeOk returns a tuple with the ChangedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedObjectType

`func (o *ObjectChange) SetChangedObjectType(v string)`

SetChangedObjectType sets ChangedObjectType field to given value.


### GetChangedObjectId

`func (o *ObjectChange) GetChangedObjectId() string`

GetChangedObjectId returns the ChangedObjectId field if non-nil, zero value otherwise.

### GetChangedObjectIdOk

`func (o *ObjectChange) GetChangedObjectIdOk() (*string, bool)`

GetChangedObjectIdOk returns a tuple with the ChangedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedObjectId

`func (o *ObjectChange) SetChangedObjectId(v string)`

SetChangedObjectId sets ChangedObjectId field to given value.


### GetChangedObject

`func (o *ObjectChange) GetChangedObject() map[string]interface{}`

GetChangedObject returns the ChangedObject field if non-nil, zero value otherwise.

### GetChangedObjectOk

`func (o *ObjectChange) GetChangedObjectOk() (*map[string]interface{}, bool)`

GetChangedObjectOk returns a tuple with the ChangedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedObject

`func (o *ObjectChange) SetChangedObject(v map[string]interface{})`

SetChangedObject sets ChangedObject field to given value.


### SetChangedObjectNil

`func (o *ObjectChange) SetChangedObjectNil(b bool)`

 SetChangedObjectNil sets the value for ChangedObject to be an explicit nil

### UnsetChangedObject
`func (o *ObjectChange) UnsetChangedObject()`

UnsetChangedObject ensures that no value is present for ChangedObject, not even an explicit nil
### GetObjectData

`func (o *ObjectChange) GetObjectData() map[string]interface{}`

GetObjectData returns the ObjectData field if non-nil, zero value otherwise.

### GetObjectDataOk

`func (o *ObjectChange) GetObjectDataOk() (*map[string]interface{}, bool)`

GetObjectDataOk returns a tuple with the ObjectData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectData

`func (o *ObjectChange) SetObjectData(v map[string]interface{})`

SetObjectData sets ObjectData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


