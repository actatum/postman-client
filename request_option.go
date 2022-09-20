package postman

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

func WithWorkspace(w string) RequestOption {
	return workspaceOption(w)
}
