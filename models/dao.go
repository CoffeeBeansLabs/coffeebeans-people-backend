package models

import (
	"context"
)

type Dao interface {
	CreateUser(ctx context.Context, user User) error
}
