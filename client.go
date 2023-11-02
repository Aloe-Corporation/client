// Package client will provide a reusable *http.Client with Let's Encrypt RootCAs for secured TLS connections
package client

import (
	"fmt"
	"net/http"

	"github.com/WeGetFunded-com/conf"
)

const (
	PREFIX_ENV = "CLIENT_"
	CLIENT_URL = PREFIX_ENV + "URL"
)

var (
	instanciatedClient = make(map[string]*http.Client)
	CONF_ENV_VAR       = []string{
		CLIENT_URL,
	}

	_ conf.OverrideConfigWithEnv = (*Conf)(nil)
)

type Conf struct {
	URL string `yaml:"url"`
}

// OverrideWithEnv override all config with environement variable.
func (c *Conf) OverrideWithEnv(prefix string) error {
	env := map[string]interface{}{
		prefix + CLIENT_URL: &c.URL,
	}

	err := conf.ParseEnvVarMap(env)
	if err != nil {
		return fmt.Errorf("fail to parse env client: %w", err)
	}

	return nil
}

// Get a reusable *http.Client with Let's Encrypt RootCAs for secured TLS connections.
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
