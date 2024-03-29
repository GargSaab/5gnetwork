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

// checks if the ModificationNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModificationNotification{}

// ModificationNotification struct for ModificationNotification
type ModificationNotification struct {
	NotifyItems []NotifyItem `json:"notifyItems"`
}

type _ModificationNotification ModificationNotification

// NewModificationNotification instantiates a new ModificationNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModificationNotification(notifyItems []NotifyItem) *ModificationNotification {
	this := ModificationNotification{}
	this.NotifyItems = notifyItems
	return &this
}

// NewModificationNotificationWithDefaults instantiates a new ModificationNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModificationNotificationWithDefaults() *ModificationNotification {
	this := ModificationNotification{}
	return &this
}

// GetNotifyItems returns the NotifyItems field value
func (o *ModificationNotification) GetNotifyItems() []NotifyItem {
	if o == nil {
		var ret []NotifyItem
		return ret
	}

	return o.NotifyItems
}

// GetNotifyItemsOk returns a tuple with the NotifyItems field value
// and a boolean to check if the value has been set.
func (o *ModificationNotification) GetNotifyItemsOk() ([]NotifyItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotifyItems, true
}

// SetNotifyItems sets field value
func (o *ModificationNotification) SetNotifyItems(v []NotifyItem) {
	o.NotifyItems = v
}

func (o ModificationNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModificationNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifyItems"] = o.NotifyItems
	return toSerialize, nil
}

func (o *ModificationNotification) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"notifyItems",
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

	varModificationNotification := _ModificationNotification{}

	err = json.Unmarshal(bytes, &varModificationNotification)

	if err != nil {
		return err
	}

	*o = ModificationNotification(varModificationNotification)

	return err
}

type NullableModificationNotification struct {
	value *ModificationNotification
	isSet bool
}

func (v NullableModificationNotification) Get() *ModificationNotification {
	return v.value
}

func (v *NullableModificationNotification) Set(val *ModificationNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableModificationNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableModificationNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModificationNotification(val *ModificationNotification) *NullableModificationNotification {
	return &NullableModificationNotification{value: val, isSet: true}
}

func (v NullableModificationNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModificationNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


