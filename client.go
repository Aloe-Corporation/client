package client

import (
	"net/http"
)

var (
	instanciatedClient = make(map[string]*http.Client)
)

type Conf struct {
	URL          string `yaml:"url"`
	PingEndpoint string `yaml:"ping_endpoint"`
}

func ProxyFactoryHttpClient(key string) *http.Client {
	if c, ok := instanciatedClient[key]; ok {
		return c
	}

	c := FactoryHttpClient()
	instanciatedClient[key] = c

	return c
}

func FactoryHttpClient() *http.Client {
	return &http.Client{}
}
