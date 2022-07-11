# PatchedCustomLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**TargetUrl** | Pointer to **string** | Jinja2 template code for link URL. Reference the object as &lt;code&gt;{{ obj }}&lt;/code&gt; such as &lt;code&gt;{{ obj.platform.slug }}&lt;/code&gt;. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** | Jinja2 template code for link text. Reference the object as &lt;code&gt;{{ obj }}&lt;/code&gt; such as &lt;code&gt;{{ obj.platform.slug }}&lt;/code&gt;. Links which render as empty text will not be displayed. | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**GroupName** | Pointer to **string** | Links with the same group will appear as a dropdown menu | [optional] 
**ButtonClass** | Pointer to [**NullableButtonClassEnum**](ButtonClassEnum.md) | The class of the first link in a group will be used for the dropdown button | [optional] 
**NewWindow** | Pointer to **bool** | Force link to open in a new window | [optional] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedCustomLink

`func NewPatchedCustomLink() *PatchedCustomLink`

NewPatchedCustomLink instantiates a new PatchedCustomLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCustomLinkWithDefaults

`func NewPatchedCustomLinkWithDefaults() *PatchedCustomLink`

NewPatchedCustomLinkWithDefaults instantiates a new PatchedCustomLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedCustomLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedCustomLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedCustomLink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedCustomLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedCustomLink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedCustomLink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedCustomLink) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedCustomLink) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTargetUrl

`func (o *PatchedCustomLink) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *PatchedCustomLink) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *PatchedCustomLink) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *PatchedCustomLink) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetName

`func (o *PatchedCustomLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedCustomLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedCustomLink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedCustomLink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContentType

`func (o *PatchedCustomLink) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedCustomLink) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedCustomLink) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedCustomLink) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetText

`func (o *PatchedCustomLink) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PatchedCustomLink) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PatchedCustomLink) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PatchedCustomLink) HasText() bool`

HasText returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedCustomLink) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedCustomLink) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedCustomLink) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedCustomLink) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetGroupName

`func (o *PatchedCustomLink) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *PatchedCustomLink) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *PatchedCustomLink) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *PatchedCustomLink) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetButtonClass

`func (o *PatchedCustomLink) GetButtonClass() ButtonClassEnum`

GetButtonClass returns the ButtonClass field if non-nil, zero value otherwise.

### GetButtonClassOk

`func (o *PatchedCustomLink) GetButtonClassOk() (*ButtonClassEnum, bool)`

GetButtonClassOk returns a tuple with the ButtonClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonClass

`func (o *PatchedCustomLink) SetButtonClass(v ButtonClassEnum)`

SetButtonClass sets ButtonClass field to given value.

### HasButtonClass

`func (o *PatchedCustomLink) HasButtonClass() bool`

HasButtonClass returns a boolean if a field has been set.

### SetButtonClassNil

`func (o *PatchedCustomLink) SetButtonClassNil(b bool)`

 SetButtonClassNil sets the value for ButtonClass to be an explicit nil

### UnsetButtonClass
`func (o *PatchedCustomLink) UnsetButtonClass()`

UnsetButtonClass ensures that no value is present for ButtonClass, not even an explicit nil
### GetNewWindow

`func (o *PatchedCustomLink) GetNewWindow() bool`

GetNewWindow returns the NewWindow field if non-nil, zero value otherwise.

### GetNewWindowOk

`func (o *PatchedCustomLink) GetNewWindowOk() (*bool, bool)`

GetNewWindowOk returns a tuple with the NewWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewWindow

`func (o *PatchedCustomLink) SetNewWindow(v bool)`

SetNewWindow sets NewWindow field to given value.

### HasNewWindow

`func (o *PatchedCustomLink) HasNewWindow() bool`

HasNewWindow returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedCustomLink) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedCustomLink) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedCustomLink) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedCustomLink) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


