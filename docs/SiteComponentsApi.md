# \SiteComponentsApi

All URIs are relative to *https://api.mach.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SiteComponentCreate**](SiteComponentsApi.md#SiteComponentCreate) | **Post** /organizations/{organization}/projects/{project}/site-components | Create site component
[**SiteComponentDelete**](SiteComponentsApi.md#SiteComponentDelete) | **Delete** /organizations/{organization}/projects/{project}/site-components/{site-component} | Delete a site component
[**SiteComponentGet**](SiteComponentsApi.md#SiteComponentGet) | **Get** /organizations/{organization}/projects/{project}/site-components/{site-component} | Get an existing site component
[**SiteComponentPatch**](SiteComponentsApi.md#SiteComponentPatch) | **Patch** /organizations/{organization}/projects/{project}/site-components/{site-component} | Patch an existing site component
[**SiteComponentQuery**](SiteComponentsApi.md#SiteComponentQuery) | **Get** /organizations/{organization}/projects/{project}/site-components | List all sites components



## SiteComponentCreate

> SiteComponent SiteComponentCreate(ctx, organization, project).SiteComponentDraft(siteComponentDraft).Execute()

Create site component

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
	siteComponentDraft := *openapiclient.NewSiteComponentDraft("Key_example", "Name_example", "SiteKey_example", "ComponentKey_example") // SiteComponentDraft | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteComponentsApi.SiteComponentCreate(context.Background(), organization, project).SiteComponentDraft(siteComponentDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteComponentsApi.SiteComponentCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SiteComponentCreate`: SiteComponent
	fmt.Fprintf(os.Stdout, "Response from `SiteComponentsApi.SiteComponentCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSiteComponentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteComponentDraft** | [**SiteComponentDraft**](SiteComponentDraft.md) |  | 

### Return type

[**SiteComponent**](SiteComponent.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteComponentDelete

> SiteComponent SiteComponentDelete(ctx, organization, project, siteComponent).Execute()

Delete a site component



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
	siteComponent := "siteComponent_example" // string | Site component key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteComponentsApi.SiteComponentDelete(context.Background(), organization, project, siteComponent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteComponentsApi.SiteComponentDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SiteComponentDelete`: SiteComponent
	fmt.Fprintf(os.Stdout, "Response from `SiteComponentsApi.SiteComponentDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**siteComponent** | **string** | Site component key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSiteComponentDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SiteComponent**](SiteComponent.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteComponentGet

> SiteComponent SiteComponentGet(ctx, organization, project, siteComponent).Execute()

Get an existing site component

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
	siteComponent := "siteComponent_example" // string | Site component key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteComponentsApi.SiteComponentGet(context.Background(), organization, project, siteComponent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteComponentsApi.SiteComponentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SiteComponentGet`: SiteComponent
	fmt.Fprintf(os.Stdout, "Response from `SiteComponentsApi.SiteComponentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**siteComponent** | **string** | Site component key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSiteComponentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SiteComponent**](SiteComponent.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteComponentPatch

> SiteComponent SiteComponentPatch(ctx, organization, project, siteComponent).PatchRequestInner(patchRequestInner).Execute()

Patch an existing site component

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
	siteComponent := "siteComponent_example" // string | Site component key
	patchRequestInner := []openapiclient.PatchRequestInner{openapiclient.PatchRequest_inner{JSONPatchRequestAddReplaceTest: openapiclient.NewJSONPatchRequestAddReplaceTest("Path_example", interface{}(123), "Op_example")}} // []PatchRequestInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteComponentsApi.SiteComponentPatch(context.Background(), organization, project, siteComponent).PatchRequestInner(patchRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteComponentsApi.SiteComponentPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SiteComponentPatch`: SiteComponent
	fmt.Fprintf(os.Stdout, "Response from `SiteComponentsApi.SiteComponentPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**siteComponent** | **string** | Site component key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSiteComponentPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchRequestInner** | [**[]PatchRequestInner**](PatchRequestInner.md) |  | 

### Return type

[**SiteComponent**](SiteComponent.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteComponentQuery

> SiteComponentPaginator SiteComponentQuery(ctx, organization, project).Offset(offset).Limit(limit).Sort(sort).Execute()

List all sites components

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
	resp, r, err := apiClient.SiteComponentsApi.SiteComponentQuery(context.Background(), organization, project).Offset(offset).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteComponentsApi.SiteComponentQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SiteComponentQuery`: SiteComponentPaginator
	fmt.Fprintf(os.Stdout, "Response from `SiteComponentsApi.SiteComponentQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSiteComponentQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 

### Return type

[**SiteComponentPaginator**](SiteComponentPaginator.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

