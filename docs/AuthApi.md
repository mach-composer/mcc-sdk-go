# \AuthApi

All URIs are relative to *https://api.mach.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authorize**](AuthApi.md#Authorize) | **Get** /authorize | Start authorization flow
[**GetAuthToken**](AuthApi.md#GetAuthToken) | **Post** /oauth/token | Return a new token
[**IntrospectToken**](AuthApi.md#IntrospectToken) | **Post** /oauth/introspect | Introspect an existing token



## Authorize

> Authorize(ctx).ResponseType(responseType).Scope(scope).RedirectUri(redirectUri).State(state).ClientId(clientId).CodeChallenge(codeChallenge).CodeChallengeMethod(codeChallengeMethod).Provider(provider).Execute()

Start authorization flow

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    responseType := "responseType_example" // string | 
    scope := "scope_example" // string | 
    redirectUri := "redirectUri_example" // string | 
    state := "state_example" // string | 
    clientId := "clientId_example" // string | 
    codeChallenge := "codeChallenge_example" // string | 
    codeChallengeMethod := "S256" // string | 
    provider := "provider_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.Authorize(context.Background()).ResponseType(responseType).Scope(scope).RedirectUri(redirectUri).State(state).ClientId(clientId).CodeChallenge(codeChallenge).CodeChallengeMethod(codeChallengeMethod).Provider(provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.Authorize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **responseType** | **string** |  | 
 **scope** | **string** |  | 
 **redirectUri** | **string** |  | 
 **state** | **string** |  | 
 **clientId** | **string** |  | 
 **codeChallenge** | **string** |  | 
 **codeChallengeMethod** | **string** |  | 
 **provider** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthToken

> GetAuthToken(ctx).GrantType(grantType).ClientId(clientId).CodeVerifier(codeVerifier).Code(code).RedirectUri(redirectUri).RefreshToken(refreshToken).Execute()

Return a new token

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    grantType := "grantType_example" // string | 
    clientId := "clientId_example" // string | 
    codeVerifier := "codeVerifier_example" // string |  (optional)
    code := "code_example" // string |  (optional)
    redirectUri := "redirectUri_example" // string |  (optional)
    refreshToken := "refreshToken_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.GetAuthToken(context.Background()).GrantType(grantType).ClientId(clientId).CodeVerifier(codeVerifier).Code(code).RedirectUri(redirectUri).RefreshToken(refreshToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetAuthToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** |  | 
 **clientId** | **string** |  | 
 **codeVerifier** | **string** |  | 
 **code** | **string** |  | 
 **redirectUri** | **string** |  | 
 **refreshToken** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntrospectToken

> IntrospectToken(ctx).Token(token).Execute()

Introspect an existing token

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.IntrospectToken(context.Background()).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.IntrospectToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntrospectTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

