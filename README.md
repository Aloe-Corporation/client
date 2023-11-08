# Client

The client module defines a supercharged HTTP Client that simplify HTTP requests. It aims to minify the code repetition
that `net/http` module creates when dealing with multiple API calls.

This project has been developped by the [Aloe](https://www.aloe-corp.com/) team and is now open source.

![tests](https://github.com/Aloe-Corporation/client/actions/workflows/tests.yml/badge.svg)

## Overview

Client supercharge the native http client.

The client module provides:

- Unique indication of the target API base URL
- Simple error checking and response body reading
- Ping method with timeout
- Infinite compatibility because it embed a native `net/http` client
- Proxy of instanciated clients

## Concepts

### One client by target

One of the main concept is that one client is bound to one API. We do this to simplify HTTP calls by only setting the endpoint path as function parameter and not the base URL + the path.

### Status code range

To ease the error checking we've defined a `StatusCodeRange` struct.

```go
type StatusCodeRange struct {
	Min int // Lower bound
	Max int // Max bound excluded
}
```
All status code received from target API within the range will be considered as valid response.


## Usage

### Instanciation

Use the `Conf` struct then call the `FactoryConnector(conf Conf)` function to instanciate a new HTTP connector.
Don't forget to call the  `*Connector.Ping(t int)`` method to verify that the target API is availble.

```go
conf := Conf{
    URL: "https://myserver.com"
    PingEndpoint: "/ping"
}

connector := FactoryConnector(conf)

if err := connector.Ping(10); err != nil {
    return fmt.Errorf("target api is not available: %w", err)
}
```

### Send HTTP requests
Once the connector is instanciates, you can use it to make your HTTP requests.

Example:

Send the GET "https://myserver.com/data" will result in the following code:

``` go
responseBody, err := connector.SimpleGet("/data")
if err != nil {
    return nil, fmt.Errorf("can't call /data endpoint: %w", err)
}

// Do what you wan't with the responseBody such as unmarshalling to JSON etc.
fmt.Println(string(responseBody))
```

## Contributing

## Licence