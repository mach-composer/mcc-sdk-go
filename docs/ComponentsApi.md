# \ComponentsApi

All URIs are relative to *https://api.mach.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentCommitQuery**](ComponentsApi.md#ComponentCommitQuery) | **Get** /organizations/{organization}/projects/{project}/components/{component}/commits | List all commits of a component between two versions ordered by creation date. If &#x60;to&#x60; is not provided, it will list all commits from &#x60;from&#x60; to the latest version.
[**ComponentCreate**](ComponentsApi.md#ComponentCreate) | **Post** /organizations/{organization}/projects/{project}/components | Create component
[**ComponentDelete**](ComponentsApi.md#ComponentDelete) | **Delete** /organizations/{organization}/projects/{project}/components/{component} | Delete a component
[**ComponentGet**](ComponentsApi.md#ComponentGet) | **Get** /organizations/{organization}/projects/{project}/components/{component} | Get component details
[**ComponentLatestVersion**](ComponentsApi.md#ComponentLatestVersion) | **Get** /organizations/{organization}/projects/{project}/components/{component}/latest | Get last component version
[**ComponentPatch**](ComponentsApi.md#ComponentPatch) | **Patch** /organizations/{organization}/projects/{project}/components/{component} | Patch an existing component
[**ComponentQuery**](ComponentsApi.md#ComponentQuery) | **Get** /organizations/{organization}/projects/{project}/components | List all components
[**ComponentUpdate**](ComponentsApi.md#ComponentUpdate) | **Put** /organizations/{organization}/projects/{project}/components/{component} | Update a component
[**ComponentVersionCreate**](ComponentsApi.md#ComponentVersionCreate) | **Post** /organizations/{organization}/projects/{project}/components/{component}/versions | Create component version
[**ComponentVersionDelete**](ComponentsApi.md#ComponentVersionDelete) | **Delete** /organizations/{organization}/projects/{project}/components/{component}/versions/{version} | Delete component version
[**ComponentVersionDeleteCommit**](ComponentsApi.md#ComponentVersionDeleteCommit) | **Delete** /organizations/{organization}/projects/{project}/components/{component}/versions/{version}/commits/{commit} | Delete commit
[**ComponentVersionGet**](ComponentsApi.md#ComponentVersionGet) | **Get** /organizations/{organization}/projects/{project}/components/{component}/versions/{version} | Get component version
[**ComponentVersionGetCommit**](ComponentsApi.md#ComponentVersionGetCommit) | **Get** /organizations/{organization}/projects/{project}/components/{component}/versions/{version}/commits/{commit} | Get commit details
[**ComponentVersionPatch**](ComponentsApi.md#ComponentVersionPatch) | **Patch** /organizations/{organization}/projects/{project}/components/{component}/versions/{version} | Patch component version
[**ComponentVersionPatchCommit**](ComponentsApi.md#ComponentVersionPatchCommit) | **Patch** /organizations/{organization}/projects/{project}/components/{component}/versions/{version}/commits/{commit} | Patch commit
[**ComponentVersionPushCommits**](ComponentsApi.md#ComponentVersionPushCommits) | **Post** /organizations/{organization}/projects/{project}/components/{component}/versions/{version}/commits | Push commits for this component version
[**ComponentVersionQuery**](ComponentsApi.md#ComponentVersionQuery) | **Get** /organizations/{organization}/projects/{project}/components/{component}/versions | List all versions of a component
[**ComponentVersionQueryCommits**](ComponentsApi.md#ComponentVersionQueryCommits) | **Get** /organizations/{organization}/projects/{project}/components/{component}/versions/{version}/commits | Get commits for this component version
[**ComponentVersionUpdate**](ComponentsApi.md#ComponentVersionUpdate) | **Put** /organizations/{organization}/projects/{project}/components/{component}/versions/{version} | Update component version
[**ComponentVersionUpdateCommit**](ComponentsApi.md#ComponentVersionUpdateCommit) | **Put** /organizations/{organization}/projects/{project}/components/{component}/versions/{version}/commits/{commit} | Update commit



## ComponentCommitQuery

> CommitDataPaginator ComponentCommitQuery(ctx, organization, project, component).From(from).To(to).Offset(offset).Limit(limit).Execute()

List all commits of a component between two versions ordered by creation date. If `to` is not provided, it will list all commits from `from` to the latest version.

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
	component := "component_example" // string | 
	from := "from_example" // string |  (optional)
	to := "to_example" // string |  (optional)
	offset := float32(8.14) // float32 |  (optional)
	limit := float32(8.14) // float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsApi.ComponentCommitQuery(context.Background(), organization, project, component).From(from).To(to).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentCommitQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComponentCommitQuery`: CommitDataPaginator
	fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentCommitQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**component** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentCommitQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **from** | **string** |  | 
 **to** | **string** |  | 
 **offset** | **float32** |  | 
 **limit** | **float32** |  | 

### Return type

[**CommitDataPaginator**](CommitDataPaginator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	componentDraft := *openapiclient.NewComponentDraft("Key_example", "Name_example") // ComponentDraft | 

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
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **componentDraft** | [**ComponentDraft**](ComponentDraft.md) |  | 

### Return type

[**Component**](Component.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentDelete

> ComponentDelete(ctx, organization, project, component).Execute()

Delete a component



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
	component := "component_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComponentsApi.ComponentDelete(context.Background(), organization, project, component).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentDelete``: %v\n", err)
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
**component** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentDeleteRequest struct via the builder pattern


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


## ComponentGet

> Component ComponentGet(ctx, organization, project, component).Execute()

Get component details

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
	component := "component_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsApi.ComponentGet(context.Background(), organization, project, component).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComponentGet`: Component
	fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**component** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Component**](Component.md)

### Authorization

No authorization required

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	component := "component_example" // string | 
	branch := "branch_example" // string |  (optional)

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
**organization** | **string** |  | 
**project** | **string** |  | 
**component** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentLatestVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **branch** | **string** |  | 

### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentPatch

> Component ComponentPatch(ctx, organization, project, component).PatchedComponentDraft(patchedComponentDraft).Execute()

Patch an existing component

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
	component := "component_example" // string | 
	patchedComponentDraft := *openapiclient.NewPatchedComponentDraft() // PatchedComponentDraft |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsApi.ComponentPatch(context.Background(), organization, project, component).PatchedComponentDraft(patchedComponentDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComponentPatch`: Component
	fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**component** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchedComponentDraft** | [**PatchedComponentDraft**](PatchedComponentDraft.md) |  | 

### Return type

[**Component**](Component.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentQuery

> ComponentPaginator ComponentQuery(ctx, organization, project).Key(key).Limit(limit).Offset(offset).Execute()

List all components

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
	key := "key_example" // string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsApi.ComponentQuery(context.Background(), organization, project).Key(key).Limit(limit).Offset(offset).Execute()
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
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**ComponentPaginator**](ComponentPaginator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentUpdate

> Component ComponentUpdate(ctx, organization, project, component).ComponentDraft(componentDraft).Execute()

Update a component

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
	component := "component_example" // string | 
	componentDraft := *openapiclient.NewComponentDraft("Key_example", "Name_example") // ComponentDraft | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsApi.ComponentUpdate(context.Background(), organization, project, component).ComponentDraft(componentDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComponentUpdate`: Component
	fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**component** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **componentDraft** | [**ComponentDraft**](ComponentDraft.md) |  | 

### Return type

[**Component**](Component.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	component := "component_example" // string | 
	componentVersionDraft := *openapiclient.NewComponentVersionDraft("Version_example") // ComponentVersionDraft | 

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
**organization** | **string** |  | 
**project** | **string** |  | 
**component** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **componentVersionDraft** | [**ComponentVersionDraft**](ComponentVersionDraft.md) |  | 

### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionDelete

> ComponentVersionDelete(ctx, organization, project, component, version).Execute()

Delete component version

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
	component := "component_example" // string | 
	version := "version_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComponentsApi.ComponentVersionDelete(context.Background(), organization, project, component, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentVersionDelete``: %v\n", err)
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
**component** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionDeleteRequest struct via the builder pattern


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


## ComponentVersionDeleteCommit

> ComponentVersionDeleteCommit(ctx, organization, project, component, version, commit).Execute()

Delete commit

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
	component := "component_example" // string | 
	version := "version_example" // string | 
	commit := "commit_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComponentsApi.ComponentVersionDeleteCommit(context.Background(), organization, project, component, version, commit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentVersionDeleteCommit``: %v\n", err)
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
**component** | **string** |  | 
**version** | **string** |  | 
**commit** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionDeleteCommitRequest struct via the builder pattern


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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	component := "component_example" // string | 
	version := "version_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsApi.ComponentVersionGet(context.Background(), organization, project, component, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentVersionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComponentVersionGet`: ComponentVersion
	fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentVersionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**component** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionGetCommit

> CommitData ComponentVersionGetCommit(ctx, organization, project, component, version, commit).Execute()

Get commit details

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
	component := "component_example" // string | 
	version := "version_example" // string | 
	commit := "commit_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsApi.ComponentVersionGetCommit(context.Background(), organization, project, component, version, commit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentVersionGetCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComponentVersionGetCommit`: CommitData
	fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentVersionGetCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**component** | **string** |  | 
**version** | **string** |  | 
**commit** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionGetCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**CommitData**](CommitData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionPatch

> ComponentVersion ComponentVersionPatch(ctx, organization, project, component, version).PatchedComponentVersionDraft(patchedComponentVersionDraft).Execute()

Patch component version

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
	component := "component_example" // string | 
	version := "version_example" // string | 
	patchedComponentVersionDraft := *openapiclient.NewPatchedComponentVersionDraft() // PatchedComponentVersionDraft |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsApi.ComponentVersionPatch(context.Background(), organization, project, component, version).PatchedComponentVersionDraft(patchedComponentVersionDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentVersionPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComponentVersionPatch`: ComponentVersion
	fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentVersionPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**component** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **patchedComponentVersionDraft** | [**PatchedComponentVersionDraft**](PatchedComponentVersionDraft.md) |  | 

### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionPatchCommit

> CommitData ComponentVersionPatchCommit(ctx, organization, project, component, version, commit).PatchedCommitDataDraft(patchedCommitDataDraft).Execute()

Patch commit

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
	component := "component_example" // string | 
	version := "version_example" // string | 
	commit := "commit_example" // string | 
	patchedCommitDataDraft := *openapiclient.NewPatchedCommitDataDraft() // PatchedCommitDataDraft |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsApi.ComponentVersionPatchCommit(context.Background(), organization, project, component, version, commit).PatchedCommitDataDraft(patchedCommitDataDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentVersionPatchCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComponentVersionPatchCommit`: CommitData
	fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentVersionPatchCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**component** | **string** |  | 
**version** | **string** |  | 
**commit** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionPatchCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **patchedCommitDataDraft** | [**PatchedCommitDataDraft**](PatchedCommitDataDraft.md) |  | 

### Return type

[**CommitData**](CommitData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionPushCommits

> CommitDataPaginator ComponentVersionPushCommits(ctx, organization, project, component, version).ComponentCommitCreateDraft(componentCommitCreateDraft).Execute()

Push commits for this component version



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
	component := "component_example" // string | 
	version := "version_example" // string | 
	componentCommitCreateDraft := *openapiclient.NewComponentCommitCreateDraft() // ComponentCommitCreateDraft |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsApi.ComponentVersionPushCommits(context.Background(), organization, project, component, version).ComponentCommitCreateDraft(componentCommitCreateDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentVersionPushCommits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComponentVersionPushCommits`: CommitDataPaginator
	fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentVersionPushCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**component** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionPushCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **componentCommitCreateDraft** | [**ComponentCommitCreateDraft**](ComponentCommitCreateDraft.md) |  | 

### Return type

[**CommitDataPaginator**](CommitDataPaginator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionQuery

> ComponentVersionPaginator ComponentVersionQuery(ctx, organization, project, component).Limit(limit).Offset(offset).Execute()

List all versions of a component

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
	component := "component_example" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsApi.ComponentVersionQuery(context.Background(), organization, project, component).Limit(limit).Offset(offset).Execute()
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
**organization** | **string** |  | 
**project** | **string** |  | 
**component** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**ComponentVersionPaginator**](ComponentVersionPaginator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionQueryCommits

> CommitDataPaginator ComponentVersionQueryCommits(ctx, organization, project, component, version).Limit(limit).Offset(offset).Execute()

Get commits for this component version

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
	component := "component_example" // string | 
	version := "version_example" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsApi.ComponentVersionQueryCommits(context.Background(), organization, project, component, version).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentVersionQueryCommits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComponentVersionQueryCommits`: CommitDataPaginator
	fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentVersionQueryCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**component** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionQueryCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**CommitDataPaginator**](CommitDataPaginator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionUpdate

> ComponentVersion ComponentVersionUpdate(ctx, organization, project, component, version).ComponentVersionDraft(componentVersionDraft).Execute()

Update component version

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
	component := "component_example" // string | 
	version := "version_example" // string | 
	componentVersionDraft := *openapiclient.NewComponentVersionDraft("Version_example") // ComponentVersionDraft | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsApi.ComponentVersionUpdate(context.Background(), organization, project, component, version).ComponentVersionDraft(componentVersionDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentVersionUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComponentVersionUpdate`: ComponentVersion
	fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentVersionUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**component** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **componentVersionDraft** | [**ComponentVersionDraft**](ComponentVersionDraft.md) |  | 

### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentVersionUpdateCommit

> CommitData ComponentVersionUpdateCommit(ctx, organization, project, component, version, commit).CommitDataDraft(commitDataDraft).Execute()

Update commit

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	component := "component_example" // string | 
	version := "version_example" // string | 
	commit := "commit_example" // string | 
	commitDataDraft := *openapiclient.NewCommitDataDraft("Subject_example", "Commit_example", *openapiclient.NewCommitDataAuthorDraft("Name_example", "Email_example", time.Now()), *openapiclient.NewCommitDataAuthorDraft("Name_example", "Email_example", time.Now())) // CommitDataDraft | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsApi.ComponentVersionUpdateCommit(context.Background(), organization, project, component, version, commit).CommitDataDraft(commitDataDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.ComponentVersionUpdateCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComponentVersionUpdateCommit`: CommitData
	fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.ComponentVersionUpdateCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**component** | **string** |  | 
**version** | **string** |  | 
**commit** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentVersionUpdateCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **commitDataDraft** | [**CommitDataDraft**](CommitDataDraft.md) |  | 

### Return type

[**CommitData**](CommitData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

