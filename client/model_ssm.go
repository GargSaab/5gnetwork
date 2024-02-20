/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the Ssm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ssm{}

// Ssm Source specific IP multicast address
type Ssm struct {
	SourceIpAddr NullableIpAddr `json:"sourceIpAddr"`
	DestIpAddr NullableIpAddr `json:"destIpAddr"`
}

type _Ssm Ssm

// NewSsm instantiates a new Ssm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsm(sourceIpAddr NullableIpAddr, destIpAddr NullableIpAddr) *Ssm {
	this := Ssm{}
	this.SourceIpAddr = sourceIpAddr
	this.DestIpAddr = destIpAddr
	return &this
}

// NewSsmWithDefaults instantiates a new Ssm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsmWithDefaults() *Ssm {
	this := Ssm{}
	return &this
}

// GetSourceIpAddr returns the SourceIpAddr field value
// If the value is explicit nil, the zero value for IpAddr will be returned
func (o *Ssm) GetSourceIpAddr() IpAddr {
	if o == nil || o.SourceIpAddr.Get() == nil {
		var ret IpAddr
		return ret
	}

	return *o.SourceIpAddr.Get()
}

// GetSourceIpAddrOk returns a tuple with the SourceIpAddr field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ssm) GetSourceIpAddrOk() (*IpAddr, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceIpAddr.Get(), o.SourceIpAddr.IsSet()
}

// SetSourceIpAddr sets field value
func (o *Ssm) SetSourceIpAddr(v IpAddr) {
	o.SourceIpAddr.Set(&v)
}

// GetDestIpAddr returns the DestIpAddr field value
// If the value is explicit nil, the zero value for IpAddr will be returned
func (o *Ssm) GetDestIpAddr() IpAddr {
	if o == nil || o.DestIpAddr.Get() == nil {
		var ret IpAddr
		return ret
	}

	return *o.DestIpAddr.Get()
}

// GetDestIpAddrOk returns a tuple with the DestIpAddr field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ssm) GetDestIpAddrOk() (*IpAddr, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestIpAddr.Get(), o.DestIpAddr.IsSet()
}

// SetDestIpAddr sets field value
func (o *Ssm) SetDestIpAddr(v IpAddr) {
	o.DestIpAddr.Set(&v)
}

func (o Ssm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ssm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceIpAddr"] = o.SourceIpAddr.Get()
	toSerialize["destIpAddr"] = o.DestIpAddr.Get()
	return toSerialize, nil
}

func (o *Ssm) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sourceIpAddr",
		"destIpAddr",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSsm := _Ssm{}

	err = json.Unmarshal(bytes, &varSsm)

	if err != nil {
		return err
	}

	*o = Ssm(varSsm)

	return err
}

type NullableSsm struct {
	value *Ssm
	isSet bool
}

func (v NullableSsm) Get() *Ssm {
	return v.value
}

func (v *NullableSsm) Set(val *Ssm) {
	v.value = val
	v.isSet = true
}

func (v NullableSsm) IsSet() bool {
	return v.isSet
}

func (v *NullableSsm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsm(val *Ssm) *NullableSsm {
	return &NullableSsm{value: val, isSet: true}
}

func (v NullableSsm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


