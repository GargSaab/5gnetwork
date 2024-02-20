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

// ProseSubscriptionDataRetrievalAPIController binds http requests to an api service and writes the service results to the http response
type ProseSubscriptionDataRetrievalAPIController struct {
	service      ProseSubscriptionDataRetrievalAPIServicer
	errorHandler ErrorHandler
}

// ProseSubscriptionDataRetrievalAPIOption for how the controller is set up.
type ProseSubscriptionDataRetrievalAPIOption func(*ProseSubscriptionDataRetrievalAPIController)

// WithProseSubscriptionDataRetrievalAPIErrorHandler inject ErrorHandler into controller
func WithProseSubscriptionDataRetrievalAPIErrorHandler(h ErrorHandler) ProseSubscriptionDataRetrievalAPIOption {
	return func(c *ProseSubscriptionDataRetrievalAPIController) {
		c.errorHandler = h
	}
}

// NewProseSubscriptionDataRetrievalAPIController creates a default api controller
func NewProseSubscriptionDataRetrievalAPIController(s ProseSubscriptionDataRetrievalAPIServicer, opts ...ProseSubscriptionDataRetrievalAPIOption) Router {
	controller := &ProseSubscriptionDataRetrievalAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ProseSubscriptionDataRetrievalAPIController
func (c *ProseSubscriptionDataRetrievalAPIController) Routes() Routes {
	return Routes{
		"GetProseData": Route{
			strings.ToUpper("Get"),
			"/nudm-sdm/v2/{supi}/prose-data",
			c.GetProseData,
		},
	}
}

// GetProseData - retrieve a UE's ProSe Subscription Data
func (c *ProseSubscriptionDataRetrievalAPIController) GetProseData(w http.ResponseWriter, r *http.Request) {
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
	ifNoneMatchParam := r.Header.Get("If-None-Match")
	ifModifiedSinceParam := r.Header.Get("If-Modified-Since")
	result, err := c.service.GetProseData(r.Context(), supiParam, supportedFeaturesParam, ifNoneMatchParam, ifModifiedSinceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
