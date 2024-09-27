# \MyAccountApi

All URIs are relative to *https://api.mach.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MyAccountInformation**](MyAccountApi.md#MyAccountInformation) | **Get** /account/me | Get user information
[**UserCreate**](MyAccountApi.md#UserCreate) | **Post** /account/me | Create new user
[**UserDelete**](MyAccountApi.md#UserDelete) | **Delete** /account/me/{id} | Delete user
[**UserGet**](MyAccountApi.md#UserGet) | **Get** /account/me/{id} | Get user information
[**UserPatch**](MyAccountApi.md#UserPatch) | **Patch** /account/me/{id} | Update user information
[**UserUpdate**](MyAccountApi.md#UserUpdate) | **Put** /account/me/{id} | Update user information



## MyAccountInformation

> OrganizationUserPaginator MyAccountInformation(ctx).Limit(limit).Offset(offset).Execute()

Get user information



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
	resp, r, err := apiClient.MyAccountApi.MyAccountInformation(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountApi.MyAccountInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MyAccountInformation`: OrganizationUserPaginator
	fmt.Fprintf(os.Stdout, "Response from `MyAccountApi.MyAccountInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMyAccountInformationRequest struct via the builder pattern


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


## UserCreate

> OrganizationUser UserCreate(ctx).OrganizationUserDraft(organizationUserDraft).Execute()

Create new user

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
	organizationUserDraft := *openapiclient.NewOrganizationUserDraft("AccountStatus_example", "Name_example", "Email_example") // OrganizationUserDraft | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccountApi.UserCreate(context.Background()).OrganizationUserDraft(organizationUserDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountApi.UserCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserCreate`: OrganizationUser
	fmt.Fprintf(os.Stdout, "Response from `MyAccountApi.UserCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationUserDraft** | [**OrganizationUserDraft**](OrganizationUserDraft.md) |  | 

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


## UserDelete

> UserDelete(ctx, id).Execute()

Delete user

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyAccountApi.UserDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountApi.UserDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserDeleteRequest struct via the builder pattern


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


## UserGet

> OrganizationUser UserGet(ctx, id).Execute()

Get user information

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccountApi.UserGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountApi.UserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserGet`: OrganizationUser
	fmt.Fprintf(os.Stdout, "Response from `MyAccountApi.UserGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetRequest struct via the builder pattern


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


## UserPatch

> OrganizationUser UserPatch(ctx, id).PatchedOrganizationUserDraft(patchedOrganizationUserDraft).Execute()

Update user information

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
	id := "id_example" // string | 
	patchedOrganizationUserDraft := *openapiclient.NewPatchedOrganizationUserDraft() // PatchedOrganizationUserDraft |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccountApi.UserPatch(context.Background(), id).PatchedOrganizationUserDraft(patchedOrganizationUserDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountApi.UserPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserPatch`: OrganizationUser
	fmt.Fprintf(os.Stdout, "Response from `MyAccountApi.UserPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedOrganizationUserDraft** | [**PatchedOrganizationUserDraft**](PatchedOrganizationUserDraft.md) |  | 

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


## UserUpdate

> OrganizationUser UserUpdate(ctx, id).OrganizationUserDraft(organizationUserDraft).Execute()

Update user information

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
	id := "id_example" // string | 
	organizationUserDraft := *openapiclient.NewOrganizationUserDraft("AccountStatus_example", "Name_example", "Email_example") // OrganizationUserDraft | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccountApi.UserUpdate(context.Background(), id).OrganizationUserDraft(organizationUserDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountApi.UserUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUpdate`: OrganizationUser
	fmt.Fprintf(os.Stdout, "Response from `MyAccountApi.UserUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationUserDraft** | [**OrganizationUserDraft**](OrganizationUserDraft.md) |  | 

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

