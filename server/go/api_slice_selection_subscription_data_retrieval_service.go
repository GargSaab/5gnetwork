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

// SliceSelectionSubscriptionDataRetrievalAPIService is a service that implements the logic for the SliceSelectionSubscriptionDataRetrievalAPIServicer
// This service should implement the business logic for every endpoint for the SliceSelectionSubscriptionDataRetrievalAPI API.
// Include any external packages or services that will be required by this service.
type SliceSelectionSubscriptionDataRetrievalAPIService struct {
}

// NewSliceSelectionSubscriptionDataRetrievalAPIService creates a default api service
func NewSliceSelectionSubscriptionDataRetrievalAPIService() SliceSelectionSubscriptionDataRetrievalAPIServicer {
	return &SliceSelectionSubscriptionDataRetrievalAPIService{}
}

// GetNSSAI - retrieve a UE&#39;s subscribed NSSAI
func (s *SliceSelectionSubscriptionDataRetrievalAPIService) GetNSSAI(ctx context.Context, supi string, supportedFeatures string, plmnId PlmnId, disasterRoamingInd bool, ifNoneMatch string, ifModifiedSince string) (ImplResponse, error) {
	// TODO - update GetNSSAI with the required logic for this service method.
	// Add api_slice_selection_subscription_data_retrieval_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Nssai{}) or use other options such as http.Ok ...
	// return Response(200, Nssai{}), nil

	// TODO: Uncomment the next line to return response Response(400, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(400, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(404, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(404, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(500, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(500, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(503, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(503, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetNSSAI method not implemented")
}
