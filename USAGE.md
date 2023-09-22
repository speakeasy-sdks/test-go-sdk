<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	testgosdk "github.com/speakeasy-sdks/test-go-sdk"
	"github.com/speakeasy-sdks/test-go-sdk/pkg/models/shared"
)

func main() {
    s := testgosdk.New()

    ctx := context.Background()
    res, err := s.CreateUserv1(ctx, shared.UserInput{
        Country: "Malta",
        Email: "Micheal_Sporer@yahoo.com",
        Firstname: "Karley",
        Lastname: "Stamm",
        Nickname: "vel",
        Password: "error",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->