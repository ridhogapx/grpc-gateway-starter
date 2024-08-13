package api

import (
	pb "api-service/server/proto"
)

type GRPCService struct {
	pb.ApiServiceServer
}

func New() *GRPCService {
	return &GRPCService{}
}
