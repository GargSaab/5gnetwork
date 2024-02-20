/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type AppDescriptor struct {

	// Represents the Operating System of the served UE.
	OsId string `json:"osId,omitempty"`

	AppId string `json:"appId,omitempty"`
}

// AssertAppDescriptorRequired checks if the required fields are not zero-ed
func AssertAppDescriptorRequired(obj AppDescriptor) error {
	return nil
}

// AssertAppDescriptorConstraints checks if the values respects the defined constraints
func AssertAppDescriptorConstraints(obj AppDescriptor) error {
	return nil
}
