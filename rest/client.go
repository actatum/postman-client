// Package rest provides types/client for making requests to the postman REST api.
package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

const (
	baseURL = "https://api.getpostman.com"
)

// Client handles interacting with the postman api.
type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
	logOutput  io.Writer
}

// NewClient returns a new instance of the Client.
func NewClient(apiKey string, opts ...Option) *Client {
	options := options{
		httpClient: &http.Client{},
		debugLog:   nil,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &Client{
		httpClient: options.httpClient,
		apiKey:     apiKey,
		baseURL:    baseURL,
		logOutput:  options.debugLog,
	}
}

// BaseURL returns the baseURL for the rest client.
func (c *Client) BaseURL() string {
	return c.baseURL
}

// NewRequest creates a new http request.
func (c *Client) NewRequest(ctx context.Context, method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	return http.NewRequestWithContext(ctx, method, url, buf)
}

// DoRequest makes the http request and unmarshalls the response into the result interface.
func (c *Client) DoRequest(r *http.Request, result interface{}, opts ...RequestOption) error {
	options := requestOptions{
		contentType: "application/json",
	}

	for _, o := range opts {
		o.apply(&options)
	}

	query := r.URL.Query()

	if options.workspace != "" {
		query.Add("workspace", options.workspace)
	}
	r.URL.RawQuery = query.Encode()

	// Set default headers
	r.Header.Set("X-Api-Key", c.apiKey)
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Content-Type", options.contentType)

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	c.log(r, resp)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var errResp ErrorResponse
		if err = json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return err
		}

		return errResp.Error
	}

	if result == nil {
		return nil
	}

	if w, ok := result.(io.Writer); ok {
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			return err
		}
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

func (c *Client) log(r *http.Request, resp *http.Response) {
	if c.logOutput == nil {
		return
	}

	var (
		reqDump  string
		respDump []byte
	)
	if r != nil {
		reqDump = fmt.Sprintf("%s %s Data: %s", r.Method, r.URL.String(), r.Form.Encode())
	}
	if resp != nil {
		respDump, _ = httputil.DumpResponse(resp, true)
	}

	_, _ = fmt.Fprintf(c.logOutput, "Request: %s\nResponse: %s\n", reqDump, string(respDump))
}
