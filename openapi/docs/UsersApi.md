# \UsersApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /v1/users | Create a new user
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /v1/users/{userId} | Delete existing user by id
[**GenerateClientSecret**](UsersApi.md#GenerateClientSecret) | **Put** /v1/users/client/secret | Generate new client secret
[**ResendEmailVerification**](UsersApi.md#ResendEmailVerification) | **Post** /v1/users/{userId}/resend-email-verification | Resend email verification
[**ResetPassword**](UsersApi.md#ResetPassword) | **Post** /v1/users/{userId}/reset-password | Send reset password email
[**RetrieveUser**](UsersApi.md#RetrieveUser) | **Get** /v1/users/{userId} | Get specific user by id
[**RetrieveUserClient**](UsersApi.md#RetrieveUserClient) | **Get** /v1/users/client | Get client
[**RetrieveUserIsPasswordSet**](UsersApi.md#RetrieveUserIsPasswordSet) | **Get** /v1/users/is-password-set | Get user is password set status
[**RetrieveUserList**](UsersApi.md#RetrieveUserList) | **Get** /v1/users | List users
[**UpdateUser**](UsersApi.md#UpdateUser) | **Patch** /v1/users/{userId} | Update specific user by id



## CreateUser

> CreateUserResponse CreateUser(ctx).XRequestId(xRequestId).CreateUserRequest(createUserRequest).XTraceId(xTraceId).Execute()

Create a new user



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
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    createUserRequest := *openapiclient.NewCreateUserRequest("john.doe@example.com", false, false, "de") // CreateUserRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.CreateUser(context.Background()).XRequestId(xRequestId).CreateUserRequest(createUserRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: CreateUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **createUserRequest** | [**CreateUserRequest**](CreateUserRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CreateUserResponse**](CreateUserResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Delete existing user by id



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
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    userId := "6cdf5968-f9fe-4192-97c2-f349e813c5e8" // string | The identifier of the user.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.DeleteUser(context.Background(), userId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateClientSecret

> GenerateClientSecretResponse GenerateClientSecret(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Generate new client secret



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
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GenerateClientSecret(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GenerateClientSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateClientSecret`: GenerateClientSecretResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GenerateClientSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**GenerateClientSecretResponse**](GenerateClientSecretResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendEmailVerification

> ResendEmailVerification(ctx, userId).XRequestId(xRequestId).XTraceId(xTraceId).RedirectUrl(redirectUrl).Execute()

Resend email verification



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
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    userId := "6cdf5968-f9fe-4192-97c2-f349e813c5e8" // string | The identifier of the user.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    redirectUrl := "https://test.contabo.de" // string | The redirect url used for email verification (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.ResendEmailVerification(context.Background(), userId).XRequestId(xRequestId).XTraceId(xTraceId).RedirectUrl(redirectUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ResendEmailVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendEmailVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **redirectUrl** | **string** | The redirect url used for email verification | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> ResetPassword(ctx, userId).XRequestId(xRequestId).XTraceId(xTraceId).RedirectUrl(redirectUrl).Execute()

Send reset password email



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
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    userId := "6cdf5968-f9fe-4192-97c2-f349e813c5e8" // string | The identifier of the user.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    redirectUrl := "https://test.contabo.de" // string | The redirect url used for resetting password (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.ResetPassword(context.Background(), userId).XRequestId(xRequestId).XTraceId(xTraceId).RedirectUrl(redirectUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ResetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **redirectUrl** | **string** | The redirect url used for resetting password | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveUser

> FindUserResponse RetrieveUser(ctx, userId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get specific user by id



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
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    userId := "6cdf5968-f9fe-4192-97c2-f349e813c5e8" // string | The identifier of the user.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.RetrieveUser(context.Background(), userId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.RetrieveUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveUser`: FindUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.RetrieveUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**FindUserResponse**](FindUserResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveUserClient

> FindClientResponse RetrieveUserClient(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get client



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
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.RetrieveUserClient(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.RetrieveUserClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveUserClient`: FindClientResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.RetrieveUserClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveUserClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**FindClientResponse**](FindClientResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveUserIsPasswordSet

> FindUserIsPasswordSetResponse RetrieveUserIsPasswordSet(ctx).XRequestId(xRequestId).XTraceId(xTraceId).UserId(userId).Execute()

Get user is password set status



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
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    userId := "6cdf5968-f9fe-4192-97c2-f349e813c5e8" // string | The user ID for checking if password is set for him (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.RetrieveUserIsPasswordSet(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.RetrieveUserIsPasswordSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveUserIsPasswordSet`: FindUserIsPasswordSetResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.RetrieveUserIsPasswordSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveUserIsPasswordSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **userId** | **string** | The user ID for checking if password is set for him | 

### Return type

[**FindUserIsPasswordSetResponse**](FindUserIsPasswordSetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveUserList

> ListUserResponse RetrieveUserList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Email(email).Enabled(enabled).Owner(owner).Execute()

List users



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
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    page := int64(1) // int64 | Number of page to be fetched. (optional)
    size := int64(10) // int64 | Number of elements per page. (optional)
    orderBy := []string{"Inner_example"} // []string | Specify fields and ordering (ASC for ascending, DESC for descending) in following format `field:ASC|DESC`. (optional)
    email := "john.doe@example.com" // string | Filter as substring match for user emails. (optional)
    enabled := true // bool | Filter if user is enabled or not. (optional)
    owner := true // bool | Filter if user is owner or not. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.RetrieveUserList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Email(email).Enabled(enabled).Owner(owner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.RetrieveUserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveUserList`: ListUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.RetrieveUserList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **email** | **string** | Filter as substring match for user emails. | 
 **enabled** | **bool** | Filter if user is enabled or not. | 
 **owner** | **bool** | Filter if user is owner or not. | 

### Return type

[**ListUserResponse**](ListUserResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UpdateUserResponse UpdateUser(ctx, userId).XRequestId(xRequestId).UpdateUserRequest(updateUserRequest).XTraceId(xTraceId).Execute()

Update specific user by id



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
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    userId := "6cdf5968-f9fe-4192-97c2-f349e813c5e8" // string | The identifier of the user.
    updateUserRequest := *openapiclient.NewUpdateUserRequest() // UpdateUserRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UpdateUser(context.Background(), userId).XRequestId(xRequestId).UpdateUserRequest(updateUserRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: UpdateUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **updateUserRequest** | [**UpdateUserRequest**](UpdateUserRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**UpdateUserResponse**](UpdateUserResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

