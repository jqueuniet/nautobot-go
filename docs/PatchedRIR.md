# PatchedRIR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**IsPrivate** | Pointer to **bool** | IP space managed by this RIR is considered private | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AggregateCount** | Pointer to **int32** |  | [optional] [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedRIR

`func NewPatchedRIR() *PatchedRIR`

NewPatchedRIR instantiates a new PatchedRIR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRIRWithDefaults

`func NewPatchedRIRWithDefaults() *PatchedRIR`

NewPatchedRIRWithDefaults instantiates a new PatchedRIR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedRIR) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedRIR) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedRIR) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedRIR) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedRIR) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedRIR) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedRIR) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedRIR) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *PatchedRIR) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedRIR) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedRIR) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedRIR) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedRIR) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedRIR) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedRIR) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedRIR) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetIsPrivate

`func (o *PatchedRIR) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *PatchedRIR) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *PatchedRIR) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *PatchedRIR) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedRIR) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedRIR) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedRIR) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedRIR) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAggregateCount

`func (o *PatchedRIR) GetAggregateCount() int32`

GetAggregateCount returns the AggregateCount field if non-nil, zero value otherwise.

### GetAggregateCountOk

`func (o *PatchedRIR) GetAggregateCountOk() (*int32, bool)`

GetAggregateCountOk returns a tuple with the AggregateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateCount

`func (o *PatchedRIR) SetAggregateCount(v int32)`

SetAggregateCount sets AggregateCount field to given value.

### HasAggregateCount

`func (o *PatchedRIR) HasAggregateCount() bool`

HasAggregateCount returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedRIR) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedRIR) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedRIR) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedRIR) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedRIR) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedRIR) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedRIR) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedRIR) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PatchedRIR) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PatchedRIR) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PatchedRIR) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PatchedRIR) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedRIR) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedRIR) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedRIR) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedRIR) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


