package postman

import "testing"

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
