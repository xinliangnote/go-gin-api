package proxy

import (
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go w.relayWebSocketToTCP(ctx, backendConn, clientConn)
	w.relayTCPToWebSocket(backendConn, clientConn)
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

func (w *WebSocketProxier) relayTCPToWebSocket(src *websocket.Conn, dst net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := dst.Read(buf)
		if err != nil {
			return
		}

		if err := src.WriteMessage(websocket.BinaryMessage, buf[:n]); err != nil {
			return
		}
	}
}