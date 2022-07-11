# \GraphqlApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphqlCreate**](GraphqlApi.md#GraphqlCreate) | **Post** /graphql/ | 



## GraphqlCreate

> GraphqlCreate200Response GraphqlCreate(ctx).GraphQLAPI(graphQLAPI).Execute()





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
    graphQLAPI := *openapiclient.NewGraphQLAPI("Query_example") // GraphQLAPI | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphqlApi.GraphqlCreate(context.Background()).GraphQLAPI(graphQLAPI).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphqlApi.GraphqlCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphqlCreate`: GraphqlCreate200Response
    fmt.Fprintf(os.Stdout, "Response from `GraphqlApi.GraphqlCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphqlCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphQLAPI** | [**GraphQLAPI**](GraphQLAPI.md) |  | 

### Return type

[**GraphqlCreate200Response**](GraphqlCreate200Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

