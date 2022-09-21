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

// ListConnectorsResponseValueStatusConnector Connector status.
type ListConnectorsResponseValueStatusConnector struct {
	// ID of the worker.
	Id *string `json:"id,omitempty"`
	State *State `json:"state,omitempty"`
}

// NewListConnectorsResponseValueStatusConnector instantiates a new ListConnectorsResponseValueStatusConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectorsResponseValueStatusConnector() *ListConnectorsResponseValueStatusConnector {
	this := ListConnectorsResponseValueStatusConnector{}
	return &this
}

// NewListConnectorsResponseValueStatusConnectorWithDefaults instantiates a new ListConnectorsResponseValueStatusConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectorsResponseValueStatusConnectorWithDefaults() *ListConnectorsResponseValueStatusConnector {
	this := ListConnectorsResponseValueStatusConnector{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListConnectorsResponseValueStatusConnector) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorsResponseValueStatusConnector) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListConnectorsResponseValueStatusConnector) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListConnectorsResponseValueStatusConnector) SetId(v string) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ListConnectorsResponseValueStatusConnector) GetState() State {
	if o == nil || o.State == nil {
		var ret State
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorsResponseValueStatusConnector) GetStateOk() (*State, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ListConnectorsResponseValueStatusConnector) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given State and assigns it to the State field.
func (o *ListConnectorsResponseValueStatusConnector) SetState(v State) {
	o.State = &v
}

func (o ListConnectorsResponseValueStatusConnector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableListConnectorsResponseValueStatusConnector struct {
	value *ListConnectorsResponseValueStatusConnector
	isSet bool
}

func (v NullableListConnectorsResponseValueStatusConnector) Get() *ListConnectorsResponseValueStatusConnector {
	return v.value
}

func (v *NullableListConnectorsResponseValueStatusConnector) Set(val *ListConnectorsResponseValueStatusConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectorsResponseValueStatusConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectorsResponseValueStatusConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectorsResponseValueStatusConnector(val *ListConnectorsResponseValueStatusConnector) *NullableListConnectorsResponseValueStatusConnector {
	return &NullableListConnectorsResponseValueStatusConnector{value: val, isSet: true}
}

func (v NullableListConnectorsResponseValueStatusConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectorsResponseValueStatusConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


