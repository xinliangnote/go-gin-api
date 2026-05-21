package service

import (
	"context"

	"github.com/xinliangnote/go-gin-api/internal/grpc/pb"
)

// HealthService 健康检查服务实现
type HealthService struct {
	pb.UnimplementedHealthServiceServer
}

// NewHealthService 创建健康检查服务
func NewHealthService() *HealthService {
	return &HealthService{}
}

// Check 执行健康检查
func (s *HealthService) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}
