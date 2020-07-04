package models

import (
	"context"
)

type Dao interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByEmployeeId(ctx context.Context, employeeId int64) (User, error)
	GetUserByCredentials(ctx context.Context, email string, password string) (User, error)
}

type ApiSvc interface {
	RegisterUser(ctx context.Context, user User) error
	LoginUser(ctx context.Context, email string, password string) (User, bool, error)
}
