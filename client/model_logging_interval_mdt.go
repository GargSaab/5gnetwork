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

// LoggingIntervalMdt The enumeration LoggingIntervalMdt defines Logging Interval for MDT in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.12-1. 
type LoggingIntervalMdt struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LoggingIntervalMdt) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(LoggingIntervalMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LoggingIntervalMdt) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLoggingIntervalMdt struct {
	value *LoggingIntervalMdt
	isSet bool
}

func (v NullableLoggingIntervalMdt) Get() *LoggingIntervalMdt {
	return v.value
}

func (v *NullableLoggingIntervalMdt) Set(val *LoggingIntervalMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableLoggingIntervalMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableLoggingIntervalMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoggingIntervalMdt(val *LoggingIntervalMdt) *NullableLoggingIntervalMdt {
	return &NullableLoggingIntervalMdt{value: val, isSet: true}
}

func (v NullableLoggingIntervalMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoggingIntervalMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


