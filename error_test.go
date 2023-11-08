package client

import "testing"

func TestFailRequestError_Error(t *testing.T) {
	type fields struct {
		Code         int
		ResponseBody []byte
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Success case without response body",
			fields: fields{
				Code:         200,
				ResponseBody: nil,
			},
			want: "200 fail request",
		},
		{
			name: "Success case with response body",
			fields: fields{
				Code:         400,
				ResponseBody: []byte("error message"),
			},
			want: "400 fail request, error message: error message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &FailRequestError{
				Code:         tt.fields.Code,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("FailRequestError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
