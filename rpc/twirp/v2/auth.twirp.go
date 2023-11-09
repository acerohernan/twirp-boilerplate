// Code generated by protoc-gen-twirp v8.1.3, DO NOT EDIT.
// source: twirp/v2/auth.proto

package twirp

import context "context"
import fmt "fmt"
import http "net/http"
import io "io"
import json "encoding/json"
import strconv "strconv"
import strings "strings"

import protojson "google.golang.org/protobuf/encoding/protojson"
import proto "google.golang.org/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// Version compatibility assertion.
// If the constant is not defined in the package, that likely means
// the package needs to be updated to work with this generated code.
// See https://twitchtv.github.io/twirp/docs/version_matrix.html
const _ = twirp.TwirpPackageMinVersion_8_1_0

// =====================
// AuthService Interface
// =====================

type AuthService interface {
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)

	ListSessions(context.Context, *ListSessionsRequest) (*ListSessionsResponse, error)
}

// ===========================
// AuthService Protobuf Client
// ===========================

type authServiceProtobufClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewAuthServiceProtobufClient creates a Protobuf client that implements the AuthService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewAuthServiceProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) AuthService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwards compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "twirp.v2", "AuthService")
	urls := [2]string{
		serviceURL + "CreateSession",
		serviceURL + "ListSessions",
	}

	return &authServiceProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *authServiceProtobufClient) CreateSession(ctx context.Context, in *CreateSessionRequest) (*CreateSessionResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "twirp.v2")
	ctx = ctxsetters.WithServiceName(ctx, "AuthService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateSession")
	caller := c.callCreateSession
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CreateSessionRequest) (*CreateSessionResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateSessionRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateSessionRequest) when calling interceptor")
					}
					return c.callCreateSession(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateSessionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateSessionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *authServiceProtobufClient) callCreateSession(ctx context.Context, in *CreateSessionRequest) (*CreateSessionResponse, error) {
	out := new(CreateSessionResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *authServiceProtobufClient) ListSessions(ctx context.Context, in *ListSessionsRequest) (*ListSessionsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "twirp.v2")
	ctx = ctxsetters.WithServiceName(ctx, "AuthService")
	ctx = ctxsetters.WithMethodName(ctx, "ListSessions")
	caller := c.callListSessions
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *ListSessionsRequest) (*ListSessionsResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListSessionsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListSessionsRequest) when calling interceptor")
					}
					return c.callListSessions(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListSessionsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListSessionsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *authServiceProtobufClient) callListSessions(ctx context.Context, in *ListSessionsRequest) (*ListSessionsResponse, error) {
	out := new(ListSessionsResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// =======================
// AuthService JSON Client
// =======================

type authServiceJSONClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewAuthServiceJSONClient creates a JSON client that implements the AuthService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewAuthServiceJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) AuthService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwards compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "twirp.v2", "AuthService")
	urls := [2]string{
		serviceURL + "CreateSession",
		serviceURL + "ListSessions",
	}

	return &authServiceJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *authServiceJSONClient) CreateSession(ctx context.Context, in *CreateSessionRequest) (*CreateSessionResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "twirp.v2")
	ctx = ctxsetters.WithServiceName(ctx, "AuthService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateSession")
	caller := c.callCreateSession
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CreateSessionRequest) (*CreateSessionResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateSessionRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateSessionRequest) when calling interceptor")
					}
					return c.callCreateSession(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateSessionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateSessionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *authServiceJSONClient) callCreateSession(ctx context.Context, in *CreateSessionRequest) (*CreateSessionResponse, error) {
	out := new(CreateSessionResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *authServiceJSONClient) ListSessions(ctx context.Context, in *ListSessionsRequest) (*ListSessionsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "twirp.v2")
	ctx = ctxsetters.WithServiceName(ctx, "AuthService")
	ctx = ctxsetters.WithMethodName(ctx, "ListSessions")
	caller := c.callListSessions
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *ListSessionsRequest) (*ListSessionsResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListSessionsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListSessionsRequest) when calling interceptor")
					}
					return c.callListSessions(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListSessionsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListSessionsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *authServiceJSONClient) callListSessions(ctx context.Context, in *ListSessionsRequest) (*ListSessionsResponse, error) {
	out := new(ListSessionsResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ==========================
// AuthService Server Handler
// ==========================

type authServiceServer struct {
	AuthService
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
	jsonCamelCase    bool   // JSON fields are serialized as lowerCamelCase rather than keeping the original proto names
}

// NewAuthServiceServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewAuthServiceServer(svc AuthService, opts ...interface{}) TwirpServer {
	serverOpts := newServerOpts(opts)

	// Using ReadOpt allows backwards and forwards compatibility with new options in the future
	jsonSkipDefaults := false
	_ = serverOpts.ReadOpt("jsonSkipDefaults", &jsonSkipDefaults)
	jsonCamelCase := false
	_ = serverOpts.ReadOpt("jsonCamelCase", &jsonCamelCase)
	var pathPrefix string
	if ok := serverOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	return &authServiceServer{
		AuthService:      svc,
		hooks:            serverOpts.Hooks,
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		pathPrefix:       pathPrefix,
		jsonSkipDefaults: jsonSkipDefaults,
		jsonCamelCase:    jsonCamelCase,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *authServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *authServiceServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
	if context.Canceled == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.Canceled, "failed to read request: context canceled"))
		return
	}
	if context.DeadlineExceeded == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.DeadlineExceeded, "failed to read request: deadline exceeded"))
		return
	}
	s.writeError(ctx, resp, twirp.WrapError(malformedRequestError(msg), err))
}

// AuthServicePathPrefix is a convenience constant that may identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// with the default "/twirp" prefix and default CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const AuthServicePathPrefix = "/twirp/twirp.v2.AuthService/"

func (s *authServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "twirp.v2")
	ctx = ctxsetters.WithServiceName(ctx, "AuthService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	// Verify path format: [<prefix>]/<package>.<Service>/<Method>
	prefix, pkgService, method := parseTwirpPath(req.URL.Path)
	if pkgService != "twirp.v2.AuthService" {
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
	if prefix != s.pathPrefix {
		msg := fmt.Sprintf("invalid path prefix %q, expected %q, on path %q", prefix, s.pathPrefix, req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	switch method {
	case "CreateSession":
		s.serveCreateSession(ctx, resp, req)
		return
	case "ListSessions":
		s.serveListSessions(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *authServiceServer) serveCreateSession(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCreateSessionJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveCreateSessionProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *authServiceServer) serveCreateSessionJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateSession")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(CreateSessionRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.AuthService.CreateSession
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CreateSessionRequest) (*CreateSessionResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateSessionRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateSessionRequest) when calling interceptor")
					}
					return s.AuthService.CreateSession(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateSessionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateSessionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CreateSessionResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateSessionResponse and nil error while calling CreateSession. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *authServiceServer) serveCreateSessionProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateSession")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := io.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(CreateSessionRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.AuthService.CreateSession
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CreateSessionRequest) (*CreateSessionResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateSessionRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateSessionRequest) when calling interceptor")
					}
					return s.AuthService.CreateSession(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateSessionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateSessionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CreateSessionResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateSessionResponse and nil error while calling CreateSession. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *authServiceServer) serveListSessions(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveListSessionsJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveListSessionsProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *authServiceServer) serveListSessionsJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListSessions")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(ListSessionsRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.AuthService.ListSessions
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *ListSessionsRequest) (*ListSessionsResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListSessionsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListSessionsRequest) when calling interceptor")
					}
					return s.AuthService.ListSessions(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListSessionsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListSessionsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *ListSessionsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ListSessionsResponse and nil error while calling ListSessions. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *authServiceServer) serveListSessionsProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListSessions")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := io.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(ListSessionsRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.AuthService.ListSessions
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *ListSessionsRequest) (*ListSessionsResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListSessionsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListSessionsRequest) when calling interceptor")
					}
					return s.AuthService.ListSessions(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListSessionsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListSessionsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *ListSessionsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ListSessionsResponse and nil error while calling ListSessions. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *authServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor3, 0
}

func (s *authServiceServer) ProtocGenTwirpVersion() string {
	return "v8.1.3"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *authServiceServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "twirp.v2", "AuthService")
}

var twirpFileDescriptor3 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4e, 0xf3, 0x30,
	0x10, 0xc4, 0xbf, 0xe8, 0x13, 0x50, 0xb6, 0xe1, 0xe2, 0xa6, 0x52, 0x15, 0x89, 0x52, 0xf9, 0x80,
	0x7a, 0x00, 0x07, 0x85, 0x07, 0xe0, 0xdf, 0x15, 0x09, 0x29, 0xbd, 0x71, 0x73, 0xc2, 0x8a, 0x58,
	0x4a, 0x62, 0xe3, 0x75, 0xc2, 0x7b, 0xf1, 0x84, 0x48, 0x71, 0x02, 0x14, 0xda, 0xa3, 0x7f, 0xb3,
	0x33, 0x9e, 0xb5, 0x61, 0xe6, 0xde, 0x95, 0x35, 0x49, 0x97, 0x26, 0xb2, 0x75, 0xa5, 0x30, 0x56,
	0x3b, 0xcd, 0x26, 0x3d, 0x14, 0x5d, 0x1a, 0x87, 0xb5, 0x7e, 0xc1, 0x8a, 0x3c, 0xe7, 0x17, 0x10,
	0x3d, 0x58, 0x94, 0x0e, 0x37, 0x48, 0xa4, 0x74, 0x93, 0xe1, 0x5b, 0x8b, 0xe4, 0x58, 0x04, 0x07,
	0x58, 0x4b, 0x55, 0x2d, 0x82, 0x55, 0xb0, 0x3e, 0xce, 0xfc, 0x81, 0xdf, 0xc0, 0xfc, 0xd7, 0x34,
	0x19, 0xdd, 0x10, 0xb2, 0x73, 0x38, 0x22, 0x8f, 0x7a, 0xc3, 0x34, 0x0d, 0x85, 0x35, 0x85, 0x18,
	0xc7, 0x46, 0x91, 0xcf, 0x61, 0xf6, 0xa8, 0xc8, 0x0d, 0x9c, 0x86, 0xdb, 0xf8, 0x2d, 0x44, 0xdb,
	0x78, 0x88, 0x5d, 0xc3, 0x64, 0x70, 0xd2, 0x22, 0x58, 0xfd, 0xff, 0x93, 0xfb, 0xa5, 0xa6, 0x1f,
	0x01, 0x4c, 0xef, 0x5a, 0x57, 0x6e, 0xd0, 0x76, 0xaa, 0x40, 0x96, 0xc1, 0xc9, 0x56, 0x53, 0xb6,
	0x14, 0xe3, 0x0b, 0x88, 0x5d, 0x0b, 0xc7, 0x67, 0x7b, 0x75, 0xdf, 0x85, 0xff, 0x63, 0x4f, 0x10,
	0xfe, 0x6c, 0xc9, 0x4e, 0xbf, 0x2d, 0x3b, 0x96, 0x8a, 0x97, 0xfb, 0xe4, 0x31, 0xf0, 0xfe, 0xea,
	0x59, 0xbc, 0x2a, 0x57, 0xb6, 0xb9, 0x28, 0x74, 0x9d, 0xc8, 0x02, 0xad, 0x2e, 0xd1, 0x36, 0xb2,
	0x49, 0x7a, 0xe7, 0x65, 0xae, 0x55, 0x85, 0xd6, 0x54, 0xd2, 0xa1, 0x27, 0xf9, 0x61, 0xff, 0x6b,
	0xd7, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x84, 0xc0, 0x90, 0x10, 0xe4, 0x01, 0x00, 0x00,
}
