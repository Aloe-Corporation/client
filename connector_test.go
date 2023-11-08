package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/Aloe-Corporation/client/test"
	"github.com/Aloe-Corporation/docker"
	"github.com/stretchr/testify/assert"
)

const (
	PING_TIMEOUT = 10
)

type doTestData struct {
	Conf            Conf
	Verbe, Path     string
	Header          *http.Header
	Body            io.Reader
	HasResponseBody bool
	ShouldFail      bool
	ExpectedErr     string
}

var (
	simpleGetTestCases = [...]doTestData{
		{ // success
			Conf: Conf{
				URL: "http://localhost:18080",
			},
			Path:            "/get",
			HasResponseBody: true,
			ShouldFail:      false,
		},
		{ // fail : invalid URL
			Conf: Conf{
				URL: "http://test.test",
			},
			ShouldFail:  true,
			ExpectedErr: "fail to execute HTTP request: Get \"http://test.test\": dial tcp: lookup test.test: no such host",
		},
		{ // fail : invalid URL
			Conf: Conf{
				URL: "http://goog\tle.com",
			},
			ShouldFail:  true,
			ExpectedErr: "can't create the request : parse \"http://goog\\tle.com\": net/url: invalid control character in URL",
		},
	}
)

func TestSimpleGet(t *testing.T) {
	for i, testCase := range simpleGetTestCases {
		t.Run("TestSimpleGet : "+strconv.Itoa(i), func(t *testing.T) {
			c := Connector{
				URL:    testCase.Conf.URL,
				Client: http.DefaultClient,
			}

			data, err := c.SimpleGet(testCase.Path, testCase.HasResponseBody)
			if testCase.ShouldFail {
				assert.Error(t, err)
				assert.Equal(t, testCase.ExpectedErr, err.Error())
			} else {
				assert.NoError(t, err)
				if testCase.HasResponseBody {
					assert.True(t, len(data) > 0)
				}
			}
		})
	}
}

var (
	simplePostTestCases = [...]doTestData{
		{ // success
			Conf: Conf{
				URL: "http://localhost:18080",
			},
			Path:            "/post",
			HasResponseBody: true,
			ShouldFail:      false,
		},
		{ // fail : invalid URL
			Conf: Conf{
				URL: "http://test.test",
			},
			ShouldFail:  true,
			ExpectedErr: "fail to execute HTTP request: Post \"http://test.test\": dial tcp: lookup test.test: no such host",
		},
		{ // fail : invalid URL
			Conf: Conf{
				URL: "http://goog\tle.com",
			},
			ShouldFail:  true,
			ExpectedErr: "can't create the request : parse \"http://goog\\tle.com\": net/url: invalid control character in URL",
		},
	}
)

func TestSimplePost(t *testing.T) {
	for i, testCase := range simplePostTestCases {
		t.Run("TestSimplePost : "+strconv.Itoa(i), func(t *testing.T) {
			c := Connector{
				URL:    testCase.Conf.URL,
				Client: http.DefaultClient,
			}

			data, err := c.SimplePost(testCase.Path, testCase.Body, testCase.HasResponseBody)
			if testCase.ShouldFail {
				assert.Error(t, err)
				assert.Equal(t, testCase.ExpectedErr, err.Error())
			} else {
				assert.NoError(t, err)
				if testCase.HasResponseBody {
					assert.True(t, len(data) > 0)
				}
			}
		})
	}
}

var (
	simplePutTestCases = [...]doTestData{
		{ // success
			Conf: Conf{
				URL: "http://localhost:18080",
			},
			Path:            "/put",
			HasResponseBody: true,
			ShouldFail:      false,
		},
		{ // fail : invalid URL
			Conf: Conf{
				URL: "http://test.test",
			},
			ShouldFail:  true,
			ExpectedErr: "fail to execute HTTP request: Put \"http://test.test\": dial tcp: lookup test.test: no such host",
		},
		{ // fail : invalid URL
			Conf: Conf{
				URL: "http://goog\tle.com",
			},
			ShouldFail:  true,
			ExpectedErr: "can't create the request : parse \"http://goog\\tle.com\": net/url: invalid control character in URL",
		},
	}
)

func TestSimplePut(t *testing.T) {
	for i, testCase := range simplePutTestCases {
		t.Run("TestSimplePut : "+strconv.Itoa(i), func(t *testing.T) {
			c := Connector{
				URL:    testCase.Conf.URL,
				Client: http.DefaultClient,
			}

			data, err := c.SimplePut(testCase.Path, testCase.Body, testCase.HasResponseBody)
			if testCase.ShouldFail {
				assert.Error(t, err)
				assert.Equal(t, testCase.ExpectedErr, err.Error())
			} else {
				assert.NoError(t, err)
				if testCase.HasResponseBody {
					assert.True(t, len(data) > 0)
				}
			}
		})
	}
}

var (
	simpleDeleteTestCases = [...]doTestData{
		{ // success
			Conf: Conf{
				URL: "http://localhost:18080",
			},
			Path:            "/delete",
			HasResponseBody: true,
			ShouldFail:      false,
		},
		{ // fail : invalid URL
			Conf: Conf{
				URL: "http://test.test",
			},
			ShouldFail:  true,
			ExpectedErr: "fail to execute HTTP request: Delete \"http://test.test\": dial tcp: lookup test.test: no such host",
		},
		{ // fail : invalid URL
			Conf: Conf{
				URL: "http://goog\tle.com",
			},
			ShouldFail:  true,
			ExpectedErr: "can't create the request : parse \"http://goog\\tle.com\": net/url: invalid control character in URL",
		},
	}
)

func TestSimpleDelete(t *testing.T) {
	for i, testCase := range simpleDeleteTestCases {
		t.Run("TestSimpleDelete : "+strconv.Itoa(i), func(t *testing.T) {
			c := Connector{
				URL:    testCase.Conf.URL,
				Client: http.DefaultClient,
			}

			data, err := c.SimpleDelete(testCase.Path, testCase.Body, testCase.HasResponseBody)
			if testCase.ShouldFail {
				assert.Error(t, err)
				assert.Equal(t, testCase.ExpectedErr, err.Error())
			} else {
				assert.NoError(t, err)
				if testCase.HasResponseBody {
					assert.True(t, len(data) > 0)
				}
			}
		})
	}
}

var (
	simpleDoTestCases = [...]doTestData{
		{ // success
			Conf: Conf{
				URL: "http://localhost:18080",
			},
			Verbe:           http.MethodGet,
			Path:            "/get",
			HasResponseBody: true,
			ShouldFail:      false,
		},
		{ // success
			Conf: Conf{
				URL: "http://localhost:18080",
			},
			Verbe:           http.MethodPost,
			Path:            "/post",
			HasResponseBody: true,
			ShouldFail:      false,
		},
		{ // success
			Conf: Conf{
				URL: "http://localhost:18080",
			},
			Verbe:           http.MethodPut,
			Path:            "/put",
			HasResponseBody: true,
			ShouldFail:      false,
		},
		{ // success
			Conf: Conf{
				URL: "http://localhost:18080",
			},
			Verbe:           http.MethodDelete,
			Path:            "/delete",
			HasResponseBody: true,
			ShouldFail:      false,
		},
		{ // fail : invalid URL
			Conf: Conf{
				URL: "http://test.test",
			},
			Verbe:       http.MethodGet,
			ShouldFail:  true,
			ExpectedErr: "fail to execute HTTP request: Get \"http://test.test\": dial tcp: lookup test.test: no such host",
		},
		{ // fail : invalid URL
			Conf: Conf{
				URL: "http://goog\tle.com",
			},
			Verbe:       http.MethodGet,
			ShouldFail:  true,
			ExpectedErr: "can't create the request : parse \"http://goog\\tle.com\": net/url: invalid control character in URL",
		},
	}
)

func TestSimpleDo(t *testing.T) {
	for i, testCase := range simpleDoTestCases {
		t.Run("TestSimpleDo : "+strconv.Itoa(i), func(t *testing.T) {
			c := Connector{
				URL:    testCase.Conf.URL,
				Client: http.DefaultClient,
			}

			data, err := c.SimpleDo(testCase.Verbe, testCase.Path, testCase.Body, testCase.HasResponseBody)
			if testCase.ShouldFail {
				assert.Error(t, err)
				assert.Equal(t, testCase.ExpectedErr, err.Error())
			} else {
				assert.NoError(t, err)
				if testCase.HasResponseBody {
					assert.True(t, len(data) > 0)
				}
			}
		})
	}
}

var (
	connectorDoWithHeaderTestCases = [...]doTestData{
		{ // success
			Conf: Conf{
				URL: "http://localhost:18080",
			},
			Verbe: http.MethodGet,
			Path:  "/header",
			Header: &http.Header{
				"foo": []string{"bar"},
			},
			HasResponseBody: true,
			ShouldFail:      false,
		},
		{ // fail case default header
			Conf: Conf{
				URL: "http://localhost:18080",
			},
			Verbe:           http.MethodGet,
			Path:            "/header",
			Header:          &http.Header{},
			HasResponseBody: true,
			ShouldFail:      true,
			ExpectedErr:     "400 fail request",
		},
		{ // fail case empty header
			Conf: Conf{
				URL: "http://localhost:18080",
			},
			Verbe: http.MethodGet,
			Path:  "/header",
			Header: &http.Header{
				"foo": []string{},
			},
			HasResponseBody: true,
			ShouldFail:      true,
			ExpectedErr:     "400 fail request",
		},
		{ // fail case no header
			Conf: Conf{
				URL: "http://localhost:18080",
			},
			Verbe:           http.MethodGet,
			Path:            "/header",
			Header:          nil,
			HasResponseBody: true,
			ShouldFail:      true,
			ExpectedErr:     "400 fail request",
		},
		{ // fail : invalid URL
			Conf: Conf{
				URL: "http://test.test",
			},
			Verbe:       http.MethodGet,
			ShouldFail:  true,
			ExpectedErr: "fail to execute HTTP request: Get \"http://test.test\": dial tcp: lookup test.test: no such host",
		},
		{ // fail : invalid URL
			Conf: Conf{
				URL: "http://goog\tle.com",
			},
			Verbe:       http.MethodGet,
			ShouldFail:  true,
			ExpectedErr: "can't create the request : parse \"http://goog\\tle.com\": net/url: invalid control character in URL",
		},
	}
)

func TestConnectorDoWithHeader(t *testing.T) {
	for i, testCase := range connectorDoWithHeaderTestCases {
		t.Run("TestConnectorDoWithHeader : "+strconv.Itoa(i), func(t *testing.T) {
			c := Connector{
				URL:    testCase.Conf.URL,
				Client: http.DefaultClient,
			}

			data, err := c.DoWithHeader(testCase.Verbe, testCase.Path, testCase.Header, testCase.Body, DefaultStatusRange, testCase.HasResponseBody)
			if testCase.ShouldFail {
				assert.Error(t, err)
				assert.Equal(t, testCase.ExpectedErr, err.Error())
			} else {
				assert.NoError(t, err)
				if testCase.HasResponseBody {
					assert.True(t, len(data) > 0)
				}
			}
		})
	}
}

var (
	doWithStatusCheckTestCases = [...]doTestData{
		{ // success : ping google.com
			Conf: Conf{
				URL: "http://google.com",
			},
			Verbe:           http.MethodGet,
			Path:            "",
			Body:            nil,
			HasResponseBody: true,
			ShouldFail:      false,
		},
		{ // fail : invalid URL
			Conf: Conf{
				URL: "http://test.test",
			},
			Verbe:           http.MethodGet,
			Path:            "",
			Body:            nil,
			HasResponseBody: false,
			ShouldFail:      true,
			ExpectedErr:     "fail to execute HTTP request: Get \"http://test.test\": dial tcp: lookup test.test: no such host",
		},
	}
)

func TestDoWithStatusCheck(t *testing.T) {
	for i, testCase := range doWithStatusCheckTestCases {
		t.Run("TestDoWithStatusCheck : "+strconv.Itoa(i), func(t *testing.T) {
			c := Connector{
				URL:    testCase.Conf.URL,
				Client: http.DefaultClient,
			}

			req, err := http.NewRequest(testCase.Verbe, c.URL+testCase.Path, testCase.Body)
			assert.NoError(t, err)

			data, err := c.DoWithStatusCheck(req, DefaultStatusRange, testCase.HasResponseBody)
			if testCase.ShouldFail {
				assert.Error(t, err)
				assert.Equal(t, testCase.ExpectedErr, err.Error())
			} else {
				assert.NoError(t, err)
				if testCase.HasResponseBody {
					assert.True(t, len(data) > 0)
				}
			}
		})
	}
}

type ConnectorPingTestData struct {
	Conf       Conf
	ShouldFail bool
}

var connectorPingTestCases = [...]ConnectorPingTestData{
	{ // Success
		Conf: Conf{
			URL: "http://localhost:18080",
		},
		ShouldFail: false,
	},
	{ // Fail to ping
		Conf: Conf{
			URL: "http://localhost:666",
		},
		ShouldFail: true,
	},
}

func TestConnectorPing(t *testing.T) {
	for i, testCase := range connectorPingTestCases {
		t.Run("TestConnectorPing : "+strconv.Itoa(i), func(t *testing.T) {
			c := Connector{
				URL:    testCase.Conf.URL,
				Client: http.DefaultClient,
			}

			err := c.Ping(PING_TIMEOUT)
			if testCase.ShouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

var factoryConnectorTestsCases = [...]Conf{
	{
		URL: "http://localhost:8080",
	},
}

func TestFactoryConnector(t *testing.T) {
	for i, testCase := range factoryConnectorTestsCases {
		t.Run("TestFactoryConnector : "+strconv.Itoa(i), func(t *testing.T) {
			result := FactoryConnector(testCase)
			assert.NotNil(t, result)
		})
	}
}

func TestMain(m *testing.M) {
	// Up docker-compose
	dc := docker.Compose{PathFile: test.DockerCompose}
	err := dc.Up()
	if err != nil {
		_ = dc.Down()
		panic(fmt.Errorf("fail to up docker-compose: %w", err))
	}

	// Run
	r := m.Run()

	// Down docker-compose
	if err := dc.Down(); err != nil {
		panic(fmt.Errorf("fail to down docker-compose: %w", err))
	}
	os.Exit(r)
}
