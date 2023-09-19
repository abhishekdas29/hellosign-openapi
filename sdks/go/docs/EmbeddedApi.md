# dropboxsign\EmbeddedApi

All URIs are relative to *https://api.hellosign.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmbeddedEditUrl**](EmbeddedApi.md#EmbeddedEditUrl) | **Post** /embedded/edit_url/{template_id} | Get Embedded Template Edit URL
[**EmbeddedSignUrl**](EmbeddedApi.md#EmbeddedSignUrl) | **Get** /embedded/sign_url/{signature_id} | Get Embedded Sign URL



## EmbeddedEditUrl

> EmbeddedEditUrlResponse EmbeddedEditUrl(ctx, templateId).EmbeddedEditUrlRequest(embeddedEditUrlRequest).Execute()

Get Embedded Template Edit URL



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
    templateId := "f57db65d3f933b5316d398057a36176831451a35" // string | The id of the template to edit.
    embeddedEditUrlRequest := *openapiclient.NewEmbeddedEditUrlRequest() // EmbeddedEditUrlRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EmbeddedApi.EmbeddedEditUrl(context.Background(), templateId).EmbeddedEditUrlRequest(embeddedEditUrlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmbeddedApi.EmbeddedEditUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmbeddedEditUrl`: EmbeddedEditUrlResponse
    fmt.Fprintf(os.Stdout, "Response from `EmbeddedApi.EmbeddedEditUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The id of the template to edit. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmbeddedEditUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **embeddedEditUrlRequest** | [**EmbeddedEditUrlRequest**](EmbeddedEditUrlRequest.md) |  | 

### Return type

[**EmbeddedEditUrlResponse**](EmbeddedEditUrlResponse.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmbeddedSignUrl

> EmbeddedSignUrlResponse EmbeddedSignUrl(ctx, signatureId).Execute()

Get Embedded Sign URL



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
    signatureId := "50e3542f738adfa7ddd4cbd4c00d2a8ab6e4194b" // string | The id of the signature to get a signature url for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EmbeddedApi.EmbeddedSignUrl(context.Background(), signatureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmbeddedApi.EmbeddedSignUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmbeddedSignUrl`: EmbeddedSignUrlResponse
    fmt.Fprintf(os.Stdout, "Response from `EmbeddedApi.EmbeddedSignUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signatureId** | **string** | The id of the signature to get a signature url for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmbeddedSignUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmbeddedSignUrlResponse**](EmbeddedSignUrlResponse.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

