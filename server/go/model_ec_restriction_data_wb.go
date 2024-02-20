/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type EcRestrictionDataWb struct {

	EcModeARestricted bool `json:"ecModeARestricted,omitempty"`

	EcModeBRestricted bool `json:"ecModeBRestricted,omitempty"`
}

// AssertEcRestrictionDataWbRequired checks if the required fields are not zero-ed
func AssertEcRestrictionDataWbRequired(obj EcRestrictionDataWb) error {
	return nil
}

// AssertEcRestrictionDataWbConstraints checks if the values respects the defined constraints
func AssertEcRestrictionDataWbConstraints(obj EcRestrictionDataWb) error {
	return nil
}
