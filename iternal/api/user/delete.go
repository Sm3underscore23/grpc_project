package user

import (
	"context"

	desc "project/pkg/user_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	id := req.GetId()

	err := i.userService.Delete(ctx, id)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"database error: %s",
			err,
		)
	}

	return &emptypb.Empty{}, nil
}
