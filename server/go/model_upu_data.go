/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




// UpuData - Contains UE parameters update data set (e.g., the updated Routing ID Data or the Default configured NSSAI).
type UpuData struct {

	// Contains a secure packet.
	SecPacket string `json:"secPacket,omitempty"`

	DefaultConfNssai []Snssai `json:"defaultConfNssai,omitempty"`

	// Represents a routing indicator.
	RoutingId string `json:"routingId,omitempty"`
}

// AssertUpuDataRequired checks if the required fields are not zero-ed
func AssertUpuDataRequired(obj UpuData) error {
	for _, el := range obj.DefaultConfNssai {
		if err := AssertSnssaiRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertUpuDataConstraints checks if the values respects the defined constraints
func AssertUpuDataConstraints(obj UpuData) error {
	return nil
}
