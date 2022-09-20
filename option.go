package postman

import "net/http"

type options struct {
	httpClient *http.Client
}

// Option are functional options for configuring the client.
type Option interface {
	apply(*options)
}

type httpClientOption struct {
	c *http.Client
}

func (h httpClientOption) apply(opts *options) {
	opts.httpClient = h.c
}

// WithHTTPClient configures the client to use the given http client.
func WithHTTPClient(client *http.Client) Option {
	return httpClientOption{c: client}
}
