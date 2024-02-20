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
	"errors"
)



// MdtConfiguration - contains contain MDT configuration data.
type MdtConfiguration struct {

	JobType JobType `json:"jobType"`

	ReportType ReportTypeMdt `json:"reportType,omitempty"`

	AreaScope AreaScope `json:"areaScope,omitempty"`

	MeasurementLteList []MeasurementLteForMdt `json:"measurementLteList,omitempty"`

	MeasurementNrList []MeasurementNrForMdt `json:"measurementNrList,omitempty"`

	SensorMeasurementList []SensorMeasurement `json:"sensorMeasurementList,omitempty"`

	ReportingTriggerList []ReportingTrigger `json:"reportingTriggerList,omitempty"`

	ReportInterval ReportIntervalMdt `json:"reportInterval,omitempty"`

	ReportIntervalNr ReportIntervalNrMdt `json:"reportIntervalNr,omitempty"`

	ReportAmount ReportAmountMdt `json:"reportAmount,omitempty"`

	// This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in LTE. When present, this IE shall indicate the Event Threshold for RSRP, and the value shall be between 0-97. 
	EventThresholdRsrp int32 `json:"eventThresholdRsrp,omitempty"`

	// This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in NR. When present, this IE shall indicate the Event Threshold for RSRP, and the value shall be between 0-127. 
	EventThresholdRsrpNr int32 `json:"eventThresholdRsrpNr,omitempty"`

	// This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in LTE.When present, this IE shall indicate the Event Threshold for RSRQ, and the value shall be between 0-34. 
	EventThresholdRsrq int32 `json:"eventThresholdRsrq,omitempty"`

	// This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in NR.When present, this IE shall indicate the Event Threshold for RSRQ, and the value shall be between 0-127. 
	EventThresholdRsrqNr int32 `json:"eventThresholdRsrqNr,omitempty"`

	EventList []EventForMdt `json:"eventList,omitempty"`

	LoggingInterval LoggingIntervalMdt `json:"loggingInterval,omitempty"`

	LoggingIntervalNr LoggingIntervalNrMdt `json:"loggingIntervalNr,omitempty"`

	LoggingDuration LoggingDurationMdt `json:"loggingDuration,omitempty"`

	LoggingDurationNr LoggingDurationNrMdt `json:"loggingDurationNr,omitempty"`

	PositioningMethod PositioningMethodMdt `json:"positioningMethod,omitempty"`

	AddPositioningMethodList []PositioningMethodMdt `json:"addPositioningMethodList,omitempty"`

	CollectionPeriodRmmLte CollectionPeriodRmmLteMdt `json:"collectionPeriodRmmLte,omitempty"`

	CollectionPeriodRmmNr CollectionPeriodRmmNrMdt `json:"collectionPeriodRmmNr,omitempty"`

	MeasurementPeriodLte MeasurementPeriodLteMdt `json:"measurementPeriodLte,omitempty"`

	MdtAllowedPlmnIdList []PlmnId `json:"mdtAllowedPlmnIdList,omitempty"`

	MbsfnAreaList []MbsfnArea `json:"mbsfnAreaList,omitempty"`

	InterFreqTargetList []InterFreqTargetInfo `json:"interFreqTargetList,omitempty"`
}

// AssertMdtConfigurationRequired checks if the required fields are not zero-ed
func AssertMdtConfigurationRequired(obj MdtConfiguration) error {
	elements := map[string]interface{}{
		"jobType": obj.JobType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertJobTypeRequired(obj.JobType); err != nil {
		return err
	}
	if err := AssertReportTypeMdtRequired(obj.ReportType); err != nil {
		return err
	}
	if err := AssertAreaScopeRequired(obj.AreaScope); err != nil {
		return err
	}
	for _, el := range obj.MeasurementLteList {
		if err := AssertMeasurementLteForMdtRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.MeasurementNrList {
		if err := AssertMeasurementNrForMdtRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.SensorMeasurementList {
		if err := AssertSensorMeasurementRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ReportingTriggerList {
		if err := AssertReportingTriggerRequired(el); err != nil {
			return err
		}
	}
	if err := AssertReportIntervalMdtRequired(obj.ReportInterval); err != nil {
		return err
	}
	if err := AssertReportIntervalNrMdtRequired(obj.ReportIntervalNr); err != nil {
		return err
	}
	if err := AssertReportAmountMdtRequired(obj.ReportAmount); err != nil {
		return err
	}
	for _, el := range obj.EventList {
		if err := AssertEventForMdtRequired(el); err != nil {
			return err
		}
	}
	if err := AssertLoggingIntervalMdtRequired(obj.LoggingInterval); err != nil {
		return err
	}
	if err := AssertLoggingIntervalNrMdtRequired(obj.LoggingIntervalNr); err != nil {
		return err
	}
	if err := AssertLoggingDurationMdtRequired(obj.LoggingDuration); err != nil {
		return err
	}
	if err := AssertLoggingDurationNrMdtRequired(obj.LoggingDurationNr); err != nil {
		return err
	}
	if err := AssertPositioningMethodMdtRequired(obj.PositioningMethod); err != nil {
		return err
	}
	for _, el := range obj.AddPositioningMethodList {
		if err := AssertPositioningMethodMdtRequired(el); err != nil {
			return err
		}
	}
	if err := AssertCollectionPeriodRmmLteMdtRequired(obj.CollectionPeriodRmmLte); err != nil {
		return err
	}
	if err := AssertCollectionPeriodRmmNrMdtRequired(obj.CollectionPeriodRmmNr); err != nil {
		return err
	}
	if err := AssertMeasurementPeriodLteMdtRequired(obj.MeasurementPeriodLte); err != nil {
		return err
	}
	for _, el := range obj.MdtAllowedPlmnIdList {
		if err := AssertPlmnIdRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.MbsfnAreaList {
		if err := AssertMbsfnAreaRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.InterFreqTargetList {
		if err := AssertInterFreqTargetInfoRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertMdtConfigurationConstraints checks if the values respects the defined constraints
func AssertMdtConfigurationConstraints(obj MdtConfiguration) error {
	if obj.EventThresholdRsrp < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.EventThresholdRsrp > 97 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.EventThresholdRsrpNr < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.EventThresholdRsrpNr > 127 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.EventThresholdRsrq < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.EventThresholdRsrq > 34 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.EventThresholdRsrqNr < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.EventThresholdRsrqNr > 127 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
