# ListConnectorsResponseValueStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the connector. | [optional] 
**Type** | Pointer to [**ConnectorType**](ConnectorType.md) |  | [optional] 
**Connector** | Pointer to [**ListConnectorsResponseValueStatusConnector**](ListConnectorsResponseValueStatusConnector.md) |  | [optional] 
**Tasks** | Pointer to [**[]ListConnectorsResponseValueStatusTasksInner**](ListConnectorsResponseValueStatusTasksInner.md) | List of task statuses for the connector. | [optional] 

## Methods

### NewListConnectorsResponseValueStatus

`func NewListConnectorsResponseValueStatus() *ListConnectorsResponseValueStatus`

NewListConnectorsResponseValueStatus instantiates a new ListConnectorsResponseValueStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectorsResponseValueStatusWithDefaults

`func NewListConnectorsResponseValueStatusWithDefaults() *ListConnectorsResponseValueStatus`

NewListConnectorsResponseValueStatusWithDefaults instantiates a new ListConnectorsResponseValueStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListConnectorsResponseValueStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListConnectorsResponseValueStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListConnectorsResponseValueStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListConnectorsResponseValueStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ListConnectorsResponseValueStatus) GetType() ConnectorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListConnectorsResponseValueStatus) GetTypeOk() (*ConnectorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListConnectorsResponseValueStatus) SetType(v ConnectorType)`

SetType sets Type field to given value.

### HasType

`func (o *ListConnectorsResponseValueStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConnector

`func (o *ListConnectorsResponseValueStatus) GetConnector() ListConnectorsResponseValueStatusConnector`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ListConnectorsResponseValueStatus) GetConnectorOk() (*ListConnectorsResponseValueStatusConnector, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ListConnectorsResponseValueStatus) SetConnector(v ListConnectorsResponseValueStatusConnector)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *ListConnectorsResponseValueStatus) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetTasks

`func (o *ListConnectorsResponseValueStatus) GetTasks() []ListConnectorsResponseValueStatusTasksInner`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ListConnectorsResponseValueStatus) GetTasksOk() (*[]ListConnectorsResponseValueStatusTasksInner, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ListConnectorsResponseValueStatus) SetTasks(v []ListConnectorsResponseValueStatusTasksInner)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ListConnectorsResponseValueStatus) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


