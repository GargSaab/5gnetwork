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



// PointAltitudeUncertainty - Ellipsoid point with altitude and uncertainty ellipsoid.
type PointAltitudeUncertainty struct {
	GadShape

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`

	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`

	// Indicates value of uncertainty.
	UncertaintyAltitude float32 `json:"uncertaintyAltitude"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// AssertPointAltitudeUncertaintyRequired checks if the required fields are not zero-ed
func AssertPointAltitudeUncertaintyRequired(obj PointAltitudeUncertainty) error {
	elements := map[string]interface{}{
		"point": obj.Point,
		"altitude": obj.Altitude,
		"uncertaintyEllipse": obj.UncertaintyEllipse,
		"uncertaintyAltitude": obj.UncertaintyAltitude,
		"confidence": obj.Confidence,
		"shape": obj.Shape,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertGadShapeRequired(obj.GadShape); err != nil {
		return err
	}

	if err := AssertGeographicalCoordinatesRequired(obj.Point); err != nil {
		return err
	}
	if err := AssertUncertaintyEllipseRequired(obj.UncertaintyEllipse); err != nil {
		return err
	}
	return nil
}

// AssertPointAltitudeUncertaintyConstraints checks if the values respects the defined constraints
func AssertPointAltitudeUncertaintyConstraints(obj PointAltitudeUncertainty) error {
	if obj.Altitude < -32767 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Altitude > 32767 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.UncertaintyAltitude < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Confidence < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Confidence > 100 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
