package router

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/proxy"
)

func setProxyRouter(r *resource) {
	proxyHandler := proxy.NewProxy(r.logger)

	proxyHandler.AddRoute(proxy.Route{
		PathPrefix: "/api/v1/backend",
		BackendURL: "http://localhost:8081",
		Protocol:   proxy.ProtocolHTTP,
	})

	proxyHandler.AddRoute(proxy.Route{
		PathPrefix: "/proxy/socket/ws",
		BackendURL: "ws://localhost:8081",
		Protocol:   proxy.ProtocolWebSocket,
	})

	proxyGroup := r.mux.Group("/api/v1/backend", core.DisableTraceLog)
	{
		proxyGroup.Any("/*path", func(ctx core.Context) {
			proxyHandler.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
		})
	}

	socketProxyGroup := r.mux.Group("/proxy/socket", core.DisableTraceLog)
	{
		socketProxyGroup.GET("/*path", func(ctx core.Context) {
			proxyHandler.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
		})
	}
}