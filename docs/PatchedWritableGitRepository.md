# PatchedWritableGitRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**RemoteUrl** | Pointer to **string** | Only HTTP and HTTPS URLs are presently supported | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**SecretsGroup** | Pointer to **NullableString** |  | [optional] 
**CurrentHead** | Pointer to **string** | Commit hash of the most recent fetch from the selected branch. Used for syncing between workers. | [optional] 
**ProvidedContents** | Pointer to [**[]GitRepositoryProvidedContentsInner**](GitRepositoryProvidedContentsInner.md) |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ComputedFields** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedWritableGitRepository

`func NewPatchedWritableGitRepository() *PatchedWritableGitRepository`

NewPatchedWritableGitRepository instantiates a new PatchedWritableGitRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableGitRepositoryWithDefaults

`func NewPatchedWritableGitRepositoryWithDefaults() *PatchedWritableGitRepository`

NewPatchedWritableGitRepositoryWithDefaults instantiates a new PatchedWritableGitRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedWritableGitRepository) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedWritableGitRepository) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedWritableGitRepository) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedWritableGitRepository) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedWritableGitRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedWritableGitRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedWritableGitRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedWritableGitRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *PatchedWritableGitRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableGitRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableGitRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableGitRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedWritableGitRepository) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedWritableGitRepository) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedWritableGitRepository) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedWritableGitRepository) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetRemoteUrl

`func (o *PatchedWritableGitRepository) GetRemoteUrl() string`

GetRemoteUrl returns the RemoteUrl field if non-nil, zero value otherwise.

### GetRemoteUrlOk

`func (o *PatchedWritableGitRepository) GetRemoteUrlOk() (*string, bool)`

GetRemoteUrlOk returns a tuple with the RemoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUrl

`func (o *PatchedWritableGitRepository) SetRemoteUrl(v string)`

SetRemoteUrl sets RemoteUrl field to given value.

### HasRemoteUrl

`func (o *PatchedWritableGitRepository) HasRemoteUrl() bool`

HasRemoteUrl returns a boolean if a field has been set.

### GetBranch

`func (o *PatchedWritableGitRepository) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PatchedWritableGitRepository) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PatchedWritableGitRepository) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *PatchedWritableGitRepository) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetToken

`func (o *PatchedWritableGitRepository) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PatchedWritableGitRepository) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PatchedWritableGitRepository) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PatchedWritableGitRepository) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsername

`func (o *PatchedWritableGitRepository) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedWritableGitRepository) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedWritableGitRepository) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedWritableGitRepository) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *PatchedWritableGitRepository) GetSecretsGroup() string`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *PatchedWritableGitRepository) GetSecretsGroupOk() (*string, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *PatchedWritableGitRepository) SetSecretsGroup(v string)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *PatchedWritableGitRepository) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *PatchedWritableGitRepository) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *PatchedWritableGitRepository) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetCurrentHead

`func (o *PatchedWritableGitRepository) GetCurrentHead() string`

GetCurrentHead returns the CurrentHead field if non-nil, zero value otherwise.

### GetCurrentHeadOk

`func (o *PatchedWritableGitRepository) GetCurrentHeadOk() (*string, bool)`

GetCurrentHeadOk returns a tuple with the CurrentHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentHead

`func (o *PatchedWritableGitRepository) SetCurrentHead(v string)`

SetCurrentHead sets CurrentHead field to given value.

### HasCurrentHead

`func (o *PatchedWritableGitRepository) HasCurrentHead() bool`

HasCurrentHead returns a boolean if a field has been set.

### GetProvidedContents

`func (o *PatchedWritableGitRepository) GetProvidedContents() []GitRepositoryProvidedContentsInner`

GetProvidedContents returns the ProvidedContents field if non-nil, zero value otherwise.

### GetProvidedContentsOk

`func (o *PatchedWritableGitRepository) GetProvidedContentsOk() (*[]GitRepositoryProvidedContentsInner, bool)`

GetProvidedContentsOk returns a tuple with the ProvidedContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidedContents

`func (o *PatchedWritableGitRepository) SetProvidedContents(v []GitRepositoryProvidedContentsInner)`

SetProvidedContents sets ProvidedContents field to given value.

### HasProvidedContents

`func (o *PatchedWritableGitRepository) HasProvidedContents() bool`

HasProvidedContents returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedWritableGitRepository) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedWritableGitRepository) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedWritableGitRepository) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedWritableGitRepository) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PatchedWritableGitRepository) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PatchedWritableGitRepository) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PatchedWritableGitRepository) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PatchedWritableGitRepository) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableGitRepository) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableGitRepository) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableGitRepository) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableGitRepository) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetComputedFields

`func (o *PatchedWritableGitRepository) GetComputedFields() map[string]interface{}`

GetComputedFields returns the ComputedFields field if non-nil, zero value otherwise.

### GetComputedFieldsOk

`func (o *PatchedWritableGitRepository) GetComputedFieldsOk() (*map[string]interface{}, bool)`

GetComputedFieldsOk returns a tuple with the ComputedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFields

`func (o *PatchedWritableGitRepository) SetComputedFields(v map[string]interface{})`

SetComputedFields sets ComputedFields field to given value.

### HasComputedFields

`func (o *PatchedWritableGitRepository) HasComputedFields() bool`

HasComputedFields returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedWritableGitRepository) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedWritableGitRepository) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedWritableGitRepository) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedWritableGitRepository) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


