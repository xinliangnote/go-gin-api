# Support HTTP and WebSocket Reverse Proxy

## Problem

The go-gin-api framework currently lacks native support for HTTP and WebSocket reverse proxy functionality, which limits its deployment flexibility in production environments where a proxy layer is required.

## Background

As a modular API framework based on Gin, go-gin-api includes WebSocket support via gorilla/websocket for real-time communication, but does not provide built-in reverse proxy capabilities for either HTTP or WebSocket connections. This means users must implement their own proxy solutions or rely on external proxies like Nginx.

## Solution

This implementation adds comprehensive reverse proxy support for both HTTP and WebSocket protocols directly within the framework, enabling:

1. **HTTP Reverse Proxy**: Proxy HTTP requests to backend services using Go's `net/http/httputil.ReverseProxy`
2. **WebSocket Reverse Proxy**: Forward WebSocket connections through the proxy while maintaining protocol handshakes and message relay

## Changes

### Core Proxy Implementation

- Added `internal/pkg/proxy/` package with:
  - `proxy.go`: Main proxy handler supporting both HTTP and WebSocket protocols
  - `websocket_proxier.go`: WebSocket-specific proxy logic handling upgrade requests
  - `http_proxier.go`: HTTP proxy using standard reverse proxy infrastructure

### Router Integration

- Updated `internal/router/` to include proxy configuration
- Added proxy initialization in `router/` with support for dynamic route configuration

### Configuration

- Extended configuration files (`configs/*.toml`) to support proxy settings:
  - Backend service URLs
  - Timeout configurations
  - Path-based routing rules
  - WebSocket upgrade headers

### Middleware Support

- Added proxy middleware to handle protocol detection
- Automatic routing to appropriate proxy handler based on connection type

## Technical Details

### HTTP Proxy

```go
type HTTPProxier struct {
    target *url.URL
    proxy  *httputil.ReverseProxy
}
```

Uses Go's built-in `ReverseProxy` with custom director to:
- Rewrite headers
- Preserve original request information
- Add trace ID forwarding

### WebSocket Proxy

```go
type WebSocketProxier struct {
    target        *url.URL
    dialer        *websocket.Dialer
    handshakeTimeout time.Duration
}
```

Handles WebSocket protocol by:
- Detecting WebSocket upgrade requests
- Performing handshake with backend
- Bidirectional message relay
- Proxy close/error handling

## Usage Example

### Configuration (dev_configs.toml)

```toml
[proxy]
enabled = true

[[proxy.routes]]
path_prefix = "/api/v1/backend"
backend_url = "http://backend-service:8080"
protocol = "http"
timeout = "10s"

[[proxy.routes]]
path_prefix = "/socket/ws"
backend_url = "ws://websocket-service:8081"
protocol = "websocket"
```

### Programmatic Setup

```go
proxyHandler := proxy.NewDefaultProxy()

// Add HTTP route
proxyHandler.AddRoute(proxy.Route{
    PathPrefix: "/api/v1/backend",
    BackendURL: "http://localhost:8080",
    Protocol:   proxy.ProtocolHTTP,
})

// Add WebSocket route
proxyHandler.AddRoute(proxy.Route{
    PathPrefix: "/socket",
    BackendURL: "ws://localhost:9090",
    Protocol:   proxy.ProtocolWebSocket,
})
```

## Benefits

1. **Simplified Deployment**: No need for external proxy configuration in simple setups
2. **Unified Logging**: Proxy traffic captured within the existing logging infrastructure
3. **Trace Integration**: Maintains trace IDs across proxy hops
4. **Performance**: Native Go implementation without external dependencies
5. **Flexibility**: Easy programmatic configuration for complex routing scenarios

## Testing

- Unit tests for HTTP proxy routing
- WebSocket proxy handoff tests
- Error handling verification
- Load testing for concurrent connections

## Breaking Changes

None. This feature is fully optional and enabled through configuration only.

## Compatibility

- Requires Go 1.16+ (for current Go standard library features)
- Compatible with existing WebSocket implementations
- Works with all existing middleware and features

## Future Enhancements

- Circuit breaker support for backend services
- Load balancing across multiple backends
- Request/response transformation
- WebSocket message filtering and modification

## Related Issues

- Closes #91

## Notes

- This implementation prioritizes correctness and ease of use over extreme performance
- For very high throughput scenarios, consider external solutions like Envoy or Nginx
- The proxy respects the framework's rate limiting, authentication, and logging systems
