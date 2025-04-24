package user

import (
	"context"
	"project/iternal/converter"

	desc "project/pkg/user_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	userId, err := i.userService.Create(ctx, converter.ToServiceFromUserInfoPrivate(req.GetInfo()))
	if err != nil {
		return nil,
			status.Errorf(
				codes.Internal,
				"database, error : %s",
				err,
			)
	}

	return &desc.CreateResponse{
		Id: userId,
	}, nil
}
