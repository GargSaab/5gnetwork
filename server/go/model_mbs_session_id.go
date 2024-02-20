/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




// MbsSessionId - MBS Session Identifier
type MbsSessionId struct {

	Tmgi Tmgi `json:"tmgi,omitempty"`

	Ssm Ssm `json:"ssm,omitempty"`

	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).  
	Nid string `json:"nid,omitempty"`
}

// AssertMbsSessionIdRequired checks if the required fields are not zero-ed
func AssertMbsSessionIdRequired(obj MbsSessionId) error {
	if err := AssertTmgiRequired(obj.Tmgi); err != nil {
		return err
	}
	if err := AssertSsmRequired(obj.Ssm); err != nil {
		return err
	}
	return nil
}

// AssertMbsSessionIdConstraints checks if the values respects the defined constraints
func AssertMbsSessionIdConstraints(obj MbsSessionId) error {
	return nil
}
