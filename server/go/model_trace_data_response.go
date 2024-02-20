/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type TraceDataResponse struct {

	TraceData *TraceData `json:"traceData,omitempty"`

	SharedTraceDataId string `json:"sharedTraceDataId,omitempty"`
}

// AssertTraceDataResponseRequired checks if the required fields are not zero-ed
func AssertTraceDataResponseRequired(obj TraceDataResponse) error {
	if obj.TraceData != nil {
		if err := AssertTraceDataRequired(*obj.TraceData); err != nil {
			return err
		}
	}
	return nil
}

// AssertTraceDataResponseConstraints checks if the values respects the defined constraints
func AssertTraceDataResponseConstraints(obj TraceDataResponse) error {
	return nil
}
