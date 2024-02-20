/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type PduSession struct {

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SmfInstanceId string `json:"smfInstanceId"`

	PlmnId PlmnId `json:"plmnId"`

	SingleNssai Snssai `json:"singleNssai,omitempty"`
}

// AssertPduSessionRequired checks if the required fields are not zero-ed
func AssertPduSessionRequired(obj PduSession) error {
	elements := map[string]interface{}{
		"dnn": obj.Dnn,
		"smfInstanceId": obj.SmfInstanceId,
		"plmnId": obj.PlmnId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPlmnIdRequired(obj.PlmnId); err != nil {
		return err
	}
	if err := AssertSnssaiRequired(obj.SingleNssai); err != nil {
		return err
	}
	return nil
}

// AssertPduSessionConstraints checks if the values respects the defined constraints
func AssertPduSessionConstraints(obj PduSession) error {
	return nil
}
