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

// checks if the ProseServiceAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProseServiceAuth{}

// ProseServiceAuth Indicates whether the UE is authorized to use ProSe related services. 
type ProseServiceAuth struct {
	ProseDirectDiscoveryAuth *UeAuth `json:"proseDirectDiscoveryAuth,omitempty"`
	ProseDirectCommunicationAuth *UeAuth `json:"proseDirectCommunicationAuth,omitempty"`
	ProseL2RelayAuth *UeAuth `json:"proseL2RelayAuth,omitempty"`
	ProseL3RelayAuth *UeAuth `json:"proseL3RelayAuth,omitempty"`
	ProseL2RemoteAuth *UeAuth `json:"proseL2RemoteAuth,omitempty"`
	ProseL3RemoteAuth *UeAuth `json:"proseL3RemoteAuth,omitempty"`
}

// NewProseServiceAuth instantiates a new ProseServiceAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProseServiceAuth() *ProseServiceAuth {
	this := ProseServiceAuth{}
	return &this
}

// NewProseServiceAuthWithDefaults instantiates a new ProseServiceAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProseServiceAuthWithDefaults() *ProseServiceAuth {
	this := ProseServiceAuth{}
	return &this
}

// GetProseDirectDiscoveryAuth returns the ProseDirectDiscoveryAuth field value if set, zero value otherwise.
func (o *ProseServiceAuth) GetProseDirectDiscoveryAuth() UeAuth {
	if o == nil || IsNil(o.ProseDirectDiscoveryAuth) {
		var ret UeAuth
		return ret
	}
	return *o.ProseDirectDiscoveryAuth
}

// GetProseDirectDiscoveryAuthOk returns a tuple with the ProseDirectDiscoveryAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProseServiceAuth) GetProseDirectDiscoveryAuthOk() (*UeAuth, bool) {
	if o == nil || IsNil(o.ProseDirectDiscoveryAuth) {
		return nil, false
	}
	return o.ProseDirectDiscoveryAuth, true
}

// HasProseDirectDiscoveryAuth returns a boolean if a field has been set.
func (o *ProseServiceAuth) HasProseDirectDiscoveryAuth() bool {
	if o != nil && !IsNil(o.ProseDirectDiscoveryAuth) {
		return true
	}

	return false
}

// SetProseDirectDiscoveryAuth gets a reference to the given UeAuth and assigns it to the ProseDirectDiscoveryAuth field.
func (o *ProseServiceAuth) SetProseDirectDiscoveryAuth(v UeAuth) {
	o.ProseDirectDiscoveryAuth = &v
}

// GetProseDirectCommunicationAuth returns the ProseDirectCommunicationAuth field value if set, zero value otherwise.
func (o *ProseServiceAuth) GetProseDirectCommunicationAuth() UeAuth {
	if o == nil || IsNil(o.ProseDirectCommunicationAuth) {
		var ret UeAuth
		return ret
	}
	return *o.ProseDirectCommunicationAuth
}

// GetProseDirectCommunicationAuthOk returns a tuple with the ProseDirectCommunicationAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProseServiceAuth) GetProseDirectCommunicationAuthOk() (*UeAuth, bool) {
	if o == nil || IsNil(o.ProseDirectCommunicationAuth) {
		return nil, false
	}
	return o.ProseDirectCommunicationAuth, true
}

// HasProseDirectCommunicationAuth returns a boolean if a field has been set.
func (o *ProseServiceAuth) HasProseDirectCommunicationAuth() bool {
	if o != nil && !IsNil(o.ProseDirectCommunicationAuth) {
		return true
	}

	return false
}

// SetProseDirectCommunicationAuth gets a reference to the given UeAuth and assigns it to the ProseDirectCommunicationAuth field.
func (o *ProseServiceAuth) SetProseDirectCommunicationAuth(v UeAuth) {
	o.ProseDirectCommunicationAuth = &v
}

// GetProseL2RelayAuth returns the ProseL2RelayAuth field value if set, zero value otherwise.
func (o *ProseServiceAuth) GetProseL2RelayAuth() UeAuth {
	if o == nil || IsNil(o.ProseL2RelayAuth) {
		var ret UeAuth
		return ret
	}
	return *o.ProseL2RelayAuth
}

// GetProseL2RelayAuthOk returns a tuple with the ProseL2RelayAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProseServiceAuth) GetProseL2RelayAuthOk() (*UeAuth, bool) {
	if o == nil || IsNil(o.ProseL2RelayAuth) {
		return nil, false
	}
	return o.ProseL2RelayAuth, true
}

// HasProseL2RelayAuth returns a boolean if a field has been set.
func (o *ProseServiceAuth) HasProseL2RelayAuth() bool {
	if o != nil && !IsNil(o.ProseL2RelayAuth) {
		return true
	}

	return false
}

// SetProseL2RelayAuth gets a reference to the given UeAuth and assigns it to the ProseL2RelayAuth field.
func (o *ProseServiceAuth) SetProseL2RelayAuth(v UeAuth) {
	o.ProseL2RelayAuth = &v
}

// GetProseL3RelayAuth returns the ProseL3RelayAuth field value if set, zero value otherwise.
func (o *ProseServiceAuth) GetProseL3RelayAuth() UeAuth {
	if o == nil || IsNil(o.ProseL3RelayAuth) {
		var ret UeAuth
		return ret
	}
	return *o.ProseL3RelayAuth
}

// GetProseL3RelayAuthOk returns a tuple with the ProseL3RelayAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProseServiceAuth) GetProseL3RelayAuthOk() (*UeAuth, bool) {
	if o == nil || IsNil(o.ProseL3RelayAuth) {
		return nil, false
	}
	return o.ProseL3RelayAuth, true
}

// HasProseL3RelayAuth returns a boolean if a field has been set.
func (o *ProseServiceAuth) HasProseL3RelayAuth() bool {
	if o != nil && !IsNil(o.ProseL3RelayAuth) {
		return true
	}

	return false
}

// SetProseL3RelayAuth gets a reference to the given UeAuth and assigns it to the ProseL3RelayAuth field.
func (o *ProseServiceAuth) SetProseL3RelayAuth(v UeAuth) {
	o.ProseL3RelayAuth = &v
}

// GetProseL2RemoteAuth returns the ProseL2RemoteAuth field value if set, zero value otherwise.
func (o *ProseServiceAuth) GetProseL2RemoteAuth() UeAuth {
	if o == nil || IsNil(o.ProseL2RemoteAuth) {
		var ret UeAuth
		return ret
	}
	return *o.ProseL2RemoteAuth
}

// GetProseL2RemoteAuthOk returns a tuple with the ProseL2RemoteAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProseServiceAuth) GetProseL2RemoteAuthOk() (*UeAuth, bool) {
	if o == nil || IsNil(o.ProseL2RemoteAuth) {
		return nil, false
	}
	return o.ProseL2RemoteAuth, true
}

// HasProseL2RemoteAuth returns a boolean if a field has been set.
func (o *ProseServiceAuth) HasProseL2RemoteAuth() bool {
	if o != nil && !IsNil(o.ProseL2RemoteAuth) {
		return true
	}

	return false
}

// SetProseL2RemoteAuth gets a reference to the given UeAuth and assigns it to the ProseL2RemoteAuth field.
func (o *ProseServiceAuth) SetProseL2RemoteAuth(v UeAuth) {
	o.ProseL2RemoteAuth = &v
}

// GetProseL3RemoteAuth returns the ProseL3RemoteAuth field value if set, zero value otherwise.
func (o *ProseServiceAuth) GetProseL3RemoteAuth() UeAuth {
	if o == nil || IsNil(o.ProseL3RemoteAuth) {
		var ret UeAuth
		return ret
	}
	return *o.ProseL3RemoteAuth
}

// GetProseL3RemoteAuthOk returns a tuple with the ProseL3RemoteAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProseServiceAuth) GetProseL3RemoteAuthOk() (*UeAuth, bool) {
	if o == nil || IsNil(o.ProseL3RemoteAuth) {
		return nil, false
	}
	return o.ProseL3RemoteAuth, true
}

// HasProseL3RemoteAuth returns a boolean if a field has been set.
func (o *ProseServiceAuth) HasProseL3RemoteAuth() bool {
	if o != nil && !IsNil(o.ProseL3RemoteAuth) {
		return true
	}

	return false
}

// SetProseL3RemoteAuth gets a reference to the given UeAuth and assigns it to the ProseL3RemoteAuth field.
func (o *ProseServiceAuth) SetProseL3RemoteAuth(v UeAuth) {
	o.ProseL3RemoteAuth = &v
}

func (o ProseServiceAuth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProseServiceAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProseDirectDiscoveryAuth) {
		toSerialize["proseDirectDiscoveryAuth"] = o.ProseDirectDiscoveryAuth
	}
	if !IsNil(o.ProseDirectCommunicationAuth) {
		toSerialize["proseDirectCommunicationAuth"] = o.ProseDirectCommunicationAuth
	}
	if !IsNil(o.ProseL2RelayAuth) {
		toSerialize["proseL2RelayAuth"] = o.ProseL2RelayAuth
	}
	if !IsNil(o.ProseL3RelayAuth) {
		toSerialize["proseL3RelayAuth"] = o.ProseL3RelayAuth
	}
	if !IsNil(o.ProseL2RemoteAuth) {
		toSerialize["proseL2RemoteAuth"] = o.ProseL2RemoteAuth
	}
	if !IsNil(o.ProseL3RemoteAuth) {
		toSerialize["proseL3RemoteAuth"] = o.ProseL3RemoteAuth
	}
	return toSerialize, nil
}

type NullableProseServiceAuth struct {
	value *ProseServiceAuth
	isSet bool
}

func (v NullableProseServiceAuth) Get() *ProseServiceAuth {
	return v.value
}

func (v *NullableProseServiceAuth) Set(val *ProseServiceAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableProseServiceAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableProseServiceAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProseServiceAuth(val *ProseServiceAuth) *NullableProseServiceAuth {
	return &NullableProseServiceAuth{value: val, isSet: true}
}

func (v NullableProseServiceAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProseServiceAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


