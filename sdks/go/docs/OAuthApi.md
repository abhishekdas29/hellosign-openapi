# dropboxsign\OAuthApi

All URIs are relative to *https://api.hellosign.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OauthTokenGenerate**](OAuthApi.md#OauthTokenGenerate) | **Post** /oauth/token | OAuth Token Generate
[**OauthTokenRefresh**](OAuthApi.md#OauthTokenRefresh) | **Post** /oauth/token?refresh | OAuth Token Refresh



## OauthTokenGenerate

> OAuthTokenResponse OauthTokenGenerate(ctx).OAuthTokenGenerateRequest(oAuthTokenGenerateRequest).Execute()

OAuth Token Generate



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
    oAuthTokenGenerateRequest := *openapiclient.NewOAuthTokenGenerateRequest("ClientId_example", "ClientSecret_example", "Code_example", "GrantType_example", "State_example") // OAuthTokenGenerateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OAuthApi.OauthTokenGenerate(context.Background()).OAuthTokenGenerateRequest(oAuthTokenGenerateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.OauthTokenGenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OauthTokenGenerate`: OAuthTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `OAuthApi.OauthTokenGenerate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauthTokenGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oAuthTokenGenerateRequest** | [**OAuthTokenGenerateRequest**](OAuthTokenGenerateRequest.md) |  | 

### Return type

[**OAuthTokenResponse**](OAuthTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthTokenRefresh

> OAuthTokenResponse OauthTokenRefresh(ctx).OAuthTokenRefreshRequest(oAuthTokenRefreshRequest).Execute()

OAuth Token Refresh



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
    oAuthTokenRefreshRequest := *openapiclient.NewOAuthTokenRefreshRequest("GrantType_example", "RefreshToken_example") // OAuthTokenRefreshRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OAuthApi.OauthTokenRefresh(context.Background()).OAuthTokenRefreshRequest(oAuthTokenRefreshRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthApi.OauthTokenRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OauthTokenRefresh`: OAuthTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `OAuthApi.OauthTokenRefresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauthTokenRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oAuthTokenRefreshRequest** | [**OAuthTokenRefreshRequest**](OAuthTokenRefreshRequest.md) |  | 

### Return type

[**OAuthTokenResponse**](OAuthTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

