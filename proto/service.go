package proto

import (
	"context"

	"github.com/oliver7100/user-service/database"
	"github.com/oliver7100/user-service/internal"
)

type service struct {
	UnimplementedUserServiceServer
	Conn *database.Connection
}

func (s *service) CanUserLogin(ctx context.Context, req *CanUserLoginRequest) (*CanUserLoginResponse, error) {
	var user database.User

	if tx := s.Conn.Instance.First(&user, "email = ?", req.GetEmail()); tx.Error != nil {
		return nil, tx.Error
	}

	if ok, err := internal.HashCompare(req.Password, user.Password); ok {
		return &CanUserLoginResponse{
			Valid: true,
		}, nil
	} else {
		return nil, err
	}
}

func (s *service) GetUser(ctx context.Context, req *GetUserRequest) (*CreateUserResponse, error) {

	var user database.User

	if tx := s.Conn.Instance.First(&user, "email = ?", req.GetEmail()); tx.Error != nil {
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

	hashedPw, err := internal.HashPassword(req.User.Password)

	if err != nil {
		return nil, err
	}

	model.Name = req.User.Name
	model.Email = req.User.Email
	model.Password = hashedPw

	if res := s.Conn.Instance.Create(&model); res.Error != nil {
		return nil, res.Error
	}

	return &CreateUserResponse{
		User: &User{
			Name:     model.Name,
			Password: model.Password,
			Email:    model.Email,
		},
	}, nil
}

func CreateNewService(conn *database.Connection) *service {
	return &service{
		Conn: conn,
	}
}
