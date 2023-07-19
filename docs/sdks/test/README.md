# Test SDK

## Overview

User API for Speakeasy template service: The Rest Template API is an API used for instrucitonal purposes.

### Available Operations

* [CreateUserv1](#createuserv1) - Create user
* [DeleteUserv1](#deleteuserv1) - Delete a user by ID
* [GetHealth](#gethealth) - Healthcheck
* [GetUserv1](#getuserv1) - Get a user by ID
* [SearchUsersv1](#searchusersv1) - Search users
* [UpdateUserv1](#updateuserv1)

## CreateUserv1

Create user

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/test-go-sdk"
	"github.com/speakeasy-sdks/test-go-sdk/pkg/models/shared"
)

func main() {
    s := test.New()

    ctx := context.Background()
    res, err := s.Test.CreateUserv1(ctx, shared.UserInput{
        Country: "Niger",
        Email: "Hunter.Gulgowski96@yahoo.com",
        Firstname: "Donny",
        Lastname: "Hoppe",
        Nickname: "molestiae",
        Password: "minus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.UserInput](../../models/shared/userinput.md)  | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateUserv1Response](../../models/operations/createuserv1response.md), error**


## DeleteUserv1

Delete a user by ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/test-go-sdk"
	"github.com/speakeasy-sdks/test-go-sdk/pkg/models/operations"
)

func main() {
    s := test.New()

    ctx := context.Background()
    res, err := s.Test.DeleteUserv1(ctx, operations.DeleteUserv1Request{
        ID: "c8796ed1-51a0-45df-82dd-f7cc78ca1ba9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Success != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteUserv1Request](../../models/operations/deleteuserv1request.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.DeleteUserv1Response](../../models/operations/deleteuserv1response.md), error**


## GetHealth

Healthcheck

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/test-go-sdk"
)

func main() {
    s := test.New()

    ctx := context.Background()
    res, err := s.Test.GetHealth(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetHealthResponse](../../models/operations/gethealthresponse.md), error**


## GetUserv1

Get a user by ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/test-go-sdk"
	"github.com/speakeasy-sdks/test-go-sdk/pkg/models/operations"
)

func main() {
    s := test.New()

    ctx := context.Background()
    res, err := s.Test.GetUserv1(ctx, operations.GetUserv1Request{
        ID: "28fc8167-42cb-4739-a059-29396fea7596",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetUserv1Request](../../models/operations/getuserv1request.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetUserv1Response](../../models/operations/getuserv1response.md), error**


## SearchUsersv1

Search users

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/test-go-sdk"
	"github.com/speakeasy-sdks/test-go-sdk/pkg/models/shared"
)

func main() {
    s := test.New()

    ctx := context.Background()
    res, err := s.Test.SearchUsersv1(ctx, shared.Filters{
        Filters: []shared.Filter{
            shared.Filter{
                Field: "quidem",
                Matchtype: "architecto",
                Value: "ipsa",
            },
            shared.Filter{
                Field: "reiciendis",
                Matchtype: "est",
                Value: "mollitia",
            },
            shared.Filter{
                Field: "laborum",
                Matchtype: "dolores",
                Value: "dolorem",
            },
            shared.Filter{
                Field: "corporis",
                Matchtype: "explicabo",
                Value: "nobis",
            },
        },
        Limit: 315428,
        Offset: 607831,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Users != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Filters](../../models/shared/filters.md)      | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.SearchUsersv1Response](../../models/operations/searchusersv1response.md), error**


## UpdateUserv1

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/test-go-sdk"
	"github.com/speakeasy-sdks/test-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/test-go-sdk/pkg/models/shared"
)

func main() {
    s := test.New()

    ctx := context.Background()
    res, err := s.Test.UpdateUserv1(ctx, operations.UpdateUserv1Request{
        UserInput: shared.UserInput{
            Country: "Guinea",
            Email: "Keyon_Batz@gmail.com",
            Firstname: "Yasmeen",
            Lastname: "Williamson",
            Nickname: "architecto",
            Password: "mollitia",
        },
        ID: "3a2fa946-7739-4251-aa52-c3f5ad019da1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UpdateUserv1Request](../../models/operations/updateuserv1request.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.UpdateUserv1Response](../../models/operations/updateuserv1response.md), error**

