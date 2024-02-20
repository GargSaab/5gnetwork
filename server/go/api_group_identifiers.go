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
	"net/http"
	"strings"
)

// GroupIdentifiersAPIController binds http requests to an api service and writes the service results to the http response
type GroupIdentifiersAPIController struct {
	service      GroupIdentifiersAPIServicer
	errorHandler ErrorHandler
}

// GroupIdentifiersAPIOption for how the controller is set up.
type GroupIdentifiersAPIOption func(*GroupIdentifiersAPIController)

// WithGroupIdentifiersAPIErrorHandler inject ErrorHandler into controller
func WithGroupIdentifiersAPIErrorHandler(h ErrorHandler) GroupIdentifiersAPIOption {
	return func(c *GroupIdentifiersAPIController) {
		c.errorHandler = h
	}
}

// NewGroupIdentifiersAPIController creates a default api controller
func NewGroupIdentifiersAPIController(s GroupIdentifiersAPIServicer, opts ...GroupIdentifiersAPIOption) Router {
	controller := &GroupIdentifiersAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the GroupIdentifiersAPIController
func (c *GroupIdentifiersAPIController) Routes() Routes {
	return Routes{
		"GetGroupIdentifiers": Route{
			strings.ToUpper("Get"),
			"/nudm-sdm/v2/group-data/group-identifiers",
			c.GetGroupIdentifiers,
		},
	}
}

// GetGroupIdentifiers - Mapping of Group Identifiers
func (c *GroupIdentifiersAPIController) GetGroupIdentifiers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var extGroupIdParam string
	if query.Has("ext-group-id") {
		extGroupIdParam = query.Get("ext-group-id")
	}
	var intGroupIdParam string
	if query.Has("int-group-id") {
		intGroupIdParam = query.Get("int-group-id")
	}
	ueIdIndParam, err := parseBoolParameter(
		query.Get("ue-id-ind"),
		WithDefaultOrParse[bool](false, parseBool),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var supportedFeaturesParam string
	if query.Has("supported-features") {
		supportedFeaturesParam = query.Get("supported-features")
	}
	var afIdParam string
	if query.Has("af-id") {
		afIdParam = query.Get("af-id")
	}
	ifNoneMatchParam := r.Header.Get("If-None-Match")
	ifModifiedSinceParam := r.Header.Get("If-Modified-Since")
	result, err := c.service.GetGroupIdentifiers(r.Context(), extGroupIdParam, intGroupIdParam, ueIdIndParam, supportedFeaturesParam, afIdParam, ifNoneMatchParam, ifModifiedSinceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
