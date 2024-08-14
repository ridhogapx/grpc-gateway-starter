package api

import (
	"api-service/server/db/sql"
	pb "api-service/server/proto"
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type GRPCService struct {
	repos *sql.Repository
	pb.ApiServiceServer
}

func New(repos *sql.Repository) *GRPCService {
	return &GRPCService{
		repos: repos,
	}
}

func (s *GRPCService) ValidatePayload(data interface{}, rules map[string]interface{}) []*pb.Error {

	result := []*pb.Error{}

	dataMap := map[string]interface{}{}

	validate := validator.New()
	dataBytes, _ := json.Marshal(data)

	err := json.Unmarshal(dataBytes, &dataMap)
	if err != nil {
		return result
	}

	errs := validate.ValidateMap(dataMap, rules)
	if len(errs) > 0 {

		for k, v := range errs {
			errValidator := v.(validator.ValidationErrors)

			for _, el := range errValidator {
				result = append(result, &pb.Error{
					Field: k,
					Tag:   el.ActualTag(),
				})
			}
		}

	}

	return result

}
