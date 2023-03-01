# \MyAccountApi

All URIs are relative to *https://api.mach.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MyAccountInformation**](MyAccountApi.md#MyAccountInformation) | **Get** /account/me | Get user information



## MyAccountInformation

> MyAccountInformation200Response MyAccountInformation(ctx).Execute()

Get user information



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MyAccountApi.MyAccountInformation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MyAccountApi.MyAccountInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MyAccountInformation`: MyAccountInformation200Response
    fmt.Fprintf(os.Stdout, "Response from `MyAccountApi.MyAccountInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMyAccountInformationRequest struct via the builder pattern


### Return type

[**MyAccountInformation200Response**](MyAccountInformation200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

