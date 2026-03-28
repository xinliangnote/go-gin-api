# Implementation and Testing Guide

## Step 1: Create Proxy Package Structure

```bash
mkdir -p /home/calelin/dev/go-gin-api/internal/pkg/proxy
```

## Step 2: Create HTTP Proxier

Create `internal/pkg/proxy/http_proxier.go`:

```go
package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/xinliangnote/go-gin-api/pkg/errors"
	"go.uber.org/zap"
)

type HTTPProxier struct {
	target *url.URL
	logger *zap.Logger
}

func NewHTTPProxier(target string, logger *zap.Logger) (*HTTPProxier, error) {
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse target URL")
	}

	return &HTTPProxier{
		target: targetURL,
		logger: logger,
	}, nil
}

func (h *HTTPProxier) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	proxy := httputil.NewSingleHostReverseProxy(h.target)
	
	// Customize error handler
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		h.logger.Error("proxy error", 
			zap.String("target", h.target.String()),
			zap.String("path", r.URL.Path),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Backend unreachable"))
	}

	// Modify request
	proxy.Transport = http.DefaultTransport
	proxy.Transport.(*http.Transport).ResponseHeaderTimeout = 10 * time.Second

	proxy.ServeHTTP(w, r)
}
```

## Step 3: Create WebSocket Proxier

Create `internal/pkg/proxy/websocket_proxier.go`:

```go
package proxy

import (
	"bufio"
	"context"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/xinliangnote/go-gin-api/pkg/errors"
	"go.uber.org/zap"
)

type WebSocketProxier struct {
	target           *url.URL
	dialer           *websocket.Dialer
	handshakeTimeout time.Duration
	logger           *zap.Logger
}

func NewWebSocketProxier(target string, logger *zap.Logger) (*WebSocketProxier, error) {
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse target URL")
	}

	return &WebSocketProxier{
		target: targetURL,
		dialer: &websocket.Dialer{
			HandshakeTimeout: 5 * time.Second,
		},
		handshakeTimeout: 5 * time.Second,
		logger:           logger,
	}, nil
}

func (w *WebSocketProxier) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// Hijack connection
	hijacker, ok := resp.(http.Hijacker)
	if !ok {
		w.logger.Error("connection does not support hijacking")
		http.Error(resp, "Cannot hijack connection", http.StatusInternalServerError)
		return
	}

	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		w.logger.Error("failed to hijack connection", zap.Error(err))
		http.Error(resp, "Failed to hijack connection", http.StatusInternalServerError)
		return
	}
	defer clientConn.Close()

	// Connect to backend
	backendURL := w.target.String() + req.URL.Path + "?" + req.URL.RawQuery
	backendConn, _, err := w.dialer.Dial(backendURL, req.Header)
	if err != nil {
		w.logger.Error("failed to dial backend", 
			zap.String("backend", backendURL),
			zap.Error(err),
		)
		return
	}
	defer backendConn.Close()

	// Bidirectional message relay
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go w.relayWebSocketToTCP(ctx, backendConn, clientConn)
	w.relayTCPToWebSocket(backendConn, clientConn)
}

func (w *WebSocketProxier) relayWebSocketToWebSocket(ctx context.Context, src, dst *websocket.Conn) {
	defer src.Close()
	defer dst.Close()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			messageType, data, err := src.ReadMessage()
			if err != nil {
				return
			}
			if err := dst.WriteMessage(messageType, data); err != nil {
				return
			}
		}
	}
}

func (w *WebSocketProxier) relayTCPToWebSocket(src *websocket.Conn, dst net.Conn) {
	reader := bufio.NewReader(dst)

	for {
		messageType, data, err := src.ReadMessage()
		if err != nil {
			return
		}

		if _, err := dst.Write(data); err != nil {
			return
		}

		// Read response from client if needed
		// (simplified implementation)
	}
}

func (w *WebSocketProxier) relayWebSocketToTCP(ctx context.Context, src *websocket.Conn, dst net.Conn) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, data, err := src.ReadMessage()
			if err != nil {
				return
			}

			if _, err := dst.Write(data); err != nil {
				return
			}
		}
	}
}
```

## Step 4: Create Main Proxy Handler

Create `internal/pkg/proxy/proxy.go`:

```go
package proxy

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/pkg/errors"
	"go.uber.org/zap"
)

const (
	ProtocolHTTP     = "http"
	ProtocolWebSocket = "websocket"
)

type Route struct {
	PathPrefix string
	BackendURL string
	Protocol   string
}

type Proxy struct {
	routes  []Route
	httpProxiers    map[string]*HTTPProxier
	wsProxiers      map[string]*WebSocketProxier
	logger *zap.Logger
}

func NewDefaultProxy() *Proxy {
	return &Proxy{
		httpProxiers: make(map[string]*HTTPProxier),
		wsProxiers:   make(map[string]*WebSocketProxier),
		logger:       zap.NewNop(), // Replace with actual logger
	}
}

func (p *Proxy) AddRoute(route Route) error {
	p.routes = append(p.routes, route)

	switch route.Protocol {
	case ProtocolHTTP:
		proxier, err := NewHTTPProxier(route.BackendURL, p.logger)
		if err != nil {
			return err
		}
		p.httpProxiers[route.PathPrefix] = proxier

	case ProtocolWebSocket:
		proxier, err := NewWebSocketProxier(route.BackendURL, p.logger)
		if err != nil {
			return err
		}
		p.wsProxiers[route.PathPrefix] = proxier

	default:
		return errors.New("unsupported protocol: %s", route.Protocol)
	}

	return nil
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range p.routes {
		if len(r.URL.Path) >= len(route.PathPrefix) && 
		   r.URL.Path[:len(route.PathPrefix)] == route.PathPrefix {
			
			switch route.Protocol {
			case ProtocolHTTP:
				if proxier, ok := p.httpProxiers[route.PathPrefix]; ok {
					proxier.ServeHTTP(w, r)
					return
				}
			case ProtocolWebSocket:
				if proxier, ok := p.wsProxiers[route.PathPrefix]; ok {
					proxier.ServeHTTP(w, r)
					return
				}
			}
		}
	}

	http.NotFound(w, r)
}
```

## Step 5: Integrate with Router

Update `internal/router/router.go` to add proxy routes after existing routes:

```go
// Add this to your router setup
proxyHandler := proxy.NewDefaultProxy()

proxyHandler.AddRoute(proxy.Route{
	PathPrefix: "/api/v1/backend",
	BackendURL: "http://localhost:8080",
	Protocol:   proxy.ProtocolHTTP,
})

proxyHandler.AddRoute(proxy.Route{
	PathPrefix: "/socket",
	BackendURL: "ws://localhost:9090",
	Protocol:   proxy.ProtocolWebSocket,
})

// Register proxy handler
r.mux.Any("/api/v1/backend/*path", func(ctx core.Context) {
	proxyHandler.ServeHTTP(ctx.Writer(), ctx.Request())
})
r.mux.GET("/socket/*path", func(ctx core.Context) {
	proxyHandler.ServeHTTP(ctx.Writer(), ctx.Request())
})
```

## Step 6: Local Testing

### Test 1: Simple HTTP Backend Server

Create `test_backend.go`:

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/backend/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from backend! Path: %s", r.URL.Path)
	})

	http.HandleFunc("/socket/ws", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("WebSocket endpoint"))
	})

	fmt.Println("Backend server running on :8080")
	http.ListenAndServe(":8080", nil)
}
```

Run backend:
```bash
go run test_backend.go
```

### Test 2: HTTP Proxy Test

```bash
curl http://localhost:8080/api/v1/backend/test
```

Then through proxy (assuming main app runs on :9999):
```bash
curl http://localhost:9999/api/v1/backend/test
```

### Test 3: WebSocket Test

Create `test_websocket.html`:

```html
<!DOCTYPE html>
<html>
<head>
    <title>WebSocket Test</title>
</head>
<body>
    <script>
        const ws = new WebSocket('ws://localhost:9090/socket/ws');
        
        ws.onopen = () => {
            console.log('WebSocket connected');
            ws.send('Hello from client');
        };
        
        ws.onmessage = (event) => {
            console.log('Received:', event.data);
        };
        
        ws.onerror = (error) => {
            console.error('WebSocket error:', error);
        };
    </script>
</body>
</html>
```

### Test 4: Load Testing with Apache Bench

```bash
# HTTP load test
ab -n 1000 -c 10 http://localhost:9999/api/v1/backend/test
```

## Step 7: Run Main Application

```bash
cd /home/calelin/dev/go-gin-api
go run main.go
```

## Step 8: Verify Proxy Functionality

```bash
# Test HTTP proxy
curl -v http://localhost:9999/api/v1/backend/test

# Check logs for proxy activity
tail -f logs/go-gin-api-access.log
```

## Common Issues & Solutions

1. **Connection Refused**: Ensure backend servers are running
2. **WebSocket Upgrade Failures**: Verify CORS headers are allowed
3. **Timeout Errors**: Increase timeout values in configuration
4. **Path Not Found**: Check path prefix matching logic

## Performance Benchmarks

Test with varying concurrency:
```bash
# Single connection
ab -n 100 -c 1 http://localhost:9999/api/v1/backend/test

# 10 concurrent connections
ab -n 1000 -c 10 http://localhost:9999/api/v1/backend/test

# 100 concurrent connections
ab -n 10000 -c 100 http://localhost:9999/api/v1/backend/test
```

Monitor resource usage:
```bash
htop
```

## Cleanup

```bash
# Stop applications
pkill -f "go run main.go"
pkill -f "go run test_backend.go"
```