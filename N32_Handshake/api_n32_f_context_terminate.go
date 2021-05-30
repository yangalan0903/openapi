/*
 * N32 Handshake API
 *
 * N32-c Handshake Service.  © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).  All rights reserved. 
 *
 * Source file: 3GPP TS 29.573 V16.5.0; 5G System; Public Land Mobile Network (PLMN) Interconnection; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.573/
 *
 * API version: 1.1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package N32_Handshake

import (
    "github.com/yangalan0903/openapi"
	"github.com/yangalan0903/openapi/models"

	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type N32FContextTerminateApiService service

/*
N32FContextTerminateApiService N32-f Context Terminate
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param n32fContextInfo Custom operation for n32-f context termination
@return models.N32fContextInfo
*/

func (a *N32FContextTerminateApiService) PostN32fTerminate(ctx context.Context, n32fContextInfo models.N32fContextInfo) ( models.N32fContextInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  models.N32fContextInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/n32f-terminate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


    localVarHTTPContentTypes := []string{ "application/json" }


    localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{ "application/json", "application/problem+json" }

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = &n32fContextInfo

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

    apiError := openapi.GenericOpenAPIError{
        RawBody:     localVarBody,
        ErrorStatus: localVarHTTPResponse.Status,
    }

    switch localVarHTTPResponse.StatusCode {
        case 200:
            err = openapi.Deserialize(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
            }
            return localVarReturnValue, localVarHTTPResponse, nil
        case 400:
            var v models.ProblemDetails
            err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        case 411:
            var v models.ProblemDetails
            err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        case 413:
            var v models.ProblemDetails
            err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        case 415:
            var v models.ProblemDetails
            err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        case 429:
            var v models.ProblemDetails
            err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        case 500:
            var v models.ProblemDetails
            err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        case 503:
            var v models.ProblemDetails
            err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
            if err != nil {
                apiError.ErrorStatus = err.Error()
                return localVarReturnValue, localVarHTTPResponse, apiError
            }
            apiError.ErrorModel = v
            return localVarReturnValue, localVarHTTPResponse, apiError
        default:
        return localVarReturnValue, localVarHTTPResponse, nil
    }
}
