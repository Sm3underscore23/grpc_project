package user

import (
	"context"
	"project/iternal/model"
)

func (s *serv) Create(ctx context.Context, info *model.UserInfoPrivate) (int64, error) {
	userId, err := s.userRepository.Create(ctx, info)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
