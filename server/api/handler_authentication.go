package api

import (
	"api-service/server/db/model"
	pb "api-service/server/proto"
	"context"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GRPCService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {

	result := &pb.SignUpResponse{
		Code:    codes.OK.String(),
		Message: "Sign-up is complete",
	}

	rulesValidator := map[string]interface{}{
		"email":    "required,email,min=5,max=90",
		"username": "required,min=4,max=50",
		"name":     "required,max=50",
	}

	errValidator := s.ValidatePayload(req, rulesValidator)
	if len(errValidator) > 0 {

		result.Code = codes.InvalidArgument.String()
		result.Message = "Invalid input"
		result.Errors = errValidator

	}

	validateUser, err := s.repos.FindUser(&model.User{Username: req.GetUsername()}, []string{"users.id", "users.username", "users.email"})
	if err != nil {
		return nil, err
	}

	if len(validateUser) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Username has been already taken. Please use other username!")
	}

	userORM := &model.User{}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), 14)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	userORM.Username = req.GetUsername()
	userORM.Name = req.GetName()
	userORM.Email = req.GetEmail()
	userORM.Password = string(hashedPassword)
	userORM.ProfilePict = req.GetProfilePicture()
	userORM.AllowNSFW = req.GetAllowNsfw()
	userORM.Status = "PENDING"

	_, err = s.repos.UpdateOrCreateUser(userORM)
	if err != nil {
		return nil, err
	}

	return result, nil

}
