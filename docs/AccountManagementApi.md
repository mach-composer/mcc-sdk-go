# \AccountManagementApi

All URIs are relative to *https://api.mach.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MyAccountInformation**](AccountManagementApi.md#MyAccountInformation) | **Get** /account/me | Return user information from current user
[**OrganizationCreate**](AccountManagementApi.md#OrganizationCreate) | **Post** /account/organizations | Create new organization
[**OrganizationQuery**](AccountManagementApi.md#OrganizationQuery) | **Get** /account/organizations | List all organizations
[**OrganizationUserInvite**](AccountManagementApi.md#OrganizationUserInvite) | **Post** /account/organizations/{organization}/users | Add user to an organization
[**OrganizationUserQuery**](AccountManagementApi.md#OrganizationUserQuery) | **Get** /account/organizations/{organization}/users | List all users in an organization
[**ProjectCreate**](AccountManagementApi.md#ProjectCreate) | **Post** /account/organizations/{organization}/projects | Create new project in an organization
[**ProjectQuery**](AccountManagementApi.md#ProjectQuery) | **Get** /account/organizations/{organization}/projects | List all projects in an organization



## MyAccountInformation

> MyAccountInformation200Response MyAccountInformation(ctx).Execute()

Return user information from current user

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
    resp, r, err := apiClient.AccountManagementApi.MyAccountInformation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.MyAccountInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MyAccountInformation`: MyAccountInformation200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.MyAccountInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMyAccountInformationRequest struct via the builder pattern


### Return type

[**MyAccountInformation200Response**](MyAccountInformation200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationCreate

> Organization OrganizationCreate(ctx).OrganizationDraft(organizationDraft).Execute()

Create new organization

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
    organizationDraft := *openapiclient.NewOrganizationDraft("Key_example", "Name_example") // OrganizationDraft | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountManagementApi.OrganizationCreate(context.Background()).OrganizationDraft(organizationDraft).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.OrganizationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationCreate`: Organization
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.OrganizationCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationDraft** | [**OrganizationDraft**](OrganizationDraft.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationQuery

> OrganizationPaginator OrganizationQuery(ctx).Execute()

List all organizations

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
    resp, r, err := apiClient.AccountManagementApi.OrganizationQuery(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.OrganizationQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationQuery`: OrganizationPaginator
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.OrganizationQuery`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationQueryRequest struct via the builder pattern


### Return type

[**OrganizationPaginator**](OrganizationPaginator.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationUserInvite

> OrganizationUserInvite OrganizationUserInvite(ctx, organization).OrganizationUserInviteDraft(organizationUserInviteDraft).Execute()

Add user to an organization

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
    organizationUserInviteDraft := *openapiclient.NewOrganizationUserInviteDraft("Email_example") // OrganizationUserInviteDraft | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountManagementApi.OrganizationUserInvite(context.Background(), organization).OrganizationUserInviteDraft(organizationUserInviteDraft).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.OrganizationUserInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationUserInvite`: OrganizationUserInvite
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.OrganizationUserInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationUserInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationUserInviteDraft** | [**OrganizationUserInviteDraft**](OrganizationUserInviteDraft.md) |  | 

### Return type

[**OrganizationUserInvite**](OrganizationUserInvite.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationUserQuery

> OrganizationUserPaginator OrganizationUserQuery(ctx, organization).Execute()

List all users in an organization

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountManagementApi.OrganizationUserQuery(context.Background(), organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.OrganizationUserQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationUserQuery`: OrganizationUserPaginator
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.OrganizationUserQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationUserQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationUserPaginator**](OrganizationUserPaginator.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectCreate

> Project ProjectCreate(ctx, organization).ProjectDraft(projectDraft).Execute()

Create new project in an organization

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
    projectDraft := *openapiclient.NewProjectDraft("Key_example", "Name_example") // ProjectDraft | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountManagementApi.ProjectCreate(context.Background(), organization).ProjectDraft(projectDraft).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.ProjectCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectCreate`: Project
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.ProjectCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectDraft** | [**ProjectDraft**](ProjectDraft.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectQuery

> ProjectPaginator ProjectQuery(ctx, organization).Execute()

List all projects in an organization

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountManagementApi.ProjectQuery(context.Background(), organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.ProjectQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectQuery`: ProjectPaginator
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.ProjectQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectPaginator**](ProjectPaginator.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

