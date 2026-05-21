package interceptor

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryLogInterceptor 日志拦截器，记录每次 unary RPC 调用
func UnaryLogInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ts := time.Now()

		resp, err := handler(ctx, req)

		cost := time.Since(ts)
		fields := []zap.Field{
			zap.String("method", info.FullMethod),
			zap.Duration("cost", cost),
		}

		if err != nil {
			fields = append(fields, zap.Error(err))
			logger.Warn("grpc-request", fields...)
		} else {
			logger.Info("grpc-request", fields...)
		}

		return resp, err
	}
}

// UnaryRecoveryInterceptor panic 恢复拦截器
func UnaryRecoveryInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				stackInfo := string(debug.Stack())
				logger.Error("grpc panic recovered",
					zap.String("method", info.FullMethod),
					zap.String("panic", fmt.Sprintf("%+v", r)),
					zap.String("stack", stackInfo),
				)
				err = status.Errorf(codes.Internal, "internal server error")
			}
		}()

		return handler(ctx, req)
	}
}
