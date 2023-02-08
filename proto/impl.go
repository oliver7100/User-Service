package proto

import "context"

type IService interface {
	CreateUser(context.Context, CreateUserRequest) (CreateUserResponse, error)
	GetUser(context.Context, GetUserRequest) (CreateUserResponse, error)
}
