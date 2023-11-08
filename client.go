package client

import (
	"net/http"
)

var (
	instanciatedClient = make(map[string]*http.Client)
)

// Conf for the client. All parameters are required.
type Conf struct {
	URL          string `yaml:"url"`           // Base url of the target HTTP server such as https://myserver.com
	PingEndpoint string `yaml:"ping_endpoint"` // Path of the ping endpoint of the target HTTP server
}

// ProxyFactoryHttpClient creates a new client if it does not exists in
// `instanciatedClient` map. If the client key is already defined, the
// function returns the associated client.
func ProxyFactoryHttpClient(key string) *http.Client {
	if c, ok := instanciatedClient[key]; ok {
		return c
	}

	c := FactoryHttpClient()
	instanciatedClient[key] = c

	return c
}

// FactoryHttpClient returns a fresh new http.Client instance.
func FactoryHttpClient() *http.Client {
	return &http.Client{}
}
