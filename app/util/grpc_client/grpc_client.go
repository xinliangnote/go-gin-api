package grpc_client

import (
	"context"
	"fmt"
	grpc_middeware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go-gin-api/app/config"
	"go-gin-api/app/route/middleware/jaeger"
	"go-gin-api/app/util/grpc_log"
	"go-gin-api/app/util/jaeger_trace"
	"google.golang.org/grpc"
	"time"
)

func CreateServiceListenConn() *grpc.ClientConn {
	return createGrpcConn("127.0.0.1:9901")
}

func CreateServiceSpeakConn() *grpc.ClientConn {
	return createGrpcConn("127.0.0.1:9902")
}

func CreateServiceReadConn() *grpc.ClientConn {
	return createGrpcConn("127.0.0.1:9903")
}

func CreateServiceWriteConn() *grpc.ClientConn {
	return createGrpcConn("127.0.0.1:9904")
}

func createGrpcConn(serviceAddress string) *grpc.ClientConn {

	var conn *grpc.ClientConn
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 500)
	defer cancel()

	if config.JaegerOpen == 1 {
		conn, err = grpc.DialContext(
			ctx,
			serviceAddress,
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithUnaryInterceptor(
				grpc_middeware.ChainUnaryClient(
					jaeger_trace.ClientInterceptor(jaeger.Tracer, jaeger.ParentSpan.Context()),
					grpc_log.ClientInterceptor(),
				),
			),
		)
	} else {
		conn, err = grpc.DialContext(
			ctx,
			serviceAddress,
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithUnaryInterceptor(
				grpc_middeware.ChainUnaryClient(
					grpc_log.ClientInterceptor(),
				),
			),
		)
	}

	if err != nil {
		fmt.Println(serviceAddress, "grpc conn err:", err)
	}
	return conn
}
