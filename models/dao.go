package models

import (
	"context"
)

type Dao interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByEmployeeId(ctx context.Context, employeeId int64) (User, error)
}

type ApiSvc interface {
	RegisterUser(ctx context.Context, user User) error
}