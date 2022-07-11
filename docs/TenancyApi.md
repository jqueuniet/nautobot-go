# \TenancyApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenancyTenantGroupsBulkDestroy**](TenancyApi.md#TenancyTenantGroupsBulkDestroy) | **Delete** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsBulkPartialUpdate**](TenancyApi.md#TenancyTenantGroupsBulkPartialUpdate) | **Patch** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsBulkUpdate**](TenancyApi.md#TenancyTenantGroupsBulkUpdate) | **Put** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsCreate**](TenancyApi.md#TenancyTenantGroupsCreate) | **Post** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsDestroy**](TenancyApi.md#TenancyTenantGroupsDestroy) | **Delete** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantGroupsList**](TenancyApi.md#TenancyTenantGroupsList) | **Get** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsPartialUpdate**](TenancyApi.md#TenancyTenantGroupsPartialUpdate) | **Patch** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantGroupsRetrieve**](TenancyApi.md#TenancyTenantGroupsRetrieve) | **Get** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantGroupsUpdate**](TenancyApi.md#TenancyTenantGroupsUpdate) | **Put** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantsBulkDestroy**](TenancyApi.md#TenancyTenantsBulkDestroy) | **Delete** /tenancy/tenants/ | 
[**TenancyTenantsBulkPartialUpdate**](TenancyApi.md#TenancyTenantsBulkPartialUpdate) | **Patch** /tenancy/tenants/ | 
[**TenancyTenantsBulkUpdate**](TenancyApi.md#TenancyTenantsBulkUpdate) | **Put** /tenancy/tenants/ | 
[**TenancyTenantsCreate**](TenancyApi.md#TenancyTenantsCreate) | **Post** /tenancy/tenants/ | 
[**TenancyTenantsDestroy**](TenancyApi.md#TenancyTenantsDestroy) | **Delete** /tenancy/tenants/{id}/ | 
[**TenancyTenantsList**](TenancyApi.md#TenancyTenantsList) | **Get** /tenancy/tenants/ | 
[**TenancyTenantsPartialUpdate**](TenancyApi.md#TenancyTenantsPartialUpdate) | **Patch** /tenancy/tenants/{id}/ | 
[**TenancyTenantsRetrieve**](TenancyApi.md#TenancyTenantsRetrieve) | **Get** /tenancy/tenants/{id}/ | 
[**TenancyTenantsUpdate**](TenancyApi.md#TenancyTenantsUpdate) | **Put** /tenancy/tenants/{id}/ | 



## TenancyTenantGroupsBulkDestroy

> TenancyTenantGroupsBulkDestroy(ctx).Execute()





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
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsBulkDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsBulkDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsBulkDestroyRequest struct via the builder pattern


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


## TenancyTenantGroupsBulkPartialUpdate

> TenantGroup TenancyTenantGroupsBulkPartialUpdate(ctx).PatchedWritableTenantGroup(patchedWritableTenantGroup).Execute()





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
    patchedWritableTenantGroup := *openapiclient.NewPatchedWritableTenantGroup() // PatchedWritableTenantGroup |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsBulkPartialUpdate(context.Background()).PatchedWritableTenantGroup(patchedWritableTenantGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantGroupsBulkPartialUpdate`: TenantGroup
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantGroupsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedWritableTenantGroup** | [**PatchedWritableTenantGroup**](PatchedWritableTenantGroup.md) |  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsBulkUpdate

> TenantGroup TenancyTenantGroupsBulkUpdate(ctx).WritableTenantGroup(writableTenantGroup).Execute()





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
    writableTenantGroup := *openapiclient.NewWritableTenantGroup("Id_example", "Url_example", "Name_example", int32(123), int32(123), time.Now(), time.Now(), map[string]interface{}{"key": interface{}(123)}, "Display_example") // WritableTenantGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsBulkUpdate(context.Background()).WritableTenantGroup(writableTenantGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantGroupsBulkUpdate`: TenantGroup
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantGroupsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableTenantGroup** | [**WritableTenantGroup**](WritableTenantGroup.md) |  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsCreate

> TenantGroup TenancyTenantGroupsCreate(ctx).WritableTenantGroup(writableTenantGroup).Execute()





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
    writableTenantGroup := *openapiclient.NewWritableTenantGroup("Id_example", "Url_example", "Name_example", int32(123), int32(123), time.Now(), time.Now(), map[string]interface{}{"key": interface{}(123)}, "Display_example") // WritableTenantGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsCreate(context.Background()).WritableTenantGroup(writableTenantGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantGroupsCreate`: TenantGroup
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantGroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableTenantGroup** | [**WritableTenantGroup**](WritableTenantGroup.md) |  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsDestroy

> TenancyTenantGroupsDestroy(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsDestroyRequest struct via the builder pattern


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


## TenancyTenantGroupsList

> PaginatedTenantGroupList TenancyTenantGroupsList(ctx).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).Description(description).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIre(descriptionIre).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNire(descriptionNire).DescriptionNisw(descriptionNisw).DescriptionNre(descriptionNre).DescriptionRe(descriptionRe).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).LastUpdated(lastUpdated).LastUpdatedGte(lastUpdatedGte).LastUpdatedLte(lastUpdatedLte).Limit(limit).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).Offset(offset).Parent(parent).ParentN(parentN).ParentId(parentId).ParentIdN(parentIdN).Q(q).Slug(slug).SlugIc(slugIc).SlugIe(slugIe).SlugIew(slugIew).SlugIre(slugIre).SlugIsw(slugIsw).SlugN(slugN).SlugNic(slugNic).SlugNie(slugNie).SlugNiew(slugNiew).SlugNire(slugNire).SlugNisw(slugNisw).SlugNre(slugNre).SlugRe(slugRe).Execute()





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
    created := time.Now() // string |  (optional)
    createdGte := time.Now() // string |  (optional)
    createdLte := time.Now() // string |  (optional)
    description := []string{"Inner_example"} // []string |  (optional)
    descriptionIc := []string{"Inner_example"} // []string |  (optional)
    descriptionIe := []string{"Inner_example"} // []string |  (optional)
    descriptionIew := []string{"Inner_example"} // []string |  (optional)
    descriptionIre := []string{"Inner_example"} // []string |  (optional)
    descriptionIsw := []string{"Inner_example"} // []string |  (optional)
    descriptionN := []string{"Inner_example"} // []string |  (optional)
    descriptionNic := []string{"Inner_example"} // []string |  (optional)
    descriptionNie := []string{"Inner_example"} // []string |  (optional)
    descriptionNiew := []string{"Inner_example"} // []string |  (optional)
    descriptionNire := []string{"Inner_example"} // []string |  (optional)
    descriptionNisw := []string{"Inner_example"} // []string |  (optional)
    descriptionNre := []string{"Inner_example"} // []string |  (optional)
    descriptionRe := []string{"Inner_example"} // []string |  (optional)
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
    lastUpdated := time.Now() // time.Time |  (optional)
    lastUpdatedGte := time.Now() // time.Time |  (optional)
    lastUpdatedLte := time.Now() // time.Time |  (optional)
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
    parent := []string{"Inner_example"} // []string | Tenant group group (slug) (optional)
    parentN := []string{"Inner_example"} // []string | Tenant group group (slug) (optional)
    parentId := []*string{"Inner_example"} // []*string | Tenant group (ID) (optional)
    parentIdN := []*string{"Inner_example"} // []*string | Tenant group (ID) (optional)
    q := "q_example" // string | Search (optional)
    slug := []string{"Inner_example"} // []string |  (optional)
    slugIc := []string{"Inner_example"} // []string |  (optional)
    slugIe := []string{"Inner_example"} // []string |  (optional)
    slugIew := []string{"Inner_example"} // []string |  (optional)
    slugIre := []string{"Inner_example"} // []string |  (optional)
    slugIsw := []string{"Inner_example"} // []string |  (optional)
    slugN := []string{"Inner_example"} // []string |  (optional)
    slugNic := []string{"Inner_example"} // []string |  (optional)
    slugNie := []string{"Inner_example"} // []string |  (optional)
    slugNiew := []string{"Inner_example"} // []string |  (optional)
    slugNire := []string{"Inner_example"} // []string |  (optional)
    slugNisw := []string{"Inner_example"} // []string |  (optional)
    slugNre := []string{"Inner_example"} // []string |  (optional)
    slugRe := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsList(context.Background()).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).Description(description).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIre(descriptionIre).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNire(descriptionNire).DescriptionNisw(descriptionNisw).DescriptionNre(descriptionNre).DescriptionRe(descriptionRe).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).LastUpdated(lastUpdated).LastUpdatedGte(lastUpdatedGte).LastUpdatedLte(lastUpdatedLte).Limit(limit).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).Offset(offset).Parent(parent).ParentN(parentN).ParentId(parentId).ParentIdN(parentIdN).Q(q).Slug(slug).SlugIc(slugIc).SlugIe(slugIe).SlugIew(slugIew).SlugIre(slugIre).SlugIsw(slugIsw).SlugN(slugN).SlugNic(slugNic).SlugNie(slugNie).SlugNiew(slugNiew).SlugNire(slugNire).SlugNisw(slugNisw).SlugNre(slugNre).SlugRe(slugRe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantGroupsList`: PaginatedTenantGroupList
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantGroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created** | **string** |  | 
 **createdGte** | **string** |  | 
 **createdLte** | **string** |  | 
 **description** | **[]string** |  | 
 **descriptionIc** | **[]string** |  | 
 **descriptionIe** | **[]string** |  | 
 **descriptionIew** | **[]string** |  | 
 **descriptionIre** | **[]string** |  | 
 **descriptionIsw** | **[]string** |  | 
 **descriptionN** | **[]string** |  | 
 **descriptionNic** | **[]string** |  | 
 **descriptionNie** | **[]string** |  | 
 **descriptionNiew** | **[]string** |  | 
 **descriptionNire** | **[]string** |  | 
 **descriptionNisw** | **[]string** |  | 
 **descriptionNre** | **[]string** |  | 
 **descriptionRe** | **[]string** |  | 
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
 **lastUpdated** | **time.Time** |  | 
 **lastUpdatedGte** | **time.Time** |  | 
 **lastUpdatedLte** | **time.Time** |  | 
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
 **parent** | **[]string** | Tenant group group (slug) | 
 **parentN** | **[]string** | Tenant group group (slug) | 
 **parentId** | **[]string** | Tenant group (ID) | 
 **parentIdN** | **[]string** | Tenant group (ID) | 
 **q** | **string** | Search | 
 **slug** | **[]string** |  | 
 **slugIc** | **[]string** |  | 
 **slugIe** | **[]string** |  | 
 **slugIew** | **[]string** |  | 
 **slugIre** | **[]string** |  | 
 **slugIsw** | **[]string** |  | 
 **slugN** | **[]string** |  | 
 **slugNic** | **[]string** |  | 
 **slugNie** | **[]string** |  | 
 **slugNiew** | **[]string** |  | 
 **slugNire** | **[]string** |  | 
 **slugNisw** | **[]string** |  | 
 **slugNre** | **[]string** |  | 
 **slugRe** | **[]string** |  | 

### Return type

[**PaginatedTenantGroupList**](PaginatedTenantGroupList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsPartialUpdate

> TenantGroup TenancyTenantGroupsPartialUpdate(ctx, id).PatchedWritableTenantGroup(patchedWritableTenantGroup).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant group.
    patchedWritableTenantGroup := *openapiclient.NewPatchedWritableTenantGroup() // PatchedWritableTenantGroup |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsPartialUpdate(context.Background(), id).PatchedWritableTenantGroup(patchedWritableTenantGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantGroupsPartialUpdate`: TenantGroup
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantGroupsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWritableTenantGroup** | [**PatchedWritableTenantGroup**](PatchedWritableTenantGroup.md) |  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsRetrieve

> TenantGroup TenancyTenantGroupsRetrieve(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantGroupsRetrieve`: TenantGroup
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantGroupsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantGroupsUpdate

> TenantGroup TenancyTenantGroupsUpdate(ctx, id).WritableTenantGroup(writableTenantGroup).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant group.
    writableTenantGroup := *openapiclient.NewWritableTenantGroup("Id_example", "Url_example", "Name_example", int32(123), int32(123), time.Now(), time.Now(), map[string]interface{}{"key": interface{}(123)}, "Display_example") // WritableTenantGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantGroupsUpdate(context.Background(), id).WritableTenantGroup(writableTenantGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantGroupsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantGroupsUpdate`: TenantGroup
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantGroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writableTenantGroup** | [**WritableTenantGroup**](WritableTenantGroup.md) |  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsBulkDestroy

> TenancyTenantsBulkDestroy(ctx).Execute()





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
    resp, r, err := apiClient.TenancyApi.TenancyTenantsBulkDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsBulkDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsBulkDestroyRequest struct via the builder pattern


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


## TenancyTenantsBulkPartialUpdate

> Tenant TenancyTenantsBulkPartialUpdate(ctx).PatchedWritableTenant(patchedWritableTenant).Execute()





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
    patchedWritableTenant := *openapiclient.NewPatchedWritableTenant() // PatchedWritableTenant |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsBulkPartialUpdate(context.Background()).PatchedWritableTenant(patchedWritableTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantsBulkPartialUpdate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedWritableTenant** | [**PatchedWritableTenant**](PatchedWritableTenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsBulkUpdate

> Tenant TenancyTenantsBulkUpdate(ctx).WritableTenant(writableTenant).Execute()





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
    writableTenant := *openapiclient.NewWritableTenant("Id_example", "Url_example", "Name_example", time.Now(), time.Now(), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), map[string]interface{}{"key": interface{}(123)}, "Display_example") // WritableTenant | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsBulkUpdate(context.Background()).WritableTenant(writableTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantsBulkUpdate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableTenant** | [**WritableTenant**](WritableTenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsCreate

> Tenant TenancyTenantsCreate(ctx).WritableTenant(writableTenant).Execute()





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
    writableTenant := *openapiclient.NewWritableTenant("Id_example", "Url_example", "Name_example", time.Now(), time.Now(), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), map[string]interface{}{"key": interface{}(123)}, "Display_example") // WritableTenant | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsCreate(context.Background()).WritableTenant(writableTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantsCreate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableTenant** | [**WritableTenant**](WritableTenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsDestroy

> TenancyTenantsDestroy(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsDestroyRequest struct via the builder pattern


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


## TenancyTenantsList

> PaginatedTenantList TenancyTenantsList(ctx).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).LastUpdated(lastUpdated).LastUpdatedGte(lastUpdatedGte).LastUpdatedLte(lastUpdatedLte).Limit(limit).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).Offset(offset).Q(q).Slug(slug).SlugIc(slugIc).SlugIe(slugIe).SlugIew(slugIew).SlugIre(slugIre).SlugIsw(slugIsw).SlugN(slugN).SlugNic(slugNic).SlugNie(slugNie).SlugNiew(slugNiew).SlugNire(slugNire).SlugNisw(slugNisw).SlugNre(slugNre).SlugRe(slugRe).Tag(tag).TagN(tagN).Execute()





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
    created := time.Now() // string |  (optional)
    createdGte := time.Now() // string |  (optional)
    createdLte := time.Now() // string |  (optional)
    group := []string{"Inner_example"} // []string | Tenant group (slug) (optional)
    groupN := []string{"Inner_example"} // []string | Tenant group (slug) (optional)
    groupId := []string{"Inner_example"} // []string | Tenant group (ID) (optional)
    groupIdN := []string{"Inner_example"} // []string | Tenant group (ID) (optional)
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
    lastUpdated := time.Now() // time.Time |  (optional)
    lastUpdatedGte := time.Now() // time.Time |  (optional)
    lastUpdatedLte := time.Now() // time.Time |  (optional)
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
    slug := []string{"Inner_example"} // []string |  (optional)
    slugIc := []string{"Inner_example"} // []string |  (optional)
    slugIe := []string{"Inner_example"} // []string |  (optional)
    slugIew := []string{"Inner_example"} // []string |  (optional)
    slugIre := []string{"Inner_example"} // []string |  (optional)
    slugIsw := []string{"Inner_example"} // []string |  (optional)
    slugN := []string{"Inner_example"} // []string |  (optional)
    slugNic := []string{"Inner_example"} // []string |  (optional)
    slugNie := []string{"Inner_example"} // []string |  (optional)
    slugNiew := []string{"Inner_example"} // []string |  (optional)
    slugNire := []string{"Inner_example"} // []string |  (optional)
    slugNisw := []string{"Inner_example"} // []string |  (optional)
    slugNre := []string{"Inner_example"} // []string |  (optional)
    slugRe := []string{"Inner_example"} // []string |  (optional)
    tag := []string{"Inner_example"} // []string |  (optional)
    tagN := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsList(context.Background()).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdIc(idIc).IdIe(idIe).IdIew(idIew).IdIre(idIre).IdIsw(idIsw).IdN(idN).IdNic(idNic).IdNie(idNie).IdNiew(idNiew).IdNire(idNire).IdNisw(idNisw).IdNre(idNre).IdRe(idRe).LastUpdated(lastUpdated).LastUpdatedGte(lastUpdatedGte).LastUpdatedLte(lastUpdatedLte).Limit(limit).Name(name).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIre(nameIre).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNire(nameNire).NameNisw(nameNisw).NameNre(nameNre).NameRe(nameRe).Offset(offset).Q(q).Slug(slug).SlugIc(slugIc).SlugIe(slugIe).SlugIew(slugIew).SlugIre(slugIre).SlugIsw(slugIsw).SlugN(slugN).SlugNic(slugNic).SlugNie(slugNie).SlugNiew(slugNiew).SlugNire(slugNire).SlugNisw(slugNisw).SlugNre(slugNre).SlugRe(slugRe).Tag(tag).TagN(tagN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantsList`: PaginatedTenantList
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created** | **string** |  | 
 **createdGte** | **string** |  | 
 **createdLte** | **string** |  | 
 **group** | **[]string** | Tenant group (slug) | 
 **groupN** | **[]string** | Tenant group (slug) | 
 **groupId** | **[]string** | Tenant group (ID) | 
 **groupIdN** | **[]string** | Tenant group (ID) | 
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
 **lastUpdated** | **time.Time** |  | 
 **lastUpdatedGte** | **time.Time** |  | 
 **lastUpdatedLte** | **time.Time** |  | 
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
 **slug** | **[]string** |  | 
 **slugIc** | **[]string** |  | 
 **slugIe** | **[]string** |  | 
 **slugIew** | **[]string** |  | 
 **slugIre** | **[]string** |  | 
 **slugIsw** | **[]string** |  | 
 **slugN** | **[]string** |  | 
 **slugNic** | **[]string** |  | 
 **slugNie** | **[]string** |  | 
 **slugNiew** | **[]string** |  | 
 **slugNire** | **[]string** |  | 
 **slugNisw** | **[]string** |  | 
 **slugNre** | **[]string** |  | 
 **slugRe** | **[]string** |  | 
 **tag** | **[]string** |  | 
 **tagN** | **[]string** |  | 

### Return type

[**PaginatedTenantList**](PaginatedTenantList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsPartialUpdate

> Tenant TenancyTenantsPartialUpdate(ctx, id).PatchedWritableTenant(patchedWritableTenant).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant.
    patchedWritableTenant := *openapiclient.NewPatchedWritableTenant() // PatchedWritableTenant |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsPartialUpdate(context.Background(), id).PatchedWritableTenant(patchedWritableTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantsPartialUpdate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWritableTenant** | [**PatchedWritableTenant**](PatchedWritableTenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsRetrieve

> Tenant TenancyTenantsRetrieve(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantsRetrieve`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenancyTenantsUpdate

> Tenant TenancyTenantsUpdate(ctx, id).WritableTenant(writableTenant).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this tenant.
    writableTenant := *openapiclient.NewWritableTenant("Id_example", "Url_example", "Name_example", time.Now(), time.Now(), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), map[string]interface{}{"key": interface{}(123)}, "Display_example") // WritableTenant | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenancyApi.TenancyTenantsUpdate(context.Background(), id).WritableTenant(writableTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenancyApi.TenancyTenantsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenancyTenantsUpdate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenancyApi.TenancyTenantsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenancyTenantsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writableTenant** | [**WritableTenant**](WritableTenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

