# WritableSecretsGroupAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Group** | **string** |  | 
**AccessType** | [**AccessTypeEnum**](AccessTypeEnum.md) |  | 
**SecretType** | [**SecretTypeEnum**](SecretTypeEnum.md) |  | 
**Secret** | **string** |  | 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewWritableSecretsGroupAssociation

`func NewWritableSecretsGroupAssociation(id string, url string, group string, accessType AccessTypeEnum, secretType SecretTypeEnum, secret string, display string, ) *WritableSecretsGroupAssociation`

NewWritableSecretsGroupAssociation instantiates a new WritableSecretsGroupAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableSecretsGroupAssociationWithDefaults

`func NewWritableSecretsGroupAssociationWithDefaults() *WritableSecretsGroupAssociation`

NewWritableSecretsGroupAssociationWithDefaults instantiates a new WritableSecretsGroupAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableSecretsGroupAssociation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableSecretsGroupAssociation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableSecretsGroupAssociation) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *WritableSecretsGroupAssociation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableSecretsGroupAssociation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableSecretsGroupAssociation) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetGroup

`func (o *WritableSecretsGroupAssociation) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *WritableSecretsGroupAssociation) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *WritableSecretsGroupAssociation) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetAccessType

`func (o *WritableSecretsGroupAssociation) GetAccessType() AccessTypeEnum`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *WritableSecretsGroupAssociation) GetAccessTypeOk() (*AccessTypeEnum, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *WritableSecretsGroupAssociation) SetAccessType(v AccessTypeEnum)`

SetAccessType sets AccessType field to given value.


### GetSecretType

`func (o *WritableSecretsGroupAssociation) GetSecretType() SecretTypeEnum`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *WritableSecretsGroupAssociation) GetSecretTypeOk() (*SecretTypeEnum, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *WritableSecretsGroupAssociation) SetSecretType(v SecretTypeEnum)`

SetSecretType sets SecretType field to given value.


### GetSecret

`func (o *WritableSecretsGroupAssociation) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WritableSecretsGroupAssociation) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WritableSecretsGroupAssociation) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetDisplay

`func (o *WritableSecretsGroupAssociation) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableSecretsGroupAssociation) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableSecretsGroupAssociation) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


