# \ComponentsAPI

All URIs are relative to *https://api.mach.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentCreate**](ComponentsAPI.md#ComponentCreate) | **Post** /organizations/{organization}/projects/{project}/components | Create component
[**ComponentDelete**](ComponentsAPI.md#ComponentDelete) | **Delete** /organizations/{organization}/projects/{project}/components/{component} | Delete a component
[**ComponentLatestVersion**](ComponentsAPI.md#ComponentLatestVersion) | **Get** /organizations/{organization}/projects/{project}/components/{component}/latest | Get last component version
[**ComponentPatch**](ComponentsAPI.md#ComponentPatch) | **Patch** /organizations/{organization}/projects/{project}/components/{component} | Patch an existing component
[**ComponentQuery**](ComponentsAPI.md#ComponentQuery) | **Get** /organizations/{organization}/projects/{project}/components | List all components
[**ComponentVersionCreate**](ComponentsAPI.md#ComponentVersionCreate) | **Post** /organizations/{organization}/projects/{project}/components/{component}/versions | Create component version
[**ComponentVersionDelete**](ComponentsAPI.md#ComponentVersionDelete) | **Delete** /organizations/{organization}/projects/{project}/components/{component}/versions/{version} | Delete component version
[**ComponentVersionGet**](ComponentsAPI.md#ComponentVersionGet) | **Get** /organizations/{organization}/projects/{project}/components/{component}/versions/{version} | Get component version
[**ComponentVersionPushCommits**](ComponentsAPI.md#ComponentVersionPushCommits) | **Post** /organizations/{organization}/projects/{project}/components/{component}/versions/{version}/commits | Push commits for this component version
[**ComponentVersionQuery**](ComponentsAPI.md#ComponentVersionQuery) | **Get** /organizations/{organization}/projects/{project}/components/{component}/versions | List all versions of a component
[**ComponentVersionQueryCommits**](ComponentsAPI.md#ComponentVersionQueryCommits) | **Get** /organizations/{organization}/projects/{project}/components/{component}/versions/{version}/commits | Get commits for this component version



## ComponentCreate

> Component ComponentCreate(ctx, organization, project).ComponentDraft(componentDraft).Execute()

Create component

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
    componentDraft := *openapiclient.NewComponentDraft("Key_example", "Name_example") // ComponentDraft | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsAPI.ComponentCreate(context.Background(), organization, project).ComponentDraft(componentDraft).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.ComponentCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentCreate`: Component
    fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.ComponentCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **componentDraft** | [**ComponentDraft**](ComponentDraft.md) |  | 

### Return type

[**Component**](Component.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentDelete

> Component ComponentDelete(ctx, organization, project, component).Execute()

Delete a component



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
    component := "component_example" // string | Component key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsAPI.ComponentDelete(context.Background(), organization, project, component).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.ComponentDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentDelete`: Component
    fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.ComponentDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**component** | **string** | Component key | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Component**](Component.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentLatestVersion

> ComponentVersion ComponentLatestVersion(ctx, organization, project, component).Branch(branch).Execute()

Get last component version

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
    component := "component_example" // string | Component key
    branch := "branch_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsAPI.ComponentLatestVersion(context.Background(), organization, project, component).Branch(branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.ComponentLatestVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentLatestVersion`: ComponentVersion
    fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.ComponentLatestVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**component** | **string** | Component key | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentLatestVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **branch** | **string** |  | 

### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentPatch

> Component ComponentPatch(ctx, organization, project, component).PatchRequestInner(patchRequestInner).Execute()

Patch an existing component

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
    component := "component_example" // string | Component key
    patchRequestInner := []openapiclient.PatchRequestInner{openapiclient.PatchRequest_inner{JSONPatchRequestAddReplaceTest: openapiclient.NewJSONPatchRequestAddReplaceTest("Path_example", interface{}(123), "Op_example")}} // []PatchRequestInner |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsAPI.ComponentPatch(context.Background(), organization, project, component).PatchRequestInner(patchRequestInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.ComponentPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentPatch`: Component
    fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.ComponentPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**component** | **string** | Component key | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchRequestInner** | [**[]PatchRequestInner**](PatchRequestInner.md) |  | 

### Return type

[**Component**](Component.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentQuery

> ComponentPaginator ComponentQuery(ctx, organization, project).Offset(offset).Limit(limit).Sort(sort).Execute()

List all components

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
    resp, r, err := apiClient.ComponentsAPI.ComponentQuery(context.Background(), organization, project).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.ComponentQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentQuery`: ComponentPaginator
    fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.ComponentQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 

### Return type

[**ComponentPaginator**](ComponentPaginator.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionCreate

> ComponentVersion ComponentVersionCreate(ctx, organization, project, component).ComponentVersionDraft(componentVersionDraft).Execute()

Create component version

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
    component := "component_example" // string | Component key
    componentVersionDraft := *openapiclient.NewComponentVersionDraft("Version_example", "Branch_example") // ComponentVersionDraft | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsAPI.ComponentVersionCreate(context.Background(), organization, project, component).ComponentVersionDraft(componentVersionDraft).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.ComponentVersionCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentVersionCreate`: ComponentVersion
    fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.ComponentVersionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**component** | **string** | Component key | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **componentVersionDraft** | [**ComponentVersionDraft**](ComponentVersionDraft.md) |  | 

### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionDelete

> ComponentVersion ComponentVersionDelete(ctx, organization, project, component, version).Execute()

Delete component version

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
    component := "component_example" // string | Component key
    version := "version_example" // string | Version

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsAPI.ComponentVersionDelete(context.Background(), organization, project, component, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.ComponentVersionDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentVersionDelete`: ComponentVersion
    fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.ComponentVersionDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**component** | **string** | Component key | 
**version** | **string** | Version | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionGet

> ComponentVersion ComponentVersionGet(ctx, organization, project, component, version).Execute()

Get component version

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
    component := "component_example" // string | Component key
    version := "version_example" // string | Version

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsAPI.ComponentVersionGet(context.Background(), organization, project, component, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.ComponentVersionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentVersionGet`: ComponentVersion
    fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.ComponentVersionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**component** | **string** | Component key | 
**version** | **string** | Version | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionPushCommits

> ComponentVersionPushCommits(ctx, organization, project, component, version).ComponentVersionCommits(componentVersionCommits).Execute()

Push commits for this component version

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/mach-composer/mcc-sdk-go"
)

func main() {
    organization := "my-organization" // string | Organization Key
    project := "my-project" // string | Project Key
    component := "component_example" // string | Component key
    version := "version_example" // string | Component version
    componentVersionCommits := *openapiclient.NewComponentVersionCommits([]openapiclient.CommitData{*openapiclient.NewCommitData("Commit_example", "Subject_example", *openapiclient.NewCommitDataAuthor("Name_example", "Email_example", time.Now()), *openapiclient.NewCommitDataAuthor("Name_example", "Email_example", time.Now()))}) // ComponentVersionCommits | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ComponentsAPI.ComponentVersionPushCommits(context.Background(), organization, project, component, version).ComponentVersionCommits(componentVersionCommits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.ComponentVersionPushCommits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**component** | **string** | Component key | 
**version** | **string** | Component version | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionPushCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **componentVersionCommits** | [**ComponentVersionCommits**](ComponentVersionCommits.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionQuery

> ComponentVersionPaginator ComponentVersionQuery(ctx, organization, project, component).Offset(offset).Limit(limit).Sort(sort).Execute()

List all versions of a component

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
    component := "component_example" // string | Component key
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional)
    sort := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsAPI.ComponentVersionQuery(context.Background(), organization, project, component).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.ComponentVersionQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentVersionQuery`: ComponentVersionPaginator
    fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.ComponentVersionQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**component** | **string** | Component key | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 

### Return type

[**ComponentVersionPaginator**](ComponentVersionPaginator.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionQueryCommits

> CommitDataPaginator ComponentVersionQueryCommits(ctx, organization, project, component, version).Offset(offset).Limit(limit).Sort(sort).Execute()

Get commits for this component version

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
    component := "component_example" // string | Component key
    version := "version_example" // string | Component version
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional)
    sort := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsAPI.ComponentVersionQueryCommits(context.Background(), organization, project, component, version).Offset(offset).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.ComponentVersionQueryCommits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentVersionQueryCommits`: CommitDataPaginator
    fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.ComponentVersionQueryCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 
**component** | **string** | Component key | 
**version** | **string** | Component version | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionQueryCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | 
 **sort** | **[]string** |  | 

### Return type

[**CommitDataPaginator**](CommitDataPaginator.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

