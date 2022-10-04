// Package postman provides a client set with handles to all the different postman endpoints.
package postman

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
