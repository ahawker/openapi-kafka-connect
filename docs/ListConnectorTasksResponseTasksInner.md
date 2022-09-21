# ListConnectorTasksResponseTasksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**ListConnectorTasksResponseTasksInnerId**](ListConnectorTasksResponseTasksInnerId.md) |  | [optional] 
**Config** | Pointer to **map[string]string** | Configuration of a connector/task/worker. All keys/values should be strings. | [optional] 

## Methods

### NewListConnectorTasksResponseTasksInner

`func NewListConnectorTasksResponseTasksInner() *ListConnectorTasksResponseTasksInner`

NewListConnectorTasksResponseTasksInner instantiates a new ListConnectorTasksResponseTasksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectorTasksResponseTasksInnerWithDefaults

`func NewListConnectorTasksResponseTasksInnerWithDefaults() *ListConnectorTasksResponseTasksInner`

NewListConnectorTasksResponseTasksInnerWithDefaults instantiates a new ListConnectorTasksResponseTasksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListConnectorTasksResponseTasksInner) GetId() ListConnectorTasksResponseTasksInnerId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListConnectorTasksResponseTasksInner) GetIdOk() (*ListConnectorTasksResponseTasksInnerId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListConnectorTasksResponseTasksInner) SetId(v ListConnectorTasksResponseTasksInnerId)`

SetId sets Id field to given value.

### HasId

`func (o *ListConnectorTasksResponseTasksInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfig

`func (o *ListConnectorTasksResponseTasksInner) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListConnectorTasksResponseTasksInner) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListConnectorTasksResponseTasksInner) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListConnectorTasksResponseTasksInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


