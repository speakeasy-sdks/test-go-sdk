<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	testgosdk "github.com/speakeasy-sdks/test-go-sdk"
	"github.com/speakeasy-sdks/test-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := testgosdk.New()

	ctx := context.Background()
	res, err := s.Test.CreateUserv1(ctx, shared.UserInput{
		Country:   "Benin",
		Email:     "Della67@yahoo.com",
		Firstname: "Enrique",
		Lastname:  "Ernser",
		Nickname:  "panel",
		Password:  "OXx1B29WwlhtAAe",
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