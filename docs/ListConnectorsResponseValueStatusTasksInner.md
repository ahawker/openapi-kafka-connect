# ListConnectorsResponseValueStatusTasksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the task. | [optional] 
**State** | Pointer to [**State**](State.md) |  | [optional] 
**Trace** | Pointer to **string** | Stack trace information if the task has failed. | [optional] 
**WorkerId** | Pointer to **string** | ID of the worker. | [optional] 

## Methods

### NewListConnectorsResponseValueStatusTasksInner

`func NewListConnectorsResponseValueStatusTasksInner() *ListConnectorsResponseValueStatusTasksInner`

NewListConnectorsResponseValueStatusTasksInner instantiates a new ListConnectorsResponseValueStatusTasksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectorsResponseValueStatusTasksInnerWithDefaults

`func NewListConnectorsResponseValueStatusTasksInnerWithDefaults() *ListConnectorsResponseValueStatusTasksInner`

NewListConnectorsResponseValueStatusTasksInnerWithDefaults instantiates a new ListConnectorsResponseValueStatusTasksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListConnectorsResponseValueStatusTasksInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListConnectorsResponseValueStatusTasksInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListConnectorsResponseValueStatusTasksInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ListConnectorsResponseValueStatusTasksInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *ListConnectorsResponseValueStatusTasksInner) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ListConnectorsResponseValueStatusTasksInner) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ListConnectorsResponseValueStatusTasksInner) SetState(v State)`

SetState sets State field to given value.

### HasState

`func (o *ListConnectorsResponseValueStatusTasksInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTrace

`func (o *ListConnectorsResponseValueStatusTasksInner) GetTrace() string`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *ListConnectorsResponseValueStatusTasksInner) GetTraceOk() (*string, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *ListConnectorsResponseValueStatusTasksInner) SetTrace(v string)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *ListConnectorsResponseValueStatusTasksInner) HasTrace() bool`

HasTrace returns a boolean if a field has been set.

### GetWorkerId

`func (o *ListConnectorsResponseValueStatusTasksInner) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *ListConnectorsResponseValueStatusTasksInner) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *ListConnectorsResponseValueStatusTasksInner) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *ListConnectorsResponseValueStatusTasksInner) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


