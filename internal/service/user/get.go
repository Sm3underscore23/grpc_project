package user

import (
	"context"
	"project/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.UserInfoPublic, error) {
	getedUser, err := s.userRepository.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return getedUser, nil
}
