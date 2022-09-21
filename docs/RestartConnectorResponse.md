# RestartConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the connector. | [optional] 
**Connector** | Pointer to [**GetConnectorStatusResponseConnector**](GetConnectorStatusResponseConnector.md) |  | [optional] 
**Tasks** | Pointer to [**[]GetConnectorStatusResponseTasksInner**](GetConnectorStatusResponseTasksInner.md) | List of task status for the connector. | [optional] 

## Methods

### NewRestartConnectorResponse

`func NewRestartConnectorResponse() *RestartConnectorResponse`

NewRestartConnectorResponse instantiates a new RestartConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartConnectorResponseWithDefaults

`func NewRestartConnectorResponseWithDefaults() *RestartConnectorResponse`

NewRestartConnectorResponseWithDefaults instantiates a new RestartConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RestartConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestartConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestartConnectorResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestartConnectorResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConnector

`func (o *RestartConnectorResponse) GetConnector() GetConnectorStatusResponseConnector`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *RestartConnectorResponse) GetConnectorOk() (*GetConnectorStatusResponseConnector, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *RestartConnectorResponse) SetConnector(v GetConnectorStatusResponseConnector)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *RestartConnectorResponse) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetTasks

`func (o *RestartConnectorResponse) GetTasks() []GetConnectorStatusResponseTasksInner`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *RestartConnectorResponse) GetTasksOk() (*[]GetConnectorStatusResponseTasksInner, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *RestartConnectorResponse) SetTasks(v []GetConnectorStatusResponseTasksInner)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *RestartConnectorResponse) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


