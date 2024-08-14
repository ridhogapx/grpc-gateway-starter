package sql

import (
	"api-service/server/db/model"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *Repository) UpdateOrCreateUser(data *model.User) (*model.User, error) {

	if data.ID < 1 {

		err := r.db.Create(&data).Error
		if err != nil {
			r.log.Debug("[func: UpdateOrCreateUser] Something went wrong in query execution", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		return data, nil

	} else {

		loc, _ := time.LoadLocation("Asia/Jakarta")
		t := time.Now().In(loc)

		v := map[string]interface{}{
			"Name":        data.Name,
			"Username":    data.Username,
			"Status":      data.Status,
			"ProfilePict": data.ProfilePict,
			"AllowNSFW":   data.AllowNSFW,
			"Bio":         data.Bio,
			"SocialMedia": data.SocialMedia,
			"UpdatedAt":   t,
		}

		err := r.db.Where("id = ?", data.ID).Updates(&v).Error
		if err != nil {
			r.log.Debug("[func: UpdateOrCreateUser] Something went wrong in query execution", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "Internal Error")
		}

		return data, nil

	}

}

func (r *Repository) FindUser(v *model.User, field []string) (data []*model.User, err error) {

	query := r.db.Model(&model.User{})

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
