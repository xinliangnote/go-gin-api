package service

import (
	"context"
	"fmt"

	"github.com/xinliangnote/go-gin-api/internal/grpc/pb"
	"github.com/xinliangnote/go-gin-api/pkg/trace"

	"google.golang.org/grpc/metadata"
)

// HelloService 示例 gRPC 服务实现
type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

// NewHelloService 创建示例服务
func NewHelloService() *HelloService {
	return &HelloService{}
}

// SayHello 问候
func (s *HelloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := req.Name
	if name == "" {
		name = "World"
	}

	// 从 metadata 中获取 trace id
	var traceId string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if vals := md.Get(trace.Header); len(vals) > 0 {
			traceId = vals[0]
		}
	}

	if traceId == "" {
		t := trace.New("")
		traceId = t.ID()
	}

	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", name),
		TraceId: traceId,
	}, nil
}
