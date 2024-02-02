// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/test-go-sdk/v3/pkg/models/shared"
	"net/http"
)

type GetUserv1Request struct {
	// Numeric ID of the user to get
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetUserv1Request) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetUserv1Response struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	User *shared.User
}

func (o *GetUserv1Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUserv1Response) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetUserv1Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUserv1Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUserv1Response) GetUser() *shared.User {
	if o == nil {
		return nil
	}
	return o.User
}
