# dropboxsign\TemplateApi

All URIs are relative to *https://api.hellosign.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplateAddUser**](TemplateApi.md#TemplateAddUser) | **Post** /template/add_user/{template_id} | Add User to Template
[**TemplateCreateEmbedded**](TemplateApi.md#TemplateCreateEmbedded) | **Post** /template/create_embedded | Create Embedded Template
[**TemplateCreateEmbeddedDraft**](TemplateApi.md#TemplateCreateEmbeddedDraft) | **Post** /template/create_embedded_draft | Create Embedded Template Draft
[**TemplateDelete**](TemplateApi.md#TemplateDelete) | **Post** /template/delete/{template_id} | Delete Template
[**TemplateFiles**](TemplateApi.md#TemplateFiles) | **Get** /template/files/{template_id} | Get Template Files
[**TemplateFilesAsDataUri**](TemplateApi.md#TemplateFilesAsDataUri) | **Get** /template/files_as_data_uri/{template_id} | Get Template Files as Data Uri
[**TemplateFilesAsFileUrl**](TemplateApi.md#TemplateFilesAsFileUrl) | **Get** /template/files_as_file_url/{template_id} | Get Template Files as File Url
[**TemplateGet**](TemplateApi.md#TemplateGet) | **Get** /template/{template_id} | Get Template
[**TemplateList**](TemplateApi.md#TemplateList) | **Get** /template/list | List Templates
[**TemplateRemoveUser**](TemplateApi.md#TemplateRemoveUser) | **Post** /template/remove_user/{template_id} | Remove User from Template
[**TemplateUpdateFiles**](TemplateApi.md#TemplateUpdateFiles) | **Post** /template/update_files/{template_id} | Update Template Files



## TemplateAddUser

> TemplateGetResponse TemplateAddUser(ctx, templateId).TemplateAddUserRequest(templateAddUserRequest).Execute()

Add User to Template



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
    templateId := "f57db65d3f933b5316d398057a36176831451a35" // string | The id of the Template to give the Account access to.
    templateAddUserRequest := *openapiclient.NewTemplateAddUserRequest() // TemplateAddUserRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateApi.TemplateAddUser(context.Background(), templateId).TemplateAddUserRequest(templateAddUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateAddUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateAddUser`: TemplateGetResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateAddUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The id of the Template to give the Account access to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateAddUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateAddUserRequest** | [**TemplateAddUserRequest**](TemplateAddUserRequest.md) |  | 

### Return type

[**TemplateGetResponse**](TemplateGetResponse.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateCreateEmbedded

> TemplateCreateEmbeddedResponse TemplateCreateEmbedded(ctx).TemplateCreateEmbeddedRequest(templateCreateEmbeddedRequest).Execute()

Create Embedded Template



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
    templateCreateEmbeddedRequest := *openapiclient.NewTemplateCreateEmbeddedRequest("ClientId_example") // TemplateCreateEmbeddedRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateApi.TemplateCreateEmbedded(context.Background()).TemplateCreateEmbeddedRequest(templateCreateEmbeddedRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateCreateEmbedded``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateCreateEmbedded`: TemplateCreateEmbeddedResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateCreateEmbedded`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplateCreateEmbeddedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateCreateEmbeddedRequest** | [**TemplateCreateEmbeddedRequest**](TemplateCreateEmbeddedRequest.md) |  | 

### Return type

[**TemplateCreateEmbeddedResponse**](TemplateCreateEmbeddedResponse.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateCreateEmbeddedDraft

> TemplateCreateEmbeddedDraftResponse TemplateCreateEmbeddedDraft(ctx).TemplateCreateEmbeddedDraftRequest(templateCreateEmbeddedDraftRequest).Execute()

Create Embedded Template Draft



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
    templateCreateEmbeddedDraftRequest := *openapiclient.NewTemplateCreateEmbeddedDraftRequest("ClientId_example") // TemplateCreateEmbeddedDraftRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateApi.TemplateCreateEmbeddedDraft(context.Background()).TemplateCreateEmbeddedDraftRequest(templateCreateEmbeddedDraftRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateCreateEmbeddedDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateCreateEmbeddedDraft`: TemplateCreateEmbeddedDraftResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateCreateEmbeddedDraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplateCreateEmbeddedDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateCreateEmbeddedDraftRequest** | [**TemplateCreateEmbeddedDraftRequest**](TemplateCreateEmbeddedDraftRequest.md) |  | 

### Return type

[**TemplateCreateEmbeddedDraftResponse**](TemplateCreateEmbeddedDraftResponse.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateDelete

> TemplateDelete(ctx, templateId).Execute()

Delete Template



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
    templateId := "f57db65d3f933b5316d398057a36176831451a35" // string | The id of the Template to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateApi.TemplateDelete(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The id of the Template to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateFiles

> *os.File TemplateFiles(ctx, templateId).FileType(fileType).Execute()

Get Template Files



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
    templateId := "f57db65d3f933b5316d398057a36176831451a35" // string | The id of the template files to retrieve.
    fileType := "fileType_example" // string | Set to `pdf` for a single merged document or `zip` for a collection of individual documents. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateApi.TemplateFiles(context.Background(), templateId).FileType(fileType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFiles`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The id of the template files to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileType** | **string** | Set to &#x60;pdf&#x60; for a single merged document or &#x60;zip&#x60; for a collection of individual documents. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateFilesAsDataUri

> FileResponseDataUri TemplateFilesAsDataUri(ctx, templateId).Execute()

Get Template Files as Data Uri



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
    templateId := "f57db65d3f933b5316d398057a36176831451a35" // string | The id of the template files to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateApi.TemplateFilesAsDataUri(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateFilesAsDataUri``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFilesAsDataUri`: FileResponseDataUri
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateFilesAsDataUri`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The id of the template files to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFilesAsDataUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileResponseDataUri**](FileResponseDataUri.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateFilesAsFileUrl

> FileResponse TemplateFilesAsFileUrl(ctx, templateId).Execute()

Get Template Files as File Url



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
    templateId := "f57db65d3f933b5316d398057a36176831451a35" // string | The id of the template files to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateApi.TemplateFilesAsFileUrl(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateFilesAsFileUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFilesAsFileUrl`: FileResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateFilesAsFileUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The id of the template files to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFilesAsFileUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileResponse**](FileResponse.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateGet

> TemplateGetResponse TemplateGet(ctx, templateId).Execute()

Get Template



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
    templateId := "f57db65d3f933b5316d398057a36176831451a35" // string | The id of the Template to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateApi.TemplateGet(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateGet`: TemplateGetResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The id of the Template to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateGetResponse**](TemplateGetResponse.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateList

> TemplateListResponse TemplateList(ctx).AccountId(accountId).Page(page).PageSize(pageSize).Query(query).Execute()

List Templates



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
    accountId := "accountId_example" // string | Which account to return Templates for. Must be a team member. Use `all` to indicate all team members. Defaults to your account. (optional)
    page := int32(56) // int32 | Which page number of the Template List to return. Defaults to `1`. (optional) (default to 1)
    pageSize := int32(56) // int32 | Number of objects to be returned per page. Must be between `1` and `100`. Default is `20`. (optional) (default to 20)
    query := "query_example" // string | String that includes search terms and/or fields to be used to filter the Template objects. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateApi.TemplateList(context.Background()).AccountId(accountId).Page(page).PageSize(pageSize).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateList`: TemplateListResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Which account to return Templates for. Must be a team member. Use &#x60;all&#x60; to indicate all team members. Defaults to your account. | 
 **page** | **int32** | Which page number of the Template List to return. Defaults to &#x60;1&#x60;. | [default to 1]
 **pageSize** | **int32** | Number of objects to be returned per page. Must be between &#x60;1&#x60; and &#x60;100&#x60;. Default is &#x60;20&#x60;. | [default to 20]
 **query** | **string** | String that includes search terms and/or fields to be used to filter the Template objects. | 

### Return type

[**TemplateListResponse**](TemplateListResponse.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateRemoveUser

> TemplateGetResponse TemplateRemoveUser(ctx, templateId).TemplateRemoveUserRequest(templateRemoveUserRequest).Execute()

Remove User from Template



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
    templateId := "f57db65d3f933b5316d398057a36176831451a35" // string | The id of the Template to remove the Account's access to.
    templateRemoveUserRequest := *openapiclient.NewTemplateRemoveUserRequest() // TemplateRemoveUserRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateApi.TemplateRemoveUser(context.Background(), templateId).TemplateRemoveUserRequest(templateRemoveUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateRemoveUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateRemoveUser`: TemplateGetResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateRemoveUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The id of the Template to remove the Account&#39;s access to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateRemoveUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateRemoveUserRequest** | [**TemplateRemoveUserRequest**](TemplateRemoveUserRequest.md) |  | 

### Return type

[**TemplateGetResponse**](TemplateGetResponse.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateUpdateFiles

> TemplateUpdateFilesResponse TemplateUpdateFiles(ctx, templateId).TemplateUpdateFilesRequest(templateUpdateFilesRequest).Execute()

Update Template Files



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
    templateId := "f57db65d3f933b5316d398057a36176831451a35" // string | The ID of the template whose files to update.
    templateUpdateFilesRequest := *openapiclient.NewTemplateUpdateFilesRequest() // TemplateUpdateFilesRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateApi.TemplateUpdateFiles(context.Background(), templateId).TemplateUpdateFilesRequest(templateUpdateFilesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApi.TemplateUpdateFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateUpdateFiles`: TemplateUpdateFilesResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateApi.TemplateUpdateFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the template whose files to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateUpdateFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateUpdateFilesRequest** | [**TemplateUpdateFilesRequest**](TemplateUpdateFilesRequest.md) |  | 

### Return type

[**TemplateUpdateFilesResponse**](TemplateUpdateFilesResponse.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
