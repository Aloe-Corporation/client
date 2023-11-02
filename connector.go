package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	TICK_INTERVAL = 50 * time.Millisecond
)

var (
	DefaultStatusRange = StatusCodeRange{
		Min: http.StatusOK,
		Max: http.StatusBadRequest,
	}
)

type FailRequestError struct {
	Code int
}

func (e *FailRequestError) Error() string {
	return fmt.Sprintf("%d fail request", e.Code)
}

type StatusCodeRange struct {
	Min int
	Max int
}

type Connector struct {
	*http.Client
	URL string
}

// SimpleGet allow to use Connector.SimpleDo easily.
// You have to specify the path and if the response have a body, if not []byte will be not nil and error will be nil.
// The StatusCodeRange use in Connector.DoWithStatusCheck will be DefautStatusRange [200,400[.
func (c *Connector) SimpleGet(path string, hasResponseBody bool) ([]byte, error) {
	return c.SimpleDo(http.MethodGet, path, nil, hasResponseBody)
}

// SimplePost allow to use Connector.SimpleDo easily.
// You have to specify the path, the request body and if the response have a body, if not []byte will be not nil and error will be nil.
// The StatusCodeRange use in Connector.DoWithStatusCheck will be DefautStatusRange [200,400[.
func (c *Connector) SimplePost(path string, body io.Reader, hasResponseBody bool) ([]byte, error) {
	return c.SimpleDo(http.MethodPost, path, body, hasResponseBody)
}

// SimplePut allow to use Connector.SimpleDo easily.
// You have to specify the path, the request body and if the response have a body, if not []byte will be not nil and error will be nil.
// The StatusCodeRange use in Connector.DoWithStatusCheck will be DefautStatusRange [200,400[.
func (c *Connector) SimplePut(path string, body io.Reader, hasResponseBody bool) ([]byte, error) {
	return c.SimpleDo(http.MethodPut, path, body, hasResponseBody)
}

// SimpleDelete allow to use Connector.SimpleDo easily.
// You have to specify the path, the request body and if the response have a body, if not []byte will be not nil and error will be nil.
// The StatusCodeRange use in Connector.DoWithStatusCheck will be DefautStatusRange [200,400[.
func (c *Connector) SimpleDelete(path string, body io.Reader, hasResponseBody bool) ([]byte, error) {
	return c.SimpleDo(http.MethodDelete, path, body, hasResponseBody)
}

// SimpleDo allow to use Connector.DoWithHeader easily.
// You have to specify the method, the path, the body and if the response have a body, if not []byte will be not nil and error will be nil.
// The StatusCodeRange use in Connector.DoWithStatusCheck will be DefautStatusRange [200,400[.
func (c *Connector) SimpleDo(method, path string, body io.Reader, hasResponseBody bool) ([]byte, error) {
	return c.DoWithHeader(method, path, nil, body, DefaultStatusRange, hasResponseBody)
}

// DoWithHeader allow to use Connector.DoWithStatusCheck easily.
// You have to specify the method, the path, the header, the body, the excepted status range.
// The excepted status range Min will be included and Max will be excluded
// You have to specify if the response have a body if not []byte will be not nil and error will be nil.
func (c *Connector) DoWithHeader(method, path string, header *http.Header, body io.Reader, exceptedStatusCode StatusCodeRange, hasResponseBody bool) ([]byte, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, c.URL+path, body)
	if err != nil {
		return nil, fmt.Errorf("can't create the request : %w", err)
	}

	if header != nil {
		req.Header = *header
	}

	return c.DoWithStatusCheck(req, exceptedStatusCode, hasResponseBody)
}

// DoWithStatusCheck a HTTP request with the given request.
// The caller should use Connector.URL as base URL when building the request.
// You have to provide a status code range to validate if the request was succesfull.
// You have to specify if the response has a body, if not []byte will be not nil and error will be nil.
func (c *Connector) DoWithStatusCheck(req *http.Request, exceptedStatusCode StatusCodeRange, hasResponseBody bool) ([]byte, error) {
	response, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fail to execute HTTP request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode < exceptedStatusCode.Min || response.StatusCode >= exceptedStatusCode.Max {
		return nil, &FailRequestError{Code: response.StatusCode}
	}

	data := make([]byte, 0)
	if hasResponseBody {
		data, err = io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("can't read response body : %w", err)
		}
	}

	if err := response.Body.Close(); err != nil {
		return nil, fmt.Errorf("can't close response body : %w", err)
	}

	return data, nil
}

// Ping test one ping every 50ms with timeout of t second, it end if the ping is a success or timeout.
// Use ping on the We Get Funded API only.
func (c *Connector) Ping(t int) error {
	ticker := time.NewTicker(TICK_INTERVAL)
	defer ticker.Stop()

	timeout := time.After(time.Duration(t) * time.Second)

	msg := struct {
		Service   string `json:"service"`
		Copyright string `json:"copyright"`
	}{}

	for {
		select {
		case <-ticker.C:
			data, httpErr := c.SimpleGet("/", true)
			jsonErr := json.Unmarshal(data, &msg)

			if httpErr == nil && jsonErr == nil && msg.Service != "" {
				return nil
			}

		case <-timeout:
			return fmt.Errorf("can't ping API (%s): timeout after %d s", c.URL, t)
		}
	}
}

// FactoryConnector instantiate and return a *Connector.
// You MUST use *Connector.Ping() BEFORE using it.
func FactoryConnector(config Conf) *Connector {
	c := &Connector{
		URL:    config.URL,
		Client: FactoryHttpClient(),
	}

	return c
}
