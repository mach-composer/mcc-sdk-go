# \ComponentsApi

All URIs are relative to *https://api.mach.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentCreate**](ComponentsApi.md#ComponentCreate) | **Post** /{organization}/projects/{project}/components | Create component
[**ComponentQuery**](ComponentsApi.md#ComponentQuery) | **Get** /{organization}/projects/{project}/components | List all components
[**ComponentVersionCreate**](ComponentsApi.md#ComponentVersionCreate) | **Post** /{organization}/projects/{project}/components/{component}/versions | Create component
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
    componentVersionDraft := *openapiclient.NewComponentVersionDraft("Component_example", "Version_example") // ComponentVersionDraft | 

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

