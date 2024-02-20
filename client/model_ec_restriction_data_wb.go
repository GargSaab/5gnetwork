/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the EcRestrictionDataWb type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EcRestrictionDataWb{}

// EcRestrictionDataWb struct for EcRestrictionDataWb
type EcRestrictionDataWb struct {
	EcModeARestricted *bool `json:"ecModeARestricted,omitempty"`
	EcModeBRestricted *bool `json:"ecModeBRestricted,omitempty"`
}

// NewEcRestrictionDataWb instantiates a new EcRestrictionDataWb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEcRestrictionDataWb() *EcRestrictionDataWb {
	this := EcRestrictionDataWb{}
	return &this
}

// NewEcRestrictionDataWbWithDefaults instantiates a new EcRestrictionDataWb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEcRestrictionDataWbWithDefaults() *EcRestrictionDataWb {
	this := EcRestrictionDataWb{}
	return &this
}

// GetEcModeARestricted returns the EcModeARestricted field value if set, zero value otherwise.
func (o *EcRestrictionDataWb) GetEcModeARestricted() bool {
	if o == nil || IsNil(o.EcModeARestricted) {
		var ret bool
		return ret
	}
	return *o.EcModeARestricted
}

// GetEcModeARestrictedOk returns a tuple with the EcModeARestricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcRestrictionDataWb) GetEcModeARestrictedOk() (*bool, bool) {
	if o == nil || IsNil(o.EcModeARestricted) {
		return nil, false
	}
	return o.EcModeARestricted, true
}

// HasEcModeARestricted returns a boolean if a field has been set.
func (o *EcRestrictionDataWb) HasEcModeARestricted() bool {
	if o != nil && !IsNil(o.EcModeARestricted) {
		return true
	}

	return false
}

// SetEcModeARestricted gets a reference to the given bool and assigns it to the EcModeARestricted field.
func (o *EcRestrictionDataWb) SetEcModeARestricted(v bool) {
	o.EcModeARestricted = &v
}

// GetEcModeBRestricted returns the EcModeBRestricted field value if set, zero value otherwise.
func (o *EcRestrictionDataWb) GetEcModeBRestricted() bool {
	if o == nil || IsNil(o.EcModeBRestricted) {
		var ret bool
		return ret
	}
	return *o.EcModeBRestricted
}

// GetEcModeBRestrictedOk returns a tuple with the EcModeBRestricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcRestrictionDataWb) GetEcModeBRestrictedOk() (*bool, bool) {
	if o == nil || IsNil(o.EcModeBRestricted) {
		return nil, false
	}
	return o.EcModeBRestricted, true
}

// HasEcModeBRestricted returns a boolean if a field has been set.
func (o *EcRestrictionDataWb) HasEcModeBRestricted() bool {
	if o != nil && !IsNil(o.EcModeBRestricted) {
		return true
	}

	return false
}

// SetEcModeBRestricted gets a reference to the given bool and assigns it to the EcModeBRestricted field.
func (o *EcRestrictionDataWb) SetEcModeBRestricted(v bool) {
	o.EcModeBRestricted = &v
}

func (o EcRestrictionDataWb) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EcRestrictionDataWb) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EcModeARestricted) {
		toSerialize["ecModeARestricted"] = o.EcModeARestricted
	}
	if !IsNil(o.EcModeBRestricted) {
		toSerialize["ecModeBRestricted"] = o.EcModeBRestricted
	}
	return toSerialize, nil
}

type NullableEcRestrictionDataWb struct {
	value *EcRestrictionDataWb
	isSet bool
}

func (v NullableEcRestrictionDataWb) Get() *EcRestrictionDataWb {
	return v.value
}

func (v *NullableEcRestrictionDataWb) Set(val *EcRestrictionDataWb) {
	v.value = val
	v.isSet = true
}

func (v NullableEcRestrictionDataWb) IsSet() bool {
	return v.isSet
}

func (v *NullableEcRestrictionDataWb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcRestrictionDataWb(val *EcRestrictionDataWb) *NullableEcRestrictionDataWb {
	return &NullableEcRestrictionDataWb{value: val, isSet: true}
}

func (v NullableEcRestrictionDataWb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcRestrictionDataWb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


