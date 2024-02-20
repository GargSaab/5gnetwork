/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




// MbsSubscriptionData - Contains the 5MBS Subscription Data.
type MbsSubscriptionData struct {

	MbsAllowed bool `json:"mbsAllowed,omitempty"`

	MbsSessionIdList []MbsSessionId `json:"mbsSessionIdList,omitempty"`
}

// AssertMbsSubscriptionDataRequired checks if the required fields are not zero-ed
func AssertMbsSubscriptionDataRequired(obj MbsSubscriptionData) error {
	for _, el := range obj.MbsSessionIdList {
		if err := AssertMbsSessionIdRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertMbsSubscriptionDataConstraints checks if the values respects the defined constraints
func AssertMbsSubscriptionDataConstraints(obj MbsSubscriptionData) error {
	return nil
}
