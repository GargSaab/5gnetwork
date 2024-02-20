/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ValidTimePeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidTimePeriod{}

// ValidTimePeriod struct for ValidTimePeriod
type ValidTimePeriod struct {
	// string with format 'date-time' as defined in OpenAPI.
	StartTime *time.Time `json:"startTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	EndTime *time.Time `json:"endTime,omitempty"`
}

// NewValidTimePeriod instantiates a new ValidTimePeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidTimePeriod() *ValidTimePeriod {
	this := ValidTimePeriod{}
	return &this
}

// NewValidTimePeriodWithDefaults instantiates a new ValidTimePeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidTimePeriodWithDefaults() *ValidTimePeriod {
	this := ValidTimePeriod{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ValidTimePeriod) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidTimePeriod) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ValidTimePeriod) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ValidTimePeriod) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ValidTimePeriod) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidTimePeriod) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ValidTimePeriod) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *ValidTimePeriod) SetEndTime(v time.Time) {
	o.EndTime = &v
}

func (o ValidTimePeriod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidTimePeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	return toSerialize, nil
}

type NullableValidTimePeriod struct {
	value *ValidTimePeriod
	isSet bool
}

func (v NullableValidTimePeriod) Get() *ValidTimePeriod {
	return v.value
}

func (v *NullableValidTimePeriod) Set(val *ValidTimePeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableValidTimePeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableValidTimePeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidTimePeriod(val *ValidTimePeriod) *NullableValidTimePeriod {
	return &NullableValidTimePeriod{value: val, isSet: true}
}

func (v NullableValidTimePeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidTimePeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


