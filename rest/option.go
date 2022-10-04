// Package rest provides types/client for making requests to the postman REST api.
package rest

import (
	"io"
	"net/http"
	"os"
)

type options struct {
	httpClient *http.Client
	debugLog   io.Writer
}

// Option represents functional options for configuring the client.
type Option interface {
	apply(*options)
}

type httpClientOption struct {
	c *http.Client
}

func (h httpClientOption) apply(opts *options) {
	opts.httpClient = h.c
}

// WithHTTPClient configures the client to use the given http.Client.
func WithHTTPClient(client *http.Client) Option {
	return httpClientOption{c: client}
}

type debugLogOption struct {
	w io.Writer
}

func (d debugLogOption) apply(opts *options) {
	if d.w == nil {
		d.w = os.Stdout
	}
	opts.debugLog = d.w
}

// WithDebugLog configures the io.Writer to send debug logging output to.
func WithDebugLog(w io.Writer) Option {
	return debugLogOption{w: w}
}

type requestOptions struct {
	workspace   string
	contentType string
}

// RequestOption represents functional options for configuring client requests.
type RequestOption interface {
	apply(*requestOptions)
}

type workspaceOption string

func (w workspaceOption) apply(opts *requestOptions) {
	opts.workspace = string(w)
}

// WithWorkspace specifies the target workspace for the api operation.
func WithWorkspace(w string) RequestOption {
	return workspaceOption(w)
}

type contentTypeOption string

func (c contentTypeOption) apply(opts *requestOptions) {
	opts.contentType = string(c)
}

// WithContentType specifies the content type for the request
func WithContentType(c string) RequestOption {
	return contentTypeOption(c)
}
