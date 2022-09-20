package postman

import (
	"net/http"
)

const postmanBaseURL = "https://api.getpostman.com"

// ClientSet contains the clients for groups. Each group has a single client.
type ClientSet struct {
	collections *CollectionsClient
}

// New returns a new instance of the ClientSet.
func New(apiKey string, opts ...Option) *ClientSet {
	options := options{
		httpClient: &http.Client{},
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &ClientSet{
		collections: NewCollectionsClient(apiKey, options.httpClient),
	}
}

// Collections returns a handle to a collections.CollectionsClient.
func (cs *ClientSet) Collections() *CollectionsClient {
	return cs.collections
}
