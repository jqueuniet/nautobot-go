# PatchedImageAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**ContentType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**ImageHeight** | Pointer to **int32** |  | [optional] 
**ImageWidth** | Pointer to **int32** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedImageAttachment

`func NewPatchedImageAttachment() *PatchedImageAttachment`

NewPatchedImageAttachment instantiates a new PatchedImageAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedImageAttachmentWithDefaults

`func NewPatchedImageAttachmentWithDefaults() *PatchedImageAttachment`

NewPatchedImageAttachmentWithDefaults instantiates a new PatchedImageAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedImageAttachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedImageAttachment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedImageAttachment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedImageAttachment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedImageAttachment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedImageAttachment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedImageAttachment) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedImageAttachment) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetContentType

`func (o *PatchedImageAttachment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedImageAttachment) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedImageAttachment) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedImageAttachment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedImageAttachment) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedImageAttachment) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedImageAttachment) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedImageAttachment) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetParent

`func (o *PatchedImageAttachment) GetParent() map[string]interface{}`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedImageAttachment) GetParentOk() (*map[string]interface{}, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedImageAttachment) SetParent(v map[string]interface{})`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedImageAttachment) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetName

`func (o *PatchedImageAttachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedImageAttachment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedImageAttachment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedImageAttachment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *PatchedImageAttachment) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PatchedImageAttachment) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PatchedImageAttachment) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PatchedImageAttachment) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImageHeight

`func (o *PatchedImageAttachment) GetImageHeight() int32`

GetImageHeight returns the ImageHeight field if non-nil, zero value otherwise.

### GetImageHeightOk

`func (o *PatchedImageAttachment) GetImageHeightOk() (*int32, bool)`

GetImageHeightOk returns a tuple with the ImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageHeight

`func (o *PatchedImageAttachment) SetImageHeight(v int32)`

SetImageHeight sets ImageHeight field to given value.

### HasImageHeight

`func (o *PatchedImageAttachment) HasImageHeight() bool`

HasImageHeight returns a boolean if a field has been set.

### GetImageWidth

`func (o *PatchedImageAttachment) GetImageWidth() int32`

GetImageWidth returns the ImageWidth field if non-nil, zero value otherwise.

### GetImageWidthOk

`func (o *PatchedImageAttachment) GetImageWidthOk() (*int32, bool)`

GetImageWidthOk returns a tuple with the ImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageWidth

`func (o *PatchedImageAttachment) SetImageWidth(v int32)`

SetImageWidth sets ImageWidth field to given value.

### HasImageWidth

`func (o *PatchedImageAttachment) HasImageWidth() bool`

HasImageWidth returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedImageAttachment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedImageAttachment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedImageAttachment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedImageAttachment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedImageAttachment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedImageAttachment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedImageAttachment) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedImageAttachment) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


