# \AuthApi

All URIs are relative to *https://api.mach.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authorize**](AuthApi.md#Authorize) | **Get** /authorize | Start authorization flow
[**GetAuthToken**](AuthApi.md#GetAuthToken) | **Post** /oauth/token | Return a new token
[**IntrospectToken**](AuthApi.md#IntrospectToken) | **Post** /oauth/introspect | Introspect an existing token



## Authorize

> Authorize(ctx).Audience(audience).ResponseType(responseType).Provider(provider).CodeChallenge(codeChallenge).CodeChallengeMethod(codeChallengeMethod).RedirectUri(redirectUri).Execute()

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
    audience := "audience_example" // string |  (optional)
    responseType := "responseType_example" // string |  (optional)
    provider := "provider_example" // string |  (optional)
    codeChallenge := "codeChallenge_example" // string |  (optional)
    codeChallengeMethod := "S256" // string |  (optional)
    redirectUri := "redirectUri_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.Authorize(context.Background()).Audience(audience).ResponseType(responseType).Provider(provider).CodeChallenge(codeChallenge).CodeChallengeMethod(codeChallengeMethod).RedirectUri(redirectUri).Execute()
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
 **audience** | **string** |  | 
 **responseType** | **string** |  | 
 **provider** | **string** |  | 
 **codeChallenge** | **string** |  | 
 **codeChallengeMethod** | **string** |  | 
 **redirectUri** | **string** |  | 

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

> GetAuthToken(ctx).GrantType(grantType).ClientId(clientId).CodeVerifier(codeVerifier).Code(code).RedirectUri(redirectUri).Execute()

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
    grantType := "client_credentials" // string |  (optional)
    clientId := "clientId_example" // string |  (optional)
    codeVerifier := "codeVerifier_example" // string |  (optional)
    code := "code_example" // string |  (optional)
    redirectUri := "redirectUri_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.GetAuthToken(context.Background()).GrantType(grantType).ClientId(clientId).CodeVerifier(codeVerifier).Code(code).RedirectUri(redirectUri).Execute()
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

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
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

