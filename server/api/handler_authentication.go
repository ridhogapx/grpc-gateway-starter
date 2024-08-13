package api

import (
	pb "api-service/server/proto"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GRPCService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {

	result := &pb.SignUpResponse{
		Code:    codes.OK.String(),
		Message: "Sign-up is complete",
	}

	if req.GetEmail() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Email field can't be empty.")
	}

	if req.GetUsername() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Username can't be empty.")
	}

	if req.GetName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name field can't be empty.")
	}

	return result, nil

}
