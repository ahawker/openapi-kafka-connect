# GetConnectorStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the connector. | [optional] 
**Connector** | Pointer to [**GetConnectorStatusResponseConnector**](GetConnectorStatusResponseConnector.md) |  | [optional] 
**Tasks** | Pointer to [**[]GetConnectorStatusResponseTasksInner**](GetConnectorStatusResponseTasksInner.md) | List of task status for the connector. | [optional] 

## Methods

### NewGetConnectorStatusResponse

`func NewGetConnectorStatusResponse() *GetConnectorStatusResponse`

NewGetConnectorStatusResponse instantiates a new GetConnectorStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectorStatusResponseWithDefaults

`func NewGetConnectorStatusResponseWithDefaults() *GetConnectorStatusResponse`

NewGetConnectorStatusResponseWithDefaults instantiates a new GetConnectorStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetConnectorStatusResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetConnectorStatusResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetConnectorStatusResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetConnectorStatusResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConnector

`func (o *GetConnectorStatusResponse) GetConnector() GetConnectorStatusResponseConnector`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *GetConnectorStatusResponse) GetConnectorOk() (*GetConnectorStatusResponseConnector, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *GetConnectorStatusResponse) SetConnector(v GetConnectorStatusResponseConnector)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *GetConnectorStatusResponse) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetTasks

`func (o *GetConnectorStatusResponse) GetTasks() []GetConnectorStatusResponseTasksInner`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *GetConnectorStatusResponse) GetTasksOk() (*[]GetConnectorStatusResponseTasksInner, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *GetConnectorStatusResponse) SetTasks(v []GetConnectorStatusResponseTasksInner)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *GetConnectorStatusResponse) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


