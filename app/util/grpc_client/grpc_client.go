package grpc_client

import (
	"fmt"
	"go-gin-api/app/config"
	"go-gin-api/app/route/middleware/jaeger"
	"go-gin-api/app/util/jaeger_trace"
	"google.golang.org/grpc"
)

func CreateServiceListenConn() *grpc.ClientConn {
	return createGrpcClient("127.0.0.1:9901")
}

func CreateServiceSpeakConn() *grpc.ClientConn {
	return createGrpcClient("127.0.0.1:9902")
}

func CreateServiceReadConn() *grpc.ClientConn {
	return createGrpcClient("127.0.0.1:9903")
}

func CreateServiceWriteConn() *grpc.ClientConn {
	return createGrpcClient("127.0.0.1:9904")
}

func createGrpcClient(serviceAddress string ) *grpc.ClientConn {

	var conn *grpc.ClientConn
	var err error
	if config.JaegerOpen == 1 {
		conn, err = grpc.Dial(
			serviceAddress,
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(jaeger_trace.ClientInterceptor(jaeger.Tracer, jaeger.ParentSpan.Context())))
	} else {
		conn, err = grpc.Dial(
			serviceAddress,
			grpc.WithInsecure())
	}
	if err != nil {
		fmt.Println(serviceAddress, "grpc conn err:", err)
	}
	return conn
}
