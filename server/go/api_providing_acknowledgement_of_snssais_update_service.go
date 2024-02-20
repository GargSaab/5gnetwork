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
	"context"
	"net/http"
	"errors"
)

// ProvidingAcknowledgementOfSNSSAIsUpdateAPIService is a service that implements the logic for the ProvidingAcknowledgementOfSNSSAIsUpdateAPIServicer
// This service should implement the business logic for every endpoint for the ProvidingAcknowledgementOfSNSSAIsUpdateAPI API.
// Include any external packages or services that will be required by this service.
type ProvidingAcknowledgementOfSNSSAIsUpdateAPIService struct {
}

// NewProvidingAcknowledgementOfSNSSAIsUpdateAPIService creates a default api service
func NewProvidingAcknowledgementOfSNSSAIsUpdateAPIService() ProvidingAcknowledgementOfSNSSAIsUpdateAPIServicer {
	return &ProvidingAcknowledgementOfSNSSAIsUpdateAPIService{}
}

// SNSSAIsAck - Nudm_Sdm Info operation for S-NSSAIs acknowledgement
func (s *ProvidingAcknowledgementOfSNSSAIsUpdateAPIService) SNSSAIsAck(ctx context.Context, supi string, acknowledgeInfo AcknowledgeInfo) (ImplResponse, error) {
	// TODO - update SNSSAIsAck with the required logic for this service method.
	// Add api_providing_acknowledgement_of_snssais_update_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(400, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(500, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(500, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(503, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(503, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("SNSSAIsAck method not implemented")
}