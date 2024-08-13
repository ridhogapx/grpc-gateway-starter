package api

import (
	pb "api-service/server/proto"
	"context"

	"google.golang.org/grpc/codes"
)

type GRPCService struct {
	pb.ApiServiceServer
}

func New() *GRPCService {
	return &GRPCService{}
}

func (s *GRPCService) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {

	result := &pb.HealthCheckResponse{
		Code:    codes.OK.String(),
		Message: "API is running!",
	}

	return result, nil

}
