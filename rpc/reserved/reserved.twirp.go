// Code generated by protoc-gen-twirp v8.0.0, DO NOT EDIT.
// source: reserved.proto

/*
Package reserved is a generated twirp stub package.
This code was generated with github.com/twitchtv/twirp/protoc-gen-twirp v8.0.0.

It is generated from these files:
	reserved.proto
*/
package reserved

import context "context"
import fmt "fmt"
import http "net/http"
import ioutil "io/ioutil"
import json "encoding/json"
import strconv "strconv"
import strings "strings"

import protojson "google.golang.org/protobuf/encoding/protojson"
import proto "google.golang.org/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

import bytes "bytes"
import io "io"
import path "path"
import url "net/url"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the twirp package used in your project.
// A compilation error at this line likely means your copy of the
// twirp package needs to be updated.
const _ = twirp.TwirpPackageIsVersion7

// ==================
// Reserved Interface
// ==================

// Master service.
type Reserved interface {
	// GetAircraft returns a Aircraft from a Query.
	GetAircraft(context.Context, *Query) (*Aircraft, error)
}

// ========================
// Reserved Protobuf Client
// ========================

type reservedProtobufClient struct {
	client      HTTPClient
	urls        [1]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewReservedProtobufClient creates a Protobuf client that implements the Reserved interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewReservedProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) Reserved {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(clientOpts.PathPrefix(), "engelsjk.faadb.reserved", "Reserved")
	urls := [1]string{
		serviceURL + "GetAircraft",
	}

	return &reservedProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *reservedProtobufClient) GetAircraft(ctx context.Context, in *Query) (*Aircraft, error) {
	ctx = ctxsetters.WithPackageName(ctx, "engelsjk.faadb.reserved")
	ctx = ctxsetters.WithServiceName(ctx, "Reserved")
	ctx = ctxsetters.WithMethodName(ctx, "GetAircraft")
	caller := c.callGetAircraft
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *Query) (*Aircraft, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Query)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Query) when calling interceptor")
					}
					return c.callGetAircraft(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Aircraft)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Aircraft) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *reservedProtobufClient) callGetAircraft(ctx context.Context, in *Query) (*Aircraft, error) {
	out := new(Aircraft)
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

// ====================
// Reserved JSON Client
// ====================

type reservedJSONClient struct {
	client      HTTPClient
	urls        [1]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewReservedJSONClient creates a JSON client that implements the Reserved interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewReservedJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) Reserved {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(clientOpts.PathPrefix(), "engelsjk.faadb.reserved", "Reserved")
	urls := [1]string{
		serviceURL + "GetAircraft",
	}

	return &reservedJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *reservedJSONClient) GetAircraft(ctx context.Context, in *Query) (*Aircraft, error) {
	ctx = ctxsetters.WithPackageName(ctx, "engelsjk.faadb.reserved")
	ctx = ctxsetters.WithServiceName(ctx, "Reserved")
	ctx = ctxsetters.WithMethodName(ctx, "GetAircraft")
	caller := c.callGetAircraft
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *Query) (*Aircraft, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Query)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Query) when calling interceptor")
					}
					return c.callGetAircraft(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Aircraft)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Aircraft) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *reservedJSONClient) callGetAircraft(ctx context.Context, in *Query) (*Aircraft, error) {
	out := new(Aircraft)
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

// =======================
// Reserved Server Handler
// =======================

type reservedServer struct {
	Reserved
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
}

// NewReservedServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewReservedServer(svc Reserved, opts ...interface{}) TwirpServer {
	serverOpts := twirp.ServerOptions{}
	for _, opt := range opts {
		switch o := opt.(type) {
		case twirp.ServerOption:
			o(&serverOpts)
		case *twirp.ServerHooks: // backwards compatibility, allow to specify hooks as an argument
			twirp.WithServerHooks(o)(&serverOpts)
		case nil: // backwards compatibility, allow nil value for the argument
			continue
		default:
			panic(fmt.Sprintf("Invalid option type %T on NewReservedServer", o))
		}
	}

	return &reservedServer{
		Reserved:         svc,
		pathPrefix:       serverOpts.PathPrefix(),
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		hooks:            serverOpts.Hooks,
		jsonSkipDefaults: serverOpts.JSONSkipDefaults,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *reservedServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *reservedServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
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

// ReservedPathPrefix is a convenience constant that could used to identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// that add a "/twirp" prefix by default, and use CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const ReservedPathPrefix = "/twirp/engelsjk.faadb.reserved.Reserved/"

func (s *reservedServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "engelsjk.faadb.reserved")
	ctx = ctxsetters.WithServiceName(ctx, "Reserved")
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
	if pkgService != "engelsjk.faadb.reserved.Reserved" {
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
	case "GetAircraft":
		s.serveGetAircraft(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *reservedServer) serveGetAircraft(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetAircraftJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveGetAircraftProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *reservedServer) serveGetAircraftJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetAircraft")
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
	reqContent := new(Query)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.Reserved.GetAircraft
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *Query) (*Aircraft, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Query)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Query) when calling interceptor")
					}
					return s.Reserved.GetAircraft(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Aircraft)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Aircraft) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *Aircraft
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Aircraft and nil error while calling GetAircraft. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: true, EmitUnpopulated: !s.jsonSkipDefaults}
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

func (s *reservedServer) serveGetAircraftProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetAircraft")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(Query)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.Reserved.GetAircraft
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *Query) (*Aircraft, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Query)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Query) when calling interceptor")
					}
					return s.Reserved.GetAircraft(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Aircraft)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Aircraft) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *Aircraft
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Aircraft and nil error while calling GetAircraft. nil responses are not supported"))
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

func (s *reservedServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor0, 0
}

func (s *reservedServer) ProtocGenTwirpVersion() string {
	return "v8.0.0"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *reservedServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "engelsjk.faadb.reserved", "Reserved")
}

// =====
// Utils
// =====

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
//
// HTTPClient implementations should not follow redirects. Redirects are
// automatically disabled if *(net/http).Client is passed to client
// constructors. See the withoutRedirects function in this file for more
// details.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// TwirpServer is the interface generated server structs will support: they're
// HTTP handlers with additional methods for accessing metadata about the
// service. Those accessors are a low-level API for building reflection tools.
// Most people can think of TwirpServers as just http.Handlers.
type TwirpServer interface {
	http.Handler

	// ServiceDescriptor returns gzipped bytes describing the .proto file that
	// this service was generated from. Once unzipped, the bytes can be
	// unmarshalled as a
	// google.golang.org/protobuf/types/descriptorpb.FileDescriptorProto.
	//
	// The returned integer is the index of this particular service within that
	// FileDescriptorProto's 'Service' slice of ServiceDescriptorProtos. This is a
	// low-level field, expected to be used for reflection.
	ServiceDescriptor() ([]byte, int)

	// ProtocGenTwirpVersion is the semantic version string of the version of
	// twirp used to generate this file.
	ProtocGenTwirpVersion() string

	// PathPrefix returns the HTTP URL path prefix for all methods handled by this
	// service. This can be used with an HTTP mux to route Twirp requests.
	// The path prefix is in the form: "/<prefix>/<package>.<Service>/"
	// that is, everything in a Twirp route except for the <Method> at the end.
	PathPrefix() string
}

// WriteError writes an HTTP response with a valid Twirp error format (code, msg, meta).
// Useful outside of the Twirp server (e.g. http middleware), but does not trigger hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func WriteError(resp http.ResponseWriter, err error) {
	writeError(context.Background(), resp, err, nil)
}

// writeError writes Twirp errors in the response and triggers hooks.
func writeError(ctx context.Context, resp http.ResponseWriter, err error, hooks *twirp.ServerHooks) {
	// Non-twirp errors are wrapped as Internal (default)
	twerr, ok := err.(twirp.Error)
	if !ok {
		twerr = twirp.InternalErrorWith(err)
	}

	statusCode := twirp.ServerHTTPStatusFromErrorCode(twerr.Code())
	ctx = ctxsetters.WithStatusCode(ctx, statusCode)
	ctx = callError(ctx, hooks, twerr)

	respBody := marshalErrorToJSON(twerr)

	resp.Header().Set("Content-Type", "application/json") // Error responses are always JSON
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBody)))
	resp.WriteHeader(statusCode) // set HTTP status code and send response

	_, writeErr := resp.Write(respBody)
	if writeErr != nil {
		// We have three options here. We could log the error, call the Error
		// hook, or just silently ignore the error.
		//
		// Logging is unacceptable because we don't have a user-controlled
		// logger; writing out to stderr without permission is too rude.
		//
		// Calling the Error hook would confuse users: it would mean the Error
		// hook got called twice for one request, which is likely to lead to
		// duplicated log messages and metrics, no matter how well we document
		// the behavior.
		//
		// Silently ignoring the error is our least-bad option. It's highly
		// likely that the connection is broken and the original 'err' says
		// so anyway.
		_ = writeErr
	}

	callResponseSent(ctx, hooks)
}

// sanitizeBaseURL parses the the baseURL, and adds the "http" scheme if needed.
// If the URL is unparsable, the baseURL is returned unchaged.
func sanitizeBaseURL(baseURL string) string {
	u, err := url.Parse(baseURL)
	if err != nil {
		return baseURL // invalid URL will fail later when making requests
	}
	if u.Scheme == "" {
		u.Scheme = "http"
	}
	return u.String()
}

// baseServicePath composes the path prefix for the service (without <Method>).
// e.g.: baseServicePath("/twirp", "my.pkg", "MyService")
//       returns => "/twirp/my.pkg.MyService/"
// e.g.: baseServicePath("", "", "MyService")
//       returns => "/MyService/"
func baseServicePath(prefix, pkg, service string) string {
	fullServiceName := service
	if pkg != "" {
		fullServiceName = pkg + "." + service
	}
	return path.Join("/", prefix, fullServiceName) + "/"
}

// parseTwirpPath extracts path components form a valid Twirp route.
// Expected format: "[<prefix>]/<package>.<Service>/<Method>"
// e.g.: prefix, pkgService, method := parseTwirpPath("/twirp/pkg.Svc/MakeHat")
func parseTwirpPath(path string) (string, string, string) {
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		return "", "", ""
	}
	method := parts[len(parts)-1]
	pkgService := parts[len(parts)-2]
	prefix := strings.Join(parts[0:len(parts)-2], "/")
	return prefix, pkgService, method
}

// getCustomHTTPReqHeaders retrieves a copy of any headers that are set in
// a context through the twirp.WithHTTPRequestHeaders function.
// If there are no headers set, or if they have the wrong type, nil is returned.
func getCustomHTTPReqHeaders(ctx context.Context) http.Header {
	header, ok := twirp.HTTPRequestHeaders(ctx)
	if !ok || header == nil {
		return nil
	}
	copied := make(http.Header)
	for k, vv := range header {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}
	return copied
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if customHeader := getCustomHTTPReqHeaders(ctx); customHeader != nil {
		req.Header = customHeader
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Twirp-Version", "v8.0.0")
	return req, nil
}

// JSON serialization for errors
type twerrJSON struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Meta map[string]string `json:"meta,omitempty"`
}

// marshalErrorToJSON returns JSON from a twirp.Error, that can be used as HTTP error response body.
// If serialization fails, it will use a descriptive Internal error instead.
func marshalErrorToJSON(twerr twirp.Error) []byte {
	// make sure that msg is not too large
	msg := twerr.Msg()
	if len(msg) > 1e6 {
		msg = msg[:1e6]
	}

	tj := twerrJSON{
		Code: string(twerr.Code()),
		Msg:  msg,
		Meta: twerr.MetaMap(),
	}

	buf, err := json.Marshal(&tj)
	if err != nil {
		buf = []byte("{\"type\": \"" + twirp.Internal + "\", \"msg\": \"There was an error but it could not be serialized into JSON\"}") // fallback
	}

	return buf
}

// errorFromResponse builds a twirp.Error from a non-200 HTTP response.
// If the response has a valid serialized Twirp error, then it's returned.
// If not, the response status code is used to generate a similar twirp
// error. See twirpErrorFromIntermediary for more info on intermediary errors.
func errorFromResponse(resp *http.Response) twirp.Error {
	statusCode := resp.StatusCode
	statusText := http.StatusText(statusCode)

	if isHTTPRedirect(statusCode) {
		// Unexpected redirect: it must be an error from an intermediary.
		// Twirp clients don't follow redirects automatically, Twirp only handles
		// POST requests, redirects should only happen on GET and HEAD requests.
		location := resp.Header.Get("Location")
		msg := fmt.Sprintf("unexpected HTTP status code %d %q received, Location=%q", statusCode, statusText, location)
		return twirpErrorFromIntermediary(statusCode, msg, location)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return wrapInternal(err, "failed to read server error response body")
	}

	var tj twerrJSON
	dec := json.NewDecoder(bytes.NewReader(respBodyBytes))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&tj); err != nil || tj.Code == "" {
		// Invalid JSON response; it must be an error from an intermediary.
		msg := fmt.Sprintf("Error from intermediary with HTTP status code %d %q", statusCode, statusText)
		return twirpErrorFromIntermediary(statusCode, msg, string(respBodyBytes))
	}

	errorCode := twirp.ErrorCode(tj.Code)
	if !twirp.IsValidErrorCode(errorCode) {
		msg := "invalid type returned from server error response: " + tj.Code
		return twirp.InternalError(msg).WithMeta("body", string(respBodyBytes))
	}

	twerr := twirp.NewError(errorCode, tj.Msg)
	for k, v := range tj.Meta {
		twerr = twerr.WithMeta(k, v)
	}
	return twerr
}

// twirpErrorFromIntermediary maps HTTP errors from non-twirp sources to twirp errors.
// The mapping is similar to gRPC: https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md.
// Returned twirp Errors have some additional metadata for inspection.
func twirpErrorFromIntermediary(status int, msg string, bodyOrLocation string) twirp.Error {
	var code twirp.ErrorCode
	if isHTTPRedirect(status) { // 3xx
		code = twirp.Internal
	} else {
		switch status {
		case 400: // Bad Request
			code = twirp.Internal
		case 401: // Unauthorized
			code = twirp.Unauthenticated
		case 403: // Forbidden
			code = twirp.PermissionDenied
		case 404: // Not Found
			code = twirp.BadRoute
		case 429: // Too Many Requests
			code = twirp.ResourceExhausted
		case 502, 503, 504: // Bad Gateway, Service Unavailable, Gateway Timeout
			code = twirp.Unavailable
		default: // All other codes
			code = twirp.Unknown
		}
	}

	twerr := twirp.NewError(code, msg)
	twerr = twerr.WithMeta("http_error_from_intermediary", "true") // to easily know if this error was from intermediary
	twerr = twerr.WithMeta("status_code", strconv.Itoa(status))
	if isHTTPRedirect(status) {
		twerr = twerr.WithMeta("location", bodyOrLocation)
	} else {
		twerr = twerr.WithMeta("body", bodyOrLocation)
	}
	return twerr
}

func isHTTPRedirect(status int) bool {
	return status >= 300 && status <= 399
}

// wrapInternal wraps an error with a prefix as an Internal error.
// The original error cause is accessible by github.com/pkg/errors.Cause.
func wrapInternal(err error, prefix string) twirp.Error {
	return twirp.InternalErrorWith(&wrappedError{prefix: prefix, cause: err})
}

type wrappedError struct {
	prefix string
	cause  error
}

func (e *wrappedError) Error() string { return e.prefix + ": " + e.cause.Error() }
func (e *wrappedError) Unwrap() error { return e.cause } // for go1.13 + errors.Is/As
func (e *wrappedError) Cause() error  { return e.cause } // for github.com/pkg/errors

// ensurePanicResponses makes sure that rpc methods causing a panic still result in a Twirp Internal
// error response (status 500), and error hooks are properly called with the panic wrapped as an error.
// The panic is re-raised so it can be handled normally with middleware.
func ensurePanicResponses(ctx context.Context, resp http.ResponseWriter, hooks *twirp.ServerHooks) {
	if r := recover(); r != nil {
		// Wrap the panic as an error so it can be passed to error hooks.
		// The original error is accessible from error hooks, but not visible in the response.
		err := errFromPanic(r)
		twerr := &internalWithCause{msg: "Internal service panic", cause: err}
		// Actually write the error
		writeError(ctx, resp, twerr, hooks)
		// If possible, flush the error to the wire.
		f, ok := resp.(http.Flusher)
		if ok {
			f.Flush()
		}

		panic(r)
	}
}

// errFromPanic returns the typed error if the recovered panic is an error, otherwise formats as error.
func errFromPanic(p interface{}) error {
	if err, ok := p.(error); ok {
		return err
	}
	return fmt.Errorf("panic: %v", p)
}

// internalWithCause is a Twirp Internal error wrapping an original error cause,
// but the original error message is not exposed on Msg(). The original error
// can be checked with go1.13+ errors.Is/As, and also by (github.com/pkg/errors).Unwrap
type internalWithCause struct {
	msg   string
	cause error
}

func (e *internalWithCause) Unwrap() error                               { return e.cause } // for go1.13 + errors.Is/As
func (e *internalWithCause) Cause() error                                { return e.cause } // for github.com/pkg/errors
func (e *internalWithCause) Error() string                               { return e.msg + ": " + e.cause.Error() }
func (e *internalWithCause) Code() twirp.ErrorCode                       { return twirp.Internal }
func (e *internalWithCause) Msg() string                                 { return e.msg }
func (e *internalWithCause) Meta(key string) string                      { return "" }
func (e *internalWithCause) MetaMap() map[string]string                  { return nil }
func (e *internalWithCause) WithMeta(key string, val string) twirp.Error { return e }

// malformedRequestError is used when the twirp server cannot unmarshal a request
func malformedRequestError(msg string) twirp.Error {
	return twirp.NewError(twirp.Malformed, msg)
}

// badRouteError is used when the twirp server cannot route a request
func badRouteError(msg string, method, url string) twirp.Error {
	err := twirp.NewError(twirp.BadRoute, msg)
	err = err.WithMeta("twirp_invalid_route", method+" "+url)
	return err
}

// withoutRedirects makes sure that the POST request can not be redirected.
// The standard library will, by default, redirect requests (including POSTs) if it gets a 302 or
// 303 response, and also 301s in go1.8. It redirects by making a second request, changing the
// method to GET and removing the body. This produces very confusing error messages, so instead we
// set a redirect policy that always errors. This stops Go from executing the redirect.
//
// We have to be a little careful in case the user-provided http.Client has its own CheckRedirect
// policy - if so, we'll run through that policy first.
//
// Because this requires modifying the http.Client, we make a new copy of the client and return it.
func withoutRedirects(in *http.Client) *http.Client {
	copy := *in
	copy.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if in.CheckRedirect != nil {
			// Run the input's redirect if it exists, in case it has side effects, but ignore any error it
			// returns, since we want to use ErrUseLastResponse.
			err := in.CheckRedirect(req, via)
			_ = err // Silly, but this makes sure generated code passes errcheck -blank, which some people use.
		}
		return http.ErrUseLastResponse
	}
	return &copy
}

// doProtobufRequest makes a Protobuf request to the remote Twirp service.
func doProtobufRequest(ctx context.Context, client HTTPClient, hooks *twirp.ClientHooks, url string, in, out proto.Message) (_ context.Context, err error) {
	reqBodyBytes, err := proto.Marshal(in)
	if err != nil {
		return ctx, wrapInternal(err, "failed to marshal proto request")
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)
	if err = ctx.Err(); err != nil {
		return ctx, wrapInternal(err, "aborted because context was done")
	}

	req, err := newRequest(ctx, url, reqBody, "application/protobuf")
	if err != nil {
		return ctx, wrapInternal(err, "could not build request")
	}
	ctx, err = callClientRequestPrepared(ctx, hooks, req)
	if err != nil {
		return ctx, err
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return ctx, wrapInternal(err, "failed to do request")
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = wrapInternal(cerr, "failed to close response body")
		}
	}()

	if err = ctx.Err(); err != nil {
		return ctx, wrapInternal(err, "aborted because context was done")
	}

	if resp.StatusCode != 200 {
		return ctx, errorFromResponse(resp)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ctx, wrapInternal(err, "failed to read response body")
	}
	if err = ctx.Err(); err != nil {
		return ctx, wrapInternal(err, "aborted because context was done")
	}

	if err = proto.Unmarshal(respBodyBytes, out); err != nil {
		return ctx, wrapInternal(err, "failed to unmarshal proto response")
	}
	return ctx, nil
}

// doJSONRequest makes a JSON request to the remote Twirp service.
func doJSONRequest(ctx context.Context, client HTTPClient, hooks *twirp.ClientHooks, url string, in, out proto.Message) (_ context.Context, err error) {
	marshaler := &protojson.MarshalOptions{UseProtoNames: true}
	reqBytes, err := marshaler.Marshal(in)
	if err != nil {
		return ctx, wrapInternal(err, "failed to marshal json request")
	}
	if err = ctx.Err(); err != nil {
		return ctx, wrapInternal(err, "aborted because context was done")
	}

	req, err := newRequest(ctx, url, bytes.NewReader(reqBytes), "application/json")
	if err != nil {
		return ctx, wrapInternal(err, "could not build request")
	}
	ctx, err = callClientRequestPrepared(ctx, hooks, req)
	if err != nil {
		return ctx, err
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return ctx, wrapInternal(err, "failed to do request")
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = wrapInternal(cerr, "failed to close response body")
		}
	}()

	if err = ctx.Err(); err != nil {
		return ctx, wrapInternal(err, "aborted because context was done")
	}

	if resp.StatusCode != 200 {
		return ctx, errorFromResponse(resp)
	}

	d := json.NewDecoder(resp.Body)
	rawRespBody := json.RawMessage{}
	if err := d.Decode(&rawRespBody); err != nil {
		return ctx, wrapInternal(err, "failed to unmarshal json response")
	}
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawRespBody, out); err != nil {
		return ctx, wrapInternal(err, "failed to unmarshal json response")
	}
	if err = ctx.Err(); err != nil {
		return ctx, wrapInternal(err, "aborted because context was done")
	}
	return ctx, nil
}

// Call twirp.ServerHooks.RequestReceived if the hook is available
func callRequestReceived(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestReceived == nil {
		return ctx, nil
	}
	return h.RequestReceived(ctx)
}

// Call twirp.ServerHooks.RequestRouted if the hook is available
func callRequestRouted(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestRouted == nil {
		return ctx, nil
	}
	return h.RequestRouted(ctx)
}

// Call twirp.ServerHooks.ResponsePrepared if the hook is available
func callResponsePrepared(ctx context.Context, h *twirp.ServerHooks) context.Context {
	if h == nil || h.ResponsePrepared == nil {
		return ctx
	}
	return h.ResponsePrepared(ctx)
}

// Call twirp.ServerHooks.ResponseSent if the hook is available
func callResponseSent(ctx context.Context, h *twirp.ServerHooks) {
	if h == nil || h.ResponseSent == nil {
		return
	}
	h.ResponseSent(ctx)
}

// Call twirp.ServerHooks.Error if the hook is available
func callError(ctx context.Context, h *twirp.ServerHooks, err twirp.Error) context.Context {
	if h == nil || h.Error == nil {
		return ctx
	}
	return h.Error(ctx, err)
}

func callClientResponseReceived(ctx context.Context, h *twirp.ClientHooks) {
	if h == nil || h.ResponseReceived == nil {
		return
	}
	h.ResponseReceived(ctx)
}

func callClientRequestPrepared(ctx context.Context, h *twirp.ClientHooks, req *http.Request) (context.Context, error) {
	if h == nil || h.RequestPrepared == nil {
		return ctx, nil
	}
	return h.RequestPrepared(ctx, req)
}

func callClientError(ctx context.Context, h *twirp.ClientHooks, err twirp.Error) {
	if h == nil || h.Error == nil {
		return
	}
	h.Error(ctx, err)
}

var twirpFileDescriptor0 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0x57, 0x16, 0xb6, 0xb5, 0xaf, 0xd3, 0x00, 0x0b, 0x84, 0x35, 0x21, 0x28, 0x3d, 0xa0, 0x0a,
	0xa1, 0x44, 0x2b, 0xbb, 0x71, 0x0a, 0xe5, 0xcf, 0x2e, 0x94, 0xd1, 0x72, 0x9a, 0xb8, 0xb8, 0xc9,
	0x6b, 0x6a, 0x68, 0x63, 0xeb, 0xc5, 0x1d, 0xeb, 0x27, 0xe1, 0x43, 0xf2, 0x25, 0x50, 0xec, 0x84,
	0x76, 0x4d, 0x0b, 0xec, 0x16, 0xff, 0xfe, 0xc5, 0x79, 0x3f, 0x3b, 0x70, 0x4c, 0x98, 0x23, 0x5d,
	0x61, 0x12, 0x68, 0x52, 0x46, 0xb1, 0x47, 0x98, 0xa5, 0x38, 0xcb, 0xbf, 0x7d, 0x0f, 0x26, 0x42,
	0x24, 0xe3, 0xa0, 0xa2, 0x3b, 0x3f, 0x7d, 0xd8, 0xff, 0xbc, 0x40, 0x5a, 0x32, 0x0e, 0x87, 0x83,
	0xc1, 0x62, 0x3e, 0x46, 0xe2, 0x5e, 0xdb, 0xeb, 0x36, 0x87, 0xd5, 0x92, 0x75, 0xe0, 0x68, 0x84,
	0x24, 0xc5, 0xac, 0xa4, 0xf7, 0x2c, 0x7d, 0x03, 0x2b, 0x34, 0x1f, 0x55, 0x82, 0xa3, 0xbe, 0x4a,
	0xf0, 0x1c, 0xaf, 0xb9, 0xef, 0x34, 0xeb, 0x18, 0x7b, 0x0e, 0xc7, 0x43, 0x4c, 0x65, 0x6e, 0x48,
	0x64, 0x66, 0x20, 0xe6, 0xc8, 0xef, 0x58, 0xd5, 0x06, 0xca, 0x5e, 0xc2, 0xfd, 0x15, 0x32, 0x32,
	0x84, 0x68, 0x4e, 0xf9, 0xbe, 0x95, 0xd6, 0x09, 0xd6, 0x85, 0xbb, 0xeb, 0xa0, 0x30, 0xc8, 0x0f,
	0xac, 0x76, 0x13, 0x2e, 0x72, 0x23, 0x49, 0x31, 0x89, 0x89, 0x29, 0xf6, 0x35, 0x2b, 0xf6, 0xc5,
	0x0f, 0x5d, 0x6e, 0x8d, 0x60, 0x67, 0xf0, 0x30, 0xd2, 0x9a, 0xd4, 0x15, 0x26, 0x9f, 0x34, 0x92,
	0x30, 0x52, 0x65, 0xd6, 0xd1, 0xb0, 0x8e, 0xed, 0x24, 0x3b, 0x87, 0xa7, 0x91, 0xa4, 0x1f, 0x8a,
	0xcc, 0x54, 0x66, 0x98, 0xe7, 0xfd, 0x99, 0xc8, 0x73, 0x39, 0x91, 0xf1, 0xca, 0xdf, 0xb4, 0xfe,
	0x7f, 0xc9, 0x3a, 0xbf, 0x7c, 0xf0, 0xa2, 0xbf, 0xb4, 0x52, 0x9f, 0xe6, 0xde, 0xff, 0x4f, 0xd3,
	0xdf, 0x35, 0xcd, 0x2d, 0xea, 0x5e, 0x59, 0x53, 0x9d, 0xb8, 0xb9, 0x87, 0xbe, 0x34, 0xcb, 0xb2,
	0xa6, 0x0d, 0xf4, 0x76, 0x1d, 0xad, 0xa0, 0x4b, 0xa9, 0xe3, 0xb5, 0x8e, 0x6a, 0x04, 0x6b, 0x43,
	0x6b, 0xe8, 0x4e, 0xf2, 0xdb, 0x22, 0xd3, 0x35, 0xb3, 0x0e, 0xb9, 0x37, 0x17, 0x4b, 0x3b, 0xd8,
	0x2f, 0x4b, 0x5d, 0xcd, 0x7f, 0x13, 0x66, 0x3d, 0x78, 0xf0, 0xee, 0x5a, 0x4b, 0xd7, 0xe5, 0x40,
	0x19, 0x19, 0xbb, 0x50, 0xb0, 0xf2, 0xad, 0x1c, 0x7b, 0x01, 0xf7, 0xca, 0x3a, 0xde, 0x2b, 0xea,
	0x4f, 0x45, 0x96, 0x22, 0x6f, 0x59, 0x7d, 0x0d, 0x67, 0x8f, 0xa1, 0x79, 0xb1, 0xa0, 0xd4, 0x85,
	0x1e, 0x59, 0xd1, 0x0a, 0xe8, 0x9c, 0x41, 0xa3, 0x3a, 0x82, 0xac, 0x0b, 0x5e, 0xc4, 0xbd, 0xb6,
	0xdf, 0x6d, 0xf5, 0x4e, 0x82, 0x1d, 0x17, 0x37, 0x88, 0x86, 0x5e, 0xd4, 0xfb, 0x0a, 0x8d, 0xf2,
	0x63, 0x13, 0x76, 0x01, 0xad, 0x0f, 0x68, 0xfe, 0x84, 0x3c, 0xd9, 0xe9, 0xb4, 0xd7, 0xfd, 0xe4,
	0xd9, 0xee, 0xe4, 0x32, 0xe2, 0xcd, 0xe9, 0x65, 0x98, 0x4a, 0x33, 0x5d, 0x8c, 0x83, 0x58, 0xcd,
	0xc3, 0x4a, 0x1e, 0x5a, 0x79, 0x48, 0x3a, 0x0e, 0x2b, 0xcb, 0xeb, 0xea, 0x61, 0x7c, 0x60, 0x7f,
	0x37, 0xaf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x69, 0xb4, 0xdc, 0x74, 0x80, 0x04, 0x00, 0x00,
}
