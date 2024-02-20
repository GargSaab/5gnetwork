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
)

// SubscriptionDeletionAPIService is a service that implements the logic for the SubscriptionDeletionAPIServicer
// This service should implement the business logic for every endpoint for the SubscriptionDeletionAPI API.
// Include any external packages or services that will be required by this service.
type SubscriptionDeletionAPIService struct {
}

// NewSubscriptionDeletionAPIService creates a default api service
func NewSubscriptionDeletionAPIService() SubscriptionDeletionAPIServicer {
	return &SubscriptionDeletionAPIService{}
}

// Unsubscribe - unsubscribe from notifications
func (s *SubscriptionDeletionAPIService) Unsubscribe(ctx context.Context, ueId string, subscriptionId string) (ImplResponse, error) {
	// TODO - update Unsubscribe with the required logic for this service method.
	// Add api_subscription_deletion_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	if ueId == "" {
		return Response(http.StatusBadRequest, ProblemDetails{
			Title:  "Bad Request",
			Detail: "UE ID is required",
		}), nil
	}

	// Check if the subscription ID is valid
	if subscriptionId == "" {
		return Response(http.StatusBadRequest, ProblemDetails{
			Title:  "Bad Request",
			Detail: "Subscription ID is required",
		}), nil
	}

	// TODO: Implement the actual logic for subscription deletion based on the provided requirements

	// Uncomment the next line to return response Response(http.StatusNoContent, nil), nil

	// Return a simple response indicating successful deletion
	//TODO, delete it from the store or database
	return Response(http.StatusNoContent, "Subscription deleted successfully"), nil

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

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

	// return Response(http.StatusNotImplemented, nil), errors.New("Unsubscribe method not implemented")
}
