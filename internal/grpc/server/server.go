package server

import (
	"context"
	"net"

	grpcInterceptor "github.com/xinliangnote/go-gin-api/internal/grpc/interceptor"
	"github.com/xinliangnote/go-gin-api/internal/grpc/pb"
	"github.com/xinliangnote/go-gin-api/internal/grpc/service"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// GRPCServer gRPC 服务器封装
type GRPCServer struct {
	server   *grpc.Server
	listener net.Listener
	logger   *zap.Logger
}

// New 创建 gRPC 服务器
func New(logger *zap.Logger, addr string) (*GRPCServer, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	// 链式拦截器：recovery -> logging
	chainedInterceptor := chainUnaryInterceptors(
		grpcInterceptor.UnaryRecoveryInterceptor(logger),
		grpcInterceptor.UnaryLogInterceptor(logger),
	)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(chainedInterceptor),
	)

	// 注册服务
	pb.RegisterHealthServiceServer(s, service.NewHealthService())
	pb.RegisterHelloServiceServer(s, service.NewHelloService())

	return &GRPCServer{
		server:   s,
		listener: lis,
		logger:   logger,
	}, nil
}

// Serve 启动 gRPC 服务（阻塞）
func (s *GRPCServer) Serve() error {
	s.logger.Info("grpc server starting", zap.String("addr", s.listener.Addr().String()))
	return s.server.Serve(s.listener)
}

// GracefulStop 优雅停止 gRPC 服务
func (s *GRPCServer) GracefulStop() {
	s.logger.Info("grpc server stopping")
	s.server.GracefulStop()
}

// chainUnaryInterceptors 将多个 unary 拦截器串联
func chainUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		chain := handler
		for i := len(interceptors) - 1; i >= 0; i-- {
			currentInterceptor := interceptors[i]
			next := chain
			chain = func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return currentInterceptor(currentCtx, currentReq, info, func(ctx context.Context, req interface{}) (interface{}, error) {
					return next(ctx, req)
				})
			}
		}
		return chain(ctx, req)
	}
}
