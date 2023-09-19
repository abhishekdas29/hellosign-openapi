# dropboxsign\AccountApi

All URIs are relative to *https://api.hellosign.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountCreate**](AccountApi.md#AccountCreate) | **Post** /account/create | Create Account
[**AccountGet**](AccountApi.md#AccountGet) | **Get** /account | Get Account
[**AccountUpdate**](AccountApi.md#AccountUpdate) | **Put** /account | Update Account
[**AccountVerify**](AccountApi.md#AccountVerify) | **Post** /account/verify | Verify Account



## AccountCreate

> AccountCreateResponse AccountCreate(ctx).AccountCreateRequest(accountCreateRequest).Execute()

Create Account



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
    accountCreateRequest := *openapiclient.NewAccountCreateRequest("EmailAddress_example") // AccountCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.AccountCreate(context.Background()).AccountCreateRequest(accountCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountCreate`: AccountCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountCreateRequest** | [**AccountCreateRequest**](AccountCreateRequest.md) |  | 

### Return type

[**AccountCreateResponse**](AccountCreateResponse.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountGet

> AccountGetResponse AccountGet(ctx).AccountId(accountId).EmailAddress(emailAddress).Execute()

Get Account



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
    accountId := "accountId_example" // string | `account_id` or `email_address` is required. If both are provided, the account id prevails.  The ID of the Account. (optional)
    emailAddress := "emailAddress_example" // string | `account_id` or `email_address` is required, If both are provided, the account id prevails.  The email address of the Account. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.AccountGet(context.Background()).AccountId(accountId).EmailAddress(emailAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountGet`: AccountGetResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | &#x60;account_id&#x60; or &#x60;email_address&#x60; is required. If both are provided, the account id prevails.  The ID of the Account. | 
 **emailAddress** | **string** | &#x60;account_id&#x60; or &#x60;email_address&#x60; is required, If both are provided, the account id prevails.  The email address of the Account. | 

### Return type

[**AccountGetResponse**](AccountGetResponse.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountUpdate

> AccountGetResponse AccountUpdate(ctx).AccountUpdateRequest(accountUpdateRequest).Execute()

Update Account



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
    accountUpdateRequest := *openapiclient.NewAccountUpdateRequest() // AccountUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.AccountUpdate(context.Background()).AccountUpdateRequest(accountUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountUpdate`: AccountGetResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountUpdateRequest** | [**AccountUpdateRequest**](AccountUpdateRequest.md) |  | 

### Return type

[**AccountGetResponse**](AccountGetResponse.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountVerify

> AccountVerifyResponse AccountVerify(ctx).AccountVerifyRequest(accountVerifyRequest).Execute()

Verify Account



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
    accountVerifyRequest := *openapiclient.NewAccountVerifyRequest("EmailAddress_example") // AccountVerifyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.AccountVerify(context.Background()).AccountVerifyRequest(accountVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountVerify`: AccountVerifyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountVerify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountVerifyRequest** | [**AccountVerifyRequest**](AccountVerifyRequest.md) |  | 

### Return type

[**AccountVerifyResponse**](AccountVerifyResponse.md)

### Authorization

[api_key](../README.md#api_key), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

