# GetConnectorTaskStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**State**](State.md) |  | [optional] 
**WorkerId** | Pointer to **string** | ID of the worker. | [optional] 

## Methods

### NewGetConnectorTaskStatusResponse

`func NewGetConnectorTaskStatusResponse() *GetConnectorTaskStatusResponse`

NewGetConnectorTaskStatusResponse instantiates a new GetConnectorTaskStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectorTaskStatusResponseWithDefaults

`func NewGetConnectorTaskStatusResponseWithDefaults() *GetConnectorTaskStatusResponse`

NewGetConnectorTaskStatusResponseWithDefaults instantiates a new GetConnectorTaskStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *GetConnectorTaskStatusResponse) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetConnectorTaskStatusResponse) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetConnectorTaskStatusResponse) SetState(v State)`

SetState sets State field to given value.

### HasState

`func (o *GetConnectorTaskStatusResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWorkerId

`func (o *GetConnectorTaskStatusResponse) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *GetConnectorTaskStatusResponse) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *GetConnectorTaskStatusResponse) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *GetConnectorTaskStatusResponse) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


