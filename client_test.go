package client

import (
	"strconv"
	"testing"

	"github.com/WeGetFunded-com/conf"
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

type confOverrideWithEnvTestData struct {
	Prefix      string
	EnvVarSet   map[string]string
	InitialConf Conf
	IsFail      bool
	ResultConf  Conf
}

var confOverrideWithEnvTestsCases = []confOverrideWithEnvTestData{
	{
		Prefix:      "CONF_TC_1_",
		EnvVarSet:   map[string]string{},
		InitialConf: Conf{},
		ResultConf:  Conf{},
	},
	{
		Prefix: "CONF_TC_2_",
		EnvVarSet: map[string]string{
			"CONF_TC_2_" + CLIENT_URL: "http://127.0.0.1:8081",
		},
		InitialConf: Conf{
			URL: "http://localhost:8080",
		},
		ResultConf: Conf{
			URL: "http://127.0.0.1:8081",
		},
	},
}

func TestConfOverrideWithEnvVar(t *testing.T) {
	for i, testCase := range confOverrideWithEnvTestsCases {
		t.Run("Case "+strconv.Itoa(i)+"]", func(t *testing.T) {
			conf.SetEnvWithMap(testCase.EnvVarSet)
			config := testCase.InitialConf

			err := config.OverrideWithEnv(testCase.Prefix)

			if testCase.IsFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.ResultConf, config)
			}

			conf.UnsetEnvWithMap(testCase.EnvVarSet)
		})
	}
}
