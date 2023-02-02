/*
MACH composer Cloud (MCC) Public API

# Introduction  MACH composer Cloud is a platform and API to facilitate and coordinate work across teams that build composable architectures using MACH technology.   All operations available in MACH composer cloud are available through this API. For more information about using it in your MACH architecture, have a look at the [documentation website](https://docs.machcomposer.io/cloud/index.html).

API version: 0.1.0
Contact: mach@labdigital.nl
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mccsdk

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AuthApi interface {

	/*
		Authorize Start authorization flow

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiAuthorizeRequest
	*/
	Authorize(ctx context.Context) ApiAuthorizeRequest

	// AuthorizeExecute executes the request
	AuthorizeExecute(r ApiAuthorizeRequest) (*http.Response, error)

	/*
		GetAuthToken Return a new token

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetAuthTokenRequest
	*/
	GetAuthToken(ctx context.Context) ApiGetAuthTokenRequest

	// GetAuthTokenExecute executes the request
	GetAuthTokenExecute(r ApiGetAuthTokenRequest) (*http.Response, error)

	/*
		IntrospectToken Introspect an existing token

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiIntrospectTokenRequest
	*/
	IntrospectToken(ctx context.Context) ApiIntrospectTokenRequest

	// IntrospectTokenExecute executes the request
	IntrospectTokenExecute(r ApiIntrospectTokenRequest) (*http.Response, error)
}

// AuthApiService AuthApi service
type AuthApiService service

type ApiAuthorizeRequest struct {
	ctx                 context.Context
	ApiService          AuthApi
	responseType        *string
	scope               *string
	redirectUri         *string
	state               *string
	clientId            *string
	codeChallenge       *string
	codeChallengeMethod *string
	provider            *string
}

func (r ApiAuthorizeRequest) ResponseType(responseType string) ApiAuthorizeRequest {
	r.responseType = &responseType
	return r
}

func (r ApiAuthorizeRequest) Scope(scope string) ApiAuthorizeRequest {
	r.scope = &scope
	return r
}

func (r ApiAuthorizeRequest) RedirectUri(redirectUri string) ApiAuthorizeRequest {
	r.redirectUri = &redirectUri
	return r
}

func (r ApiAuthorizeRequest) State(state string) ApiAuthorizeRequest {
	r.state = &state
	return r
}

func (r ApiAuthorizeRequest) ClientId(clientId string) ApiAuthorizeRequest {
	r.clientId = &clientId
	return r
}

func (r ApiAuthorizeRequest) CodeChallenge(codeChallenge string) ApiAuthorizeRequest {
	r.codeChallenge = &codeChallenge
	return r
}

func (r ApiAuthorizeRequest) CodeChallengeMethod(codeChallengeMethod string) ApiAuthorizeRequest {
	r.codeChallengeMethod = &codeChallengeMethod
	return r
}

func (r ApiAuthorizeRequest) Provider(provider string) ApiAuthorizeRequest {
	r.provider = &provider
	return r
}

func (r ApiAuthorizeRequest) Execute() (*http.Response, error) {
	return r.ApiService.AuthorizeExecute(r)
}

/*
Authorize Start authorization flow

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAuthorizeRequest
*/
func (a *AuthApiService) Authorize(ctx context.Context) ApiAuthorizeRequest {
	return ApiAuthorizeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *AuthApiService) AuthorizeExecute(r ApiAuthorizeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.Authorize")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/authorize"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.responseType == nil {
		return nil, reportError("responseType is required and must be specified")
	}
	if r.scope == nil {
		return nil, reportError("scope is required and must be specified")
	}
	if r.redirectUri == nil {
		return nil, reportError("redirectUri is required and must be specified")
	}
	if r.state == nil {
		return nil, reportError("state is required and must be specified")
	}
	if r.clientId == nil {
		return nil, reportError("clientId is required and must be specified")
	}
	if r.codeChallenge == nil {
		return nil, reportError("codeChallenge is required and must be specified")
	}
	if r.codeChallengeMethod == nil {
		return nil, reportError("codeChallengeMethod is required and must be specified")
	}

	localVarQueryParams.Add("response_type", parameterToString(*r.responseType, ""))
	localVarQueryParams.Add("scope", parameterToString(*r.scope, ""))
	localVarQueryParams.Add("redirect_uri", parameterToString(*r.redirectUri, ""))
	localVarQueryParams.Add("state", parameterToString(*r.state, ""))
	localVarQueryParams.Add("client_id", parameterToString(*r.clientId, ""))
	localVarQueryParams.Add("code_challenge", parameterToString(*r.codeChallenge, ""))
	localVarQueryParams.Add("code_challenge_method", parameterToString(*r.codeChallengeMethod, ""))
	if r.provider != nil {
		localVarQueryParams.Add("provider", parameterToString(*r.provider, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetAuthTokenRequest struct {
	ctx          context.Context
	ApiService   AuthApi
	grantType    *string
	clientId     *string
	codeVerifier *string
	code         *string
	redirectUri  *string
	refreshToken *string
}

func (r ApiGetAuthTokenRequest) GrantType(grantType string) ApiGetAuthTokenRequest {
	r.grantType = &grantType
	return r
}

func (r ApiGetAuthTokenRequest) ClientId(clientId string) ApiGetAuthTokenRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetAuthTokenRequest) CodeVerifier(codeVerifier string) ApiGetAuthTokenRequest {
	r.codeVerifier = &codeVerifier
	return r
}

func (r ApiGetAuthTokenRequest) Code(code string) ApiGetAuthTokenRequest {
	r.code = &code
	return r
}

func (r ApiGetAuthTokenRequest) RedirectUri(redirectUri string) ApiGetAuthTokenRequest {
	r.redirectUri = &redirectUri
	return r
}

func (r ApiGetAuthTokenRequest) RefreshToken(refreshToken string) ApiGetAuthTokenRequest {
	r.refreshToken = &refreshToken
	return r
}

func (r ApiGetAuthTokenRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetAuthTokenExecute(r)
}

/*
GetAuthToken Return a new token

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAuthTokenRequest
*/
func (a *AuthApiService) GetAuthToken(ctx context.Context) ApiGetAuthTokenRequest {
	return ApiGetAuthTokenRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *AuthApiService) GetAuthTokenExecute(r ApiGetAuthTokenRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.GetAuthToken")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.grantType == nil {
		return nil, reportError("grantType is required and must be specified")
	}
	if r.clientId == nil {
		return nil, reportError("clientId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("grant_type", parameterToString(*r.grantType, ""))
	localVarFormParams.Add("client_id", parameterToString(*r.clientId, ""))
	if r.codeVerifier != nil {
		localVarFormParams.Add("code_verifier", parameterToString(*r.codeVerifier, ""))
	}
	if r.code != nil {
		localVarFormParams.Add("code", parameterToString(*r.code, ""))
	}
	if r.redirectUri != nil {
		localVarFormParams.Add("redirect_uri", parameterToString(*r.redirectUri, ""))
	}
	if r.refreshToken != nil {
		localVarFormParams.Add("refresh_token", parameterToString(*r.refreshToken, ""))
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiIntrospectTokenRequest struct {
	ctx        context.Context
	ApiService AuthApi
	token      *string
}

func (r ApiIntrospectTokenRequest) Token(token string) ApiIntrospectTokenRequest {
	r.token = &token
	return r
}

func (r ApiIntrospectTokenRequest) Execute() (*http.Response, error) {
	return r.ApiService.IntrospectTokenExecute(r)
}

/*
IntrospectToken Introspect an existing token

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiIntrospectTokenRequest
*/
func (a *AuthApiService) IntrospectToken(ctx context.Context) ApiIntrospectTokenRequest {
	return ApiIntrospectTokenRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *AuthApiService) IntrospectTokenExecute(r ApiIntrospectTokenRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.IntrospectToken")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/introspect"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.token != nil {
		localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
