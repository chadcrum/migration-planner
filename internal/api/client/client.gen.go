// Package client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	. "github.com/kubev2v/migration-planner/api/v1alpha1"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// DeleteSources request
	DeleteSources(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListSources request
	ListSources(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateSourceWithBody request with any body
	CreateSourceWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateSource(ctx context.Context, body CreateSourceJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteSource request
	DeleteSource(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ReadSource request
	ReadSource(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetSourceImage request
	GetSourceImage(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) DeleteSources(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteSourcesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListSources(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListSourcesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateSourceWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateSourceRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateSource(ctx context.Context, body CreateSourceJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateSourceRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteSource(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteSourceRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ReadSource(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewReadSourceRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetSourceImage(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSourceImageRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewDeleteSourcesRequest generates requests for DeleteSources
func NewDeleteSourcesRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/sources")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListSourcesRequest generates requests for ListSources
func NewListSourcesRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/sources")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateSourceRequest calls the generic CreateSource builder with application/json body
func NewCreateSourceRequest(server string, body CreateSourceJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateSourceRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateSourceRequestWithBody generates requests for CreateSource with any type of body
func NewCreateSourceRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/sources")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteSourceRequest generates requests for DeleteSource
func NewDeleteSourceRequest(server string, id openapi_types.UUID) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/sources/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewReadSourceRequest generates requests for ReadSource
func NewReadSourceRequest(server string, id openapi_types.UUID) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/sources/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetSourceImageRequest generates requests for GetSourceImage
func NewGetSourceImageRequest(server string, id openapi_types.UUID) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/sources/%s/image", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// DeleteSourcesWithResponse request
	DeleteSourcesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*DeleteSourcesResponse, error)

	// ListSourcesWithResponse request
	ListSourcesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListSourcesResponse, error)

	// CreateSourceWithBodyWithResponse request with any body
	CreateSourceWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateSourceResponse, error)

	CreateSourceWithResponse(ctx context.Context, body CreateSourceJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateSourceResponse, error)

	// DeleteSourceWithResponse request
	DeleteSourceWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteSourceResponse, error)

	// ReadSourceWithResponse request
	ReadSourceWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*ReadSourceResponse, error)

	// GetSourceImageWithResponse request
	GetSourceImageWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetSourceImageResponse, error)
}

type DeleteSourcesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Status
	JSON401      *Error
}

// Status returns HTTPResponse.Status
func (r DeleteSourcesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteSourcesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListSourcesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SourceList
	JSON401      *Error
}

// Status returns HTTPResponse.Status
func (r ListSourcesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListSourcesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateSourceResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *Source
	JSON400      *Error
	JSON401      *Error
}

// Status returns HTTPResponse.Status
func (r CreateSourceResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateSourceResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteSourceResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Source
	JSON400      *Error
	JSON401      *Error
	JSON404      *Error
}

// Status returns HTTPResponse.Status
func (r DeleteSourceResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteSourceResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ReadSourceResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Source
	JSON400      *Error
	JSON401      *Error
	JSON404      *Error
}

// Status returns HTTPResponse.Status
func (r ReadSourceResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ReadSourceResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetSourceImageResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *Error
	JSON401      *Error
	JSON404      *Error
	JSON500      *Error
}

// Status returns HTTPResponse.Status
func (r GetSourceImageResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSourceImageResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// DeleteSourcesWithResponse request returning *DeleteSourcesResponse
func (c *ClientWithResponses) DeleteSourcesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*DeleteSourcesResponse, error) {
	rsp, err := c.DeleteSources(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteSourcesResponse(rsp)
}

// ListSourcesWithResponse request returning *ListSourcesResponse
func (c *ClientWithResponses) ListSourcesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListSourcesResponse, error) {
	rsp, err := c.ListSources(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListSourcesResponse(rsp)
}

// CreateSourceWithBodyWithResponse request with arbitrary body returning *CreateSourceResponse
func (c *ClientWithResponses) CreateSourceWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateSourceResponse, error) {
	rsp, err := c.CreateSourceWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateSourceResponse(rsp)
}

func (c *ClientWithResponses) CreateSourceWithResponse(ctx context.Context, body CreateSourceJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateSourceResponse, error) {
	rsp, err := c.CreateSource(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateSourceResponse(rsp)
}

// DeleteSourceWithResponse request returning *DeleteSourceResponse
func (c *ClientWithResponses) DeleteSourceWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteSourceResponse, error) {
	rsp, err := c.DeleteSource(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteSourceResponse(rsp)
}

// ReadSourceWithResponse request returning *ReadSourceResponse
func (c *ClientWithResponses) ReadSourceWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*ReadSourceResponse, error) {
	rsp, err := c.ReadSource(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseReadSourceResponse(rsp)
}

// GetSourceImageWithResponse request returning *GetSourceImageResponse
func (c *ClientWithResponses) GetSourceImageWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetSourceImageResponse, error) {
	rsp, err := c.GetSourceImage(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSourceImageResponse(rsp)
}

// ParseDeleteSourcesResponse parses an HTTP response from a DeleteSourcesWithResponse call
func ParseDeleteSourcesResponse(rsp *http.Response) (*DeleteSourcesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteSourcesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Status
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseListSourcesResponse parses an HTTP response from a ListSourcesWithResponse call
func ParseListSourcesResponse(rsp *http.Response) (*ListSourcesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListSourcesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SourceList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseCreateSourceResponse parses an HTTP response from a CreateSourceWithResponse call
func ParseCreateSourceResponse(rsp *http.Response) (*CreateSourceResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateSourceResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest Source
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseDeleteSourceResponse parses an HTTP response from a DeleteSourceWithResponse call
func ParseDeleteSourceResponse(rsp *http.Response) (*DeleteSourceResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteSourceResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Source
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseReadSourceResponse parses an HTTP response from a ReadSourceWithResponse call
func ParseReadSourceResponse(rsp *http.Response) (*ReadSourceResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ReadSourceResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Source
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseGetSourceImageResponse parses an HTTP response from a GetSourceImageWithResponse call
func ParseGetSourceImageResponse(rsp *http.Response) (*GetSourceImageResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetSourceImageResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
