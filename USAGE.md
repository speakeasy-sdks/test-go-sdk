<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	testgosdk "github.com/speakeasy-sdks/test-go-sdk/v3"
	"github.com/speakeasy-sdks/test-go-sdk/v3/pkg/models/shared"
	"log"
)

func main() {
	s := testgosdk.New()

	ctx := context.Background()
	res, err := s.CreateUserv1(ctx, shared.UserInput{
		Country:   "Benin",
		Email:     "Della67@yahoo.com",
		Firstname: "Enrique",
		Lastname:  "Ernser",
		Nickname:  "<value>",
		Password:  "TOXx1B29WwlhtAA",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.User != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->