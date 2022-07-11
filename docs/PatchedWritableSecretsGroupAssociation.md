# PatchedWritableSecretsGroupAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Group** | Pointer to **string** |  | [optional] 
**AccessType** | Pointer to [**AccessTypeEnum**](AccessTypeEnum.md) |  | [optional] 
**SecretType** | Pointer to [**SecretTypeEnum**](SecretTypeEnum.md) |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedWritableSecretsGroupAssociation

`func NewPatchedWritableSecretsGroupAssociation() *PatchedWritableSecretsGroupAssociation`

NewPatchedWritableSecretsGroupAssociation instantiates a new PatchedWritableSecretsGroupAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableSecretsGroupAssociationWithDefaults

`func NewPatchedWritableSecretsGroupAssociationWithDefaults() *PatchedWritableSecretsGroupAssociation`

NewPatchedWritableSecretsGroupAssociationWithDefaults instantiates a new PatchedWritableSecretsGroupAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedWritableSecretsGroupAssociation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedWritableSecretsGroupAssociation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedWritableSecretsGroupAssociation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedWritableSecretsGroupAssociation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedWritableSecretsGroupAssociation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedWritableSecretsGroupAssociation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedWritableSecretsGroupAssociation) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedWritableSecretsGroupAssociation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetGroup

`func (o *PatchedWritableSecretsGroupAssociation) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedWritableSecretsGroupAssociation) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedWritableSecretsGroupAssociation) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedWritableSecretsGroupAssociation) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetAccessType

`func (o *PatchedWritableSecretsGroupAssociation) GetAccessType() AccessTypeEnum`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *PatchedWritableSecretsGroupAssociation) GetAccessTypeOk() (*AccessTypeEnum, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *PatchedWritableSecretsGroupAssociation) SetAccessType(v AccessTypeEnum)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *PatchedWritableSecretsGroupAssociation) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetSecretType

`func (o *PatchedWritableSecretsGroupAssociation) GetSecretType() SecretTypeEnum`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *PatchedWritableSecretsGroupAssociation) GetSecretTypeOk() (*SecretTypeEnum, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *PatchedWritableSecretsGroupAssociation) SetSecretType(v SecretTypeEnum)`

SetSecretType sets SecretType field to given value.

### HasSecretType

`func (o *PatchedWritableSecretsGroupAssociation) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.

### GetSecret

`func (o *PatchedWritableSecretsGroupAssociation) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PatchedWritableSecretsGroupAssociation) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PatchedWritableSecretsGroupAssociation) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PatchedWritableSecretsGroupAssociation) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedWritableSecretsGroupAssociation) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedWritableSecretsGroupAssociation) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedWritableSecretsGroupAssociation) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedWritableSecretsGroupAssociation) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


