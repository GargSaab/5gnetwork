/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi


import (
	"time"
)



type ValidTimePeriod struct {

	// string with format 'date-time' as defined in OpenAPI.
	StartTime time.Time `json:"startTime,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	EndTime time.Time `json:"endTime,omitempty"`
}

// AssertValidTimePeriodRequired checks if the required fields are not zero-ed
func AssertValidTimePeriodRequired(obj ValidTimePeriod) error {
	return nil
}

// AssertValidTimePeriodConstraints checks if the values respects the defined constraints
func AssertValidTimePeriodConstraints(obj ValidTimePeriod) error {
	return nil
}
