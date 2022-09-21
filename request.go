package postman

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type requestOptions struct {
	workspace string
}

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

func doRequest(r *http.Request, apiKey string, httpClient *http.Client, opts ...RequestOption) (data []byte, err error) {
	options := requestOptions{}

	for _, o := range opts {
		o.apply(&options)
	}

	query := r.URL.Query()

	if options.workspace != "" {
		query.Add("workspace", options.workspace)
	}

	r.URL.RawQuery = query.Encode()
	r.Header.Set("x-api-key", apiKey)
	r.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %d\n", resp.StatusCode)
	if resp.StatusCode >= http.StatusBadRequest {
		var postmanError Error
		if err = json.NewDecoder(resp.Body).Decode(&postmanError); err != nil {
			return nil, err
		}

		return nil, postmanError
	}

	return io.ReadAll(resp.Body)
}
