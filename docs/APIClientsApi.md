# \APIClientsApi

All URIs are relative to *https://api.mach.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiClientCreate**](APIClientsApi.md#ApiClientCreate) | **Post** /organizations/{organization}/projects/{project}/api-clients | Create new API Client
[**ApiClientDelete**](APIClientsApi.md#ApiClientDelete) | **Delete** /organizations/{organization}/projects/{project}/api-clients/{id} | Delete API Client
[**ApiClientGet**](APIClientsApi.md#ApiClientGet) | **Get** /organizations/{organization}/projects/{project}/api-clients/{id} | Get API Client details
[**ApiClientPatch**](APIClientsApi.md#ApiClientPatch) | **Patch** /organizations/{organization}/projects/{project}/api-clients/{id} | Update API Client
[**ApiClientQuery**](APIClientsApi.md#ApiClientQuery) | **Get** /organizations/{organization}/projects/{project}/api-clients | List all API Clients
[**ApiClientUpdate**](APIClientsApi.md#ApiClientUpdate) | **Put** /organizations/{organization}/projects/{project}/api-clients/{id} | Update API Client



## ApiClientCreate

> ApiClient ApiClientCreate(ctx, organization, project).ApiClientDraft(apiClientDraft).Execute()

Create new API Client

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	apiClientDraft := *openapiclient.NewApiClientDraft("ClientId_example", "ClientSecret_example") // ApiClientDraft | 

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
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClientCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiClientDraft** | [**ApiClientDraft**](ApiClientDraft.md) |  | 

### Return type

[**ApiClient**](ApiClient.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiClientDelete

> ApiClientDelete(ctx, organization, project, id).Execute()

Delete API Client

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.APIClientsApi.ApiClientDelete(context.Background(), organization, project, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIClientsApi.ApiClientDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClientDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiClientGet

> ApiClient ApiClientGet(ctx, organization, project, id).Execute()

Get API Client details

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIClientsApi.ApiClientGet(context.Background(), organization, project, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIClientsApi.ApiClientGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiClientGet`: ApiClient
	fmt.Fprintf(os.Stdout, "Response from `APIClientsApi.ApiClientGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClientGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiClient**](ApiClient.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiClientPatch

> ApiClient ApiClientPatch(ctx, organization, project, id).PatchedApiClientDraft(patchedApiClientDraft).Execute()

Update API Client

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	id := "id_example" // string | 
	patchedApiClientDraft := *openapiclient.NewPatchedApiClientDraft() // PatchedApiClientDraft |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIClientsApi.ApiClientPatch(context.Background(), organization, project, id).PatchedApiClientDraft(patchedApiClientDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIClientsApi.ApiClientPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiClientPatch`: ApiClient
	fmt.Fprintf(os.Stdout, "Response from `APIClientsApi.ApiClientPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClientPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchedApiClientDraft** | [**PatchedApiClientDraft**](PatchedApiClientDraft.md) |  | 

### Return type

[**ApiClient**](ApiClient.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiClientQuery

> ApiClientPaginator ApiClientQuery(ctx, organization, project).Limit(limit).Offset(offset).Execute()

List all API Clients

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIClientsApi.ApiClientQuery(context.Background(), organization, project).Limit(limit).Offset(offset).Execute()
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
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClientQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**ApiClientPaginator**](ApiClientPaginator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiClientUpdate

> ApiClient ApiClientUpdate(ctx, organization, project, id).ApiClientDraft(apiClientDraft).Execute()

Update API Client

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	id := "id_example" // string | 
	apiClientDraft := *openapiclient.NewApiClientDraft("ClientId_example", "ClientSecret_example") // ApiClientDraft | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIClientsApi.ApiClientUpdate(context.Background(), organization, project, id).ApiClientDraft(apiClientDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIClientsApi.ApiClientUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiClientUpdate`: ApiClient
	fmt.Fprintf(os.Stdout, "Response from `APIClientsApi.ApiClientUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClientUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiClientDraft** | [**ApiClientDraft**](ApiClientDraft.md) |  | 

### Return type

[**ApiClient**](ApiClient.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

