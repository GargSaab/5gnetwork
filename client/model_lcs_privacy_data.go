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

// checks if the LcsPrivacyData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LcsPrivacyData{}

// LcsPrivacyData struct for LcsPrivacyData
type LcsPrivacyData struct {
	Lpi *Lpi `json:"lpi,omitempty"`
	UnrelatedClass *UnrelatedClass `json:"unrelatedClass,omitempty"`
	PlmnOperatorClasses []PlmnOperatorClass `json:"plmnOperatorClasses,omitempty"`
}

// NewLcsPrivacyData instantiates a new LcsPrivacyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLcsPrivacyData() *LcsPrivacyData {
	this := LcsPrivacyData{}
	return &this
}

// NewLcsPrivacyDataWithDefaults instantiates a new LcsPrivacyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLcsPrivacyDataWithDefaults() *LcsPrivacyData {
	this := LcsPrivacyData{}
	return &this
}

// GetLpi returns the Lpi field value if set, zero value otherwise.
func (o *LcsPrivacyData) GetLpi() Lpi {
	if o == nil || IsNil(o.Lpi) {
		var ret Lpi
		return ret
	}
	return *o.Lpi
}

// GetLpiOk returns a tuple with the Lpi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LcsPrivacyData) GetLpiOk() (*Lpi, bool) {
	if o == nil || IsNil(o.Lpi) {
		return nil, false
	}
	return o.Lpi, true
}

// HasLpi returns a boolean if a field has been set.
func (o *LcsPrivacyData) HasLpi() bool {
	if o != nil && !IsNil(o.Lpi) {
		return true
	}

	return false
}

// SetLpi gets a reference to the given Lpi and assigns it to the Lpi field.
func (o *LcsPrivacyData) SetLpi(v Lpi) {
	o.Lpi = &v
}

// GetUnrelatedClass returns the UnrelatedClass field value if set, zero value otherwise.
func (o *LcsPrivacyData) GetUnrelatedClass() UnrelatedClass {
	if o == nil || IsNil(o.UnrelatedClass) {
		var ret UnrelatedClass
		return ret
	}
	return *o.UnrelatedClass
}

// GetUnrelatedClassOk returns a tuple with the UnrelatedClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LcsPrivacyData) GetUnrelatedClassOk() (*UnrelatedClass, bool) {
	if o == nil || IsNil(o.UnrelatedClass) {
		return nil, false
	}
	return o.UnrelatedClass, true
}

// HasUnrelatedClass returns a boolean if a field has been set.
func (o *LcsPrivacyData) HasUnrelatedClass() bool {
	if o != nil && !IsNil(o.UnrelatedClass) {
		return true
	}

	return false
}

// SetUnrelatedClass gets a reference to the given UnrelatedClass and assigns it to the UnrelatedClass field.
func (o *LcsPrivacyData) SetUnrelatedClass(v UnrelatedClass) {
	o.UnrelatedClass = &v
}

// GetPlmnOperatorClasses returns the PlmnOperatorClasses field value if set, zero value otherwise.
func (o *LcsPrivacyData) GetPlmnOperatorClasses() []PlmnOperatorClass {
	if o == nil || IsNil(o.PlmnOperatorClasses) {
		var ret []PlmnOperatorClass
		return ret
	}
	return o.PlmnOperatorClasses
}

// GetPlmnOperatorClassesOk returns a tuple with the PlmnOperatorClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LcsPrivacyData) GetPlmnOperatorClassesOk() ([]PlmnOperatorClass, bool) {
	if o == nil || IsNil(o.PlmnOperatorClasses) {
		return nil, false
	}
	return o.PlmnOperatorClasses, true
}

// HasPlmnOperatorClasses returns a boolean if a field has been set.
func (o *LcsPrivacyData) HasPlmnOperatorClasses() bool {
	if o != nil && !IsNil(o.PlmnOperatorClasses) {
		return true
	}

	return false
}

// SetPlmnOperatorClasses gets a reference to the given []PlmnOperatorClass and assigns it to the PlmnOperatorClasses field.
func (o *LcsPrivacyData) SetPlmnOperatorClasses(v []PlmnOperatorClass) {
	o.PlmnOperatorClasses = v
}

func (o LcsPrivacyData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LcsPrivacyData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lpi) {
		toSerialize["lpi"] = o.Lpi
	}
	if !IsNil(o.UnrelatedClass) {
		toSerialize["unrelatedClass"] = o.UnrelatedClass
	}
	if !IsNil(o.PlmnOperatorClasses) {
		toSerialize["plmnOperatorClasses"] = o.PlmnOperatorClasses
	}
	return toSerialize, nil
}

type NullableLcsPrivacyData struct {
	value *LcsPrivacyData
	isSet bool
}

func (v NullableLcsPrivacyData) Get() *LcsPrivacyData {
	return v.value
}

func (v *NullableLcsPrivacyData) Set(val *LcsPrivacyData) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsPrivacyData) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsPrivacyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsPrivacyData(val *LcsPrivacyData) *NullableLcsPrivacyData {
	return &NullableLcsPrivacyData{value: val, isSet: true}
}

func (v NullableLcsPrivacyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsPrivacyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


