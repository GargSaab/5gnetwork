/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type SubscriptionDataSets struct {

	AmData AccessAndMobilitySubscriptionData `json:"amData,omitempty"`

	SmfSelData SmfSelectionSubscriptionData `json:"smfSelData,omitempty"`

	UecAmfData UeContextInAmfData `json:"uecAmfData,omitempty"`

	UecSmfData UeContextInSmfData `json:"uecSmfData,omitempty"`

	UecSmsfData UeContextInSmsfData `json:"uecSmsfData,omitempty"`

	SmsSubsData SmsSubscriptionData `json:"smsSubsData,omitempty"`

	SmData SmSubsData `json:"smData,omitempty"`

	TraceData *TraceData `json:"traceData,omitempty"`

	SmsMngData SmsManagementSubscriptionData `json:"smsMngData,omitempty"`

	LcsPrivacyData LcsPrivacyData `json:"lcsPrivacyData,omitempty"`

	LcsMoData LcsMoData `json:"lcsMoData,omitempty"`

	V2xData V2xSubscriptionData `json:"v2xData,omitempty"`

	LcsBroadcastAssistanceTypesData LcsBroadcastAssistanceTypesData `json:"lcsBroadcastAssistanceTypesData,omitempty"`

	ProseData ProseSubscriptionData `json:"proseData,omitempty"`

	MbsData MbsSubscriptionData `json:"mbsData,omitempty"`

	UcData UcSubscriptionData `json:"ucData,omitempty"`
}

// AssertSubscriptionDataSetsRequired checks if the required fields are not zero-ed
func AssertSubscriptionDataSetsRequired(obj SubscriptionDataSets) error {
	if err := AssertAccessAndMobilitySubscriptionDataRequired(obj.AmData); err != nil {
		return err
	}
	if err := AssertSmfSelectionSubscriptionDataRequired(obj.SmfSelData); err != nil {
		return err
	}
	if err := AssertUeContextInAmfDataRequired(obj.UecAmfData); err != nil {
		return err
	}
	if err := AssertUeContextInSmfDataRequired(obj.UecSmfData); err != nil {
		return err
	}
	if err := AssertUeContextInSmsfDataRequired(obj.UecSmsfData); err != nil {
		return err
	}
	if err := AssertSmsSubscriptionDataRequired(obj.SmsSubsData); err != nil {
		return err
	}
	if err := AssertSmSubsDataRequired(obj.SmData); err != nil {
		return err
	}
	if obj.TraceData != nil {
		if err := AssertTraceDataRequired(*obj.TraceData); err != nil {
			return err
		}
	}
	if err := AssertSmsManagementSubscriptionDataRequired(obj.SmsMngData); err != nil {
		return err
	}
	if err := AssertLcsPrivacyDataRequired(obj.LcsPrivacyData); err != nil {
		return err
	}
	if err := AssertLcsMoDataRequired(obj.LcsMoData); err != nil {
		return err
	}
	if err := AssertV2xSubscriptionDataRequired(obj.V2xData); err != nil {
		return err
	}
	if err := AssertLcsBroadcastAssistanceTypesDataRequired(obj.LcsBroadcastAssistanceTypesData); err != nil {
		return err
	}
	if err := AssertProseSubscriptionDataRequired(obj.ProseData); err != nil {
		return err
	}
	if err := AssertMbsSubscriptionDataRequired(obj.MbsData); err != nil {
		return err
	}
	if err := AssertUcSubscriptionDataRequired(obj.UcData); err != nil {
		return err
	}
	return nil
}

// AssertSubscriptionDataSetsConstraints checks if the values respects the defined constraints
func AssertSubscriptionDataSetsConstraints(obj SubscriptionDataSets) error {
	return nil
}
