package varorequest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func New() RequestBuilder {
	client := &http.Client{}
	r := RequestInput{
		Client: client,
	}

	return RequestBuilder{
		RequestInput: r,
	}
}

// The purpose for this function is to make proxying request easier. Insteading of creating a new http client, we can use the request's http.Client to save time to mimick request.
func NewWithClient(c *http.Client) RequestBuilder {
	var client *http.Client
	if c == nil {
		client = &http.Client{}
	} else {
		client = c
	}
	r := RequestInput{
		Client: client,
	}
	return RequestBuilder{
		RequestInput: r,
	}
}
func (r *RequestBuilder) SetURL(b string) *RequestBuilder {
	r.RequestInput.URL = &b
	return r
}
func (r *RequestBuilder) SetBodyBytes(b []byte) *RequestBuilder {
	r.RequestInput.BodyBytes = b
	return r
}
func (r *RequestBuilder) SetHeaders(h map[string]string) *RequestBuilder {
	r.RequestInput.Headers = h
	return r
}
func (r *RequestBuilder) SetRawHeaders(h http.Header) *RequestBuilder {
	r.RequestInput.RawHeaders = h
	return r
}
func (r *RequestBuilder) SetTimeOut(h int) *RequestBuilder {
	//	r.RequestInput.Client.Timeout = time.Duration(h) * time.Second
	r.RequestInput.TimeOut = h
	return r
}
func (r *RequestBuilder) SetMethod(m string) *RequestBuilder {
	r.RequestInput.Method = &m
	return r
}
func (r *RequestBuilder) NoVerify() *RequestBuilder {
	r.RequestInput.NoVerify = true
	return r
}
func (r *RequestBuilder) SetJson(j interface{}) *RequestBuilder {
	r.RequestInput.Json = &j
	return r
}
func (r *RequestBuilder) Do() (*http.Response, error) {
	var client *http.Client
	if r.RequestInput.Client == nil {
		client = &http.Client{}
	} else {
		client = r.RequestInput.Client
	}
	var (
		h     *http.Response
		ebody *bytes.Reader
	)
	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = r.RequestInput.NoVerify
	err := r._check()
	if err != nil {
		return h, err
	}

	// check if json exist
	if r.Json != nil {
		eJson, err := json.Marshal(*r.RequestInput.Json)
		if err != nil {
			return h, err
		}
		ebody = bytes.NewReader(eJson)
	} else {
		ebody = bytes.NewReader([]byte(""))
	}
	// check if bodyBytes exist bodybytes overwrite everything
	if r.BodyBytes != nil {
		ebody = bytes.NewReader(r.BodyBytes)
	}
	req, err := http.NewRequest(*r.RequestInput.Method, *r.RequestInput.URL, ebody)
	if err != nil {
		return h, err
	}

	if r.RequestInput.Headers != nil {
		for k, v := range r.RequestInput.Headers {
			req.Header.Set(k, v)
		}
	}
	if r.RequestInput.RawHeaders != nil {
		req.Header = r.RequestInput.RawHeaders
	}

	client.Transport = &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	// Set Timeout
	client.Timeout = time.Duration(r.RequestInput.TimeOut) * time.Second
	h, err = client.Do(req)
	if err != nil {
		return h, err
	}
	return h, nil
}

func (r *RequestBuilder) _check() error {
	// make GET as default
	if r.RequestInput.Method == nil {
		method := "GET"
		r.RequestInput.Method = &method
	}

	// check if url is valid
	_, err := url.ParseRequestURI(*r.RequestInput.URL)
	if err != nil {
		return err
	}
	if r.RequestInput.URL == nil {
		return fmt.Errorf("url not valid")
	}
	return nil
}
