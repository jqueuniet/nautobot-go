# WritableGitRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**RemoteUrl** | **string** | Only HTTP and HTTPS URLs are presently supported | 
**Branch** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**SecretsGroup** | Pointer to **NullableString** |  | [optional] 
**CurrentHead** | Pointer to **string** | Commit hash of the most recent fetch from the selected branch. Used for syncing between workers. | [optional] 
**ProvidedContents** | Pointer to [**[]GitRepositoryProvidedContentsInner**](GitRepositoryProvidedContentsInner.md) |  | [optional] 
**Created** | **string** |  | [readonly] 
**LastUpdated** | **time.Time** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ComputedFields** | **map[string]interface{}** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewWritableGitRepository

`func NewWritableGitRepository(id string, url string, name string, remoteUrl string, created string, lastUpdated time.Time, computedFields map[string]interface{}, display string, ) *WritableGitRepository`

NewWritableGitRepository instantiates a new WritableGitRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableGitRepositoryWithDefaults

`func NewWritableGitRepositoryWithDefaults() *WritableGitRepository`

NewWritableGitRepositoryWithDefaults instantiates a new WritableGitRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableGitRepository) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableGitRepository) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableGitRepository) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *WritableGitRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableGitRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableGitRepository) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *WritableGitRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableGitRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableGitRepository) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *WritableGitRepository) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WritableGitRepository) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WritableGitRepository) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *WritableGitRepository) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetRemoteUrl

`func (o *WritableGitRepository) GetRemoteUrl() string`

GetRemoteUrl returns the RemoteUrl field if non-nil, zero value otherwise.

### GetRemoteUrlOk

`func (o *WritableGitRepository) GetRemoteUrlOk() (*string, bool)`

GetRemoteUrlOk returns a tuple with the RemoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUrl

`func (o *WritableGitRepository) SetRemoteUrl(v string)`

SetRemoteUrl sets RemoteUrl field to given value.


### GetBranch

`func (o *WritableGitRepository) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *WritableGitRepository) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *WritableGitRepository) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *WritableGitRepository) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetToken

`func (o *WritableGitRepository) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *WritableGitRepository) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *WritableGitRepository) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *WritableGitRepository) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsername

`func (o *WritableGitRepository) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *WritableGitRepository) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *WritableGitRepository) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *WritableGitRepository) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *WritableGitRepository) GetSecretsGroup() string`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *WritableGitRepository) GetSecretsGroupOk() (*string, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *WritableGitRepository) SetSecretsGroup(v string)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *WritableGitRepository) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *WritableGitRepository) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *WritableGitRepository) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetCurrentHead

`func (o *WritableGitRepository) GetCurrentHead() string`

GetCurrentHead returns the CurrentHead field if non-nil, zero value otherwise.

### GetCurrentHeadOk

`func (o *WritableGitRepository) GetCurrentHeadOk() (*string, bool)`

GetCurrentHeadOk returns a tuple with the CurrentHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentHead

`func (o *WritableGitRepository) SetCurrentHead(v string)`

SetCurrentHead sets CurrentHead field to given value.

### HasCurrentHead

`func (o *WritableGitRepository) HasCurrentHead() bool`

HasCurrentHead returns a boolean if a field has been set.

### GetProvidedContents

`func (o *WritableGitRepository) GetProvidedContents() []GitRepositoryProvidedContentsInner`

GetProvidedContents returns the ProvidedContents field if non-nil, zero value otherwise.

### GetProvidedContentsOk

`func (o *WritableGitRepository) GetProvidedContentsOk() (*[]GitRepositoryProvidedContentsInner, bool)`

GetProvidedContentsOk returns a tuple with the ProvidedContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidedContents

`func (o *WritableGitRepository) SetProvidedContents(v []GitRepositoryProvidedContentsInner)`

SetProvidedContents sets ProvidedContents field to given value.

### HasProvidedContents

`func (o *WritableGitRepository) HasProvidedContents() bool`

HasProvidedContents returns a boolean if a field has been set.

### GetCreated

`func (o *WritableGitRepository) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritableGitRepository) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritableGitRepository) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *WritableGitRepository) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritableGitRepository) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritableGitRepository) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetCustomFields

`func (o *WritableGitRepository) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableGitRepository) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableGitRepository) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableGitRepository) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetComputedFields

`func (o *WritableGitRepository) GetComputedFields() map[string]interface{}`

GetComputedFields returns the ComputedFields field if non-nil, zero value otherwise.

### GetComputedFieldsOk

`func (o *WritableGitRepository) GetComputedFieldsOk() (*map[string]interface{}, bool)`

GetComputedFieldsOk returns a tuple with the ComputedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFields

`func (o *WritableGitRepository) SetComputedFields(v map[string]interface{})`

SetComputedFields sets ComputedFields field to given value.


### GetDisplay

`func (o *WritableGitRepository) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableGitRepository) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableGitRepository) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


