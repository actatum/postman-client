// Package rest provides types/client for making requests to the postman REST api.
package rest

import (
	"reflect"
	"testing"
)

func TestError_Error(t *testing.T) {
	type fields struct {
		Name    string
		Message string
		Details map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "error string",
			fields: fields{
				Name:    "instanceNotFoundError",
				Message: "We could not find the collection you are looking for",
			},
			want: "instanceNotFoundError: We could not find the collection you are looking for",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{
				Name:    tt.fields.Name,
				Message: tt.fields.Message,
				Details: tt.fields.Details,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    Error
	}{
		{
			name: "regular api error",
			args: args{
				[]byte(`{"name":"AuthenticationError",
				"message":"Invalid API Key. Every request requires a valid API Key to be sent."}`),
			},
			wantErr: false,
			want: Error{
				Name:    "AuthenticationError",
				Message: "Invalid API Key. Every request requires a valid API Key to be sent.",
			},
		},
		{
			name: "api security schema validation error",
			args: args{
				[]byte(`{"name":{"name":"Invalid schema","reason":"Provided schema type is not supported."}}`),
			},
			wantErr: false,
			want: Error{
				Name:    "Invalid schema",
				Message: "Provided schema type is not supported.",
			},
		},
		{
			name: "invalid postman api error response (missing name)",
			args: args{
				[]byte(`{"cheese":"gouda"}`),
			},
			wantErr: true,
			want:    Error{},
		},
		{
			name: "invalid postman api error response (missing message)",
			args: args{
				[]byte(`{"name":"doodlebob"}`),
			},
			wantErr: true,
			want:    Error{},
		},
		{
			name: "invalid postman api validation error response",
			args: args{
				[]byte(`{"name":{"cheese":"gouda","reason":"Provided schema type is not supported."}}`),
			},
			wantErr: true,
			want:    Error{},
		},
		{
			name: "invalid postman api validation error response",
			args: args{
				[]byte(`{"name":{"name":"Invalid schema","cheese":"gouda"}}`),
			},
			wantErr: true,
			want:    Error{},
		},
		{
			name: "invalid json",
			args: args{
				data: []byte(`"im not j{s{o}n}"`),
			},
			wantErr: true,
			want:    Error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{}
			if err := e.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}

			if !reflect.DeepEqual(e, tt.want) {
				t.Errorf("UnmarshalJSON() got = %v, want %v", e, tt.want)
			}
		})
	}
}
