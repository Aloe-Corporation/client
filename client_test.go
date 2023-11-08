package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testProxyFactoryHttpClient struct {
	key           string
	isInstanciate bool
}

var testProxyFactoryHttpClientData = []testProxyFactoryHttpClient{
	{ // Success case
		key:           "c1",
		isInstanciate: true,
	},
	{ // test a second time with the same client name
		key:           "c1",
		isInstanciate: true,
	},
}

func TestFactoryHttpClient(t *testing.T) {
	for _, d := range testProxyFactoryHttpClientData {
		c := ProxyFactoryHttpClient(d.key)
		assert.NotZero(t, d.isInstanciate, c != nil)
	}
}
