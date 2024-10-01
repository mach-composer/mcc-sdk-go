# \OrganizationManagementApi

All URIs are relative to *https://api.mach.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationCreate**](OrganizationManagementApi.md#OrganizationCreate) | **Post** /organizations | Create new organization
[**OrganizationDelete**](OrganizationManagementApi.md#OrganizationDelete) | **Delete** /organizations/{organization} | Delete an organization
[**OrganizationGet**](OrganizationManagementApi.md#OrganizationGet) | **Get** /organizations/{organization} | Get organization details
[**OrganizationPatch**](OrganizationManagementApi.md#OrganizationPatch) | **Patch** /organizations/{organization} | Update an organization
[**OrganizationQuery**](OrganizationManagementApi.md#OrganizationQuery) | **Get** /organizations | List all organizations
[**OrganizationUpdate**](OrganizationManagementApi.md#OrganizationUpdate) | **Put** /organizations/{organization} | Update an organization
[**OrganizationUserInvite**](OrganizationManagementApi.md#OrganizationUserInvite) | **Post** /organizations/{organization}/users/invite | Invite a user to the organization
[**OrganizationUserInviteDelete**](OrganizationManagementApi.md#OrganizationUserInviteDelete) | **Delete** /organizations/{organization}/users/invite/{id} | Cancel an invite
[**OrganizationUserInvitePatch**](OrganizationManagementApi.md#OrganizationUserInvitePatch) | **Patch** /organizations/{organization}/users/invite/{id} | Accept or reject an invite
[**OrganizationUserInviteUpdate**](OrganizationManagementApi.md#OrganizationUserInviteUpdate) | **Put** /organizations/{organization}/users/invite/{id} | Accept or reject an invite
[**OrganizationUserInviteView**](OrganizationManagementApi.md#OrganizationUserInviteView) | **Get** /organizations/{organization}/users/invite/{id} | View invite information
[**OrganizationUserQuery**](OrganizationManagementApi.md#OrganizationUserQuery) | **Get** /organizations/{organization}/users/invite | List all users in an organization
[**ProjectCreate**](OrganizationManagementApi.md#ProjectCreate) | **Post** /organizations/{organization}/projects | Create new project in an organization
[**ProjectDelete**](OrganizationManagementApi.md#ProjectDelete) | **Delete** /organizations/{organization}/projects/{project} | Delete a project
[**ProjectGet**](OrganizationManagementApi.md#ProjectGet) | **Get** /organizations/{organization}/projects/{project} | Get project details
[**ProjectPatch**](OrganizationManagementApi.md#ProjectPatch) | **Patch** /organizations/{organization}/projects/{project} | Update a project
[**ProjectQuery**](OrganizationManagementApi.md#ProjectQuery) | **Get** /organizations/{organization}/projects | List all projects in an organization
[**ProjectUpdate**](OrganizationManagementApi.md#ProjectUpdate) | **Put** /organizations/{organization}/projects/{project} | Update a project



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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationDelete

> OrganizationDelete(ctx, organization).Execute()

Delete an organization

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationManagementApi.OrganizationDelete(context.Background(), organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationDeleteRequest struct via the builder pattern


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


## OrganizationGet

> Organization OrganizationGet(ctx, organization).Execute()

Get organization details

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationGet(context.Background(), organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationGet`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.OrganizationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Organization**](Organization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationPatch

> Organization OrganizationPatch(ctx, organization).PatchedOrganizationDraft(patchedOrganizationDraft).Execute()

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
	organization := "organization_example" // string | 
	patchedOrganizationDraft := *openapiclient.NewPatchedOrganizationDraft() // PatchedOrganizationDraft |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationPatch(context.Background(), organization).PatchedOrganizationDraft(patchedOrganizationDraft).Execute()
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
**organization** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedOrganizationDraft** | [**PatchedOrganizationDraft**](PatchedOrganizationDraft.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationQuery

> OrganizationPaginator OrganizationQuery(ctx).Limit(limit).Offset(offset).Execute()

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
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationQuery(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationQuery`: OrganizationPaginator
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.OrganizationQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**OrganizationPaginator**](OrganizationPaginator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationUpdate

> Organization OrganizationUpdate(ctx, organization).OrganizationDraft(organizationDraft).Execute()

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
	organization := "organization_example" // string | 
	organizationDraft := *openapiclient.NewOrganizationDraft("Key_example", "Name_example") // OrganizationDraft | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationUpdate(context.Background(), organization).OrganizationDraft(organizationDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationUpdate`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.OrganizationUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationDraft** | [**OrganizationDraft**](OrganizationDraft.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationUserInvite

> OrganizationUser OrganizationUserInvite(ctx, organization).UserInviteDraft(userInviteDraft).Execute()

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
	organization := "organization_example" // string | 
	userInviteDraft := *openapiclient.NewUserInviteDraft("CreatedBy_example", "Email_example", "Organization_example", "Role_example") // UserInviteDraft | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationUserInvite(context.Background(), organization).UserInviteDraft(userInviteDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationUserInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationUserInvite`: OrganizationUser
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.OrganizationUserInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationUserInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userInviteDraft** | [**UserInviteDraft**](UserInviteDraft.md) |  | 

### Return type

[**OrganizationUser**](OrganizationUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationUserInviteDelete

> OrganizationUserInviteDelete(ctx, organization, id).Execute()

Cancel an invite

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationManagementApi.OrganizationUserInviteDelete(context.Background(), organization, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationUserInviteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationUserInviteDeleteRequest struct via the builder pattern


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


## OrganizationUserInvitePatch

> OrganizationUser OrganizationUserInvitePatch(ctx, organization, id).PatchedUserInviteDraft(patchedUserInviteDraft).Execute()

Accept or reject an invite

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
	id := "id_example" // string | 
	patchedUserInviteDraft := *openapiclient.NewPatchedUserInviteDraft() // PatchedUserInviteDraft |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationUserInvitePatch(context.Background(), organization, id).PatchedUserInviteDraft(patchedUserInviteDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationUserInvitePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationUserInvitePatch`: OrganizationUser
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.OrganizationUserInvitePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationUserInvitePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedUserInviteDraft** | [**PatchedUserInviteDraft**](PatchedUserInviteDraft.md) |  | 

### Return type

[**OrganizationUser**](OrganizationUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationUserInviteUpdate

> OrganizationUser OrganizationUserInviteUpdate(ctx, organization, id).UserInviteDraft(userInviteDraft).Execute()

Accept or reject an invite

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
	id := "id_example" // string | 
	userInviteDraft := *openapiclient.NewUserInviteDraft("CreatedBy_example", "Email_example", "Organization_example", "Role_example") // UserInviteDraft | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationUserInviteUpdate(context.Background(), organization, id).UserInviteDraft(userInviteDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationUserInviteUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationUserInviteUpdate`: OrganizationUser
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.OrganizationUserInviteUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationUserInviteUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userInviteDraft** | [**UserInviteDraft**](UserInviteDraft.md) |  | 

### Return type

[**OrganizationUser**](OrganizationUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationUserInviteView

> OrganizationUser OrganizationUserInviteView(ctx, organization, id).Execute()

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
	organization := "organization_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationUserInviteView(context.Background(), organization, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.OrganizationUserInviteView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationUserInviteView`: OrganizationUser
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.OrganizationUserInviteView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationUserInviteViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationUser**](OrganizationUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationUserQuery

> OrganizationUserPaginator OrganizationUserQuery(ctx, organization).Limit(limit).Offset(offset).Execute()

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
	organization := "organization_example" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.OrganizationUserQuery(context.Background(), organization).Limit(limit).Offset(offset).Execute()
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
**organization** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationUserQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**OrganizationUserPaginator**](OrganizationUserPaginator.md)

### Authorization

No authorization required

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
	organization := "organization_example" // string | 
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
**organization** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectDraft** | [**ProjectDraft**](ProjectDraft.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectDelete

> ProjectDelete(ctx, organization, project).Execute()

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationManagementApi.ProjectDelete(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.ProjectDelete``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiProjectDeleteRequest struct via the builder pattern


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


## ProjectGet

> Project ProjectGet(ctx, organization, project).Execute()

Get project details

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.ProjectGet(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.ProjectGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectGet`: Project
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.ProjectGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectPatch

> Project ProjectPatch(ctx, organization, project).PatchedProjectDraft(patchedProjectDraft).Execute()

Update a project

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
	patchedProjectDraft := *openapiclient.NewPatchedProjectDraft() // PatchedProjectDraft |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.ProjectPatch(context.Background(), organization, project).PatchedProjectDraft(patchedProjectDraft).Execute()
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
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedProjectDraft** | [**PatchedProjectDraft**](PatchedProjectDraft.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectQuery

> ProjectPaginator ProjectQuery(ctx, organization).Limit(limit).Offset(offset).Execute()

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
	organization := "organization_example" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.ProjectQuery(context.Background(), organization).Limit(limit).Offset(offset).Execute()
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
**organization** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**ProjectPaginator**](ProjectPaginator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectUpdate

> Project ProjectUpdate(ctx, organization, project).ProjectDraft(projectDraft).Execute()

Update a project

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
	projectDraft := *openapiclient.NewProjectDraft("Key_example", "Name_example") // ProjectDraft | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationManagementApi.ProjectUpdate(context.Background(), organization, project).ProjectDraft(projectDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationManagementApi.ProjectUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectUpdate`: Project
	fmt.Fprintf(os.Stdout, "Response from `OrganizationManagementApi.ProjectUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectDraft** | [**ProjectDraft**](ProjectDraft.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

