package user

import (
	"context"
	"project/internal/converter"

	desc "project/pkg/user_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	id := req.GetId()

	createdUser, err := i.userService.Get(ctx, id)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"database error: %s",
			err,
		)
	}

	return &desc.GetResponse{
		InfoPublic: converter.ToUserInfoPublicFromService(createdUser),
	}, nil
}
