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

// checks if the PduSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PduSession{}

// PduSession struct for PduSession
type PduSession struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SmfInstanceId string `json:"smfInstanceId"`
	PlmnId PlmnId `json:"plmnId"`
	SingleNssai *Snssai `json:"singleNssai,omitempty"`
}

type _PduSession PduSession

// NewPduSession instantiates a new PduSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduSession(dnn string, smfInstanceId string, plmnId PlmnId) *PduSession {
	this := PduSession{}
	this.Dnn = dnn
	this.SmfInstanceId = smfInstanceId
	this.PlmnId = plmnId
	return &this
}

// NewPduSessionWithDefaults instantiates a new PduSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduSessionWithDefaults() *PduSession {
	this := PduSession{}
	return &this
}

// GetDnn returns the Dnn field value
func (o *PduSession) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *PduSession) GetDnnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *PduSession) SetDnn(v string) {
	o.Dnn = v
}

// GetSmfInstanceId returns the SmfInstanceId field value
func (o *PduSession) GetSmfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmfInstanceId
}

// GetSmfInstanceIdOk returns a tuple with the SmfInstanceId field value
// and a boolean to check if the value has been set.
func (o *PduSession) GetSmfInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmfInstanceId, true
}

// SetSmfInstanceId sets field value
func (o *PduSession) SetSmfInstanceId(v string) {
	o.SmfInstanceId = v
}

// GetPlmnId returns the PlmnId field value
func (o *PduSession) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *PduSession) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *PduSession) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetSingleNssai returns the SingleNssai field value if set, zero value otherwise.
func (o *PduSession) GetSingleNssai() Snssai {
	if o == nil || IsNil(o.SingleNssai) {
		var ret Snssai
		return ret
	}
	return *o.SingleNssai
}

// GetSingleNssaiOk returns a tuple with the SingleNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSession) GetSingleNssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.SingleNssai) {
		return nil, false
	}
	return o.SingleNssai, true
}

// HasSingleNssai returns a boolean if a field has been set.
func (o *PduSession) HasSingleNssai() bool {
	if o != nil && !IsNil(o.SingleNssai) {
		return true
	}

	return false
}

// SetSingleNssai gets a reference to the given Snssai and assigns it to the SingleNssai field.
func (o *PduSession) SetSingleNssai(v Snssai) {
	o.SingleNssai = &v
}

func (o PduSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PduSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dnn"] = o.Dnn
	toSerialize["smfInstanceId"] = o.SmfInstanceId
	toSerialize["plmnId"] = o.PlmnId
	if !IsNil(o.SingleNssai) {
		toSerialize["singleNssai"] = o.SingleNssai
	}
	return toSerialize, nil
}

func (o *PduSession) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dnn",
		"smfInstanceId",
		"plmnId",
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

	varPduSession := _PduSession{}

	err = json.Unmarshal(bytes, &varPduSession)

	if err != nil {
		return err
	}

	*o = PduSession(varPduSession)

	return err
}

type NullablePduSession struct {
	value *PduSession
	isSet bool
}

func (v NullablePduSession) Get() *PduSession {
	return v.value
}

func (v *NullablePduSession) Set(val *PduSession) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSession) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSession(val *PduSession) *NullablePduSession {
	return &NullablePduSession{value: val, isSet: true}
}

func (v NullablePduSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


