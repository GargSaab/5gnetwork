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

	"github.com/gorilla/mux"
)

// SliceSelectionSubscriptionDataRetrievalAPIController binds http requests to an api service and writes the service results to the http response
type SliceSelectionSubscriptionDataRetrievalAPIController struct {
	service      SliceSelectionSubscriptionDataRetrievalAPIServicer
	errorHandler ErrorHandler
}

// SliceSelectionSubscriaptionDataRetrievalAPIOption for how the controller is set up.
type SliceSelectionSubscriptionDataRetrievalAPIOption func(*SliceSelectionSubscriptionDataRetrievalAPIController)

// WithSliceSelectionSubscriptionDataRetrievalAPIErrorHandler inject ErrorHandler into controller
func WithSliceSelectionSubscriptionDataRetrievalAPIErrorHandler(h ErrorHandler) SliceSelectionSubscriptionDataRetrievalAPIOption {
	return func(c *SliceSelectionSubscriptionDataRetrievalAPIController) {
		c.errorHandler = h
	}
}

// NewSliceSelectionSubscriptionDataRetrievalAPIController creates a default api controller
func NewSliceSelectionSubscriptionDataRetrievalAPIController(s SliceSelectionSubscriptionDataRetrievalAPIServicer, opts ...SliceSelectionSubscriptionDataRetrievalAPIOption) Router {
	controller := &SliceSelectionSubscriptionDataRetrievalAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SliceSelectionSubscriptionDataRetrievalAPIController
func (c *SliceSelectionSubscriptionDataRetrievalAPIController) Routes() Routes {
	return Routes{
		"GetNSSAI": Route{
			strings.ToUpper("Get"),
			"/nudm-sdm/v2/{supi}/nssai",
			c.GetNSSAI,
		},
	}
}

// GetNSSAI - retrieve a UE's subscribed NSSAI
func (c *SliceSelectionSubscriptionDataRetrievalAPIController) GetNSSAI(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	supiParam := params["supi"]
	if supiParam == "" {
		c.errorHandler(w, r, &RequiredError{"supi"}, nil)
		return
	}
	var supportedFeaturesParam string
	if query.Has("supported-features") {
		supportedFeaturesParam = query.Get("supported-features")
	}
	var plmnIdParam PlmnId
	if query.Has("plmn-id") {
		plmnIdParam = PlmnId{}
		//Logically incorrect
		//query.Get("plmn-id")
	}
	disasterRoamingIndParam, err := parseBoolParameter(
		query.Get("disaster-roaming-ind"),
		WithDefaultOrParse[bool](false, parseBool),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	ifNoneMatchParam := r.Header.Get("If-None-Match")
	ifModifiedSinceParam := r.Header.Get("If-Modified-Since")
	result, err := c.service.GetNSSAI(r.Context(), supiParam, supportedFeaturesParam, plmnIdParam, disasterRoamingIndParam, ifNoneMatchParam, ifModifiedSinceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
