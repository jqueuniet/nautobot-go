# InterfaceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceA** | [**InterfaceConnectionInterfaceA**](InterfaceConnectionInterfaceA.md) |  | 
**InterfaceB** | [**NestedInterface**](NestedInterface.md) |  | 
**ConnectedEndpointReachable** | **NullableBool** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewInterfaceConnection

`func NewInterfaceConnection(interfaceA InterfaceConnectionInterfaceA, interfaceB NestedInterface, connectedEndpointReachable NullableBool, display string, ) *InterfaceConnection`

NewInterfaceConnection instantiates a new InterfaceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceConnectionWithDefaults

`func NewInterfaceConnectionWithDefaults() *InterfaceConnection`

NewInterfaceConnectionWithDefaults instantiates a new InterfaceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceA

`func (o *InterfaceConnection) GetInterfaceA() InterfaceConnectionInterfaceA`

GetInterfaceA returns the InterfaceA field if non-nil, zero value otherwise.

### GetInterfaceAOk

`func (o *InterfaceConnection) GetInterfaceAOk() (*InterfaceConnectionInterfaceA, bool)`

GetInterfaceAOk returns a tuple with the InterfaceA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceA

`func (o *InterfaceConnection) SetInterfaceA(v InterfaceConnectionInterfaceA)`

SetInterfaceA sets InterfaceA field to given value.


### GetInterfaceB

`func (o *InterfaceConnection) GetInterfaceB() NestedInterface`

GetInterfaceB returns the InterfaceB field if non-nil, zero value otherwise.

### GetInterfaceBOk

`func (o *InterfaceConnection) GetInterfaceBOk() (*NestedInterface, bool)`

GetInterfaceBOk returns a tuple with the InterfaceB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceB

`func (o *InterfaceConnection) SetInterfaceB(v NestedInterface)`

SetInterfaceB sets InterfaceB field to given value.


### GetConnectedEndpointReachable

`func (o *InterfaceConnection) GetConnectedEndpointReachable() bool`

GetConnectedEndpointReachable returns the ConnectedEndpointReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointReachableOk

`func (o *InterfaceConnection) GetConnectedEndpointReachableOk() (*bool, bool)`

GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointReachable

`func (o *InterfaceConnection) SetConnectedEndpointReachable(v bool)`

SetConnectedEndpointReachable sets ConnectedEndpointReachable field to given value.


### SetConnectedEndpointReachableNil

`func (o *InterfaceConnection) SetConnectedEndpointReachableNil(b bool)`

 SetConnectedEndpointReachableNil sets the value for ConnectedEndpointReachable to be an explicit nil

### UnsetConnectedEndpointReachable
`func (o *InterfaceConnection) UnsetConnectedEndpointReachable()`

UnsetConnectedEndpointReachable ensures that no value is present for ConnectedEndpointReachable, not even an explicit nil
### GetDisplay

`func (o *InterfaceConnection) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *InterfaceConnection) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *InterfaceConnection) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


