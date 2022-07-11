# GitRepository

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
**SecretsGroup** | Pointer to [**NullableDeviceSecretsGroup**](DeviceSecretsGroup.md) |  | [optional] 
**CurrentHead** | Pointer to **string** | Commit hash of the most recent fetch from the selected branch. Used for syncing between workers. | [optional] 
**ProvidedContents** | Pointer to [**[]GitRepositoryProvidedContentsInner**](GitRepositoryProvidedContentsInner.md) |  | [optional] 
**Created** | **string** |  | [readonly] 
**LastUpdated** | **time.Time** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewGitRepository

`func NewGitRepository(id string, url string, name string, remoteUrl string, created string, lastUpdated time.Time, display string, ) *GitRepository`

NewGitRepository instantiates a new GitRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRepositoryWithDefaults

`func NewGitRepositoryWithDefaults() *GitRepository`

NewGitRepositoryWithDefaults instantiates a new GitRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GitRepository) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitRepository) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitRepository) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *GitRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitRepository) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *GitRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitRepository) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *GitRepository) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *GitRepository) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *GitRepository) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *GitRepository) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetRemoteUrl

`func (o *GitRepository) GetRemoteUrl() string`

GetRemoteUrl returns the RemoteUrl field if non-nil, zero value otherwise.

### GetRemoteUrlOk

`func (o *GitRepository) GetRemoteUrlOk() (*string, bool)`

GetRemoteUrlOk returns a tuple with the RemoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUrl

`func (o *GitRepository) SetRemoteUrl(v string)`

SetRemoteUrl sets RemoteUrl field to given value.


### GetBranch

`func (o *GitRepository) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GitRepository) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GitRepository) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *GitRepository) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetToken

`func (o *GitRepository) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GitRepository) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GitRepository) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GitRepository) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsername

`func (o *GitRepository) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GitRepository) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GitRepository) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GitRepository) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *GitRepository) GetSecretsGroup() DeviceSecretsGroup`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *GitRepository) GetSecretsGroupOk() (*DeviceSecretsGroup, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *GitRepository) SetSecretsGroup(v DeviceSecretsGroup)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *GitRepository) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *GitRepository) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *GitRepository) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetCurrentHead

`func (o *GitRepository) GetCurrentHead() string`

GetCurrentHead returns the CurrentHead field if non-nil, zero value otherwise.

### GetCurrentHeadOk

`func (o *GitRepository) GetCurrentHeadOk() (*string, bool)`

GetCurrentHeadOk returns a tuple with the CurrentHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentHead

`func (o *GitRepository) SetCurrentHead(v string)`

SetCurrentHead sets CurrentHead field to given value.

### HasCurrentHead

`func (o *GitRepository) HasCurrentHead() bool`

HasCurrentHead returns a boolean if a field has been set.

### GetProvidedContents

`func (o *GitRepository) GetProvidedContents() []GitRepositoryProvidedContentsInner`

GetProvidedContents returns the ProvidedContents field if non-nil, zero value otherwise.

### GetProvidedContentsOk

`func (o *GitRepository) GetProvidedContentsOk() (*[]GitRepositoryProvidedContentsInner, bool)`

GetProvidedContentsOk returns a tuple with the ProvidedContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidedContents

`func (o *GitRepository) SetProvidedContents(v []GitRepositoryProvidedContentsInner)`

SetProvidedContents sets ProvidedContents field to given value.

### HasProvidedContents

`func (o *GitRepository) HasProvidedContents() bool`

HasProvidedContents returns a boolean if a field has been set.

### GetCreated

`func (o *GitRepository) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GitRepository) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GitRepository) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *GitRepository) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GitRepository) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GitRepository) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetCustomFields

`func (o *GitRepository) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GitRepository) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GitRepository) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GitRepository) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDisplay

`func (o *GitRepository) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *GitRepository) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *GitRepository) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


