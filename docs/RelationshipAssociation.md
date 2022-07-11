# RelationshipAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Relationship** | [**NestedRelationship**](NestedRelationship.md) |  | 
**SourceType** | **string** |  | 
**SourceId** | **string** |  | 
**DestinationType** | **string** |  | 
**DestinationId** | **string** |  | 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewRelationshipAssociation

`func NewRelationshipAssociation(id string, relationship NestedRelationship, sourceType string, sourceId string, destinationType string, destinationId string, display string, ) *RelationshipAssociation`

NewRelationshipAssociation instantiates a new RelationshipAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipAssociationWithDefaults

`func NewRelationshipAssociationWithDefaults() *RelationshipAssociation`

NewRelationshipAssociationWithDefaults instantiates a new RelationshipAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelationshipAssociation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipAssociation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipAssociation) SetId(v string)`

SetId sets Id field to given value.


### GetRelationship

`func (o *RelationshipAssociation) GetRelationship() NestedRelationship`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *RelationshipAssociation) GetRelationshipOk() (*NestedRelationship, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *RelationshipAssociation) SetRelationship(v NestedRelationship)`

SetRelationship sets Relationship field to given value.


### GetSourceType

`func (o *RelationshipAssociation) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *RelationshipAssociation) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *RelationshipAssociation) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetSourceId

`func (o *RelationshipAssociation) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *RelationshipAssociation) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *RelationshipAssociation) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetDestinationType

`func (o *RelationshipAssociation) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *RelationshipAssociation) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *RelationshipAssociation) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.


### GetDestinationId

`func (o *RelationshipAssociation) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *RelationshipAssociation) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *RelationshipAssociation) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetDisplay

`func (o *RelationshipAssociation) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RelationshipAssociation) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RelationshipAssociation) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


