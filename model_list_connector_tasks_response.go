/*
Kafka Connect REST API

Kafka Connect REST API https://docs.confluent.io/platform/current/connect/references/restapi.html

API version: 0.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ListConnectorTasksResponse Connector tasks fetched.
type ListConnectorTasksResponse struct {
	// List of tasks for the connector.
	Tasks []ListConnectorTasksResponseTasksInner `json:"tasks,omitempty"`
}

// NewListConnectorTasksResponse instantiates a new ListConnectorTasksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectorTasksResponse() *ListConnectorTasksResponse {
	this := ListConnectorTasksResponse{}
	return &this
}

// NewListConnectorTasksResponseWithDefaults instantiates a new ListConnectorTasksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectorTasksResponseWithDefaults() *ListConnectorTasksResponse {
	this := ListConnectorTasksResponse{}
	return &this
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *ListConnectorTasksResponse) GetTasks() []ListConnectorTasksResponseTasksInner {
	if o == nil || o.Tasks == nil {
		var ret []ListConnectorTasksResponseTasksInner
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorTasksResponse) GetTasksOk() ([]ListConnectorTasksResponseTasksInner, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *ListConnectorTasksResponse) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []ListConnectorTasksResponseTasksInner and assigns it to the Tasks field.
func (o *ListConnectorTasksResponse) SetTasks(v []ListConnectorTasksResponseTasksInner) {
	o.Tasks = v
}

func (o ListConnectorTasksResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableListConnectorTasksResponse struct {
	value *ListConnectorTasksResponse
	isSet bool
}

func (v NullableListConnectorTasksResponse) Get() *ListConnectorTasksResponse {
	return v.value
}

func (v *NullableListConnectorTasksResponse) Set(val *ListConnectorTasksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectorTasksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectorTasksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectorTasksResponse(val *ListConnectorTasksResponse) *NullableListConnectorTasksResponse {
	return &NullableListConnectorTasksResponse{value: val, isSet: true}
}

func (v NullableListConnectorTasksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectorTasksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


