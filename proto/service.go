package proto

import (
	"context"

	"github.com/users/oliver7100/user-service/database"
)

type service struct {
	UnimplementedUserServiceServer
	Conn *database.Connection
}

func (s *service) GetUser(ctx context.Context, req *GetUserRequest) (*CreateUserResponse, error) {

	var user database.User

	if tx := s.Conn.Instance.Find(&user, "email = ?", req.GetEmail()); tx.Error != nil {
		return nil, tx.Error
	}

	return &CreateUserResponse{
		User: &User{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		},
	}, nil
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	var model database.User

	model.Name = req.User.Name
	model.Email = req.User.Email
	model.Password = req.User.Password

	if res := s.Conn.Instance.Create(&model); res.Error != nil {
		return nil, res.Error
	}

	return &CreateUserResponse{
		User: req.User,
	}, nil
}

func CreateNewService(conn *database.Connection) *service {
	return &service{
		Conn: conn,
	}
}
