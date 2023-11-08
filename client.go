package client

import (
	"net/http"
)

var (
	instanciatedClient = make(map[string]*http.Client)
)

// ProxyFactoryHTTPClient creates a new client if it does not exists in
// `instanciatedClient` map. If the client key is already defined, the
// function returns the associated client.
func ProxyFactoryHTTPClient(key string) *http.Client {
	if c, ok := instanciatedClient[key]; ok {
		return c
	}

	c := FactoryHTTPClient()
	instanciatedClient[key] = c

	return c
}

// FactoryHTTPClient returns a fresh new http.Client instance.
func FactoryHTTPClient() *http.Client {
	return &http.Client{}
}
