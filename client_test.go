package client

import (
	"testing"
)

func TestProxyFactoryHTTPClient(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		args      args
		newClient bool
	}{
		{
			name: "Success case: new client",
			args: args{
				key: "c1",
			},
			newClient: true,
		},
		{
			name: "Success case: already instanciated client ",
			args: args{
				key: "c0",
			},
			newClient: false,
		},
	}

	// Instanciate a client before test run
	instanciatedClient["c0"] = FactoryHTTPClient()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProxyFactoryHTTPClient(tt.args.key)
			if tt.newClient {
				if got == instanciatedClient["c0"] {
					t.Errorf("ProxyFactoryHTTPClient() = %v, should be a new client but was equal to already instanciated one.", got)
				}
			} else {
				if got == nil {
					t.Errorf("ProxyFactoryHTTPClient() = %v, client should not be nil", got)
				}
			}
		})
	}
}
