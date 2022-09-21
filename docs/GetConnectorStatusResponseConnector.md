# GetConnectorStatusResponseConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**State**](State.md) |  | [optional] 
**WorkerId** | Pointer to **string** | ID of the worker. | [optional] 

## Methods

### NewGetConnectorStatusResponseConnector

`func NewGetConnectorStatusResponseConnector() *GetConnectorStatusResponseConnector`

NewGetConnectorStatusResponseConnector instantiates a new GetConnectorStatusResponseConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectorStatusResponseConnectorWithDefaults

`func NewGetConnectorStatusResponseConnectorWithDefaults() *GetConnectorStatusResponseConnector`

NewGetConnectorStatusResponseConnectorWithDefaults instantiates a new GetConnectorStatusResponseConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *GetConnectorStatusResponseConnector) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetConnectorStatusResponseConnector) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetConnectorStatusResponseConnector) SetState(v State)`

SetState sets State field to given value.

### HasState

`func (o *GetConnectorStatusResponseConnector) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWorkerId

`func (o *GetConnectorStatusResponseConnector) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *GetConnectorStatusResponseConnector) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *GetConnectorStatusResponseConnector) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *GetConnectorStatusResponseConnector) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


