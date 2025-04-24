package service

import (
	"context"
	"project/iternal/model"
)

type UserService interface {
	Create(ctx context.Context, info *model.UserInfoPrivate) (int64, error)
	Get(ctx context.Context, id int64) (*model.UserInfoPublic, error)
	Update(ctx context.Context, id int64, info *model.UpdateUserInfo) error
	Delete(ctx context.Context, id int64) error
}
