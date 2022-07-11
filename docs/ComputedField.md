# ComputedField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Slug** | Pointer to **string** | Internal field name | [optional] 
**Label** | **string** | Name of the field as displayed to users | 
**Description** | Pointer to **string** |  | [optional] 
**ContentType** | **string** |  | 
**Template** | **string** | Jinja2 template code for field value | 
**FallbackValue** | Pointer to **string** | Fallback value (if any) to be output for the field in the case of a template rendering error. | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewComputedField

`func NewComputedField(id string, url string, label string, contentType string, template string, display string, ) *ComputedField`

NewComputedField instantiates a new ComputedField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputedFieldWithDefaults

`func NewComputedFieldWithDefaults() *ComputedField`

NewComputedFieldWithDefaults instantiates a new ComputedField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputedField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputedField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputedField) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ComputedField) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ComputedField) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ComputedField) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSlug

`func (o *ComputedField) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ComputedField) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ComputedField) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ComputedField) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetLabel

`func (o *ComputedField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ComputedField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ComputedField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *ComputedField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComputedField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComputedField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComputedField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContentType

`func (o *ComputedField) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ComputedField) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ComputedField) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetTemplate

`func (o *ComputedField) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ComputedField) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ComputedField) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetFallbackValue

`func (o *ComputedField) GetFallbackValue() string`

GetFallbackValue returns the FallbackValue field if non-nil, zero value otherwise.

### GetFallbackValueOk

`func (o *ComputedField) GetFallbackValueOk() (*string, bool)`

GetFallbackValueOk returns a tuple with the FallbackValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackValue

`func (o *ComputedField) SetFallbackValue(v string)`

SetFallbackValue sets FallbackValue field to given value.

### HasFallbackValue

`func (o *ComputedField) HasFallbackValue() bool`

HasFallbackValue returns a boolean if a field has been set.

### GetWeight

`func (o *ComputedField) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ComputedField) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ComputedField) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ComputedField) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisplay

`func (o *ComputedField) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ComputedField) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ComputedField) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


