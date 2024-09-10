package sql

import (
	pb "api-service/server/proto"
	"time"

	"github.com/fatih/structs"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *Repository) UpdateOrCreateUser(data *pb.UserORM) (*pb.UserORM, error) {

	if data.Id < 1 {

		err := r.db.Create(&data).Error
		if err != nil {
			r.log.Debug("[func: UpdateOrCreateUser] Something went wrong in query execution", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		return data, nil

	} else {

		loc, _ := time.LoadLocation("Asia/Jakarta")
		t := time.Now().In(loc)

		data.UpdatedAt = &t

		dataMap := structs.Map(data)

		err := r.db.Where("id = ?", data.Id).Updates(&dataMap).Error
		if err != nil {
			r.log.Debug("[func: UpdateOrCreateUser] Something went wrong in query execution", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		return data, nil

	}

}

func (r *Repository) FindUser(v *pb.UserORM, field []string) (data []*pb.UserORM, err error) {

	query := r.db.Model(&pb.UserORM{})

	if len(field) > 0 {
		query = query.Select(field)
	}

	if v != nil {
		query = query.Where(v)
	}

	if err := query.Find(&data).Error; err != nil {
		r.log.Debug("[func: UpdateOrCreateUser] Something went wrong in query execution", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	return data, nil

}
