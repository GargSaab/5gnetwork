/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// RetrievalOfSharedDataAPIService RetrievalOfSharedDataAPI service
type RetrievalOfSharedDataAPIService service

type ApiGetSharedDataRequest struct {
	ctx context.Context
	ApiService *RetrievalOfSharedDataAPIService
	sharedDataIds *[]string
	supportedFeatures *string
	supportedFeatures2 *string
	ifNoneMatch *string
	ifModifiedSince *string
}

// List of shared data ids
func (r ApiGetSharedDataRequest) SharedDataIds(sharedDataIds []string) ApiGetSharedDataRequest {
	r.sharedDataIds = &sharedDataIds
	return r
}

// Supported Features; this query parameter should not be used
// Deprecated
func (r ApiGetSharedDataRequest) SupportedFeatures(supportedFeatures string) ApiGetSharedDataRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

// Supported Features
func (r ApiGetSharedDataRequest) SupportedFeatures2(supportedFeatures2 string) ApiGetSharedDataRequest {
	r.supportedFeatures2 = &supportedFeatures2
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.2
func (r ApiGetSharedDataRequest) IfNoneMatch(ifNoneMatch string) ApiGetSharedDataRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.3
func (r ApiGetSharedDataRequest) IfModifiedSince(ifModifiedSince string) ApiGetSharedDataRequest {
	r.ifModifiedSince = &ifModifiedSince
	return r
}

func (r ApiGetSharedDataRequest) Execute() ([]SharedData, *http.Response, error) {
	return r.ApiService.GetSharedDataExecute(r)
}

/*
GetSharedData retrieve shared data

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSharedDataRequest
*/
func (a *RetrievalOfSharedDataAPIService) GetSharedData(ctx context.Context) ApiGetSharedDataRequest {
	return ApiGetSharedDataRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SharedData
func (a *RetrievalOfSharedDataAPIService) GetSharedDataExecute(r ApiGetSharedDataRequest) ([]SharedData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SharedData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RetrievalOfSharedDataAPIService.GetSharedData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shared-data"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sharedDataIds == nil {
		return localVarReturnValue, nil, reportError("sharedDataIds is required and must be specified")
	}
	if len(*r.sharedDataIds) < 1 {
		return localVarReturnValue, nil, reportError("sharedDataIds must have at least 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "shared-data-ids", r.sharedDataIds, "csv")
	if r.supportedFeatures != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "supportedFeatures", r.supportedFeatures, "")
	}
	if r.supportedFeatures2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "supported-features", r.supportedFeatures2, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ifNoneMatch != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-None-Match", r.ifNoneMatch, "")
	}
	if r.ifModifiedSince != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Modified-Since", r.ifModifiedSince, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
