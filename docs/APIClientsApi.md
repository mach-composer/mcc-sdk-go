# \APIClientsApi

All URIs are relative to *https://api.mach.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiClientCreate**](APIClientsApi.md#ApiClientCreate) | **Post** /organizations/{organization}/projects/{project}/api-clients | Create new api client
[**ApiClientDelete**](APIClientsApi.md#ApiClientDelete) | **Delete** /organizations/{organization}/projects/{project}/api-clients/{id} | Delete an API Client
[**ApiClientQuery**](APIClientsApi.md#ApiClientQuery) | **Get** /organizations/{organization}/projects/{project}/api-clients | List all api clients



## ApiClientCreate

> ApiClient ApiClientCreate(ctx, organization, project).ApiClientDraft(apiClientDraft).Execute()

Create new api client

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "my-organization" // string | Organization Key
	project := "my-project" // string | Project Key
	apiClientDraft := *openapiclient.NewApiClientDraft() // ApiClientDraft | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIClientsApi.ApiClientCreate(context.Background(), organization, project).ApiClientDraft(apiClientDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIClientsApi.ApiClientCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiClientCreate`: ApiClient
	fmt.Fprintf(os.Stdout, "Response from `APIClientsApi.ApiClientCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClientCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiClientDraft** | [**ApiClientDraft**](ApiClientDraft.md) |  | 

### Return type

[**ApiClient**](ApiClient.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiClientDelete

> ApiClient ApiClientDelete(ctx, organization, project, id).Execute()

Delete an API Client

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "my-organization" // string | Organization Key
	project := "my-project" // string | Project Key
	id := "id_example" // string | API Client ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIClientsApi.ApiClientDelete(context.Background(), organization, project, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIClientsApi.ApiClientDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiClientDelete`: ApiClient
	fmt.Fprintf(os.Stdout, "Response from `APIClientsApi.ApiClientDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**id** | **string** | API Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClientDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiClient**](ApiClient.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiClientQuery

> ApiClientPaginator ApiClientQuery(ctx, organization, project).Offset(offset).Limit(limit).Sort(sort).Execute()

List all api clients

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "my-organization" // string | Organization Key
	project := "my-project" // string | Project Key
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional)
	sort := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIClientsApi.ApiClientQuery(context.Background(), organization, project).Offset(offset).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIClientsApi.ApiClientQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiClientQuery`: ApiClientPaginator
	fmt.Fprintf(os.Stdout, "Response from `APIClientsApi.ApiClientQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClientQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 

### Return type

[**ApiClientPaginator**](ApiClientPaginator.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

