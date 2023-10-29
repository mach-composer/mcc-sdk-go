# \SitesAPI

All URIs are relative to *https://api.mach.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SiteComponentSiteQuery**](SitesAPI.md#SiteComponentSiteQuery) | **Get** /organizations/{organization}/projects/{project}/sites/{site}/components | List all components for the given site
[**SiteCreate**](SitesAPI.md#SiteCreate) | **Post** /organizations/{organization}/projects/{project}/sites | Create site
[**SiteDelete**](SitesAPI.md#SiteDelete) | **Delete** /organizations/{organization}/projects/{project}/sites/{site} | Delete a site
[**SiteGet**](SitesAPI.md#SiteGet) | **Get** /organizations/{organization}/projects/{project}/sites/{site} | Get an existing site
[**SitePatch**](SitesAPI.md#SitePatch) | **Patch** /organizations/{organization}/projects/{project}/sites/{site} | Patch an existing site
[**SiteQuery**](SitesAPI.md#SiteQuery) | **Get** /organizations/{organization}/projects/{project}/sites | List all sites



## SiteComponentSiteQuery

> SiteComponentPaginator SiteComponentSiteQuery(ctx, organization, project, site).Offset(offset).Limit(limit).Sort(sort).Execute()

List all components for the given site

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mach-composer/mcc-sdk-go"
)

func main() {
    organization := "my-organization" // string | Organization Key
    project := "my-project" // string | Project Key
    site := "site_example" // string | Site key
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional)
    sort := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPI.SiteComponentSiteQuery(context.Background(), organization, project, site).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.SiteComponentSiteQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiteComponentSiteQuery`: SiteComponentPaginator
    fmt.Fprintf(os.Stdout, "Response from `SitesAPI.SiteComponentSiteQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**site** | **string** | Site key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSiteComponentSiteQueryRequest struct via the builder pattern


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


## SiteCreate

> Site SiteCreate(ctx, organization, project).SiteDraft(siteDraft).Execute()

Create site

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mach-composer/mcc-sdk-go"
)

func main() {
    organization := "my-organization" // string | Organization Key
    project := "my-project" // string | Project Key
    siteDraft := *openapiclient.NewSiteDraft("Key_example", "Name_example") // SiteDraft | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPI.SiteCreate(context.Background(), organization, project).SiteDraft(siteDraft).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.SiteCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiteCreate`: Site
    fmt.Fprintf(os.Stdout, "Response from `SitesAPI.SiteCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSiteCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteDraft** | [**SiteDraft**](SiteDraft.md) |  | 

### Return type

[**Site**](Site.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteDelete

> Site SiteDelete(ctx, organization, project, site).Execute()

Delete a site



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mach-composer/mcc-sdk-go"
)

func main() {
    organization := "my-organization" // string | Organization Key
    project := "my-project" // string | Project Key
    site := "site_example" // string | Site key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPI.SiteDelete(context.Background(), organization, project, site).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.SiteDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiteDelete`: Site
    fmt.Fprintf(os.Stdout, "Response from `SitesAPI.SiteDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**site** | **string** | Site key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSiteDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Site**](Site.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteGet

> Site SiteGet(ctx, organization, project, site).Execute()

Get an existing site

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mach-composer/mcc-sdk-go"
)

func main() {
    organization := "my-organization" // string | Organization Key
    project := "my-project" // string | Project Key
    site := "site_example" // string | Site key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPI.SiteGet(context.Background(), organization, project, site).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.SiteGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiteGet`: Site
    fmt.Fprintf(os.Stdout, "Response from `SitesAPI.SiteGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**site** | **string** | Site key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSiteGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Site**](Site.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitePatch

> Site SitePatch(ctx, organization, project, site).PatchRequestInner(patchRequestInner).Execute()

Patch an existing site

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mach-composer/mcc-sdk-go"
)

func main() {
    organization := "my-organization" // string | Organization Key
    project := "my-project" // string | Project Key
    site := "site_example" // string | Site key
    patchRequestInner := []openapiclient.PatchRequestInner{openapiclient.PatchRequest_inner{JSONPatchRequestAddReplaceTest: openapiclient.NewJSONPatchRequestAddReplaceTest("Path_example", interface{}(123), "Op_example")}} // []PatchRequestInner |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPI.SitePatch(context.Background(), organization, project, site).PatchRequestInner(patchRequestInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.SitePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitePatch`: Site
    fmt.Fprintf(os.Stdout, "Response from `SitesAPI.SitePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**site** | **string** | Site key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchRequestInner** | [**[]PatchRequestInner**](PatchRequestInner.md) |  | 

### Return type

[**Site**](Site.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteQuery

> SitePaginator SiteQuery(ctx, organization, project).Offset(offset).Limit(limit).Sort(sort).Execute()

List all sites

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mach-composer/mcc-sdk-go"
)

func main() {
    organization := "my-organization" // string | Organization Key
    project := "my-project" // string | Project Key
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional)
    sort := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesAPI.SiteQuery(context.Background(), organization, project).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.SiteQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiteQuery`: SitePaginator
    fmt.Fprintf(os.Stdout, "Response from `SitesAPI.SiteQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSiteQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 

### Return type

[**SitePaginator**](SitePaginator.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

