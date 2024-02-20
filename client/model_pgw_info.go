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
	"fmt"
)

// checks if the PgwInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PgwInfo{}

// PgwInfo struct for PgwInfo
type PgwInfo struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn"`
	// Fully Qualified Domain Name
	PgwFqdn string `json:"pgwFqdn"`
	PgwIpAddr NullableIpAddress `json:"pgwIpAddr,omitempty"`
	PlmnId *PlmnId `json:"plmnId,omitempty"`
	EpdgInd *bool `json:"epdgInd,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	PcfId *string `json:"pcfId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RegistrationTime *time.Time `json:"registrationTime,omitempty"`
}

type _PgwInfo PgwInfo

// NewPgwInfo instantiates a new PgwInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPgwInfo(dnn string, pgwFqdn string) *PgwInfo {
	this := PgwInfo{}
	this.Dnn = dnn
	this.PgwFqdn = pgwFqdn
	var epdgInd bool = false
	this.EpdgInd = &epdgInd
	return &this
}

// NewPgwInfoWithDefaults instantiates a new PgwInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPgwInfoWithDefaults() *PgwInfo {
	this := PgwInfo{}
	var epdgInd bool = false
	this.EpdgInd = &epdgInd
	return &this
}

// GetDnn returns the Dnn field value
func (o *PgwInfo) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *PgwInfo) GetDnnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *PgwInfo) SetDnn(v string) {
	o.Dnn = v
}

// GetPgwFqdn returns the PgwFqdn field value
func (o *PgwInfo) GetPgwFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PgwFqdn
}

// GetPgwFqdnOk returns a tuple with the PgwFqdn field value
// and a boolean to check if the value has been set.
func (o *PgwInfo) GetPgwFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PgwFqdn, true
}

// SetPgwFqdn sets field value
func (o *PgwInfo) SetPgwFqdn(v string) {
	o.PgwFqdn = v
}

// GetPgwIpAddr returns the PgwIpAddr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PgwInfo) GetPgwIpAddr() IpAddress {
	if o == nil || IsNil(o.PgwIpAddr.Get()) {
		var ret IpAddress
		return ret
	}
	return *o.PgwIpAddr.Get()
}

// GetPgwIpAddrOk returns a tuple with the PgwIpAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PgwInfo) GetPgwIpAddrOk() (*IpAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.PgwIpAddr.Get(), o.PgwIpAddr.IsSet()
}

// HasPgwIpAddr returns a boolean if a field has been set.
func (o *PgwInfo) HasPgwIpAddr() bool {
	if o != nil && o.PgwIpAddr.IsSet() {
		return true
	}

	return false
}

// SetPgwIpAddr gets a reference to the given NullableIpAddress and assigns it to the PgwIpAddr field.
func (o *PgwInfo) SetPgwIpAddr(v IpAddress) {
	o.PgwIpAddr.Set(&v)
}
// SetPgwIpAddrNil sets the value for PgwIpAddr to be an explicit nil
func (o *PgwInfo) SetPgwIpAddrNil() {
	o.PgwIpAddr.Set(nil)
}

// UnsetPgwIpAddr ensures that no value is present for PgwIpAddr, not even an explicit nil
func (o *PgwInfo) UnsetPgwIpAddr() {
	o.PgwIpAddr.Unset()
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *PgwInfo) GetPlmnId() PlmnId {
	if o == nil || IsNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgwInfo) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.PlmnId) {
		return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *PgwInfo) HasPlmnId() bool {
	if o != nil && !IsNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *PgwInfo) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

// GetEpdgInd returns the EpdgInd field value if set, zero value otherwise.
func (o *PgwInfo) GetEpdgInd() bool {
	if o == nil || IsNil(o.EpdgInd) {
		var ret bool
		return ret
	}
	return *o.EpdgInd
}

// GetEpdgIndOk returns a tuple with the EpdgInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgwInfo) GetEpdgIndOk() (*bool, bool) {
	if o == nil || IsNil(o.EpdgInd) {
		return nil, false
	}
	return o.EpdgInd, true
}

// HasEpdgInd returns a boolean if a field has been set.
func (o *PgwInfo) HasEpdgInd() bool {
	if o != nil && !IsNil(o.EpdgInd) {
		return true
	}

	return false
}

// SetEpdgInd gets a reference to the given bool and assigns it to the EpdgInd field.
func (o *PgwInfo) SetEpdgInd(v bool) {
	o.EpdgInd = &v
}

// GetPcfId returns the PcfId field value if set, zero value otherwise.
func (o *PgwInfo) GetPcfId() string {
	if o == nil || IsNil(o.PcfId) {
		var ret string
		return ret
	}
	return *o.PcfId
}

// GetPcfIdOk returns a tuple with the PcfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgwInfo) GetPcfIdOk() (*string, bool) {
	if o == nil || IsNil(o.PcfId) {
		return nil, false
	}
	return o.PcfId, true
}

// HasPcfId returns a boolean if a field has been set.
func (o *PgwInfo) HasPcfId() bool {
	if o != nil && !IsNil(o.PcfId) {
		return true
	}

	return false
}

// SetPcfId gets a reference to the given string and assigns it to the PcfId field.
func (o *PgwInfo) SetPcfId(v string) {
	o.PcfId = &v
}

// GetRegistrationTime returns the RegistrationTime field value if set, zero value otherwise.
func (o *PgwInfo) GetRegistrationTime() time.Time {
	if o == nil || IsNil(o.RegistrationTime) {
		var ret time.Time
		return ret
	}
	return *o.RegistrationTime
}

// GetRegistrationTimeOk returns a tuple with the RegistrationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgwInfo) GetRegistrationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RegistrationTime) {
		return nil, false
	}
	return o.RegistrationTime, true
}

// HasRegistrationTime returns a boolean if a field has been set.
func (o *PgwInfo) HasRegistrationTime() bool {
	if o != nil && !IsNil(o.RegistrationTime) {
		return true
	}

	return false
}

// SetRegistrationTime gets a reference to the given time.Time and assigns it to the RegistrationTime field.
func (o *PgwInfo) SetRegistrationTime(v time.Time) {
	o.RegistrationTime = &v
}

func (o PgwInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PgwInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dnn"] = o.Dnn
	toSerialize["pgwFqdn"] = o.PgwFqdn
	if o.PgwIpAddr.IsSet() {
		toSerialize["pgwIpAddr"] = o.PgwIpAddr.Get()
	}
	if !IsNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !IsNil(o.EpdgInd) {
		toSerialize["epdgInd"] = o.EpdgInd
	}
	if !IsNil(o.PcfId) {
		toSerialize["pcfId"] = o.PcfId
	}
	if !IsNil(o.RegistrationTime) {
		toSerialize["registrationTime"] = o.RegistrationTime
	}
	return toSerialize, nil
}

func (o *PgwInfo) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dnn",
		"pgwFqdn",
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

	varPgwInfo := _PgwInfo{}

	err = json.Unmarshal(bytes, &varPgwInfo)

	if err != nil {
		return err
	}

	*o = PgwInfo(varPgwInfo)

	return err
}

type NullablePgwInfo struct {
	value *PgwInfo
	isSet bool
}

func (v NullablePgwInfo) Get() *PgwInfo {
	return v.value
}

func (v *NullablePgwInfo) Set(val *PgwInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePgwInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePgwInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePgwInfo(val *PgwInfo) *NullablePgwInfo {
	return &NullablePgwInfo{value: val, isSet: true}
}

func (v NullablePgwInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePgwInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


