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
	"errors"
	"net/http"
)

// GroupIdentifiersAPIService is a service that implements the logic for the GroupIdentifiersAPIServicer
// This service should implement the business logic for every endpoint for the GroupIdentifiersAPI API.
// Include any external packages or services that will be required by this service.
type GroupIdentifiersAPIService struct {
}

// NewGroupIdentifiersAPIService creates a default api service
func NewGroupIdentifiersAPIService() GroupIdentifiersAPIServicer {
	return &GroupIdentifiersAPIService{}
}

// GetGroupIdentifiers - Mapping of Group Identifiers
func (s *GroupIdentifiersAPIService) GetGroupIdentifiers(ctx context.Context, extGroupId string, intGroupId string, ueIdInd bool, supportedFeatures string, afId string, ifNoneMatch string, ifModifiedSince string) (ImplResponse, error) {
	// TODO - update GetGroupIdentifiers with the required logic for this service method.
	// Add api_group_identifiers_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GroupIdentifiers{}) or use other options such as http.Ok ...
	// return Response(200, GroupIdentifiers{}), nil

	// TODO: Uncomment the next line to return response Response(400, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(400, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(403, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(403, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(404, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(404, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(500, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(500, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(503, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(503, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetGroupIdentifiers method not implemented")
}
