# ListConnectorsResponseValueInfoTasksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to **string** | Name of the connector. | [optional] 
**Task** | Pointer to **int32** | ID of the task. | [optional] 
**Config** | Pointer to **map[string]string** | Configuration of a connector/task/worker. All keys/values should be strings. | [optional] 

## Methods

### NewListConnectorsResponseValueInfoTasksInner

`func NewListConnectorsResponseValueInfoTasksInner() *ListConnectorsResponseValueInfoTasksInner`

NewListConnectorsResponseValueInfoTasksInner instantiates a new ListConnectorsResponseValueInfoTasksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectorsResponseValueInfoTasksInnerWithDefaults

`func NewListConnectorsResponseValueInfoTasksInnerWithDefaults() *ListConnectorsResponseValueInfoTasksInner`

NewListConnectorsResponseValueInfoTasksInnerWithDefaults instantiates a new ListConnectorsResponseValueInfoTasksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *ListConnectorsResponseValueInfoTasksInner) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ListConnectorsResponseValueInfoTasksInner) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ListConnectorsResponseValueInfoTasksInner) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *ListConnectorsResponseValueInfoTasksInner) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetTask

`func (o *ListConnectorsResponseValueInfoTasksInner) GetTask() int32`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ListConnectorsResponseValueInfoTasksInner) GetTaskOk() (*int32, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ListConnectorsResponseValueInfoTasksInner) SetTask(v int32)`

SetTask sets Task field to given value.

### HasTask

`func (o *ListConnectorsResponseValueInfoTasksInner) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetConfig

`func (o *ListConnectorsResponseValueInfoTasksInner) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListConnectorsResponseValueInfoTasksInner) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListConnectorsResponseValueInfoTasksInner) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListConnectorsResponseValueInfoTasksInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


