package varorequest

import (
	"net/http"
)

// This is request Input.
type RequestInput struct {
	URL        *string
	Headers    map[string]string
	RawHeaders http.Header
	Method     *string
	TimeOut    int
	NoVerify   bool
	BodyBytes  []byte
	Json       *interface{}
	Client     *http.Client
}

// The request builder struct for method chaining
type RequestBuilder struct {
	RequestInput
}
