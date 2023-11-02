# client

This project is a module that handles HTTP client.

## Usage

### Factory

`FactoryHttpsClient()` build a new client with certifiacte

```go
clientUAExtractor := client.FactoryHttpsClient()
```

### Proxy Factory

`FactoryHttpsClient()` build a new client with certifiacte if not exist, the parameter is an identifier

```go
clientUAExtractor := client.ProxyFactoryHttpsClient("ua_extractor")
```
