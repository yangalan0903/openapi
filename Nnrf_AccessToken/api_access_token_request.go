/*
 * NRF OAuth2
 *
 * NRF OAuth2 Authorization
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_AccessToken

import (
	"github.com/yangalan0903/openapi"
	"github.com/yangalan0903/openapi/models"

	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type AccessTokenRequestApiService service

/*
AccessTokenRequestApiService Access Token Request
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param grantType
 * @param nfInstanceId
 * @param scope
 * @param optional nil or *AccessTokenRequestParamOpts - Optional Parameters:
 * @param "NfType" (optional.Interface of models.NfType) -
 * @param "TargetNfType" (optional.Interface of models.NfType) -
 * @param "TargetNfInstanceId" (optional.Interface of string) -
 * @param "RequesterPlmn" (optional.Interface of models.PlmnId) -
 * @param "TargetPlmn" (optional.Interface of models.PlmnId) -
@return models.AccessTokenRsp
*/

type AccessTokenRequestParamOpts struct {
	NfType             optional.Interface
	TargetNfType       optional.Interface
	TargetNfInstanceId optional.Interface
	RequesterPlmn      optional.Interface
	TargetPlmn         optional.Interface
}

func (a *AccessTokenRequestApiService) AccessTokenRequest(ctx context.Context, grantType string, nfInstanceId string, scope string, localVarOptionals *AccessTokenRequestParamOpts) (models.AccessTokenRsp, *http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  models.AccessTokenRsp
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/oauth2/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	localVarFormParams.Add("grant_type", openapi.ParameterToString(grantType, ""))
	localVarFormParams.Add("nfInstanceId", openapi.ParameterToString(nfInstanceId, ""))
	if localVarOptionals != nil && localVarOptionals.NfType.IsSet() {
		localVarFormParams.Add("nfType", openapi.ParameterToString(localVarOptionals.NfType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TargetNfType.IsSet() {
		localVarFormParams.Add("targetNfType", openapi.ParameterToString(localVarOptionals.TargetNfType.Value(), ""))
	}
	localVarFormParams.Add("scope", openapi.ParameterToString(scope, ""))
	if localVarOptionals != nil && localVarOptionals.TargetNfInstanceId.IsSet() {
		localVarFormParams.Add("targetNfInstanceId", openapi.ParameterToString(localVarOptionals.TargetNfInstanceId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RequesterPlmn.IsSet() {
		localVarFormParams.Add("requesterPlmn", openapi.ParameterToString(localVarOptionals.RequesterPlmn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TargetPlmn.IsSet() {
		localVarFormParams.Add("targetPlmn", openapi.ParameterToString(localVarOptionals.TargetPlmn.Value(), ""))
	}

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
		var v models.AccessTokenErr
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	default:
		return localVarReturnValue, localVarHTTPResponse, openapi.ReportError("%d is not a valid status code in AccessTokenRequest", localVarHTTPResponse.StatusCode)
	}
}
