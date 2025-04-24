package user

import (
	"context"
	"project/internal/converter"

	desc "project/pkg/user_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	id := req.GetId()
	info := req.GetInfo()

	err := i.userService.Update(ctx, id, converter.ToServiceFromUpdateUserInfo(info))
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"database error: %s",
			err,
		)
	}

	return &emptypb.Empty{}, nil
}
