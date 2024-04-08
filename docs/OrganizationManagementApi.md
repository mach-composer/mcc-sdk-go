# \OrganizationManagementApi

All URIs are relative to *https://api.mach.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationCreate**](OrganizationManagementApi.md#OrganizationCreate) | **Post** /organizations | Create new organization
[**OrganizationPatch**](OrganizationManagementApi.md#OrganizationPatch) | **Patch** /organizations/{organization} | Update an organization
[**OrganizationQuery**](OrganizationManagementApi.md#OrganizationQuery) | **Get** /organizations | List all organizations
[**OrganizationUserInvite**](OrganizationManagementApi.md#OrganizationUserInvite) | **Post** /organizations/{organization}/users/invite | Invite a user to the organization
[**OrganizationUserInviteAccept**](OrganizationManagementApi.md#OrganizationUserInviteAccept) | **Post** /organizations/{organization}/users/invite/{id} | Accept a user invite
[**OrganizationUserInviteView**](OrganizationManagementApi.md#OrganizationUserInviteView) | **Get** /organizations/{organization}/users/invite/{id} | View invite information
[**OrganizationUserQuery**](OrganizationManagementApi.md#OrganizationUserQuery) | **Get** /organizations/{organization}/users | List all users in an organization
[**ProjectCreate**](OrganizationManagementApi.md#ProjectCreate) | **Post** /organizations/{organization}/projects | Create new project in an organization
[**ProjectDelete**](OrganizationManagementApi.md#ProjectDelete) | **Delete** /organizations/{organization}/projects/{project} | Delete a project
[**ProjectPatch**](OrganizationManagementApi.md#ProjectPatch) | **Patch** /organizations/{organization}/projects/{project} | Update a Project
[**ProjectQuery**](OrganizationManagementApi.md#ProjectQuery) | **Get** /organizations/{organization}/projects | List all projects in an organization



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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationDraft := *openapiclient.NewOrganizationDraft("Key_example", "Name_example") // OrganizationDraft | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationCreate(context.Background()).OrganizationDraft(organizationDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationCreate`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.OrganizationCreate`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationPatch

> Organization OrganizationPatch(ctx, organization).PatchRequestInner(patchRequestInner).Execute()

Update an organization

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
	patchRequestInner := []openapiclient.PatchRequestInner{openapiclient.PatchRequest_inner{JSONPatchRequestAddReplaceTest: openapiclient.NewJSONPatchRequestAddReplaceTest("Path_example", interface{}(123), "Op_example")}} // []PatchRequestInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationPatch(context.Background(), organization).PatchRequestInner(patchRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationPatch`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.OrganizationPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchRequestInner** | [**[]PatchRequestInner**](PatchRequestInner.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationQuery(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationQuery`: OrganizationPaginator
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.OrganizationQuery`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationQueryRequest struct via the builder pattern


### Return type

[**OrganizationPaginator**](OrganizationPaginator.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationUserInvite

> OrganizationUserInvite OrganizationUserInvite(ctx, organization).OrganizationUserInviteDraft(organizationUserInviteDraft).Execute()

Invite a user to the organization

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
	organizationUserInviteDraft := *openapiclient.NewOrganizationUserInviteDraft("Email_example") // OrganizationUserInviteDraft | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationUserInvite(context.Background(), organization).OrganizationUserInviteDraft(organizationUserInviteDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationUserInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationUserInvite`: OrganizationUserInvite
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.OrganizationUserInvite`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationUserInviteAccept

> OrganizationUserInviteData OrganizationUserInviteAccept(ctx, organization, id).Execute()

Accept a user invite

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Invite ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationUserInviteAccept(context.Background(), organization, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationUserInviteAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationUserInviteAccept`: OrganizationUserInviteData
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.OrganizationUserInviteAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**id** | **string** | Invite ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationUserInviteAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationUserInviteData**](OrganizationUserInviteData.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationUserInviteView

> OrganizationUserInviteData OrganizationUserInviteView(ctx, organization, id).Execute()

View invite information

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Invite ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationUserInviteView(context.Background(), organization, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationUserInviteView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationUserInviteView`: OrganizationUserInviteData
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.OrganizationUserInviteView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**id** | **string** | Invite ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationUserInviteViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationUserInviteData**](OrganizationUserInviteData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "my-organization" // string | Organization Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationUserQuery(context.Background(), organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationUserQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationUserQuery`: OrganizationUserPaginator
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.OrganizationUserQuery`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "my-organization" // string | Organization Key
	projectDraft := *openapiclient.NewProjectDraft("Key_example", "Name_example") // ProjectDraft | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.ProjectCreate(context.Background(), organization).ProjectDraft(projectDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.ProjectCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectCreate`: Project
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.ProjectCreate`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectDelete

> Project ProjectDelete(ctx, organization, project).Execute()

Delete a project



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.ProjectDelete(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.ProjectDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectDelete`: Project
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.ProjectDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Project**](Project.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectPatch

> Project ProjectPatch(ctx, organization, project).PatchRequestInner(patchRequestInner).Execute()

Update a Project

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
	patchRequestInner := []openapiclient.PatchRequestInner{openapiclient.PatchRequest_inner{JSONPatchRequestAddReplaceTest: openapiclient.NewJSONPatchRequestAddReplaceTest("Path_example", interface{}(123), "Op_example")}} // []PatchRequestInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.ProjectPatch(context.Background(), organization, project).PatchRequestInner(patchRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.ProjectPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectPatch`: Project
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.ProjectPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization Key | 
**project** | **string** | Project Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchRequestInner** | [**[]PatchRequestInner**](PatchRequestInner.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "my-organization" // string | Organization Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.ProjectQuery(context.Background(), organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.ProjectQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectQuery`: ProjectPaginator
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.ProjectQuery`: %v\n", resp)
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

[OAuth2](../README.md#OAuth2), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

