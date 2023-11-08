package client

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/Aloe-Corporation/client/test"
)

func TestConnector_SimpleGet(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Success case",
			fields: fields{
				Client: FactoryHttpClient(),
			},
			args: args{
				path: "/get",
			},
			want:    []byte("This is data"),
			wantErr: false,
		},
		{
			name: "Fail case: wrong path",
			fields: fields{
				Client: FactoryHttpClient(),
			},
			args: args{
				path: "/wrong",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := test.GetEndpoint()
			defer server.Close()

			c := &Connector{
				Client: tt.fields.Client,
				URL:    server.URL,
			}
			got, err := c.SimpleGet(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connector.SimpleGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connector.SimpleGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_SimplePost(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	type args struct {
		path string
		body io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Success case",
			fields: fields{
				Client: FactoryHttpClient(),
			},
			args: args{
				path: "/post",
				body: bytes.NewReader([]byte("data")),
			},
			want:    []byte("This is data"),
			wantErr: false,
		},
		{
			name: "Fail case: wrong path",
			fields: fields{
				Client: FactoryHttpClient(),
			},
			args: args{
				path: "/wrong",
				body: bytes.NewReader([]byte("data")),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := test.PostEndpoint()
			defer server.Close()
			c := &Connector{
				Client: tt.fields.Client,
				URL:    server.URL,
			}
			got, err := c.SimplePost(tt.args.path, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connector.SimplePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connector.SimplePost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_SimplePut(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	type args struct {
		path string
		body io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Success case",
			fields: fields{
				Client: FactoryHttpClient(),
			},
			args: args{
				path: "/put",
				body: bytes.NewReader([]byte("data")),
			},
			want:    []byte("This is data"),
			wantErr: false,
		},
		{
			name: "Fail case: wrong path",
			fields: fields{
				Client: FactoryHttpClient(),
			},
			args: args{
				path: "/wrong",
				body: bytes.NewReader([]byte("data")),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := test.PutEndpoint()
			defer server.Close()
			c := &Connector{
				Client: tt.fields.Client,
				URL:    server.URL,
			}
			got, err := c.SimplePut(tt.args.path, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connector.SimplePut() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connector.SimplePut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_SimpleDelete(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	type args struct {
		path string
		body io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Success case",
			fields: fields{
				Client: FactoryHttpClient(),
			},
			args: args{
				path: "/delete",
				body: bytes.NewReader([]byte("data")),
			},
			want:    []byte("This is data"),
			wantErr: false,
		},
		{
			name: "Fail case: wrong path",
			fields: fields{
				Client: FactoryHttpClient(),
			},
			args: args{
				path: "/wrong",
				body: bytes.NewReader([]byte("data")),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := test.DeleteEndpoint()
			defer server.Close()
			c := &Connector{
				Client: tt.fields.Client,
				URL:    server.URL,
			}
			got, err := c.SimpleDelete(tt.args.path, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connector.SimpleDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connector.SimpleDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_SimpleDo(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	type args struct {
		method string
		path   string
		body   io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Success case",
			fields: fields{
				Client: FactoryHttpClient(),
			},
			args: args{
				method: "POST",
				path:   "/post",
				body:   bytes.NewReader([]byte("data")),
			},
			want:    []byte("This is data"),
			wantErr: false,
		},
		{
			name: "Fail case: wrong path",
			fields: fields{
				Client: FactoryHttpClient(),
			},
			args: args{
				method: "POST",
				path:   "/wrong",
				body:   bytes.NewReader([]byte("data")),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Fail case: wrong http method",
			fields: fields{
				Client: FactoryHttpClient(),
			},
			args: args{
				method: "GET",
				path:   "/post",
				body:   bytes.NewReader([]byte("data")),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := test.PostEndpoint()
			defer server.Close()
			c := &Connector{
				Client: tt.fields.Client,
				URL:    server.URL,
			}
			got, err := c.SimpleDo(tt.args.method, tt.args.path, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connector.SimpleDo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connector.SimpleDo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_DoWithHeader(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	type args struct {
		method             string
		path               string
		header             *http.Header
		body               io.Reader
		exceptedStatusCode StatusCodeRange
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Success case",
			fields: fields{
				Client: FactoryHttpClient(),
			},
			args: args{
				method: "GET",
				path:   "/get",
				body:   bytes.NewReader([]byte("data")),
				header: &http.Header{
					"test-header": []string{"value"},
				},
				exceptedStatusCode: DefaultStatusRange,
			},
			want:    []byte("This is data"),
			wantErr: false,
		},
		{
			name: "Fail case: wrong header",
			fields: fields{
				Client: FactoryHttpClient(),
			},
			args: args{
				method: "GET",
				path:   "/get",
				body:   bytes.NewReader([]byte("data")),
				header: &http.Header{
					"unknown": []string{"value"},
				},
				exceptedStatusCode: DefaultStatusRange,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Fail case: no header, traget endpoint is waiting for header to be set",
			fields: fields{
				Client: FactoryHttpClient(),
			},
			args: args{
				method:             "GET",
				path:               "/get",
				body:               bytes.NewReader([]byte("data")),
				exceptedStatusCode: DefaultStatusRange,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Fail case: forbidden char in method",
			fields: fields{
				Client: FactoryHttpClient(),
			},
			args: args{
				method: "GE\tT",
				path:   "/get",
				body:   bytes.NewReader([]byte("data")),
				header: &http.Header{
					"test-header": []string{"value"},
				},
				exceptedStatusCode: DefaultStatusRange,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := test.GetEndpointWithHeader()
			defer server.Close()
			c := &Connector{
				Client: tt.fields.Client,
				URL:    server.URL,
			}
			got, err := c.DoWithHeader(tt.args.method, tt.args.path, tt.args.header, tt.args.body, tt.args.exceptedStatusCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connector.DoWithHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connector.DoWithHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_Ping(t *testing.T) {
	type fields struct {
		Client       *http.Client
		pingEndpoint string
	}
	type args struct {
		t int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success case",
			fields: fields{
				Client:       FactoryHttpClient(),
				pingEndpoint: "/",
			},
			args: args{
				t: 1,
			},
			wantErr: false,
		},
		{
			name: "Fail case: wrong ping endpoint",
			fields: fields{
				Client:       FactoryHttpClient(),
				pingEndpoint: "/wrong",
			},
			args: args{
				t: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := test.GetPingEndpoint()
			defer server.Close()

			c := &Connector{
				Client:       tt.fields.Client,
				URL:          server.URL,
				pingEndpoint: tt.fields.pingEndpoint,
			}

			if err := c.Ping(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Connector.Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFactoryConnector(t *testing.T) {
	type args struct {
		config Conf
	}
	tests := []struct {
		name string
		args args
		want *Connector
	}{
		{
			name: "Success case",
			args: args{Conf{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FactoryConnector(tt.args.config); got == nil {
				t.Errorf("FactoryConnector() = %v, should not be nil", got)
			}
		})
	}
}
