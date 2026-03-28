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
	
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = h.target.Host
		req.URL.Scheme = h.target.Scheme
		req.URL.Host = h.target.Host
	}
	
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		h.logger.Error("proxy error", 
			zap.String("target", h.target.String()),
			zap.String("path", r.URL.Path),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Backend unreachable"))
	}

	transport := &http.Transport{
		ResponseHeaderTimeout: 10 * time.Second,
	}
	proxy.Transport = transport

	proxy.ServeHTTP(w, r)
}