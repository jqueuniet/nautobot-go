# \UsersApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersConfigRetrieve**](UsersApi.md#UsersConfigRetrieve) | **Get** /users/config/ | 
[**UsersGroupsBulkDestroy**](UsersApi.md#UsersGroupsBulkDestroy) | **Delete** /users/groups/ | 
[**UsersGroupsBulkPartialUpdate**](UsersApi.md#UsersGroupsBulkPartialUpdate) | **Patch** /users/groups/ | 
[**UsersGroupsBulkUpdate**](UsersApi.md#UsersGroupsBulkUpdate) | **Put** /users/groups/ | 
[**UsersGroupsCreate**](UsersApi.md#UsersGroupsCreate) | **Post** /users/groups/ | 
[**UsersGroupsDestroy**](UsersApi.md#UsersGroupsDestroy) | **Delete** /users/groups/{id}/ | 
[**UsersGroupsList**](UsersApi.md#UsersGroupsList) | **Get** /users/groups/ | 
[**UsersGroupsPartialUpdate**](UsersApi.md#UsersGroupsPartialUpdate) | **Patch** /users/groups/{id}/ | 
[**UsersGroupsRetrieve**](UsersApi.md#UsersGroupsRetrieve) | **Get** /users/groups/{id}/ | 
[**UsersGroupsUpdate**](UsersApi.md#UsersGroupsUpdate) | **Put** /users/groups/{id}/ | 
[**UsersPermissionsBulkDestroy**](UsersApi.md#UsersPermissionsBulkDestroy) | **Delete** /users/permissions/ | 
[**UsersPermissionsBulkPartialUpdate**](UsersApi.md#UsersPermissionsBulkPartialUpdate) | **Patch** /users/permissions/ | 
[**UsersPermissionsBulkUpdate**](UsersApi.md#UsersPermissionsBulkUpdate) | **Put** /users/permissions/ | 
[**UsersPermissionsCreate**](UsersApi.md#UsersPermissionsCreate) | **Post** /users/permissions/ | 
[**UsersPermissionsDestroy**](UsersApi.md#UsersPermissionsDestroy) | **Delete** /users/permissions/{id}/ | 
[**UsersPermissionsList**](UsersApi.md#UsersPermissionsList) | **Get** /users/permissions/ | 
[**UsersPermissionsPartialUpdate**](UsersApi.md#UsersPermissionsPartialUpdate) | **Patch** /users/permissions/{id}/ | 
[**UsersPermissionsRetrieve**](UsersApi.md#UsersPermissionsRetrieve) | **Get** /users/permissions/{id}/ | 
[**UsersPermissionsUpdate**](UsersApi.md#UsersPermissionsUpdate) | **Put** /users/permissions/{id}/ | 
[**UsersTokensBulkDestroy**](UsersApi.md#UsersTokensBulkDestroy) | **Delete** /users/tokens/ | 
[**UsersTokensBulkPartialUpdate**](UsersApi.md#UsersTokensBulkPartialUpdate) | **Patch** /users/tokens/ | 
[**UsersTokensBulkUpdate**](UsersApi.md#UsersTokensBulkUpdate) | **Put** /users/tokens/ | 
[**UsersTokensCreate**](UsersApi.md#UsersTokensCreate) | **Post** /users/tokens/ | 
[**UsersTokensDestroy**](UsersApi.md#UsersTokensDestroy) | **Delete** /users/tokens/{id}/ | 
[**UsersTokensList**](UsersApi.md#UsersTokensList) | **Get** /users/tokens/ | 
[**UsersTokensPartialUpdate**](UsersApi.md#UsersTokensPartialUpdate) | **Patch** /users/tokens/{id}/ | 
[**UsersTokensRetrieve**](UsersApi.md#UsersTokensRetrieve) | **Get** /users/tokens/{id}/ | 
[**UsersTokensUpdate**](UsersApi.md#UsersTokensUpdate) | **Put** /users/tokens/{id}/ | 
[**UsersUsersBulkDestroy**](UsersApi.md#UsersUsersBulkDestroy) | **Delete** /users/users/ | 
[**UsersUsersBulkPartialUpdate**](UsersApi.md#UsersUsersBulkPartialUpdate) | **Patch** /users/users/ | 
[**UsersUsersBulkUpdate**](UsersApi.md#UsersUsersBulkUpdate) | **Put** /users/users/ | 
[**UsersUsersCreate**](UsersApi.md#UsersUsersCreate) | **Post** /users/users/ | 
[**UsersUsersDestroy**](UsersApi.md#UsersUsersDestroy) | **Delete** /users/users/{id}/ | 
[**UsersUsersList**](UsersApi.md#UsersUsersList) | **Get** /users/users/ | 
[**UsersUsersPartialUpdate**](UsersApi.md#UsersUsersPartialUpdate) | **Patch** /users/users/{id}/ | 
[**UsersUsersRetrieve**](UsersApi.md#UsersUsersRetrieve) | **Get** /users/users/{id}/ | 
[**UsersUsersUpdate**](UsersApi.md#UsersUsersUpdate) | **Put** /users/users/{id}/ | 



## UsersConfigRetrieve

> map[string]interface{} UsersConfigRetrieve(ctx).Execute()





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
    resp, r, err := apiClient.UsersApi.UsersConfigRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersConfigRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersConfigRetrieve`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersConfigRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersConfigRetrieveRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsBulkDestroy

> UsersGroupsBulkDestroy(ctx).Execute()





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
    resp, r, err := apiClient.UsersApi.UsersGroupsBulkDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsBulkDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsBulkDestroyRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsBulkPartialUpdate

> Group UsersGroupsBulkPartialUpdate(ctx).PatchedGroup(patchedGroup).Execute()





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
    patchedGroup := *openapiclient.NewPatchedGroup() // PatchedGroup |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsBulkPartialUpdate(context.Background()).PatchedGroup(patchedGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsBulkPartialUpdate`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedGroup** | [**PatchedGroup**](PatchedGroup.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsBulkUpdate

> Group UsersGroupsBulkUpdate(ctx).Group(group).Execute()





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
    group := *openapiclient.NewGroup(int32(123), "Url_example", "Name_example", int32(123), "Display_example") // Group | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsBulkUpdate(context.Background()).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsBulkUpdate`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | [**Group**](Group.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsCreate

> Group UsersGroupsCreate(ctx).Group(group).Execute()





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
    group := *openapiclient.NewGroup(int32(123), "Url_example", "Name_example", int32(123), "Display_example") // Group | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsCreate(context.Background()).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsCreate`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | [**Group**](Group.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsDestroy

> UsersGroupsDestroy(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsList

> PaginatedGroupList UsersGroupsList(ctx).Id(id).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Limit(limit).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).Offset(offset).Q(q).Execute()





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
    id := []int32{int32(123)} // []int32 |  (optional)
    idGt := []int32{int32(123)} // []int32 |  (optional)
    idGte := []int32{int32(123)} // []int32 |  (optional)
    idLt := []int32{int32(123)} // []int32 |  (optional)
    idLte := []int32{int32(123)} // []int32 |  (optional)
    idN := []int32{int32(123)} // []int32 |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := []string{"Inner_example"} // []string |  (optional)
    nameIc := []string{"Inner_example"} // []string |  (optional)
    nameIe := []string{"Inner_example"} // []string |  (optional)
    nameIew := []string{"Inner_example"} // []string |  (optional)
    nameIre := []string{"Inner_example"} // []string |  (optional)
    nameIsw := []string{"Inner_example"} // []string |  (optional)
    nameN := []string{"Inner_example"} // []string |  (optional)
    nameNic := []string{"Inner_example"} // []string |  (optional)
    nameNie := []string{"Inner_example"} // []string |  (optional)
    nameNiew := []string{"Inner_example"} // []string |  (optional)
    nameNire := []string{"Inner_example"} // []string |  (optional)
    nameNisw := []string{"Inner_example"} // []string |  (optional)
    nameNre := []string{"Inner_example"} // []string |  (optional)
    nameRe := []string{"Inner_example"} // []string |  (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    q := "q_example" // string | Search (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsList(context.Background()).Id(id).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Limit(limit).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).Offset(offset).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsList`: PaginatedGroupList
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]int32** |  | 
 **idGt** | **[]int32** |  | 
 **idGte** | **[]int32** |  | 
 **idLt** | **[]int32** |  | 
 **idLte** | **[]int32** |  | 
 **idN** | **[]int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **[]string** |  | 
 **nameIc** | **[]string** |  | 
 **nameIe** | **[]string** |  | 
 **nameIew** | **[]string** |  | 
 **nameIre** | **[]string** |  | 
 **nameIsw** | **[]string** |  | 
 **nameN** | **[]string** |  | 
 **nameNic** | **[]string** |  | 
 **nameNie** | **[]string** |  | 
 **nameNiew** | **[]string** |  | 
 **nameNire** | **[]string** |  | 
 **nameNisw** | **[]string** |  | 
 **nameNre** | **[]string** |  | 
 **nameRe** | **[]string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **q** | **string** | Search | 

### Return type

[**PaginatedGroupList**](PaginatedGroupList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsPartialUpdate

> Group UsersGroupsPartialUpdate(ctx, id).PatchedGroup(patchedGroup).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this group.
    patchedGroup := *openapiclient.NewPatchedGroup() // PatchedGroup |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsPartialUpdate(context.Background(), id).PatchedGroup(patchedGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsPartialUpdate`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedGroup** | [**PatchedGroup**](PatchedGroup.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsRetrieve

> Group UsersGroupsRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsRetrieve`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsUpdate

> Group UsersGroupsUpdate(ctx, id).Group(group).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this group.
    group := *openapiclient.NewGroup(int32(123), "Url_example", "Name_example", int32(123), "Display_example") // Group | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsUpdate(context.Background(), id).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsUpdate`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | [**Group**](Group.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsBulkDestroy

> UsersPermissionsBulkDestroy(ctx).Execute()





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
    resp, r, err := apiClient.UsersApi.UsersPermissionsBulkDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsBulkDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsBulkDestroyRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsBulkPartialUpdate

> ObjectPermission UsersPermissionsBulkPartialUpdate(ctx).PatchedWritableObjectPermission(patchedWritableObjectPermission).Execute()





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
    patchedWritableObjectPermission := *openapiclient.NewPatchedWritableObjectPermission() // PatchedWritableObjectPermission |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsBulkPartialUpdate(context.Background()).PatchedWritableObjectPermission(patchedWritableObjectPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsBulkPartialUpdate`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedWritableObjectPermission** | [**PatchedWritableObjectPermission**](PatchedWritableObjectPermission.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsBulkUpdate

> ObjectPermission UsersPermissionsBulkUpdate(ctx).WritableObjectPermission(writableObjectPermission).Execute()





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
    writableObjectPermission := *openapiclient.NewWritableObjectPermission("Id_example", "Url_example", "Name_example", []string{"ObjectTypes_example"}, map[string]interface{}{"key": interface{}(123)}, "Display_example") // WritableObjectPermission | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsBulkUpdate(context.Background()).WritableObjectPermission(writableObjectPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsBulkUpdate`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableObjectPermission** | [**WritableObjectPermission**](WritableObjectPermission.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsCreate

> ObjectPermission UsersPermissionsCreate(ctx).WritableObjectPermission(writableObjectPermission).Execute()





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
    writableObjectPermission := *openapiclient.NewWritableObjectPermission("Id_example", "Url_example", "Name_example", []string{"ObjectTypes_example"}, map[string]interface{}{"key": interface{}(123)}, "Display_example") // WritableObjectPermission | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsCreate(context.Background()).WritableObjectPermission(writableObjectPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsCreate`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableObjectPermission** | [**WritableObjectPermission**](WritableObjectPermission.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsDestroy

> UsersPermissionsDestroy(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this permission.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsList

> PaginatedObjectPermissionList UsersPermissionsList(ctx).Enabled(enabled).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).Limit(limit).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).ObjectTypes(objectTypes).ObjectTypesN(objectTypesN).Offset(offset).User(user).UserN(userN).UserId(userId).UserIdN(userIdN).Execute()





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
    enabled := true // bool |  (optional)
    group := []string{"Inner_example"} // []string | Group (name) (optional)
    groupN := []string{"Inner_example"} // []string | Group (name) (optional)
    groupId := []int32{int32(123)} // []int32 | Group (optional)
    groupIdN := []int32{int32(123)} // []int32 | Group (optional)
    id := []string{"Inner_example"} // []string |  (optional)
    idIc := []string{"Inner_example"} // []string |  (optional)
    idIe := []string{"Inner_example"} // []string |  (optional)
    idIew := []string{"Inner_example"} // []string |  (optional)
    idIre := []string{"Inner_example"} // []string |  (optional)
    idIsw := []string{"Inner_example"} // []string |  (optional)
    idN := []string{"Inner_example"} // []string |  (optional)
    idNic := []string{"Inner_example"} // []string |  (optional)
    idNie := []string{"Inner_example"} // []string |  (optional)
    idNiew := []string{"Inner_example"} // []string |  (optional)
    idNire := []string{"Inner_example"} // []string |  (optional)
    idNisw := []string{"Inner_example"} // []string |  (optional)
    idNre := []string{"Inner_example"} // []string |  (optional)
    idRe := []string{"Inner_example"} // []string |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := []string{"Inner_example"} // []string |  (optional)
    nameIc := []string{"Inner_example"} // []string |  (optional)
    nameIe := []string{"Inner_example"} // []string |  (optional)
    nameIew := []string{"Inner_example"} // []string |  (optional)
    nameIre := []string{"Inner_example"} // []string |  (optional)
    nameIsw := []string{"Inner_example"} // []string |  (optional)
    nameN := []string{"Inner_example"} // []string |  (optional)
    nameNic := []string{"Inner_example"} // []string |  (optional)
    nameNie := []string{"Inner_example"} // []string |  (optional)
    nameNiew := []string{"Inner_example"} // []string |  (optional)
    nameNire := []string{"Inner_example"} // []string |  (optional)
    nameNisw := []string{"Inner_example"} // []string |  (optional)
    nameNre := []string{"Inner_example"} // []string |  (optional)
    nameRe := []string{"Inner_example"} // []string |  (optional)
    objectTypes := []int32{int32(123)} // []int32 |  (optional)
    objectTypesN := []int32{int32(123)} // []int32 |  (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    user := []string{"Inner_example"} // []string | User (name) (optional)
    userN := []string{"Inner_example"} // []string | User (name) (optional)
    userId := []string{"Inner_example"} // []string | User (optional)
    userIdN := []string{"Inner_example"} // []string | User (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsList(context.Background()).Enabled(enabled).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).Limit(limit).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).ObjectTypes(objectTypes).ObjectTypesN(objectTypesN).Offset(offset).User(user).UserN(userN).UserId(userId).UserIdN(userIdN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsList`: PaginatedObjectPermissionList
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **bool** |  | 
 **group** | **[]string** | Group (name) | 
 **groupN** | **[]string** | Group (name) | 
 **groupId** | **[]int32** | Group | 
 **groupIdN** | **[]int32** | Group | 
 **id** | **[]string** |  | 
 **idIc** | **[]string** |  | 
 **idIe** | **[]string** |  | 
 **idIew** | **[]string** |  | 
 **idIre** | **[]string** |  | 
 **idIsw** | **[]string** |  | 
 **idN** | **[]string** |  | 
 **idNic** | **[]string** |  | 
 **idNie** | **[]string** |  | 
 **idNiew** | **[]string** |  | 
 **idNire** | **[]string** |  | 
 **idNisw** | **[]string** |  | 
 **idNre** | **[]string** |  | 
 **idRe** | **[]string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **[]string** |  | 
 **nameIc** | **[]string** |  | 
 **nameIe** | **[]string** |  | 
 **nameIew** | **[]string** |  | 
 **nameIre** | **[]string** |  | 
 **nameIsw** | **[]string** |  | 
 **nameN** | **[]string** |  | 
 **nameNic** | **[]string** |  | 
 **nameNie** | **[]string** |  | 
 **nameNiew** | **[]string** |  | 
 **nameNire** | **[]string** |  | 
 **nameNisw** | **[]string** |  | 
 **nameNre** | **[]string** |  | 
 **nameRe** | **[]string** |  | 
 **objectTypes** | **[]int32** |  | 
 **objectTypesN** | **[]int32** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **user** | **[]string** | User (name) | 
 **userN** | **[]string** | User (name) | 
 **userId** | **[]string** | User | 
 **userIdN** | **[]string** | User | 

### Return type

[**PaginatedObjectPermissionList**](PaginatedObjectPermissionList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsPartialUpdate

> ObjectPermission UsersPermissionsPartialUpdate(ctx, id).PatchedWritableObjectPermission(patchedWritableObjectPermission).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this permission.
    patchedWritableObjectPermission := *openapiclient.NewPatchedWritableObjectPermission() // PatchedWritableObjectPermission |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsPartialUpdate(context.Background(), id).PatchedWritableObjectPermission(patchedWritableObjectPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsPartialUpdate`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWritableObjectPermission** | [**PatchedWritableObjectPermission**](PatchedWritableObjectPermission.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsRetrieve

> ObjectPermission UsersPermissionsRetrieve(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this permission.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsRetrieve`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsUpdate

> ObjectPermission UsersPermissionsUpdate(ctx, id).WritableObjectPermission(writableObjectPermission).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this permission.
    writableObjectPermission := *openapiclient.NewWritableObjectPermission("Id_example", "Url_example", "Name_example", []string{"ObjectTypes_example"}, map[string]interface{}{"key": interface{}(123)}, "Display_example") // WritableObjectPermission | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsUpdate(context.Background(), id).WritableObjectPermission(writableObjectPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsUpdate`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writableObjectPermission** | [**WritableObjectPermission**](WritableObjectPermission.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensBulkDestroy

> UsersTokensBulkDestroy(ctx).Execute()





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
    resp, r, err := apiClient.UsersApi.UsersTokensBulkDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensBulkDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensBulkDestroyRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensBulkPartialUpdate

> Token UsersTokensBulkPartialUpdate(ctx).PatchedToken(patchedToken).Execute()





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
    patchedToken := *openapiclient.NewPatchedToken() // PatchedToken |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensBulkPartialUpdate(context.Background()).PatchedToken(patchedToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTokensBulkPartialUpdate`: Token
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersTokensBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedToken** | [**PatchedToken**](PatchedToken.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensBulkUpdate

> Token UsersTokensBulkUpdate(ctx).Token(token).Execute()





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
    token := *openapiclient.NewToken("Id_example", "Url_example", "Display_example", time.Now()) // Token |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensBulkUpdate(context.Background()).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTokensBulkUpdate`: Token
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersTokensBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | [**Token**](Token.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensCreate

> Token UsersTokensCreate(ctx).Token(token).Execute()





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
    token := *openapiclient.NewToken("Id_example", "Url_example", "Display_example", time.Now()) // Token |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensCreate(context.Background()).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTokensCreate`: Token
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersTokensCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | [**Token**](Token.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensDestroy

> UsersTokensDestroy(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensList

> PaginatedTokenList UsersTokensList(ctx).Created(created).CreatedGt(createdGt).CreatedGte(createdGte).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).Expires(expires).ExpiresGt(expiresGt).ExpiresGte(expiresGte).ExpiresLt(expiresLt).ExpiresLte(expiresLte).ExpiresN(expiresN).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).Key(key).KeyIc(keyIc).KeyIe(keyIe).KeyIew(keyIew).KeyIre(keyIre).KeyIsw(keyIsw).KeyN(keyN).KeyNic(keyNic).KeyNie(keyNie).KeyNiew(keyNiew).KeyNire(keyNire).KeyNisw(keyNisw).KeyNre(keyNre).KeyRe(keyRe).Limit(limit).Offset(offset).Q(q).WriteEnabled(writeEnabled).Execute()





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
    created := []time.Time{time.Now()} // []time.Time |  (optional)
    createdGt := []time.Time{time.Now()} // []time.Time |  (optional)
    createdGte := []time.Time{time.Now()} // []time.Time |  (optional)
    createdLt := []time.Time{time.Now()} // []time.Time |  (optional)
    createdLte := []time.Time{time.Now()} // []time.Time |  (optional)
    createdN := []time.Time{time.Now()} // []time.Time |  (optional)
    expires := []time.Time{time.Now()} // []time.Time |  (optional)
    expiresGt := []time.Time{time.Now()} // []time.Time |  (optional)
    expiresGte := []time.Time{time.Now()} // []time.Time |  (optional)
    expiresLt := []time.Time{time.Now()} // []time.Time |  (optional)
    expiresLte := []time.Time{time.Now()} // []time.Time |  (optional)
    expiresN := []time.Time{time.Now()} // []time.Time |  (optional)
    id := []string{"Inner_example"} // []string |  (optional)
    idIc := []string{"Inner_example"} // []string |  (optional)
    idIe := []string{"Inner_example"} // []string |  (optional)
    idIew := []string{"Inner_example"} // []string |  (optional)
    idIre := []string{"Inner_example"} // []string |  (optional)
    idIsw := []string{"Inner_example"} // []string |  (optional)
    idN := []string{"Inner_example"} // []string |  (optional)
    idNic := []string{"Inner_example"} // []string |  (optional)
    idNie := []string{"Inner_example"} // []string |  (optional)
    idNiew := []string{"Inner_example"} // []string |  (optional)
    idNire := []string{"Inner_example"} // []string |  (optional)
    idNisw := []string{"Inner_example"} // []string |  (optional)
    idNre := []string{"Inner_example"} // []string |  (optional)
    idRe := []string{"Inner_example"} // []string |  (optional)
    key := []string{"Inner_example"} // []string |  (optional)
    keyIc := []string{"Inner_example"} // []string |  (optional)
    keyIe := []string{"Inner_example"} // []string |  (optional)
    keyIew := []string{"Inner_example"} // []string |  (optional)
    keyIre := []string{"Inner_example"} // []string |  (optional)
    keyIsw := []string{"Inner_example"} // []string |  (optional)
    keyN := []string{"Inner_example"} // []string |  (optional)
    keyNic := []string{"Inner_example"} // []string |  (optional)
    keyNie := []string{"Inner_example"} // []string |  (optional)
    keyNiew := []string{"Inner_example"} // []string |  (optional)
    keyNire := []string{"Inner_example"} // []string |  (optional)
    keyNisw := []string{"Inner_example"} // []string |  (optional)
    keyNre := []string{"Inner_example"} // []string |  (optional)
    keyRe := []string{"Inner_example"} // []string |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    q := "q_example" // string | Search (optional)
    writeEnabled := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensList(context.Background()).Created(created).CreatedGt(createdGt).CreatedGte(createdGte).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).Expires(expires).ExpiresGt(expiresGt).ExpiresGte(expiresGte).ExpiresLt(expiresLt).ExpiresLte(expiresLte).ExpiresN(expiresN).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).Key(key).KeyIc(keyIc).KeyIe(keyIe).KeyIew(keyIew).KeyIre(keyIre).KeyIsw(keyIsw).KeyN(keyN).KeyNic(keyNic).KeyNie(keyNie).KeyNiew(keyNiew).KeyNire(keyNire).KeyNisw(keyNisw).KeyNre(keyNre).KeyRe(keyRe).Limit(limit).Offset(offset).Q(q).WriteEnabled(writeEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTokensList`: PaginatedTokenList
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersTokensList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created** | [**[]time.Time**](time.Time.md) |  | 
 **createdGt** | [**[]time.Time**](time.Time.md) |  | 
 **createdGte** | [**[]time.Time**](time.Time.md) |  | 
 **createdLt** | [**[]time.Time**](time.Time.md) |  | 
 **createdLte** | [**[]time.Time**](time.Time.md) |  | 
 **createdN** | [**[]time.Time**](time.Time.md) |  | 
 **expires** | [**[]time.Time**](time.Time.md) |  | 
 **expiresGt** | [**[]time.Time**](time.Time.md) |  | 
 **expiresGte** | [**[]time.Time**](time.Time.md) |  | 
 **expiresLt** | [**[]time.Time**](time.Time.md) |  | 
 **expiresLte** | [**[]time.Time**](time.Time.md) |  | 
 **expiresN** | [**[]time.Time**](time.Time.md) |  | 
 **id** | **[]string** |  | 
 **idIc** | **[]string** |  | 
 **idIe** | **[]string** |  | 
 **idIew** | **[]string** |  | 
 **idIre** | **[]string** |  | 
 **idIsw** | **[]string** |  | 
 **idN** | **[]string** |  | 
 **idNic** | **[]string** |  | 
 **idNie** | **[]string** |  | 
 **idNiew** | **[]string** |  | 
 **idNire** | **[]string** |  | 
 **idNisw** | **[]string** |  | 
 **idNre** | **[]string** |  | 
 **idRe** | **[]string** |  | 
 **key** | **[]string** |  | 
 **keyIc** | **[]string** |  | 
 **keyIe** | **[]string** |  | 
 **keyIew** | **[]string** |  | 
 **keyIre** | **[]string** |  | 
 **keyIsw** | **[]string** |  | 
 **keyN** | **[]string** |  | 
 **keyNic** | **[]string** |  | 
 **keyNie** | **[]string** |  | 
 **keyNiew** | **[]string** |  | 
 **keyNire** | **[]string** |  | 
 **keyNisw** | **[]string** |  | 
 **keyNre** | **[]string** |  | 
 **keyRe** | **[]string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **q** | **string** | Search | 
 **writeEnabled** | **bool** |  | 

### Return type

[**PaginatedTokenList**](PaginatedTokenList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensPartialUpdate

> Token UsersTokensPartialUpdate(ctx, id).PatchedToken(patchedToken).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this token.
    patchedToken := *openapiclient.NewPatchedToken() // PatchedToken |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensPartialUpdate(context.Background(), id).PatchedToken(patchedToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTokensPartialUpdate`: Token
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersTokensPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedToken** | [**PatchedToken**](PatchedToken.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensRetrieve

> Token UsersTokensRetrieve(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTokensRetrieve`: Token
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersTokensRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensUpdate

> Token UsersTokensUpdate(ctx, id).Token(token).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this token.
    token := *openapiclient.NewToken("Id_example", "Url_example", "Display_example", time.Now()) // Token |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensUpdate(context.Background(), id).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTokensUpdate`: Token
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersTokensUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | [**Token**](Token.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersBulkDestroy

> UsersUsersBulkDestroy(ctx).Execute()





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
    resp, r, err := apiClient.UsersApi.UsersUsersBulkDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersBulkDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersBulkDestroyRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersBulkPartialUpdate

> User UsersUsersBulkPartialUpdate(ctx).PatchedWritableUser(patchedWritableUser).Execute()





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
    patchedWritableUser := *openapiclient.NewPatchedWritableUser() // PatchedWritableUser |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersBulkPartialUpdate(context.Background()).PatchedWritableUser(patchedWritableUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersBulkPartialUpdate`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedWritableUser** | [**PatchedWritableUser**](PatchedWritableUser.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersBulkUpdate

> User UsersUsersBulkUpdate(ctx).WritableUser(writableUser).Execute()





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
    writableUser := *openapiclient.NewWritableUser("Id_example", "Url_example", "Username_example", "Password_example", "Display_example") // WritableUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersBulkUpdate(context.Background()).WritableUser(writableUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersBulkUpdate`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableUser** | [**WritableUser**](WritableUser.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersCreate

> User UsersUsersCreate(ctx).WritableUser(writableUser).Execute()





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
    writableUser := *openapiclient.NewWritableUser("Id_example", "Url_example", "Username_example", "Password_example", "Display_example") // WritableUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersCreate(context.Background()).WritableUser(writableUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersCreate`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableUser** | [**WritableUser**](WritableUser.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersDestroy

> UsersUsersDestroy(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersList

> PaginatedUserList UsersUsersList(ctx).Email(email).EmailIc(emailIc).EmailIe(emailIe).EmailIew(emailIew).EmailIre(emailIre).EmailIsw(emailIsw).EmailN(emailN).EmailNic(emailNic).EmailNie(emailNie).EmailNiew(emailNiew).EmailNire(emailNire).EmailNisw(emailNisw).EmailNre(emailNre).EmailRe(emailRe).FirstName(firstName).FirstNameIc(firstNameIc).FirstNameIe(firstNameIe).FirstNameIew(firstNameIew).FirstNameIre(firstNameIre).FirstNameIsw(firstNameIsw).FirstNameN(firstNameN).FirstNameNic(firstNameNic).FirstNameNie(firstNameNie).FirstNameNiew(firstNameNiew).FirstNameNire(firstNameNire).FirstNameNisw(firstNameNisw).FirstNameNre(firstNameNre).FirstNameRe(firstNameRe).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).IsActive(isActive).IsStaff(isStaff).LastName(lastName).LastNameIc(lastNameIc).LastNameIe(lastNameIe).LastNameIew(lastNameIew).LastNameIre(lastNameIre).LastNameIsw(lastNameIsw).LastNameN(lastNameN).LastNameNic(lastNameNic).LastNameNie(lastNameNie).LastNameNiew(lastNameNiew).LastNameNire(lastNameNire).LastNameNisw(lastNameNisw).LastNameNre(lastNameNre).LastNameRe(lastNameRe).Limit(limit).Offset(offset).Q(q).Username(username).UsernameIc(usernameIc).UsernameIe(usernameIe).UsernameIew(usernameIew).UsernameIre(usernameIre).UsernameIsw(usernameIsw).UsernameN(usernameN).UsernameNic(usernameNic).UsernameNie(usernameNie).UsernameNiew(usernameNiew).UsernameNire(usernameNire).UsernameNisw(usernameNisw).UsernameNre(usernameNre).UsernameRe(usernameRe).Execute()





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
    email := []string{"Inner_example"} // []string |  (optional)
    emailIc := []string{"Inner_example"} // []string |  (optional)
    emailIe := []string{"Inner_example"} // []string |  (optional)
    emailIew := []string{"Inner_example"} // []string |  (optional)
    emailIre := []string{"Inner_example"} // []string |  (optional)
    emailIsw := []string{"Inner_example"} // []string |  (optional)
    emailN := []string{"Inner_example"} // []string |  (optional)
    emailNic := []string{"Inner_example"} // []string |  (optional)
    emailNie := []string{"Inner_example"} // []string |  (optional)
    emailNiew := []string{"Inner_example"} // []string |  (optional)
    emailNire := []string{"Inner_example"} // []string |  (optional)
    emailNisw := []string{"Inner_example"} // []string |  (optional)
    emailNre := []string{"Inner_example"} // []string |  (optional)
    emailRe := []string{"Inner_example"} // []string |  (optional)
    firstName := []string{"Inner_example"} // []string |  (optional)
    firstNameIc := []string{"Inner_example"} // []string |  (optional)
    firstNameIe := []string{"Inner_example"} // []string |  (optional)
    firstNameIew := []string{"Inner_example"} // []string |  (optional)
    firstNameIre := []string{"Inner_example"} // []string |  (optional)
    firstNameIsw := []string{"Inner_example"} // []string |  (optional)
    firstNameN := []string{"Inner_example"} // []string |  (optional)
    firstNameNic := []string{"Inner_example"} // []string |  (optional)
    firstNameNie := []string{"Inner_example"} // []string |  (optional)
    firstNameNiew := []string{"Inner_example"} // []string |  (optional)
    firstNameNire := []string{"Inner_example"} // []string |  (optional)
    firstNameNisw := []string{"Inner_example"} // []string |  (optional)
    firstNameNre := []string{"Inner_example"} // []string |  (optional)
    firstNameRe := []string{"Inner_example"} // []string |  (optional)
    group := []string{"Inner_example"} // []string | Group (name) (optional)
    groupN := []string{"Inner_example"} // []string | Group (name) (optional)
    groupId := []int32{int32(123)} // []int32 | Group (optional)
    groupIdN := []int32{int32(123)} // []int32 | Group (optional)
    id := []string{"Inner_example"} // []string |  (optional)
    idIc := []string{"Inner_example"} // []string |  (optional)
    idIe := []string{"Inner_example"} // []string |  (optional)
    idIew := []string{"Inner_example"} // []string |  (optional)
    idIre := []string{"Inner_example"} // []string |  (optional)
    idIsw := []string{"Inner_example"} // []string |  (optional)
    idN := []string{"Inner_example"} // []string |  (optional)
    idNic := []string{"Inner_example"} // []string |  (optional)
    idNie := []string{"Inner_example"} // []string |  (optional)
    idNiew := []string{"Inner_example"} // []string |  (optional)
    idNire := []string{"Inner_example"} // []string |  (optional)
    idNisw := []string{"Inner_example"} // []string |  (optional)
    idNre := []string{"Inner_example"} // []string |  (optional)
    idRe := []string{"Inner_example"} // []string |  (optional)
    isActive := true // bool |  (optional)
    isStaff := true // bool |  (optional)
    lastName := []string{"Inner_example"} // []string |  (optional)
    lastNameIc := []string{"Inner_example"} // []string |  (optional)
    lastNameIe := []string{"Inner_example"} // []string |  (optional)
    lastNameIew := []string{"Inner_example"} // []string |  (optional)
    lastNameIre := []string{"Inner_example"} // []string |  (optional)
    lastNameIsw := []string{"Inner_example"} // []string |  (optional)
    lastNameN := []string{"Inner_example"} // []string |  (optional)
    lastNameNic := []string{"Inner_example"} // []string |  (optional)
    lastNameNie := []string{"Inner_example"} // []string |  (optional)
    lastNameNiew := []string{"Inner_example"} // []string |  (optional)
    lastNameNire := []string{"Inner_example"} // []string |  (optional)
    lastNameNisw := []string{"Inner_example"} // []string |  (optional)
    lastNameNre := []string{"Inner_example"} // []string |  (optional)
    lastNameRe := []string{"Inner_example"} // []string |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    q := "q_example" // string | Search (optional)
    username := []string{"Inner_example"} // []string |  (optional)
    usernameIc := []string{"Inner_example"} // []string |  (optional)
    usernameIe := []string{"Inner_example"} // []string |  (optional)
    usernameIew := []string{"Inner_example"} // []string |  (optional)
    usernameIre := []string{"Inner_example"} // []string |  (optional)
    usernameIsw := []string{"Inner_example"} // []string |  (optional)
    usernameN := []string{"Inner_example"} // []string |  (optional)
    usernameNic := []string{"Inner_example"} // []string |  (optional)
    usernameNie := []string{"Inner_example"} // []string |  (optional)
    usernameNiew := []string{"Inner_example"} // []string |  (optional)
    usernameNire := []string{"Inner_example"} // []string |  (optional)
    usernameNisw := []string{"Inner_example"} // []string |  (optional)
    usernameNre := []string{"Inner_example"} // []string |  (optional)
    usernameRe := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersList(context.Background()).Email(email).EmailIc(emailIc).EmailIe(emailIe).EmailIew(emailIew).EmailIre(emailIre).EmailIsw(emailIsw).EmailN(emailN).EmailNic(emailNic).EmailNie(emailNie).EmailNiew(emailNiew).EmailNire(emailNire).EmailNisw(emailNisw).EmailNre(emailNre).EmailRe(emailRe).FirstName(firstName).FirstNameIc(firstNameIc).FirstNameIe(firstNameIe).FirstNameIew(firstNameIew).FirstNameIre(firstNameIre).FirstNameIsw(firstNameIsw).FirstNameN(firstNameN).FirstNameNic(firstNameNic).FirstNameNie(firstNameNie).FirstNameNiew(firstNameNiew).FirstNameNire(firstNameNire).FirstNameNisw(firstNameNisw).FirstNameNre(firstNameNre).FirstNameRe(firstNameRe).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).IsActive(isActive).IsStaff(isStaff).LastName(lastName).LastNameIc(lastNameIc).LastNameIe(lastNameIe).LastNameIew(lastNameIew).LastNameIre(lastNameIre).LastNameIsw(lastNameIsw).LastNameN(lastNameN).LastNameNic(lastNameNic).LastNameNie(lastNameNie).LastNameNiew(lastNameNiew).LastNameNire(lastNameNire).LastNameNisw(lastNameNisw).LastNameNre(lastNameNre).LastNameRe(lastNameRe).Limit(limit).Offset(offset).Q(q).Username(username).UsernameIc(usernameIc).UsernameIe(usernameIe).UsernameIew(usernameIew).UsernameIre(usernameIre).UsernameIsw(usernameIsw).UsernameN(usernameN).UsernameNic(usernameNic).UsernameNie(usernameNie).UsernameNiew(usernameNiew).UsernameNire(usernameNire).UsernameNisw(usernameNisw).UsernameNre(usernameNre).UsernameRe(usernameRe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersList`: PaginatedUserList
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **[]string** |  | 
 **emailIc** | **[]string** |  | 
 **emailIe** | **[]string** |  | 
 **emailIew** | **[]string** |  | 
 **emailIre** | **[]string** |  | 
 **emailIsw** | **[]string** |  | 
 **emailN** | **[]string** |  | 
 **emailNic** | **[]string** |  | 
 **emailNie** | **[]string** |  | 
 **emailNiew** | **[]string** |  | 
 **emailNire** | **[]string** |  | 
 **emailNisw** | **[]string** |  | 
 **emailNre** | **[]string** |  | 
 **emailRe** | **[]string** |  | 
 **firstName** | **[]string** |  | 
 **firstNameIc** | **[]string** |  | 
 **firstNameIe** | **[]string** |  | 
 **firstNameIew** | **[]string** |  | 
 **firstNameIre** | **[]string** |  | 
 **firstNameIsw** | **[]string** |  | 
 **firstNameN** | **[]string** |  | 
 **firstNameNic** | **[]string** |  | 
 **firstNameNie** | **[]string** |  | 
 **firstNameNiew** | **[]string** |  | 
 **firstNameNire** | **[]string** |  | 
 **firstNameNisw** | **[]string** |  | 
 **firstNameNre** | **[]string** |  | 
 **firstNameRe** | **[]string** |  | 
 **group** | **[]string** | Group (name) | 
 **groupN** | **[]string** | Group (name) | 
 **groupId** | **[]int32** | Group | 
 **groupIdN** | **[]int32** | Group | 
 **id** | **[]string** |  | 
 **idIc** | **[]string** |  | 
 **idIe** | **[]string** |  | 
 **idIew** | **[]string** |  | 
 **idIre** | **[]string** |  | 
 **idIsw** | **[]string** |  | 
 **idN** | **[]string** |  | 
 **idNic** | **[]string** |  | 
 **idNie** | **[]string** |  | 
 **idNiew** | **[]string** |  | 
 **idNire** | **[]string** |  | 
 **idNisw** | **[]string** |  | 
 **idNre** | **[]string** |  | 
 **idRe** | **[]string** |  | 
 **isActive** | **bool** |  | 
 **isStaff** | **bool** |  | 
 **lastName** | **[]string** |  | 
 **lastNameIc** | **[]string** |  | 
 **lastNameIe** | **[]string** |  | 
 **lastNameIew** | **[]string** |  | 
 **lastNameIre** | **[]string** |  | 
 **lastNameIsw** | **[]string** |  | 
 **lastNameN** | **[]string** |  | 
 **lastNameNic** | **[]string** |  | 
 **lastNameNie** | **[]string** |  | 
 **lastNameNiew** | **[]string** |  | 
 **lastNameNire** | **[]string** |  | 
 **lastNameNisw** | **[]string** |  | 
 **lastNameNre** | **[]string** |  | 
 **lastNameRe** | **[]string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **q** | **string** | Search | 
 **username** | **[]string** |  | 
 **usernameIc** | **[]string** |  | 
 **usernameIe** | **[]string** |  | 
 **usernameIew** | **[]string** |  | 
 **usernameIre** | **[]string** |  | 
 **usernameIsw** | **[]string** |  | 
 **usernameN** | **[]string** |  | 
 **usernameNic** | **[]string** |  | 
 **usernameNie** | **[]string** |  | 
 **usernameNiew** | **[]string** |  | 
 **usernameNire** | **[]string** |  | 
 **usernameNisw** | **[]string** |  | 
 **usernameNre** | **[]string** |  | 
 **usernameRe** | **[]string** |  | 

### Return type

[**PaginatedUserList**](PaginatedUserList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersPartialUpdate

> User UsersUsersPartialUpdate(ctx, id).PatchedWritableUser(patchedWritableUser).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this user.
    patchedWritableUser := *openapiclient.NewPatchedWritableUser() // PatchedWritableUser |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersPartialUpdate(context.Background(), id).PatchedWritableUser(patchedWritableUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersPartialUpdate`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWritableUser** | [**PatchedWritableUser**](PatchedWritableUser.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersRetrieve

> User UsersUsersRetrieve(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersRetrieve`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersUpdate

> User UsersUsersUpdate(ctx, id).WritableUser(writableUser).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this user.
    writableUser := *openapiclient.NewWritableUser("Id_example", "Url_example", "Username_example", "Password_example", "Display_example") // WritableUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersUpdate(context.Background(), id).WritableUser(writableUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersUpdate`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writableUser** | [**WritableUser**](WritableUser.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

