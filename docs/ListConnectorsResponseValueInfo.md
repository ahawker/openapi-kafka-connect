# ListConnectorsResponseValueInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the connector. | [optional] 
**Type** | Pointer to [**ConnectorType**](ConnectorType.md) |  | [optional] 
**Config** | Pointer to **map[string]string** | Configuration of a connector/task/worker. All keys/values should be strings. | [optional] 
**Tasks** | Pointer to [**[]ListConnectorsResponseValueInfoTasksInner**](ListConnectorsResponseValueInfoTasksInner.md) | List of tasks for the connector. | [optional] 

## Methods

### NewListConnectorsResponseValueInfo

`func NewListConnectorsResponseValueInfo() *ListConnectorsResponseValueInfo`

NewListConnectorsResponseValueInfo instantiates a new ListConnectorsResponseValueInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectorsResponseValueInfoWithDefaults

`func NewListConnectorsResponseValueInfoWithDefaults() *ListConnectorsResponseValueInfo`

NewListConnectorsResponseValueInfoWithDefaults instantiates a new ListConnectorsResponseValueInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListConnectorsResponseValueInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListConnectorsResponseValueInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListConnectorsResponseValueInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListConnectorsResponseValueInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ListConnectorsResponseValueInfo) GetType() ConnectorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListConnectorsResponseValueInfo) GetTypeOk() (*ConnectorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListConnectorsResponseValueInfo) SetType(v ConnectorType)`

SetType sets Type field to given value.

### HasType

`func (o *ListConnectorsResponseValueInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *ListConnectorsResponseValueInfo) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListConnectorsResponseValueInfo) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListConnectorsResponseValueInfo) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListConnectorsResponseValueInfo) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTasks

`func (o *ListConnectorsResponseValueInfo) GetTasks() []ListConnectorsResponseValueInfoTasksInner`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ListConnectorsResponseValueInfo) GetTasksOk() (*[]ListConnectorsResponseValueInfoTasksInner, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ListConnectorsResponseValueInfo) SetTasks(v []ListConnectorsResponseValueInfoTasksInner)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ListConnectorsResponseValueInfo) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


