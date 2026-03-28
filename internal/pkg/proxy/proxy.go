package proxy

import (
	"net/http"
	"strings"

	"github.com/xinliangnote/go-gin-api/pkg/errors"
	"go.uber.org/zap"
)

const (
	ProtocolHTTP      = "http"
	ProtocolWebSocket = "websocket"
)

type Route struct {
	PathPrefix string
	BackendURL string
	Protocol   string
}

type Proxy struct {
	routes       []Route
	httpProxiers map[string]*HTTPProxier
	wsProxiers   map[string]*WebSocketProxier
	logger       *zap.Logger
}

func NewProxy(logger *zap.Logger) *Proxy {
	return &Proxy{
		httpProxiers: make(map[string]*HTTPProxier),
		wsProxiers:   make(map[string]*WebSocketProxier),
		logger:       logger,
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
		return errors.Errorf("unsupported protocol: %s", route.Protocol)
	}

	return nil
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path

	for _, route := range p.routes {
		prefix := strings.TrimSuffix(route.PathPrefix, "/")
		if strings.HasPrefix(requestPath, prefix) {
			
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

func (p *Proxy) GetRoutes() []Route {
	return p.routes
}