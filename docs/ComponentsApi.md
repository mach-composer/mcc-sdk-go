# \ComponentsApi

All URIs are relative to *https://api.mach.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentCreate**](ComponentsApi.md#ComponentCreate) | **Post** /{organization}/projects/{project}/components | Create component
[**ComponentLatestVersion**](ComponentsApi.md#ComponentLatestVersion) | **Get** /{organization}/projects/{project}/components/{component}/latest | Get last component version.
[**ComponentQuery**](ComponentsApi.md#ComponentQuery) | **Get** /{organization}/projects/{project}/components | List all components
[**ComponentVersionCreate**](ComponentsApi.md#ComponentVersionCreate) | **Post** /{organization}/projects/{project}/components/{component}/versions | Create component
[**ComponentVersionGetCommits**](ComponentsApi.md#ComponentVersionGetCommits) | **Get** /{organization}/projects/{project}/components/{component}/versions/{version}/commits | Get commits for this component version
[**ComponentVersionPushCommits**](ComponentsApi.md#ComponentVersionPushCommits) | **Post** /{organization}/projects/{project}/components/{component}/versions/{version}/commits | Push commits for this component version
[**ComponentVersionQuery**](ComponentsApi.md#ComponentVersionQuery) | **Get** /{organization}/projects/{project}/components/{component}/versions | List all versions of a component



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
    openapiclient "./openapi"
)

func main() {
    organization := "my-organization" // string | Organization Key
    project := "my-project" // string | Project Key
    componentDraft := *openapiclient.NewComponentDraft("Key_example") // ComponentDraft | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsApi.ComponentCreate(context.Background(), organization, project).ComponentDraft(componentDraft).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentCreate`: Component
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentCreate`: %v\n", resp)
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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentLatestVersion

> ComponentVersion ComponentLatestVersion(ctx, organization, project, component).Branch(branch).Execute()

Get last component version.

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
    organization := "my-organization" // string | Organization Key
    project := "my-project" // string | Project Key
    component := "component_example" // string | Component key
    branch := "branch_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsApi.ComponentLatestVersion(context.Background(), organization, project, component).Branch(branch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentLatestVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentLatestVersion`: ComponentVersion
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentLatestVersion`: %v\n", resp)
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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentQuery

> ComponentPaginator ComponentQuery(ctx, organization, project).Execute()

List all components

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
    organization := "my-organization" // string | Organization Key
    project := "my-project" // string | Project Key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsApi.ComponentQuery(context.Background(), organization, project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentQuery`: ComponentPaginator
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentQuery`: %v\n", resp)
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



### Return type

[**ComponentPaginator**](ComponentPaginator.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionCreate

> ComponentVersion ComponentVersionCreate(ctx, organization, project, component).ComponentVersionDraft(componentVersionDraft).Execute()

Create component

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
    organization := "my-organization" // string | Organization Key
    project := "my-project" // string | Project Key
    component := "component_example" // string | Component key
    componentVersionDraft := *openapiclient.NewComponentVersionDraft("Component_example", "Version_example", "Branch_example") // ComponentVersionDraft | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsApi.ComponentVersionCreate(context.Background(), organization, project, component).ComponentVersionDraft(componentVersionDraft).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentVersionCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentVersionCreate`: ComponentVersion
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentVersionCreate`: %v\n", resp)
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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionGetCommits

> ComponentVersionCommits ComponentVersionGetCommits(ctx, organization, project, component, version).Execute()

Get commits for this component version

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
    organization := "my-organization" // string | Organization Key
    project := "my-project" // string | Project Key
    component := "component_example" // string | Component key
    version := "version_example" // string | Component version

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsApi.ComponentVersionGetCommits(context.Background(), organization, project, component, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentVersionGetCommits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentVersionGetCommits`: ComponentVersionCommits
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentVersionGetCommits`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiComponentVersionGetCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ComponentVersionCommits**](ComponentVersionCommits.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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
    openapiclient "./openapi"
)

func main() {
    organization := "my-organization" // string | Organization Key
    project := "my-project" // string | Project Key
    component := "component_example" // string | Component key
    version := "version_example" // string | Component version
    componentVersionCommits := *openapiclient.NewComponentVersionCommits([]openapiclient.CommitData{*openapiclient.NewCommitData("Commit_example", "Subject_example", *openapiclient.NewCommitDataAuthor("Name_example", "Email_example", time.Now()), *openapiclient.NewCommitDataAuthor("Name_example", "Email_example", time.Now()))}) // ComponentVersionCommits | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsApi.ComponentVersionPushCommits(context.Background(), organization, project, component, version).ComponentVersionCommits(componentVersionCommits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentVersionPushCommits``: %v\n", err)
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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionQuery

> ComponentVersionPaginator ComponentVersionQuery(ctx, organization, project, component).Execute()

List all versions of a component

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
    organization := "my-organization" // string | Organization Key
    project := "my-project" // string | Project Key
    component := "component_example" // string | Component key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsApi.ComponentVersionQuery(context.Background(), organization, project, component).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentVersionQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentVersionQuery`: ComponentVersionPaginator
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentVersionQuery`: %v\n", resp)
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




### Return type

[**ComponentVersionPaginator**](ComponentVersionPaginator.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

