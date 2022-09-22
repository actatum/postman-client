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

func doRequest(
	r *http.Request,
	apiKey string,
	httpClient *http.Client,
	opts ...RequestOption,
) (data []byte, err error) {
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

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Result: \n%s\n", string(data))
	if resp.StatusCode >= http.StatusBadRequest {
		var wrapper errorWrapper
		if err = json.Unmarshal(data, &wrapper); err != nil {
			return nil, err
		}

		return nil, wrapper.Error
	}

	return data, nil
}
