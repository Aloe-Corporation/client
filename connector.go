package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// tickInterval is interval between each request sent by the Ping() method.
	tickInterval = 50 * time.Millisecond
)

var (
	// DefaultStatusRange that encompass most of all cases.
	DefaultStatusRange = StatusCodeRange{
		Min: http.StatusOK,
		Max: http.StatusBadRequest,
	}
)

// Conf for the client. All parameters are required.
type Conf struct {
	URL          string `yaml:"url"`           // Base url of the target HTTP server such as https://myserver.com
	PingEndpoint string `yaml:"ping_endpoint"` // Path of the ping endpoint of the target HTTP server
}

// StatusCodeRange defines the range of valid status codes.
// Status codes within the range will be considered as expected
// codes when received from the target API.
type StatusCodeRange struct {
	Min int // Lower bound
	Max int // Max bound excluded
}

// Connector is a supercharged HTTP client.
// It embed a native http.Client so it can be used as native client.
type Connector struct {
	*http.Client        // Natice http client
	URL          string // Base url of the target HTTP server such as https://myserver.com
	pingEndpoint string // Path of the ping endpoint of the target HTTP server
}

// SimpleGet eases the Connector.SimpleDo use.
// You have to specify the path.
// The StatusCodeRange use in Connector.DoWithStatusCheck will be DefautStatusRange [200,400[.
func (c *Connector) SimpleGet(path string) ([]byte, error) {
	return c.SimpleDo(http.MethodGet, path, nil)
}

// SimplePost eases the Connector.SimpleDo use.
// You have to specify the path and the request body as an io.Reader.
// The StatusCodeRange use in Connector.DoWithStatusCheck will be DefautStatusRange [200,400[.
func (c *Connector) SimplePost(path string, body io.Reader) ([]byte, error) {
	return c.SimpleDo(http.MethodPost, path, body)
}

// SimplePut eases the Connector.SimpleDo use.
// You have to specify the path and the request body as an io.Reader.
// The StatusCodeRange use in Connector.DoWithStatusCheck will be DefautStatusRange [200,400[.
func (c *Connector) SimplePut(path string, body io.Reader) ([]byte, error) {
	return c.SimpleDo(http.MethodPut, path, body)
}

// SimpleDelete eases the Connector.SimpleDo use.
// You have to specify the path and the request body as an io.Reader.
// The StatusCodeRange use in Connector.DoWithStatusCheck will be DefautStatusRange [200,400[.
func (c *Connector) SimpleDelete(path string, body io.Reader) ([]byte, error) {
	return c.SimpleDo(http.MethodDelete, path, body)
}

// SimpleDo eases the Connector.SimpleDo use.
// You have to specify the method, the path and the body as an io.Reader.
// The StatusCodeRange use in Connector.DoWithStatusCheck will be DefautStatusRange [200,400[.
func (c *Connector) SimpleDo(method, path string, body io.Reader) ([]byte, error) {
	return c.DoWithHeader(method, path, nil, body, DefaultStatusRange)
}

// DoWithHeader  eases the Connector.DoWithStatusCheck use.
// You have to specify the method, the path, the header, the body, the excepted status range.
// The excepted status range Min will be included and Max will be excluded
func (c *Connector) DoWithHeader(method, path string, header *http.Header, body io.Reader, exceptedStatusCode StatusCodeRange) ([]byte, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, c.URL+path, body)
	if err != nil {
		return nil, fmt.Errorf("can't create the request : %w", err)
	}

	if header != nil {
		req.Header = *header
	}

	return c.DoWithStatusCheck(req, exceptedStatusCode)
}

// DoWithStatusCheck a HTTP request with the given request.
// The caller should use Connector.URL as base URL when building the request.
// You have to provide a status code range to validate if the request was succesfull.
func (c *Connector) DoWithStatusCheck(req *http.Request, exceptedStatusCode StatusCodeRange) ([]byte, error) {
	response, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fail to execute HTTP request: %w", err)
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read response body : %w", err)
	}

	if response.StatusCode < exceptedStatusCode.Min || response.StatusCode >= exceptedStatusCode.Max {
		return nil, &FailRequestError{Code: response.StatusCode, ResponseBody: data}
	}

	return data, nil
}

// Ping sends one ping every 50ms with timeout of t second, it ends if the ping is a success or timeout.
func (c *Connector) Ping(t int) error {
	ticker := time.NewTicker(tickInterval)
	defer ticker.Stop()

	timeout := time.After(time.Duration(t) * time.Second)

	for {
		select {
		case <-ticker.C:
			_, err := c.SimpleGet(c.pingEndpoint)
			if err == nil {
				return nil
			}

		case <-timeout:
			return fmt.Errorf("can't ping API (%s): timeout after %d s", c.URL, t)
		}
	}
}

// FactoryConnector instantiate and return a *Connector.
// Call *Connector.Ping() to ensure that the target API is available.
func FactoryConnector(config Conf) *Connector {
	c := &Connector{
		URL:          config.URL,
		pingEndpoint: config.PingEndpoint,
		Client:       FactoryHTTPClient(),
	}

	return c
}
